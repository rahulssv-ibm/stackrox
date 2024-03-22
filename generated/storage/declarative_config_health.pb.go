// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/declarative_config_health.proto

package storage

import (
	fmt "fmt"
	types "github.com/gogo/protobuf/types"
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

type DeclarativeConfigHealth_Status int32

const (
	DeclarativeConfigHealth_UNHEALTHY DeclarativeConfigHealth_Status = 0
	DeclarativeConfigHealth_HEALTHY   DeclarativeConfigHealth_Status = 1
)

var DeclarativeConfigHealth_Status_name = map[int32]string{
	0: "UNHEALTHY",
	1: "HEALTHY",
}

var DeclarativeConfigHealth_Status_value = map[string]int32{
	"UNHEALTHY": 0,
	"HEALTHY":   1,
}

func (x DeclarativeConfigHealth_Status) String() string {
	return proto.EnumName(DeclarativeConfigHealth_Status_name, int32(x))
}

func (DeclarativeConfigHealth_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_671c3ab7741e96a4, []int{0, 0}
}

type DeclarativeConfigHealth_ResourceType int32

const (
	DeclarativeConfigHealth_CONFIG_MAP     DeclarativeConfigHealth_ResourceType = 0
	DeclarativeConfigHealth_ACCESS_SCOPE   DeclarativeConfigHealth_ResourceType = 1
	DeclarativeConfigHealth_PERMISSION_SET DeclarativeConfigHealth_ResourceType = 2
	DeclarativeConfigHealth_ROLE           DeclarativeConfigHealth_ResourceType = 3
	DeclarativeConfigHealth_AUTH_PROVIDER  DeclarativeConfigHealth_ResourceType = 4
	DeclarativeConfigHealth_GROUP          DeclarativeConfigHealth_ResourceType = 5
	DeclarativeConfigHealth_NOTIFIER       DeclarativeConfigHealth_ResourceType = 6
)

var DeclarativeConfigHealth_ResourceType_name = map[int32]string{
	0: "CONFIG_MAP",
	1: "ACCESS_SCOPE",
	2: "PERMISSION_SET",
	3: "ROLE",
	4: "AUTH_PROVIDER",
	5: "GROUP",
	6: "NOTIFIER",
}

var DeclarativeConfigHealth_ResourceType_value = map[string]int32{
	"CONFIG_MAP":     0,
	"ACCESS_SCOPE":   1,
	"PERMISSION_SET": 2,
	"ROLE":           3,
	"AUTH_PROVIDER":  4,
	"GROUP":          5,
	"NOTIFIER":       6,
}

func (x DeclarativeConfigHealth_ResourceType) String() string {
	return proto.EnumName(DeclarativeConfigHealth_ResourceType_name, int32(x))
}

func (DeclarativeConfigHealth_ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_671c3ab7741e96a4, []int{0, 1}
}

type DeclarativeConfigHealth struct {
	Id           string                               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,type(uuid)"` // @gotags: sql:"pk,type(uuid)"
	Name         string                               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status       DeclarativeConfigHealth_Status       `protobuf:"varint,4,opt,name=status,proto3,enum=storage.DeclarativeConfigHealth_Status" json:"status,omitempty"`
	ErrorMessage string                               `protobuf:"bytes,5,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	ResourceName string                               `protobuf:"bytes,6,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	ResourceType DeclarativeConfigHealth_ResourceType `protobuf:"varint,7,opt,name=resource_type,json=resourceType,proto3,enum=storage.DeclarativeConfigHealth_ResourceType" json:"resource_type,omitempty"`
	// Timestamp when the current status was set.
	LastTimestamp        *types.Timestamp `protobuf:"bytes,8,opt,name=last_timestamp,json=lastTimestamp,proto3" json:"last_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DeclarativeConfigHealth) Reset()         { *m = DeclarativeConfigHealth{} }
func (m *DeclarativeConfigHealth) String() string { return proto.CompactTextString(m) }
func (*DeclarativeConfigHealth) ProtoMessage()    {}
func (*DeclarativeConfigHealth) Descriptor() ([]byte, []int) {
	return fileDescriptor_671c3ab7741e96a4, []int{0}
}
func (m *DeclarativeConfigHealth) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeclarativeConfigHealth) XXX_MarshalVT(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeclarativeConfigHealth.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeclarativeConfigHealth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeclarativeConfigHealth.Merge(m, src)
}
func (m *DeclarativeConfigHealth) XXX_Size() int {
	return m.Size()
}
func (m *DeclarativeConfigHealth) XXX_DiscardUnknown() {
	xxx_messageInfo_DeclarativeConfigHealth.DiscardUnknown(m)
}

var xxx_messageInfo_DeclarativeConfigHealth proto.InternalMessageInfo

func (m *DeclarativeConfigHealth) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeclarativeConfigHealth) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeclarativeConfigHealth) GetStatus() DeclarativeConfigHealth_Status {
	if m != nil {
		return m.Status
	}
	return DeclarativeConfigHealth_UNHEALTHY
}

func (m *DeclarativeConfigHealth) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *DeclarativeConfigHealth) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *DeclarativeConfigHealth) GetResourceType() DeclarativeConfigHealth_ResourceType {
	if m != nil {
		return m.ResourceType
	}
	return DeclarativeConfigHealth_CONFIG_MAP
}

func (m *DeclarativeConfigHealth) GetLastTimestamp() *types.Timestamp {
	if m != nil {
		return m.LastTimestamp
	}
	return nil
}

func (m *DeclarativeConfigHealth) MessageClone() proto.Message {
	return m.Clone()
}
func (m *DeclarativeConfigHealth) Clone() *DeclarativeConfigHealth {
	if m == nil {
		return nil
	}
	cloned := new(DeclarativeConfigHealth)
	*cloned = *m

	cloned.LastTimestamp = m.LastTimestamp.Clone()
	return cloned
}

func init() {
	proto.RegisterEnum("storage.DeclarativeConfigHealth_Status", DeclarativeConfigHealth_Status_name, DeclarativeConfigHealth_Status_value)
	proto.RegisterEnum("storage.DeclarativeConfigHealth_ResourceType", DeclarativeConfigHealth_ResourceType_name, DeclarativeConfigHealth_ResourceType_value)
	proto.RegisterType((*DeclarativeConfigHealth)(nil), "storage.DeclarativeConfigHealth")
}

func init() {
	proto.RegisterFile("storage/declarative_config_health.proto", fileDescriptor_671c3ab7741e96a4)
}

var fileDescriptor_671c3ab7741e96a4 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x14, 0x45, 0xe3, 0xe0, 0x38, 0xc9, 0x6b, 0x62, 0x99, 0xd9, 0x60, 0xba, 0x08, 0x51, 0x40, 0x6a,
	0x36, 0x38, 0x52, 0x61, 0x8f, 0x8c, 0xeb, 0x36, 0x96, 0x1a, 0xdb, 0x1a, 0x3b, 0x48, 0xb0, 0x19,
	0x4d, 0x93, 0xa9, 0x6b, 0x91, 0x74, 0xa2, 0x99, 0x09, 0xa2, 0xec, 0xf8, 0x0b, 0x3e, 0x89, 0x25,
	0x9f, 0x80, 0xc2, 0x8f, 0xa0, 0x8c, 0xe3, 0xb6, 0x1b, 0xc4, 0xce, 0xef, 0xea, 0xbc, 0xeb, 0xfb,
	0xae, 0x06, 0x4e, 0xa4, 0xe2, 0x82, 0x16, 0x6c, 0xb2, 0x64, 0x8b, 0x15, 0x15, 0x54, 0x95, 0x5f,
	0x18, 0x59, 0xf0, 0xdb, 0xeb, 0xb2, 0x20, 0x37, 0x8c, 0xae, 0xd4, 0x8d, 0xb7, 0x11, 0x5c, 0x71,
	0xd4, 0x3e, 0x80, 0xc7, 0x2f, 0x0a, 0xce, 0x8b, 0x15, 0x9b, 0x68, 0xf9, 0x6a, 0x7b, 0x3d, 0x51,
	0xe5, 0x9a, 0x49, 0x45, 0xd7, 0x9b, 0x8a, 0x1c, 0x7d, 0x37, 0xe1, 0xd9, 0xd9, 0x83, 0x5b, 0xa0,
	0xcd, 0xa6, 0xda, 0x0b, 0xd9, 0xd0, 0x2c, 0x97, 0xae, 0x31, 0x34, 0xc6, 0x5d, 0xdc, 0x2c, 0x97,
	0x08, 0x81, 0x79, 0x4b, 0xd7, 0xcc, 0x6d, 0x6a, 0x45, 0x7f, 0xa3, 0x77, 0x60, 0x49, 0x45, 0xd5,
	0x56, 0xba, 0xe6, 0xd0, 0x18, 0xdb, 0xa7, 0x27, 0xde, 0xe1, 0xd7, 0xde, 0x3f, 0x5c, 0xbd, 0x4c,
	0xe3, 0xf8, 0xb0, 0x86, 0x5e, 0x42, 0x9f, 0x09, 0xc1, 0x05, 0x59, 0x33, 0x29, 0x69, 0xc1, 0xdc,
	0x96, 0x76, 0xef, 0x69, 0x71, 0x56, 0x69, 0x7b, 0x48, 0x30, 0xc9, 0xb7, 0x62, 0xc1, 0x88, 0x8e,
	0x60, 0x55, 0x50, 0x2d, 0xc6, 0xfb, 0x28, 0xf8, 0x11, 0xa4, 0xee, 0x36, 0xcc, 0x6d, 0xeb, 0x44,
	0xaf, 0xff, 0x9b, 0x08, 0x1f, 0xb6, 0xf2, 0xbb, 0x0d, 0x7b, 0xf0, 0xdc, 0x4f, 0xc8, 0x07, 0x7b,
	0x45, 0xa5, 0x22, 0xf7, 0xb5, 0xb9, 0x9d, 0xa1, 0x31, 0x3e, 0x3a, 0x3d, 0xf6, 0xaa, 0x62, 0xbd,
	0xba, 0x58, 0x2f, 0xaf, 0x09, 0xdc, 0xdf, 0x6f, 0xdc, 0x8f, 0xa3, 0x57, 0x60, 0x55, 0x27, 0xa3,
	0x3e, 0x74, 0xe7, 0xf1, 0x34, 0xf4, 0x2f, 0xf3, 0xe9, 0x47, 0xa7, 0x81, 0x8e, 0xa0, 0x5d, 0x0f,
	0xc6, 0xe8, 0x1b, 0xf4, 0x1e, 0xc7, 0x40, 0x36, 0x40, 0x90, 0xc4, 0xe7, 0xd1, 0x05, 0x99, 0xf9,
	0xa9, 0xd3, 0x40, 0x0e, 0xf4, 0xfc, 0x20, 0x08, 0xb3, 0x8c, 0x64, 0x41, 0x92, 0x86, 0x8e, 0x81,
	0x10, 0xd8, 0x69, 0x88, 0x67, 0x51, 0x96, 0x45, 0x49, 0x4c, 0xb2, 0x30, 0x77, 0x9a, 0xa8, 0x03,
	0x26, 0x4e, 0x2e, 0x43, 0xe7, 0x09, 0x7a, 0x0a, 0x7d, 0x7f, 0x9e, 0x4f, 0x49, 0x8a, 0x93, 0x0f,
	0xd1, 0x59, 0x88, 0x1d, 0x13, 0x75, 0xa1, 0x75, 0x81, 0x93, 0x79, 0xea, 0xb4, 0x50, 0x0f, 0x3a,
	0x71, 0x92, 0x47, 0xe7, 0x51, 0x88, 0x1d, 0xeb, 0xfd, 0xdb, 0x9f, 0xbb, 0x81, 0xf1, 0x6b, 0x37,
	0x30, 0x7e, 0xef, 0x06, 0xc6, 0x8f, 0x3f, 0x83, 0x06, 0x3c, 0x2f, 0xb9, 0x27, 0x15, 0x5d, 0x7c,
	0x16, 0xfc, 0x6b, 0x75, 0x62, 0x5d, 0xe2, 0xa7, 0xfa, 0x69, 0x5d, 0x59, 0x5a, 0x7f, 0xf3, 0x37,
	0x00, 0x00, 0xff, 0xff, 0xb3, 0x60, 0x0a, 0xa6, 0x95, 0x02, 0x00, 0x00,
}

func (m *DeclarativeConfigHealth) MarshalVT() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeclarativeConfigHealth) MarshalVTTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeclarativeConfigHealth) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.LastTimestamp != nil {
		{
			size, err := m.LastTimestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDeclarativeConfigHealth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.ResourceType != 0 {
		i = encodeVarintDeclarativeConfigHealth(dAtA, i, uint64(m.ResourceType))
		i--
		dAtA[i] = 0x38
	}
	if len(m.ResourceName) > 0 {
		i -= len(m.ResourceName)
		copy(dAtA[i:], m.ResourceName)
		i = encodeVarintDeclarativeConfigHealth(dAtA, i, uint64(len(m.ResourceName)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ErrorMessage) > 0 {
		i -= len(m.ErrorMessage)
		copy(dAtA[i:], m.ErrorMessage)
		i = encodeVarintDeclarativeConfigHealth(dAtA, i, uint64(len(m.ErrorMessage)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Status != 0 {
		i = encodeVarintDeclarativeConfigHealth(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDeclarativeConfigHealth(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDeclarativeConfigHealth(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeclarativeConfigHealth(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeclarativeConfigHealth(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeclarativeConfigHealth) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDeclarativeConfigHealth(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDeclarativeConfigHealth(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovDeclarativeConfigHealth(uint64(m.Status))
	}
	l = len(m.ErrorMessage)
	if l > 0 {
		n += 1 + l + sovDeclarativeConfigHealth(uint64(l))
	}
	l = len(m.ResourceName)
	if l > 0 {
		n += 1 + l + sovDeclarativeConfigHealth(uint64(l))
	}
	if m.ResourceType != 0 {
		n += 1 + sovDeclarativeConfigHealth(uint64(m.ResourceType))
	}
	if m.LastTimestamp != nil {
		l = m.LastTimestamp.Size()
		n += 1 + l + sovDeclarativeConfigHealth(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDeclarativeConfigHealth(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeclarativeConfigHealth(x uint64) (n int) {
	return sovDeclarativeConfigHealth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeclarativeConfigHealth) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeclarativeConfigHealth
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
			return fmt.Errorf("proto: DeclarativeConfigHealth: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeclarativeConfigHealth: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeclarativeConfigHealth
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
				return ErrInvalidLengthDeclarativeConfigHealth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeclarativeConfigHealth
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
					return ErrIntOverflowDeclarativeConfigHealth
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
				return ErrInvalidLengthDeclarativeConfigHealth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeclarativeConfigHealth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeclarativeConfigHealth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= DeclarativeConfigHealth_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeclarativeConfigHealth
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
				return ErrInvalidLengthDeclarativeConfigHealth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeclarativeConfigHealth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeclarativeConfigHealth
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
				return ErrInvalidLengthDeclarativeConfigHealth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeclarativeConfigHealth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceType", wireType)
			}
			m.ResourceType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeclarativeConfigHealth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResourceType |= DeclarativeConfigHealth_ResourceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeclarativeConfigHealth
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
				return ErrInvalidLengthDeclarativeConfigHealth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeclarativeConfigHealth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastTimestamp == nil {
				m.LastTimestamp = &types.Timestamp{}
			}
			if err := m.LastTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeclarativeConfigHealth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeclarativeConfigHealth
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
func skipDeclarativeConfigHealth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeclarativeConfigHealth
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
					return 0, ErrIntOverflowDeclarativeConfigHealth
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
					return 0, ErrIntOverflowDeclarativeConfigHealth
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
				return 0, ErrInvalidLengthDeclarativeConfigHealth
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeclarativeConfigHealth
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeclarativeConfigHealth
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeclarativeConfigHealth        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeclarativeConfigHealth          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeclarativeConfigHealth = fmt.Errorf("proto: unexpected end of group")
)
