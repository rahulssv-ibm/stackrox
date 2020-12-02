package service

import (
	"context"
	"fmt"

	timestamp "github.com/gogo/protobuf/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/detection"
	healthDatastore "github.com/stackrox/rox/central/integrationhealth/datastore"
	"github.com/stackrox/rox/central/notifier/datastore"
	"github.com/stackrox/rox/central/notifier/processor"
	"github.com/stackrox/rox/central/notifiers"
	"github.com/stackrox/rox/central/role/resources"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/auth/permissions"
	"github.com/stackrox/rox/pkg/endpoints"
	"github.com/stackrox/rox/pkg/errorhelpers"
	"github.com/stackrox/rox/pkg/grpc/authz"
	"github.com/stackrox/rox/pkg/grpc/authz/perrpc"
	"github.com/stackrox/rox/pkg/grpc/authz/user"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/secrets"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	authorizer = perrpc.FromMap(map[authz.Authorizer][]string{
		user.With(permissions.View(resources.Notifier)): {
			"/v1.NotifierService/GetNotifier",
			"/v1.NotifierService/GetNotifiers",
		},
		user.With(permissions.Modify(resources.Notifier)): {
			"/v1.NotifierService/PutNotifier",
			"/v1.NotifierService/PostNotifier",
			"/v1.NotifierService/TestNotifier",
			"/v1.NotifierService/DeleteNotifier",
			"/v1.NotifierService/TestUpdatedNotifier",
			"/v1.NotifierService/UpdateNotifier",
		},
	})
)

// ClusterService is the struct that manages the cluster API
type serviceImpl struct {
	storage         datastore.DataStore
	processor       processor.Processor
	healthDatastore healthDatastore.DataStore

	buildTimePolicies  detection.PolicySet
	deployTimePolicies detection.PolicySet
	runTimePolicies    detection.PolicySet
}

// RegisterServiceServer registers this service with the given gRPC Server.
func (s *serviceImpl) RegisterServiceServer(grpcServer *grpc.Server) {
	v1.RegisterNotifierServiceServer(grpcServer, s)
}

// RegisterServiceHandler registers this service with the given gRPC Gateway endpoint.
func (s *serviceImpl) RegisterServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return v1.RegisterNotifierServiceHandler(ctx, mux, conn)
}

// AuthFuncOverride specifies the auth criteria for this API.
func (s *serviceImpl) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, authorizer.Authorized(ctx, fullMethodName)
}

// GetNotifier retrieves all registries that matches the request filters
func (s *serviceImpl) GetNotifier(ctx context.Context, request *v1.ResourceByID) (*storage.Notifier, error) {
	if request.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "notifier id must be provided")
	}
	notifier, exists, err := s.storage.GetNotifier(ctx, request.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !exists {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("notifier %v not found", request.GetId()))
	}
	secrets.ScrubSecretsFromStructWithReplacement(notifier, secrets.ScrubReplacementStr)
	return notifier, nil
}

