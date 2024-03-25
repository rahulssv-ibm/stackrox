// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: internalapi/sensor/cert_distribution_iservice.proto

package sensor

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
	CertDistributionService_FetchCertificate_FullMethodName = "/sensor.CertDistributionService/FetchCertificate"
)

// CertDistributionServiceClient is the client API for CertDistributionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertDistributionServiceClient interface {
	FetchCertificate(ctx context.Context, in *FetchCertificateRequest, opts ...grpc.CallOption) (*FetchCertificateResponse, error)
}

type certDistributionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertDistributionServiceClient(cc grpc.ClientConnInterface) CertDistributionServiceClient {
	return &certDistributionServiceClient{cc}
}

func (c *certDistributionServiceClient) FetchCertificate(ctx context.Context, in *FetchCertificateRequest, opts ...grpc.CallOption) (*FetchCertificateResponse, error) {
	out := new(FetchCertificateResponse)
	err := c.cc.Invoke(ctx, CertDistributionService_FetchCertificate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertDistributionServiceServer is the server API for CertDistributionService service.
// All implementations must embed UnimplementedCertDistributionServiceServer
// for forward compatibility
type CertDistributionServiceServer interface {
	FetchCertificate(context.Context, *FetchCertificateRequest) (*FetchCertificateResponse, error)
	mustEmbedUnimplementedCertDistributionServiceServer()
}

// UnimplementedCertDistributionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCertDistributionServiceServer struct {
}

func (UnimplementedCertDistributionServiceServer) FetchCertificate(context.Context, *FetchCertificateRequest) (*FetchCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCertificate not implemented")
}
func (UnimplementedCertDistributionServiceServer) mustEmbedUnimplementedCertDistributionServiceServer() {
}

// UnsafeCertDistributionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertDistributionServiceServer will
// result in compilation errors.
type UnsafeCertDistributionServiceServer interface {
	mustEmbedUnimplementedCertDistributionServiceServer()
}

func RegisterCertDistributionServiceServer(s grpc.ServiceRegistrar, srv CertDistributionServiceServer) {
	s.RegisterService(&CertDistributionService_ServiceDesc, srv)
}

func _CertDistributionService_FetchCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertDistributionServiceServer).FetchCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertDistributionService_FetchCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertDistributionServiceServer).FetchCertificate(ctx, req.(*FetchCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CertDistributionService_ServiceDesc is the grpc.ServiceDesc for CertDistributionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertDistributionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sensor.CertDistributionService",
	HandlerType: (*CertDistributionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchCertificate",
			Handler:    _CertDistributionService_FetchCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internalapi/sensor/cert_distribution_iservice.proto",
}
