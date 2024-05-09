// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/integration_health.proto

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

type IntegrationHealth_Status int32

const (
	IntegrationHealth_UNINITIALIZED IntegrationHealth_Status = 0
	IntegrationHealth_UNHEALTHY     IntegrationHealth_Status = 1
	IntegrationHealth_HEALTHY       IntegrationHealth_Status = 2
)

var IntegrationHealth_Status_name = map[int32]string{
	0: "UNINITIALIZED",
	1: "UNHEALTHY",
	2: "HEALTHY",
}

var IntegrationHealth_Status_value = map[string]int32{
	"UNINITIALIZED": 0,
	"UNHEALTHY":     1,
	"HEALTHY":       2,
}

func (x IntegrationHealth_Status) String() string {
	return proto.EnumName(IntegrationHealth_Status_name, int32(x))
}

func (IntegrationHealth_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2c35f22f95df38b9, []int{0, 0}
}

type IntegrationHealth_Type int32

const (
	IntegrationHealth_UNKNOWN            IntegrationHealth_Type = 0
	IntegrationHealth_IMAGE_INTEGRATION  IntegrationHealth_Type = 1
	IntegrationHealth_NOTIFIER           IntegrationHealth_Type = 2
	IntegrationHealth_BACKUP             IntegrationHealth_Type = 3
	IntegrationHealth_DECLARATIVE_CONFIG IntegrationHealth_Type = 4
)

var IntegrationHealth_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "IMAGE_INTEGRATION",
	2: "NOTIFIER",
	3: "BACKUP",
	4: "DECLARATIVE_CONFIG",
}

var IntegrationHealth_Type_value = map[string]int32{
	"UNKNOWN":            0,
	"IMAGE_INTEGRATION":  1,
	"NOTIFIER":           2,
	"BACKUP":             3,
	"DECLARATIVE_CONFIG": 4,
}

func (x IntegrationHealth_Type) String() string {
	return proto.EnumName(IntegrationHealth_Type_name, int32(x))
}

func (IntegrationHealth_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2c35f22f95df38b9, []int{0, 1}
}

type IntegrationHealth struct {
	Id           string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk"` // @gotags: sql:"pk"
	Name         string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type         IntegrationHealth_Type   `protobuf:"varint,3,opt,name=type,proto3,enum=storage.IntegrationHealth_Type" json:"type,omitempty"`
	Status       IntegrationHealth_Status `protobuf:"varint,4,opt,name=status,proto3,enum=storage.IntegrationHealth_Status" json:"status,omitempty"`
	ErrorMessage string                   `protobuf:"bytes,5,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	// Timestamp when the status was ascertained
	LastTimestamp        *types.Timestamp `protobuf:"bytes,6,opt,name=last_timestamp,json=lastTimestamp,proto3" json:"last_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IntegrationHealth) Reset()         { *m = IntegrationHealth{} }
func (m *IntegrationHealth) String() string { return proto.CompactTextString(m) }
func (*IntegrationHealth) ProtoMessage()    {}
func (*IntegrationHealth) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c35f22f95df38b9, []int{0}
}
func (m *IntegrationHealth) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IntegrationHealth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IntegrationHealth.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IntegrationHealth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegrationHealth.Merge(m, src)
}
func (m *IntegrationHealth) XXX_Size() int {
	return m.Size()
}
func (m *IntegrationHealth) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegrationHealth.DiscardUnknown(m)
}

var xxx_messageInfo_IntegrationHealth proto.InternalMessageInfo

func (m *IntegrationHealth) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *IntegrationHealth) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IntegrationHealth) GetType() IntegrationHealth_Type {
	if m != nil {
		return m.Type
	}
	return IntegrationHealth_UNKNOWN
}

func (m *IntegrationHealth) GetStatus() IntegrationHealth_Status {
	if m != nil {
		return m.Status
	}
	return IntegrationHealth_UNINITIALIZED
}

func (m *IntegrationHealth) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *IntegrationHealth) GetLastTimestamp() *types.Timestamp {
	if m != nil {
		return m.LastTimestamp
	}
	return nil
}

func (m *IntegrationHealth) MessageClone() proto.Message {
	return m.Clone()
}
func (m *IntegrationHealth) Clone() *IntegrationHealth {
	if m == nil {
		return nil
	}
	cloned := new(IntegrationHealth)
	*cloned = *m

	cloned.LastTimestamp = m.LastTimestamp.Clone()
	return cloned
}

