// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/pod_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	storage "github.com/stackrox/rox/generated/storage"
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

type PodsResponse struct {
	Pods                 []*storage.Pod `protobuf:"bytes,1,rep,name=pods,proto3" json:"pods,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PodsResponse) Reset()         { *m = PodsResponse{} }
func (m *PodsResponse) String() string { return proto.CompactTextString(m) }
func (*PodsResponse) ProtoMessage()    {}
func (*PodsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5092c07627247a99, []int{0}
}
func (m *PodsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PodsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PodsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PodsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodsResponse.Merge(m, src)
}
func (m *PodsResponse) XXX_Size() int {
	return m.Size()
}
func (m *PodsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PodsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PodsResponse proto.InternalMessageInfo

func (m *PodsResponse) GetPods() []*storage.Pod {
	if m != nil {
		return m.Pods
	}
	return nil
}

func (m *PodsResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *PodsResponse) Clone() *PodsResponse {
	if m == nil {
		return nil
	}
	cloned := new(PodsResponse)
	*cloned = *m

	if m.Pods != nil {
		cloned.Pods = make([]*storage.Pod, len(m.Pods))
		for idx, v := range m.Pods {
			cloned.Pods[idx] = v.Clone()
		}
	}
	return cloned
}

type ExportPodRequest struct {
	Timeout              int32    `protobuf:"varint,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Query                string   `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportPodRequest) Reset()         { *m = ExportPodRequest{} }
func (m *ExportPodRequest) String() string { return proto.CompactTextString(m) }
func (*ExportPodRequest) ProtoMessage()    {}
func (*ExportPodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5092c07627247a99, []int{1}
}
func (m *ExportPodRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExportPodRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExportPodRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExportPodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportPodRequest.Merge(m, src)
}
func (m *ExportPodRequest) XXX_Size() int {
	return m.Size()
}
func (m *ExportPodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportPodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportPodRequest proto.InternalMessageInfo

func (m *ExportPodRequest) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *ExportPodRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *ExportPodRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ExportPodRequest) Clone() *ExportPodRequest {
	if m == nil {
		return nil
	}
	cloned := new(ExportPodRequest)
	*cloned = *m

	return cloned
}

type ExportPodResponse struct {
	Pod                  *storage.Pod `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ExportPodResponse) Reset()         { *m = ExportPodResponse{} }
func (m *ExportPodResponse) String() string { return proto.CompactTextString(m) }
func (*ExportPodResponse) ProtoMessage()    {}
func (*ExportPodResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5092c07627247a99, []int{2}
}
func (m *ExportPodResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExportPodResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExportPodResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExportPodResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportPodResponse.Merge(m, src)
}
func (m *ExportPodResponse) XXX_Size() int {
	return m.Size()
}
func (m *ExportPodResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportPodResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportPodResponse proto.InternalMessageInfo

func (m *ExportPodResponse) GetPod() *storage.Pod {
	if m != nil {
		return m.Pod
	}
	return nil
}

func (m *ExportPodResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ExportPodResponse) Clone() *ExportPodResponse {
	if m == nil {
		return nil
	}
	cloned := new(ExportPodResponse)
	*cloned = *m

	cloned.Pod = m.Pod.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*PodsResponse)(nil), "v1.PodsResponse")
	proto.RegisterType((*ExportPodRequest)(nil), "v1.ExportPodRequest")
	proto.RegisterType((*ExportPodResponse)(nil), "v1.ExportPodResponse")
}

func init() { proto.RegisterFile("api/v1/pod_service.proto", fileDescriptor_5092c07627247a99) }

var fileDescriptor_5092c07627247a99 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4e, 0xf2, 0x40,
	0x10, 0xc7, 0xbf, 0x2d, 0x1f, 0x1f, 0x9f, 0x2b, 0x89, 0xb0, 0xc1, 0xd8, 0x54, 0xd3, 0x34, 0x3d,
	0x71, 0xda, 0x52, 0x38, 0x7a, 0x23, 0x31, 0x5e, 0xb1, 0x7a, 0x30, 0x5e, 0xcc, 0x4a, 0x27, 0xd8,
	0x08, 0x9d, 0xa5, 0xbb, 0x54, 0xb8, 0xfa, 0x0a, 0x5e, 0xbc, 0xf9, 0x3a, 0x1e, 0x4d, 0x7c, 0x01,
	0x83, 0x3e, 0x88, 0xd9, 0xb6, 0x1a, 0xe4, 0x38, 0xf3, 0xdb, 0xf9, 0xcd, 0xfe, 0x87, 0xda, 0x42,
	0x26, 0x41, 0x1e, 0x06, 0x12, 0xe3, 0x6b, 0x05, 0x59, 0x9e, 0x8c, 0x81, 0xcb, 0x0c, 0x35, 0x32,
	0x2b, 0x0f, 0x9d, 0xc3, 0x8a, 0x2a, 0x10, 0xd9, 0xf8, 0xf6, 0xf7, 0x03, 0xe7, 0x68, 0x82, 0x38,
	0x99, 0x42, 0x60, 0xde, 0x88, 0x34, 0x45, 0x2d, 0x74, 0x82, 0xa9, 0xaa, 0xa8, 0xad, 0x34, 0x66,
	0x62, 0x02, 0x41, 0x0c, 0x72, 0x8a, 0xab, 0x19, 0xa4, 0xba, 0x24, 0x7e, 0x8f, 0x36, 0x47, 0x18,
	0xab, 0x08, 0x94, 0xc4, 0x54, 0x01, 0xf3, 0xe8, 0x5f, 0x89, 0xb1, 0xb2, 0x89, 0x57, 0xeb, 0xee,
	0xf6, 0x9b, 0xbc, 0x1a, 0xe4, 0x23, 0x8c, 0xa3, 0x82, 0xf8, 0x43, 0xda, 0x3a, 0x59, 0x4a, 0xcc,
	0xb4, 0x69, 0xc1, 0x7c, 0x01, 0x4a, 0x33, 0x9b, 0x36, 0x74, 0x32, 0x03, 0x5c, 0x68, 0x9b, 0x78,
	0xa4, 0x5b, 0x8f, 0xbe, 0x4b, 0xd6, 0xa1, 0xf5, 0xf9, 0x02, 0xb2, 0x95, 0x6d, 0x79, 0xa4, 0xbb,
	0x13, 0x95, 0x85, 0x3f, 0xa0, 0xed, 0x0d, 0x47, 0xb5, 0xda, 0xa5, 0x35, 0x89, 0x71, 0x21, 0xd8,
	0xde, 0x6c, 0x40, 0xff, 0x99, 0x50, 0x3a, 0xc2, 0xf8, 0xbc, 0xcc, 0xcd, 0x8e, 0x69, 0xe3, 0x14,
	0x8c, 0x40, 0xb1, 0x26, 0xcf, 0x43, 0x1e, 0x89, 0xfb, 0x33, 0x23, 0x77, 0x5a, 0xa6, 0xda, 0x0c,
	0xe5, 0xb7, 0x1e, 0xde, 0x3e, 0x1f, 0x2d, 0xca, 0xfe, 0x57, 0xc7, 0x55, 0xec, 0x82, 0xd2, 0x9f,
	0x0f, 0x28, 0xd6, 0x31, 0x13, 0xdb, 0xa1, 0x9c, 0xfd, 0xad, 0x6e, 0x25, 0x3b, 0x28, 0x64, 0x6d,
	0xb6, 0x67, 0x64, 0x50, 0xe0, 0xc2, 0xd9, 0x23, 0x43, 0xfe, 0xb2, 0x76, 0xc9, 0xeb, 0xda, 0x25,
	0xef, 0x6b, 0x97, 0x3c, 0x7d, 0xb8, 0x7f, 0xa8, 0x9d, 0x20, 0x57, 0x5a, 0x8c, 0xef, 0x32, 0x5c,
	0x96, 0x17, 0xe7, 0x42, 0x26, 0x3c, 0x0f, 0xaf, 0xac, 0x3c, 0xbc, 0x24, 0x37, 0xff, 0x8a, 0xde,
	0xe0, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x34, 0x96, 0xdb, 0xfa, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PodServiceClient is the client API for PodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
type PodServiceClient interface {
	// GetPods returns the pods.
	GetPods(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*PodsResponse, error)
	ExportPods(ctx context.Context, in *ExportPodRequest, opts ...grpc.CallOption) (PodService_ExportPodsClient, error)
}

type podServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPodServiceClient(cc grpc.ClientConnInterface) PodServiceClient {
	return &podServiceClient{cc}
}

func (c *podServiceClient) GetPods(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*PodsResponse, error) {
	out := new(PodsResponse)
	err := c.cc.Invoke(ctx, "/v1.PodService/GetPods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podServiceClient) ExportPods(ctx context.Context, in *ExportPodRequest, opts ...grpc.CallOption) (PodService_ExportPodsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PodService_serviceDesc.Streams[0], "/v1.PodService/ExportPods", opts...)
	if err != nil {
		return nil, err
	}
	x := &podServiceExportPodsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PodService_ExportPodsClient interface {
	Recv() (*ExportPodResponse, error)
	grpc.ClientStream
}

type podServiceExportPodsClient struct {
	grpc.ClientStream
}

func (x *podServiceExportPodsClient) Recv() (*ExportPodResponse, error) {
	m := new(ExportPodResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PodServiceServer is the server API for PodService service.
type PodServiceServer interface {
	// GetPods returns the pods.
	GetPods(context.Context, *RawQuery) (*PodsResponse, error)
	ExportPods(*ExportPodRequest, PodService_ExportPodsServer) error
}

// UnimplementedPodServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPodServiceServer struct {
}

func (*UnimplementedPodServiceServer) GetPods(ctx context.Context, req *RawQuery) (*PodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPods not implemented")
}
func (*UnimplementedPodServiceServer) ExportPods(req *ExportPodRequest, srv PodService_ExportPodsServer) error {
	return status.Errorf(codes.Unimplemented, "method ExportPods not implemented")
}

func RegisterPodServiceServer(s *grpc.Server, srv PodServiceServer) {
	s.RegisterService(&_PodService_serviceDesc, srv)
}

func _PodService_GetPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodServiceServer).GetPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PodService/GetPods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodServiceServer).GetPods(ctx, req.(*RawQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodService_ExportPods_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExportPodRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PodServiceServer).ExportPods(m, &podServiceExportPodsServer{stream})
}

type PodService_ExportPodsServer interface {
	Send(*ExportPodResponse) error
	grpc.ServerStream
}

type podServiceExportPodsServer struct {
	grpc.ServerStream
}

func (x *podServiceExportPodsServer) Send(m *ExportPodResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _PodService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PodService",
	HandlerType: (*PodServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPods",
			Handler:    _PodService_GetPods_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExportPods",
			Handler:       _PodService_ExportPods_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/v1/pod_service.proto",
}

func (m *PodsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PodsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Pods) > 0 {
		for iNdEx := len(m.Pods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPodService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ExportPodRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportPodRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExportPodRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Query) > 0 {
		i -= len(m.Query)
		copy(dAtA[i:], m.Query)
		i = encodeVarintPodService(dAtA, i, uint64(len(m.Query)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timeout != 0 {
		i = encodeVarintPodService(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ExportPodResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportPodResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExportPodResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Pod != nil {
		{
			size, err := m.Pod.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPodService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPodService(dAtA []byte, offset int, v uint64) int {
	offset -= sovPodService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PodsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pods) > 0 {
		for _, e := range m.Pods {
			l = e.Size()
			n += 1 + l + sovPodService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ExportPodRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timeout != 0 {
		n += 1 + sovPodService(uint64(m.Timeout))
	}
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovPodService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ExportPodResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pod != nil {
		l = m.Pod.Size()
		n += 1 + l + sovPodService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPodService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPodService(x uint64) (n int) {
	return sovPodService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PodsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPodService
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
			return fmt.Errorf("proto: PodsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPodService
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
				return ErrInvalidLengthPodService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPodService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pods = append(m.Pods, &storage.Pod{})
			if err := m.Pods[len(m.Pods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPodService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPodService
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
func (m *ExportPodRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPodService
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
			return fmt.Errorf("proto: ExportPodRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportPodRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPodService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPodService
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
				return ErrInvalidLengthPodService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPodService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPodService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPodService
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
func (m *ExportPodResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPodService
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
			return fmt.Errorf("proto: ExportPodResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportPodResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPodService
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
				return ErrInvalidLengthPodService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPodService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pod == nil {
				m.Pod = &storage.Pod{}
			}
			if err := m.Pod.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPodService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPodService
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
func skipPodService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPodService
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
					return 0, ErrIntOverflowPodService
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
					return 0, ErrIntOverflowPodService
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
				return 0, ErrInvalidLengthPodService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPodService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPodService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPodService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPodService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPodService = fmt.Errorf("proto: unexpected end of group")
)
