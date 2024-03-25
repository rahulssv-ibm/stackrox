// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v1/network_graph_service.proto

package v1

import (
	context "context"
	storage "github.com/stackrox/rox/generated/storage"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	NetworkGraphService_GetNetworkGraph_FullMethodName             = "/v1.NetworkGraphService/GetNetworkGraph"
	NetworkGraphService_GetExternalNetworkEntities_FullMethodName  = "/v1.NetworkGraphService/GetExternalNetworkEntities"
	NetworkGraphService_CreateExternalNetworkEntity_FullMethodName = "/v1.NetworkGraphService/CreateExternalNetworkEntity"
	NetworkGraphService_PatchExternalNetworkEntity_FullMethodName  = "/v1.NetworkGraphService/PatchExternalNetworkEntity"
	NetworkGraphService_DeleteExternalNetworkEntity_FullMethodName = "/v1.NetworkGraphService/DeleteExternalNetworkEntity"
	NetworkGraphService_GetNetworkGraphConfig_FullMethodName       = "/v1.NetworkGraphService/GetNetworkGraphConfig"
	NetworkGraphService_PutNetworkGraphConfig_FullMethodName       = "/v1.NetworkGraphService/PutNetworkGraphConfig"
)

// NetworkGraphServiceClient is the client API for NetworkGraphService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkGraphServiceClient interface {
	GetNetworkGraph(ctx context.Context, in *NetworkGraphRequest, opts ...grpc.CallOption) (*NetworkGraph, error)
	GetExternalNetworkEntities(ctx context.Context, in *GetExternalNetworkEntitiesRequest, opts ...grpc.CallOption) (*GetExternalNetworkEntitiesResponse, error)
	CreateExternalNetworkEntity(ctx context.Context, in *CreateNetworkEntityRequest, opts ...grpc.CallOption) (*storage.NetworkEntity, error)
	PatchExternalNetworkEntity(ctx context.Context, in *PatchNetworkEntityRequest, opts ...grpc.CallOption) (*storage.NetworkEntity, error)
	DeleteExternalNetworkEntity(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*Empty, error)
	GetNetworkGraphConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*storage.NetworkGraphConfig, error)
	PutNetworkGraphConfig(ctx context.Context, in *PutNetworkGraphConfigRequest, opts ...grpc.CallOption) (*storage.NetworkGraphConfig, error)
}

type networkGraphServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkGraphServiceClient(cc grpc.ClientConnInterface) NetworkGraphServiceClient {
	return &networkGraphServiceClient{cc}
}

