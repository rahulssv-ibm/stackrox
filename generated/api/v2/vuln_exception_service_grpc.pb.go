// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v2/vuln_exception_service.proto

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
	VulnerabilityExceptionService_GetVulnerabilityException_FullMethodName                 = "/v2.VulnerabilityExceptionService/GetVulnerabilityException"
	VulnerabilityExceptionService_ListVulnerabilityExceptions_FullMethodName               = "/v2.VulnerabilityExceptionService/ListVulnerabilityExceptions"
	VulnerabilityExceptionService_CreateDeferVulnerabilityException_FullMethodName         = "/v2.VulnerabilityExceptionService/CreateDeferVulnerabilityException"
	VulnerabilityExceptionService_CreateFalsePositiveVulnerabilityException_FullMethodName = "/v2.VulnerabilityExceptionService/CreateFalsePositiveVulnerabilityException"
	VulnerabilityExceptionService_ApproveVulnerabilityException_FullMethodName             = "/v2.VulnerabilityExceptionService/ApproveVulnerabilityException"
	VulnerabilityExceptionService_DenyVulnerabilityException_FullMethodName                = "/v2.VulnerabilityExceptionService/DenyVulnerabilityException"
	VulnerabilityExceptionService_UpdateVulnerabilityException_FullMethodName              = "/v2.VulnerabilityExceptionService/UpdateVulnerabilityException"
	VulnerabilityExceptionService_CancelVulnerabilityException_FullMethodName              = "/v2.VulnerabilityExceptionService/CancelVulnerabilityException"
	VulnerabilityExceptionService_DeleteVulnerabilityException_FullMethodName              = "/v2.VulnerabilityExceptionService/DeleteVulnerabilityException"
)

// VulnerabilityExceptionServiceClient is the client API for VulnerabilityExceptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VulnerabilityExceptionServiceClient interface {
	// GetVulnerabilityException returns the vulnerability exception with specified ID.
	GetVulnerabilityException(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*GetVulnerabilityExceptionResponse, error)
	// ListVulnerabilityExceptions returns a list of vulnerability exceptions.
	ListVulnerabilityExceptions(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*ListVulnerabilityExceptionsResponse, error)
	// CreateDeferVulnerabilityException creates an exception request to defer specified vulnerabilities.
	// Once an exception is created, it remains in the PENDING state until the approval. The exception is enforced
	// only after it is approved. Once the exception expires, it is garbage collected as per the retention configuration
	// `.expiredVulnReqRetentionDurationDays` (GET`/v1/config/`)
	CreateDeferVulnerabilityException(ctx context.Context, in *CreateDeferVulnerabilityExceptionRequest, opts ...grpc.CallOption) (*CreateDeferVulnerabilityExceptionResponse, error)
	// CreateFalsePositiveVulnerabilityException creates an exception request to mark specified vulnerabilities as false positive.
	// Once an exception is created, it remains in the PENDING state until the approval. The exception is enforced only after it is approved.
	CreateFalsePositiveVulnerabilityException(ctx context.Context, in *CreateFalsePositiveVulnerabilityExceptionRequest, opts ...grpc.CallOption) (*CreateFalsePositiveVulnerabilityExceptionResponse, error)
	// ApproveVulnerabilityException approves a vulnerability exception. Once approved, the exception is enforced.
	// The associated vulnerabilities are excluded from policy evaluation and risk evaluation, and the vulnerabilities
	// may not appear in certain APIs responses by default.
	ApproveVulnerabilityException(ctx context.Context, in *ApproveVulnerabilityExceptionRequest, opts ...grpc.CallOption) (*ApproveVulnerabilityExceptionResponse, error)
	// DenyVulnerabilityException denies a vulnerability exception. Denied exceptions are inactive and are garbage
	// collected as per the retention configuration `.expiredVulnReqRetentionDurationDays` (GET`/v1/config/`)
	DenyVulnerabilityException(ctx context.Context, in *DenyVulnerabilityExceptionRequest, opts ...grpc.CallOption) (*DenyVulnerabilityExceptionResponse, error)
	// UpdateVulnerabilityException updates an existing vulnerability exception. The update is enforced only once it is approved.
	// Currently only the following can be updated:
	// - CVEs and expiry of the deferral exceptions
	// - CVEs of the false positive exception
	UpdateVulnerabilityException(ctx context.Context, in *UpdateVulnerabilityExceptionRequest, opts ...grpc.CallOption) (*UpdateVulnerabilityExceptionResponse, error)
	// CancelVulnerabilityException cancels a vulnerability exception. Once cancelled, an approved exception is no longer
	// enforced. Cancelled exceptions are garbage collected as per the retention configuration
	// `.expiredVulnReqRetentionDurationDays` (GET `/v1/config/`).
	CancelVulnerabilityException(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*CancelVulnerabilityExceptionResponse, error)
	// DeleteVulnerabilityException deletes a vulnerability exception. Only pending exceptions and pending updates
	// to an enforced exception can be deleted. To revert an exception use cancel API. All exceptions are retained
	// in the system according to the retention configuration.
	DeleteVulnerabilityException(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*Empty, error)
}

type vulnerabilityExceptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVulnerabilityExceptionServiceClient(cc grpc.ClientConnInterface) VulnerabilityExceptionServiceClient {
	return &vulnerabilityExceptionServiceClient{cc}
}

func (c *vulnerabilityExceptionServiceClient) GetVulnerabilityException(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*GetVulnerabilityExceptionResponse, error) {
	out := new(GetVulnerabilityExceptionResponse)
	err := c.cc.Invoke(ctx, VulnerabilityExceptionService_GetVulnerabilityException_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vulnerabilityExceptionServiceClient) ListVulnerabilityExceptions(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*ListVulnerabilityExceptionsResponse, error) {
	out := new(ListVulnerabilityExceptionsResponse)
	err := c.cc.Invoke(ctx, VulnerabilityExceptionService_ListVulnerabilityExceptions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vulnerabilityExceptionServiceClient) CreateDeferVulnerabilityException(ctx context.Context, in *CreateDeferVulnerabilityExceptionRequest, opts ...grpc.CallOption) (*CreateDeferVulnerabilityExceptionResponse, error) {
	out := new(CreateDeferVulnerabilityExceptionResponse)
	err := c.cc.Invoke(ctx, VulnerabilityExceptionService_CreateDeferVulnerabilityException_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vulnerabilityExceptionServiceClient) CreateFalsePositiveVulnerabilityException(ctx context.Context, in *CreateFalsePositiveVulnerabilityExceptionRequest, opts ...grpc.CallOption) (*CreateFalsePositiveVulnerabilityExceptionResponse, error) {
	out := new(CreateFalsePositiveVulnerabilityExceptionResponse)
	err := c.cc.Invoke(ctx, VulnerabilityExceptionService_CreateFalsePositiveVulnerabilityException_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vulnerabilityExceptionServiceClient) ApproveVulnerabilityException(ctx context.Context, in *ApproveVulnerabilityExceptionRequest, opts ...grpc.CallOption) (*ApproveVulnerabilityExceptionResponse, error) {
	out := new(ApproveVulnerabilityExceptionResponse)
	err := c.cc.Invoke(ctx, VulnerabilityExceptionService_ApproveVulnerabilityException_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vulnerabilityExceptionServiceClient) DenyVulnerabilityException(ctx context.Context, in *DenyVulnerabilityExceptionRequest, opts ...grpc.CallOption) (*DenyVulnerabilityExceptionResponse, error) {
	out := new(DenyVulnerabilityExceptionResponse)
	err := c.cc.Invoke(ctx, VulnerabilityExceptionService_DenyVulnerabilityException_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vulnerabilityExceptionServiceClient) UpdateVulnerabilityException(ctx context.Context, in *UpdateVulnerabilityExceptionRequest, opts ...grpc.CallOption) (*UpdateVulnerabilityExceptionResponse, error) {
	out := new(UpdateVulnerabilityExceptionResponse)
	err := c.cc.Invoke(ctx, VulnerabilityExceptionService_UpdateVulnerabilityException_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vulnerabilityExceptionServiceClient) CancelVulnerabilityException(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*CancelVulnerabilityExceptionResponse, error) {
	out := new(CancelVulnerabilityExceptionResponse)
	err := c.cc.Invoke(ctx, VulnerabilityExceptionService_CancelVulnerabilityException_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vulnerabilityExceptionServiceClient) DeleteVulnerabilityException(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, VulnerabilityExceptionService_DeleteVulnerabilityException_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VulnerabilityExceptionServiceServer is the server API for VulnerabilityExceptionService service.
// All implementations must embed UnimplementedVulnerabilityExceptionServiceServer
// for forward compatibility
type VulnerabilityExceptionServiceServer interface {
	// GetVulnerabilityException returns the vulnerability exception with specified ID.
	GetVulnerabilityException(context.Context, *ResourceByID) (*GetVulnerabilityExceptionResponse, error)
	// ListVulnerabilityExceptions returns a list of vulnerability exceptions.
	ListVulnerabilityExceptions(context.Context, *RawQuery) (*ListVulnerabilityExceptionsResponse, error)
	// CreateDeferVulnerabilityException creates an exception request to defer specified vulnerabilities.
	// Once an exception is created, it remains in the PENDING state until the approval. The exception is enforced
	// only after it is approved. Once the exception expires, it is garbage collected as per the retention configuration
	// `.expiredVulnReqRetentionDurationDays` (GET`/v1/config/`)
	CreateDeferVulnerabilityException(context.Context, *CreateDeferVulnerabilityExceptionRequest) (*CreateDeferVulnerabilityExceptionResponse, error)
	// CreateFalsePositiveVulnerabilityException creates an exception request to mark specified vulnerabilities as false positive.
	// Once an exception is created, it remains in the PENDING state until the approval. The exception is enforced only after it is approved.
	CreateFalsePositiveVulnerabilityException(context.Context, *CreateFalsePositiveVulnerabilityExceptionRequest) (*CreateFalsePositiveVulnerabilityExceptionResponse, error)
	// ApproveVulnerabilityException approves a vulnerability exception. Once approved, the exception is enforced.
	// The associated vulnerabilities are excluded from policy evaluation and risk evaluation, and the vulnerabilities
	// may not appear in certain APIs responses by default.
	ApproveVulnerabilityException(context.Context, *ApproveVulnerabilityExceptionRequest) (*ApproveVulnerabilityExceptionResponse, error)
	// DenyVulnerabilityException denies a vulnerability exception. Denied exceptions are inactive and are garbage
	// collected as per the retention configuration `.expiredVulnReqRetentionDurationDays` (GET`/v1/config/`)
	DenyVulnerabilityException(context.Context, *DenyVulnerabilityExceptionRequest) (*DenyVulnerabilityExceptionResponse, error)
	// UpdateVulnerabilityException updates an existing vulnerability exception. The update is enforced only once it is approved.
	// Currently only the following can be updated:
	// - CVEs and expiry of the deferral exceptions
	// - CVEs of the false positive exception
	UpdateVulnerabilityException(context.Context, *UpdateVulnerabilityExceptionRequest) (*UpdateVulnerabilityExceptionResponse, error)
	// CancelVulnerabilityException cancels a vulnerability exception. Once cancelled, an approved exception is no longer
	// enforced. Cancelled exceptions are garbage collected as per the retention configuration
	// `.expiredVulnReqRetentionDurationDays` (GET `/v1/config/`).
	CancelVulnerabilityException(context.Context, *ResourceByID) (*CancelVulnerabilityExceptionResponse, error)
	// DeleteVulnerabilityException deletes a vulnerability exception. Only pending exceptions and pending updates
	// to an enforced exception can be deleted. To revert an exception use cancel API. All exceptions are retained
	// in the system according to the retention configuration.
	DeleteVulnerabilityException(context.Context, *ResourceByID) (*Empty, error)
	mustEmbedUnimplementedVulnerabilityExceptionServiceServer()
}

// UnimplementedVulnerabilityExceptionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVulnerabilityExceptionServiceServer struct {
}

func (UnimplementedVulnerabilityExceptionServiceServer) GetVulnerabilityException(context.Context, *ResourceByID) (*GetVulnerabilityExceptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVulnerabilityException not implemented")
}
func (UnimplementedVulnerabilityExceptionServiceServer) ListVulnerabilityExceptions(context.Context, *RawQuery) (*ListVulnerabilityExceptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVulnerabilityExceptions not implemented")
}
func (UnimplementedVulnerabilityExceptionServiceServer) CreateDeferVulnerabilityException(context.Context, *CreateDeferVulnerabilityExceptionRequest) (*CreateDeferVulnerabilityExceptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeferVulnerabilityException not implemented")
}
func (UnimplementedVulnerabilityExceptionServiceServer) CreateFalsePositiveVulnerabilityException(context.Context, *CreateFalsePositiveVulnerabilityExceptionRequest) (*CreateFalsePositiveVulnerabilityExceptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFalsePositiveVulnerabilityException not implemented")
}
func (UnimplementedVulnerabilityExceptionServiceServer) ApproveVulnerabilityException(context.Context, *ApproveVulnerabilityExceptionRequest) (*ApproveVulnerabilityExceptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveVulnerabilityException not implemented")
}
func (UnimplementedVulnerabilityExceptionServiceServer) DenyVulnerabilityException(context.Context, *DenyVulnerabilityExceptionRequest) (*DenyVulnerabilityExceptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenyVulnerabilityException not implemented")
}
func (UnimplementedVulnerabilityExceptionServiceServer) UpdateVulnerabilityException(context.Context, *UpdateVulnerabilityExceptionRequest) (*UpdateVulnerabilityExceptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVulnerabilityException not implemented")
}
func (UnimplementedVulnerabilityExceptionServiceServer) CancelVulnerabilityException(context.Context, *ResourceByID) (*CancelVulnerabilityExceptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelVulnerabilityException not implemented")
}
func (UnimplementedVulnerabilityExceptionServiceServer) DeleteVulnerabilityException(context.Context, *ResourceByID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVulnerabilityException not implemented")
}
func (UnimplementedVulnerabilityExceptionServiceServer) mustEmbedUnimplementedVulnerabilityExceptionServiceServer() {
}

// UnsafeVulnerabilityExceptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VulnerabilityExceptionServiceServer will
// result in compilation errors.
type UnsafeVulnerabilityExceptionServiceServer interface {
	mustEmbedUnimplementedVulnerabilityExceptionServiceServer()
}

func RegisterVulnerabilityExceptionServiceServer(s grpc.ServiceRegistrar, srv VulnerabilityExceptionServiceServer) {
	s.RegisterService(&VulnerabilityExceptionService_ServiceDesc, srv)
}

func _VulnerabilityExceptionService_GetVulnerabilityException_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VulnerabilityExceptionServiceServer).GetVulnerabilityException(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VulnerabilityExceptionService_GetVulnerabilityException_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VulnerabilityExceptionServiceServer).GetVulnerabilityException(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _VulnerabilityExceptionService_ListVulnerabilityExceptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VulnerabilityExceptionServiceServer).ListVulnerabilityExceptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VulnerabilityExceptionService_ListVulnerabilityExceptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VulnerabilityExceptionServiceServer).ListVulnerabilityExceptions(ctx, req.(*RawQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _VulnerabilityExceptionService_CreateDeferVulnerabilityException_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeferVulnerabilityExceptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VulnerabilityExceptionServiceServer).CreateDeferVulnerabilityException(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VulnerabilityExceptionService_CreateDeferVulnerabilityException_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VulnerabilityExceptionServiceServer).CreateDeferVulnerabilityException(ctx, req.(*CreateDeferVulnerabilityExceptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VulnerabilityExceptionService_CreateFalsePositiveVulnerabilityException_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFalsePositiveVulnerabilityExceptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VulnerabilityExceptionServiceServer).CreateFalsePositiveVulnerabilityException(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VulnerabilityExceptionService_CreateFalsePositiveVulnerabilityException_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VulnerabilityExceptionServiceServer).CreateFalsePositiveVulnerabilityException(ctx, req.(*CreateFalsePositiveVulnerabilityExceptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VulnerabilityExceptionService_ApproveVulnerabilityException_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveVulnerabilityExceptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VulnerabilityExceptionServiceServer).ApproveVulnerabilityException(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VulnerabilityExceptionService_ApproveVulnerabilityException_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VulnerabilityExceptionServiceServer).ApproveVulnerabilityException(ctx, req.(*ApproveVulnerabilityExceptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VulnerabilityExceptionService_DenyVulnerabilityException_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DenyVulnerabilityExceptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VulnerabilityExceptionServiceServer).DenyVulnerabilityException(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VulnerabilityExceptionService_DenyVulnerabilityException_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VulnerabilityExceptionServiceServer).DenyVulnerabilityException(ctx, req.(*DenyVulnerabilityExceptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VulnerabilityExceptionService_UpdateVulnerabilityException_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVulnerabilityExceptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VulnerabilityExceptionServiceServer).UpdateVulnerabilityException(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VulnerabilityExceptionService_UpdateVulnerabilityException_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VulnerabilityExceptionServiceServer).UpdateVulnerabilityException(ctx, req.(*UpdateVulnerabilityExceptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VulnerabilityExceptionService_CancelVulnerabilityException_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VulnerabilityExceptionServiceServer).CancelVulnerabilityException(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VulnerabilityExceptionService_CancelVulnerabilityException_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VulnerabilityExceptionServiceServer).CancelVulnerabilityException(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _VulnerabilityExceptionService_DeleteVulnerabilityException_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VulnerabilityExceptionServiceServer).DeleteVulnerabilityException(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VulnerabilityExceptionService_DeleteVulnerabilityException_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VulnerabilityExceptionServiceServer).DeleteVulnerabilityException(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

// VulnerabilityExceptionService_ServiceDesc is the grpc.ServiceDesc for VulnerabilityExceptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VulnerabilityExceptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v2.VulnerabilityExceptionService",
	HandlerType: (*VulnerabilityExceptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVulnerabilityException",
			Handler:    _VulnerabilityExceptionService_GetVulnerabilityException_Handler,
		},
		{
			MethodName: "ListVulnerabilityExceptions",
			Handler:    _VulnerabilityExceptionService_ListVulnerabilityExceptions_Handler,
		},
		{
			MethodName: "CreateDeferVulnerabilityException",
			Handler:    _VulnerabilityExceptionService_CreateDeferVulnerabilityException_Handler,
		},
		{
			MethodName: "CreateFalsePositiveVulnerabilityException",
			Handler:    _VulnerabilityExceptionService_CreateFalsePositiveVulnerabilityException_Handler,
		},
		{
			MethodName: "ApproveVulnerabilityException",
			Handler:    _VulnerabilityExceptionService_ApproveVulnerabilityException_Handler,
		},
		{
			MethodName: "DenyVulnerabilityException",
			Handler:    _VulnerabilityExceptionService_DenyVulnerabilityException_Handler,
		},
		{
			MethodName: "UpdateVulnerabilityException",
			Handler:    _VulnerabilityExceptionService_UpdateVulnerabilityException_Handler,
		},
		{
			MethodName: "CancelVulnerabilityException",
			Handler:    _VulnerabilityExceptionService_CancelVulnerabilityException_Handler,
		},
		{
			MethodName: "DeleteVulnerabilityException",
			Handler:    _VulnerabilityExceptionService_DeleteVulnerabilityException_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v2/vuln_exception_service.proto",
}
