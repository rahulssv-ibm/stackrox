// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/taints.proto

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

type TaintEffect int32

const (
	TaintEffect_UNKNOWN_TAINT_EFFECT            TaintEffect = 0
	TaintEffect_NO_SCHEDULE_TAINT_EFFECT        TaintEffect = 1
	TaintEffect_PREFER_NO_SCHEDULE_TAINT_EFFECT TaintEffect = 2
	TaintEffect_NO_EXECUTE_TAINT_EFFECT         TaintEffect = 3
)

var TaintEffect_name = map[int32]string{
	0: "UNKNOWN_TAINT_EFFECT",
	1: "NO_SCHEDULE_TAINT_EFFECT",
	2: "PREFER_NO_SCHEDULE_TAINT_EFFECT",
	3: "NO_EXECUTE_TAINT_EFFECT",
}

var TaintEffect_value = map[string]int32{
	"UNKNOWN_TAINT_EFFECT":            0,
	"NO_SCHEDULE_TAINT_EFFECT":        1,
	"PREFER_NO_SCHEDULE_TAINT_EFFECT": 2,
	"NO_EXECUTE_TAINT_EFFECT":         3,
}

func (x TaintEffect) String() string {
	return proto.EnumName(TaintEffect_name, int32(x))
}

func (TaintEffect) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ecc716fc97bf932b, []int{0}
}

type Toleration_Operator int32

const (
	Toleration_TOLERATION_OPERATION_UNKNOWN Toleration_Operator = 0
	Toleration_TOLERATION_OPERATOR_EXISTS   Toleration_Operator = 1
	Toleration_TOLERATION_OPERATOR_EQUAL    Toleration_Operator = 2
)

var Toleration_Operator_name = map[int32]string{
	0: "TOLERATION_OPERATION_UNKNOWN",
	1: "TOLERATION_OPERATOR_EXISTS",
	2: "TOLERATION_OPERATOR_EQUAL",
}

var Toleration_Operator_value = map[string]int32{
	"TOLERATION_OPERATION_UNKNOWN": 0,
	"TOLERATION_OPERATOR_EXISTS":   1,
	"TOLERATION_OPERATOR_EQUAL":    2,
}

func (x Toleration_Operator) String() string {
	return proto.EnumName(Toleration_Operator_name, int32(x))
}

func (Toleration_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ecc716fc97bf932b, []int{1, 0}
}

type Taint struct {
	Key                  string      `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" search:"Taint Key"`                                                              // @gotags: search:"Taint Key"
	Value                string      `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty" search:"Taint Value"`                                                          // @gotags: search:"Taint Value"
	TaintEffect          TaintEffect `protobuf:"varint,3,opt,name=taint_effect,json=taintEffect,proto3,enum=storage.TaintEffect" json:"taint_effect,omitempty" search:"Taint Effect"` // @gotags: search:"Taint Effect"
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Taint) Reset()         { *m = Taint{} }
func (m *Taint) String() string { return proto.CompactTextString(m) }
func (*Taint) ProtoMessage()    {}
func (*Taint) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecc716fc97bf932b, []int{0}
}
func (m *Taint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Taint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Taint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Taint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Taint.Merge(m, src)
}
func (m *Taint) XXX_Size() int {
	return m.Size()
}
func (m *Taint) XXX_DiscardUnknown() {
	xxx_messageInfo_Taint.DiscardUnknown(m)
}

var xxx_messageInfo_Taint proto.InternalMessageInfo

func (m *Taint) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Taint) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Taint) GetTaintEffect() TaintEffect {
	if m != nil {
		return m.TaintEffect
	}
	return TaintEffect_UNKNOWN_TAINT_EFFECT
}

func (m *Taint) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Taint) Clone() *Taint {
	if m == nil {
		return nil
	}
	cloned := new(Taint)
	*cloned = *m

	return cloned
}

