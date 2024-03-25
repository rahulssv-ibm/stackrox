// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v1/alert_service.proto

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
	AlertService_GetAlert_FullMethodName           = "/v1.AlertService/GetAlert"
	AlertService_CountAlerts_FullMethodName        = "/v1.AlertService/CountAlerts"
	AlertService_ListAlerts_FullMethodName         = "/v1.AlertService/ListAlerts"
	AlertService_GetAlertsGroup_FullMethodName     = "/v1.AlertService/GetAlertsGroup"
	AlertService_GetAlertsCounts_FullMethodName    = "/v1.AlertService/GetAlertsCounts"
	AlertService_GetAlertTimeseries_FullMethodName = "/v1.AlertService/GetAlertTimeseries"
	AlertService_ResolveAlert_FullMethodName       = "/v1.AlertService/ResolveAlert"
	AlertService_ResolveAlerts_FullMethodName      = "/v1.AlertService/ResolveAlerts"
	AlertService_SnoozeAlert_FullMethodName        = "/v1.AlertService/SnoozeAlert"
	AlertService_DeleteAlerts_FullMethodName       = "/v1.AlertService/DeleteAlerts"
)

// AlertServiceClient is the client API for AlertService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertServiceClient interface {
	// GetAlert returns the alert given its id.
	GetAlert(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*storage.Alert, error)
	// CountAlerts counts how many alerts match the get request.
	CountAlerts(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*CountAlertsResponse, error)
	// List returns the slim list version of the alerts.
	ListAlerts(ctx context.Context, in *ListAlertsRequest, opts ...grpc.CallOption) (*ListAlertsResponse, error)
	// GetAlertsGroup returns alerts grouped by policy.
	GetAlertsGroup(ctx context.Context, in *ListAlertsRequest, opts ...grpc.CallOption) (*GetAlertsGroupResponse, error)
	// GetAlertsCounts returns the number of alerts in the requested cluster or category.
	GetAlertsCounts(ctx context.Context, in *GetAlertsCountsRequest, opts ...grpc.CallOption) (*GetAlertsCountsResponse, error)
	// GetAlertTimeseries returns the alerts sorted by time.
	GetAlertTimeseries(ctx context.Context, in *ListAlertsRequest, opts ...grpc.CallOption) (*GetAlertTimeseriesResponse, error)
	// ResolveAlert marks the given alert (by ID) as resolved.
	ResolveAlert(ctx context.Context, in *ResolveAlertRequest, opts ...grpc.CallOption) (*Empty, error)
	// ResolveAlertsByQuery marks alerts matching search query as resolved.
	ResolveAlerts(ctx context.Context, in *ResolveAlertsRequest, opts ...grpc.CallOption) (*Empty, error)
	// SnoozeAlert is deprecated.
	SnoozeAlert(ctx context.Context, in *SnoozeAlertRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteAlerts(ctx context.Context, in *DeleteAlertsRequest, opts ...grpc.CallOption) (*DeleteAlertsResponse, error)
}

type alertServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertServiceClient(cc grpc.ClientConnInterface) AlertServiceClient {
	return &alertServiceClient{cc}
}

