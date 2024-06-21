// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v1/vuln_mgmt_service.proto

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
	VulnMgmtService_VulnMgmtExportWorkloads_FullMethodName = "/v1.VulnMgmtService/VulnMgmtExportWorkloads"
)

// VulnMgmtServiceClient is the client API for VulnMgmtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VulnMgmtServiceClient interface {
	// Streams vulnerability data upon request. Each entry consists of a deployment and the associated
	// container images. The response is structured as
	//
	//	{
	//	  {"result": {"deployment": {...}, "images": [...]}},
	//	  {"result": {"deployment": {...}, "images": [...]}},
	//	  ...
	//	}
	VulnMgmtExportWorkloads(ctx context.Context, in *VulnMgmtExportWorkloadsRequest, opts ...grpc.CallOption) (VulnMgmtService_VulnMgmtExportWorkloadsClient, error)
}

type vulnMgmtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVulnMgmtServiceClient(cc grpc.ClientConnInterface) VulnMgmtServiceClient {
	return &vulnMgmtServiceClient{cc}
}

func (c *vulnMgmtServiceClient) VulnMgmtExportWorkloads(ctx context.Context, in *VulnMgmtExportWorkloadsRequest, opts ...grpc.CallOption) (VulnMgmtService_VulnMgmtExportWorkloadsClient, error) {
	stream, err := c.cc.NewStream(ctx, &VulnMgmtService_ServiceDesc.Streams[0], VulnMgmtService_VulnMgmtExportWorkloads_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &vulnMgmtServiceVulnMgmtExportWorkloadsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VulnMgmtService_VulnMgmtExportWorkloadsClient interface {
	Recv() (*VulnMgmtExportWorkloadsResponse, error)
	grpc.ClientStream
}

type vulnMgmtServiceVulnMgmtExportWorkloadsClient struct {
	grpc.ClientStream
}

func (x *vulnMgmtServiceVulnMgmtExportWorkloadsClient) Recv() (*VulnMgmtExportWorkloadsResponse, error) {
	m := new(VulnMgmtExportWorkloadsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VulnMgmtServiceServer is the server API for VulnMgmtService service.
// All implementations should embed UnimplementedVulnMgmtServiceServer
// for forward compatibility
type VulnMgmtServiceServer interface {
	// Streams vulnerability data upon request. Each entry consists of a deployment and the associated
	// container images. The response is structured as
	//
	//	{
	//	  {"result": {"deployment": {...}, "images": [...]}},
	//	  {"result": {"deployment": {...}, "images": [...]}},
	//	  ...
	//	}
	VulnMgmtExportWorkloads(*VulnMgmtExportWorkloadsRequest, VulnMgmtService_VulnMgmtExportWorkloadsServer) error
}

// UnimplementedVulnMgmtServiceServer should be embedded to have forward compatible implementations.
type UnimplementedVulnMgmtServiceServer struct {
}

func (UnimplementedVulnMgmtServiceServer) VulnMgmtExportWorkloads(*VulnMgmtExportWorkloadsRequest, VulnMgmtService_VulnMgmtExportWorkloadsServer) error {
	return status.Errorf(codes.Unimplemented, "method VulnMgmtExportWorkloads not implemented")
}

// UnsafeVulnMgmtServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VulnMgmtServiceServer will
// result in compilation errors.
type UnsafeVulnMgmtServiceServer interface {
	mustEmbedUnimplementedVulnMgmtServiceServer()
}

func RegisterVulnMgmtServiceServer(s grpc.ServiceRegistrar, srv VulnMgmtServiceServer) {
	s.RegisterService(&VulnMgmtService_ServiceDesc, srv)
}

func _VulnMgmtService_VulnMgmtExportWorkloads_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(VulnMgmtExportWorkloadsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VulnMgmtServiceServer).VulnMgmtExportWorkloads(m, &vulnMgmtServiceVulnMgmtExportWorkloadsServer{stream})
}

type VulnMgmtService_VulnMgmtExportWorkloadsServer interface {
	Send(*VulnMgmtExportWorkloadsResponse) error
	grpc.ServerStream
}

type vulnMgmtServiceVulnMgmtExportWorkloadsServer struct {
	grpc.ServerStream
}

func (x *vulnMgmtServiceVulnMgmtExportWorkloadsServer) Send(m *VulnMgmtExportWorkloadsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// VulnMgmtService_ServiceDesc is the grpc.ServiceDesc for VulnMgmtService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VulnMgmtService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.VulnMgmtService",
	HandlerType: (*VulnMgmtServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "VulnMgmtExportWorkloads",
			Handler:       _VulnMgmtService_VulnMgmtExportWorkloads_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/v1/vuln_mgmt_service.proto",
}
