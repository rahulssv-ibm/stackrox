// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/image_integration_service.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	storage "github.com/stackrox/rox/generated/storage"
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

type GetImageIntegrationsRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cluster              string   `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetImageIntegrationsRequest) Reset()         { *m = GetImageIntegrationsRequest{} }
func (m *GetImageIntegrationsRequest) String() string { return proto.CompactTextString(m) }
func (*GetImageIntegrationsRequest) ProtoMessage()    {}
func (*GetImageIntegrationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_263887791ba1d8ff, []int{0}
}
func (m *GetImageIntegrationsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetImageIntegrationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetImageIntegrationsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetImageIntegrationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetImageIntegrationsRequest.Merge(m, src)
}
func (m *GetImageIntegrationsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetImageIntegrationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetImageIntegrationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetImageIntegrationsRequest proto.InternalMessageInfo

func (m *GetImageIntegrationsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetImageIntegrationsRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *GetImageIntegrationsRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GetImageIntegrationsRequest) Clone() *GetImageIntegrationsRequest {
	if m == nil {
		return nil
	}
	cloned := new(GetImageIntegrationsRequest)
	*cloned = *m

	return cloned
}

type GetImageIntegrationsResponse struct {
	Integrations         []*storage.ImageIntegration `protobuf:"bytes,1,rep,name=integrations,proto3" json:"integrations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *GetImageIntegrationsResponse) Reset()         { *m = GetImageIntegrationsResponse{} }
