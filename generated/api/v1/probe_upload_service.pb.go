// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/probe_upload_service.proto

package v1

import (
	context "context"
	fmt "fmt"
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

type ProbeUploadManifest struct {
	Files                []*ProbeUploadManifest_File `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ProbeUploadManifest) Reset()         { *m = ProbeUploadManifest{} }
func (m *ProbeUploadManifest) String() string { return proto.CompactTextString(m) }
func (*ProbeUploadManifest) ProtoMessage()    {}
func (*ProbeUploadManifest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b730ffd5769ce00f, []int{0}
}
func (m *ProbeUploadManifest) XXX_UnmarshalVT(b []byte) error {
	return m.UnmarshalVT(b)
}
func (m *ProbeUploadManifest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProbeUploadManifest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProbeUploadManifest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeUploadManifest.Merge(m, src)
}
func (m *ProbeUploadManifest) XXX_Size() int {
	return m.Size()
}
func (m *ProbeUploadManifest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeUploadManifest.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeUploadManifest proto.InternalMessageInfo

func (m *ProbeUploadManifest) GetFiles() []*ProbeUploadManifest_File {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *ProbeUploadManifest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ProbeUploadManifest) Clone() *ProbeUploadManifest {
	if m == nil {
		return nil
	}
	cloned := new(ProbeUploadManifest)
	*cloned = *m

	if m.Files != nil {
		cloned.Files = make([]*ProbeUploadManifest_File, len(m.Files))
		for idx, v := range m.Files {
			cloned.Files[idx] = v.Clone()
		}
	}
	return cloned
}

type ProbeUploadManifest_File struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size_                int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Crc32                uint32   `protobuf:"varint,3,opt,name=crc32,proto3" json:"crc32,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbeUploadManifest_File) Reset()         { *m = ProbeUploadManifest_File{} }
func (m *ProbeUploadManifest_File) String() string { return proto.CompactTextString(m) }
func (*ProbeUploadManifest_File) ProtoMessage()    {}
func (*ProbeUploadManifest_File) Descriptor() ([]byte, []int) {
	return fileDescriptor_b730ffd5769ce00f, []int{0, 0}
}
func (m *ProbeUploadManifest_File) XXX_UnmarshalVT(b []byte) error {
	return m.UnmarshalVT(b)
}
func (m *ProbeUploadManifest_File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProbeUploadManifest_File.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProbeUploadManifest_File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeUploadManifest_File.Merge(m, src)
}
func (m *ProbeUploadManifest_File) XXX_Size() int {
	return m.Size()
}
func (m *ProbeUploadManifest_File) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeUploadManifest_File.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeUploadManifest_File proto.InternalMessageInfo

func (m *ProbeUploadManifest_File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProbeUploadManifest_File) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *ProbeUploadManifest_File) GetCrc32() uint32 {
	if m != nil {
		return m.Crc32
	}
	return 0
}

func (m *ProbeUploadManifest_File) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ProbeUploadManifest_File) Clone() *ProbeUploadManifest_File {
	if m == nil {
		return nil
	}
	cloned := new(ProbeUploadManifest_File)
	*cloned = *m

	return cloned
}

