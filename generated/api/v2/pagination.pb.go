// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v2/pagination.proto

package v2

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

type Aggregation int32

const (
	Aggregation_UNSET Aggregation = 0
	Aggregation_COUNT Aggregation = 1
	Aggregation_MIN   Aggregation = 2
	Aggregation_MAX   Aggregation = 3
)

var Aggregation_name = map[int32]string{
	0: "UNSET",
	1: "COUNT",
	2: "MIN",
	3: "MAX",
}

var Aggregation_value = map[string]int32{
	"UNSET": 0,
	"COUNT": 1,
	"MIN":   2,
	"MAX":   3,
}

func (x Aggregation) String() string {
	return proto.EnumName(Aggregation_name, int32(x))
}

func (Aggregation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_789c41507a0ed95d, []int{0}
}

type AggregateBy struct {
	AggrFunc             Aggregation `protobuf:"varint,1,opt,name=aggrFunc,proto3,enum=v2.Aggregation" json:"aggrFunc,omitempty"`
	Distinct             bool        `protobuf:"varint,2,opt,name=distinct,proto3" json:"distinct,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AggregateBy) Reset()         { *m = AggregateBy{} }
func (m *AggregateBy) String() string { return proto.CompactTextString(m) }
func (*AggregateBy) ProtoMessage()    {}
func (*AggregateBy) Descriptor() ([]byte, []int) {
	return fileDescriptor_789c41507a0ed95d, []int{0}
}
func (m *AggregateBy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AggregateBy) XXX_MarshalVT(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AggregateBy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AggregateBy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateBy.Merge(m, src)
}
func (m *AggregateBy) XXX_Size() int {
	return m.Size()
}
func (m *AggregateBy) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateBy.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateBy proto.InternalMessageInfo

func (m *AggregateBy) GetAggrFunc() Aggregation {
	if m != nil {
		return m.AggrFunc
	}
	return Aggregation_UNSET
}

func (m *AggregateBy) GetDistinct() bool {
	if m != nil {
		return m.Distinct
	}
	return false
}

func (m *AggregateBy) MessageClone() proto.Message {
	return m.Clone()
}
func (m *AggregateBy) Clone() *AggregateBy {
	if m == nil {
		return nil
	}
	cloned := new(AggregateBy)
	*cloned = *m

	return cloned
}

type SortOption struct {
	Field    string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Reversed bool   `protobuf:"varint,2,opt,name=reversed,proto3" json:"reversed,omitempty"`
	// This field is under development. It is not supported on any REST APIs.
	AggregateBy          *AggregateBy `protobuf:"bytes,3,opt,name=aggregate_by,json=aggregateBy,proto3" json:"aggregate_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SortOption) Reset()         { *m = SortOption{} }
func (m *SortOption) String() string { return proto.CompactTextString(m) }
func (*SortOption) ProtoMessage()    {}
func (*SortOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_789c41507a0ed95d, []int{1}
}
func (m *SortOption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SortOption) XXX_MarshalVT(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SortOption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SortOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SortOption.Merge(m, src)
}
func (m *SortOption) XXX_Size() int {
	return m.Size()
}
func (m *SortOption) XXX_DiscardUnknown() {
	xxx_messageInfo_SortOption.DiscardUnknown(m)
}

var xxx_messageInfo_SortOption proto.InternalMessageInfo

func (m *SortOption) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *SortOption) GetReversed() bool {
	if m != nil {
		return m.Reversed
	}
	return false
}

func (m *SortOption) GetAggregateBy() *AggregateBy {
	if m != nil {
		return m.AggregateBy
	}
	return nil
}

func (m *SortOption) MessageClone() proto.Message {
	return m.Clone()
}
func (m *SortOption) Clone() *SortOption {
	if m == nil {
		return nil
	}
	cloned := new(SortOption)
	*cloned = *m

	cloned.AggregateBy = m.AggregateBy.Clone()
	return cloned
}