func init() {
	proto.RegisterEnum("storage.IntegrationHealth_Status", IntegrationHealth_Status_name, IntegrationHealth_Status_value)
	proto.RegisterEnum("storage.IntegrationHealth_Type", IntegrationHealth_Type_name, IntegrationHealth_Type_value)
	proto.RegisterType((*IntegrationHealth)(nil), "storage.IntegrationHealth")
}

func init() { proto.RegisterFile("storage/integration_health.proto", fileDescriptor_2c35f22f95df38b9) }

var fileDescriptor_2c35f22f95df38b9 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb3, 0x8e, 0x71, 0xe9, 0xb4, 0x89, 0x9c, 0x91, 0x40, 0xa6, 0x87, 0x34, 0x84, 0x4b,
	0x4e, 0x1b, 0xa9, 0x3d, 0x20, 0x84, 0x84, 0xe4, 0xa6, 0x6e, 0xb2, 0x6a, 0xba, 0x41, 0xc6, 0x01,
	0x51, 0x21, 0x59, 0x5b, 0xba, 0xb8, 0x16, 0x49, 0x36, 0xf2, 0x6e, 0x25, 0xfa, 0x26, 0x3c, 0x01,
	0xcf, 0xc2, 0x91, 0x47, 0x40, 0xe1, 0x45, 0x50, 0xd6, 0x76, 0x39, 0x20, 0xf5, 0xb4, 0x33, 0xbf,
	0xbe, 0x7f, 0x66, 0xe7, 0x87, 0x9e, 0x36, 0xaa, 0x10, 0x99, 0x1c, 0xe6, 0x2b, 0x23, 0xb3, 0x42,
	0x98, 0x5c, 0xad, 0xd2, 0x1b, 0x29, 0x16, 0xe6, 0x86, 0xae, 0x0b, 0x65, 0x14, 0xee, 0x54, 0xc4,
	0xc1, 0x61, 0xa6, 0x54, 0xb6, 0x90, 0x43, 0x2b, 0x5f, 0xdd, 0x7e, 0x19, 0x9a, 0x7c, 0x29, 0xb5,
	0x11, 0xcb, 0x75, 0x49, 0xf6, 0x7f, 0x34, 0xa1, 0xc3, 0xfe, 0x8d, 0x99, 0xd8, 0x29, 0xd8, 0x06,
	0x27, 0xbf, 0x0e, 0x48, 0x8f, 0x0c, 0x76, 0x63, 0x27, 0xbf, 0x46, 0x04, 0x77, 0x25, 0x96, 0x32,
	0x70, 0xac, 0x62, 0x6b, 0x3c, 0x06, 0xd7, 0xdc, 0xad, 0x65, 0xd0, 0xec, 0x91, 0x41, 0xfb, 0xe8,
	0x90, 0x56, 0x2b, 0xe9, 0x7f, 0xd3, 0x68, 0x72, 0xb7, 0x96, 0xb1, 0x85, 0xf1, 0x15, 0x78, 0xda,
	0x08, 0x73, 0xab, 0x03, 0xd7, 0xda, 0x9e, 0x3f, 0x60, 0x7b, 0x67, 0xc1, 0xb8, 0x32, 0xe0, 0x0b,
	0x68, 0xc9, 0xa2, 0x50, 0x45, 0xba, 0x94, 0x5a, 0x8b, 0x4c, 0x06, 0x8f, 0xec, 0x67, 0xf6, 0xad,
	0x78, 0x51, 0x6a, 0x18, 0x42, 0x7b, 0x21, 0xb4, 0x49, 0xef, 0xcf, 0x0c, 0xbc, 0x1e, 0x19, 0xec,
	0x1d, 0x1d, 0xd0, 0x32, 0x08, 0x5a, 0x07, 0x41, 0x93, 0x9a, 0x88, 0x5b, 0x5b, 0xc7, 0x7d, 0xdb,
	0x7f, 0x09, 0x5e, 0xb9, 0x19, 0x3b, 0xd0, 0x9a, 0x73, 0xc6, 0x59, 0xc2, 0xc2, 0x29, 0xbb, 0x8c,
	0x4e, 0xfd, 0x06, 0xb6, 0x60, 0x77, 0xce, 0x27, 0x51, 0x38, 0x4d, 0x26, 0x1f, 0x7d, 0x82, 0x7b,
	0xb0, 0x53, 0x37, 0x4e, 0xff, 0x13, 0xb8, 0xdb, 0x4b, 0xb7, 0xe2, 0x9c, 0x9f, 0xf3, 0xd9, 0x07,
	0xee, 0x37, 0xf0, 0x09, 0x74, 0xd8, 0x45, 0x38, 0x8e, 0x52, 0xc6, 0x93, 0x68, 0x1c, 0x87, 0x09,
	0x9b, 0x71, 0x9f, 0xe0, 0x3e, 0x3c, 0xe6, 0xb3, 0x84, 0x9d, 0xb1, 0x28, 0xf6, 0x1d, 0x04, 0xf0,
	0x4e, 0xc2, 0xd1, 0xf9, 0xfc, 0xad, 0xdf, 0xc4, 0xa7, 0x80, 0xa7, 0xd1, 0x68, 0x1a, 0x6e, 0xd1,
	0xf7, 0x51, 0x3a, 0x9a, 0xf1, 0x33, 0x36, 0xf6, 0xdd, 0x93, 0x37, 0x3f, 0x37, 0x5d, 0xf2, 0x6b,
	0xd3, 0x25, 0xbf, 0x37, 0x5d, 0xf2, 0xfd, 0x4f, 0xb7, 0x01, 0xcf, 0x72, 0x45, 0xb5, 0x11, 0x9f,
	0xbf, 0x16, 0xea, 0x5b, 0x79, 0x57, 0x1d, 0xe6, 0x65, 0x87, 0x0e, 0xab, 0xf2, 0x75, 0xf5, 0x5e,
	0x79, 0x96, 0x38, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x30, 0xcd, 0xaa, 0x54, 0x3d, 0x02, 0x00,
	0x00,
}

