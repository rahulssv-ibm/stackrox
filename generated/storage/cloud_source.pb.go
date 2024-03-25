// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/cloud_source.proto

package storage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CloudSource_Type int32

const (
	CloudSource_TYPE_UNSPECIFIED   CloudSource_Type = 0
	CloudSource_TYPE_PALADIN_CLOUD CloudSource_Type = 1
	CloudSource_TYPE_OCM           CloudSource_Type = 2
)

var CloudSource_Type_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "TYPE_PALADIN_CLOUD",
	2: "TYPE_OCM",
}

var CloudSource_Type_value = map[string]int32{
	"TYPE_UNSPECIFIED":   0,
	"TYPE_PALADIN_CLOUD": 1,
	"TYPE_OCM":           2,
}

func (x CloudSource_Type) String() string {
	return proto.EnumName(CloudSource_Type_name, int32(x))
}

func (CloudSource_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d0f224372f8cbe44, []int{0, 0}
}

type CloudSource struct {
	Id                  string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,type(uuid)"`                                    // @gotags: sql:"pk,type(uuid)"
	Name                string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" sql:"unique" search:"Integration Name,hidden"`                                // @gotags: sql:"unique" search:"Integration Name,hidden"
	Type                CloudSource_Type         `protobuf:"varint,3,opt,name=type,proto3,enum=storage.CloudSource_Type" json:"type,omitempty" search:"Integration Type,hidden"` // @gotags: search:"Integration Type,hidden"
	Credentials         *CloudSource_Credentials `protobuf:"bytes,4,opt,name=credentials,proto3" json:"credentials,omitempty"`
	SkipTestIntegration bool                     `protobuf:"varint,5,opt,name=skip_test_integration,json=skipTestIntegration,proto3" json:"skip_test_integration,omitempty"`
	// Types that are valid to be assigned to Config:
	//	*CloudSource_PaladinCloud
	//	*CloudSource_Ocm
	Config               isCloudSource_Config `protobuf_oneof:"Config"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CloudSource) Reset()         { *m = CloudSource{} }
func (m *CloudSource) String() string { return proto.CompactTextString(m) }
func (*CloudSource) ProtoMessage()    {}
func (*CloudSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0f224372f8cbe44, []int{0}
}
func (m *CloudSource) XXX_UnmarshalVT(b []byte) error {
	return m.UnmarshalVT(b)
}
func (m *CloudSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CloudSource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CloudSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudSource.Merge(m, src)
}
func (m *CloudSource) XXX_Size() int {
	return m.Size()
}
func (m *CloudSource) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudSource.DiscardUnknown(m)
}

var xxx_messageInfo_CloudSource proto.InternalMessageInfo

type isCloudSource_Config interface {
	isCloudSource_Config()
	MarshalTo([]byte) (int, error)
	Size() int
	Clone() isCloudSource_Config
}

type CloudSource_PaladinCloud struct {
	PaladinCloud *PaladinCloudConfig `protobuf:"bytes,6,opt,name=paladin_cloud,json=paladinCloud,proto3,oneof" json:"paladin_cloud,omitempty"`
}
type CloudSource_Ocm struct {
	Ocm *OCMConfig `protobuf:"bytes,7,opt,name=ocm,proto3,oneof" json:"ocm,omitempty"`
}

func (*CloudSource_PaladinCloud) isCloudSource_Config() {}
func (m *CloudSource_PaladinCloud) Clone() isCloudSource_Config {
	if m == nil {
		return nil
	}
	cloned := new(CloudSource_PaladinCloud)
	*cloned = *m

	cloned.PaladinCloud = m.PaladinCloud.Clone()
	return cloned
}
func (*CloudSource_Ocm) isCloudSource_Config() {}
func (m *CloudSource_Ocm) Clone() isCloudSource_Config {
	if m == nil {
		return nil
	}
	cloned := new(CloudSource_Ocm)
	*cloned = *m

	cloned.Ocm = m.Ocm.Clone()
	return cloned
}

func (m *CloudSource) GetConfig() isCloudSource_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *CloudSource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CloudSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CloudSource) GetType() CloudSource_Type {
	if m != nil {
		return m.Type
	}
	return CloudSource_TYPE_UNSPECIFIED
}

func (m *CloudSource) GetCredentials() *CloudSource_Credentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (m *CloudSource) GetSkipTestIntegration() bool {
	if m != nil {
		return m.SkipTestIntegration
	}
	return false
}

func (m *CloudSource) GetPaladinCloud() *PaladinCloudConfig {
	if x, ok := m.GetConfig().(*CloudSource_PaladinCloud); ok {
		return x.PaladinCloud
	}
	return nil
}

func (m *CloudSource) GetOcm() *OCMConfig {
	if x, ok := m.GetConfig().(*CloudSource_Ocm); ok {
		return x.Ocm
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CloudSource) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CloudSource_PaladinCloud)(nil),
		(*CloudSource_Ocm)(nil),
	}
}

func (m *CloudSource) MessageClone() proto.Message {
	return m.Clone()
}
func (m *CloudSource) Clone() *CloudSource {
	if m == nil {
		return nil
	}
	cloned := new(CloudSource)
	*cloned = *m

	cloned.Credentials = m.Credentials.Clone()
	if m.Config != nil {
		cloned.Config = m.Config.Clone()
	}
	return cloned
}

type CloudSource_Credentials struct {
	Secret               string   `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty" scrub:"always"` // @gotags: scrub:"always"
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloudSource_Credentials) Reset()         { *m = CloudSource_Credentials{} }
func (m *CloudSource_Credentials) String() string { return proto.CompactTextString(m) }
func (*CloudSource_Credentials) ProtoMessage()    {}
func (*CloudSource_Credentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0f224372f8cbe44, []int{0, 0}
}
func (m *CloudSource_Credentials) XXX_UnmarshalVT(b []byte) error {
	return m.UnmarshalVT(b)
}
func (m *CloudSource_Credentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CloudSource_Credentials.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CloudSource_Credentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudSource_Credentials.Merge(m, src)
}
func (m *CloudSource_Credentials) XXX_Size() int {
	return m.Size()
}
func (m *CloudSource_Credentials) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudSource_Credentials.DiscardUnknown(m)
}

