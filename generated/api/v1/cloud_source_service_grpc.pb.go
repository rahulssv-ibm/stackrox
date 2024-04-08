// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v1/cloud_source_service.proto

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
	CloudSourcesService_CountCloudSources_FullMethodName = "/v1.CloudSourcesService/CountCloudSources"
	CloudSourcesService_GetCloudSource_FullMethodName    = "/v1.CloudSourcesService/GetCloudSource"
	CloudSourcesService_ListCloudSources_FullMethodName  = "/v1.CloudSourcesService/ListCloudSources"
	CloudSourcesService_CreateCloudSource_FullMethodName = "/v1.CloudSourcesService/CreateCloudSource"
	CloudSourcesService_UpdateCloudSource_FullMethodName = "/v1.CloudSourcesService/UpdateCloudSource"
	CloudSourcesService_DeleteCloudSource_FullMethodName = "/v1.CloudSourcesService/DeleteCloudSource"
	CloudSourcesService_TestCloudSource_FullMethodName   = "/v1.CloudSourcesService/TestCloudSource"
)

// CloudSourcesServiceClient is the client API for CloudSourcesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudSourcesServiceClient interface {
	// CountCloudSources returns the number of cloud sources after filtering by requested fields.
	CountCloudSources(ctx context.Context, in *CountCloudSourcesRequest, opts ...grpc.CallOption) (*CountCloudSourcesResponse, error)
	// GetCloudSource retrieves a cloud source by ID.
	GetCloudSource(ctx context.Context, in *GetCloudSourceRequest, opts ...grpc.CallOption) (*GetCloudSourceResponse, error)
	// ListCloudSources returns the list of cloud sources after filtered by requested fields.
	ListCloudSources(ctx context.Context, in *ListCloudSourcesRequest, opts ...grpc.CallOption) (*ListCloudSourcesResponse, error)
	// CreateCloudSource creates a cloud source.
	CreateCloudSource(ctx context.Context, in *CreateCloudSourceRequest, opts ...grpc.CallOption) (*CreateCloudSourceResponse, error)
	// UpdateCloudSource creates or replaces a cloud source.
	UpdateCloudSource(ctx context.Context, in *UpdateCloudSourceRequest, opts ...grpc.CallOption) (*Empty, error)
	// DeleteCloudSource removes a cloud source.
	DeleteCloudSource(ctx context.Context, in *DeleteCloudSourceRequest, opts ...grpc.CallOption) (*Empty, error)
	// TestCloudSource tests a cloud source.
	TestCloudSource(ctx context.Context, in *TestCloudSourceRequest, opts ...grpc.CallOption) (*Empty, error)
}

type cloudSourcesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudSourcesServiceClient(cc grpc.ClientConnInterface) CloudSourcesServiceClient {
	return &cloudSourcesServiceClient{cc}
}