func (m *IntegrationHealth) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IntegrationHealth) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IntegrationHealth) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintIntegrationHealth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.ErrorMessage) > 0 {
		i -= len(m.ErrorMessage)
		copy(dAtA[i:], m.ErrorMessage)
		i = encodeVarintIntegrationHealth(dAtA, i, uint64(len(m.ErrorMessage)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Status != 0 {
		i = encodeVarintIntegrationHealth(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.Type != 0 {
		i = encodeVarintIntegrationHealth(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintIntegrationHealth(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintIntegrationHealth(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIntegrationHealth(dAtA []byte, offset int, v uint64) int {
	offset -= sovIntegrationHealth(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IntegrationHealth) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovIntegrationHealth(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovIntegrationHealth(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovIntegrationHealth(uint64(m.Type))
	}
	if m.Status != 0 {
		n += 1 + sovIntegrationHealth(uint64(m.Status))
	}
	l = len(m.ErrorMessage)
	if l > 0 {
		n += 1 + l + sovIntegrationHealth(uint64(l))
	}
	if m.LastTimestamp != nil {
		l = m.LastTimestamp.Size()
		n += 1 + l + sovIntegrationHealth(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIntegrationHealth(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIntegrationHealth(x uint64) (n int) {
	return sovIntegrationHealth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IntegrationHealth) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIntegrationHealth
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
			return fmt.Errorf("proto: IntegrationHealth: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IntegrationHealth: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntegrationHealth
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
				return ErrInvalidLengthIntegrationHealth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIntegrationHealth
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
					return ErrIntOverflowIntegrationHealth
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
				return ErrInvalidLengthIntegrationHealth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIntegrationHealth
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
					return ErrIntOverflowIntegrationHealth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= IntegrationHealth_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntegrationHealth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= IntegrationHealth_Status(b&0x7F) << shift
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
					return ErrIntOverflowIntegrationHealth
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
				return ErrInvalidLengthIntegrationHealth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIntegrationHealth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntegrationHealth
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
				return ErrInvalidLengthIntegrationHealth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIntegrationHealth
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
			skippy, err := skipIntegrationHealth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIntegrationHealth
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
func skipIntegrationHealth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIntegrationHealth
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
					return 0, ErrIntOverflowIntegrationHealth
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
					return 0, ErrIntOverflowIntegrationHealth
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
				return 0, ErrInvalidLengthIntegrationHealth
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIntegrationHealth
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIntegrationHealth
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIntegrationHealth        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIntegrationHealth          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIntegrationHealth = fmt.Errorf("proto: unexpected end of group")
)
