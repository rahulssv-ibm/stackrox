// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: api/v1/image_integration_service.proto

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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ImageIntegrationService_GetImageIntegration_FullMethodName         = "/v1.ImageIntegrationService/GetImageIntegration"
	ImageIntegrationService_GetImageIntegrations_FullMethodName        = "/v1.ImageIntegrationService/GetImageIntegrations"
	ImageIntegrationService_PostImageIntegration_FullMethodName        = "/v1.ImageIntegrationService/PostImageIntegration"
	ImageIntegrationService_PutImageIntegration_FullMethodName         = "/v1.ImageIntegrationService/PutImageIntegration"
	ImageIntegrationService_TestImageIntegration_FullMethodName        = "/v1.ImageIntegrationService/TestImageIntegration"
	ImageIntegrationService_DeleteImageIntegration_FullMethodName      = "/v1.ImageIntegrationService/DeleteImageIntegration"
	ImageIntegrationService_UpdateImageIntegration_FullMethodName      = "/v1.ImageIntegrationService/UpdateImageIntegration"
	ImageIntegrationService_TestUpdatedImageIntegration_FullMethodName = "/v1.ImageIntegrationService/TestUpdatedImageIntegration"
)

// ImageIntegrationServiceClient is the client API for ImageIntegrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ImageIntegrationService APIs manage image registry and image scanner integration.
type ImageIntegrationServiceClient interface {
	// GetImageIntegration returns the image integration given its ID.
	GetImageIntegration(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*storage.ImageIntegration, error)
	// GetImageIntegrations returns all image integrations that match the request filters.
	GetImageIntegrations(ctx context.Context, in *GetImageIntegrationsRequest, opts ...grpc.CallOption) (*GetImageIntegrationsResponse, error)
	// PostImageIntegration creates a image integration.
	PostImageIntegration(ctx context.Context, in *storage.ImageIntegration, opts ...grpc.CallOption) (*storage.ImageIntegration, error)
	// PutImageIntegration modifies a given image integration, without using stored credential reconciliation.
	PutImageIntegration(ctx context.Context, in *storage.ImageIntegration, opts ...grpc.CallOption) (*Empty, error)
	// TestImageIntegration checks if the given image integration is correctly configured, without using stored credential reconciliation.
	TestImageIntegration(ctx context.Context, in *storage.ImageIntegration, opts ...grpc.CallOption) (*Empty, error)
	// DeleteImageIntegration removes a image integration given its ID.
	DeleteImageIntegration(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*Empty, error)
	// UpdateImageIntegration modifies a given image integration, with optional stored credential reconciliation.
	UpdateImageIntegration(ctx context.Context, in *UpdateImageIntegrationRequest, opts ...grpc.CallOption) (*Empty, error)
	// TestUpdatedImageIntegration checks if the given image integration is correctly configured, with optional stored credential reconciliation.
	TestUpdatedImageIntegration(ctx context.Context, in *UpdateImageIntegrationRequest, opts ...grpc.CallOption) (*Empty, error)
}

type imageIntegrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageIntegrationServiceClient(cc grpc.ClientConnInterface) ImageIntegrationServiceClient {
	return &imageIntegrationServiceClient{cc}
}

