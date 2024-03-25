// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/cve_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SuppressCVERequest struct {
	// These are (NVD) vulnerability identifiers, `cve` field of `storage.CVE`, and *not* the `id` field.
	// For example, CVE-2021-44832.
	Cves                 []string        `protobuf:"bytes,1,rep,name=cves,proto3" json:"cves,omitempty"`
	Duration             *types.Duration `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SuppressCVERequest) Reset()         { *m = SuppressCVERequest{} }
func (m *SuppressCVERequest) String() string { return proto.CompactTextString(m) }
func (*SuppressCVERequest) ProtoMessage()    {}
func (*SuppressCVERequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_257c48a61826e50f, []int{0}
}
func (m *SuppressCVERequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SuppressCVERequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SuppressCVERequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SuppressCVERequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuppressCVERequest.Merge(m, src)
}
func (m *SuppressCVERequest) XXX_Size() int {
	return m.Size()
}
func (m *SuppressCVERequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SuppressCVERequest.DiscardUnknown(m)
}

var xxx_messageInfo_SuppressCVERequest proto.InternalMessageInfo

func (m *SuppressCVERequest) GetCves() []string {
	if m != nil {
		return m.Cves
	}
	return nil
}

func (m *SuppressCVERequest) GetDuration() *types.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *SuppressCVERequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *SuppressCVERequest) Clone() *SuppressCVERequest {
	if m == nil {
		return nil
	}
	cloned := new(SuppressCVERequest)
	*cloned = *m

	if m.Cves != nil {
		cloned.Cves = make([]string, len(m.Cves))
		copy(cloned.Cves, m.Cves)
	}
	cloned.Duration = m.Duration.Clone()
	return cloned
}

type UnsuppressCVERequest struct {
	// These are (NVD) vulnerability identifiers, `cve` field of `storage.CVE`, and *not* the `id` field.
	// For example, CVE-2021-44832.
	Cves                 []string `protobuf:"bytes,1,rep,name=cves,proto3" json:"cves,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnsuppressCVERequest) Reset()         { *m = UnsuppressCVERequest{} }