// GetNotifiers retrieves all notifiers that match the request filters
func (s *serviceImpl) GetNotifiers(ctx context.Context, request *v1.GetNotifiersRequest) (*v1.GetNotifiersResponse, error) {
	notifiers, err := s.storage.GetNotifiers(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	for _, n := range notifiers {
		secrets.ScrubSecretsFromStructWithReplacement(n, secrets.ScrubReplacementStr)
	}
	return &v1.GetNotifiersResponse{Notifiers: notifiers}, nil
}

func validateNotifier(notifier *storage.Notifier) error {
	if notifier == nil {
		return errors.New("empty notifier")
	}
	errorList := errorhelpers.NewErrorList("Validation")
	if notifier.GetName() == "" {
		errorList.AddString("notifier name must be defined")
	}
	if notifier.GetType() == "" {
		errorList.AddString("notifier type must be defined")
	}
	if notifier.GetUiEndpoint() == "" {
		errorList.AddString("notifier UI endpoint must be defined")
	}
	if err := endpoints.ValidateEndpoints(notifier.Config); err != nil {
		errorList.AddWrap(err, "invalid endpoint")
	}
	return errorList.ToError()
}

// PutNotifier updates a notifier configuration, without stored credential reconciliation
func (s *serviceImpl) PutNotifier(ctx context.Context, notifier *storage.Notifier) (*v1.Empty, error) {
	return s.UpdateNotifier(ctx, &v1.UpdateNotifierRequest{Notifier: notifier, UpdatePassword: true})
}

// UpdateNotifier updates a notifier configuration
func (s *serviceImpl) UpdateNotifier(ctx context.Context, request *v1.UpdateNotifierRequest) (*v1.Empty, error) {
	if err := validateNotifier(request.GetNotifier()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.reconcileUpdateNotifierRequest(ctx, request); err != nil {
		return nil, err
	}
	notifierCreator, ok := notifiers.Registry[request.GetNotifier().GetType()]
	if !ok {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("notifier type %v is not a valid notifier type", request.GetNotifier().GetType()))
	}
	notifier, err := notifierCreator(request.GetNotifier())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err := s.storage.UpdateNotifier(ctx, request.GetNotifier()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	s.processor.UpdateNotifier(ctx, notifier)
	currentHealth, found, err := s.healthDatastore.GetIntegrationHealth(ctx, request.GetNotifier().Id)

	if !found || err != nil {
		errMsg := fmt.Sprintf("unable to get integration health for integration %s.", request.GetNotifier().Id)
		if err != nil {
			errMsg = fmt.Sprintf("%s Error: %+v", errMsg, err)
		}
		return nil, status.Error(codes.Internal, errMsg)
	}

	newHealth := currentHealth.Clone()
	newHealth.Name = request.GetNotifier().GetName()
	if err := s.healthDatastore.UpdateIntegrationHealth(ctx, newHealth); err != nil {
		return nil, err
	}
	return &v1.Empty{}, nil
}

// PostNotifier inserts a new registry into the system if it doesn't already exist
func (s *serviceImpl) PostNotifier(ctx context.Context, request *storage.Notifier) (*storage.Notifier, error) {
	if err := validateNotifier(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if request.GetId() != "" {
		return nil, status.Error(codes.InvalidArgument, "id field should be empty when posting a new notifier")
	}
	notifier, err := notifiers.CreateNotifier(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	id, err := s.storage.AddNotifier(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	request.Id = id
	s.processor.UpdateNotifier(ctx, notifier)

	if err := s.healthDatastore.UpdateIntegrationHealth(ctx, &storage.IntegrationHealth{
		Id:            request.Id,
		Name:          request.Name,
		Type:          storage.IntegrationHealth_NOTIFIER,
		Status:        storage.IntegrationHealth_UNINITIALIZED,
		ErrorMessage:  "",
		LastTimestamp: timestamp.TimestampNow(),
	}); err != nil {
		return nil, err
	}

	return request, nil
}

// TestNotifier tests to see if the config is setup properly, without stored credential reconciliation
func (s *serviceImpl) TestNotifier(ctx context.Context, notifier *storage.Notifier) (*v1.Empty, error) {
	return s.TestUpdatedNotifier(ctx, &v1.UpdateNotifierRequest{Notifier: notifier, UpdatePassword: true})
}

// TestUpdatedNotifier tests to see if the config is setup properly
func (s *serviceImpl) TestUpdatedNotifier(ctx context.Context, request *v1.UpdateNotifierRequest) (*v1.Empty, error) {
	if err := validateNotifier(request.GetNotifier()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.reconcileUpdateNotifierRequest(ctx, request); err != nil {
		return nil, err
	}
	notifier, err := notifiers.CreateNotifier(request.GetNotifier())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	defer func() {
		if err := notifier.Close(ctx); err != nil {
			log.Warn("failed to close temporary notifier instance", logging.Err(err))
		}
	}()

	if err := notifier.Test(ctx); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &v1.Empty{}, nil
}

// DeleteNotifier deletes a notifier from the system
func (s *serviceImpl) DeleteNotifier(ctx context.Context, request *v1.DeleteNotifierRequest) (*v1.Empty, error) {
	if request.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "notifier id must be provided")
	}

	n, err := s.GetNotifier(ctx, &v1.ResourceByID{Id: request.GetId()})
	if err != nil {
		return nil, err
	}

	err = s.deleteNotifiersFromPolicies(n.GetId())
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.FailedPrecondition, fmt.Sprintf("notifier is still in use by policies. Error: %s", err))
	}

	if err := s.storage.RemoveNotifier(ctx, request.GetId()); err != nil {
		return nil, err
	}

	s.processor.RemoveNotifier(ctx, request.GetId())
	if err := s.healthDatastore.RemoveIntegrationHealth(ctx, request.GetId()); err != nil {
		return nil, err
	}
	return &v1.Empty{}, nil
}

func (s *serviceImpl) deleteNotifiersFromPolicies(notifierID string) error {

	err := s.buildTimePolicies.RemoveNotifier(notifierID)
	if err != nil {
		return err
	}

	err = s.deployTimePolicies.RemoveNotifier(notifierID)
	if err != nil {
		return err
	}

	err = s.runTimePolicies.RemoveNotifier(notifierID)
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceImpl) reconcileUpdateNotifierRequest(ctx context.Context, updateRequest *v1.UpdateNotifierRequest) error {
	if updateRequest.GetUpdatePassword() {
		return nil
	}
	if updateRequest.GetNotifier() == nil {
		return status.Error(codes.InvalidArgument, "request is missing notifier config")
	}
	if updateRequest.GetNotifier().GetId() == "" {
		return status.Error(codes.InvalidArgument, "id required for stored credential reconciliation")
	}
	existingNotifierConfig, exists, err := s.storage.GetNotifier(ctx, updateRequest.GetNotifier().GetId())
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	if !exists {
		return status.Errorf(codes.NotFound, "notifier integration %s not found", updateRequest.GetNotifier().GetId())
	}
	if err := reconcileNotifierConfigWithExisting(updateRequest.GetNotifier(), existingNotifierConfig); err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}

func reconcileNotifierConfigWithExisting(updated, existing *storage.Notifier) error {
	if updated.GetConfig() == nil {
		return errors.New("the request doesn't have a valid notifier config")
	}
	return secrets.ReconcileScrubbedStructWithExisting(updated, existing)
}