var xxx_messageInfo_CloudSource_Credentials proto.InternalMessageInfo

func (m *CloudSource_Credentials) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *CloudSource_Credentials) MessageClone() proto.Message {
	return m.Clone()
}
func (m *CloudSource_Credentials) Clone() *CloudSource_Credentials {
	if m == nil {
		return nil
	}
	cloned := new(CloudSource_Credentials)
	*cloned = *m

	return cloned
}

type PaladinCloudConfig struct {
	Endpoint             string   `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty" validate:"nolocalendpoint"` // @gotags: validate:"nolocalendpoint"
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaladinCloudConfig) Reset()         { *m = PaladinCloudConfig{} }
func (m *PaladinCloudConfig) String() string { return proto.CompactTextString(m) }
func (*PaladinCloudConfig) ProtoMessage()    {}
func (*PaladinCloudConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0f224372f8cbe44, []int{1}
}
func (m *PaladinCloudConfig) XXX_UnmarshalVT(b []byte) error {
	return m.UnmarshalVT(b)
}
func (m *PaladinCloudConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PaladinCloudConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PaladinCloudConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaladinCloudConfig.Merge(m, src)
}
func (m *PaladinCloudConfig) XXX_Size() int {
	return m.Size()
}
func (m *PaladinCloudConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PaladinCloudConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PaladinCloudConfig proto.InternalMessageInfo

func (m *PaladinCloudConfig) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *PaladinCloudConfig) MessageClone() proto.Message {
	return m.Clone()
}
func (m *PaladinCloudConfig) Clone() *PaladinCloudConfig {
	if m == nil {
		return nil
	}
	cloned := new(PaladinCloudConfig)
	*cloned = *m

	return cloned
}