func (m *GetImageIntegrationsResponse) String() string { return proto.CompactTextString(m) }
func (*GetImageIntegrationsResponse) ProtoMessage()    {}
func (*GetImageIntegrationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_263887791ba1d8ff, []int{1}
}
func (m *GetImageIntegrationsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetImageIntegrationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetImageIntegrationsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetImageIntegrationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetImageIntegrationsResponse.Merge(m, src)
}
func (m *GetImageIntegrationsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetImageIntegrationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetImageIntegrationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetImageIntegrationsResponse proto.InternalMessageInfo

func (m *GetImageIntegrationsResponse) GetIntegrations() []*storage.ImageIntegration {
	if m != nil {
		return m.Integrations
	}
	return nil
}

func (m *GetImageIntegrationsResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GetImageIntegrationsResponse) Clone() *GetImageIntegrationsResponse {
	if m == nil {
		return nil
	}
	cloned := new(GetImageIntegrationsResponse)
	*cloned = *m

	if m.Integrations != nil {
		cloned.Integrations = make([]*storage.ImageIntegration, len(m.Integrations))
		for idx, v := range m.Integrations {
			cloned.Integrations[idx] = v.Clone()
		}
	}
	return cloned
}

type UpdateImageIntegrationRequest struct {
	Config *storage.ImageIntegration `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// When false, use the stored credentials of an existing image integration given its ID.
	UpdatePassword       bool     `protobuf:"varint,2,opt,name=updatePassword,proto3" json:"updatePassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateImageIntegrationRequest) Reset()         { *m = UpdateImageIntegrationRequest{} }
func (m *UpdateImageIntegrationRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateImageIntegrationRequest) ProtoMessage()    {}
func (*UpdateImageIntegrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_263887791ba1d8ff, []int{2}
}
func (m *UpdateImageIntegrationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateImageIntegrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateImageIntegrationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateImageIntegrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateImageIntegrationRequest.Merge(m, src)
}
func (m *UpdateImageIntegrationRequest) XXX_Size() int {
	return m.Size()
}
func (m *UpdateImageIntegrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateImageIntegrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateImageIntegrationRequest proto.InternalMessageInfo

func (m *UpdateImageIntegrationRequest) GetConfig() *storage.ImageIntegration {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *UpdateImageIntegrationRequest) GetUpdatePassword() bool {
	if m != nil {
		return m.UpdatePassword
	}
	return false
}

func (m *UpdateImageIntegrationRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *UpdateImageIntegrationRequest) Clone() *UpdateImageIntegrationRequest {
	if m == nil {
		return nil
	}
	cloned := new(UpdateImageIntegrationRequest)
	*cloned = *m

	cloned.Config = m.Config.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*GetImageIntegrationsRequest)(nil), "v1.GetImageIntegrationsRequest")
	proto.RegisterType((*GetImageIntegrationsResponse)(nil), "v1.GetImageIntegrationsResponse")
	proto.RegisterType((*UpdateImageIntegrationRequest)(nil), "v1.UpdateImageIntegrationRequest")
}

func init() {
	proto.RegisterFile("api/v1/image_integration_service.proto", fileDescriptor_263887791ba1d8ff)
}

var fileDescriptor_263887791ba1d8ff = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x51, 0x6b, 0xd3, 0x50,
	0x14, 0xc7, 0x4d, 0x94, 0xcd, 0x5d, 0x45, 0xe4, 0xb6, 0x6e, 0x31, 0xeb, 0xda, 0xee, 0x82, 0x75,
	0x14, 0xbd, 0x21, 0xf5, 0x6d, 0xe0, 0xcb, 0x98, 0x48, 0xf1, 0xa5, 0x44, 0x05, 0x15, 0x64, 0xdc,
	0x25, 0x67, 0x21, 0xd8, 0xe4, 0xc6, 0xdc, 0x9b, 0xb8, 0x4d, 0x7c, 0xd0, 0xaf, 0xe0, 0x8b, 0x1f,
	0xc9, 0x47, 0xc1, 0x2f, 0x20, 0xd5, 0x0f, 0x22, 0xb9, 0x89, 0x5a, 0xdb, 0x24, 0x30, 0xfa, 0xd6,
	0x9e, 0x73, 0xfa, 0xff, 0x9d, 0xff, 0xb9, 0xe7, 0x14, 0x0d, 0x58, 0x1c, 0x58, 0x99, 0x6d, 0x05,
	0x21, 0xf3, 0xe1, 0x28, 0x88, 0x24, 0xf8, 0x09, 0x93, 0x01, 0x8f, 0x8e, 0x04, 0x24, 0x59, 0xe0,
	0x02, 0x8d, 0x13, 0x2e, 0x39, 0xd6, 0x33, 0xdb, 0x6c, 0x95, 0xb5, 0x2e, 0x0f, 0x43, 0x1e, 0x15,
	0x09, 0x13, 0x97, 0x41, 0x08, 0x63, 0x79, 0x56, 0xc6, 0x3a, 0x3e, 0xe7, 0xfe, 0x14, 0xac, 0x3c,
	0xc5, 0xa2, 0x88, 0x4b, 0xa5, 0x28, 0xca, 0x6c, 0x4f, 0x48, 0x9e, 0x30, 0x1f, 0x96, 0x99, 0x45,
	0x01, 0x79, 0x82, 0xb6, 0x1f, 0x83, 0x1c, 0xe7, 0xd9, 0xf1, 0xbf, 0xa4, 0x70, 0xe0, 0x6d, 0x0a,
	0x42, 0x62, 0x8c, 0xae, 0x44, 0x2c, 0x04, 0x43, 0xeb, 0x6b, 0x7b, 0x1b, 0x8e, 0xfa, 0x8c, 0x0d,
	0xb4, 0xee, 0x4e, 0x53, 0x21, 0x21, 0x31, 0x74, 0x15, 0xfe, 0xf3, 0x95, 0xbc, 0x46, 0x9d, 0x6a,
	0x31, 0x11, 0xf3, 0x48, 0x00, 0x7e, 0x88, 0xae, 0xcf, 0x75, 0x20, 0x0c, 0xad, 0x7f, 0x79, 0xef,
	0xda, 0xe8, 0x36, 0x2d, 0x9b, 0xa4, 0x8b, 0xbf, 0x74, 0xfe, 0x2b, 0x27, 0xe7, 0x68, 0xe7, 0x79,
	0xec, 0x31, 0x09, 0x4b, 0x75, 0x65, 0xb7, 0x36, 0x5a, 0x73, 0x79, 0x74, 0x12, 0xf8, 0xaa, 0xdf,
	0x46, 0xe5, 0xb2, 0x10, 0x0f, 0xd0, 0x8d, 0x54, 0x69, 0x4e, 0x98, 0x10, 0xef, 0x78, 0xe2, 0x29,
	0x4f, 0x57, 0x9d, 0x85, 0xe8, 0xe8, 0xe3, 0x3a, 0xda, 0x5a, 0x14, 0x79, 0x5a, 0xbc, 0x1a, 0x3e,
	0x41, 0xad, 0x0a, 0xdb, 0xf8, 0x26, 0xcd, 0x6c, 0xea, 0x80, 0xe0, 0x69, 0xe2, 0xc2, 0xc1, 0xd9,
	0xf8, 0xd0, 0xac, 0xef, 0x87, 0x90, 0x4f, 0xdf, 0x7f, 0x7d, 0xd6, 0x3b, 0xd8, 0xfc, 0xbb, 0x20,
	0xf3, 0xd6, 0xad, 0xf7, 0x81, 0xf7, 0x01, 0x9f, 0xa2, 0x76, 0xd5, 0x78, 0x71, 0x2f, 0x07, 0x35,
	0xbc, 0xa2, 0xd9, 0xaf, 0x2f, 0x28, 0x5e, 0x86, 0xec, 0x28, 0xfc, 0x16, 0xbe, 0x55, 0x89, 0xc7,
	0x11, 0x6a, 0x4f, 0xb8, 0x58, 0xb6, 0x58, 0x6f, 0xa8, 0xc9, 0x6b, 0x5f, 0xc1, 0x4c, 0x52, 0x0d,
	0xdb, 0xd7, 0x86, 0xf8, 0x18, 0xb5, 0x26, 0xe9, 0x85, 0x70, 0x1b, 0xb9, 0xc5, 0x47, 0xf9, 0x5d,
	0x90, 0x3b, 0x4a, 0xbe, 0x67, 0x36, 0x8c, 0x32, 0x67, 0xb8, 0xa8, 0xfd, 0x0c, 0xc4, 0x8a, 0x10,
	0x52, 0x03, 0x91, 0x20, 0x64, 0x0e, 0x79, 0x89, 0x36, 0x0f, 0x61, 0x0a, 0xcb, 0x2b, 0x5b, 0xb1,
	0x1d, 0x73, 0xea, 0xe5, 0x36, 0x0c, 0x9b, 0xb6, 0x41, 0xa2, 0xcd, 0xea, 0x6b, 0xc0, 0xbb, 0xb9,
	0x50, 0xe3, 0xa5, 0xcc, 0xb3, 0xee, 0x29, 0xd6, 0x60, 0xb4, 0x5b, 0xc3, 0x2a, 0x0e, 0x85, 0x96,
	0x53, 0x3b, 0x47, 0xdb, 0xf9, 0xd4, 0x0a, 0x75, 0x6f, 0x45, 0xf4, 0x7d, 0x85, 0xbe, 0x4b, 0x48,
	0xfd, 0x10, 0xad, 0xe2, 0x0e, 0xbd, 0x7d, 0x6d, 0x78, 0x40, 0xbf, 0xce, 0xba, 0xda, 0xb7, 0x59,
	0x57, 0xfb, 0x31, 0xeb, 0x6a, 0x5f, 0x7e, 0x76, 0x2f, 0x21, 0x23, 0xe0, 0x54, 0x48, 0xe6, 0xbe,
	0x49, 0xf8, 0x69, 0xf1, 0x87, 0x46, 0x59, 0x1c, 0xd0, 0xcc, 0x7e, 0xa5, 0x67, 0xf6, 0x0b, 0xfd,
	0x78, 0x4d, 0xc5, 0x1e, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xad, 0x34, 0xe3, 0xc7, 0x7a, 0x05,
	0x00, 0x00,
}

