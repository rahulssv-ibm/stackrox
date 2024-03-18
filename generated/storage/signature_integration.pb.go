// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/signature_integration.proto

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

type SignatureIntegration struct {
	Id                   string                       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk"`     // @gotags: sql:"pk"
	Name                 string                       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" sql:"unique"` // @gotags: sql:"unique"
	Cosign               *CosignPublicKeyVerification `protobuf:"bytes,3,opt,name=cosign,proto3" json:"cosign,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SignatureIntegration) Reset()         { *m = SignatureIntegration{} }
func (m *SignatureIntegration) String() string { return proto.CompactTextString(m) }
func (*SignatureIntegration) ProtoMessage()    {}
func (*SignatureIntegration) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3165e7a4c19e14a, []int{0}
}
func (m *SignatureIntegration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignatureIntegration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignatureIntegration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignatureIntegration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureIntegration.Merge(m, src)
}
func (m *SignatureIntegration) XXX_Size() int {
	return m.Size()
}
func (m *SignatureIntegration) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureIntegration.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureIntegration proto.InternalMessageInfo

func (m *SignatureIntegration) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SignatureIntegration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SignatureIntegration) GetCosign() *CosignPublicKeyVerification {
	if m != nil {
		return m.Cosign
	}
	return nil
}

func (m *SignatureIntegration) MessageClone() proto.Message {
	return m.Clone()
}
func (m *SignatureIntegration) Clone() *SignatureIntegration {
	if m == nil {
		return nil
	}
	cloned := new(SignatureIntegration)
	*cloned = *m

	cloned.Cosign = m.Cosign.Clone()
	return cloned
}

type CosignPublicKeyVerification struct {
	PublicKeys           []*CosignPublicKeyVerification_PublicKey `protobuf:"bytes,3,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *CosignPublicKeyVerification) Reset()         { *m = CosignPublicKeyVerification{} }
func (m *CosignPublicKeyVerification) String() string { return proto.CompactTextString(m) }
func (*CosignPublicKeyVerification) ProtoMessage()    {}
func (*CosignPublicKeyVerification) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3165e7a4c19e14a, []int{1}
}
func (m *CosignPublicKeyVerification) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CosignPublicKeyVerification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CosignPublicKeyVerification.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CosignPublicKeyVerification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CosignPublicKeyVerification.Merge(m, src)
}
func (m *CosignPublicKeyVerification) XXX_Size() int {
	return m.Size()
}
func (m *CosignPublicKeyVerification) XXX_DiscardUnknown() {
	xxx_messageInfo_CosignPublicKeyVerification.DiscardUnknown(m)
}

var xxx_messageInfo_CosignPublicKeyVerification proto.InternalMessageInfo

