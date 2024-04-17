package service

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	deleConnection "github.com/stackrox/rox/central/delegatedregistryconfig/util/connection"
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

func (s *serviceImpl) broadcast(ctx context.Context, msg *central.MsgToSensor) error {
	for _, conn := range s.connManager.GetActiveConnections() {
		if !deleConnection.ValidForDelegation(conn) {
			continue
		}

		err := conn.InjectMessage(ctx, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *serviceImpl) PostCollectorRuntimeConfiguration(
	ctx context.Context,
	request *v1.PostCollectorRuntimeConfigurationRequest,
) (*v1.Empty, error) {

	log.Infof("request.CollectorRuntimeConfiguration= %+v", request.CollectorRuntimeConfiguration)
	err := s.dataStore.SetRuntimeConfiguration(ctx, request.CollectorRuntimeConfiguration)

	msg := &central.MsgToSensor{
		Msg: &central.MsgToSensor_RuntimeFilteringConfiguration{
			RuntimeFilteringConfiguration: request.CollectorRuntimeConfiguration,
		},
	}

	err2 := s.broadcast(ctx, msg)

	if err2 != nil {
		return nil, err2
	}

	return &v1.Empty{}, err
}
