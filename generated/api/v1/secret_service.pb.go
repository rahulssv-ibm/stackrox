// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/secret_service.proto

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

// A list of secrets (free of scoped information)
// Next Tag: 2
type SecretList struct {
	Secrets              []*storage.Secret `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SecretList) Reset()         { *m = SecretList{} }
func (m *SecretList) String() string { return proto.CompactTextString(m) }
func (*SecretList) ProtoMessage()    {}
func (*SecretList) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a81c8939a98fd13, []int{0}
}
func (m *SecretList) XXX_UnmarshalVT(b []byte) error {
	return m.UnmarshalVT(b)
}
func (m *SecretList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SecretList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SecretList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretList.Merge(m, src)
}
func (m *SecretList) XXX_Size() int {
	return m.Size()
}
func (m *SecretList) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretList.DiscardUnknown(m)
}

var xxx_messageInfo_SecretList proto.InternalMessageInfo

func (m *SecretList) GetSecrets() []*storage.Secret {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *SecretList) MessageClone() proto.Message {
	return m.Clone()
}
func (m *SecretList) Clone() *SecretList {
	if m == nil {
		return nil
	}
	cloned := new(SecretList)
	*cloned = *m

	if m.Secrets != nil {
		cloned.Secrets = make([]*storage.Secret, len(m.Secrets))
		for idx, v := range m.Secrets {
			cloned.Secrets[idx] = v.Clone()
		}
	}
	return cloned
}

// A list of secrets with their relationships.
// Next Tag: 2
type ListSecretsResponse struct {
	Secrets              []*storage.ListSecret `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListSecretsResponse) Reset()         { *m = ListSecretsResponse{} }
