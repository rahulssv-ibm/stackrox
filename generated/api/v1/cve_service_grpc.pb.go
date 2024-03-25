// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v1/cve_service.proto

package v1

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
	ImageCVEService_SuppressCVEs_FullMethodName   = "/v1.ImageCVEService/SuppressCVEs"
	ImageCVEService_UnsuppressCVEs_FullMethodName = "/v1.ImageCVEService/UnsuppressCVEs"
)

// ImageCVEServiceClient is the client API for ImageCVEService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageCVEServiceClient interface {
	// SuppressCVE suppresses image cves.
	SuppressCVEs(ctx context.Context, in *SuppressCVERequest, opts ...grpc.CallOption) (*Empty, error)
	// UnsuppressCVE unsuppresses image cves.
	UnsuppressCVEs(ctx context.Context, in *UnsuppressCVERequest, opts ...grpc.CallOption) (*Empty, error)
}

type imageCVEServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageCVEServiceClient(cc grpc.ClientConnInterface) ImageCVEServiceClient {
	return &imageCVEServiceClient{cc}
}

func (c *imageCVEServiceClient) SuppressCVEs(ctx context.Context, in *SuppressCVERequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ImageCVEService_SuppressCVEs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageCVEServiceClient) UnsuppressCVEs(ctx context.Context, in *UnsuppressCVERequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ImageCVEService_UnsuppressCVEs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageCVEServiceServer is the server API for ImageCVEService service.
// All implementations must embed UnimplementedImageCVEServiceServer
// for forward compatibility
type ImageCVEServiceServer interface {
	// SuppressCVE suppresses image cves.
	SuppressCVEs(context.Context, *SuppressCVERequest) (*Empty, error)
	// UnsuppressCVE unsuppresses image cves.
	UnsuppressCVEs(context.Context, *UnsuppressCVERequest) (*Empty, error)
	mustEmbedUnimplementedImageCVEServiceServer()
}

// UnimplementedImageCVEServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImageCVEServiceServer struct {
}

func (UnimplementedImageCVEServiceServer) SuppressCVEs(context.Context, *SuppressCVERequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuppressCVEs not implemented")
}
func (UnimplementedImageCVEServiceServer) UnsuppressCVEs(context.Context, *UnsuppressCVERequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsuppressCVEs not implemented")
}
func (UnimplementedImageCVEServiceServer) mustEmbedUnimplementedImageCVEServiceServer() {}

// UnsafeImageCVEServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageCVEServiceServer will
// result in compilation errors.
type UnsafeImageCVEServiceServer interface {
	mustEmbedUnimplementedImageCVEServiceServer()
}

func RegisterImageCVEServiceServer(s grpc.ServiceRegistrar, srv ImageCVEServiceServer) {
	s.RegisterService(&ImageCVEService_ServiceDesc, srv)
}