type Pagination struct {
	Limit      int32       `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset     int32       `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	SortOption *SortOption `protobuf:"bytes,3,opt,name=sort_option,json=sortOption,proto3" json:"sort_option,omitempty"`
	// This field is under development. It is not supported on any REST APIs.
	SortOptions          []*SortOption `protobuf:"bytes,4,rep,name=sort_options,json=sortOptions,proto3" json:"sort_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_789c41507a0ed95d, []int{2}
}
func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pagination) XXX_MarshalVT(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(m, src)
}
func (m *Pagination) XXX_Size() int {
	return m.Size()
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Pagination) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Pagination) GetSortOption() *SortOption {
	if m != nil {
		return m.SortOption
	}
	return nil
}

func (m *Pagination) GetSortOptions() []*SortOption {
	if m != nil {
		return m.SortOptions
	}
	return nil
}

func (m *Pagination) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Pagination) Clone() *Pagination {
	if m == nil {
		return nil
	}
	cloned := new(Pagination)
	*cloned = *m

	cloned.SortOption = m.SortOption.Clone()
	if m.SortOptions != nil {
		cloned.SortOptions = make([]*SortOption, len(m.SortOptions))
		for idx, v := range m.SortOptions {
			cloned.SortOptions[idx] = v.Clone()
		}
	}
	return cloned
}

func init() {
	proto.RegisterEnum("v2.Aggregation", Aggregation_name, Aggregation_value)
	proto.RegisterType((*AggregateBy)(nil), "v2.AggregateBy")
	proto.RegisterType((*SortOption)(nil), "v2.SortOption")
	proto.RegisterType((*Pagination)(nil), "v2.Pagination")
}

func init() { proto.RegisterFile("api/v2/pagination.proto", fileDescriptor_789c41507a0ed95d) }

var fileDescriptor_789c41507a0ed95d = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x3b, 0xc9, 0x4d, 0x6f, 0x7b, 0x52, 0x7a, 0xc3, 0x70, 0xb9, 0x37, 0xb8, 0x08, 0xa1,
	0xab, 0xa2, 0x90, 0x62, 0xc4, 0x07, 0x68, 0x45, 0xc1, 0x85, 0xad, 0x4c, 0x5b, 0x11, 0x37, 0x65,
	0xda, 0x4e, 0xc3, 0x60, 0xcd, 0x84, 0x99, 0x31, 0xd8, 0x37, 0x71, 0xe5, 0xf3, 0xb8, 0xf4, 0x11,
	0xa4, 0xbe, 0x88, 0x24, 0x69, 0x1a, 0x05, 0x77, 0xe7, 0x0f, 0xe7, 0x3b, 0xff, 0xff, 0x67, 0xe0,
	0x3f, 0x4d, 0x78, 0x2f, 0x0d, 0x7b, 0x09, 0x8d, 0x78, 0x4c, 0x35, 0x17, 0x71, 0x90, 0x48, 0xa1,
	0x05, 0x36, 0xd2, 0xb0, 0x73, 0x03, 0x76, 0x3f, 0x8a, 0x24, 0x8b, 0xa8, 0x66, 0x83, 0x0d, 0x3e,
	0x82, 0x06, 0x8d, 0x22, 0x79, 0xf1, 0x18, 0x2f, 0x5c, 0xe4, 0xa3, 0x6e, 0x3b, 0xfc, 0x13, 0xa4,
	0x61, 0x50, 0xae, 0x70, 0x11, 0x93, 0xfd, 0x02, 0x3e, 0x80, 0xc6, 0x92, 0x2b, 0xcd, 0xe3, 0x85,
	0x76, 0x0d, 0x1f, 0x75, 0x1b, 0x64, 0xaf, 0x3b, 0x12, 0x60, 0x2c, 0xa4, 0x1e, 0x25, 0x19, 0x83,
	0xff, 0x82, 0xb5, 0xe2, 0x6c, 0xbd, 0xcc, 0x6f, 0x36, 0x49, 0x21, 0x32, 0x5e, 0xb2, 0x94, 0x49,
	0xc5, 0x96, 0x25, 0x5f, 0x6a, 0x1c, 0x42, 0x8b, 0x96, 0xb9, 0x66, 0xf3, 0x8d, 0x6b, 0xfa, 0xa8,
	0x6b, 0x7f, 0x0f, 0xc3, 0x06, 0x1b, 0x62, 0xd3, 0x4a, 0x74, 0x5e, 0x10, 0xc0, 0xf5, 0xbe, 0x64,
	0x66, 0xba, 0xe6, 0x0f, 0x5c, 0xe7, 0xa6, 0x16, 0x29, 0x04, 0xfe, 0x07, 0x75, 0xb1, 0x5a, 0x29,
	0x56, 0x44, 0xb6, 0xc8, 0x4e, 0xe1, 0x1e, 0xd8, 0x4a, 0x48, 0x3d, 0x13, 0x79, 0xe2, 0x9d, 0x5f,
	0x3b, 0xf3, 0xab, 0x7a, 0x10, 0x50, 0x55, 0xa7, 0x63, 0x68, 0x7d, 0x01, 0x94, 0xfb, 0xcb, 0x37,
	0x7f, 0x20, 0xec, 0x8a, 0x50, 0x87, 0xa7, 0xd5, 0xcf, 0xce, 0x2e, 0x34, 0xc1, 0x9a, 0x0e, 0xc7,
	0xe7, 0x13, 0xa7, 0x96, 0x8d, 0x67, 0xa3, 0xe9, 0x70, 0xe2, 0x20, 0xfc, 0x1b, 0xcc, 0xab, 0xcb,
	0xa1, 0x63, 0xe4, 0x43, 0xff, 0xd6, 0x31, 0x07, 0xc1, 0xeb, 0xd6, 0x43, 0x6f, 0x5b, 0x0f, 0xbd,
	0x6f, 0x3d, 0xf4, 0xfc, 0xe1, 0xd5, 0xc0, 0xe5, 0x22, 0x50, 0x9a, 0x2e, 0xee, 0xa5, 0x78, 0x2a,
	0x1e, 0x33, 0xa0, 0x09, 0x0f, 0xd2, 0xf0, 0xce, 0x48, 0xc3, 0x79, 0x3d, 0xff, 0x72, 0xf2, 0x19,
	0x00, 0x00, 0xff, 0xff, 0x30, 0x67, 0xe2, 0xcd, 0xf9, 0x01, 0x00, 0x00,
}

func (m *AggregateBy) MarshalVT() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AggregateBy) MarshalVTTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AggregateBy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Distinct {
		i--
		if m.Distinct {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.AggrFunc != 0 {
		i = encodeVarintPagination(dAtA, i, uint64(m.AggrFunc))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SortOption) MarshalVT() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SortOption) MarshalVTTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SortOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.AggregateBy != nil {
		{
			size, err := m.AggregateBy.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPagination(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Reversed {
		i--
		if m.Reversed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Field) > 0 {
		i -= len(m.Field)
		copy(dAtA[i:], m.Field)
		i = encodeVarintPagination(dAtA, i, uint64(len(m.Field)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Pagination) MarshalVT() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pagination) MarshalVTTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pagination) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.SortOptions) > 0 {
		for iNdEx := len(m.SortOptions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SortOptions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPagination(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.SortOption != nil {
		{
			size, err := m.SortOption.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPagination(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Offset != 0 {
		i = encodeVarintPagination(dAtA, i, uint64(m.Offset))
		i--
		dAtA[i] = 0x10
	}
	if m.Limit != 0 {
		i = encodeVarintPagination(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPagination(dAtA []byte, offset int, v uint64) int {
	offset -= sovPagination(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AggregateBy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AggrFunc != 0 {
		n += 1 + sovPagination(uint64(m.AggrFunc))
	}
	if m.Distinct {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SortOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Field)
	if l > 0 {
		n += 1 + l + sovPagination(uint64(l))
	}
	if m.Reversed {
		n += 2
	}
	if m.AggregateBy != nil {
		l = m.AggregateBy.Size()
		n += 1 + l + sovPagination(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Pagination) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Limit != 0 {
		n += 1 + sovPagination(uint64(m.Limit))
	}
	if m.Offset != 0 {
		n += 1 + sovPagination(uint64(m.Offset))
	}
	if m.SortOption != nil {
		l = m.SortOption.Size()
		n += 1 + l + sovPagination(uint64(l))
	}
	if len(m.SortOptions) > 0 {
		for _, e := range m.SortOptions {
			l = e.Size()
			n += 1 + l + sovPagination(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPagination(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPagination(x uint64) (n int) {
	return sovPagination(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AggregateBy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPagination
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
			return fmt.Errorf("proto: AggregateBy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AggregateBy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggrFunc", wireType)
			}
			m.AggrFunc = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AggrFunc |= Aggregation(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Distinct", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
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
			m.Distinct = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPagination(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPagination
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
func (m *SortOption) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPagination
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
			return fmt.Errorf("proto: SortOption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SortOption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
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
				return ErrInvalidLengthPagination
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPagination
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reversed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
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
			m.Reversed = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateBy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
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
				return ErrInvalidLengthPagination
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPagination
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AggregateBy == nil {
				m.AggregateBy = &AggregateBy{}
			}
			if err := m.AggregateBy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPagination(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPagination
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
func (m *Pagination) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPagination
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
			return fmt.Errorf("proto: Pagination: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pagination: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SortOption", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
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
				return ErrInvalidLengthPagination
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPagination
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SortOption == nil {
				m.SortOption = &SortOption{}
			}
			if err := m.SortOption.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SortOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
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
				return ErrInvalidLengthPagination
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPagination
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SortOptions = append(m.SortOptions, &SortOption{})
			if err := m.SortOptions[len(m.SortOptions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPagination(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPagination
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
func skipPagination(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPagination
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
					return 0, ErrIntOverflowPagination
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
					return 0, ErrIntOverflowPagination
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
				return 0, ErrInvalidLengthPagination
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPagination
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPagination
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPagination        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPagination          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPagination = fmt.Errorf("proto: unexpected end of group")
)