type GetExistingProbesRequest struct {
	FilesToCheck         []string `protobuf:"bytes,1,rep,name=files_to_check,json=filesToCheck,proto3" json:"files_to_check,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetExistingProbesRequest) Reset()         { *m = GetExistingProbesRequest{} }
func (m *GetExistingProbesRequest) String() string { return proto.CompactTextString(m) }
func (*GetExistingProbesRequest) ProtoMessage()    {}
func (*GetExistingProbesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b730ffd5769ce00f, []int{1}
}
func (m *GetExistingProbesRequest) XXX_UnmarshalVT(b []byte) error {
	return m.UnmarshalVT(b)
}
func (m *GetExistingProbesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetExistingProbesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetExistingProbesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExistingProbesRequest.Merge(m, src)
}
func (m *GetExistingProbesRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetExistingProbesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExistingProbesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetExistingProbesRequest proto.InternalMessageInfo

func (m *GetExistingProbesRequest) GetFilesToCheck() []string {
	if m != nil {
		return m.FilesToCheck
	}
	return nil
}

func (m *GetExistingProbesRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GetExistingProbesRequest) Clone() *GetExistingProbesRequest {
	if m == nil {
		return nil
	}
	cloned := new(GetExistingProbesRequest)
	*cloned = *m

	if m.FilesToCheck != nil {
		cloned.FilesToCheck = make([]string, len(m.FilesToCheck))
		copy(cloned.FilesToCheck, m.FilesToCheck)
	}
	return cloned
}

type GetExistingProbesResponse struct {
	ExistingFiles        []*ProbeUploadManifest_File `protobuf:"bytes,1,rep,name=existing_files,json=existingFiles,proto3" json:"existing_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *GetExistingProbesResponse) Reset()         { *m = GetExistingProbesResponse{} }
func (m *GetExistingProbesResponse) String() string { return proto.CompactTextString(m) }
func (*GetExistingProbesResponse) ProtoMessage()    {}
func (*GetExistingProbesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b730ffd5769ce00f, []int{2}
}
func (m *GetExistingProbesResponse) XXX_UnmarshalVT(b []byte) error {
	return m.UnmarshalVT(b)
}
func (m *GetExistingProbesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetExistingProbesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetExistingProbesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExistingProbesResponse.Merge(m, src)
}
func (m *GetExistingProbesResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetExistingProbesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExistingProbesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetExistingProbesResponse proto.InternalMessageInfo

func (m *GetExistingProbesResponse) GetExistingFiles() []*ProbeUploadManifest_File {
	if m != nil {
		return m.ExistingFiles
	}
	return nil
}

func (m *GetExistingProbesResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GetExistingProbesResponse) Clone() *GetExistingProbesResponse {
	if m == nil {
		return nil
	}
	cloned := new(GetExistingProbesResponse)
	*cloned = *m

	if m.ExistingFiles != nil {
		cloned.ExistingFiles = make([]*ProbeUploadManifest_File, len(m.ExistingFiles))
		for idx, v := range m.ExistingFiles {
			cloned.ExistingFiles[idx] = v.Clone()
		}
	}
	return cloned
}

func init() {
	proto.RegisterType((*ProbeUploadManifest)(nil), "v1.ProbeUploadManifest")
	proto.RegisterType((*ProbeUploadManifest_File)(nil), "v1.ProbeUploadManifest.File")
	proto.RegisterType((*GetExistingProbesRequest)(nil), "v1.GetExistingProbesRequest")
	proto.RegisterType((*GetExistingProbesResponse)(nil), "v1.GetExistingProbesResponse")
}

func init() { proto.RegisterFile("api/v1/probe_upload_service.proto", fileDescriptor_b730ffd5769ce00f) }

var fileDescriptor_b730ffd5769ce00f = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0xed, 0xa4, 0xed, 0x07, 0x9d, 0xcf, 0x16, 0x1c, 0x5d, 0xc4, 0xda, 0x86, 0x18, 0x5d, 0x64,
	0x35, 0x21, 0xe9, 0x0b, 0x88, 0xf5, 0x67, 0x25, 0x48, 0x54, 0x10, 0x37, 0x71, 0x1a, 0xa7, 0x71,
	0x68, 0xcc, 0xc4, 0xcc, 0x34, 0x14, 0x97, 0x6e, 0x5d, 0xb8, 0x70, 0xe3, 0x23, 0xb9, 0x14, 0x7c,
	0x01, 0xa9, 0x3e, 0x88, 0x64, 0x62, 0x45, 0xb0, 0x05, 0x77, 0x73, 0xcf, 0x3d, 0xf7, 0x9e, 0x73,
	0xe7, 0xc0, 0x0d, 0x92, 0x32, 0x27, 0x77, 0x9d, 0x34, 0xe3, 0x03, 0x1a, 0x8c, 0xd3, 0x98, 0x93,
	0xcb, 0x40, 0xd0, 0x2c, 0x67, 0x21, 0xc5, 0x69, 0xc6, 0x25, 0x47, 0x5a, 0xee, 0xb6, 0x3b, 0x11,
	0xe7, 0x51, 0x4c, 0x9d, 0x82, 0x4d, 0x92, 0x84, 0x4b, 0x22, 0x19, 0x4f, 0x44, 0xc9, 0xb0, 0x1e,
	0x00, 0x5c, 0x39, 0x2a, 0x16, 0x9c, 0xaa, 0xf9, 0x43, 0x92, 0xb0, 0x21, 0x15, 0x12, 0x79, 0xb0,
	0x3e, 0x64, 0x31, 0x15, 0x3a, 0x30, 0xab, 0xf6, 0x7f, 0xaf, 0x83, 0x73, 0x17, 0xcf, 0xe1, 0xe1,
	0x7d, 0x16, 0x53, 0xbf, 0xa4, 0xb6, 0x77, 0x61, 0xad, 0x28, 0x11, 0x82, 0xb5, 0x84, 0x5c, 0x53,
	0x1d, 0x98, 0xc0, 0x6e, 0xf8, 0xea, 0x5d, 0x60, 0x82, 0xdd, 0x52, 0x5d, 0x33, 0x81, 0x5d, 0xf5,
	0xd5, 0x1b, 0xad, 0xc2, 0x7a, 0x98, 0x85, 0x3d, 0x4f, 0xaf, 0x9a, 0xc0, 0x6e, 0xfa, 0x65, 0x61,
	0x6d, 0x43, 0xfd, 0x80, 0xca, 0xbd, 0x09, 0x13, 0x92, 0x25, 0x91, 0xd2, 0x14, 0x3e, 0xbd, 0x19,
	0x17, 0xae, 0xb6, 0x60, 0x4b, 0x49, 0x05, 0x92, 0x07, 0xe1, 0x15, 0x0d, 0x47, 0xca, 0x5e, 0xc3,
	0x5f, 0x52, 0xe8, 0x09, 0xef, 0x17, 0x98, 0x75, 0x01, 0xd7, 0xe6, 0x6c, 0x10, 0x29, 0x4f, 0x04,
	0x45, 0x7d, 0xd8, 0xa2, 0x5f, 0x9d, 0xe0, 0xef, 0x17, 0x36, 0x67, 0x33, 0x45, 0x25, 0xbc, 0x7b,
	0x00, 0xd1, 0x0f, 0xee, 0x71, 0xf9, 0xe9, 0x68, 0x0c, 0x97, 0x7f, 0x09, 0x23, 0xb5, 0x78, 0xd1,
	0x45, 0xed, 0xee, 0x82, 0x6e, 0xe9, 0xd6, 0xda, 0xbc, 0x7b, 0xfd, 0x78, 0xd4, 0xba, 0xd6, 0xfa,
	0x77, 0xd0, 0x65, 0xce, 0x4e, 0x44, 0xe5, 0xcc, 0xd2, 0x0e, 0x7e, 0x9e, 0x1a, 0xe0, 0x65, 0x6a,
	0x80, 0xb7, 0xa9, 0x01, 0x9e, 0xde, 0x8d, 0x0a, 0xd4, 0x19, 0xc7, 0x42, 0x92, 0x70, 0x94, 0xf1,
	0x49, 0x19, 0x34, 0x26, 0x29, 0xc3, 0xb9, 0x7b, 0xae, 0xe5, 0xee, 0x59, 0x65, 0xf0, 0x4f, 0x61,
	0xbd, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x13, 0xe7, 0xe7, 0x43, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProbeUploadServiceClient is the client API for ProbeUploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
type ProbeUploadServiceClient interface {
	GetExistingProbes(ctx context.Context, in *GetExistingProbesRequest, opts ...grpc.CallOption) (*GetExistingProbesResponse, error)
}

type probeUploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProbeUploadServiceClient(cc grpc.ClientConnInterface) ProbeUploadServiceClient {
	return &probeUploadServiceClient{cc}
}

func (c *probeUploadServiceClient) GetExistingProbes(ctx context.Context, in *GetExistingProbesRequest, opts ...grpc.CallOption) (*GetExistingProbesResponse, error) {
	out := new(GetExistingProbesResponse)
	err := c.cc.Invoke(ctx, "/v1.ProbeUploadService/GetExistingProbes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProbeUploadServiceServer is the server API for ProbeUploadService service.
type ProbeUploadServiceServer interface {
	GetExistingProbes(context.Context, *GetExistingProbesRequest) (*GetExistingProbesResponse, error)
}

// UnimplementedProbeUploadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProbeUploadServiceServer struct {
}

func (*UnimplementedProbeUploadServiceServer) GetExistingProbes(ctx context.Context, req *GetExistingProbesRequest) (*GetExistingProbesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExistingProbes not implemented")
}

func RegisterProbeUploadServiceServer(s *grpc.Server, srv ProbeUploadServiceServer) {
	s.RegisterService(&_ProbeUploadService_serviceDesc, srv)
}

func _ProbeUploadService_GetExistingProbes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExistingProbesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProbeUploadServiceServer).GetExistingProbes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProbeUploadService/GetExistingProbes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProbeUploadServiceServer).GetExistingProbes(ctx, req.(*GetExistingProbesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProbeUploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ProbeUploadService",
	HandlerType: (*ProbeUploadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExistingProbes",
			Handler:    _ProbeUploadService_GetExistingProbes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/probe_upload_service.proto",
}

func (m *ProbeUploadManifest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProbeUploadManifest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProbeUploadManifest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Files) > 0 {
		for iNdEx := len(m.Files) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Files[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProbeUploadService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ProbeUploadManifest_File) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProbeUploadManifest_File) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProbeUploadManifest_File) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Crc32 != 0 {
		i = encodeVarintProbeUploadService(dAtA, i, uint64(m.Crc32))
		i--
		dAtA[i] = 0x18
	}
	if m.Size_ != 0 {
		i = encodeVarintProbeUploadService(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProbeUploadService(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetExistingProbesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetExistingProbesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetExistingProbesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.FilesToCheck) > 0 {
		for iNdEx := len(m.FilesToCheck) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.FilesToCheck[iNdEx])
			copy(dAtA[i:], m.FilesToCheck[iNdEx])
			i = encodeVarintProbeUploadService(dAtA, i, uint64(len(m.FilesToCheck[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetExistingProbesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetExistingProbesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetExistingProbesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ExistingFiles) > 0 {
		for iNdEx := len(m.ExistingFiles) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExistingFiles[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProbeUploadService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintProbeUploadService(dAtA []byte, offset int, v uint64) int {
	offset -= sovProbeUploadService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProbeUploadManifest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Files) > 0 {
		for _, e := range m.Files {
			l = e.Size()
			n += 1 + l + sovProbeUploadService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ProbeUploadManifest_File) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProbeUploadService(uint64(l))
	}
	if m.Size_ != 0 {
		n += 1 + sovProbeUploadService(uint64(m.Size_))
	}
	if m.Crc32 != 0 {
		n += 1 + sovProbeUploadService(uint64(m.Crc32))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetExistingProbesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FilesToCheck) > 0 {
		for _, s := range m.FilesToCheck {
			l = len(s)
			n += 1 + l + sovProbeUploadService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetExistingProbesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ExistingFiles) > 0 {
		for _, e := range m.ExistingFiles {
			l = e.Size()
			n += 1 + l + sovProbeUploadService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProbeUploadService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProbeUploadService(x uint64) (n int) {
	return sovProbeUploadService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProbeUploadManifest) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProbeUploadService
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
			return fmt.Errorf("proto: ProbeUploadManifest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProbeUploadManifest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Files", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProbeUploadService
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
				return ErrInvalidLengthProbeUploadService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProbeUploadService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Files = append(m.Files, &ProbeUploadManifest_File{})
			if err := m.Files[len(m.Files)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProbeUploadService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProbeUploadService
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
func (m *ProbeUploadManifest_File) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProbeUploadService
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
			return fmt.Errorf("proto: File: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: File: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProbeUploadService
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
				return ErrInvalidLengthProbeUploadService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProbeUploadService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProbeUploadService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Crc32", wireType)
			}
			m.Crc32 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProbeUploadService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Crc32 |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProbeUploadService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProbeUploadService
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
func (m *GetExistingProbesRequest) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProbeUploadService
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
			return fmt.Errorf("proto: GetExistingProbesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetExistingProbesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilesToCheck", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProbeUploadService
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
				return ErrInvalidLengthProbeUploadService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProbeUploadService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilesToCheck = append(m.FilesToCheck, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProbeUploadService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProbeUploadService
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
func (m *GetExistingProbesResponse) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProbeUploadService
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
			return fmt.Errorf("proto: GetExistingProbesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetExistingProbesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExistingFiles", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProbeUploadService
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
				return ErrInvalidLengthProbeUploadService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProbeUploadService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExistingFiles = append(m.ExistingFiles, &ProbeUploadManifest_File{})
			if err := m.ExistingFiles[len(m.ExistingFiles)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProbeUploadService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProbeUploadService
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
func skipProbeUploadService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProbeUploadService
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
					return 0, ErrIntOverflowProbeUploadService
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
					return 0, ErrIntOverflowProbeUploadService
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
				return 0, ErrInvalidLengthProbeUploadService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProbeUploadService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProbeUploadService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProbeUploadService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProbeUploadService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProbeUploadService = fmt.Errorf("proto: unexpected end of group")
)
