// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: api/v1/administration_usage_service.proto

package v1

import (
	fmt "fmt"
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	timestamppb1 "github.com/planetscale/vtprotobuf/types/known/timestamppb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *TimeRange) CloneVT() *TimeRange {
	if m == nil {
		return (*TimeRange)(nil)
	}
	r := new(TimeRange)
	r.From = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.From).CloneVT())
	r.To = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.To).CloneVT())
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *TimeRange) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *SecuredUnitsUsageResponse) CloneVT() *SecuredUnitsUsageResponse {
	if m == nil {
		return (*SecuredUnitsUsageResponse)(nil)
	}
	r := new(SecuredUnitsUsageResponse)
	r.NumNodes = m.NumNodes
	r.NumCpuUnits = m.NumCpuUnits
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *SecuredUnitsUsageResponse) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *MaxSecuredUnitsUsageResponse) CloneVT() *MaxSecuredUnitsUsageResponse {
	if m == nil {
		return (*MaxSecuredUnitsUsageResponse)(nil)
	}
	r := new(MaxSecuredUnitsUsageResponse)
	r.MaxNodesAt = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.MaxNodesAt).CloneVT())
	r.MaxNodes = m.MaxNodes
	r.MaxCpuUnitsAt = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.MaxCpuUnitsAt).CloneVT())
	r.MaxCpuUnits = m.MaxCpuUnits
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *MaxSecuredUnitsUsageResponse) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *TimeRange) EqualVT(that *TimeRange) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.From).EqualVT((*timestamppb1.Timestamp)(that.From)) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.To).EqualVT((*timestamppb1.Timestamp)(that.To)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *TimeRange) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*TimeRange)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *SecuredUnitsUsageResponse) EqualVT(that *SecuredUnitsUsageResponse) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.NumNodes != that.NumNodes {
		return false
	}
	if this.NumCpuUnits != that.NumCpuUnits {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *SecuredUnitsUsageResponse) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*SecuredUnitsUsageResponse)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *MaxSecuredUnitsUsageResponse) EqualVT(that *MaxSecuredUnitsUsageResponse) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.MaxNodesAt).EqualVT((*timestamppb1.Timestamp)(that.MaxNodesAt)) {
		return false
	}
	if this.MaxNodes != that.MaxNodes {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.MaxCpuUnitsAt).EqualVT((*timestamppb1.Timestamp)(that.MaxCpuUnitsAt)) {
		return false
	}
	if this.MaxCpuUnits != that.MaxCpuUnits {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *MaxSecuredUnitsUsageResponse) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*MaxSecuredUnitsUsageResponse)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *TimeRange) MarshalVT() (dAtA []byte, err error) {
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

func (m *TimeRange) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *TimeRange) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.To != nil {
		size, err := (*timestamppb1.Timestamp)(m.To).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.From != nil {
		size, err := (*timestamppb1.Timestamp)(m.From).MarshalToSizedBufferVT(dAtA[:i])
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

func (m *SecuredUnitsUsageResponse) MarshalVT() (dAtA []byte, err error) {
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

func (m *SecuredUnitsUsageResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *SecuredUnitsUsageResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.NumCpuUnits != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.NumCpuUnits))
		i--
		dAtA[i] = 0x10
	}
	if m.NumNodes != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.NumNodes))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MaxSecuredUnitsUsageResponse) MarshalVT() (dAtA []byte, err error) {
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

func (m *MaxSecuredUnitsUsageResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *MaxSecuredUnitsUsageResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.MaxCpuUnits != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.MaxCpuUnits))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxCpuUnitsAt != nil {
		size, err := (*timestamppb1.Timestamp)(m.MaxCpuUnitsAt).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.MaxNodes != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.MaxNodes))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxNodesAt != nil {
		size, err := (*timestamppb1.Timestamp)(m.MaxNodesAt).MarshalToSizedBufferVT(dAtA[:i])
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

func (m *TimeRange) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.From != nil {
		l = (*timestamppb1.Timestamp)(m.From).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.To != nil {
		l = (*timestamppb1.Timestamp)(m.To).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *SecuredUnitsUsageResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumNodes != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.NumNodes))
	}
	if m.NumCpuUnits != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.NumCpuUnits))
	}
	n += len(m.unknownFields)
	return n
}

func (m *MaxSecuredUnitsUsageResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxNodesAt != nil {
		l = (*timestamppb1.Timestamp)(m.MaxNodesAt).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.MaxNodes != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.MaxNodes))
	}
	if m.MaxCpuUnitsAt != nil {
		l = (*timestamppb1.Timestamp)(m.MaxCpuUnitsAt).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.MaxCpuUnits != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.MaxCpuUnits))
	}
	n += len(m.unknownFields)
	return n
}

func (m *TimeRange) UnmarshalVTUnsafe(dAtA []byte) error {
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
			return fmt.Errorf("proto: TimeRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			if m.From == nil {
				m.From = &timestamppb.Timestamp{}
			}
			if err := (*timestamppb1.Timestamp)(m.From).UnmarshalVTUnsafe(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
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
			if m.To == nil {
				m.To = &timestamppb.Timestamp{}
			}
			if err := (*timestamppb1.Timestamp)(m.To).UnmarshalVTUnsafe(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *SecuredUnitsUsageResponse) UnmarshalVTUnsafe(dAtA []byte) error {
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
			return fmt.Errorf("proto: SecuredUnitsUsageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SecuredUnitsUsageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumNodes", wireType)
			}
			m.NumNodes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumNodes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumCpuUnits", wireType)
			}
			m.NumCpuUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumCpuUnits |= int64(b&0x7F) << shift
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
func (m *MaxSecuredUnitsUsageResponse) UnmarshalVTUnsafe(dAtA []byte) error {
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
			return fmt.Errorf("proto: MaxSecuredUnitsUsageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MaxSecuredUnitsUsageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNodesAt", wireType)
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
			if m.MaxNodesAt == nil {
				m.MaxNodesAt = &timestamppb.Timestamp{}
			}
			if err := (*timestamppb1.Timestamp)(m.MaxNodesAt).UnmarshalVTUnsafe(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNodes", wireType)
			}
			m.MaxNodes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNodes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxCpuUnitsAt", wireType)
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
			if m.MaxCpuUnitsAt == nil {
				m.MaxCpuUnitsAt = &timestamppb.Timestamp{}
			}
			if err := (*timestamppb1.Timestamp)(m.MaxCpuUnitsAt).UnmarshalVTUnsafe(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxCpuUnits", wireType)
			}
			m.MaxCpuUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxCpuUnits |= int64(b&0x7F) << shift
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
