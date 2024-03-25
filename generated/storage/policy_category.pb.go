// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/policy_category.proto

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

type PolicyCategory struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk"`     // @gotags: sql:"pk"
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" search:"Policy Category,store,hidden" sql:"unique"` // @gotags: search:"Policy Category,store,hidden" sql:"unique"
	IsDefault            bool     `protobuf:"varint,3,opt,name=isDefault,proto3" json:"isDefault,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyCategory) Reset()         { *m = PolicyCategory{} }
func (m *PolicyCategory) String() string { return proto.CompactTextString(m) }
func (*PolicyCategory) ProtoMessage()    {}
func (*PolicyCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b38511e664e6e6a, []int{0}
}
func (m *PolicyCategory) XXX_UnmarshalVT(b []byte) error {
	return m.UnmarshalVT(b)
}
func (m *PolicyCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PolicyCategory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PolicyCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyCategory.Merge(m, src)
}
func (m *PolicyCategory) XXX_Size() int {
	return m.Size()
}
func (m *PolicyCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyCategory.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyCategory proto.InternalMessageInfo

func (m *PolicyCategory) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PolicyCategory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PolicyCategory) GetIsDefault() bool {
	if m != nil {
		return m.IsDefault
	}
	return false
}

func (m *PolicyCategory) MessageClone() proto.Message {
	return m.Clone()
}
func (m *PolicyCategory) Clone() *PolicyCategory {
	if m == nil {
		return nil
	}
	cloned := new(PolicyCategory)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterType((*PolicyCategory)(nil), "storage.PolicyCategory")
}

func init() { proto.RegisterFile("storage/policy_category.proto", fileDescriptor_0b38511e664e6e6a) }

var fileDescriptor_0b38511e664e6e6a = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x2f, 0xc8, 0xcf, 0xc9, 0x4c, 0xae, 0x8c, 0x4f, 0x4e, 0x2c, 0x49, 0x4d,
	0xcf, 0x2f, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x4a, 0x2b, 0x05, 0x71,
	0xf1, 0x05, 0x80, 0x55, 0x38, 0x43, 0x15, 0x08, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a,
	0x30, 0x81, 0x45, 0xc0, 0x6c, 0x21, 0x19, 0x2e, 0xce, 0xcc, 0x62, 0x97, 0xd4, 0xb4, 0xc4, 0xd2,
	0x9c, 0x12, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x84, 0x80, 0x93, 0xf9, 0x89, 0x47, 0x72,
	0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x03, 0x97, 0x64,
	0x66, 0xbe, 0x5e, 0x71, 0x49, 0x62, 0x72, 0x76, 0x51, 0x7e, 0x05, 0xc4, 0x05, 0x7a, 0x50, 0x07,
	0x44, 0xc1, 0x5c, 0xf2, 0x83, 0x91, 0x31, 0x89, 0x0d, 0x2c, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x3c, 0xa3, 0xcc, 0xdb, 0xbd, 0x00, 0x00, 0x00,
}

func (m *PolicyCategory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PolicyCategory) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PolicyCategory) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.IsDefault {
		i--
		if m.IsDefault {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPolicyCategory(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintPolicyCategory(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPolicyCategory(dAtA []byte, offset int, v uint64) int {
	offset -= sovPolicyCategory(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PolicyCategory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPolicyCategory(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPolicyCategory(uint64(l))
	}
	if m.IsDefault {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPolicyCategory(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPolicyCategory(x uint64) (n int) {
	return sovPolicyCategory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PolicyCategory) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicyCategory
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
			return fmt.Errorf("proto: PolicyCategory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PolicyCategory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicyCategory
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
				return ErrInvalidLengthPolicyCategory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicyCategory
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
					return ErrIntOverflowPolicyCategory
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
				return ErrInvalidLengthPolicyCategory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicyCategory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsDefault", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicyCategory
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
			m.IsDefault = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPolicyCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicyCategory
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
func skipPolicyCategory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPolicyCategory
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
					return 0, ErrIntOverflowPolicyCategory
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
					return 0, ErrIntOverflowPolicyCategory
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
				return 0, ErrInvalidLengthPolicyCategory
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPolicyCategory
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPolicyCategory
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPolicyCategory        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPolicyCategory          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPolicyCategory = fmt.Errorf("proto: unexpected end of group")
)
