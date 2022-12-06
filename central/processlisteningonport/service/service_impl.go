package service

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	datastore "github.com/stackrox/rox/central/processlisteningonport/datastore"
	"github.com/stackrox/rox/central/role/resources"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/auth/permissions"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/grpc/authz"
	"github.com/stackrox/rox/pkg/grpc/authz/perrpc"
	"github.com/stackrox/rox/pkg/grpc/authz/user"
	"google.golang.org/grpc"
	// "github.com/stackrox/rox/pkg/grpc/authz/allow"
)

var (
	authorizer = perrpc.FromMap(map[authz.Authorizer][]string{
		user.With(permissions.View(resources.ProcessListeningOnPort)): {
			"/v1.ProcessesListeningOnPortsService/GetProcessesListeningOnPortsByNamespace",
			"/v1.ProcessesListeningOnPortsService/GetProcessesListeningOnPortsByNamespaceAndDeployment",
		},
	})
)

type serviceImpl struct {
	dataStore datastore.DataStore
}

// RegisterServiceServer registers this service with the given gRPC Server.
func (s *serviceImpl) RegisterServiceServer(grpcServer *grpc.Server) {
	v1.RegisterProcessesListeningOnPortsServiceServer(grpcServer, s)
}

// RegisterServiceHandler registers this service with the given gRPC Gateway endpoint.
func (s *serviceImpl) RegisterServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return v1.RegisterProcessesListeningOnPortsServiceHandler(ctx, mux, conn)
}

// AuthFuncOverride specifies the auth criteria for this API.
func (s *serviceImpl) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	//// XXX: An anonymous access only for development
	//  return ctx, allow.Anonymous().Authorized(ctx, fullMethodName)
	return ctx, authorizer.Authorized(ctx, fullMethodName)
}

func emptyProcessesListeningOnPortsWithDeploymentResponse() (*v1.GetProcessesListeningOnPortsWithDeploymentResponse, error) {
	result := &v1.GetProcessesListeningOnPortsWithDeploymentResponse{
		ProcessesListeningOnPortsWithDeployment: make([]*v1.ProcessListeningOnPortWithDeploymentId, 0),
	}
	return result, nil
}

func (s *serviceImpl) GetProcessesListeningOnPortsByNamespace(ctx context.Context, req *v1.GetProcessesListeningOnPortsByNamespaceRequest) (*v1.GetProcessesListeningOnPortsWithDeploymentResponse, error) {

	if !env.PostgresDatastoreEnabled.BooleanSetting() {
		// PLOP is a Postgres-only feature, do nothing.
		log.Warnf("Tried to request PLOP not on Postgres, ignore: %s",
			req.GetNamespace())
		return emptyProcessesListeningOnPortsWithDeploymentResponse()
	}

	namespace := req.GetNamespace()
	processesListeningOnPorts, err := s.dataStore.GetProcessListeningOnPort(
		ctx, datastore.GetOptions{Namespace: &namespace})

	if err != nil {
		log.Warnf("In processlisteningonport service query return err: %+v", err)
		return emptyProcessesListeningOnPortsWithDeploymentResponse()
	}

	if processesListeningOnPorts == nil {
		log.Debug("In processlisteningonport service query return nil")
		return emptyProcessesListeningOnPortsWithDeploymentResponse()
	}

	size := len(processesListeningOnPorts)
	result := make([]*v1.ProcessListeningOnPortWithDeploymentId, size)

	for k, v := range processesListeningOnPorts {
		plop := &v1.ProcessListeningOnPortWithDeploymentId{
			DeploymentId:              k,
			ProcessesListeningOnPorts: v,
		}
		result = append(result, plop)
	}

	return &v1.GetProcessesListeningOnPortsWithDeploymentResponse{
		ProcessesListeningOnPortsWithDeployment: result,
	}, err
}

func emptyProcessesListeningOnPortsResponse() (*v1.GetProcessesListeningOnPortsResponse, error) {
	result := &v1.GetProcessesListeningOnPortsResponse{
		ProcessesListeningOnPorts: make([]*storage.ProcessListeningOnPort, 0),
	}
	return result, nil
}

func (s *serviceImpl) GetProcessesListeningOnPortsByNamespaceAndDeployment(
	ctx context.Context,
	req *v1.GetProcessesListeningOnPortsByNamespaceAndDeploymentRequest,
) (*v1.GetProcessesListeningOnPortsResponse, error) {

	if !env.PostgresDatastoreEnabled.BooleanSetting() {
		// PLOP is a Postgres-only feature, do nothing.
		log.Warnf("Tried to request PLOP not on Postgres, ignore: %s, %s",
			req.GetNamespace(), req.GetDeploymentId())
		return emptyProcessesListeningOnPortsResponse()
	}

	namespace := req.GetNamespace()
	deployment := req.GetDeploymentId()
	processesListeningOnPorts, err := s.dataStore.GetProcessListeningOnPort(
		ctx, datastore.GetOptions{
			Namespace:    &namespace,
			DeploymentID: &deployment,
		})
	log.Info("In processlisteningonport service got processes")

	if err != nil {
		log.Warnf("In processlisteningonport service query return err: %+v", err)
		return emptyProcessesListeningOnPortsResponse()
	}

	if processesListeningOnPorts == nil {
		log.Debug("In processlisteningonport service query return nil")
		return emptyProcessesListeningOnPortsResponse()
	}

	size := len(processesListeningOnPorts)
	result := make([]*storage.ProcessListeningOnPort, size)

	// Storage returns map DeploymentID -> PLOP. Just in case verify that
	// deployment id matches.
	for k, v := range processesListeningOnPorts {
		if k != deployment {
			log.Warnf("Requested deployment %s, got %s. Skipping", deployment, k)
		} else {
			result = append(result, v...)
		}
	}

	return &v1.GetProcessesListeningOnPortsResponse{
		ProcessesListeningOnPorts: result,
	}, err
}
