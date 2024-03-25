// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: internalapi/central/development_service.proto

// This is an internal service which contains tools intended to be used only for testing.
// It will NOT be available in Central in production builds.

package central

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
	DevelopmentService_ReplicateImage_FullMethodName               = "/central.DevelopmentService/ReplicateImage"
	DevelopmentService_URLHasValidCert_FullMethodName              = "/central.DevelopmentService/URLHasValidCert"
	DevelopmentService_RandomData_FullMethodName                   = "/central.DevelopmentService/RandomData"
	DevelopmentService_EnvVars_FullMethodName                      = "/central.DevelopmentService/EnvVars"
	DevelopmentService_ReconciliationStatsByCluster_FullMethodName = "/central.DevelopmentService/ReconciliationStatsByCluster"
)

// DevelopmentServiceClient is the client API for DevelopmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DevelopmentServiceClient interface {
	ReplicateImage(ctx context.Context, in *ReplicateImageRequest, opts ...grpc.CallOption) (*Empty, error)
	URLHasValidCert(ctx context.Context, in *URLHasValidCertRequest, opts ...grpc.CallOption) (*URLHasValidCertResponse, error)
	RandomData(ctx context.Context, in *RandomDataRequest, opts ...grpc.CallOption) (*RandomDataResponse, error)
	EnvVars(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EnvVarsResponse, error)
	ReconciliationStatsByCluster(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReconciliationStatsByClusterResponse, error)
}

type developmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDevelopmentServiceClient(cc grpc.ClientConnInterface) DevelopmentServiceClient {
	return &developmentServiceClient{cc}
}

func (c *developmentServiceClient) ReplicateImage(ctx context.Context, in *ReplicateImageRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, DevelopmentService_ReplicateImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developmentServiceClient) URLHasValidCert(ctx context.Context, in *URLHasValidCertRequest, opts ...grpc.CallOption) (*URLHasValidCertResponse, error) {
	out := new(URLHasValidCertResponse)
	err := c.cc.Invoke(ctx, DevelopmentService_URLHasValidCert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developmentServiceClient) RandomData(ctx context.Context, in *RandomDataRequest, opts ...grpc.CallOption) (*RandomDataResponse, error) {
	out := new(RandomDataResponse)
	err := c.cc.Invoke(ctx, DevelopmentService_RandomData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developmentServiceClient) EnvVars(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EnvVarsResponse, error) {
	out := new(EnvVarsResponse)
	err := c.cc.Invoke(ctx, DevelopmentService_EnvVars_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developmentServiceClient) ReconciliationStatsByCluster(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReconciliationStatsByClusterResponse, error) {
	out := new(ReconciliationStatsByClusterResponse)
	err := c.cc.Invoke(ctx, DevelopmentService_ReconciliationStatsByCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevelopmentServiceServer is the server API for DevelopmentService service.
// All implementations must embed UnimplementedDevelopmentServiceServer
// for forward compatibility
type DevelopmentServiceServer interface {
	ReplicateImage(context.Context, *ReplicateImageRequest) (*Empty, error)
	URLHasValidCert(context.Context, *URLHasValidCertRequest) (*URLHasValidCertResponse, error)
	RandomData(context.Context, *RandomDataRequest) (*RandomDataResponse, error)
	EnvVars(context.Context, *Empty) (*EnvVarsResponse, error)
	ReconciliationStatsByCluster(context.Context, *Empty) (*ReconciliationStatsByClusterResponse, error)
	mustEmbedUnimplementedDevelopmentServiceServer()
}

// UnimplementedDevelopmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDevelopmentServiceServer struct {
}

func (UnimplementedDevelopmentServiceServer) ReplicateImage(context.Context, *ReplicateImageRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplicateImage not implemented")
}
func (UnimplementedDevelopmentServiceServer) URLHasValidCert(context.Context, *URLHasValidCertRequest) (*URLHasValidCertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method URLHasValidCert not implemented")
}
func (UnimplementedDevelopmentServiceServer) RandomData(context.Context, *RandomDataRequest) (*RandomDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RandomData not implemented")
}
func (UnimplementedDevelopmentServiceServer) EnvVars(context.Context, *Empty) (*EnvVarsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnvVars not implemented")
}
func (UnimplementedDevelopmentServiceServer) ReconciliationStatsByCluster(context.Context, *Empty) (*ReconciliationStatsByClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReconciliationStatsByCluster not implemented")
}
func (UnimplementedDevelopmentServiceServer) mustEmbedUnimplementedDevelopmentServiceServer() {}

// UnsafeDevelopmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DevelopmentServiceServer will
// result in compilation errors.
type UnsafeDevelopmentServiceServer interface {
	mustEmbedUnimplementedDevelopmentServiceServer()
}

func RegisterDevelopmentServiceServer(s grpc.ServiceRegistrar, srv DevelopmentServiceServer) {
	s.RegisterService(&DevelopmentService_ServiceDesc, srv)
}

func _DevelopmentService_ReplicateImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplicateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevelopmentServiceServer).ReplicateImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DevelopmentService_ReplicateImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevelopmentServiceServer).ReplicateImage(ctx, req.(*ReplicateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevelopmentService_URLHasValidCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URLHasValidCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevelopmentServiceServer).URLHasValidCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DevelopmentService_URLHasValidCert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevelopmentServiceServer).URLHasValidCert(ctx, req.(*URLHasValidCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevelopmentService_RandomData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevelopmentServiceServer).RandomData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DevelopmentService_RandomData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevelopmentServiceServer).RandomData(ctx, req.(*RandomDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevelopmentService_EnvVars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevelopmentServiceServer).EnvVars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DevelopmentService_EnvVars_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevelopmentServiceServer).EnvVars(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevelopmentService_ReconciliationStatsByCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevelopmentServiceServer).ReconciliationStatsByCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DevelopmentService_ReconciliationStatsByCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevelopmentServiceServer).ReconciliationStatsByCluster(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DevelopmentService_ServiceDesc is the grpc.ServiceDesc for DevelopmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DevelopmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "central.DevelopmentService",
	HandlerType: (*DevelopmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReplicateImage",
			Handler:    _DevelopmentService_ReplicateImage_Handler,
		},
		{
			MethodName: "URLHasValidCert",
			Handler:    _DevelopmentService_URLHasValidCert_Handler,
		},
		{
			MethodName: "RandomData",
			Handler:    _DevelopmentService_RandomData_Handler,
		},
		{
			MethodName: "EnvVars",
			Handler:    _DevelopmentService_EnvVars_Handler,
		},
		{
			MethodName: "ReconciliationStatsByCluster",
			Handler:    _DevelopmentService_ReconciliationStatsByCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internalapi/central/development_service.proto",
}