func (c *imageIntegrationServiceClient) GetImageIntegration(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*storage.ImageIntegration, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(storage.ImageIntegration)
	err := c.cc.Invoke(ctx, ImageIntegrationService_GetImageIntegration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageIntegrationServiceClient) GetImageIntegrations(ctx context.Context, in *GetImageIntegrationsRequest, opts ...grpc.CallOption) (*GetImageIntegrationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetImageIntegrationsResponse)
	err := c.cc.Invoke(ctx, ImageIntegrationService_GetImageIntegrations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageIntegrationServiceClient) PostImageIntegration(ctx context.Context, in *storage.ImageIntegration, opts ...grpc.CallOption) (*storage.ImageIntegration, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(storage.ImageIntegration)
	err := c.cc.Invoke(ctx, ImageIntegrationService_PostImageIntegration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageIntegrationServiceClient) PutImageIntegration(ctx context.Context, in *storage.ImageIntegration, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, ImageIntegrationService_PutImageIntegration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageIntegrationServiceClient) TestImageIntegration(ctx context.Context, in *storage.ImageIntegration, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, ImageIntegrationService_TestImageIntegration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageIntegrationServiceClient) DeleteImageIntegration(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, ImageIntegrationService_DeleteImageIntegration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageIntegrationServiceClient) UpdateImageIntegration(ctx context.Context, in *UpdateImageIntegrationRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, ImageIntegrationService_UpdateImageIntegration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageIntegrationServiceClient) TestUpdatedImageIntegration(ctx context.Context, in *UpdateImageIntegrationRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, ImageIntegrationService_TestUpdatedImageIntegration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageIntegrationServiceServer is the server API for ImageIntegrationService service.
// All implementations should embed UnimplementedImageIntegrationServiceServer
// for forward compatibility.
//
// ImageIntegrationService APIs manage image registry and image scanner integration.
type ImageIntegrationServiceServer interface {
	// GetImageIntegration returns the image integration given its ID.
	GetImageIntegration(context.Context, *ResourceByID) (*storage.ImageIntegration, error)
	// GetImageIntegrations returns all image integrations that match the request filters.
	GetImageIntegrations(context.Context, *GetImageIntegrationsRequest) (*GetImageIntegrationsResponse, error)
	// PostImageIntegration creates a image integration.
	PostImageIntegration(context.Context, *storage.ImageIntegration) (*storage.ImageIntegration, error)
	// PutImageIntegration modifies a given image integration, without using stored credential reconciliation.
	PutImageIntegration(context.Context, *storage.ImageIntegration) (*Empty, error)
	// TestImageIntegration checks if the given image integration is correctly configured, without using stored credential reconciliation.
	TestImageIntegration(context.Context, *storage.ImageIntegration) (*Empty, error)
	// DeleteImageIntegration removes a image integration given its ID.
	DeleteImageIntegration(context.Context, *ResourceByID) (*Empty, error)
	// UpdateImageIntegration modifies a given image integration, with optional stored credential reconciliation.
	UpdateImageIntegration(context.Context, *UpdateImageIntegrationRequest) (*Empty, error)
	// TestUpdatedImageIntegration checks if the given image integration is correctly configured, with optional stored credential reconciliation.
	TestUpdatedImageIntegration(context.Context, *UpdateImageIntegrationRequest) (*Empty, error)
}

// UnimplementedImageIntegrationServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedImageIntegrationServiceServer struct{}

func (UnimplementedImageIntegrationServiceServer) GetImageIntegration(context.Context, *ResourceByID) (*storage.ImageIntegration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImageIntegration not implemented")
}
func (UnimplementedImageIntegrationServiceServer) GetImageIntegrations(context.Context, *GetImageIntegrationsRequest) (*GetImageIntegrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImageIntegrations not implemented")
}
func (UnimplementedImageIntegrationServiceServer) PostImageIntegration(context.Context, *storage.ImageIntegration) (*storage.ImageIntegration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostImageIntegration not implemented")
}
func (UnimplementedImageIntegrationServiceServer) PutImageIntegration(context.Context, *storage.ImageIntegration) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutImageIntegration not implemented")
}
func (UnimplementedImageIntegrationServiceServer) TestImageIntegration(context.Context, *storage.ImageIntegration) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestImageIntegration not implemented")
}
func (UnimplementedImageIntegrationServiceServer) DeleteImageIntegration(context.Context, *ResourceByID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImageIntegration not implemented")
}
func (UnimplementedImageIntegrationServiceServer) UpdateImageIntegration(context.Context, *UpdateImageIntegrationRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateImageIntegration not implemented")
}
func (UnimplementedImageIntegrationServiceServer) TestUpdatedImageIntegration(context.Context, *UpdateImageIntegrationRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestUpdatedImageIntegration not implemented")
}
func (UnimplementedImageIntegrationServiceServer) testEmbeddedByValue() {}

// UnsafeImageIntegrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageIntegrationServiceServer will
// result in compilation errors.
type UnsafeImageIntegrationServiceServer interface {
	mustEmbedUnimplementedImageIntegrationServiceServer()
}

func RegisterImageIntegrationServiceServer(s grpc.ServiceRegistrar, srv ImageIntegrationServiceServer) {
	// If the following call pancis, it indicates UnimplementedImageIntegrationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ImageIntegrationService_ServiceDesc, srv)
}

func _ImageIntegrationService_GetImageIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageIntegrationServiceServer).GetImageIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageIntegrationService_GetImageIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageIntegrationServiceServer).GetImageIntegration(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageIntegrationService_GetImageIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageIntegrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageIntegrationServiceServer).GetImageIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageIntegrationService_GetImageIntegrations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageIntegrationServiceServer).GetImageIntegrations(ctx, req.(*GetImageIntegrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageIntegrationService_PostImageIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(storage.ImageIntegration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageIntegrationServiceServer).PostImageIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageIntegrationService_PostImageIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageIntegrationServiceServer).PostImageIntegration(ctx, req.(*storage.ImageIntegration))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageIntegrationService_PutImageIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(storage.ImageIntegration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageIntegrationServiceServer).PutImageIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageIntegrationService_PutImageIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageIntegrationServiceServer).PutImageIntegration(ctx, req.(*storage.ImageIntegration))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageIntegrationService_TestImageIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(storage.ImageIntegration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageIntegrationServiceServer).TestImageIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageIntegrationService_TestImageIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageIntegrationServiceServer).TestImageIntegration(ctx, req.(*storage.ImageIntegration))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageIntegrationService_DeleteImageIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageIntegrationServiceServer).DeleteImageIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageIntegrationService_DeleteImageIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageIntegrationServiceServer).DeleteImageIntegration(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageIntegrationService_UpdateImageIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateImageIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageIntegrationServiceServer).UpdateImageIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageIntegrationService_UpdateImageIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageIntegrationServiceServer).UpdateImageIntegration(ctx, req.(*UpdateImageIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageIntegrationService_TestUpdatedImageIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateImageIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageIntegrationServiceServer).TestUpdatedImageIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageIntegrationService_TestUpdatedImageIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageIntegrationServiceServer).TestUpdatedImageIntegration(ctx, req.(*UpdateImageIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageIntegrationService_ServiceDesc is the grpc.ServiceDesc for ImageIntegrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageIntegrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ImageIntegrationService",
	HandlerType: (*ImageIntegrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetImageIntegration",
			Handler:    _ImageIntegrationService_GetImageIntegration_Handler,
		},
		{
			MethodName: "GetImageIntegrations",
			Handler:    _ImageIntegrationService_GetImageIntegrations_Handler,
		},
		{
			MethodName: "PostImageIntegration",
			Handler:    _ImageIntegrationService_PostImageIntegration_Handler,
		},
		{
			MethodName: "PutImageIntegration",
			Handler:    _ImageIntegrationService_PutImageIntegration_Handler,
		},
		{
			MethodName: "TestImageIntegration",
			Handler:    _ImageIntegrationService_TestImageIntegration_Handler,
		},
		{
			MethodName: "DeleteImageIntegration",
			Handler:    _ImageIntegrationService_DeleteImageIntegration_Handler,
		},
		{
			MethodName: "UpdateImageIntegration",
			Handler:    _ImageIntegrationService_UpdateImageIntegration_Handler,
		},
		{
			MethodName: "TestUpdatedImageIntegration",
			Handler:    _ImageIntegrationService_TestUpdatedImageIntegration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/image_integration_service.proto",
}