func (c *alertServiceClient) GetAlert(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*storage.Alert, error) {
	out := new(storage.Alert)
	err := c.cc.Invoke(ctx, AlertService_GetAlert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) CountAlerts(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*CountAlertsResponse, error) {
	out := new(CountAlertsResponse)
	err := c.cc.Invoke(ctx, AlertService_CountAlerts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) ListAlerts(ctx context.Context, in *ListAlertsRequest, opts ...grpc.CallOption) (*ListAlertsResponse, error) {
	out := new(ListAlertsResponse)
	err := c.cc.Invoke(ctx, AlertService_ListAlerts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) GetAlertsGroup(ctx context.Context, in *ListAlertsRequest, opts ...grpc.CallOption) (*GetAlertsGroupResponse, error) {
	out := new(GetAlertsGroupResponse)
	err := c.cc.Invoke(ctx, AlertService_GetAlertsGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) GetAlertsCounts(ctx context.Context, in *GetAlertsCountsRequest, opts ...grpc.CallOption) (*GetAlertsCountsResponse, error) {
	out := new(GetAlertsCountsResponse)
	err := c.cc.Invoke(ctx, AlertService_GetAlertsCounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) GetAlertTimeseries(ctx context.Context, in *ListAlertsRequest, opts ...grpc.CallOption) (*GetAlertTimeseriesResponse, error) {
	out := new(GetAlertTimeseriesResponse)
	err := c.cc.Invoke(ctx, AlertService_GetAlertTimeseries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) ResolveAlert(ctx context.Context, in *ResolveAlertRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, AlertService_ResolveAlert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) ResolveAlerts(ctx context.Context, in *ResolveAlertsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, AlertService_ResolveAlerts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) SnoozeAlert(ctx context.Context, in *SnoozeAlertRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, AlertService_SnoozeAlert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) DeleteAlerts(ctx context.Context, in *DeleteAlertsRequest, opts ...grpc.CallOption) (*DeleteAlertsResponse, error) {
	out := new(DeleteAlertsResponse)
	err := c.cc.Invoke(ctx, AlertService_DeleteAlerts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertServiceServer is the server API for AlertService service.
// All implementations must embed UnimplementedAlertServiceServer
// for forward compatibility
type AlertServiceServer interface {
	// GetAlert returns the alert given its id.
	GetAlert(context.Context, *ResourceByID) (*storage.Alert, error)
	// CountAlerts counts how many alerts match the get request.
	CountAlerts(context.Context, *RawQuery) (*CountAlertsResponse, error)
	// List returns the slim list version of the alerts.
	ListAlerts(context.Context, *ListAlertsRequest) (*ListAlertsResponse, error)
	// GetAlertsGroup returns alerts grouped by policy.
	GetAlertsGroup(context.Context, *ListAlertsRequest) (*GetAlertsGroupResponse, error)
	// GetAlertsCounts returns the number of alerts in the requested cluster or category.
	GetAlertsCounts(context.Context, *GetAlertsCountsRequest) (*GetAlertsCountsResponse, error)
	// GetAlertTimeseries returns the alerts sorted by time.
	GetAlertTimeseries(context.Context, *ListAlertsRequest) (*GetAlertTimeseriesResponse, error)
	// ResolveAlert marks the given alert (by ID) as resolved.
	ResolveAlert(context.Context, *ResolveAlertRequest) (*Empty, error)
	// ResolveAlertsByQuery marks alerts matching search query as resolved.
	ResolveAlerts(context.Context, *ResolveAlertsRequest) (*Empty, error)
	// SnoozeAlert is deprecated.
	SnoozeAlert(context.Context, *SnoozeAlertRequest) (*Empty, error)
	DeleteAlerts(context.Context, *DeleteAlertsRequest) (*DeleteAlertsResponse, error)
	mustEmbedUnimplementedAlertServiceServer()
}

// UnimplementedAlertServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAlertServiceServer struct {
}

func (UnimplementedAlertServiceServer) GetAlert(context.Context, *ResourceByID) (*storage.Alert, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlert not implemented")
}
func (UnimplementedAlertServiceServer) CountAlerts(context.Context, *RawQuery) (*CountAlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAlerts not implemented")
}
func (UnimplementedAlertServiceServer) ListAlerts(context.Context, *ListAlertsRequest) (*ListAlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlerts not implemented")
}
func (UnimplementedAlertServiceServer) GetAlertsGroup(context.Context, *ListAlertsRequest) (*GetAlertsGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertsGroup not implemented")
}
func (UnimplementedAlertServiceServer) GetAlertsCounts(context.Context, *GetAlertsCountsRequest) (*GetAlertsCountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertsCounts not implemented")
}
func (UnimplementedAlertServiceServer) GetAlertTimeseries(context.Context, *ListAlertsRequest) (*GetAlertTimeseriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertTimeseries not implemented")
}
func (UnimplementedAlertServiceServer) ResolveAlert(context.Context, *ResolveAlertRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveAlert not implemented")
}
func (UnimplementedAlertServiceServer) ResolveAlerts(context.Context, *ResolveAlertsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveAlerts not implemented")
}
func (UnimplementedAlertServiceServer) SnoozeAlert(context.Context, *SnoozeAlertRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SnoozeAlert not implemented")
}
func (UnimplementedAlertServiceServer) DeleteAlerts(context.Context, *DeleteAlertsRequest) (*DeleteAlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlerts not implemented")
}
func (UnimplementedAlertServiceServer) mustEmbedUnimplementedAlertServiceServer() {}

// UnsafeAlertServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertServiceServer will
// result in compilation errors.
type UnsafeAlertServiceServer interface {
	mustEmbedUnimplementedAlertServiceServer()
}

func RegisterAlertServiceServer(s grpc.ServiceRegistrar, srv AlertServiceServer) {
	s.RegisterService(&AlertService_ServiceDesc, srv)
}

func _AlertService_GetAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).GetAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_GetAlert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).GetAlert(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_CountAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).CountAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_CountAlerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).CountAlerts(ctx, req.(*RawQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_ListAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).ListAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_ListAlerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).ListAlerts(ctx, req.(*ListAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_GetAlertsGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).GetAlertsGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_GetAlertsGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).GetAlertsGroup(ctx, req.(*ListAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_GetAlertsCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertsCountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).GetAlertsCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_GetAlertsCounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).GetAlertsCounts(ctx, req.(*GetAlertsCountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_GetAlertTimeseries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).GetAlertTimeseries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_GetAlertTimeseries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).GetAlertTimeseries(ctx, req.(*ListAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_ResolveAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).ResolveAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_ResolveAlert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).ResolveAlert(ctx, req.(*ResolveAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_ResolveAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).ResolveAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_ResolveAlerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).ResolveAlerts(ctx, req.(*ResolveAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_SnoozeAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnoozeAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).SnoozeAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_SnoozeAlert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).SnoozeAlert(ctx, req.(*SnoozeAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_DeleteAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).DeleteAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_DeleteAlerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).DeleteAlerts(ctx, req.(*DeleteAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AlertService_ServiceDesc is the grpc.ServiceDesc for AlertService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlertService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AlertService",
	HandlerType: (*AlertServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAlert",
			Handler:    _AlertService_GetAlert_Handler,
		},
		{
			MethodName: "CountAlerts",
			Handler:    _AlertService_CountAlerts_Handler,
		},
		{
			MethodName: "ListAlerts",
			Handler:    _AlertService_ListAlerts_Handler,
		},
		{
			MethodName: "GetAlertsGroup",
			Handler:    _AlertService_GetAlertsGroup_Handler,
		},
		{
			MethodName: "GetAlertsCounts",
			Handler:    _AlertService_GetAlertsCounts_Handler,
		},
		{
			MethodName: "GetAlertTimeseries",
			Handler:    _AlertService_GetAlertTimeseries_Handler,
		},
		{
			MethodName: "ResolveAlert",
			Handler:    _AlertService_ResolveAlert_Handler,
		},
		{
			MethodName: "ResolveAlerts",
			Handler:    _AlertService_ResolveAlerts_Handler,
		},
		{
			MethodName: "SnoozeAlert",
			Handler:    _AlertService_SnoozeAlert_Handler,
		},
		{
			MethodName: "DeleteAlerts",
			Handler:    _AlertService_DeleteAlerts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/alert_service.proto",
}
