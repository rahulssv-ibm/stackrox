// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: internalapi/scanner/v4/matcher_service.proto

package v4

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Matcher_GetVulnerabilities_FullMethodName = "/scanner.v4.Matcher/GetVulnerabilities"
	Matcher_GetMetadata_FullMethodName        = "/scanner.v4.Matcher/GetMetadata"
)

// MatcherClient is the client API for Matcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Matcher finds vulnerabilities in index reports.
type MatcherClient interface {
	// GetVulnerabilities returns a VulnerabilityReport for a previously indexed manifest.
	GetVulnerabilities(ctx context.Context, in *GetVulnerabilitiesRequest, opts ...grpc.CallOption) (*VulnerabilityReport, error)
	// GetMetadata returns information on vulnerability metadata, ek.g., last update timestamp.
	GetMetadata(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Metadata, error)
}

type matcherClient struct {
	cc grpc.ClientConnInterface
}

func NewMatcherClient(cc grpc.ClientConnInterface) MatcherClient {
	return &matcherClient{cc}
}

func (c *matcherClient) GetVulnerabilities(ctx context.Context, in *GetVulnerabilitiesRequest, opts ...grpc.CallOption) (*VulnerabilityReport, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VulnerabilityReport)
	err := c.cc.Invoke(ctx, Matcher_GetVulnerabilities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matcherClient) GetMetadata(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Metadata, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Metadata)
	err := c.cc.Invoke(ctx, Matcher_GetMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatcherServer is the server API for Matcher service.
// All implementations should embed UnimplementedMatcherServer
// for forward compatibility.
//
// Matcher finds vulnerabilities in index reports.
type MatcherServer interface {
	// GetVulnerabilities returns a VulnerabilityReport for a previously indexed manifest.
	GetVulnerabilities(context.Context, *GetVulnerabilitiesRequest) (*VulnerabilityReport, error)
	// GetMetadata returns information on vulnerability metadata, ek.g., last update timestamp.
	GetMetadata(context.Context, *emptypb.Empty) (*Metadata, error)
}

// UnimplementedMatcherServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMatcherServer struct{}

func (UnimplementedMatcherServer) GetVulnerabilities(context.Context, *GetVulnerabilitiesRequest) (*VulnerabilityReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVulnerabilities not implemented")
}
func (UnimplementedMatcherServer) GetMetadata(context.Context, *emptypb.Empty) (*Metadata, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadata not implemented")
}
func (UnimplementedMatcherServer) testEmbeddedByValue() {}

// UnsafeMatcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatcherServer will
// result in compilation errors.
type UnsafeMatcherServer interface {
	mustEmbedUnimplementedMatcherServer()
}

func RegisterMatcherServer(s grpc.ServiceRegistrar, srv MatcherServer) {
	// If the following call pancis, it indicates UnimplementedMatcherServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Matcher_ServiceDesc, srv)
}

func _Matcher_GetVulnerabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVulnerabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatcherServer).GetVulnerabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Matcher_GetVulnerabilities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatcherServer).GetVulnerabilities(ctx, req.(*GetVulnerabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Matcher_GetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatcherServer).GetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Matcher_GetMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatcherServer).GetMetadata(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Matcher_ServiceDesc is the grpc.ServiceDesc for Matcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Matcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scanner.v4.Matcher",
	HandlerType: (*MatcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVulnerabilities",
			Handler:    _Matcher_GetVulnerabilities_Handler,
		},
		{
			MethodName: "GetMetadata",
			Handler:    _Matcher_GetMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internalapi/scanner/v4/matcher_service.proto",
}
