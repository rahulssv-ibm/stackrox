// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/group.proto

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

// Group is a GroupProperties : Role mapping.
type Group struct {
	// GroupProperties define the properties of a group, applying to users when their properties match.
	// They also uniquely identify the group with the props.id field.
	Props *GroupProperties `protobuf:"bytes,1,opt,name=props,proto3" json:"props,omitempty"`
	// This is the name of the role that will apply to users in this group.
	RoleName             string   `protobuf:"bytes,3,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty" search:"Role,hidden" sql:"index=name:groups_unique_indicator;category:unique"` // @gotags: search:"Role,hidden" sql:"index=name:groups_unique_indicator;category:unique"
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_a11bdb88fa982b85, []int{0}
}
func (m *Group) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Group.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return m.Size()
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetProps() *GroupProperties {
	if m != nil {
		return m.Props
	}
	return nil
}

func (m *Group) GetRoleName() string {
	if m != nil {
		return m.RoleName
	}
	return ""
}

func (m *Group) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Group) Clone() *Group {
	if m == nil {
		return nil
	}
	cloned := new(Group)
	*cloned = *m

	cloned.Props = m.Props.Clone()
	return cloned
}

// GroupProperties defines the properties of a group. Groups apply to users when
// their properties match. For instance:
//   - If GroupProperties has only an auth_provider_id, then that group applies
//     to all users logged in with that auth provider.
//   - If GroupProperties in addition has a claim key, then it applies to all
//     users with that auth provider and the claim key, etc.
//
// Note: Changes to GroupProperties may require changes to v1.DeleteGroupRequest.
type GroupProperties struct {
	// Unique identifier for group properties and respectively the group.
	Id                   string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty" sql:"pk"` // @gotags: sql:"pk"
	Traits               *Traits  `protobuf:"bytes,5,opt,name=traits,proto3" json:"traits,omitempty"`
	AuthProviderId       string   `protobuf:"bytes,1,opt,name=auth_provider_id,json=authProviderId,proto3" json:"auth_provider_id,omitempty" search:"Group Auth Provider,hidden" sql:"index=category:unique;name:groups_unique_indicator"` // @gotags: search:"Group Auth Provider,hidden" sql:"index=category:unique;name:groups_unique_indicator"
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty" search:"Group Key,hidden" sql:"index=category:unique;name:groups_unique_indicator"`                                               // @gotags: search:"Group Key,hidden" sql:"index=category:unique;name:groups_unique_indicator"
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty" search:"Group Value,hidden" sql:"index=category:unique;name:groups_unique_indicator"`                                           // @gotags: search:"Group Value,hidden" sql:"index=category:unique;name:groups_unique_indicator"
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupProperties) Reset()         { *m = GroupProperties{} }
func (m *GroupProperties) String() string { return proto.CompactTextString(m) }
func (*GroupProperties) ProtoMessage()    {}
func (*GroupProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_a11bdb88fa982b85, []int{1}
}
func (m *GroupProperties) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupProperties.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupProperties.Merge(m, src)
}
func (m *GroupProperties) XXX_Size() int {
	return m.Size()
}
func (m *GroupProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupProperties.DiscardUnknown(m)
}

var xxx_messageInfo_GroupProperties proto.InternalMessageInfo

func (m *GroupProperties) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GroupProperties) GetTraits() *Traits {
	if m != nil {
		return m.Traits
	}
	return nil
}

func (m *GroupProperties) GetAuthProviderId() string {
	if m != nil {
		return m.AuthProviderId
	}
	return ""
}

func (m *GroupProperties) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GroupProperties) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *GroupProperties) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GroupProperties) Clone() *GroupProperties {
	if m == nil {
		return nil
	}
	cloned := new(GroupProperties)
	*cloned = *m

	cloned.Traits = m.Traits.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*Group)(nil), "storage.Group")
	proto.RegisterType((*GroupProperties)(nil), "storage.GroupProperties")
}

func init() { proto.RegisterFile("storage/group.proto", fileDescriptor_a11bdb88fa982b85) }

