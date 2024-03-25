// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internalapi/central/policy_sync.proto

package central

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

type PolicySync struct {
	Policies             []*storage.Policy `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PolicySync) Reset()         { *m = PolicySync{} }
func (m *PolicySync) String() string { return proto.CompactTextString(m) }
func (*PolicySync) ProtoMessage()    {}
func (*PolicySync) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c217656959c2b05, []int{0}
}
func (m *PolicySync) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PolicySync) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PolicySync.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PolicySync) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicySync.Merge(m, src)
}
func (m *PolicySync) XXX_Size() int {
	return m.Size()
}
func (m *PolicySync) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicySync.DiscardUnknown(m)
}

var xxx_messageInfo_PolicySync proto.InternalMessageInfo

func (m *PolicySync) GetPolicies() []*storage.Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

func (m *PolicySync) MessageClone() proto.Message {
	return m.Clone()
}
func (m *PolicySync) Clone() *PolicySync {
	if m == nil {
		return nil
	}
	cloned := new(PolicySync)
	*cloned = *m

	if m.Policies != nil {
		cloned.Policies = make([]*storage.Policy, len(m.Policies))
		for idx, v := range m.Policies {
			cloned.Policies[idx] = v.Clone()
		}
	}
	return cloned
}

func init() {
	proto.RegisterType((*PolicySync)(nil), "central.PolicySync")
}

func init() {
	proto.RegisterFile("internalapi/central/policy_sync.proto", fileDescriptor_4c217656959c2b05)
}

var fileDescriptor_4c217656959c2b05 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0x49, 0x2c, 0xc8, 0xd4, 0x4f, 0x4e, 0xcd, 0x2b, 0x29, 0x4a, 0xcc, 0xd1,
	0x2f, 0xc8, 0xcf, 0xc9, 0x4c, 0xae, 0x8c, 0x2f, 0xae, 0xcc, 0x4b, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x87, 0x4a, 0x49, 0x89, 0x14, 0x97, 0xe4, 0x17, 0x25, 0xa6, 0xa7, 0x42, 0xd5,
	0x40, 0xa4, 0x95, 0x2c, 0xb9, 0xb8, 0x02, 0xc0, 0xfc, 0xe0, 0xca, 0xbc, 0x64, 0x21, 0x6d, 0x2e,
	0x0e, 0xb0, 0x6c, 0x66, 0x6a, 0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0xb7, 0x11, 0xbf, 0x1e, 0x54,
	0x9b, 0x1e, 0x44, 0x59, 0x10, 0x5c, 0x81, 0x93, 0xe4, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9,
	0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x43, 0x14, 0xcc, 0xae, 0x24, 0x36, 0xb0,
	0xe1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0x11, 0x24, 0x11, 0xa4, 0x00, 0x00, 0x00,
}

func (m *PolicySync) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PolicySync) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PolicySync) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Policies) > 0 {
		for iNdEx := len(m.Policies) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Policies[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPolicySync(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPolicySync(dAtA []byte, offset int, v uint64) int {
	offset -= sovPolicySync(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PolicySync) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Policies) > 0 {
		for _, e := range m.Policies {
			l = e.Size()
			n += 1 + l + sovPolicySync(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPolicySync(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPolicySync(x uint64) (n int) {
	return sovPolicySync(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PolicySync) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicySync
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
			return fmt.Errorf("proto: PolicySync: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PolicySync: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicySync
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
				return ErrInvalidLengthPolicySync
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPolicySync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Policies = append(m.Policies, &storage.Policy{})
			if err := m.Policies[len(m.Policies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolicySync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicySync
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
func skipPolicySync(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPolicySync
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
					return 0, ErrIntOverflowPolicySync
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
					return 0, ErrIntOverflowPolicySync
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
				return 0, ErrInvalidLengthPolicySync
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPolicySync
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPolicySync
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPolicySync        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPolicySync          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPolicySync = fmt.Errorf("proto: unexpected end of group")
)
