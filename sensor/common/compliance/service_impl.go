package compliance

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/internalapi/compliance"
	"github.com/stackrox/rox/generated/internalapi/sensor"
	"github.com/stackrox/rox/pkg/grpc/authz/idcheck"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// BenchmarkResultsService is the struct that manages the benchmark results API
type serviceImpl struct {
	output chan *compliance.ComplianceReturn
}

// RegisterServiceServer registers this service with the given gRPC Server.
func (s *serviceImpl) RegisterServiceServer(grpcServer *grpc.Server) {
	sensor.RegisterComplianceServiceServer(grpcServer, s)
}

// RegisterServiceHandler implements the APIService interface, but the agent does not accept calls over the gRPC gateway
func (s *serviceImpl) RegisterServiceHandler(context.Context, *runtime.ServeMux, *grpc.ClientConn) error {
	return nil
}

// AuthFuncOverride specifies the auth criteria for this API.
func (s *serviceImpl) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, idcheck.BenchmarkOnly().Authorized(ctx, fullMethodName)
}

// Output returns the channel where the received messages are output.
func (s *serviceImpl) Output() <-chan *compliance.ComplianceReturn {
	return s.output
}

// PushComplianceReturn takes the compliance results and outputs them to the channel.
func (s *serviceImpl) PushComplianceReturn(ctx context.Context, request *compliance.ComplianceReturn) (*v1.Empty, error) {
	// Push a message to the output channel.
	s.output <- request
	return &v1.Empty{}, nil
}