type OCMConfig struct {
	Endpoint             string   `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty" validate:"nolocalendpoint"` // @gotags: validate:"nolocalendpoint"
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OCMConfig) Reset()         { *m = OCMConfig{} }
func (m *OCMConfig) String() string { return proto.CompactTextString(m) }
func (*OCMConfig) ProtoMessage()    {}
func (*OCMConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0f224372f8cbe44, []int{2}
}
func (m *OCMConfig) XXX_UnmarshalVT(b []byte) error {
	return m.UnmarshalVT(b)
}
func (m *OCMConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OCMConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OCMConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OCMConfig.Merge(m, src)
}
func (m *OCMConfig) XXX_Size() int {
	return m.Size()
}
func (m *OCMConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OCMConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OCMConfig proto.InternalMessageInfo

func (m *OCMConfig) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *OCMConfig) MessageClone() proto.Message {
	return m.Clone()
}
func (m *OCMConfig) Clone() *OCMConfig {
	if m == nil {
		return nil
	}
	cloned := new(OCMConfig)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterEnum("storage.CloudSource_Type", CloudSource_Type_name, CloudSource_Type_value)
	proto.RegisterType((*CloudSource)(nil), "storage.CloudSource")
	proto.RegisterType((*CloudSource_Credentials)(nil), "storage.CloudSource.Credentials")
	proto.RegisterType((*PaladinCloudConfig)(nil), "storage.PaladinCloudConfig")
	proto.RegisterType((*OCMConfig)(nil), "storage.OCMConfig")
}

func init() { proto.RegisterFile("storage/cloud_source.proto", fileDescriptor_d0f224372f8cbe44) }

