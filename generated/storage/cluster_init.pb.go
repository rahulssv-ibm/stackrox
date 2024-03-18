// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/cluster_init.proto

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

type InitBundleMeta struct {
	Id                   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk"` // @gotags: sql:"pk"
	Name                 string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            *types.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy            *User            `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	IsRevoked            bool             `protobuf:"varint,5,opt,name=is_revoked,json=isRevoked,proto3" json:"is_revoked,omitempty"`
	ExpiresAt            *types.Timestamp `protobuf:"bytes,6,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *InitBundleMeta) Reset()         { *m = InitBundleMeta{} }
func (m *InitBundleMeta) String() string { return proto.CompactTextString(m) }
func (*InitBundleMeta) ProtoMessage()    {}
func (*InitBundleMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd67d3248f7aceeb, []int{0}
}
func (m *InitBundleMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InitBundleMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InitBundleMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InitBundleMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitBundleMeta.Merge(m, src)
}
func (m *InitBundleMeta) XXX_Size() int {
	return m.Size()
}
func (m *InitBundleMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_InitBundleMeta.DiscardUnknown(m)
}

var xxx_messageInfo_InitBundleMeta proto.InternalMessageInfo

func (m *InitBundleMeta) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InitBundleMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InitBundleMeta) GetCreatedAt() *types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *InitBundleMeta) GetCreatedBy() *User {
	if m != nil {
		return m.CreatedBy
	}
	return nil
}

func (m *InitBundleMeta) GetIsRevoked() bool {
	if m != nil {
		return m.IsRevoked
	}
	return false
}

func (m *InitBundleMeta) GetExpiresAt() *types.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

func (m *InitBundleMeta) MessageClone() proto.Message {
	return m.Clone()
}
func (m *InitBundleMeta) Clone() *InitBundleMeta {
	if m == nil {
		return nil
	}
	cloned := new(InitBundleMeta)
	*cloned = *m

	cloned.CreatedAt = m.CreatedAt.Clone()
	cloned.CreatedBy = m.CreatedBy.Clone()
	cloned.ExpiresAt = m.ExpiresAt.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*InitBundleMeta)(nil), "storage.InitBundleMeta")
}

func init() { proto.RegisterFile("storage/cluster_init.proto", fileDescriptor_bd67d3248f7aceeb) }

var fileDescriptor_bd67d3248f7aceeb = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4f, 0xb4, 0x30,
	0x1c, 0x87, 0xdf, 0xf2, 0x9e, 0xa7, 0xd4, 0x78, 0x43, 0x27, 0x24, 0x11, 0x89, 0x13, 0x83, 0x29,
	0x89, 0xba, 0x38, 0x1e, 0x9b, 0x83, 0x0b, 0xd1, 0xc5, 0x85, 0x14, 0xf8, 0x4b, 0x9a, 0x03, 0x4a,
	0xda, 0x3f, 0xe6, 0xee, 0x9b, 0xf8, 0x91, 0x1c, 0xfd, 0x08, 0x06, 0x3f, 0x84, 0xab, 0x39, 0xa0,
	0x3a, 0xba, 0x35, 0x4f, 0x9f, 0x5f, 0xf3, 0xa4, 0xd4, 0x37, 0xa8, 0xb4, 0xa8, 0x20, 0x2e, 0xea,
	0xde, 0x20, 0xe8, 0x4c, 0xb6, 0x12, 0x79, 0xa7, 0x15, 0x2a, 0x76, 0x38, 0xdf, 0xf9, 0xe7, 0x95,
	0x52, 0x55, 0x0d, 0xf1, 0x88, 0xf3, 0xfe, 0x39, 0x46, 0xd9, 0x80, 0x41, 0xd1, 0x74, 0x93, 0xe9,
	0x33, 0xfb, 0x4a, 0x6f, 0x40, 0x4f, 0xec, 0xe2, 0x8b, 0xd0, 0xd5, 0x5d, 0x2b, 0x31, 0xe9, 0xdb,
	0xb2, 0x86, 0x7b, 0x40, 0xc1, 0x56, 0xd4, 0x91, 0xa5, 0x47, 0x42, 0x12, 0xb9, 0xa9, 0x23, 0x4b,
	0xc6, 0xe8, 0xa2, 0x15, 0x0d, 0x78, 0xce, 0x48, 0xc6, 0x33, 0xbb, 0xa5, 0xb4, 0xd0, 0x20, 0x10,
	0xca, 0x4c, 0xa0, 0xf7, 0x3f, 0x24, 0xd1, 0xf1, 0x95, 0xcf, 0xa7, 0x00, 0x6e, 0x03, 0xf8, 0x83,
	0x0d, 0x48, 0xdd, 0xd9, 0x5e, 0x23, 0xbb, 0xfc, 0x9d, 0xe6, 0x3b, 0x6f, 0x31, 0x4e, 0x4f, 0xf8,
	0x9c, 0xc6, 0x1f, 0x0d, 0xe8, 0x1f, 0x3b, 0xd9, 0xb1, 0x33, 0x4a, 0xa5, 0xc9, 0x34, 0xbc, 0xa8,
	0x0d, 0x94, 0xde, 0x41, 0x48, 0xa2, 0xa3, 0xd4, 0x95, 0x26, 0x9d, 0xc0, 0xbe, 0x03, 0xb6, 0x9d,
	0xd4, 0x60, 0xf6, 0x1d, 0xcb, 0xbf, 0x3b, 0x66, 0x7b, 0x8d, 0xc9, 0xcd, 0xdb, 0x10, 0x90, 0xf7,
	0x21, 0x20, 0x1f, 0x43, 0x40, 0x5e, 0x3f, 0x83, 0x7f, 0xf4, 0x54, 0x2a, 0x6e, 0x50, 0x14, 0x1b,
	0xad, 0xb6, 0xd3, 0xd8, 0x66, 0x3d, 0xd9, 0x4f, 0xce, 0x97, 0x23, 0xbf, 0xfe, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0xb0, 0x7c, 0xf5, 0x7e, 0x92, 0x01, 0x00, 0x00,
}

