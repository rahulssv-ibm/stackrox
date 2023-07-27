// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/usage_service.proto

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

type UsageRequest struct {
	From                 *types.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   *types.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UsageRequest) Reset()         { *m = UsageRequest{} }
func (m *UsageRequest) String() string { return proto.CompactTextString(m) }
func (*UsageRequest) ProtoMessage()    {}
func (*UsageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_13c5d48ec6b04754, []int{0}
}
func (m *UsageRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UsageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UsageRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UsageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsageRequest.Merge(m, src)
}
func (m *UsageRequest) XXX_Size() int {
	return m.Size()
}
func (m *UsageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UsageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UsageRequest proto.InternalMessageInfo

func (m *UsageRequest) GetFrom() *types.Timestamp {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *UsageRequest) GetTo() *types.Timestamp {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *UsageRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *UsageRequest) Clone() *UsageRequest {
	if m == nil {
		return nil
	}
	cloned := new(UsageRequest)
	*cloned = *m

	cloned.From = m.From.Clone()
	cloned.To = m.To.Clone()
	return cloned
}

type CurrentUsageResponse struct {
	Timestamp            *types.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	NumNodes             int32            `protobuf:"varint,2,opt,name=num_nodes,json=numNodes,proto3" json:"num_nodes,omitempty"`
	NumCpuUnits          int32            `protobuf:"varint,3,opt,name=num_cpu_units,json=numCpuUnits,proto3" json:"num_cpu_units,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CurrentUsageResponse) Reset()         { *m = CurrentUsageResponse{} }
func (m *CurrentUsageResponse) String() string { return proto.CompactTextString(m) }
func (*CurrentUsageResponse) ProtoMessage()    {}
func (*CurrentUsageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_13c5d48ec6b04754, []int{1}
}
func (m *CurrentUsageResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CurrentUsageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CurrentUsageResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CurrentUsageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentUsageResponse.Merge(m, src)
}
func (m *CurrentUsageResponse) XXX_Size() int {
	return m.Size()
}
func (m *CurrentUsageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentUsageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentUsageResponse proto.InternalMessageInfo

func (m *CurrentUsageResponse) GetTimestamp() *types.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *CurrentUsageResponse) GetNumNodes() int32 {
	if m != nil {
		return m.NumNodes
	}
	return 0
}

func (m *CurrentUsageResponse) GetNumCpuUnits() int32 {
	if m != nil {
		return m.NumCpuUnits
	}
	return 0
}

func (m *CurrentUsageResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *CurrentUsageResponse) Clone() *CurrentUsageResponse {
	if m == nil {
		return nil
	}
	cloned := new(CurrentUsageResponse)
	*cloned = *m

	cloned.Timestamp = m.Timestamp.Clone()
	return cloned
}

type MaxUsageResponse struct {
	MaxNodesAt           *types.Timestamp `protobuf:"bytes,1,opt,name=max_nodes_at,json=maxNodesAt,proto3" json:"max_nodes_at,omitempty"`
	MaxNodes             int32            `protobuf:"varint,2,opt,name=max_nodes,json=maxNodes,proto3" json:"max_nodes,omitempty"`
	MaxCpuUnitsAt        *types.Timestamp `protobuf:"bytes,3,opt,name=max_cpu_units_at,json=maxCpuUnitsAt,proto3" json:"max_cpu_units_at,omitempty"`
	MaxCpuUnits          int32            `protobuf:"varint,4,opt,name=max_cpu_units,json=maxCpuUnits,proto3" json:"max_cpu_units,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MaxUsageResponse) Reset()         { *m = MaxUsageResponse{} }
func (m *MaxUsageResponse) String() string { return proto.CompactTextString(m) }
func (*MaxUsageResponse) ProtoMessage()    {}
func (*MaxUsageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_13c5d48ec6b04754, []int{2}
}
func (m *MaxUsageResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MaxUsageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MaxUsageResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MaxUsageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaxUsageResponse.Merge(m, src)
}
func (m *MaxUsageResponse) XXX_Size() int {
	return m.Size()
}
func (m *MaxUsageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MaxUsageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MaxUsageResponse proto.InternalMessageInfo

func (m *MaxUsageResponse) GetMaxNodesAt() *types.Timestamp {
	if m != nil {
		return m.MaxNodesAt
	}
	return nil
}

func (m *MaxUsageResponse) GetMaxNodes() int32 {
	if m != nil {
		return m.MaxNodes
	}
	return 0
}

func (m *MaxUsageResponse) GetMaxCpuUnitsAt() *types.Timestamp {
	if m != nil {
		return m.MaxCpuUnitsAt
	}
	return nil
}

func (m *MaxUsageResponse) GetMaxCpuUnits() int32 {
	if m != nil {
		return m.MaxCpuUnits
	}
	return 0
}

func (m *MaxUsageResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *MaxUsageResponse) Clone() *MaxUsageResponse {
	if m == nil {
		return nil
	}
	cloned := new(MaxUsageResponse)
	*cloned = *m

	cloned.MaxNodesAt = m.MaxNodesAt.Clone()
	cloned.MaxCpuUnitsAt = m.MaxCpuUnitsAt.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*UsageRequest)(nil), "v1.UsageRequest")
	proto.RegisterType((*CurrentUsageResponse)(nil), "v1.CurrentUsageResponse")
	proto.RegisterType((*MaxUsageResponse)(nil), "v1.MaxUsageResponse")
}

func init() { proto.RegisterFile("api/v1/usage_service.proto", fileDescriptor_13c5d48ec6b04754) }

var fileDescriptor_13c5d48ec6b04754 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x75, 0xd2, 0x2a, 0x76, 0xda, 0xd2, 0x3a, 0x56, 0x88, 0x51, 0xa2, 0x64, 0x25, 0x2e, 0x12,
	0x52, 0x37, 0x2e, 0xdc, 0xd4, 0x22, 0xdd, 0xa8, 0x60, 0xb5, 0x20, 0x6e, 0xc2, 0xb4, 0x4e, 0x4b,
	0xc4, 0x99, 0x89, 0x99, 0x3b, 0x21, 0x6e, 0xfd, 0x07, 0xe2, 0xc6, 0x7f, 0xe1, 0xdf, 0x70, 0x29,
	0x08, 0xae, 0xa5, 0xef, 0xfd, 0x90, 0xc7, 0xe4, 0xab, 0xed, 0xe3, 0x41, 0xbb, 0x9b, 0xb9, 0xe7,
	0xdc, 0x73, 0xcf, 0xdc, 0x39, 0xd8, 0xa1, 0x49, 0x1c, 0x64, 0x61, 0xa0, 0x15, 0xdd, 0xb0, 0x48,
	0xb1, 0x34, 0x8b, 0x57, 0xcc, 0x4f, 0x52, 0x09, 0x92, 0x58, 0x59, 0xe8, 0x90, 0x0a, 0x67, 0x3c,
	0x81, 0xaf, 0x65, 0xdd, 0xb9, 0xbf, 0x91, 0x72, 0xf3, 0x99, 0x05, 0x06, 0xa2, 0x42, 0x48, 0xa0,
	0x10, 0x4b, 0xa1, 0x2a, 0xf4, 0x41, 0x85, 0x16, 0xb7, 0xa5, 0x5e, 0x07, 0x10, 0x73, 0xa6, 0x80,
	0xf2, 0xa4, 0x24, 0x78, 0x9f, 0x70, 0x6f, 0x61, 0xa6, 0xcd, 0xd9, 0x17, 0xcd, 0x14, 0x10, 0x1f,
	0xb7, 0xd7, 0xa9, 0xe4, 0x36, 0x7a, 0x88, 0x1e, 0x75, 0xc7, 0x8e, 0x5f, 0xf6, 0xfb, 0x75, 0xbf,
	0xff, 0xae, 0xee, 0x9f, 0x17, 0x3c, 0xf2, 0x18, 0x5b, 0x20, 0x6d, 0xeb, 0x28, 0xdb, 0x02, 0xe9,
	0x7d, 0x47, 0x78, 0x34, 0xd5, 0x69, 0xca, 0x04, 0x54, 0x33, 0x55, 0x22, 0x85, 0x62, 0xe4, 0x29,
	0xee, 0x34, 0xbe, 0x4e, 0x98, 0xbc, 0x23, 0x93, 0x7b, 0xb8, 0x23, 0x34, 0x8f, 0x84, 0xfc, 0xc8,
	0x54, 0xe1, 0xe2, 0xfa, 0xfc, 0xa6, 0xd0, 0xfc, 0xb5, 0xb9, 0x13, 0x0f, 0xf7, 0x0d, 0xb8, 0x4a,
	0x74, 0xa4, 0x45, 0x0c, 0xca, 0x6e, 0x15, 0x84, 0xae, 0xd0, 0x7c, 0x9a, 0xe8, 0x85, 0x29, 0x79,
	0xff, 0x10, 0x1e, 0xbe, 0xa2, 0xf9, 0xa1, 0x9f, 0x67, 0xb8, 0xc7, 0x69, 0x5e, 0xaa, 0x46, 0x14,
	0x4e, 0xb0, 0x84, 0x39, 0xcd, 0x8b, 0xa1, 0x13, 0x30, 0x9e, 0x9a, 0xee, 0xda, 0x53, 0x0d, 0x93,
	0x29, 0x1e, 0x1a, 0xb0, 0xf1, 0x64, 0xe4, 0x5b, 0x47, 0xe5, 0xfb, 0x9c, 0xe6, 0xb5, 0xe5, 0x09,
	0x98, 0x87, 0x1d, 0x88, 0xd8, 0xed, 0xf2, 0x61, 0x7b, 0xac, 0xf1, 0x2f, 0x54, 0xfd, 0xec, 0xdb,
	0x32, 0x46, 0xe4, 0x0d, 0x1e, 0xcc, 0x18, 0xec, 0xef, 0x9f, 0x74, 0xfc, 0x2c, 0xf4, 0x5f, 0x98,
	0x30, 0x39, 0xb6, 0x39, 0x5e, 0xf5, 0x39, 0xde, 0xdd, 0x6f, 0x7f, 0xcf, 0x7f, 0x58, 0xb7, 0xc9,
	0xad, 0x26, 0x99, 0xc1, 0xaa, 0xe4, 0x91, 0x97, 0xb8, 0x3b, 0x63, 0x50, 0xaf, 0x8f, 0x0c, 0x8d,
	0xc6, 0x7e, 0x9a, 0x9c, 0x91, 0xa9, 0x5c, 0x5e, 0xaf, 0x77, 0xa7, 0x50, 0x1c, 0x90, 0xfe, 0x4e,
	0x91, 0xd3, 0xfc, 0xf9, 0xe8, 0xf7, 0xd6, 0x45, 0x7f, 0xb6, 0x2e, 0xfa, 0xbf, 0x75, 0xd1, 0xcf,
	0x33, 0xf7, 0xda, 0x07, 0x2b, 0x0b, 0xdf, 0xa3, 0xe5, 0x8d, 0x62, 0x21, 0x4f, 0x2e, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xe3, 0x0e, 0x8d, 0x1d, 0x1e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UsageServiceClient is the client API for UsageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
type UsageServiceClient interface {
	// GetCurrentUsage
	//
	// Returns current usage, with about 5 minutes precision.
	GetCurrentUsage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CurrentUsageResponse, error)
	// GetMaxUsage
	//
	// Returns maximum, i.e. peak, usage for the given time frame together
	// with the time when this maximum was observed.
	GetMaxUsage(ctx context.Context, in *UsageRequest, opts ...grpc.CallOption) (*MaxUsageResponse, error)
}

type usageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsageServiceClient(cc grpc.ClientConnInterface) UsageServiceClient {
	return &usageServiceClient{cc}
}

func (c *usageServiceClient) GetCurrentUsage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CurrentUsageResponse, error) {
	out := new(CurrentUsageResponse)
	err := c.cc.Invoke(ctx, "/v1.UsageService/GetCurrentUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usageServiceClient) GetMaxUsage(ctx context.Context, in *UsageRequest, opts ...grpc.CallOption) (*MaxUsageResponse, error) {
	out := new(MaxUsageResponse)
	err := c.cc.Invoke(ctx, "/v1.UsageService/GetMaxUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsageServiceServer is the server API for UsageService service.
type UsageServiceServer interface {
	// GetCurrentUsage
	//
	// Returns current usage, with about 5 minutes precision.
	GetCurrentUsage(context.Context, *Empty) (*CurrentUsageResponse, error)
	// GetMaxUsage
	//
	// Returns maximum, i.e. peak, usage for the given time frame together
	// with the time when this maximum was observed.
	GetMaxUsage(context.Context, *UsageRequest) (*MaxUsageResponse, error)
}

// UnimplementedUsageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUsageServiceServer struct {
}

func (*UnimplementedUsageServiceServer) GetCurrentUsage(ctx context.Context, req *Empty) (*CurrentUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUsage not implemented")
}
func (*UnimplementedUsageServiceServer) GetMaxUsage(ctx context.Context, req *UsageRequest) (*MaxUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMaxUsage not implemented")
}

func RegisterUsageServiceServer(s *grpc.Server, srv UsageServiceServer) {
	s.RegisterService(&_UsageService_serviceDesc, srv)
}

func _UsageService_GetCurrentUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsageServiceServer).GetCurrentUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UsageService/GetCurrentUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsageServiceServer).GetCurrentUsage(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsageService_GetMaxUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsageServiceServer).GetMaxUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UsageService/GetMaxUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsageServiceServer).GetMaxUsage(ctx, req.(*UsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UsageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UsageService",
	HandlerType: (*UsageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentUsage",
			Handler:    _UsageService_GetCurrentUsage_Handler,
		},
		{
			MethodName: "GetMaxUsage",
			Handler:    _UsageService_GetMaxUsage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/usage_service.proto",
}

func (m *UsageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UsageRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UsageRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.To != nil {
		{
			size, err := m.To.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUsageService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.From != nil {
		{
			size, err := m.From.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUsageService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CurrentUsageResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CurrentUsageResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CurrentUsageResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.NumCpuUnits != 0 {
		i = encodeVarintUsageService(dAtA, i, uint64(m.NumCpuUnits))
		i--
		dAtA[i] = 0x18
	}
	if m.NumNodes != 0 {
		i = encodeVarintUsageService(dAtA, i, uint64(m.NumNodes))
		i--
		dAtA[i] = 0x10
	}
	if m.Timestamp != nil {
		{
			size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUsageService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MaxUsageResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MaxUsageResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MaxUsageResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MaxCpuUnits != 0 {
		i = encodeVarintUsageService(dAtA, i, uint64(m.MaxCpuUnits))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxCpuUnitsAt != nil {
		{
			size, err := m.MaxCpuUnitsAt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUsageService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.MaxNodes != 0 {
		i = encodeVarintUsageService(dAtA, i, uint64(m.MaxNodes))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxNodesAt != nil {
		{
			size, err := m.MaxNodesAt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUsageService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUsageService(dAtA []byte, offset int, v uint64) int {
	offset -= sovUsageService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UsageRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovUsageService(uint64(l))
	}
	if m.To != nil {
		l = m.To.Size()
		n += 1 + l + sovUsageService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CurrentUsageResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovUsageService(uint64(l))
	}
	if m.NumNodes != 0 {
		n += 1 + sovUsageService(uint64(m.NumNodes))
	}
	if m.NumCpuUnits != 0 {
		n += 1 + sovUsageService(uint64(m.NumCpuUnits))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MaxUsageResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxNodesAt != nil {
		l = m.MaxNodesAt.Size()
		n += 1 + l + sovUsageService(uint64(l))
	}
	if m.MaxNodes != 0 {
		n += 1 + sovUsageService(uint64(m.MaxNodes))
	}
	if m.MaxCpuUnitsAt != nil {
		l = m.MaxCpuUnitsAt.Size()
		n += 1 + l + sovUsageService(uint64(l))
	}
	if m.MaxCpuUnits != 0 {
		n += 1 + sovUsageService(uint64(m.MaxCpuUnits))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovUsageService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUsageService(x uint64) (n int) {
	return sovUsageService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UsageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUsageService
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
			return fmt.Errorf("proto: UsageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UsageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsageService
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
				return ErrInvalidLengthUsageService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUsageService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &types.Timestamp{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsageService
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
				return ErrInvalidLengthUsageService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUsageService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.To == nil {
				m.To = &types.Timestamp{}
			}
			if err := m.To.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUsageService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUsageService
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
func (m *CurrentUsageResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUsageService
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
			return fmt.Errorf("proto: CurrentUsageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CurrentUsageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsageService
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
				return ErrInvalidLengthUsageService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUsageService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &types.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumNodes", wireType)
			}
			m.NumNodes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsageService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumNodes |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumCpuUnits", wireType)
			}
			m.NumCpuUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsageService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumCpuUnits |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUsageService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUsageService
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
func (m *MaxUsageResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUsageService
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
			return fmt.Errorf("proto: MaxUsageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MaxUsageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNodesAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsageService
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
				return ErrInvalidLengthUsageService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUsageService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxNodesAt == nil {
				m.MaxNodesAt = &types.Timestamp{}
			}
			if err := m.MaxNodesAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNodes", wireType)
			}
			m.MaxNodes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsageService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNodes |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxCpuUnitsAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsageService
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
				return ErrInvalidLengthUsageService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUsageService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxCpuUnitsAt == nil {
				m.MaxCpuUnitsAt = &types.Timestamp{}
			}
			if err := m.MaxCpuUnitsAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxCpuUnits", wireType)
			}
			m.MaxCpuUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsageService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxCpuUnits |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUsageService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUsageService
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
func skipUsageService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUsageService
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
					return 0, ErrIntOverflowUsageService
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
					return 0, ErrIntOverflowUsageService
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
				return 0, ErrInvalidLengthUsageService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUsageService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUsageService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUsageService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUsageService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUsageService = fmt.Errorf("proto: unexpected end of group")
)
