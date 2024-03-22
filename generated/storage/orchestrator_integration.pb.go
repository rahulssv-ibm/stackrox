// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/orchestrator_integration.proto

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

// Next Tag: 5
type OrchestratorIntegration struct {
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Types that are valid to be assigned to IntegrationConfig:
	//	*OrchestratorIntegration_Clairify
	IntegrationConfig    isOrchestratorIntegration_IntegrationConfig `protobuf_oneof:"IntegrationConfig"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *OrchestratorIntegration) Reset()         { *m = OrchestratorIntegration{} }
func (m *OrchestratorIntegration) String() string { return proto.CompactTextString(m) }
func (*OrchestratorIntegration) ProtoMessage()    {}
func (*OrchestratorIntegration) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ec4a06d39fb9cc0, []int{0}
}
func (m *OrchestratorIntegration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrchestratorIntegration) XXX_MarshalVT(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrchestratorIntegration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrchestratorIntegration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrchestratorIntegration.Merge(m, src)
}
func (m *OrchestratorIntegration) XXX_Size() int {
	return m.Size()
}
func (m *OrchestratorIntegration) XXX_DiscardUnknown() {
	xxx_messageInfo_OrchestratorIntegration.DiscardUnknown(m)
}

var xxx_messageInfo_OrchestratorIntegration proto.InternalMessageInfo

type isOrchestratorIntegration_IntegrationConfig interface {
	isOrchestratorIntegration_IntegrationConfig()
	MarshalVTTo([]byte) (int, error)
	Size() int
	Clone() isOrchestratorIntegration_IntegrationConfig
}

type OrchestratorIntegration_Clairify struct {
	Clairify *ClairifyConfig `protobuf:"bytes,4,opt,name=clairify,proto3,oneof" json:"clairify,omitempty"`
}

func (*OrchestratorIntegration_Clairify) isOrchestratorIntegration_IntegrationConfig() {}
func (m *OrchestratorIntegration_Clairify) Clone() isOrchestratorIntegration_IntegrationConfig {
	if m == nil {
		return nil
	}
	cloned := new(OrchestratorIntegration_Clairify)
	*cloned = *m

	cloned.Clairify = m.Clairify.Clone()
	return cloned
}

func (m *OrchestratorIntegration) GetIntegrationConfig() isOrchestratorIntegration_IntegrationConfig {
	if m != nil {
		return m.IntegrationConfig
	}
	return nil
}

func (m *OrchestratorIntegration) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrchestratorIntegration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrchestratorIntegration) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *OrchestratorIntegration) GetClairify() *ClairifyConfig {
	if x, ok := m.GetIntegrationConfig().(*OrchestratorIntegration_Clairify); ok {
		return x.Clairify
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OrchestratorIntegration) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OrchestratorIntegration_Clairify)(nil),
	}
}

func (m *OrchestratorIntegration) MessageClone() proto.Message {
	return m.Clone()
}
func (m *OrchestratorIntegration) Clone() *OrchestratorIntegration {
	if m == nil {
		return nil
	}
	cloned := new(OrchestratorIntegration)
	*cloned = *m

	if m.IntegrationConfig != nil {
		cloned.IntegrationConfig = m.IntegrationConfig.Clone()
	}
	return cloned
}

func init() {
	proto.RegisterType((*OrchestratorIntegration)(nil), "storage.OrchestratorIntegration")
}

func init() {
	proto.RegisterFile("storage/orchestrator_integration.proto", fileDescriptor_7ec4a06d39fb9cc0)
}

var fileDescriptor_7ec4a06d39fb9cc0 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0xcf, 0x2f, 0x4a, 0xce, 0x48, 0x2d, 0x2e, 0x29, 0x4a, 0x2c, 0xc9, 0x2f,
	0x8a, 0xcf, 0xcc, 0x2b, 0x49, 0x4d, 0x2f, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x87, 0xaa, 0x93, 0x92, 0x87, 0x69, 0xc8, 0xcc, 0x4d, 0x4c, 0x4f, 0xc5,
	0x54, 0xa9, 0x34, 0x97, 0x91, 0x4b, 0xdc, 0x1f, 0xc9, 0x30, 0x4f, 0x84, 0x0a, 0x21, 0x3e, 0x2e,
	0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x21, 0x2e,
	0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0xb0, 0x08, 0x98, 0x0d, 0x12, 0x2b, 0xa9, 0x2c, 0x48,
	0x95, 0x60, 0x86, 0x88, 0x81, 0xd8, 0x42, 0xa6, 0x5c, 0x1c, 0xc9, 0x39, 0x89, 0x99, 0x45, 0x99,
	0x69, 0x95, 0x12, 0x2c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xe2, 0x7a, 0x50, 0x77, 0xe8, 0x39, 0x43,
	0x25, 0x9c, 0xf3, 0xf3, 0xd2, 0x32, 0xd3, 0x3d, 0x18, 0x82, 0xe0, 0x4a, 0x9d, 0x84, 0xb9, 0x04,
	0x91, 0x6c, 0x87, 0x28, 0x70, 0x32, 0x39, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0xe0, 0x92, 0xcc, 0xcc, 0xd7, 0x2b, 0x2e, 0x49, 0x4c,
	0xce, 0x2e, 0xca, 0xaf, 0x80, 0x78, 0x02, 0x66, 0x78, 0x14, 0xcc, 0xdb, 0x49, 0x6c, 0x60, 0x71,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0x18, 0x9d, 0xc6, 0x30, 0x01, 0x00, 0x00,
}

func (m *OrchestratorIntegration) MarshalVT() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrchestratorIntegration) MarshalVTTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrchestratorIntegration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.IntegrationConfig != nil {
		{
			size := m.IntegrationConfig.Size()
			i -= size
			if _, err := m.IntegrationConfig.MarshalVTTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintOrchestratorIntegration(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintOrchestratorIntegration(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintOrchestratorIntegration(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OrchestratorIntegration_Clairify) MarshalVTTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrchestratorIntegration_Clairify) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Clairify != nil {
		{
			size, err := m.Clairify.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOrchestratorIntegration(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func encodeVarintOrchestratorIntegration(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrchestratorIntegration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrchestratorIntegration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovOrchestratorIntegration(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovOrchestratorIntegration(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovOrchestratorIntegration(uint64(l))
	}
	if m.IntegrationConfig != nil {
		n += m.IntegrationConfig.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *OrchestratorIntegration_Clairify) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Clairify != nil {
		l = m.Clairify.Size()
		n += 1 + l + sovOrchestratorIntegration(uint64(l))
	}
	return n
}

func sovOrchestratorIntegration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrchestratorIntegration(x uint64) (n int) {
	return sovOrchestratorIntegration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrchestratorIntegration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrchestratorIntegration
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
			return fmt.Errorf("proto: OrchestratorIntegration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrchestratorIntegration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrchestratorIntegration
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
				return ErrInvalidLengthOrchestratorIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrchestratorIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrchestratorIntegration
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
				return ErrInvalidLengthOrchestratorIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrchestratorIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrchestratorIntegration
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
				return ErrInvalidLengthOrchestratorIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrchestratorIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clairify", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrchestratorIntegration
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
				return ErrInvalidLengthOrchestratorIntegration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrchestratorIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ClairifyConfig{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.IntegrationConfig = &OrchestratorIntegration_Clairify{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrchestratorIntegration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrchestratorIntegration
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
func skipOrchestratorIntegration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrchestratorIntegration
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
					return 0, ErrIntOverflowOrchestratorIntegration
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
					return 0, ErrIntOverflowOrchestratorIntegration
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
				return 0, ErrInvalidLengthOrchestratorIntegration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrchestratorIntegration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrchestratorIntegration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrchestratorIntegration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrchestratorIntegration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrchestratorIntegration = fmt.Errorf("proto: unexpected end of group")
)
