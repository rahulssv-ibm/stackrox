// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/collector_runtime_configuration_service.proto

package v1

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

type GetCollectorRuntimeConfigurationResponse struct {
	CollectorRuntimeConfiguration *storage.RuntimeFilteringConfiguration `protobuf:"bytes,1,opt,name=collector_runtime_configuration,json=collectorRuntimeConfiguration,proto3" json:"collector_runtime_configuration,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                               `json:"-"`
	XXX_unrecognized              []byte                                 `json:"-"`
	XXX_sizecache                 int32                                  `json:"-"`
}

func (m *GetCollectorRuntimeConfigurationResponse) Reset() {
	*m = GetCollectorRuntimeConfigurationResponse{}
}
func (m *GetCollectorRuntimeConfigurationResponse) String() string { return proto.CompactTextString(m) }
func (*GetCollectorRuntimeConfigurationResponse) ProtoMessage()    {}
func (*GetCollectorRuntimeConfigurationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c8f554c4c3b1305, []int{0}
}
func (m *GetCollectorRuntimeConfigurationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetCollectorRuntimeConfigurationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetCollectorRuntimeConfigurationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetCollectorRuntimeConfigurationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCollectorRuntimeConfigurationResponse.Merge(m, src)
}
func (m *GetCollectorRuntimeConfigurationResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetCollectorRuntimeConfigurationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCollectorRuntimeConfigurationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCollectorRuntimeConfigurationResponse proto.InternalMessageInfo

func (m *GetCollectorRuntimeConfigurationResponse) GetCollectorRuntimeConfiguration() *storage.RuntimeFilteringConfiguration {
	if m != nil {
		return m.CollectorRuntimeConfiguration
	}
	return nil
}

func (m *GetCollectorRuntimeConfigurationResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GetCollectorRuntimeConfigurationResponse) Clone() *GetCollectorRuntimeConfigurationResponse {
	if m == nil {
		return nil
	}
	cloned := new(GetCollectorRuntimeConfigurationResponse)
	*cloned = *m

	cloned.CollectorRuntimeConfiguration = m.CollectorRuntimeConfiguration.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*GetCollectorRuntimeConfigurationResponse)(nil), "v1.GetCollectorRuntimeConfigurationResponse")
}

func init() {
	proto.RegisterFile("api/v1/collector_runtime_configuration_service.proto", fileDescriptor_9c8f554c4c3b1305)
}

var fileDescriptor_9c8f554c4c3b1305 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3d, 0x4b, 0xf4, 0x40,
	0x10, 0xc7, 0x9f, 0x4d, 0xf1, 0x80, 0x6b, 0xb7, 0xd5, 0x11, 0xbc, 0x78, 0xc6, 0x17, 0x82, 0xca,
	0x86, 0x9c, 0x56, 0x96, 0x1e, 0x6a, 0x2b, 0x67, 0x23, 0x36, 0x61, 0x5d, 0xf6, 0xc2, 0x62, 0x6e,
	0x27, 0xec, 0xee, 0x2d, 0xda, 0x5a, 0xd9, 0xdb, 0x88, 0xf8, 0x81, 0x2c, 0x05, 0xbf, 0x80, 0x44,
	0x3f, 0x88, 0xe4, 0xe5, 0x50, 0x9b, 0xbb, 0x74, 0xc3, 0xfc, 0x67, 0xe6, 0xf7, 0x9f, 0x19, 0x7c,
	0xc8, 0x0a, 0x19, 0xbb, 0x24, 0xe6, 0x90, 0xe7, 0x82, 0x5b, 0xd0, 0xa9, 0x9e, 0x29, 0x2b, 0xa7,
	0x22, 0xe5, 0xa0, 0x26, 0x32, 0x9b, 0x69, 0x66, 0x25, 0xa8, 0xd4, 0x08, 0xed, 0x24, 0x17, 0xb4,
	0xd0, 0x60, 0x81, 0x78, 0x2e, 0xf1, 0x49, 0xdb, 0x29, 0xa6, 0x85, 0xbd, 0x6b, 0xf2, 0xfe, 0x5a,
	0x06, 0x90, 0xe5, 0x22, 0xae, 0x24, 0xa6, 0x14, 0xd8, 0xba, 0xd9, 0xb4, 0x6a, 0xdf, 0x58, 0xd0,
	0x2c, 0x13, 0xf1, 0x1c, 0x31, 0x91, 0xb9, 0x15, 0xba, 0x95, 0xc3, 0x67, 0x84, 0xa3, 0x33, 0x61,
	0x47, 0x73, 0x27, 0xe3, 0xa6, 0x6a, 0xf4, 0xdb, 0xc7, 0x58, 0x98, 0x02, 0x94, 0x11, 0x44, 0xe1,
	0xf5, 0x25, 0x96, 0x7b, 0x68, 0x80, 0xa2, 0xd5, 0xe1, 0x0e, 0x6d, 0xa9, 0xb4, 0x9d, 0x77, 0x5a,
	0x43, 0xa5, 0xca, 0xfe, 0x0e, 0xee, 0xf3, 0x45, 0xdc, 0xe1, 0x8b, 0x87, 0xb7, 0x16, 0x3a, 0xbb,
	0x68, 0x0e, 0x44, 0x1e, 0x10, 0x1e, 0x2c, 0xdb, 0x82, 0xac, 0x50, 0x97, 0xd0, 0x93, 0xea, 0x70,
	0xfe, 0x7e, 0x15, 0x76, 0x5d, 0x3b, 0xdc, 0xbb, 0x7f, 0xff, 0x7a, 0xf4, 0xb6, 0xc9, 0x66, 0x87,
	0x9f, 0x11, 0x89, 0x37, 0xce, 0xc1, 0x74, 0xb7, 0xf2, 0x13, 0x86, 0xb4, 0xe6, 0x44, 0x61, 0x17,
	0xce, 0x11, 0xda, 0x3d, 0xa6, 0xaf, 0x65, 0x80, 0xde, 0xca, 0x00, 0x7d, 0x94, 0x01, 0x7a, 0xfa,
	0x0c, 0xfe, 0xe1, 0x9e, 0x04, 0x6a, 0x2c, 0xe3, 0x37, 0x1a, 0x6e, 0x9b, 0x07, 0x53, 0x56, 0x48,
	0xea, 0x92, 0x2b, 0xcf, 0x25, 0x97, 0xe8, 0xfa, 0x7f, 0x9d, 0x3b, 0xf8, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x69, 0x77, 0x06, 0x74, 0x81, 0x02, 0x00, 0x00,
}

func (m *GetCollectorRuntimeConfigurationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetCollectorRuntimeConfigurationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetCollectorRuntimeConfigurationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CollectorRuntimeConfiguration != nil {
		{
			size, err := m.CollectorRuntimeConfiguration.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCollectorRuntimeConfigurationService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCollectorRuntimeConfigurationService(dAtA []byte, offset int, v uint64) int {
	offset -= sovCollectorRuntimeConfigurationService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetCollectorRuntimeConfigurationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CollectorRuntimeConfiguration != nil {
		l = m.CollectorRuntimeConfiguration.Size()
		n += 1 + l + sovCollectorRuntimeConfigurationService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCollectorRuntimeConfigurationService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCollectorRuntimeConfigurationService(x uint64) (n int) {
	return sovCollectorRuntimeConfigurationService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetCollectorRuntimeConfigurationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollectorRuntimeConfigurationService
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
			return fmt.Errorf("proto: GetCollectorRuntimeConfigurationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCollectorRuntimeConfigurationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectorRuntimeConfiguration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectorRuntimeConfigurationService
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
				return ErrInvalidLengthCollectorRuntimeConfigurationService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollectorRuntimeConfigurationService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CollectorRuntimeConfiguration == nil {
				m.CollectorRuntimeConfiguration = &storage.RuntimeFilteringConfiguration{}
			}
			if err := m.CollectorRuntimeConfiguration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollectorRuntimeConfigurationService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollectorRuntimeConfigurationService
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
func skipCollectorRuntimeConfigurationService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollectorRuntimeConfigurationService
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
					return 0, ErrIntOverflowCollectorRuntimeConfigurationService
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
					return 0, ErrIntOverflowCollectorRuntimeConfigurationService
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
				return 0, ErrInvalidLengthCollectorRuntimeConfigurationService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCollectorRuntimeConfigurationService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCollectorRuntimeConfigurationService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCollectorRuntimeConfigurationService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollectorRuntimeConfigurationService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCollectorRuntimeConfigurationService = fmt.Errorf("proto: unexpected end of group")
)