func (m *InitBundleMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InitBundleMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InitBundleMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ExpiresAt != nil {
		{
			size, err := m.ExpiresAt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClusterInit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.IsRevoked {
		i--
		if m.IsRevoked {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.CreatedBy != nil {
		{
			size, err := m.CreatedBy.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClusterInit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.CreatedAt != nil {
		{
			size, err := m.CreatedAt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClusterInit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintClusterInit(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintClusterInit(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClusterInit(dAtA []byte, offset int, v uint64) int {
	offset -= sovClusterInit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InitBundleMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovClusterInit(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovClusterInit(uint64(l))
	}
	if m.CreatedAt != nil {
		l = m.CreatedAt.Size()
		n += 1 + l + sovClusterInit(uint64(l))
	}
	if m.CreatedBy != nil {
		l = m.CreatedBy.Size()
		n += 1 + l + sovClusterInit(uint64(l))
	}
	if m.IsRevoked {
		n += 2
	}
	if m.ExpiresAt != nil {
		l = m.ExpiresAt.Size()
		n += 1 + l + sovClusterInit(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovClusterInit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClusterInit(x uint64) (n int) {
	return sovClusterInit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InitBundleMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClusterInit
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
			return fmt.Errorf("proto: InitBundleMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InitBundleMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterInit
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
				return ErrInvalidLengthClusterInit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClusterInit
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
					return ErrIntOverflowClusterInit
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
				return ErrInvalidLengthClusterInit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClusterInit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterInit
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
				return ErrInvalidLengthClusterInit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClusterInit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedAt == nil {
				m.CreatedAt = &types.Timestamp{}
			}
			if err := m.CreatedAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedBy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterInit
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
				return ErrInvalidLengthClusterInit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClusterInit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedBy == nil {
				m.CreatedBy = &User{}
			}
			if err := m.CreatedBy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsRevoked", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterInit
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
			m.IsRevoked = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterInit
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
				return ErrInvalidLengthClusterInit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClusterInit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExpiresAt == nil {
				m.ExpiresAt = &types.Timestamp{}
			}
			if err := m.ExpiresAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClusterInit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClusterInit
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
func skipClusterInit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClusterInit
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
					return 0, ErrIntOverflowClusterInit
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
					return 0, ErrIntOverflowClusterInit
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
				return 0, ErrInvalidLengthClusterInit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClusterInit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClusterInit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClusterInit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClusterInit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClusterInit = fmt.Errorf("proto: unexpected end of group")
)
