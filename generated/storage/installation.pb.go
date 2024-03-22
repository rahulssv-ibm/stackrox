// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/installation.proto

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

type InstallationInfo struct {
	Id                   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created              *types.Timestamp `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *InstallationInfo) Reset()         { *m = InstallationInfo{} }
func (m *InstallationInfo) String() string { return proto.CompactTextString(m) }
func (*InstallationInfo) ProtoMessage()    {}
func (*InstallationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_77152fd9fb250ba4, []int{0}
}
func (m *InstallationInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InstallationInfo) XXX_MarshalVT(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InstallationInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InstallationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallationInfo.Merge(m, src)
}
func (m *InstallationInfo) XXX_Size() int {
	return m.Size()
}
func (m *InstallationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_InstallationInfo proto.InternalMessageInfo

func (m *InstallationInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InstallationInfo) GetCreated() *types.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *InstallationInfo) MessageClone() proto.Message {
	return m.Clone()
}
func (m *InstallationInfo) Clone() *InstallationInfo {
	if m == nil {
		return nil
	}
	cloned := new(InstallationInfo)
	*cloned = *m

	cloned.Created = m.Created.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*InstallationInfo)(nil), "storage.InstallationInfo")
}

func init() { proto.RegisterFile("storage/installation.proto", fileDescriptor_77152fd9fb250ba4) }

var fileDescriptor_77152fd9fb250ba4 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0xcf, 0xcc, 0x2b, 0x2e, 0x49, 0xcc, 0xc9, 0x49, 0x2c, 0xc9, 0xcc, 0xcf,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0xca, 0x49, 0xc9, 0xa7, 0xe7, 0xe7, 0xa7,
	0xe7, 0xa4, 0xea, 0x83, 0x85, 0x93, 0x4a, 0xd3, 0xf4, 0x4b, 0x32, 0x73, 0x53, 0x8b, 0x4b, 0x12,
	0x73, 0x0b, 0x20, 0x2a, 0x95, 0x22, 0xb8, 0x04, 0x3c, 0x91, 0xf4, 0x7b, 0xe6, 0xa5, 0xe5, 0x0b,
	0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08,
	0x99, 0x70, 0xb1, 0x27, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0xa6, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70,
	0x1b, 0x49, 0xe9, 0x41, 0x8c, 0xd5, 0x83, 0x19, 0xab, 0x17, 0x02, 0x33, 0x36, 0x08, 0xa6, 0xd4,
	0xc9, 0xe4, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1,
	0x58, 0x8e, 0x81, 0x4b, 0x32, 0x33, 0x5f, 0xaf, 0xb8, 0x24, 0x31, 0x39, 0xbb, 0x28, 0xbf, 0x02,
	0xa2, 0x55, 0x0f, 0xea, 0xce, 0x28, 0x98, 0x83, 0x93, 0xd8, 0xc0, 0xe2, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x58, 0x6f, 0x1c, 0xf1, 0xde, 0x00, 0x00, 0x00,
}

func (m *InstallationInfo) MarshalVT() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstallationInfo) MarshalVTTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InstallationInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Created != nil {
		{
			size, err := m.Created.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintInstallation(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintInstallation(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInstallation(dAtA []byte, offset int, v uint64) int {
	offset -= sovInstallation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InstallationInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovInstallation(uint64(l))
	}
	if m.Created != nil {
		l = m.Created.Size()
		n += 1 + l + sovInstallation(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovInstallation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInstallation(x uint64) (n int) {
	return sovInstallation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InstallationInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInstallation
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
			return fmt.Errorf("proto: InstallationInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstallationInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInstallation
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
				return ErrInvalidLengthInstallation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInstallation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInstallation
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
				return ErrInvalidLengthInstallation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInstallation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Created == nil {
				m.Created = &types.Timestamp{}
			}
			if err := m.Created.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInstallation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInstallation
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
func skipInstallation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInstallation
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
					return 0, ErrIntOverflowInstallation
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
					return 0, ErrIntOverflowInstallation
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
				return 0, ErrInvalidLengthInstallation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInstallation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInstallation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInstallation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInstallation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInstallation = fmt.Errorf("proto: unexpected end of group")
)
