// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/integration_health_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	types "github.com/gogo/protobuf/types"
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

type VulnDefinitionsInfoRequest_Component int32

const (
	VulnDefinitionsInfoRequest_SCANNER    VulnDefinitionsInfoRequest_Component = 0
	VulnDefinitionsInfoRequest_SCANNER_V4 VulnDefinitionsInfoRequest_Component = 1
)

var VulnDefinitionsInfoRequest_Component_name = map[int32]string{
	0: "SCANNER",
	1: "SCANNER_V4",
}

var VulnDefinitionsInfoRequest_Component_value = map[string]int32{
	"SCANNER":    0,
	"SCANNER_V4": 1,
}

func (x VulnDefinitionsInfoRequest_Component) String() string {
	return proto.EnumName(VulnDefinitionsInfoRequest_Component_name, int32(x))
}

func (VulnDefinitionsInfoRequest_Component) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4c40da462e693fb8, []int{1, 0}
}

type GetIntegrationHealthResponse struct {
	IntegrationHealth    []*storage.IntegrationHealth `protobuf:"bytes,1,rep,name=integrationHealth,proto3" json:"integrationHealth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetIntegrationHealthResponse) Reset()         { *m = GetIntegrationHealthResponse{} }
func (m *GetIntegrationHealthResponse) String() string { return proto.CompactTextString(m) }
func (*GetIntegrationHealthResponse) ProtoMessage()    {}
func (*GetIntegrationHealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c40da462e693fb8, []int{0}
}
func (m *GetIntegrationHealthResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetIntegrationHealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetIntegrationHealthResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetIntegrationHealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIntegrationHealthResponse.Merge(m, src)
}
func (m *GetIntegrationHealthResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetIntegrationHealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIntegrationHealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetIntegrationHealthResponse proto.InternalMessageInfo

func (m *GetIntegrationHealthResponse) GetIntegrationHealth() []*storage.IntegrationHealth {
	if m != nil {
		return m.IntegrationHealth
	}
	return nil
}

func (m *GetIntegrationHealthResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GetIntegrationHealthResponse) Clone() *GetIntegrationHealthResponse {
	if m == nil {
		return nil
	}
	cloned := new(GetIntegrationHealthResponse)
	*cloned = *m

	if m.IntegrationHealth != nil {
		cloned.IntegrationHealth = make([]*storage.IntegrationHealth, len(m.IntegrationHealth))
		for idx, v := range m.IntegrationHealth {
			cloned.IntegrationHealth[idx] = v.Clone()
		}
	}
	return cloned
}

type VulnDefinitionsInfoRequest struct {
	Component            VulnDefinitionsInfoRequest_Component `protobuf:"varint,1,opt,name=component,proto3,enum=v1.VulnDefinitionsInfoRequest_Component" json:"component,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *VulnDefinitionsInfoRequest) Reset()         { *m = VulnDefinitionsInfoRequest{} }
func (m *VulnDefinitionsInfoRequest) String() string { return proto.CompactTextString(m) }
func (*VulnDefinitionsInfoRequest) ProtoMessage()    {}
func (*VulnDefinitionsInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c40da462e693fb8, []int{1}
}
func (m *VulnDefinitionsInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VulnDefinitionsInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VulnDefinitionsInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VulnDefinitionsInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VulnDefinitionsInfoRequest.Merge(m, src)
}
func (m *VulnDefinitionsInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *VulnDefinitionsInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VulnDefinitionsInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VulnDefinitionsInfoRequest proto.InternalMessageInfo

func (m *VulnDefinitionsInfoRequest) GetComponent() VulnDefinitionsInfoRequest_Component {
	if m != nil {
		return m.Component
	}
	return VulnDefinitionsInfoRequest_SCANNER
}

func (m *VulnDefinitionsInfoRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *VulnDefinitionsInfoRequest) Clone() *VulnDefinitionsInfoRequest {
	if m == nil {
		return nil
	}
	cloned := new(VulnDefinitionsInfoRequest)
	*cloned = *m

	return cloned
}

type VulnDefinitionsInfo struct {
	LastUpdatedTimestamp *types.Timestamp `protobuf:"bytes,1,opt,name=last_updated_timestamp,json=lastUpdatedTimestamp,proto3" json:"last_updated_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *VulnDefinitionsInfo) Reset()         { *m = VulnDefinitionsInfo{} }
func (m *VulnDefinitionsInfo) String() string { return proto.CompactTextString(m) }
func (*VulnDefinitionsInfo) ProtoMessage()    {}
func (*VulnDefinitionsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c40da462e693fb8, []int{2}
}
func (m *VulnDefinitionsInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VulnDefinitionsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VulnDefinitionsInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VulnDefinitionsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VulnDefinitionsInfo.Merge(m, src)
}
func (m *VulnDefinitionsInfo) XXX_Size() int {
	return m.Size()
}
func (m *VulnDefinitionsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VulnDefinitionsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VulnDefinitionsInfo proto.InternalMessageInfo

func (m *VulnDefinitionsInfo) GetLastUpdatedTimestamp() *types.Timestamp {
	if m != nil {
		return m.LastUpdatedTimestamp
	}
	return nil
}

func (m *VulnDefinitionsInfo) MessageClone() proto.Message {
	return m.Clone()
}
func (m *VulnDefinitionsInfo) Clone() *VulnDefinitionsInfo {
	if m == nil {
		return nil
	}
	cloned := new(VulnDefinitionsInfo)
	*cloned = *m

	cloned.LastUpdatedTimestamp = m.LastUpdatedTimestamp.Clone()
	return cloned
}

func init() {
	proto.RegisterEnum("v1.VulnDefinitionsInfoRequest_Component", VulnDefinitionsInfoRequest_Component_name, VulnDefinitionsInfoRequest_Component_value)
	proto.RegisterType((*GetIntegrationHealthResponse)(nil), "v1.GetIntegrationHealthResponse")
	proto.RegisterType((*VulnDefinitionsInfoRequest)(nil), "v1.VulnDefinitionsInfoRequest")
	proto.RegisterType((*VulnDefinitionsInfo)(nil), "v1.VulnDefinitionsInfo")
}

func init() {
	proto.RegisterFile("api/v1/integration_health_service.proto", fileDescriptor_4c40da462e693fb8)
}

var fileDescriptor_4c40da462e693fb8 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcf, 0x6e, 0xd3, 0x4c,
	0x10, 0xaf, 0xf3, 0x7d, 0x02, 0x65, 0x8b, 0xaa, 0xb0, 0x94, 0x12, 0x59, 0x55, 0x1a, 0x22, 0xa1,
	0x04, 0x24, 0xd6, 0x24, 0xf0, 0x02, 0x34, 0x2d, 0x69, 0x2f, 0x51, 0x95, 0x42, 0x85, 0xb8, 0x44,
	0x1b, 0x67, 0xe2, 0xac, 0x6a, 0xef, 0x1a, 0xef, 0xd8, 0x2a, 0x37, 0xc4, 0x03, 0x70, 0xe1, 0xc2,
	0x23, 0x71, 0x44, 0xe2, 0x05, 0x20, 0xf0, 0x20, 0xc8, 0x6b, 0x27, 0xa9, 0x1a, 0xb7, 0x28, 0x37,
	0x7b, 0xe6, 0xf7, 0x67, 0x66, 0x7f, 0x43, 0x9a, 0x3c, 0x14, 0x4e, 0xd2, 0x76, 0x84, 0x44, 0xf0,
	0x22, 0x8e, 0x42, 0xc9, 0xe1, 0x14, 0xb8, 0x8f, 0xd3, 0xa1, 0x86, 0x28, 0x11, 0x2e, 0xb0, 0x30,
	0x52, 0xa8, 0x68, 0x29, 0x69, 0xdb, 0x7b, 0x9e, 0x52, 0x9e, 0x0f, 0x8e, 0xa9, 0x8c, 0xe2, 0x89,
	0x83, 0x22, 0x00, 0x8d, 0x3c, 0x08, 0x33, 0x90, 0xbd, 0x9b, 0x03, 0x52, 0x51, 0x2e, 0xa5, 0x42,
	0x23, 0xa8, 0xf3, 0x6e, 0x5d, 0xa3, 0x8a, 0xb8, 0x07, 0x05, 0x66, 0x39, 0x82, 0xe6, 0xd3, 0x40,
	0x10, 0xe2, 0x87, 0xac, 0xd6, 0x98, 0x92, 0xdd, 0x1e, 0xe0, 0xf1, 0x92, 0x72, 0x64, 0x18, 0x03,
	0xd0, 0xa1, 0x92, 0x1a, 0xe8, 0x11, 0xb9, 0x2b, 0xae, 0x36, 0xab, 0x56, 0xfd, 0xbf, 0xd6, 0x66,
	0xc7, 0x66, 0xb9, 0x23, 0x5b, 0xa5, 0xaf, 0x92, 0x1a, 0x9f, 0x2d, 0x62, 0x9f, 0xc5, 0xbe, 0x3c,
	0x80, 0x89, 0x90, 0xc2, 0x4c, 0x7e, 0x2c, 0x27, 0x6a, 0x00, 0xef, 0x63, 0xd0, 0x48, 0x5f, 0x91,
	0xb2, 0xab, 0x82, 0x50, 0x49, 0x90, 0x58, 0xb5, 0xea, 0x56, 0x6b, 0xab, 0xd3, 0x62, 0x49, 0x9b,
	0x5d, 0x4f, 0x61, 0xdd, 0x39, 0x7e, 0xb0, 0xa4, 0x36, 0x5a, 0xa4, 0xbc, 0xa8, 0xd3, 0x4d, 0x72,
	0xfb, 0xb4, 0xfb, 0xb2, 0xdf, 0x3f, 0x1c, 0x54, 0x36, 0xe8, 0x16, 0x21, 0xf9, 0xcf, 0xf0, 0xec,
	0x45, 0xc5, 0x6a, 0x78, 0xe4, 0x5e, 0x81, 0x38, 0x3d, 0x21, 0x3b, 0x3e, 0xd7, 0x38, 0x8c, 0xc3,
	0x31, 0x47, 0x18, 0x0f, 0x17, 0x29, 0x98, 0xa9, 0xd2, 0xb5, 0xb3, 0x18, 0xd8, 0x3c, 0x27, 0xf6,
	0x7a, 0x8e, 0x18, 0x6c, 0xa7, 0xcc, 0x37, 0x19, 0x71, 0x51, 0xed, 0xfc, 0xfa, 0x9f, 0x54, 0x57,
	0x9e, 0xe8, 0x34, 0xcb, 0x9f, 0x22, 0xd9, 0x4e, 0x03, 0x08, 0xb8, 0x07, 0x97, 0x30, 0x9a, 0x96,
	0xd3, 0xe5, 0x0f, 0xd3, 0xa4, 0xec, 0x7a, 0xfa, 0x79, 0x53, 0x4a, 0x0d, 0xe7, 0xd3, 0x8f, 0x3f,
	0x5f, 0x4a, 0x8f, 0x69, 0xf3, 0xca, 0xb1, 0x65, 0xf1, 0x3b, 0x22, 0x55, 0x17, 0x97, 0xd5, 0xc7,
	0xe4, 0x4e, 0x0f, 0xb0, 0xaf, 0x50, 0x4c, 0x04, 0x44, 0x6b, 0xba, 0x35, 0x8d, 0xdb, 0x43, 0xba,
	0x57, 0xec, 0x26, 0x17, 0xaa, 0x92, 0x54, 0x7a, 0x80, 0xfb, 0xdc, 0x3d, 0x8f, 0xc3, 0x13, 0x3f,
	0xf6, 0xc4, 0xba, 0x7b, 0x3d, 0x35, 0x4e, 0x4d, 0xfa, 0xa8, 0xd8, 0x09, 0x2e, 0x10, 0x22, 0xc9,
	0xfd, 0x91, 0x71, 0xd0, 0x34, 0x21, 0xf7, 0x7b, 0x80, 0x07, 0xe0, 0xfa, 0x3c, 0xc5, 0x25, 0xd0,
	0x55, 0x72, 0x22, 0xbc, 0x35, 0x4d, 0x9f, 0x19, 0xd3, 0x27, 0xb4, 0x55, 0x6c, 0x3a, 0x5e, 0xca,
	0xbb, 0xb9, 0xfc, 0x47, 0x8b, 0xec, 0xf4, 0x00, 0x8b, 0xae, 0xa9, 0x76, 0xf3, 0x0d, 0xdb, 0x0f,
	0xae, 0xe9, 0xff, 0x6b, 0xf5, 0x24, 0xf6, 0xe5, 0x78, 0x49, 0xd9, 0x67, 0xdf, 0x66, 0x35, 0xeb,
	0xfb, 0xac, 0x66, 0xfd, 0x9c, 0xd5, 0xac, 0xaf, 0xbf, 0x6b, 0x1b, 0xa4, 0x2a, 0x14, 0xd3, 0xc8,
	0xdd, 0xf3, 0x48, 0x5d, 0x64, 0xb7, 0xca, 0x78, 0x28, 0x58, 0xd2, 0x7e, 0x57, 0x4a, 0xda, 0x6f,
	0xad, 0xd1, 0x2d, 0x53, 0x7b, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x99, 0x6e, 0xdf, 0xb5, 0xa4,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IntegrationHealthServiceClient is the client API for IntegrationHealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
type IntegrationHealthServiceClient interface {
	GetImageIntegrations(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetIntegrationHealthResponse, error)
	GetNotifiers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetIntegrationHealthResponse, error)
	GetBackupPlugins(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetIntegrationHealthResponse, error)
	GetDeclarativeConfigs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetIntegrationHealthResponse, error)
	GetVulnDefinitionsInfo(ctx context.Context, in *VulnDefinitionsInfoRequest, opts ...grpc.CallOption) (*VulnDefinitionsInfo, error)
}

type integrationHealthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIntegrationHealthServiceClient(cc grpc.ClientConnInterface) IntegrationHealthServiceClient {
	return &integrationHealthServiceClient{cc}
}

func (c *integrationHealthServiceClient) GetImageIntegrations(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetIntegrationHealthResponse, error) {
	out := new(GetIntegrationHealthResponse)
	err := c.cc.Invoke(ctx, "/v1.IntegrationHealthService/GetImageIntegrations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationHealthServiceClient) GetNotifiers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetIntegrationHealthResponse, error) {
	out := new(GetIntegrationHealthResponse)
	err := c.cc.Invoke(ctx, "/v1.IntegrationHealthService/GetNotifiers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationHealthServiceClient) GetBackupPlugins(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetIntegrationHealthResponse, error) {
	out := new(GetIntegrationHealthResponse)
	err := c.cc.Invoke(ctx, "/v1.IntegrationHealthService/GetBackupPlugins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationHealthServiceClient) GetDeclarativeConfigs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetIntegrationHealthResponse, error) {
	out := new(GetIntegrationHealthResponse)
	err := c.cc.Invoke(ctx, "/v1.IntegrationHealthService/GetDeclarativeConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationHealthServiceClient) GetVulnDefinitionsInfo(ctx context.Context, in *VulnDefinitionsInfoRequest, opts ...grpc.CallOption) (*VulnDefinitionsInfo, error) {
	out := new(VulnDefinitionsInfo)
	err := c.cc.Invoke(ctx, "/v1.IntegrationHealthService/GetVulnDefinitionsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntegrationHealthServiceServer is the server API for IntegrationHealthService service.
type IntegrationHealthServiceServer interface {
	GetImageIntegrations(context.Context, *Empty) (*GetIntegrationHealthResponse, error)
	GetNotifiers(context.Context, *Empty) (*GetIntegrationHealthResponse, error)
	GetBackupPlugins(context.Context, *Empty) (*GetIntegrationHealthResponse, error)
	GetDeclarativeConfigs(context.Context, *Empty) (*GetIntegrationHealthResponse, error)
	GetVulnDefinitionsInfo(context.Context, *VulnDefinitionsInfoRequest) (*VulnDefinitionsInfo, error)
}

// UnimplementedIntegrationHealthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIntegrationHealthServiceServer struct {
}

func (*UnimplementedIntegrationHealthServiceServer) GetImageIntegrations(ctx context.Context, req *Empty) (*GetIntegrationHealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImageIntegrations not implemented")
}
func (*UnimplementedIntegrationHealthServiceServer) GetNotifiers(ctx context.Context, req *Empty) (*GetIntegrationHealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifiers not implemented")
}
func (*UnimplementedIntegrationHealthServiceServer) GetBackupPlugins(ctx context.Context, req *Empty) (*GetIntegrationHealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBackupPlugins not implemented")
}
func (*UnimplementedIntegrationHealthServiceServer) GetDeclarativeConfigs(ctx context.Context, req *Empty) (*GetIntegrationHealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeclarativeConfigs not implemented")
}
func (*UnimplementedIntegrationHealthServiceServer) GetVulnDefinitionsInfo(ctx context.Context, req *VulnDefinitionsInfoRequest) (*VulnDefinitionsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVulnDefinitionsInfo not implemented")
}

func RegisterIntegrationHealthServiceServer(s *grpc.Server, srv IntegrationHealthServiceServer) {
	s.RegisterService(&_IntegrationHealthService_serviceDesc, srv)
}

func _IntegrationHealthService_GetImageIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationHealthServiceServer).GetImageIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.IntegrationHealthService/GetImageIntegrations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationHealthServiceServer).GetImageIntegrations(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationHealthService_GetNotifiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationHealthServiceServer).GetNotifiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.IntegrationHealthService/GetNotifiers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationHealthServiceServer).GetNotifiers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationHealthService_GetBackupPlugins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationHealthServiceServer).GetBackupPlugins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.IntegrationHealthService/GetBackupPlugins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationHealthServiceServer).GetBackupPlugins(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationHealthService_GetDeclarativeConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationHealthServiceServer).GetDeclarativeConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.IntegrationHealthService/GetDeclarativeConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationHealthServiceServer).GetDeclarativeConfigs(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationHealthService_GetVulnDefinitionsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VulnDefinitionsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationHealthServiceServer).GetVulnDefinitionsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.IntegrationHealthService/GetVulnDefinitionsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationHealthServiceServer).GetVulnDefinitionsInfo(ctx, req.(*VulnDefinitionsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IntegrationHealthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.IntegrationHealthService",
	HandlerType: (*IntegrationHealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetImageIntegrations",
			Handler:    _IntegrationHealthService_GetImageIntegrations_Handler,
		},
		{
			MethodName: "GetNotifiers",
			Handler:    _IntegrationHealthService_GetNotifiers_Handler,
		},
		{
			MethodName: "GetBackupPlugins",
			Handler:    _IntegrationHealthService_GetBackupPlugins_Handler,
		},
		{
			MethodName: "GetDeclarativeConfigs",
			Handler:    _IntegrationHealthService_GetDeclarativeConfigs_Handler,
		},
		{
			MethodName: "GetVulnDefinitionsInfo",
			Handler:    _IntegrationHealthService_GetVulnDefinitionsInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/integration_health_service.proto",
}

func (m *GetIntegrationHealthResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetIntegrationHealthResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetIntegrationHealthResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.IntegrationHealth) > 0 {
		for iNdEx := len(m.IntegrationHealth) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IntegrationHealth[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIntegrationHealthService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *VulnDefinitionsInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VulnDefinitionsInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VulnDefinitionsInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Component != 0 {
		i = encodeVarintIntegrationHealthService(dAtA, i, uint64(m.Component))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *VulnDefinitionsInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VulnDefinitionsInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VulnDefinitionsInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.LastUpdatedTimestamp != nil {
		{
			size, err := m.LastUpdatedTimestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIntegrationHealthService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIntegrationHealthService(dAtA []byte, offset int, v uint64) int {
	offset -= sovIntegrationHealthService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetIntegrationHealthResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.IntegrationHealth) > 0 {
		for _, e := range m.IntegrationHealth {
			l = e.Size()
			n += 1 + l + sovIntegrationHealthService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *VulnDefinitionsInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Component != 0 {
		n += 1 + sovIntegrationHealthService(uint64(m.Component))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *VulnDefinitionsInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LastUpdatedTimestamp != nil {
		l = m.LastUpdatedTimestamp.Size()
		n += 1 + l + sovIntegrationHealthService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIntegrationHealthService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIntegrationHealthService(x uint64) (n int) {
	return sovIntegrationHealthService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetIntegrationHealthResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIntegrationHealthService
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
			return fmt.Errorf("proto: GetIntegrationHealthResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetIntegrationHealthResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntegrationHealth", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntegrationHealthService
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
				return ErrInvalidLengthIntegrationHealthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIntegrationHealthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IntegrationHealth = append(m.IntegrationHealth, &storage.IntegrationHealth{})
			if err := m.IntegrationHealth[len(m.IntegrationHealth)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIntegrationHealthService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIntegrationHealthService
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
func (m *VulnDefinitionsInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIntegrationHealthService
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
			return fmt.Errorf("proto: VulnDefinitionsInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VulnDefinitionsInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Component", wireType)
			}
			m.Component = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntegrationHealthService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Component |= VulnDefinitionsInfoRequest_Component(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIntegrationHealthService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIntegrationHealthService
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
func (m *VulnDefinitionsInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIntegrationHealthService
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
			return fmt.Errorf("proto: VulnDefinitionsInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VulnDefinitionsInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdatedTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntegrationHealthService
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
				return ErrInvalidLengthIntegrationHealthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIntegrationHealthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastUpdatedTimestamp == nil {
				m.LastUpdatedTimestamp = &types.Timestamp{}
			}
			if err := m.LastUpdatedTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIntegrationHealthService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIntegrationHealthService
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
func skipIntegrationHealthService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIntegrationHealthService
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
					return 0, ErrIntOverflowIntegrationHealthService
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
					return 0, ErrIntOverflowIntegrationHealthService
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
				return 0, ErrInvalidLengthIntegrationHealthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIntegrationHealthService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIntegrationHealthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIntegrationHealthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIntegrationHealthService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIntegrationHealthService = fmt.Errorf("proto: unexpected end of group")
)