func (m *CosignPublicKeyVerification) GetPublicKeys() []*CosignPublicKeyVerification_PublicKey {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

func (m *CosignPublicKeyVerification) MessageClone() proto.Message {
	return m.Clone()
}
func (m *CosignPublicKeyVerification) Clone() *CosignPublicKeyVerification {
	if m == nil {
		return nil
	}
	cloned := new(CosignPublicKeyVerification)
	*cloned = *m

	if m.PublicKeys != nil {
		cloned.PublicKeys = make([]*CosignPublicKeyVerification_PublicKey, len(m.PublicKeys))
		for idx, v := range m.PublicKeys {
			cloned.PublicKeys[idx] = v.Clone()
		}
	}
	return cloned
}

type CosignPublicKeyVerification_PublicKey struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PublicKeyPemEnc      string   `protobuf:"bytes,2,opt,name=public_key_pem_enc,json=publicKeyPemEnc,proto3" json:"public_key_pem_enc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CosignPublicKeyVerification_PublicKey) Reset()         { *m = CosignPublicKeyVerification_PublicKey{} }
func (m *CosignPublicKeyVerification_PublicKey) String() string { return proto.CompactTextString(m) }
func (*CosignPublicKeyVerification_PublicKey) ProtoMessage()    {}
func (*CosignPublicKeyVerification_PublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3165e7a4c19e14a, []int{1, 0}
}
func (m *CosignPublicKeyVerification_PublicKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CosignPublicKeyVerification_PublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CosignPublicKeyVerification_PublicKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CosignPublicKeyVerification_PublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CosignPublicKeyVerification_PublicKey.Merge(m, src)
}
func (m *CosignPublicKeyVerification_PublicKey) XXX_Size() int {
	return m.Size()
}
func (m *CosignPublicKeyVerification_PublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_CosignPublicKeyVerification_PublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_CosignPublicKeyVerification_PublicKey proto.InternalMessageInfo

func (m *CosignPublicKeyVerification_PublicKey) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CosignPublicKeyVerification_PublicKey) GetPublicKeyPemEnc() string {
	if m != nil {
		return m.PublicKeyPemEnc
	}
	return ""
}

func (m *CosignPublicKeyVerification_PublicKey) MessageClone() proto.Message {
	return m.Clone()
}
func (m *CosignPublicKeyVerification_PublicKey) Clone() *CosignPublicKeyVerification_PublicKey {
	if m == nil {
		return nil
	}
	cloned := new(CosignPublicKeyVerification_PublicKey)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterType((*SignatureIntegration)(nil), "storage.SignatureIntegration")
	proto.RegisterType((*CosignPublicKeyVerification)(nil), "storage.CosignPublicKeyVerification")
	proto.RegisterType((*CosignPublicKeyVerification_PublicKey)(nil), "storage.CosignPublicKeyVerification.PublicKey")
}

func init() {
	proto.RegisterFile("storage/signature_integration.proto", fileDescriptor_b3165e7a4c19e14a)
}

var fileDescriptor_b3165e7a4c19e14a = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x2f, 0xce, 0x4c, 0xcf, 0x4b, 0x2c, 0x29, 0x2d, 0x4a, 0x8d, 0xcf, 0xcc,
	0x2b, 0x49, 0x4d, 0x2f, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x87, 0x2a, 0x52, 0xaa, 0xe0, 0x12, 0x09, 0x86, 0xa9, 0xf3, 0x44, 0x28, 0x13, 0xe2, 0xe3,
	0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe2,
	0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0x8b, 0x80, 0xd9, 0x42, 0x36, 0x5c, 0x6c, 0xc9,
	0xf9, 0x20, 0x5b, 0x24, 0x98, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x54, 0xf4, 0xa0, 0xa6, 0xea, 0x39,
	0x83, 0x85, 0x03, 0x4a, 0x93, 0x72, 0x32, 0x93, 0xbd, 0x53, 0x2b, 0xc3, 0x52, 0x8b, 0x32, 0xd3,
	0x32, 0x93, 0xc1, 0x26, 0x07, 0x41, 0xf5, 0x28, 0xed, 0x61, 0xe4, 0x92, 0xc6, 0xa3, 0x4e, 0xc8,
	0x9f, 0x8b, 0xbb, 0x00, 0x2c, 0x11, 0x9f, 0x9d, 0x5a, 0x59, 0x2c, 0xc1, 0xac, 0xc0, 0xac, 0xc1,
	0x6d, 0xa4, 0x47, 0x8c, 0x15, 0x7a, 0x70, 0xd1, 0x20, 0xae, 0x02, 0x18, 0xb3, 0x58, 0xca, 0x87,
	0x8b, 0x13, 0x2e, 0x01, 0xf7, 0x0f, 0x23, 0x92, 0x7f, 0xb4, 0xb9, 0x84, 0x10, 0x36, 0xc6, 0x17,
	0xa4, 0xe6, 0xc6, 0xa7, 0xe6, 0x25, 0x43, 0x7d, 0xcc, 0x0f, 0x37, 0x28, 0x20, 0x35, 0xd7, 0x35,
	0x2f, 0xd9, 0xc9, 0xe4, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63,
	0x9c, 0xf1, 0x58, 0x8e, 0x81, 0x4b, 0x32, 0x33, 0x5f, 0xaf, 0xb8, 0x24, 0x31, 0x39, 0xbb, 0x28,
	0xbf, 0x02, 0x12, 0xcc, 0x30, 0xc7, 0x46, 0xc1, 0x82, 0x3b, 0x89, 0x0d, 0x2c, 0x6e, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x8d, 0xc3, 0x49, 0x33, 0xa5, 0x01, 0x00, 0x00,
}

func (m *SignatureIntegration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignatureIntegration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignatureIntegration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Cosign != nil {
		{
			size, err := m.Cosign.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSignatureIntegration(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSignatureIntegration(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintSignatureIntegration(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CosignPublicKeyVerification) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CosignPublicKeyVerification) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CosignPublicKeyVerification) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PublicKeys) > 0 {
		for iNdEx := len(m.PublicKeys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PublicKeys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSignatureIntegration(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	return len(dAtA) - i, nil
}

func (m *CosignPublicKeyVerification_PublicKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CosignPublicKeyVerification_PublicKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CosignPublicKeyVerification_PublicKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PublicKeyPemEnc) > 0 {
		i -= len(m.PublicKeyPemEnc)
		copy(dAtA[i:], m.PublicKeyPemEnc)
		i = encodeVarintSignatureIntegration(dAtA, i, uint64(len(m.PublicKeyPemEnc)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSignatureIntegration(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSignatureIntegration(dAtA []byte, offset int, v uint64) int {
	offset -= sovSignatureIntegration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SignatureIntegration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSignatureIntegration(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSignatureIntegration(uint64(l))
	}
	if m.Cosign != nil {
		l = m.Cosign.Size()
		n += 1 + l + sovSignatureIntegration(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CosignPublicKeyVerification) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PublicKeys) > 0 {
		for _, e := range m.PublicKeys {
			l = e.Size()
			n += 1 + l + sovSignatureIntegration(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CosignPublicKeyVerification_PublicKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSignatureIntegration(uint64(l))
	}
	l = len(m.PublicKeyPemEnc)
	if l > 0 {
		n += 1 + l + sovSignatureIntegration(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSignatureIntegration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSignatureIntegration(x uint64) (n int) {
	return sovSignatureIntegration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignatureIntegration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignatureIntegration
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
			return fmt.Errorf("proto: SignatureIntegration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignatureIntegration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignatureIntegration
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
				return ErrInvalidLengthSignatureIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignatureIntegration
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
					return ErrIntOverflowSignatureIntegration
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
				return ErrInvalidLengthSignatureIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignatureIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cosign", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignatureIntegration
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
				return ErrInvalidLengthSignatureIntegration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSignatureIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Cosign == nil {
				m.Cosign = &CosignPublicKeyVerification{}
			}
			if err := m.Cosign.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSignatureIntegration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSignatureIntegration
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
func (m *CosignPublicKeyVerification) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignatureIntegration
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
			return fmt.Errorf("proto: CosignPublicKeyVerification: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CosignPublicKeyVerification: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignatureIntegration
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
				return ErrInvalidLengthSignatureIntegration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSignatureIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKeys = append(m.PublicKeys, &CosignPublicKeyVerification_PublicKey{})
			if err := m.PublicKeys[len(m.PublicKeys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSignatureIntegration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSignatureIntegration
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
func (m *CosignPublicKeyVerification_PublicKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignatureIntegration
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
			return fmt.Errorf("proto: PublicKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublicKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignatureIntegration
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
				return ErrInvalidLengthSignatureIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignatureIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKeyPemEnc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignatureIntegration
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
				return ErrInvalidLengthSignatureIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignatureIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKeyPemEnc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSignatureIntegration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSignatureIntegration
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
func skipSignatureIntegration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSignatureIntegration
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
					return 0, ErrIntOverflowSignatureIntegration
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
					return 0, ErrIntOverflowSignatureIntegration
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
				return 0, ErrInvalidLengthSignatureIntegration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSignatureIntegration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSignatureIntegration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSignatureIntegration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSignatureIntegration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSignatureIntegration = fmt.Errorf("proto: unexpected end of group")
)