func (c *cloudSourcesServiceClient) CountCloudSources(ctx context.Context, in *CountCloudSourcesRequest, opts ...grpc.CallOption) (*CountCloudSourcesResponse, error) {
	out := new(CountCloudSourcesResponse)
	err := c.cc.Invoke(ctx, CloudSourcesService_CountCloudSources_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudSourcesServiceClient) GetCloudSource(ctx context.Context, in *GetCloudSourceRequest, opts ...grpc.CallOption) (*GetCloudSourceResponse, error) {
	out := new(GetCloudSourceResponse)
	err := c.cc.Invoke(ctx, CloudSourcesService_GetCloudSource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudSourcesServiceClient) ListCloudSources(ctx context.Context, in *ListCloudSourcesRequest, opts ...grpc.CallOption) (*ListCloudSourcesResponse, error) {
	out := new(ListCloudSourcesResponse)
	err := c.cc.Invoke(ctx, CloudSourcesService_ListCloudSources_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudSourcesServiceClient) CreateCloudSource(ctx context.Context, in *CreateCloudSourceRequest, opts ...grpc.CallOption) (*CreateCloudSourceResponse, error) {
	out := new(CreateCloudSourceResponse)
	err := c.cc.Invoke(ctx, CloudSourcesService_CreateCloudSource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudSourcesServiceClient) UpdateCloudSource(ctx context.Context, in *UpdateCloudSourceRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, CloudSourcesService_UpdateCloudSource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudSourcesServiceClient) DeleteCloudSource(ctx context.Context, in *DeleteCloudSourceRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, CloudSourcesService_DeleteCloudSource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudSourcesServiceClient) TestCloudSource(ctx context.Context, in *TestCloudSourceRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, CloudSourcesService_TestCloudSource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudSourcesServiceServer is the server API for CloudSourcesService service.
// All implementations should embed UnimplementedCloudSourcesServiceServer
// for forward compatibility
type CloudSourcesServiceServer interface {
	// CountCloudSources returns the number of cloud sources after filtering by requested fields.
	CountCloudSources(context.Context, *CountCloudSourcesRequest) (*CountCloudSourcesResponse, error)
	// GetCloudSource retrieves a cloud source by ID.
	GetCloudSource(context.Context, *GetCloudSourceRequest) (*GetCloudSourceResponse, error)
	// ListCloudSources returns the list of cloud sources after filtered by requested fields.
	ListCloudSources(context.Context, *ListCloudSourcesRequest) (*ListCloudSourcesResponse, error)
	// CreateCloudSource creates a cloud source.
	CreateCloudSource(context.Context, *CreateCloudSourceRequest) (*CreateCloudSourceResponse, error)
	// UpdateCloudSource creates or replaces a cloud source.
	UpdateCloudSource(context.Context, *UpdateCloudSourceRequest) (*Empty, error)
	// DeleteCloudSource removes a cloud source.
	DeleteCloudSource(context.Context, *DeleteCloudSourceRequest) (*Empty, error)
	// TestCloudSource tests a cloud source.
	TestCloudSource(context.Context, *TestCloudSourceRequest) (*Empty, error)
}

// UnimplementedCloudSourcesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCloudSourcesServiceServer struct {
}

func (UnimplementedCloudSourcesServiceServer) CountCloudSources(context.Context, *CountCloudSourcesRequest) (*CountCloudSourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountCloudSources not implemented")
}
func (UnimplementedCloudSourcesServiceServer) GetCloudSource(context.Context, *GetCloudSourceRequest) (*GetCloudSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCloudSource not implemented")
}
func (UnimplementedCloudSourcesServiceServer) ListCloudSources(context.Context, *ListCloudSourcesRequest) (*ListCloudSourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCloudSources not implemented")
}
func (UnimplementedCloudSourcesServiceServer) CreateCloudSource(context.Context, *CreateCloudSourceRequest) (*CreateCloudSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCloudSource not implemented")
}
func (UnimplementedCloudSourcesServiceServer) UpdateCloudSource(context.Context, *UpdateCloudSourceRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCloudSource not implemented")
}
func (UnimplementedCloudSourcesServiceServer) DeleteCloudSource(context.Context, *DeleteCloudSourceRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCloudSource not implemented")
}
func (UnimplementedCloudSourcesServiceServer) TestCloudSource(context.Context, *TestCloudSourceRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestCloudSource not implemented")
}

// UnsafeCloudSourcesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudSourcesServiceServer will
// result in compilation errors.
type UnsafeCloudSourcesServiceServer interface {
	mustEmbedUnimplementedCloudSourcesServiceServer()
}

func RegisterCloudSourcesServiceServer(s grpc.ServiceRegistrar, srv CloudSourcesServiceServer) {
	s.RegisterService(&CloudSourcesService_ServiceDesc, srv)
}

func _CloudSourcesService_CountCloudSources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountCloudSourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudSourcesServiceServer).CountCloudSources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudSourcesService_CountCloudSources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudSourcesServiceServer).CountCloudSources(ctx, req.(*CountCloudSourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudSourcesService_GetCloudSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCloudSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudSourcesServiceServer).GetCloudSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudSourcesService_GetCloudSource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudSourcesServiceServer).GetCloudSource(ctx, req.(*GetCloudSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudSourcesService_ListCloudSources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCloudSourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudSourcesServiceServer).ListCloudSources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudSourcesService_ListCloudSources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudSourcesServiceServer).ListCloudSources(ctx, req.(*ListCloudSourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudSourcesService_CreateCloudSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCloudSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudSourcesServiceServer).CreateCloudSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudSourcesService_CreateCloudSource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudSourcesServiceServer).CreateCloudSource(ctx, req.(*CreateCloudSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudSourcesService_UpdateCloudSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCloudSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudSourcesServiceServer).UpdateCloudSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudSourcesService_UpdateCloudSource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudSourcesServiceServer).UpdateCloudSource(ctx, req.(*UpdateCloudSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudSourcesService_DeleteCloudSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCloudSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudSourcesServiceServer).DeleteCloudSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudSourcesService_DeleteCloudSource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudSourcesServiceServer).DeleteCloudSource(ctx, req.(*DeleteCloudSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudSourcesService_TestCloudSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestCloudSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudSourcesServiceServer).TestCloudSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudSourcesService_TestCloudSource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudSourcesServiceServer).TestCloudSource(ctx, req.(*TestCloudSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudSourcesService_ServiceDesc is the grpc.ServiceDesc for CloudSourcesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudSourcesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.CloudSourcesService",
	HandlerType: (*CloudSourcesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CountCloudSources",
			Handler:    _CloudSourcesService_CountCloudSources_Handler,
		},
		{
			MethodName: "GetCloudSource",
			Handler:    _CloudSourcesService_GetCloudSource_Handler,
		},
		{
			MethodName: "ListCloudSources",
			Handler:    _CloudSourcesService_ListCloudSources_Handler,
		},
		{
			MethodName: "CreateCloudSource",
			Handler:    _CloudSourcesService_CreateCloudSource_Handler,
		},
		{
			MethodName: "UpdateCloudSource",
			Handler:    _CloudSourcesService_UpdateCloudSource_Handler,
		},
		{
			MethodName: "DeleteCloudSource",
			Handler:    _CloudSourcesService_DeleteCloudSource_Handler,
		},
		{
			MethodName: "TestCloudSource",
			Handler:    _CloudSourcesService_TestCloudSource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/cloud_source_service.proto",
}