var fileDescriptor_d0f224372f8cbe44 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0xce, 0x24, 0x21, 0x4d, 0x9f, 0x4b, 0x65, 0x3d, 0xa0, 0x72, 0x83, 0x64, 0x59, 0x96, 0x00,
	0x6f, 0x30, 0x28, 0x70, 0x81, 0xda, 0x09, 0xc2, 0x52, 0xdb, 0x58, 0x6e, 0xba, 0x80, 0x8d, 0x65,
	0xec, 0x21, 0x1a, 0x35, 0x9d, 0xb1, 0x3c, 0x53, 0x89, 0xde, 0x84, 0xab, 0x70, 0x03, 0x96, 0x1c,
	0x01, 0x85, 0x8b, 0x20, 0x4f, 0x1c, 0xc7, 0x12, 0xa8, 0x3b, 0xcf, 0xf7, 0xf3, 0xbe, 0xe7, 0x4f,
	0x0f, 0x26, 0x52, 0x89, 0x2a, 0x5b, 0xd1, 0x37, 0xf9, 0x5a, 0xdc, 0x15, 0xa9, 0x14, 0x77, 0x55,
	0x4e, 0xfd, 0xb2, 0x12, 0x4a, 0xe0, 0x41, 0xc3, 0xb9, 0x3f, 0x06, 0x60, 0x84, 0x35, 0x7f, 0xa5,
	0x69, 0x3c, 0x86, 0x3e, 0x2b, 0x2c, 0xe2, 0x10, 0xef, 0x30, 0xe9, 0xb3, 0x02, 0x11, 0x86, 0x3c,
	0xbb, 0xa5, 0x56, 0x5f, 0x23, 0xfa, 0x1b, 0x5f, 0xc3, 0x50, 0xdd, 0x97, 0xd4, 0x1a, 0x38, 0xc4,
	0x3b, 0x9e, 0x9e, 0xfa, 0xcd, 0x2c, 0xbf, 0x33, 0xc7, 0x5f, 0xde, 0x97, 0x34, 0xd1, 0x32, 0x0c,
	0xc0, 0xc8, 0x2b, 0x5a, 0x50, 0xae, 0x58, 0xb6, 0x96, 0xd6, 0xd0, 0x21, 0x9e, 0x31, 0x75, 0xfe,
	0xeb, 0x0a, 0xf7, 0xba, 0xa4, 0x6b, 0xc2, 0x29, 0x3c, 0x93, 0x37, 0xac, 0x4c, 0x15, 0x95, 0x2a,
	0x65, 0x5c, 0xd1, 0x55, 0x95, 0x29, 0x26, 0xb8, 0xf5, 0xc8, 0x21, 0xde, 0x38, 0x79, 0x52, 0x93,
	0x4b, 0x2a, 0x55, 0xb4, 0xa7, 0x30, 0x80, 0xc7, 0x65, 0xb6, 0xce, 0x0a, 0xc6, 0x53, 0xdd, 0x80,
	0x35, 0xd2, 0xc9, 0xcf, 0xdb, 0xe4, 0x78, 0xcb, 0xea, 0x05, 0x42, 0xc1, 0xbf, 0xb2, 0xd5, 0xc7,
	0x5e, 0x72, 0x54, 0x76, 0x50, 0x7c, 0x09, 0x03, 0x91, 0xdf, 0x5a, 0x07, 0xda, 0x89, 0xad, 0x73,
	0x11, 0x5e, 0xb4, 0x86, 0x5a, 0x30, 0x79, 0x01, 0x46, 0x67, 0x77, 0x3c, 0x81, 0x91, 0xa4, 0x79,
	0x45, 0x55, 0xd3, 0x64, 0xf3, 0x72, 0x03, 0x18, 0xd6, 0xc5, 0xe0, 0x53, 0x30, 0x97, 0x9f, 0xe2,
	0x79, 0x7a, 0x7d, 0x79, 0x15, 0xcf, 0xc3, 0xe8, 0x43, 0x34, 0x9f, 0x99, 0x3d, 0x3c, 0x01, 0xd4,
	0x68, 0x7c, 0x76, 0x7e, 0x36, 0x8b, 0x2e, 0xd3, 0xf0, 0x7c, 0x71, 0x3d, 0x33, 0x09, 0x1e, 0xc1,
	0x58, 0xe3, 0x8b, 0xf0, 0xc2, 0xec, 0x07, 0x63, 0x18, 0x6d, 0xb3, 0xdd, 0xb7, 0x80, 0xff, 0xfe,
	0x02, 0x4e, 0x60, 0x4c, 0x79, 0x51, 0x0a, 0xc6, 0x77, 0xe9, 0xed, 0xdb, 0x7d, 0x05, 0x87, 0xed,
	0xea, 0x0f, 0x09, 0x83, 0xf7, 0x3f, 0x37, 0x36, 0xf9, 0xb5, 0xb1, 0xc9, 0xef, 0x8d, 0x4d, 0xbe,
	0xff, 0xb1, 0x7b, 0x70, 0xca, 0x84, 0x2f, 0x55, 0x96, 0xdf, 0x54, 0xe2, 0xdb, 0xf6, 0x88, 0x76,
	0x6d, 0x7c, 0xde, 0x1d, 0xd3, 0x97, 0x91, 0xc6, 0xdf, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x40,
	0x5b, 0x5a, 0xbf, 0x7a, 0x02, 0x00, 0x00,
}

