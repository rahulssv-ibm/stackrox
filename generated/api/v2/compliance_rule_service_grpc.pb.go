// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v2/compliance_rule_service.proto

package v2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ComplianceRuleService_GetComplianceRule_FullMethodName = "/v2.ComplianceRuleService/GetComplianceRule"
)

// ComplianceRuleServiceClient is the client API for ComplianceRuleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComplianceRuleServiceClient interface {
	// GetComplianceRule returns rule matching given request
	GetComplianceRule(ctx context.Context, in *RuleRequest, opts ...grpc.CallOption) (*ComplianceRule, error)
}

type complianceRuleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComplianceRuleServiceClient(cc grpc.ClientConnInterface) ComplianceRuleServiceClient {
	return &complianceRuleServiceClient{cc}
}

func (c *complianceRuleServiceClient) GetComplianceRule(ctx context.Context, in *RuleRequest, opts ...grpc.CallOption) (*ComplianceRule, error) {
	out := new(ComplianceRule)
	err := c.cc.Invoke(ctx, ComplianceRuleService_GetComplianceRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComplianceRuleServiceServer is the server API for ComplianceRuleService service.
// All implementations should embed UnimplementedComplianceRuleServiceServer
// for forward compatibility
type ComplianceRuleServiceServer interface {
	// GetComplianceRule returns rule matching given request
	GetComplianceRule(context.Context, *RuleRequest) (*ComplianceRule, error)
}

// UnimplementedComplianceRuleServiceServer should be embedded to have forward compatible implementations.
type UnimplementedComplianceRuleServiceServer struct {
}

func (UnimplementedComplianceRuleServiceServer) GetComplianceRule(context.Context, *RuleRequest) (*ComplianceRule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComplianceRule not implemented")
}

// UnsafeComplianceRuleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComplianceRuleServiceServer will
// result in compilation errors.
type UnsafeComplianceRuleServiceServer interface {
	mustEmbedUnimplementedComplianceRuleServiceServer()
}

func RegisterComplianceRuleServiceServer(s grpc.ServiceRegistrar, srv ComplianceRuleServiceServer) {
	s.RegisterService(&ComplianceRuleService_ServiceDesc, srv)
}

func _ComplianceRuleService_GetComplianceRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceRuleServiceServer).GetComplianceRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComplianceRuleService_GetComplianceRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceRuleServiceServer).GetComplianceRule(ctx, req.(*RuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ComplianceRuleService_ServiceDesc is the grpc.ServiceDesc for ComplianceRuleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComplianceRuleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v2.ComplianceRuleService",
	HandlerType: (*ComplianceRuleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetComplianceRule",
			Handler:    _ComplianceRuleService_GetComplianceRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v2/compliance_rule_service.proto",
}