type Toleration struct {
	Key                  string              `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" search:"Toleration Key"` // @gotags: search:"Toleration Key"
	Operator             Toleration_Operator `protobuf:"varint,2,opt,name=operator,proto3,enum=storage.Toleration_Operator" json:"operator,omitempty"`
	Value                string              `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty" search:"Toleration Value"` // @gotags: search:"Toleration Value"
	TaintEffect          TaintEffect         `protobuf:"varint,4,opt,name=taint_effect,json=taintEffect,proto3,enum=storage.TaintEffect" json:"taint_effect,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Toleration) Reset()         { *m = Toleration{} }
func (m *Toleration) String() string { return proto.CompactTextString(m) }
func (*Toleration) ProtoMessage()    {}
func (*Toleration) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecc716fc97bf932b, []int{1}
}
func (m *Toleration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Toleration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Toleration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Toleration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Toleration.Merge(m, src)
}
func (m *Toleration) XXX_Size() int {
	return m.Size()
}
func (m *Toleration) XXX_DiscardUnknown() {
	xxx_messageInfo_Toleration.DiscardUnknown(m)
}

var xxx_messageInfo_Toleration proto.InternalMessageInfo

func (m *Toleration) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Toleration) GetOperator() Toleration_Operator {
	if m != nil {
		return m.Operator
	}
	return Toleration_TOLERATION_OPERATION_UNKNOWN
}

func (m *Toleration) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Toleration) GetTaintEffect() TaintEffect {
	if m != nil {
		return m.TaintEffect
	}
	return TaintEffect_UNKNOWN_TAINT_EFFECT
}

func (m *Toleration) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Toleration) Clone() *Toleration {
	if m == nil {
		return nil
	}
	cloned := new(Toleration)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterEnum("storage.TaintEffect", TaintEffect_name, TaintEffect_value)
	proto.RegisterEnum("storage.Toleration_Operator", Toleration_Operator_name, Toleration_Operator_value)
	proto.RegisterType((*Taint)(nil), "storage.Taint")
	proto.RegisterType((*Toleration)(nil), "storage.Toleration")
}

func init() { proto.RegisterFile("storage/taints.proto", fileDescriptor_ecc716fc97bf932b) }

var fileDescriptor_ecc716fc97bf932b = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x2f, 0x49, 0xcc, 0xcc, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x87, 0x8a, 0x2a, 0x65, 0x70, 0xb1, 0x86, 0x80, 0x24, 0x84, 0x04, 0xb8, 0x98, 0xb3,
	0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xb2,
	0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0xb0, 0x18, 0x84, 0x23, 0x64, 0xce, 0xc5, 0x03, 0x36, 0x29,
	0x3e, 0x35, 0x2d, 0x2d, 0x35, 0xb9, 0x44, 0x82, 0x59, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x44, 0x0f,
	0x6a, 0xa0, 0x1e, 0xd8, 0x34, 0x57, 0xb0, 0x5c, 0x10, 0x77, 0x09, 0x82, 0xa3, 0x34, 0x85, 0x89,
	0x8b, 0x2b, 0x24, 0x3f, 0x27, 0xb5, 0x28, 0xb1, 0x24, 0x33, 0x3f, 0x0f, 0x8b, 0x7d, 0x16, 0x5c,
	0x1c, 0xf9, 0x05, 0x20, 0xe9, 0xfc, 0x22, 0xb0, 0x95, 0x7c, 0x46, 0x32, 0x08, 0x53, 0xe1, 0x1a,
	0xf5, 0xfc, 0xa1, 0x6a, 0x82, 0xe0, 0xaa, 0x11, 0x2e, 0x65, 0xc6, 0xe7, 0x52, 0x16, 0x62, 0x5d,
	0x9a, 0xcd, 0xc5, 0x01, 0xb3, 0x44, 0x48, 0x81, 0x4b, 0x26, 0xc4, 0xdf, 0xc7, 0x35, 0xc8, 0x31,
	0xc4, 0xd3, 0xdf, 0x2f, 0xde, 0x3f, 0x00, 0xc6, 0x0a, 0xf5, 0xf3, 0xf6, 0xf3, 0x0f, 0xf7, 0x13,
	0x60, 0x10, 0x92, 0xe3, 0x92, 0xc2, 0x50, 0xe1, 0x1f, 0x14, 0xef, 0x1a, 0xe1, 0x19, 0x1c, 0x12,
	0x2c, 0xc0, 0x28, 0x24, 0xcb, 0x25, 0x89, 0x55, 0x3e, 0x30, 0xd4, 0xd1, 0x47, 0x80, 0x49, 0xab,
	0x9d, 0x91, 0x8b, 0x1b, 0xc9, 0x25, 0x42, 0x12, 0x5c, 0x22, 0x50, 0xb3, 0xe3, 0x43, 0x1c, 0x3d,
	0xfd, 0x42, 0xe2, 0x5d, 0xdd, 0xdc, 0x5c, 0x9d, 0x43, 0x04, 0x18, 0x84, 0x64, 0xb8, 0x24, 0xfc,
	0xfc, 0xe3, 0x83, 0x9d, 0x3d, 0x5c, 0x5d, 0x42, 0x7d, 0x5c, 0x51, 0x65, 0x19, 0x85, 0x94, 0xb9,
	0xe4, 0x03, 0x82, 0x5c, 0xdd, 0x5c, 0x83, 0xe2, 0x71, 0x2a, 0x62, 0x12, 0x92, 0xe6, 0x12, 0xf7,
	0xf3, 0x8f, 0x77, 0x8d, 0x70, 0x75, 0x0e, 0x0d, 0x41, 0x93, 0x64, 0x76, 0xb2, 0x3b, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0xe0, 0x92,
	0xcc, 0xcc, 0xd7, 0x2b, 0x2e, 0x49, 0x4c, 0xce, 0x2e, 0xca, 0xaf, 0x80, 0x24, 0x1c, 0x58, 0xe0,
	0x45, 0x09, 0xea, 0xe9, 0x43, 0x99, 0xd6, 0x50, 0x3a, 0x89, 0x0d, 0xac, 0xc2, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x61, 0xb3, 0x09, 0x2a, 0x72, 0x02, 0x00, 0x00,
}

func (m *Taint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Taint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Taint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TaintEffect != 0 {
		i = encodeVarintTaints(dAtA, i, uint64(m.TaintEffect))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTaints(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTaints(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Toleration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Toleration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Toleration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TaintEffect != 0 {
		i = encodeVarintTaints(dAtA, i, uint64(m.TaintEffect))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTaints(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Operator != 0 {
		i = encodeVarintTaints(dAtA, i, uint64(m.Operator))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTaints(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTaints(dAtA []byte, offset int, v uint64) int {
	offset -= sovTaints(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Taint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTaints(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTaints(uint64(l))
	}
	if m.TaintEffect != 0 {
		n += 1 + sovTaints(uint64(m.TaintEffect))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Toleration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTaints(uint64(l))
	}
	if m.Operator != 0 {
		n += 1 + sovTaints(uint64(m.Operator))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTaints(uint64(l))
	}
	if m.TaintEffect != 0 {
		n += 1 + sovTaints(uint64(m.TaintEffect))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTaints(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTaints(x uint64) (n int) {
	return sovTaints(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Taint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTaints
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
			return fmt.Errorf("proto: Taint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Taint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaints
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
				return ErrInvalidLengthTaints
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTaints
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaints
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
				return ErrInvalidLengthTaints
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTaints
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaintEffect", wireType)
			}
			m.TaintEffect = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaints
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TaintEffect |= TaintEffect(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTaints(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTaints
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
func (m *Toleration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTaints
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
			return fmt.Errorf("proto: Toleration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Toleration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaints
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
				return ErrInvalidLengthTaints
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTaints
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			m.Operator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaints
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Operator |= Toleration_Operator(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaints
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
				return ErrInvalidLengthTaints
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTaints
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaintEffect", wireType)
			}
			m.TaintEffect = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaints
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TaintEffect |= TaintEffect(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTaints(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTaints
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
func skipTaints(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTaints
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
					return 0, ErrIntOverflowTaints
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
					return 0, ErrIntOverflowTaints
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
				return 0, ErrInvalidLengthTaints
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTaints
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTaints
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTaints        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTaints          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTaints = fmt.Errorf("proto: unexpected end of group")
)