func (m *CloudSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CloudSource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CloudSource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Config != nil {
		{
			size := m.Config.Size()
			i -= size
			if _, err := m.Config.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.SkipTestIntegration {
		i--
		if m.SkipTestIntegration {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Credentials != nil {
		{
			size, err := m.Credentials.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCloudSource(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Type != 0 {
		i = encodeVarintCloudSource(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCloudSource(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintCloudSource(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CloudSource_PaladinCloud) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CloudSource_PaladinCloud) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PaladinCloud != nil {
		{
			size, err := m.PaladinCloud.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCloudSource(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *CloudSource_Ocm) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CloudSource_Ocm) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Ocm != nil {
		{
			size, err := m.Ocm.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCloudSource(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *CloudSource_Credentials) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CloudSource_Credentials) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CloudSource_Credentials) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Secret) > 0 {
		i -= len(m.Secret)
		copy(dAtA[i:], m.Secret)
		i = encodeVarintCloudSource(dAtA, i, uint64(len(m.Secret)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PaladinCloudConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PaladinCloudConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PaladinCloudConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Endpoint) > 0 {
		i -= len(m.Endpoint)
		copy(dAtA[i:], m.Endpoint)
		i = encodeVarintCloudSource(dAtA, i, uint64(len(m.Endpoint)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OCMConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OCMConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OCMConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Endpoint) > 0 {
		i -= len(m.Endpoint)
		copy(dAtA[i:], m.Endpoint)
		i = encodeVarintCloudSource(dAtA, i, uint64(len(m.Endpoint)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCloudSource(dAtA []byte, offset int, v uint64) int {
	offset -= sovCloudSource(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CloudSource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCloudSource(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCloudSource(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovCloudSource(uint64(m.Type))
	}
	if m.Credentials != nil {
		l = m.Credentials.Size()
		n += 1 + l + sovCloudSource(uint64(l))
	}
	if m.SkipTestIntegration {
		n += 2
	}
	if m.Config != nil {
		n += m.Config.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CloudSource_PaladinCloud) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PaladinCloud != nil {
		l = m.PaladinCloud.Size()
		n += 1 + l + sovCloudSource(uint64(l))
	}
	return n
}
func (m *CloudSource_Ocm) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ocm != nil {
		l = m.Ocm.Size()
		n += 1 + l + sovCloudSource(uint64(l))
	}
	return n
}
func (m *CloudSource_Credentials) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovCloudSource(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PaladinCloudConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Endpoint)
	if l > 0 {
		n += 1 + l + sovCloudSource(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *OCMConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Endpoint)
	if l > 0 {
		n += 1 + l + sovCloudSource(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCloudSource(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCloudSource(x uint64) (n int) {
	return sovCloudSource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CloudSource) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCloudSource
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
			return fmt.Errorf("proto: CloudSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CloudSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloudSource
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
				return ErrInvalidLengthCloudSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCloudSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloudSource
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
				return ErrInvalidLengthCloudSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCloudSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloudSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= CloudSource_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credentials", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloudSource
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
				return ErrInvalidLengthCloudSource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCloudSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Credentials == nil {
				m.Credentials = &CloudSource_Credentials{}
			}
			if err := m.Credentials.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SkipTestIntegration", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloudSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SkipTestIntegration = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaladinCloud", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloudSource
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
				return ErrInvalidLengthCloudSource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCloudSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PaladinCloudConfig{}
			if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Config = &CloudSource_PaladinCloud{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ocm", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloudSource
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
				return ErrInvalidLengthCloudSource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCloudSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &OCMConfig{}
			if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Config = &CloudSource_Ocm{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCloudSource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCloudSource
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
func (m *CloudSource_Credentials) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCloudSource
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
			return fmt.Errorf("proto: Credentials: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Credentials: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloudSource
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
				return ErrInvalidLengthCloudSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCloudSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCloudSource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCloudSource
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
func (m *PaladinCloudConfig) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCloudSource
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
			return fmt.Errorf("proto: PaladinCloudConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PaladinCloudConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloudSource
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
				return ErrInvalidLengthCloudSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCloudSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCloudSource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCloudSource
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
func (m *OCMConfig) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCloudSource
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
			return fmt.Errorf("proto: OCMConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OCMConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloudSource
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
				return ErrInvalidLengthCloudSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCloudSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCloudSource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCloudSource
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
func skipCloudSource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCloudSource
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
					return 0, ErrIntOverflowCloudSource
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
					return 0, ErrIntOverflowCloudSource
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
				return 0, ErrInvalidLengthCloudSource
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCloudSource
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCloudSource
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCloudSource        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCloudSource          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCloudSource = fmt.Errorf("proto: unexpected end of group")
)