func (m *ListSecretsResponse) String() string { return proto.CompactTextString(m) }
func (*ListSecretsResponse) ProtoMessage()    {}
func (*ListSecretsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a81c8939a98fd13, []int{1}
}
func (m *ListSecretsResponse) XXX_UnmarshalVT(b []byte) error {
	return m.UnmarshalVT(b)
}
func (m *ListSecretsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListSecretsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListSecretsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSecretsResponse.Merge(m, src)
}
func (m *ListSecretsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListSecretsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSecretsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSecretsResponse proto.InternalMessageInfo

func (m *ListSecretsResponse) GetSecrets() []*storage.ListSecret {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *ListSecretsResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ListSecretsResponse) Clone() *ListSecretsResponse {
	if m == nil {
		return nil
	}
	cloned := new(ListSecretsResponse)
	*cloned = *m

	if m.Secrets != nil {
		cloned.Secrets = make([]*storage.ListSecret, len(m.Secrets))
		for idx, v := range m.Secrets {
			cloned.Secrets[idx] = v.Clone()
		}
	}
	return cloned
}

type CountSecretsResponse struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountSecretsResponse) Reset()         { *m = CountSecretsResponse{} }
func (m *CountSecretsResponse) String() string { return proto.CompactTextString(m) }
func (*CountSecretsResponse) ProtoMessage()    {}
func (*CountSecretsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a81c8939a98fd13, []int{2}
}
func (m *CountSecretsResponse) XXX_UnmarshalVT(b []byte) error {
	return m.UnmarshalVT(b)
}
func (m *CountSecretsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CountSecretsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CountSecretsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountSecretsResponse.Merge(m, src)
}
func (m *CountSecretsResponse) XXX_Size() int {
	return m.Size()
}
func (m *CountSecretsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CountSecretsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CountSecretsResponse proto.InternalMessageInfo

func (m *CountSecretsResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *CountSecretsResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *CountSecretsResponse) Clone() *CountSecretsResponse {
	if m == nil {
		return nil
	}
	cloned := new(CountSecretsResponse)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterType((*SecretList)(nil), "v1.SecretList")
	proto.RegisterType((*ListSecretsResponse)(nil), "v1.ListSecretsResponse")
	proto.RegisterType((*CountSecretsResponse)(nil), "v1.CountSecretsResponse")
}

func init() { proto.RegisterFile("api/v1/secret_service.proto", fileDescriptor_7a81c8939a98fd13) }

var fileDescriptor_7a81c8939a98fd13 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0xff, 0xe4, 0xa7, 0x8a, 0xd3, 0x16, 0xcb, 0xa4, 0x60, 0x88, 0x12, 0x4a, 0x56, 0x15,
	0x74, 0x42, 0xea, 0xc2, 0x7d, 0x2d, 0x68, 0xc1, 0x85, 0xa6, 0x1b, 0x71, 0x23, 0x63, 0x1c, 0xea,
	0xa0, 0xcd, 0x0d, 0x33, 0xd3, 0x68, 0x11, 0x37, 0xbe, 0x82, 0x1b, 0x1f, 0xc9, 0xa5, 0xe0, 0x0b,
	0x48, 0xf5, 0x19, 0x5c, 0x4b, 0x32, 0xd1, 0xd4, 0xda, 0xe5, 0xbd, 0xe7, 0x9c, 0x2f, 0x27, 0x77,
	0xd0, 0x3a, 0x4d, 0xb8, 0x9f, 0x06, 0xbe, 0x64, 0x91, 0x60, 0xea, 0x4c, 0x32, 0x91, 0xf2, 0x88,
	0x91, 0x44, 0x80, 0x02, 0x6c, 0xa6, 0x81, 0x63, 0x15, 0x86, 0x08, 0x46, 0x23, 0x88, 0xb5, 0xe0,
	0x94, 0x29, 0x2a, 0xa2, 0xcb, 0xdf, 0x29, 0x67, 0x63, 0x08, 0x30, 0xbc, 0x66, 0x7e, 0xe6, 0xa1,
	0x71, 0x0c, 0x8a, 0x2a, 0x0e, 0xb1, 0x2c, 0xd4, 0xa6, 0x54, 0x20, 0xe8, 0x90, 0x15, 0x5f, 0xd4,
	0x5b, 0x6f, 0x17, 0xa1, 0x41, 0x3e, 0x1f, 0x72, 0xa9, 0xf0, 0x26, 0x5a, 0xd6, 0xaa, 0xb4, 0x8d,
	0xd6, 0xff, 0x76, 0xb5, 0xb3, 0x4a, 0x8a, 0x14, 0xd1, 0xae, 0xf0, 0x5b, 0xf7, 0x7a, 0xc8, 0xca,
	0x22, 0x7a, 0x2d, 0x43, 0x26, 0x13, 0x88, 0x25, 0xc3, 0xdb, 0xf3, 0x04, 0xeb, 0x87, 0x50, 0xda,
	0x4b, 0xca, 0x16, 0x6a, 0xee, 0xc1, 0x38, 0xfe, 0x83, 0x69, 0xa2, 0x4a, 0x94, 0xed, 0x6d, 0xa3,
	0x65, 0xb4, 0x2b, 0xa1, 0x1e, 0x3a, 0x9f, 0x06, 0xaa, 0x6b, 0xe7, 0x40, 0xff, 0x38, 0x3e, 0x40,
	0x2b, 0xfb, 0xac, 0x48, 0xe3, 0x06, 0x49, 0x03, 0x12, 0x32, 0x09, 0x63, 0x11, 0xb1, 0xee, 0xa4,
	0xdf, 0x73, 0xe6, 0xeb, 0x7b, 0xf6, 0xc3, 0xeb, 0xc7, 0xa3, 0x89, 0x71, 0xa3, 0x3c, 0xbd, 0xf4,
	0xef, 0xf8, 0xc5, 0x3d, 0x3e, 0x42, 0xb5, 0xd9, 0x26, 0xb8, 0x96, 0xc3, 0xe8, 0xcd, 0xf1, 0x98,
	0x89, 0x89, 0x63, 0x67, 0xd3, 0xa2, 0xa6, 0x0b, 0x89, 0x79, 0x5b, 0xdc, 0x47, 0xd5, 0x99, 0x0b,
	0xcd, 0x01, 0xd7, 0xb2, 0x69, 0xc1, 0x01, 0x3d, 0x2b, 0xe7, 0xd5, 0x71, 0x75, 0x86, 0xd7, 0x25,
	0xcf, 0x53, 0xd7, 0x78, 0x99, 0xba, 0xc6, 0xdb, 0xd4, 0x35, 0x9e, 0xde, 0xdd, 0x7f, 0xc8, 0xe6,
	0x40, 0xa4, 0xa2, 0xd1, 0x95, 0x80, 0x5b, 0xfd, 0x94, 0x84, 0x26, 0x9c, 0xa4, 0xc1, 0xa9, 0x99,
	0x06, 0x27, 0xe6, 0xf9, 0x52, 0xbe, 0xdb, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x86, 0xe3, 0x37,
	0x2d, 0x67, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SecretServiceClient is the client API for SecretService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
type SecretServiceClient interface {
	// GetSecret returns a secret given its ID.
	GetSecret(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*storage.Secret, error)
	// CountSecrets returns the number of secrets.
	CountSecrets(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*CountSecretsResponse, error)
	// ListSecrets returns the list of secrets.
	ListSecrets(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*ListSecretsResponse, error)
}

type secretServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretServiceClient(cc grpc.ClientConnInterface) SecretServiceClient {
	return &secretServiceClient{cc}
}

func (c *secretServiceClient) GetSecret(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*storage.Secret, error) {
	out := new(storage.Secret)
	err := c.cc.Invoke(ctx, "/v1.SecretService/GetSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) CountSecrets(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*CountSecretsResponse, error) {
	out := new(CountSecretsResponse)
	err := c.cc.Invoke(ctx, "/v1.SecretService/CountSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) ListSecrets(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*ListSecretsResponse, error) {
	out := new(ListSecretsResponse)
	err := c.cc.Invoke(ctx, "/v1.SecretService/ListSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretServiceServer is the server API for SecretService service.
type SecretServiceServer interface {
	// GetSecret returns a secret given its ID.
	GetSecret(context.Context, *ResourceByID) (*storage.Secret, error)
	// CountSecrets returns the number of secrets.
	CountSecrets(context.Context, *RawQuery) (*CountSecretsResponse, error)
	// ListSecrets returns the list of secrets.
	ListSecrets(context.Context, *RawQuery) (*ListSecretsResponse, error)
}

// UnimplementedSecretServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSecretServiceServer struct {
}

func (*UnimplementedSecretServiceServer) GetSecret(ctx context.Context, req *ResourceByID) (*storage.Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecret not implemented")
}
func (*UnimplementedSecretServiceServer) CountSecrets(ctx context.Context, req *RawQuery) (*CountSecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountSecrets not implemented")
}
func (*UnimplementedSecretServiceServer) ListSecrets(ctx context.Context, req *RawQuery) (*ListSecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSecrets not implemented")
}

func RegisterSecretServiceServer(s *grpc.Server, srv SecretServiceServer) {
	s.RegisterService(&_SecretService_serviceDesc, srv)
}

func _SecretService_GetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).GetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SecretService/GetSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).GetSecret(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_CountSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).CountSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SecretService/CountSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).CountSecrets(ctx, req.(*RawQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_ListSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).ListSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SecretService/ListSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).ListSecrets(ctx, req.(*RawQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecretService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.SecretService",
	HandlerType: (*SecretServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSecret",
			Handler:    _SecretService_GetSecret_Handler,
		},
		{
			MethodName: "CountSecrets",
			Handler:    _SecretService_CountSecrets_Handler,
		},
		{
			MethodName: "ListSecrets",
			Handler:    _SecretService_ListSecrets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/secret_service.proto",
}

func (m *SecretList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SecretList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SecretList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Secrets) > 0 {
		for iNdEx := len(m.Secrets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Secrets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSecretService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ListSecretsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListSecretsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListSecretsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Secrets) > 0 {
		for iNdEx := len(m.Secrets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Secrets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSecretService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CountSecretsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CountSecretsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CountSecretsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Count != 0 {
		i = encodeVarintSecretService(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSecretService(dAtA []byte, offset int, v uint64) int {
	offset -= sovSecretService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SecretList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Secrets) > 0 {
		for _, e := range m.Secrets {
			l = e.Size()
			n += 1 + l + sovSecretService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ListSecretsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Secrets) > 0 {
		for _, e := range m.Secrets {
			l = e.Size()
			n += 1 + l + sovSecretService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CountSecretsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Count != 0 {
		n += 1 + sovSecretService(uint64(m.Count))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSecretService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSecretService(x uint64) (n int) {
	return sovSecretService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SecretList) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSecretService
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
			return fmt.Errorf("proto: SecretList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SecretList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secrets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecretService
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
				return ErrInvalidLengthSecretService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSecretService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secrets = append(m.Secrets, &storage.Secret{})
			if err := m.Secrets[len(m.Secrets)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSecretService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSecretService
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
func (m *ListSecretsResponse) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSecretService
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
			return fmt.Errorf("proto: ListSecretsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListSecretsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secrets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecretService
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
				return ErrInvalidLengthSecretService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSecretService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secrets = append(m.Secrets, &storage.ListSecret{})
			if err := m.Secrets[len(m.Secrets)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSecretService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSecretService
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
func (m *CountSecretsResponse) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSecretService
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
			return fmt.Errorf("proto: CountSecretsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CountSecretsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecretService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSecretService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSecretService
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
func skipSecretService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSecretService
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
					return 0, ErrIntOverflowSecretService
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
					return 0, ErrIntOverflowSecretService
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
				return 0, ErrInvalidLengthSecretService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSecretService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSecretService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSecretService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSecretService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSecretService = fmt.Errorf("proto: unexpected end of group")
)