func (m *UnsuppressCVERequest) String() string { return proto.CompactTextString(m) }
func (*UnsuppressCVERequest) ProtoMessage()    {}
func (*UnsuppressCVERequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_257c48a61826e50f, []int{1}
}
func (m *UnsuppressCVERequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnsuppressCVERequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnsuppressCVERequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnsuppressCVERequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsuppressCVERequest.Merge(m, src)
}
func (m *UnsuppressCVERequest) XXX_Size() int {
	return m.Size()
}
func (m *UnsuppressCVERequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsuppressCVERequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnsuppressCVERequest proto.InternalMessageInfo

func (m *UnsuppressCVERequest) GetCves() []string {
	if m != nil {
		return m.Cves
	}
	return nil
}

func (m *UnsuppressCVERequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *UnsuppressCVERequest) Clone() *UnsuppressCVERequest {
	if m == nil {
		return nil
	}
	cloned := new(UnsuppressCVERequest)
	*cloned = *m

	if m.Cves != nil {
		cloned.Cves = make([]string, len(m.Cves))
		copy(cloned.Cves, m.Cves)
	}
	return cloned
}

func init() {
	proto.RegisterType((*SuppressCVERequest)(nil), "v1.SuppressCVERequest")
	proto.RegisterType((*UnsuppressCVERequest)(nil), "v1.UnsuppressCVERequest")
}

func init() { proto.RegisterFile("api/v1/cve_service.proto", fileDescriptor_257c48a61826e50f) }

var fileDescriptor_257c48a61826e50f = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x2c, 0xc8, 0xd4,
	0x2f, 0x33, 0xd4, 0x4f, 0x2e, 0x4b, 0x8d, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x94, 0x12, 0x82, 0xca, 0xa6, 0xe6, 0x16, 0x94,
	0x54, 0x42, 0xc4, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x41, 0x52, 0x89, 0x79,
	0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x50, 0x59, 0x39, 0xa8, 0x2c, 0x98, 0x97,
	0x54, 0x9a, 0xa6, 0x9f, 0x52, 0x5a, 0x04, 0x56, 0x00, 0x91, 0x57, 0x4a, 0xe5, 0x12, 0x0a, 0x2e,
	0x2d, 0x28, 0x28, 0x4a, 0x2d, 0x2e, 0x76, 0x0e, 0x73, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x11, 0x12, 0xe2, 0x62, 0x49, 0x2e, 0x4b, 0x2d, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x0c, 0x02,
	0xb3, 0x85, 0x4c, 0xb9, 0x38, 0x60, 0x7a, 0x25, 0x98, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x24, 0xf5,
	0x20, 0x86, 0xeb, 0xc1, 0x0c, 0xd7, 0x73, 0x81, 0x2a, 0x08, 0x82, 0x2b, 0xf5, 0x62, 0xe1, 0x60,
	0x12, 0x60, 0x56, 0xd2, 0xe2, 0x12, 0x09, 0xcd, 0x2b, 0x26, 0xca, 0x22, 0xa3, 0xc3, 0x8c, 0x5c,
	0xfc, 0x9e, 0xb9, 0x89, 0xe9, 0xa9, 0xce, 0x61, 0xae, 0xc1, 0x90, 0x20, 0x10, 0x0a, 0xe1, 0xe2,
	0x41, 0x72, 0x66, 0xb1, 0x90, 0x98, 0x5e, 0x99, 0xa1, 0x1e, 0xa6, 0xc3, 0xa5, 0x38, 0x41, 0xe2,
	0xae, 0xa0, 0xd0, 0x51, 0x52, 0x6c, 0xba, 0xfc, 0x64, 0x32, 0x93, 0xb4, 0x91, 0x18, 0x28, 0xc0,
	0x32, 0x41, 0xe6, 0x81, 0x4c, 0xd7, 0x87, 0x39, 0xc2, 0x8a, 0x51, 0x4b, 0x28, 0x8a, 0x8b, 0x0f,
	0xc5, 0x55, 0xc5, 0x42, 0x12, 0x20, 0xfd, 0xd8, 0x5c, 0x8a, 0x6c, 0xb2, 0x32, 0xd8, 0x64, 0x59,
	0x23, 0x09, 0x54, 0x93, 0x4b, 0xf3, 0x90, 0xcc, 0x36, 0x3a, 0xc0, 0xc8, 0xc5, 0xe7, 0x97, 0x9f,
	0x82, 0xec, 0x89, 0x60, 0xd2, 0x3d, 0xa1, 0x00, 0xb6, 0x4a, 0xca, 0x48, 0x14, 0x64, 0x55, 0x5e,
	0x7e, 0x0a, 0xa6, 0x1f, 0x22, 0xc9, 0xf3, 0x83, 0x12, 0xd8, 0x60, 0x19, 0x23, 0x71, 0x14, 0x83,
	0x51, 0xbd, 0x70, 0x92, 0x91, 0x4b, 0xd0, 0x39, 0xa7, 0xb4, 0xb8, 0x24, 0xb5, 0x08, 0xc9, 0x17,
	0x61, 0xa4, 0xfb, 0x02, 0x25, 0xc0, 0x92, 0x21, 0x26, 0x62, 0x78, 0x24, 0x86, 0x3c, 0x8f, 0xa8,
	0x82, 0xcd, 0x96, 0x37, 0x92, 0x42, 0x37, 0x1b, 0xc5, 0x2f, 0x4e, 0x7a, 0x27, 0x1e, 0xc9, 0x31,
	0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x5c, 0x12, 0x99,
	0xf9, 0x7a, 0xc5, 0x25, 0x89, 0xc9, 0xd9, 0x45, 0xf9, 0x15, 0x90, 0x14, 0xac, 0x97, 0x58, 0x90,
	0xa9, 0x57, 0x66, 0x18, 0xc5, 0x54, 0x66, 0x18, 0xc1, 0x98, 0xc4, 0x06, 0x16, 0x33, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x46, 0x8e, 0xdc, 0x9a, 0x92, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ImageCVEServiceClient is the client API for ImageCVEService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
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
	err := c.cc.Invoke(ctx, "/v1.ImageCVEService/SuppressCVEs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageCVEServiceClient) UnsuppressCVEs(ctx context.Context, in *UnsuppressCVERequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/v1.ImageCVEService/UnsuppressCVEs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageCVEServiceServer is the server API for ImageCVEService service.
type ImageCVEServiceServer interface {
	// SuppressCVE suppresses image cves.
	SuppressCVEs(context.Context, *SuppressCVERequest) (*Empty, error)
	// UnsuppressCVE unsuppresses image cves.
	UnsuppressCVEs(context.Context, *UnsuppressCVERequest) (*Empty, error)
}

// UnimplementedImageCVEServiceServer can be embedded to have forward compatible implementations.
type UnimplementedImageCVEServiceServer struct {
}

func (*UnimplementedImageCVEServiceServer) SuppressCVEs(ctx context.Context, req *SuppressCVERequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuppressCVEs not implemented")
}
func (*UnimplementedImageCVEServiceServer) UnsuppressCVEs(ctx context.Context, req *UnsuppressCVERequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsuppressCVEs not implemented")
}

func RegisterImageCVEServiceServer(s *grpc.Server, srv ImageCVEServiceServer) {
	s.RegisterService(&_ImageCVEService_serviceDesc, srv)
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
		FullMethod: "/v1.ImageCVEService/SuppressCVEs",
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
		FullMethod: "/v1.ImageCVEService/UnsuppressCVEs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageCVEServiceServer).UnsuppressCVEs(ctx, req.(*UnsuppressCVERequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImageCVEService_serviceDesc = grpc.ServiceDesc{
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

// NodeCVEServiceClient is the client API for NodeCVEService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
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
	err := c.cc.Invoke(ctx, "/v1.NodeCVEService/SuppressCVEs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCVEServiceClient) UnsuppressCVEs(ctx context.Context, in *UnsuppressCVERequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/v1.NodeCVEService/UnsuppressCVEs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeCVEServiceServer is the server API for NodeCVEService service.
type NodeCVEServiceServer interface {
	// SuppressCVE suppresses node cves.
	SuppressCVEs(context.Context, *SuppressCVERequest) (*Empty, error)
	// UnsuppressCVE unsuppresses node cves.
	UnsuppressCVEs(context.Context, *UnsuppressCVERequest) (*Empty, error)
}

// UnimplementedNodeCVEServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeCVEServiceServer struct {
}

func (*UnimplementedNodeCVEServiceServer) SuppressCVEs(ctx context.Context, req *SuppressCVERequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuppressCVEs not implemented")
}
func (*UnimplementedNodeCVEServiceServer) UnsuppressCVEs(ctx context.Context, req *UnsuppressCVERequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsuppressCVEs not implemented")
}

func RegisterNodeCVEServiceServer(s *grpc.Server, srv NodeCVEServiceServer) {
	s.RegisterService(&_NodeCVEService_serviceDesc, srv)
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
		FullMethod: "/v1.NodeCVEService/SuppressCVEs",
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
		FullMethod: "/v1.NodeCVEService/UnsuppressCVEs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCVEServiceServer).UnsuppressCVEs(ctx, req.(*UnsuppressCVERequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeCVEService_serviceDesc = grpc.ServiceDesc{
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

// ClusterCVEServiceClient is the client API for ClusterCVEService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
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
	err := c.cc.Invoke(ctx, "/v1.ClusterCVEService/SuppressCVEs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterCVEServiceClient) UnsuppressCVEs(ctx context.Context, in *UnsuppressCVERequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/v1.ClusterCVEService/UnsuppressCVEs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterCVEServiceServer is the server API for ClusterCVEService service.
type ClusterCVEServiceServer interface {
	// SuppressCVE suppresses cluster cves.
	SuppressCVEs(context.Context, *SuppressCVERequest) (*Empty, error)
	// UnsuppressCVE unsuppresses cluster cves.
	UnsuppressCVEs(context.Context, *UnsuppressCVERequest) (*Empty, error)
}

// UnimplementedClusterCVEServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClusterCVEServiceServer struct {
}

func (*UnimplementedClusterCVEServiceServer) SuppressCVEs(ctx context.Context, req *SuppressCVERequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuppressCVEs not implemented")
}
func (*UnimplementedClusterCVEServiceServer) UnsuppressCVEs(ctx context.Context, req *UnsuppressCVERequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsuppressCVEs not implemented")
}

func RegisterClusterCVEServiceServer(s *grpc.Server, srv ClusterCVEServiceServer) {
	s.RegisterService(&_ClusterCVEService_serviceDesc, srv)
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
		FullMethod: "/v1.ClusterCVEService/SuppressCVEs",
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
		FullMethod: "/v1.ClusterCVEService/UnsuppressCVEs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterCVEServiceServer).UnsuppressCVEs(ctx, req.(*UnsuppressCVERequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterCVEService_serviceDesc = grpc.ServiceDesc{
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

func (m *SuppressCVERequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SuppressCVERequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SuppressCVERequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Duration != nil {
		{
			size, err := m.Duration.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCveService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Cves) > 0 {
		for iNdEx := len(m.Cves) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Cves[iNdEx])
			copy(dAtA[i:], m.Cves[iNdEx])
			i = encodeVarintCveService(dAtA, i, uint64(len(m.Cves[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *UnsuppressCVERequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnsuppressCVERequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnsuppressCVERequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Cves) > 0 {
		for iNdEx := len(m.Cves) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Cves[iNdEx])
			copy(dAtA[i:], m.Cves[iNdEx])
			i = encodeVarintCveService(dAtA, i, uint64(len(m.Cves[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCveService(dAtA []byte, offset int, v uint64) int {
	offset -= sovCveService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SuppressCVERequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Cves) > 0 {
		for _, s := range m.Cves {
			l = len(s)
			n += 1 + l + sovCveService(uint64(l))
		}
	}
	if m.Duration != nil {
		l = m.Duration.Size()
		n += 1 + l + sovCveService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UnsuppressCVERequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Cves) > 0 {
		for _, s := range m.Cves {
			l = len(s)
			n += 1 + l + sovCveService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCveService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCveService(x uint64) (n int) {
	return sovCveService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SuppressCVERequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCveService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SuppressCVERequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SuppressCVERequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cves", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCveService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCveService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCveService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cves = append(m.Cves, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCveService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCveService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCveService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Duration == nil {
				m.Duration = &types.Duration{}
			}
			if err := m.Duration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCveService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCveService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UnsuppressCVERequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCveService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UnsuppressCVERequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnsuppressCVERequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cves", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCveService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCveService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCveService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cves = append(m.Cves, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCveService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCveService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCveService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCveService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCveService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCveService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCveService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCveService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCveService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCveService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCveService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCveService = fmt.Errorf("proto: unexpected end of group")
)
