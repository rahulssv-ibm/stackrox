// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internalapi/central/network_baseline_sync.proto

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

type NetworkBaselineSync struct {
	NetworkBaselines     []*storage.NetworkBaseline `protobuf:"bytes,1,rep,name=network_baselines,json=networkBaselines,proto3" json:"network_baselines,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *NetworkBaselineSync) Reset()         { *m = NetworkBaselineSync{} }
func (m *NetworkBaselineSync) String() string { return proto.CompactTextString(m) }
func (*NetworkBaselineSync) ProtoMessage()    {}
func (*NetworkBaselineSync) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f52e7f2d7e76ab4, []int{0}
}
func (m *NetworkBaselineSync) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkBaselineSync) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkBaselineSync.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkBaselineSync) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkBaselineSync.Merge(m, src)
}
func (m *NetworkBaselineSync) XXX_Size() int {
	return m.Size()
}
func (m *NetworkBaselineSync) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkBaselineSync.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkBaselineSync proto.InternalMessageInfo

func (m *NetworkBaselineSync) GetNetworkBaselines() []*storage.NetworkBaseline {
	if m != nil {
		return m.NetworkBaselines
	}
	return nil
}

func (m *NetworkBaselineSync) MessageClone() proto.Message {
	return m.Clone()
}
func (m *NetworkBaselineSync) Clone() *NetworkBaselineSync {
	if m == nil {
		return nil
	}
	cloned := new(NetworkBaselineSync)
	*cloned = *m

	if m.NetworkBaselines != nil {
		cloned.NetworkBaselines = make([]*storage.NetworkBaseline, len(m.NetworkBaselines))
		for idx, v := range m.NetworkBaselines {
			cloned.NetworkBaselines[idx] = v.Clone()
		}
	}
	return cloned
}

func init() {
	proto.RegisterType((*NetworkBaselineSync)(nil), "central.NetworkBaselineSync")
}

func init() {
	proto.RegisterFile("internalapi/central/network_baseline_sync.proto", fileDescriptor_7f52e7f2d7e76ab4)
}

var fileDescriptor_7f52e7f2d7e76ab4 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcf, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0x49, 0x2c, 0xc8, 0xd4, 0x4f, 0x4e, 0xcd, 0x2b, 0x29, 0x4a, 0xcc, 0xd1,
	0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0x8e, 0x4f, 0x4a, 0x2c, 0x4e, 0xcd, 0xc9, 0xcc, 0x4b,
	0x8d, 0x2f, 0xae, 0xcc, 0x4b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x2a, 0x92,
	0x92, 0x2b, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0xc5, 0x50, 0x0d, 0x51, 0xa8, 0x14, 0xc3, 0x25,
	0xec, 0x07, 0x91, 0x71, 0x82, 0x4a, 0x04, 0x57, 0xe6, 0x25, 0x0b, 0xb9, 0x72, 0x09, 0xa2, 0x6b,
	0x28, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd0, 0x83, 0x1a, 0xa9, 0x87, 0xa6, 0x31,
	0x48, 0x20, 0x0f, 0x55, 0xa0, 0xd8, 0x49, 0xff, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18,
	0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x56, 0x0f, 0x9b, 0x5f, 0xac, 0xa1,
	0x74, 0x12, 0x1b, 0xd8, 0x55, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x96, 0x32, 0x22,
	0xf1, 0x00, 0x00, 0x00,
}

func (m *NetworkBaselineSync) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkBaselineSync) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkBaselineSync) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.NetworkBaselines) > 0 {
		for iNdEx := len(m.NetworkBaselines) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NetworkBaselines[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNetworkBaselineSync(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintNetworkBaselineSync(dAtA []byte, offset int, v uint64) int {
	offset -= sovNetworkBaselineSync(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NetworkBaselineSync) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.NetworkBaselines) > 0 {
		for _, e := range m.NetworkBaselines {
			l = e.Size()
			n += 1 + l + sovNetworkBaselineSync(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNetworkBaselineSync(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNetworkBaselineSync(x uint64) (n int) {
	return sovNetworkBaselineSync(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NetworkBaselineSync) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkBaselineSync
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
			return fmt.Errorf("proto: NetworkBaselineSync: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkBaselineSync: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkBaselines", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkBaselineSync
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
				return ErrInvalidLengthNetworkBaselineSync
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkBaselineSync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NetworkBaselines = append(m.NetworkBaselines, &storage.NetworkBaseline{})
			if err := m.NetworkBaselines[len(m.NetworkBaselines)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkBaselineSync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNetworkBaselineSync
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
func skipNetworkBaselineSync(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNetworkBaselineSync
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
					return 0, ErrIntOverflowNetworkBaselineSync
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
					return 0, ErrIntOverflowNetworkBaselineSync
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
				return 0, ErrInvalidLengthNetworkBaselineSync
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNetworkBaselineSync
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNetworkBaselineSync
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNetworkBaselineSync        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetworkBaselineSync          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNetworkBaselineSync = fmt.Errorf("proto: unexpected end of group")
)
