// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: api/v1/central_health_service.proto

package v1

import (
	fmt "fmt"
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *GetUpgradeStatusResponse) CloneVT() *GetUpgradeStatusResponse {
	if m == nil {
		return (*GetUpgradeStatusResponse)(nil)
	}
	r := new(GetUpgradeStatusResponse)
	r.UpgradeStatus = m.UpgradeStatus.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetUpgradeStatusResponse) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *CentralUpgradeStatus) CloneVT() *CentralUpgradeStatus {
	if m == nil {
		return (*CentralUpgradeStatus)(nil)
	}
	r := new(CentralUpgradeStatus)
	r.Version = m.Version
	r.ForceRollbackTo = m.ForceRollbackTo
	r.CanRollbackAfterUpgrade = m.CanRollbackAfterUpgrade
	r.SpaceRequiredForRollbackAfterUpgrade = m.SpaceRequiredForRollbackAfterUpgrade
	r.SpaceAvailableForRollbackAfterUpgrade = m.SpaceAvailableForRollbackAfterUpgrade
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *CentralUpgradeStatus) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *GetUpgradeStatusResponse) EqualVT(that *GetUpgradeStatusResponse) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.UpgradeStatus.EqualVT(that.UpgradeStatus) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetUpgradeStatusResponse) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetUpgradeStatusResponse)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *CentralUpgradeStatus) EqualVT(that *CentralUpgradeStatus) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	if this.ForceRollbackTo != that.ForceRollbackTo {
		return false
	}
	if this.CanRollbackAfterUpgrade != that.CanRollbackAfterUpgrade {
		return false
	}
	if this.SpaceRequiredForRollbackAfterUpgrade != that.SpaceRequiredForRollbackAfterUpgrade {
		return false
	}
	if this.SpaceAvailableForRollbackAfterUpgrade != that.SpaceAvailableForRollbackAfterUpgrade {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CentralUpgradeStatus) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*CentralUpgradeStatus)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *GetUpgradeStatusResponse) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetUpgradeStatusResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *GetUpgradeStatusResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.UpgradeStatus != nil {
		size, err := m.UpgradeStatus.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CentralUpgradeStatus) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CentralUpgradeStatus) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *CentralUpgradeStatus) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.SpaceAvailableForRollbackAfterUpgrade != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.SpaceAvailableForRollbackAfterUpgrade))
		i--
		dAtA[i] = 0x28
	}
	if m.SpaceRequiredForRollbackAfterUpgrade != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.SpaceRequiredForRollbackAfterUpgrade))
		i--
		dAtA[i] = 0x20
	}
	if m.CanRollbackAfterUpgrade {
		i--
		if m.CanRollbackAfterUpgrade {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.ForceRollbackTo) > 0 {
		i -= len(m.ForceRollbackTo)
		copy(dAtA[i:], m.ForceRollbackTo)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ForceRollbackTo)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetUpgradeStatusResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UpgradeStatus != nil {
		l = m.UpgradeStatus.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *CentralUpgradeStatus) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.ForceRollbackTo)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.CanRollbackAfterUpgrade {
		n += 2
	}
	if m.SpaceRequiredForRollbackAfterUpgrade != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.SpaceRequiredForRollbackAfterUpgrade))
	}
	if m.SpaceAvailableForRollbackAfterUpgrade != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.SpaceAvailableForRollbackAfterUpgrade))
	}
	n += len(m.unknownFields)
	return n
}

func (m *GetUpgradeStatusResponse) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
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
			return fmt.Errorf("proto: GetUpgradeStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetUpgradeStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpgradeStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpgradeStatus == nil {
				m.UpgradeStatus = &CentralUpgradeStatus{}
			}
			if err := m.UpgradeStatus.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CentralUpgradeStatus) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
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
			return fmt.Errorf("proto: CentralUpgradeStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CentralUpgradeStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForceRollbackTo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ForceRollbackTo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanRollbackAfterUpgrade", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
			m.CanRollbackAfterUpgrade = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceRequiredForRollbackAfterUpgrade", wireType)
			}
			m.SpaceRequiredForRollbackAfterUpgrade = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpaceRequiredForRollbackAfterUpgrade |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceAvailableForRollbackAfterUpgrade", wireType)
			}
			m.SpaceAvailableForRollbackAfterUpgrade = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpaceAvailableForRollbackAfterUpgrade |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
