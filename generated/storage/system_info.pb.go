// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/system_info.proto

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

type BackupInfo struct {
	BackupLastRunAt      *types.Timestamp `protobuf:"bytes,1,opt,name=backup_last_run_at,json=backupLastRunAt,proto3" json:"backup_last_run_at,omitempty"`
	Status               OperationStatus  `protobuf:"varint,2,opt,name=status,proto3,enum=storage.OperationStatus" json:"status,omitempty"`
	Requestor            *SlimUser        `protobuf:"bytes,3,opt,name=requestor,proto3" json:"requestor,omitempty" sql:"ignore_labels(User ID)"` // @gotags: sql:"ignore_labels(User ID)"
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BackupInfo) Reset()         { *m = BackupInfo{} }
func (m *BackupInfo) String() string { return proto.CompactTextString(m) }
func (*BackupInfo) ProtoMessage()    {}
func (*BackupInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_aea16395861e483e, []int{0}
}
func (m *BackupInfo) XXX_UnmarshalVT(b []byte) error {
	return m.UnmarshalVT(b)
}
func (m *BackupInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BackupInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BackupInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupInfo.Merge(m, src)
}
func (m *BackupInfo) XXX_Size() int {
	return m.Size()
}
func (m *BackupInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BackupInfo proto.InternalMessageInfo

func (m *BackupInfo) GetBackupLastRunAt() *types.Timestamp {
	if m != nil {
		return m.BackupLastRunAt
	}
	return nil
}

func (m *BackupInfo) GetStatus() OperationStatus {
	if m != nil {
		return m.Status
	}
	return OperationStatus_FAIL
}

func (m *BackupInfo) GetRequestor() *SlimUser {
	if m != nil {
		return m.Requestor
	}
	return nil
}

func (m *BackupInfo) MessageClone() proto.Message {
	return m.Clone()
}
func (m *BackupInfo) Clone() *BackupInfo {
	if m == nil {
		return nil
	}
	cloned := new(BackupInfo)
	*cloned = *m

	cloned.BackupLastRunAt = m.BackupLastRunAt.Clone()
	cloned.Requestor = m.Requestor.Clone()
	return cloned
}

type SystemInfo struct {
	BackupInfo           *BackupInfo `protobuf:"bytes,1,opt,name=backup_info,json=backupInfo,proto3" json:"backup_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SystemInfo) Reset()         { *m = SystemInfo{} }
func (m *SystemInfo) String() string { return proto.CompactTextString(m) }
func (*SystemInfo) ProtoMessage()    {}
func (*SystemInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_aea16395861e483e, []int{1}
}
func (m *SystemInfo) XXX_UnmarshalVT(b []byte) error {
	return m.UnmarshalVT(b)
}
func (m *SystemInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SystemInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SystemInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemInfo.Merge(m, src)
}
func (m *SystemInfo) XXX_Size() int {
	return m.Size()
}
func (m *SystemInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SystemInfo proto.InternalMessageInfo

func (m *SystemInfo) GetBackupInfo() *BackupInfo {
	if m != nil {
		return m.BackupInfo
	}
	return nil
}

func (m *SystemInfo) MessageClone() proto.Message {
	return m.Clone()
}
func (m *SystemInfo) Clone() *SystemInfo {
	if m == nil {
		return nil
	}
	cloned := new(SystemInfo)
	*cloned = *m

	cloned.BackupInfo = m.BackupInfo.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*BackupInfo)(nil), "storage.BackupInfo")
	proto.RegisterType((*SystemInfo)(nil), "storage.SystemInfo")
}

func init() { proto.RegisterFile("storage/system_info.proto", fileDescriptor_aea16395861e483e) }

var fileDescriptor_aea16395861e483e = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x5d, 0x85, 0x8a, 0x5b, 0x50, 0x5c, 0x2f, 0x69, 0x0f, 0xb1, 0xf4, 0xd4, 0xd3, 0x46,
	0x6a, 0x5f, 0xc0, 0x5e, 0x44, 0x10, 0x84, 0x54, 0x2f, 0x5e, 0xc2, 0xa6, 0x6c, 0x42, 0x68, 0x92,
	0x8d, 0x3b, 0xb3, 0xa0, 0x6f, 0xe2, 0xa3, 0xf8, 0x08, 0x1e, 0x7d, 0x04, 0x89, 0x2f, 0x22, 0xdd,
	0x3f, 0xe9, 0x2d, 0x99, 0xef, 0x9b, 0xf9, 0x7e, 0x33, 0x4b, 0x27, 0x80, 0x4a, 0x8b, 0x52, 0x26,
	0xf0, 0x01, 0x28, 0x9b, 0xac, 0x6a, 0x0b, 0xc5, 0x3b, 0xad, 0x50, 0xb1, 0x53, 0x2f, 0x4d, 0xaf,
	0x4b, 0xa5, 0xca, 0x5a, 0x26, 0xb6, 0x9c, 0x9b, 0x22, 0xc1, 0xaa, 0x91, 0x80, 0xa2, 0xe9, 0x9c,
	0x73, 0x1a, 0x87, 0x21, 0xaa, 0x93, 0x5a, 0x60, 0xa5, 0xda, 0x0c, 0x50, 0xa0, 0x01, 0xaf, 0xb3,
	0xa0, 0x1b, 0x90, 0xda, 0xd5, 0xe6, 0x5f, 0x84, 0xd2, 0xb5, 0xd8, 0xee, 0x4c, 0xf7, 0xd0, 0x16,
	0x8a, 0xdd, 0x53, 0x96, 0xdb, 0xbf, 0xac, 0x16, 0x80, 0x99, 0x36, 0x6d, 0x26, 0x30, 0x22, 0x33,
	0xb2, 0x18, 0x2f, 0xa7, 0xdc, 0x01, 0xf0, 0x00, 0xc0, 0x9f, 0x03, 0x40, 0x7a, 0xe1, 0xba, 0x1e,
	0x05, 0x60, 0x6a, 0xda, 0x3b, 0x64, 0x37, 0x74, 0xe4, 0xb2, 0xa3, 0xe3, 0x19, 0x59, 0x9c, 0x2f,
	0x23, 0xee, 0xc3, 0xf9, 0x53, 0x80, 0xdb, 0x58, 0x3d, 0xf5, 0x3e, 0x96, 0xd0, 0x33, 0x2d, 0xdf,
	0x8c, 0xdc, 0xfb, 0xa2, 0x13, 0x9b, 0x78, 0x39, 0x34, 0x6d, 0xea, 0xaa, 0x79, 0x01, 0xa9, 0xd3,
	0x83, 0x67, 0xbe, 0xa6, 0x74, 0x63, 0xaf, 0x65, 0xc9, 0x57, 0x74, 0xec, 0xc9, 0xf7, 0xb7, 0xf3,
	0xc8, 0x57, 0xc3, 0x80, 0xc3, 0x8e, 0x29, 0xcd, 0x87, 0xef, 0xf5, 0xea, 0xbb, 0x8f, 0xc9, 0x4f,
	0x1f, 0x93, 0xdf, 0x3e, 0x26, 0x9f, 0x7f, 0xf1, 0x11, 0x9d, 0x54, 0x8a, 0x03, 0x8a, 0xed, 0x4e,
	0xab, 0x77, 0xb7, 0x69, 0x98, 0xf1, 0x1a, 0x5e, 0x22, 0x1f, 0xd9, 0xfa, 0xed, 0x7f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x10, 0xc1, 0xf1, 0x44, 0xb6, 0x01, 0x00, 0x00,
}

func (m *BackupInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BackupInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Requestor != nil {
		{
			size, err := m.Requestor.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSystemInfo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Status != 0 {
		i = encodeVarintSystemInfo(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if m.BackupLastRunAt != nil {
		{
			size, err := m.BackupLastRunAt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSystemInfo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SystemInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SystemInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SystemInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.BackupInfo != nil {
		{
			size, err := m.BackupInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSystemInfo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSystemInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovSystemInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BackupInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BackupLastRunAt != nil {
		l = m.BackupLastRunAt.Size()
		n += 1 + l + sovSystemInfo(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovSystemInfo(uint64(m.Status))
	}
	if m.Requestor != nil {
		l = m.Requestor.Size()
		n += 1 + l + sovSystemInfo(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SystemInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BackupInfo != nil {
		l = m.BackupInfo.Size()
		n += 1 + l + sovSystemInfo(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSystemInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSystemInfo(x uint64) (n int) {
	return sovSystemInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BackupInfo) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSystemInfo
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
			return fmt.Errorf("proto: BackupInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BackupLastRunAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystemInfo
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
				return ErrInvalidLengthSystemInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSystemInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BackupLastRunAt == nil {
				m.BackupLastRunAt = &types.Timestamp{}
			}
			if err := m.BackupLastRunAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystemInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= OperationStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requestor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystemInfo
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
				return ErrInvalidLengthSystemInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSystemInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Requestor == nil {
				m.Requestor = &SlimUser{}
			}
			if err := m.Requestor.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSystemInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSystemInfo
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
func (m *SystemInfo) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSystemInfo
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
			return fmt.Errorf("proto: SystemInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SystemInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BackupInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystemInfo
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
				return ErrInvalidLengthSystemInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSystemInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BackupInfo == nil {
				m.BackupInfo = &BackupInfo{}
			}
			if err := m.BackupInfo.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSystemInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSystemInfo
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
func skipSystemInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSystemInfo
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
					return 0, ErrIntOverflowSystemInfo
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
					return 0, ErrIntOverflowSystemInfo
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
				return 0, ErrInvalidLengthSystemInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSystemInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSystemInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSystemInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSystemInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSystemInfo = fmt.Errorf("proto: unexpected end of group")
)
