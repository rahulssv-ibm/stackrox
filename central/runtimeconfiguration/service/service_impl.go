package service

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	datastore "github.com/stackrox/rox/central/runtimeconfiguration/datastore"
	"github.com/stackrox/rox/central/sensor/service/connection"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/pkg/auth/permissions"
	"github.com/stackrox/rox/pkg/grpc/authz"
	"github.com/stackrox/rox/pkg/grpc/authz/perrpc"
	"github.com/stackrox/rox/pkg/grpc/authz/user"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/sac/resources"
	"google.golang.org/grpc"
)

var (
	authorizer = perrpc.FromMap(map[authz.Authorizer][]string{
		user.With(permissions.View(resources.Administration)): {
			"/v1.CollectorRuntimeConfigurationService/GetCollectorRuntimeConfiguration",
			"/v1.CollectorRuntimeConfigurationService/PostCollectorRuntimeConfiguration",
		},
	})

	log = logging.LoggerForModule()
)

type serviceImpl struct {
	dataStore   datastore.DataStore
	connManager connection.Manager
}

// RegisterServiceServer registers this service with the given gRPC Server.
func (s *serviceImpl) RegisterServiceServer(grpcServer *grpc.Server) {
	v1.RegisterCollectorRuntimeConfigurationServiceServer(grpcServer, s)
}

// RegisterServiceHandler registers this service with the given gRPC Gateway endpoint.
func (s *serviceImpl) RegisterServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return v1.RegisterCollectorRuntimeConfigurationServiceHandler(ctx, mux, conn)
}

// AuthFuncOverride specifies the auth criteria for this API.
func (s *serviceImpl) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, authorizer.Authorized(ctx, fullMethodName)
}

// GetCollectorRuntimeConfiguration returns the runtime configuration for collector
func (s *serviceImpl) GetCollectorRuntimeConfiguration(
	ctx context.Context, _ *v1.Empty,
) (*v1.GetCollectorRuntimeConfigurationResponse, error) {

	runtimeFilteringConfiguration, err := s.dataStore.GetRuntimeConfiguration(ctx)

	getCollectorRuntimeConfigurationResponse := v1.GetCollectorRuntimeConfigurationResponse{
		CollectorRuntimeConfiguration: runtimeFilteringConfiguration,
	}

	return &getCollectorRuntimeConfigurationResponse, err
}

func (s *serviceImpl) PostCollectorRuntimeConfiguration(
	ctx context.Context,
	request *v1.PostCollectorRuntimeConfigurationRequest,
) (*v1.Empty, error) {

	log.Infof("request.CollectorRuntimeConfiguration= %+v", request.CollectorRuntimeConfiguration)
	runtimeConfig := request.CollectorRuntimeConfiguration
	runtimeFilters := runtimeConfig.RuntimeFilters
	for _, runtimeFilter := range runtimeFilters {
		log.Infof("runtimeFilter= %+v", runtimeFilter)
		log.Infof("len(runtimeFilter.rules)= %+v", len(runtimeFilter.Rules))
		for _, rule := range runtimeFilter.Rules {
			log.Infof("rule= %+v", rule)
		}
	}
	collections := runtimeConfig.ResourceCollections
	for _, collection := range collections {
		log.Infof("collection= %+v", collection)
		for _, resourceSelector := range collection.ResourceSelectors {
			log.Infof("resourceSelector= %+v", resourceSelector)
			for _, rule := range resourceSelector.Rules {
				log.Infof("rule= %+v", rule)
			}
		}
	}

	err := s.dataStore.SetRuntimeConfiguration(ctx, request.CollectorRuntimeConfiguration)

	msg := &central.MsgToSensor{
		Msg: &central.MsgToSensor_RuntimeFilteringConfiguration{
			RuntimeFilteringConfiguration: request.CollectorRuntimeConfiguration,
		},
	}

	s.connManager.BroadcastMessage(msg)

	return &v1.Empty{}, err
}
