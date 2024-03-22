// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internalapi/wrapper/splunk_alert.proto

package wrapper

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

// Splunk notification needs the source of data
// and the type of data.
type SplunkEvent struct {
	Event                *types.Any `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Source               string     `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Sourcetype           string     `protobuf:"bytes,3,opt,name=sourcetype,proto3" json:"sourcetype,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SplunkEvent) Reset()         { *m = SplunkEvent{} }
func (m *SplunkEvent) String() string { return proto.CompactTextString(m) }
func (*SplunkEvent) ProtoMessage()    {}
func (*SplunkEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f414752d1f1082, []int{0}
}
func (m *SplunkEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SplunkEvent) XXX_MarshalVT(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SplunkEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SplunkEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SplunkEvent.Merge(m, src)
}
func (m *SplunkEvent) XXX_Size() int {
	return m.Size()
}
func (m *SplunkEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SplunkEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SplunkEvent proto.InternalMessageInfo

func (m *SplunkEvent) GetEvent() *types.Any {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *SplunkEvent) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *SplunkEvent) GetSourcetype() string {
	if m != nil {
		return m.Sourcetype
	}
	return ""
}

func (m *SplunkEvent) MessageClone() proto.Message {
	return m.Clone()
}
func (m *SplunkEvent) Clone() *SplunkEvent {
	if m == nil {
		return nil
	}
	cloned := new(SplunkEvent)
	*cloned = *m

	cloned.Event = m.Event.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*SplunkEvent)(nil), "wrapper.SplunkEvent")
}

func init() {
	proto.RegisterFile("internalapi/wrapper/splunk_alert.proto", fileDescriptor_02f414752d1f1082)
}

var fileDescriptor_02f414752d1f1082 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcb, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0x49, 0x2c, 0xc8, 0xd4, 0x2f, 0x2f, 0x4a, 0x2c, 0x28, 0x48, 0x2d, 0xd2,
	0x2f, 0x2e, 0xc8, 0x29, 0xcd, 0xcb, 0x8e, 0x4f, 0xcc, 0x49, 0x2d, 0x2a, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x87, 0xca, 0x49, 0x49, 0xa6, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83,
	0x85, 0x93, 0x4a, 0xd3, 0xf4, 0x13, 0xf3, 0x2a, 0x21, 0x6a, 0x94, 0x0a, 0xb9, 0xb8, 0x83, 0xc1,
	0x3a, 0x5d, 0xcb, 0x52, 0xf3, 0x4a, 0x84, 0xb4, 0xb8, 0x58, 0x53, 0x41, 0x0c, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x6e, 0x23, 0x11, 0x3d, 0x88, 0x4e, 0x3d, 0x98, 0x4e, 0x3d, 0xc7, 0xbc, 0xca, 0x20,
	0x88, 0x12, 0x21, 0x31, 0x2e, 0xb6, 0xe2, 0xfc, 0xd2, 0xa2, 0xe4, 0x54, 0x09, 0x26, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x28, 0x4f, 0x48, 0x8e, 0x8b, 0x0b, 0xc2, 0x2a, 0xa9, 0x2c, 0x48, 0x95, 0x60,
	0x06, 0xcb, 0x21, 0x89, 0x38, 0x49, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83,
	0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x44, 0xc1, 0x1c, 0x9a, 0xc4, 0x06, 0xb6, 0xc7, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x59, 0x84, 0xac, 0xfa, 0xe2, 0x00, 0x00, 0x00,
}

func (m *SplunkEvent) MarshalVT() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SplunkEvent) MarshalVTTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SplunkEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Sourcetype) > 0 {
		i -= len(m.Sourcetype)
		copy(dAtA[i:], m.Sourcetype)
		i = encodeVarintSplunkAlert(dAtA, i, uint64(len(m.Sourcetype)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Source) > 0 {
		i -= len(m.Source)
		copy(dAtA[i:], m.Source)
		i = encodeVarintSplunkAlert(dAtA, i, uint64(len(m.Source)))
		i--
		dAtA[i] = 0x12
	}
	if m.Event != nil {
		{
			size, err := m.Event.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSplunkAlert(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSplunkAlert(dAtA []byte, offset int, v uint64) int {
	offset -= sovSplunkAlert(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SplunkEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Event != nil {
		l = m.Event.Size()
		n += 1 + l + sovSplunkAlert(uint64(l))
	}
	l = len(m.Source)
	if l > 0 {
		n += 1 + l + sovSplunkAlert(uint64(l))
	}
	l = len(m.Sourcetype)
	if l > 0 {
		n += 1 + l + sovSplunkAlert(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSplunkAlert(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSplunkAlert(x uint64) (n int) {
	return sovSplunkAlert(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SplunkEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSplunkAlert
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
			return fmt.Errorf("proto: SplunkEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SplunkEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSplunkAlert
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
				return ErrInvalidLengthSplunkAlert
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSplunkAlert
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Event == nil {
				m.Event = &types.Any{}
			}
			if err := m.Event.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSplunkAlert
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
				return ErrInvalidLengthSplunkAlert
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSplunkAlert
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sourcetype", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSplunkAlert
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
				return ErrInvalidLengthSplunkAlert
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSplunkAlert
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sourcetype = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSplunkAlert(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSplunkAlert
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
func skipSplunkAlert(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSplunkAlert
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
					return 0, ErrIntOverflowSplunkAlert
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
					return 0, ErrIntOverflowSplunkAlert
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
				return 0, ErrInvalidLengthSplunkAlert
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSplunkAlert
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSplunkAlert
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSplunkAlert        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSplunkAlert          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSplunkAlert = fmt.Errorf("proto: unexpected end of group")
)
