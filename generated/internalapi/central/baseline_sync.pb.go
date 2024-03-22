// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internalapi/central/baseline_sync.proto

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

type BaselineSync struct {
	Baselines            []*storage.ProcessBaseline `protobuf:"bytes,1,rep,name=baselines,proto3" json:"baselines,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *BaselineSync) Reset()         { *m = BaselineSync{} }
func (m *BaselineSync) String() string { return proto.CompactTextString(m) }
func (*BaselineSync) ProtoMessage()    {}
func (*BaselineSync) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb100b68bc4abe1f, []int{0}
}
func (m *BaselineSync) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaselineSync) XXX_MarshalVT(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaselineSync.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaselineSync) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaselineSync.Merge(m, src)
}
func (m *BaselineSync) XXX_Size() int {
	return m.Size()
}
func (m *BaselineSync) XXX_DiscardUnknown() {
	xxx_messageInfo_BaselineSync.DiscardUnknown(m)
}

var xxx_messageInfo_BaselineSync proto.InternalMessageInfo

func (m *BaselineSync) GetBaselines() []*storage.ProcessBaseline {
	if m != nil {
		return m.Baselines
	}
	return nil
}

func (m *BaselineSync) MessageClone() proto.Message {
	return m.Clone()
}
func (m *BaselineSync) Clone() *BaselineSync {
	if m == nil {
		return nil
	}
	cloned := new(BaselineSync)
	*cloned = *m

	if m.Baselines != nil {
		cloned.Baselines = make([]*storage.ProcessBaseline, len(m.Baselines))
		for idx, v := range m.Baselines {
			cloned.Baselines[idx] = v.Clone()
		}
	}
	return cloned
}

func init() {
	proto.RegisterType((*BaselineSync)(nil), "central.BaselineSync")
}

func init() {
	proto.RegisterFile("internalapi/central/baseline_sync.proto", fileDescriptor_fb100b68bc4abe1f)
}

var fileDescriptor_fb100b68bc4abe1f = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcf, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0x49, 0x2c, 0xc8, 0xd4, 0x4f, 0x4e, 0xcd, 0x2b, 0x29, 0x4a, 0xcc, 0xd1,
	0x4f, 0x4a, 0x2c, 0x4e, 0xcd, 0xc9, 0xcc, 0x4b, 0x8d, 0x2f, 0xae, 0xcc, 0x4b, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x4a, 0x4a, 0xc9, 0x15, 0x97, 0xe4, 0x17, 0x25, 0xa6, 0xa7,
	0xea, 0x17, 0x14, 0xe5, 0x27, 0xa7, 0x16, 0x17, 0xc7, 0xc3, 0x54, 0x43, 0x14, 0x2a, 0xb9, 0x71,
	0xf1, 0x38, 0x41, 0x45, 0x82, 0x2b, 0xf3, 0x92, 0x85, 0xcc, 0xb8, 0x38, 0x61, 0x2a, 0x8a, 0x25,
	0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x24, 0xf4, 0xa0, 0x66, 0xe8, 0x05, 0x40, 0xcc, 0x80, 0x69,
	0x08, 0x42, 0x28, 0x75, 0x92, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0x88, 0x82, 0x39, 0x21, 0x89, 0x0d, 0x6c, 0x93, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x28, 0xaf, 0xd0, 0x66, 0xbd, 0x00, 0x00, 0x00,
}

func (m *BaselineSync) MarshalVT() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaselineSync) MarshalVTTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaselineSync) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Baselines) > 0 {
		for iNdEx := len(m.Baselines) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Baselines[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBaselineSync(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintBaselineSync(dAtA []byte, offset int, v uint64) int {
	offset -= sovBaselineSync(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaselineSync) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Baselines) > 0 {
		for _, e := range m.Baselines {
			l = e.Size()
			n += 1 + l + sovBaselineSync(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBaselineSync(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBaselineSync(x uint64) (n int) {
	return sovBaselineSync(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaselineSync) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBaselineSync
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
			return fmt.Errorf("proto: BaselineSync: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaselineSync: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Baselines", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaselineSync
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
				return ErrInvalidLengthBaselineSync
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBaselineSync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Baselines = append(m.Baselines, &storage.ProcessBaseline{})
			if err := m.Baselines[len(m.Baselines)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBaselineSync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBaselineSync
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
func skipBaselineSync(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBaselineSync
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
					return 0, ErrIntOverflowBaselineSync
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
					return 0, ErrIntOverflowBaselineSync
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
				return 0, ErrInvalidLengthBaselineSync
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBaselineSync
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBaselineSync
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBaselineSync        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBaselineSync          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBaselineSync = fmt.Errorf("proto: unexpected end of group")
)