func (m *GetImageIntegrationsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetImageIntegrationsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetImageIntegrationsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Cluster) > 0 {
		i -= len(m.Cluster)
		copy(dAtA[i:], m.Cluster)
		i = encodeVarintImageIntegrationService(dAtA, i, uint64(len(m.Cluster)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintImageIntegrationService(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetImageIntegrationsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetImageIntegrationsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetImageIntegrationsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Integrations) > 0 {
		for iNdEx := len(m.Integrations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Integrations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintImageIntegrationService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *UpdateImageIntegrationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateImageIntegrationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateImageIntegrationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.UpdatePassword {
		i--
		if m.UpdatePassword {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintImageIntegrationService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintImageIntegrationService(dAtA []byte, offset int, v uint64) int {
	offset -= sovImageIntegrationService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetImageIntegrationsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovImageIntegrationService(uint64(l))
	}
	l = len(m.Cluster)
	if l > 0 {
		n += 1 + l + sovImageIntegrationService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetImageIntegrationsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Integrations) > 0 {
		for _, e := range m.Integrations {
			l = e.Size()
			n += 1 + l + sovImageIntegrationService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UpdateImageIntegrationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovImageIntegrationService(uint64(l))
	}
	if m.UpdatePassword {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovImageIntegrationService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozImageIntegrationService(x uint64) (n int) {
	return sovImageIntegrationService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetImageIntegrationsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImageIntegrationService
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
			return fmt.Errorf("proto: GetImageIntegrationsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetImageIntegrationsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageIntegrationService
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
				return ErrInvalidLengthImageIntegrationService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImageIntegrationService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageIntegrationService
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
				return ErrInvalidLengthImageIntegrationService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImageIntegrationService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImageIntegrationService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthImageIntegrationService
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
func (m *GetImageIntegrationsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImageIntegrationService
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
			return fmt.Errorf("proto: GetImageIntegrationsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetImageIntegrationsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Integrations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageIntegrationService
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
				return ErrInvalidLengthImageIntegrationService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthImageIntegrationService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Integrations = append(m.Integrations, &storage.ImageIntegration{})
			if err := m.Integrations[len(m.Integrations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImageIntegrationService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthImageIntegrationService
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
func (m *UpdateImageIntegrationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImageIntegrationService
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
			return fmt.Errorf("proto: UpdateImageIntegrationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateImageIntegrationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageIntegrationService
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
				return ErrInvalidLengthImageIntegrationService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthImageIntegrationService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &storage.ImageIntegration{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatePassword", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageIntegrationService
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
			m.UpdatePassword = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipImageIntegrationService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthImageIntegrationService
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
func skipImageIntegrationService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowImageIntegrationService
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
					return 0, ErrIntOverflowImageIntegrationService
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
					return 0, ErrIntOverflowImageIntegrationService
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
				return 0, ErrInvalidLengthImageIntegrationService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupImageIntegrationService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthImageIntegrationService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthImageIntegrationService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowImageIntegrationService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupImageIntegrationService = fmt.Errorf("proto: unexpected end of group")
)