func _ImageCVEService_SuppressCVEs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuppressCVERequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageCVEServiceServer).SuppressCVEs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageCVEService_SuppressCVEs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageCVEServiceServer).SuppressCVEs(ctx, req.(*SuppressCVERequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageCVEService_UnsuppressCVEs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsuppressCVERequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageCVEServiceServer).UnsuppressCVEs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageCVEService_UnsuppressCVEs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageCVEServiceServer).UnsuppressCVEs(ctx, req.(*UnsuppressCVERequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageCVEService_ServiceDesc is the grpc.ServiceDesc for ImageCVEService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageCVEService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ImageCVEService",
	HandlerType: (*ImageCVEServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuppressCVEs",
			Handler:    _ImageCVEService_SuppressCVEs_Handler,
		},
		{
			MethodName: "UnsuppressCVEs",
			Handler:    _ImageCVEService_UnsuppressCVEs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/cve_service.proto",
}

const (
	NodeCVEService_SuppressCVEs_FullMethodName   = "/v1.NodeCVEService/SuppressCVEs"
	NodeCVEService_UnsuppressCVEs_FullMethodName = "/v1.NodeCVEService/UnsuppressCVEs"
)

// NodeCVEServiceClient is the client API for NodeCVEService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeCVEServiceClient interface {
	// SuppressCVE suppresses node cves.
	SuppressCVEs(ctx context.Context, in *SuppressCVERequest, opts ...grpc.CallOption) (*Empty, error)
	// UnsuppressCVE unsuppresses node cves.
	UnsuppressCVEs(ctx context.Context, in *UnsuppressCVERequest, opts ...grpc.CallOption) (*Empty, error)
}

type nodeCVEServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeCVEServiceClient(cc grpc.ClientConnInterface) NodeCVEServiceClient {
	return &nodeCVEServiceClient{cc}
}

func (c *nodeCVEServiceClient) SuppressCVEs(ctx context.Context, in *SuppressCVERequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, NodeCVEService_SuppressCVEs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCVEServiceClient) UnsuppressCVEs(ctx context.Context, in *UnsuppressCVERequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, NodeCVEService_UnsuppressCVEs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeCVEServiceServer is the server API for NodeCVEService service.
// All implementations must embed UnimplementedNodeCVEServiceServer
// for forward compatibility
type NodeCVEServiceServer interface {
	// SuppressCVE suppresses node cves.
	SuppressCVEs(context.Context, *SuppressCVERequest) (*Empty, error)
	// UnsuppressCVE unsuppresses node cves.
	UnsuppressCVEs(context.Context, *UnsuppressCVERequest) (*Empty, error)
	mustEmbedUnimplementedNodeCVEServiceServer()
}

// UnimplementedNodeCVEServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeCVEServiceServer struct {
}

func (UnimplementedNodeCVEServiceServer) SuppressCVEs(context.Context, *SuppressCVERequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuppressCVEs not implemented")
}
func (UnimplementedNodeCVEServiceServer) UnsuppressCVEs(context.Context, *UnsuppressCVERequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsuppressCVEs not implemented")
}
func (UnimplementedNodeCVEServiceServer) mustEmbedUnimplementedNodeCVEServiceServer() {}

// UnsafeNodeCVEServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeCVEServiceServer will
// result in compilation errors.
type UnsafeNodeCVEServiceServer interface {
	mustEmbedUnimplementedNodeCVEServiceServer()
}

func RegisterNodeCVEServiceServer(s grpc.ServiceRegistrar, srv NodeCVEServiceServer) {
	s.RegisterService(&NodeCVEService_ServiceDesc, srv)
}

func _NodeCVEService_SuppressCVEs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuppressCVERequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCVEServiceServer).SuppressCVEs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeCVEService_SuppressCVEs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCVEServiceServer).SuppressCVEs(ctx, req.(*SuppressCVERequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCVEService_UnsuppressCVEs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsuppressCVERequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCVEServiceServer).UnsuppressCVEs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeCVEService_UnsuppressCVEs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCVEServiceServer).UnsuppressCVEs(ctx, req.(*UnsuppressCVERequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeCVEService_ServiceDesc is the grpc.ServiceDesc for NodeCVEService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeCVEService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.NodeCVEService",
	HandlerType: (*NodeCVEServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuppressCVEs",
			Handler:    _NodeCVEService_SuppressCVEs_Handler,
		},
		{
			MethodName: "UnsuppressCVEs",
			Handler:    _NodeCVEService_UnsuppressCVEs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/cve_service.proto",
}

const (
	ClusterCVEService_SuppressCVEs_FullMethodName   = "/v1.ClusterCVEService/SuppressCVEs"
	ClusterCVEService_UnsuppressCVEs_FullMethodName = "/v1.ClusterCVEService/UnsuppressCVEs"
)

// ClusterCVEServiceClient is the client API for ClusterCVEService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterCVEServiceClient interface {
	// SuppressCVE suppresses cluster cves.
	SuppressCVEs(ctx context.Context, in *SuppressCVERequest, opts ...grpc.CallOption) (*Empty, error)
	// UnsuppressCVE unsuppresses cluster cves.
	UnsuppressCVEs(ctx context.Context, in *UnsuppressCVERequest, opts ...grpc.CallOption) (*Empty, error)
}

type clusterCVEServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterCVEServiceClient(cc grpc.ClientConnInterface) ClusterCVEServiceClient {
	return &clusterCVEServiceClient{cc}
}

func (c *clusterCVEServiceClient) SuppressCVEs(ctx context.Context, in *SuppressCVERequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ClusterCVEService_SuppressCVEs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterCVEServiceClient) UnsuppressCVEs(ctx context.Context, in *UnsuppressCVERequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ClusterCVEService_UnsuppressCVEs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterCVEServiceServer is the server API for ClusterCVEService service.
// All implementations must embed UnimplementedClusterCVEServiceServer
// for forward compatibility
type ClusterCVEServiceServer interface {
	// SuppressCVE suppresses cluster cves.
	SuppressCVEs(context.Context, *SuppressCVERequest) (*Empty, error)
	// UnsuppressCVE unsuppresses cluster cves.
	UnsuppressCVEs(context.Context, *UnsuppressCVERequest) (*Empty, error)
	mustEmbedUnimplementedClusterCVEServiceServer()
}

// UnimplementedClusterCVEServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClusterCVEServiceServer struct {
}

func (UnimplementedClusterCVEServiceServer) SuppressCVEs(context.Context, *SuppressCVERequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuppressCVEs not implemented")
}
func (UnimplementedClusterCVEServiceServer) UnsuppressCVEs(context.Context, *UnsuppressCVERequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsuppressCVEs not implemented")
}
func (UnimplementedClusterCVEServiceServer) mustEmbedUnimplementedClusterCVEServiceServer() {}

// UnsafeClusterCVEServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterCVEServiceServer will
// result in compilation errors.
type UnsafeClusterCVEServiceServer interface {
	mustEmbedUnimplementedClusterCVEServiceServer()
}

func RegisterClusterCVEServiceServer(s grpc.ServiceRegistrar, srv ClusterCVEServiceServer) {
	s.RegisterService(&ClusterCVEService_ServiceDesc, srv)
}

func _ClusterCVEService_SuppressCVEs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuppressCVERequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterCVEServiceServer).SuppressCVEs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterCVEService_SuppressCVEs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterCVEServiceServer).SuppressCVEs(ctx, req.(*SuppressCVERequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterCVEService_UnsuppressCVEs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsuppressCVERequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterCVEServiceServer).UnsuppressCVEs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterCVEService_UnsuppressCVEs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterCVEServiceServer).UnsuppressCVEs(ctx, req.(*UnsuppressCVERequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterCVEService_ServiceDesc is the grpc.ServiceDesc for ClusterCVEService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterCVEService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ClusterCVEService",
	HandlerType: (*ClusterCVEServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuppressCVEs",
			Handler:    _ClusterCVEService_SuppressCVEs_Handler,
		},
		{
			MethodName: "UnsuppressCVEs",
			Handler:    _ClusterCVEService_UnsuppressCVEs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/cve_service.proto",
}