var fileDescriptor_a11bdb88fa982b85 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x71, 0x4a, 0x0a, 0x79, 0x48, 0x6d, 0x31, 0x1d, 0x0c, 0x48, 0x51, 0xd5, 0x85, 0x4c,
	0xae, 0x04, 0x23, 0x12, 0x03, 0x0b, 0x62, 0x41, 0x55, 0xd4, 0x89, 0x25, 0x32, 0xc4, 0x2a, 0x56,
	0x5b, 0x6c, 0xbd, 0x38, 0x15, 0xdc, 0x84, 0x81, 0x03, 0x31, 0x72, 0x04, 0x14, 0x2e, 0x82, 0x62,
	0x3b, 0x0c, 0x9d, 0xfc, 0xfc, 0x7f, 0x9f, 0xf5, 0x7e, 0x19, 0x4e, 0x2a, 0xab, 0x51, 0x2c, 0xe5,
	0x6c, 0x89, 0xba, 0x36, 0xdc, 0xa0, 0xb6, 0x9a, 0x1e, 0x84, 0xf0, 0x6c, 0xdc, 0x51, 0x8b, 0x42,
	0xd9, 0xca, 0xe3, 0xe9, 0x02, 0xe2, 0xbb, 0xd6, 0xa6, 0x1c, 0x62, 0x83, 0xda, 0x54, 0x8c, 0x4c,
	0x48, 0x76, 0x74, 0xc9, 0x78, 0xd0, 0xb9, 0xc3, 0x73, 0xd4, 0x46, 0xa2, 0x55, 0xb2, 0xca, 0xbd,
	0x46, 0xcf, 0x21, 0x41, 0xbd, 0x96, 0xc5, 0xab, 0xd8, 0x48, 0xd6, 0x9b, 0x90, 0x2c, 0xc9, 0x0f,
	0xdb, 0xe0, 0x41, 0x6c, 0xe4, 0xf4, 0x93, 0xc0, 0x70, 0xe7, 0x1d, 0x1d, 0x40, 0xa4, 0x4a, 0xb6,
	0xef, 0xcc, 0x48, 0x95, 0xf4, 0x02, 0xfa, 0xbe, 0x09, 0x8b, 0xdd, 0xc6, 0xe1, 0xff, 0xc6, 0x85,
	0x8b, 0xf3, 0x80, 0x69, 0x06, 0x23, 0x51, 0xdb, 0x97, 0xc2, 0xa0, 0xde, 0xaa, 0x52, 0x62, 0xa1,
	0x4a, 0x57, 0x32, 0xc9, 0x07, 0x6d, 0x3e, 0x0f, 0xf1, 0x7d, 0x49, 0x47, 0xd0, 0x5b, 0xc9, 0x77,
	0x16, 0x39, 0xd8, 0x8e, 0x74, 0x0c, 0xf1, 0x56, 0xac, 0xeb, 0xae, 0xa1, 0xbf, 0xdc, 0xde, 0x7c,
	0x35, 0x29, 0xf9, 0x6e, 0x52, 0xf2, 0xd3, 0xa4, 0xe4, 0xe3, 0x37, 0xdd, 0x83, 0x53, 0xa5, 0x79,
	0x65, 0xc5, 0xf3, 0x0a, 0xf5, 0x9b, 0xff, 0x99, 0xae, 0xcd, 0xe3, 0x31, 0x9f, 0x85, 0xf1, 0x3a,
	0x9c, 0x4f, 0x7d, 0x67, 0x5c, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xd8, 0x2a, 0xd1, 0x71,
	0x01, 0x00, 0x00,
}

func (m *Group) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Group) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Group) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.RoleName) > 0 {
		i -= len(m.RoleName)
		copy(dAtA[i:], m.RoleName)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.RoleName)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Props != nil {
		{
			size, err := m.Props.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGroup(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GroupProperties) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupProperties) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupProperties) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Traits != nil {
		{
			size, err := m.Traits.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGroup(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AuthProviderId) > 0 {
		i -= len(m.AuthProviderId)
		copy(dAtA[i:], m.AuthProviderId)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.AuthProviderId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGroup(dAtA []byte, offset int, v uint64) int {
	offset -= sovGroup(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Group) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Props != nil {
		l = m.Props.Size()
		n += 1 + l + sovGroup(uint64(l))
	}
	l = len(m.RoleName)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GroupProperties) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AuthProviderId)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	if m.Traits != nil {
		l = m.Traits.Size()
		n += 1 + l + sovGroup(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGroup(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGroup(x uint64) (n int) {
	return sovGroup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Group) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
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
			return fmt.Errorf("proto: Group: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Group: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Props", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Props == nil {
				m.Props = &GroupProperties{}
			}
			if err := m.Props.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroup
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
func (m *GroupProperties) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
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
			return fmt.Errorf("proto: GroupProperties: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupProperties: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthProviderId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthProviderId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Traits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Traits == nil {
				m.Traits = &Traits{}
			}
			if err := m.Traits.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroup
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
func skipGroup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGroup
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
					return 0, ErrIntOverflowGroup
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
					return 0, ErrIntOverflowGroup
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
				return 0, ErrInvalidLengthGroup
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGroup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGroup
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGroup        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGroup          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGroup = fmt.Errorf("proto: unexpected end of group")
)