func (c *networkGraphServiceClient) GetNetworkGraph(ctx context.Context, in *NetworkGraphRequest, opts ...grpc.CallOption) (*NetworkGraph, error) {
	out := new(NetworkGraph)
	err := c.cc.Invoke(ctx, NetworkGraphService_GetNetworkGraph_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkGraphServiceClient) GetExternalNetworkEntities(ctx context.Context, in *GetExternalNetworkEntitiesRequest, opts ...grpc.CallOption) (*GetExternalNetworkEntitiesResponse, error) {
	out := new(GetExternalNetworkEntitiesResponse)
	err := c.cc.Invoke(ctx, NetworkGraphService_GetExternalNetworkEntities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkGraphServiceClient) CreateExternalNetworkEntity(ctx context.Context, in *CreateNetworkEntityRequest, opts ...grpc.CallOption) (*storage.NetworkEntity, error) {
	out := new(storage.NetworkEntity)
	err := c.cc.Invoke(ctx, NetworkGraphService_CreateExternalNetworkEntity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkGraphServiceClient) PatchExternalNetworkEntity(ctx context.Context, in *PatchNetworkEntityRequest, opts ...grpc.CallOption) (*storage.NetworkEntity, error) {
	out := new(storage.NetworkEntity)
	err := c.cc.Invoke(ctx, NetworkGraphService_PatchExternalNetworkEntity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkGraphServiceClient) DeleteExternalNetworkEntity(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, NetworkGraphService_DeleteExternalNetworkEntity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkGraphServiceClient) GetNetworkGraphConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*storage.NetworkGraphConfig, error) {
	out := new(storage.NetworkGraphConfig)
	err := c.cc.Invoke(ctx, NetworkGraphService_GetNetworkGraphConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkGraphServiceClient) PutNetworkGraphConfig(ctx context.Context, in *PutNetworkGraphConfigRequest, opts ...grpc.CallOption) (*storage.NetworkGraphConfig, error) {
	out := new(storage.NetworkGraphConfig)
	err := c.cc.Invoke(ctx, NetworkGraphService_PutNetworkGraphConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkGraphServiceServer is the server API for NetworkGraphService service.
// All implementations must embed UnimplementedNetworkGraphServiceServer
// for forward compatibility
type NetworkGraphServiceServer interface {
	GetNetworkGraph(context.Context, *NetworkGraphRequest) (*NetworkGraph, error)
	GetExternalNetworkEntities(context.Context, *GetExternalNetworkEntitiesRequest) (*GetExternalNetworkEntitiesResponse, error)
	CreateExternalNetworkEntity(context.Context, *CreateNetworkEntityRequest) (*storage.NetworkEntity, error)
	PatchExternalNetworkEntity(context.Context, *PatchNetworkEntityRequest) (*storage.NetworkEntity, error)
	DeleteExternalNetworkEntity(context.Context, *ResourceByID) (*Empty, error)
	GetNetworkGraphConfig(context.Context, *Empty) (*storage.NetworkGraphConfig, error)
	PutNetworkGraphConfig(context.Context, *PutNetworkGraphConfigRequest) (*storage.NetworkGraphConfig, error)
	mustEmbedUnimplementedNetworkGraphServiceServer()
}

// UnimplementedNetworkGraphServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkGraphServiceServer struct {
}

func (UnimplementedNetworkGraphServiceServer) GetNetworkGraph(context.Context, *NetworkGraphRequest) (*NetworkGraph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkGraph not implemented")
}
func (UnimplementedNetworkGraphServiceServer) GetExternalNetworkEntities(context.Context, *GetExternalNetworkEntitiesRequest) (*GetExternalNetworkEntitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExternalNetworkEntities not implemented")
}
func (UnimplementedNetworkGraphServiceServer) CreateExternalNetworkEntity(context.Context, *CreateNetworkEntityRequest) (*storage.NetworkEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExternalNetworkEntity not implemented")
}
func (UnimplementedNetworkGraphServiceServer) PatchExternalNetworkEntity(context.Context, *PatchNetworkEntityRequest) (*storage.NetworkEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchExternalNetworkEntity not implemented")
}
func (UnimplementedNetworkGraphServiceServer) DeleteExternalNetworkEntity(context.Context, *ResourceByID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExternalNetworkEntity not implemented")
}
func (UnimplementedNetworkGraphServiceServer) GetNetworkGraphConfig(context.Context, *Empty) (*storage.NetworkGraphConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkGraphConfig not implemented")
}
func (UnimplementedNetworkGraphServiceServer) PutNetworkGraphConfig(context.Context, *PutNetworkGraphConfigRequest) (*storage.NetworkGraphConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutNetworkGraphConfig not implemented")
}
func (UnimplementedNetworkGraphServiceServer) mustEmbedUnimplementedNetworkGraphServiceServer() {}

// UnsafeNetworkGraphServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkGraphServiceServer will
// result in compilation errors.
type UnsafeNetworkGraphServiceServer interface {
	mustEmbedUnimplementedNetworkGraphServiceServer()
}

func RegisterNetworkGraphServiceServer(s grpc.ServiceRegistrar, srv NetworkGraphServiceServer) {
	s.RegisterService(&NetworkGraphService_ServiceDesc, srv)
}

func _NetworkGraphService_GetNetworkGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkGraphServiceServer).GetNetworkGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkGraphService_GetNetworkGraph_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkGraphServiceServer).GetNetworkGraph(ctx, req.(*NetworkGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkGraphService_GetExternalNetworkEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExternalNetworkEntitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkGraphServiceServer).GetExternalNetworkEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkGraphService_GetExternalNetworkEntities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkGraphServiceServer).GetExternalNetworkEntities(ctx, req.(*GetExternalNetworkEntitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkGraphService_CreateExternalNetworkEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNetworkEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkGraphServiceServer).CreateExternalNetworkEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkGraphService_CreateExternalNetworkEntity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkGraphServiceServer).CreateExternalNetworkEntity(ctx, req.(*CreateNetworkEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkGraphService_PatchExternalNetworkEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchNetworkEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkGraphServiceServer).PatchExternalNetworkEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkGraphService_PatchExternalNetworkEntity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkGraphServiceServer).PatchExternalNetworkEntity(ctx, req.(*PatchNetworkEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkGraphService_DeleteExternalNetworkEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkGraphServiceServer).DeleteExternalNetworkEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkGraphService_DeleteExternalNetworkEntity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkGraphServiceServer).DeleteExternalNetworkEntity(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkGraphService_GetNetworkGraphConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkGraphServiceServer).GetNetworkGraphConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkGraphService_GetNetworkGraphConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkGraphServiceServer).GetNetworkGraphConfig(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkGraphService_PutNetworkGraphConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutNetworkGraphConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkGraphServiceServer).PutNetworkGraphConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkGraphService_PutNetworkGraphConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkGraphServiceServer).PutNetworkGraphConfig(ctx, req.(*PutNetworkGraphConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkGraphService_ServiceDesc is the grpc.ServiceDesc for NetworkGraphService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkGraphService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.NetworkGraphService",
	HandlerType: (*NetworkGraphServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNetworkGraph",
			Handler:    _NetworkGraphService_GetNetworkGraph_Handler,
		},
		{
			MethodName: "GetExternalNetworkEntities",
			Handler:    _NetworkGraphService_GetExternalNetworkEntities_Handler,
		},
		{
			MethodName: "CreateExternalNetworkEntity",
			Handler:    _NetworkGraphService_CreateExternalNetworkEntity_Handler,
		},
		{
			MethodName: "PatchExternalNetworkEntity",
			Handler:    _NetworkGraphService_PatchExternalNetworkEntity_Handler,
		},
		{
			MethodName: "DeleteExternalNetworkEntity",
			Handler:    _NetworkGraphService_DeleteExternalNetworkEntity_Handler,
		},
		{
			MethodName: "GetNetworkGraphConfig",
			Handler:    _NetworkGraphService_GetNetworkGraphConfig_Handler,
		},
		{
			MethodName: "PutNetworkGraphConfig",
			Handler:    _NetworkGraphService_PutNetworkGraphConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/network_graph_service.proto",
}
