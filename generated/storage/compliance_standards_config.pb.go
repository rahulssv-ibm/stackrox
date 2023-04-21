// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/compliance_standards_config.proto

package storage

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/gogo/protobuf/types"
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

type ComplianceConfig struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk"`
	HideScanResults      bool     `protobuf:"varint,2,opt,name=hideScanResults,proto3" json:"hideScanResults,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComplianceConfig) Reset()         { *m = ComplianceConfig{} }
func (m *ComplianceConfig) String() string { return proto.CompactTextString(m) }
func (*ComplianceConfig) ProtoMessage()    {}
func (*ComplianceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3665fd4d49b11ec, []int{0}
}
func (m *ComplianceConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ComplianceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ComplianceConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ComplianceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplianceConfig.Merge(m, src)
}
func (m *ComplianceConfig) XXX_Size() int {
	return m.Size()
}
func (m *ComplianceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplianceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ComplianceConfig proto.InternalMessageInfo

func (m *ComplianceConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ComplianceConfig) GetHideScanResults() bool {
	if m != nil {
		return m.HideScanResults
	}
	return false
}

func (m *ComplianceConfig) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ComplianceConfig) Clone() *ComplianceConfig {
	if m == nil {
		return nil
	}
	cloned := new(ComplianceConfig)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterType((*ComplianceConfig)(nil), "storage.ComplianceConfig")
}

func init() {
	proto.RegisterFile("storage/compliance_standards_config.proto", fileDescriptor_c3665fd4d49b11ec)
}

var fileDescriptor_c3665fd4d49b11ec = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2c, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x4f, 0xce, 0xcf, 0x2d, 0xc8, 0xc9, 0x4c, 0xcc, 0x4b, 0x4e, 0x8d, 0x2f,
	0x2e, 0x49, 0xcc, 0x4b, 0x49, 0x2c, 0x4a, 0x29, 0x8e, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x2a, 0x95, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0x8b, 0xe9, 0x83, 0x58, 0x10, 0x69, 0x29, 0xf9, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x7d, 0x30,
	0x2f, 0xa9, 0x34, 0x4d, 0xbf, 0x24, 0x33, 0x37, 0xb5, 0xb8, 0x24, 0x31, 0xb7, 0x00, 0xaa, 0x40,
	0x18, 0x66, 0x55, 0x66, 0x6e, 0x62, 0x7a, 0x2a, 0x44, 0x50, 0x29, 0x8a, 0x4b, 0xc0, 0x19, 0x6e,
	0xb3, 0x33, 0xd8, 0x3a, 0x21, 0x19, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e,
	0x27, 0x9e, 0x4f, 0xf7, 0xe4, 0x39, 0x8a, 0x0b, 0x73, 0xac, 0x94, 0x0a, 0xb2, 0x95, 0x82, 0x98,
	0x32, 0x53, 0x84, 0x34, 0xb8, 0xf8, 0x33, 0x32, 0x53, 0x52, 0x83, 0x93, 0x13, 0xf3, 0x82, 0x52,
	0x8b, 0x4b, 0x73, 0x4a, 0x8a, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x82, 0xd0, 0x85, 0x9d, 0x4c,
	0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5,
	0x18, 0xb8, 0x24, 0x33, 0xf3, 0xf5, 0x8a, 0x4b, 0x12, 0x93, 0xb3, 0x8b, 0xf2, 0x2b, 0x20, 0x0e,
	0xd0, 0x83, 0x3a, 0x2a, 0x0a, 0xe6, 0xbb, 0x24, 0x36, 0xb0, 0xb8, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xfd, 0x16, 0x50, 0xa6, 0x1a, 0x01, 0x00, 0x00,
}

func (m *ComplianceConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ComplianceConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ComplianceConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.HideScanResults {
		i--
		if m.HideScanResults {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintComplianceStandardsConfig(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintComplianceStandardsConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovComplianceStandardsConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ComplianceConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovComplianceStandardsConfig(uint64(l))
	}
	if m.HideScanResults {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovComplianceStandardsConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozComplianceStandardsConfig(x uint64) (n int) {
	return sovComplianceStandardsConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ComplianceConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComplianceStandardsConfig
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
			return fmt.Errorf("proto: ComplianceConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ComplianceConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceStandardsConfig
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
				return ErrInvalidLengthComplianceStandardsConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceStandardsConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HideScanResults", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceStandardsConfig
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
			m.HideScanResults = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipComplianceStandardsConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComplianceStandardsConfig
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
func skipComplianceStandardsConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowComplianceStandardsConfig
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
					return 0, ErrIntOverflowComplianceStandardsConfig
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
					return 0, ErrIntOverflowComplianceStandardsConfig
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
				return 0, ErrInvalidLengthComplianceStandardsConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupComplianceStandardsConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthComplianceStandardsConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthComplianceStandardsConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowComplianceStandardsConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupComplianceStandardsConfig = fmt.Errorf("proto: unexpected end of group")
)
