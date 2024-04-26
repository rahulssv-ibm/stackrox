// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: internalapi/central/local_scanner.proto

package central

import (
	fmt "fmt"
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	storage "github.com/stackrox/rox/generated/storage"
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

func (m *LocalScannerCertsIssueError) CloneVT() *LocalScannerCertsIssueError {
	if m == nil {
		return (*LocalScannerCertsIssueError)(nil)
	}
	r := new(LocalScannerCertsIssueError)
	r.Message = m.Message
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *LocalScannerCertsIssueError) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *IssueLocalScannerCertsRequest) CloneVT() *IssueLocalScannerCertsRequest {
	if m == nil {
		return (*IssueLocalScannerCertsRequest)(nil)
	}
	r := new(IssueLocalScannerCertsRequest)
	r.RequestId = m.RequestId
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *IssueLocalScannerCertsRequest) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *IssueLocalScannerCertsResponse) CloneVT() *IssueLocalScannerCertsResponse {
	if m == nil {
		return (*IssueLocalScannerCertsResponse)(nil)
	}
	r := new(IssueLocalScannerCertsResponse)
	r.RequestId = m.RequestId
	if m.Response != nil {
		r.Response = m.Response.(interface {
			CloneVT() isIssueLocalScannerCertsResponse_Response
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *IssueLocalScannerCertsResponse) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *IssueLocalScannerCertsResponse_Certificates) CloneVT() isIssueLocalScannerCertsResponse_Response {
	if m == nil {
		return (*IssueLocalScannerCertsResponse_Certificates)(nil)
	}
	r := new(IssueLocalScannerCertsResponse_Certificates)
	if rhs := m.Certificates; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface {
			CloneVT() *storage.TypedServiceCertificateSet
		}); ok {
			r.Certificates = vtpb.CloneVT()
		} else {
			r.Certificates = proto.Clone(rhs).(*storage.TypedServiceCertificateSet)
		}
	}
	return r
}

func (m *IssueLocalScannerCertsResponse_Error) CloneVT() isIssueLocalScannerCertsResponse_Response {
	if m == nil {
		return (*IssueLocalScannerCertsResponse_Error)(nil)
	}
	r := new(IssueLocalScannerCertsResponse_Error)
	r.Error = m.Error.CloneVT()
	return r
}

func (this *LocalScannerCertsIssueError) EqualVT(that *LocalScannerCertsIssueError) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Message != that.Message {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *LocalScannerCertsIssueError) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*LocalScannerCertsIssueError)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *IssueLocalScannerCertsRequest) EqualVT(that *IssueLocalScannerCertsRequest) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.RequestId != that.RequestId {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *IssueLocalScannerCertsRequest) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*IssueLocalScannerCertsRequest)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *IssueLocalScannerCertsResponse) EqualVT(that *IssueLocalScannerCertsResponse) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Response == nil && that.Response != nil {
		return false
	} else if this.Response != nil {
		if that.Response == nil {
			return false
		}
		if !this.Response.(interface {
			EqualVT(isIssueLocalScannerCertsResponse_Response) bool
		}).EqualVT(that.Response) {
			return false
		}
	}
	if this.RequestId != that.RequestId {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *IssueLocalScannerCertsResponse) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*IssueLocalScannerCertsResponse)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *IssueLocalScannerCertsResponse_Certificates) EqualVT(thatIface isIssueLocalScannerCertsResponse_Response) bool {
	that, ok := thatIface.(*IssueLocalScannerCertsResponse_Certificates)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Certificates, that.Certificates; p != q {
		if p == nil {
			p = &storage.TypedServiceCertificateSet{}
		}
		if q == nil {
			q = &storage.TypedServiceCertificateSet{}
		}
		if equal, ok := interface{}(p).(interface {
			EqualVT(*storage.TypedServiceCertificateSet) bool
		}); ok {
			if !equal.EqualVT(q) {
				return false
			}
		} else if !proto.Equal(p, q) {
			return false
		}
	}
	return true
}

func (this *IssueLocalScannerCertsResponse_Error) EqualVT(thatIface isIssueLocalScannerCertsResponse_Response) bool {
	that, ok := thatIface.(*IssueLocalScannerCertsResponse_Error)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Error, that.Error; p != q {
		if p == nil {
			p = &LocalScannerCertsIssueError{}
		}
		if q == nil {
			q = &LocalScannerCertsIssueError{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (m *LocalScannerCertsIssueError) MarshalVT() (dAtA []byte, err error) {
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

func (m *LocalScannerCertsIssueError) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *LocalScannerCertsIssueError) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IssueLocalScannerCertsRequest) MarshalVT() (dAtA []byte, err error) {
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

func (m *IssueLocalScannerCertsRequest) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *IssueLocalScannerCertsRequest) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.RequestId) > 0 {
		i -= len(m.RequestId)
		copy(dAtA[i:], m.RequestId)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.RequestId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IssueLocalScannerCertsResponse) MarshalVT() (dAtA []byte, err error) {
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

func (m *IssueLocalScannerCertsResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *IssueLocalScannerCertsResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if vtmsg, ok := m.Response.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if len(m.RequestId) > 0 {
		i -= len(m.RequestId)
		copy(dAtA[i:], m.RequestId)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.RequestId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IssueLocalScannerCertsResponse_Certificates) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *IssueLocalScannerCertsResponse_Certificates) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Certificates != nil {
		if vtmsg, ok := interface{}(m.Certificates).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Certificates)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *IssueLocalScannerCertsResponse_Error) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *IssueLocalScannerCertsResponse_Error) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Error != nil {
		size, err := m.Error.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *LocalScannerCertsIssueError) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *IssueLocalScannerCertsRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequestId)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *IssueLocalScannerCertsResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequestId)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if vtmsg, ok := m.Response.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *IssueLocalScannerCertsResponse_Certificates) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Certificates != nil {
		if size, ok := interface{}(m.Certificates).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Certificates)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	return n
}
func (m *IssueLocalScannerCertsResponse_Error) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Error != nil {
		l = m.Error.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	return n
}
func (m *LocalScannerCertsIssueError) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: LocalScannerCertsIssueError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocalScannerCertsIssueError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
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
			m.Message = string(dAtA[iNdEx:postIndex])
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
func (m *IssueLocalScannerCertsRequest) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: IssueLocalScannerCertsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IssueLocalScannerCertsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestId", wireType)
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
			m.RequestId = string(dAtA[iNdEx:postIndex])
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
func (m *IssueLocalScannerCertsResponse) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: IssueLocalScannerCertsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IssueLocalScannerCertsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestId", wireType)
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
			m.RequestId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificates", wireType)
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
			if oneof, ok := m.Response.(*IssueLocalScannerCertsResponse_Certificates); ok {
				if unmarshal, ok := interface{}(oneof.Certificates).(interface {
					UnmarshalVT([]byte) error
				}); ok {
					if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
						return err
					}
				} else {
					if err := proto.Unmarshal(dAtA[iNdEx:postIndex], oneof.Certificates); err != nil {
						return err
					}
				}
			} else {
				v := &storage.TypedServiceCertificateSet{}
				if unmarshal, ok := interface{}(v).(interface {
					UnmarshalVT([]byte) error
				}); ok {
					if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
						return err
					}
				} else {
					if err := proto.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
						return err
					}
				}
				m.Response = &IssueLocalScannerCertsResponse_Certificates{Certificates: v}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
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
			if oneof, ok := m.Response.(*IssueLocalScannerCertsResponse_Error); ok {
				if err := oneof.Error.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &LocalScannerCertsIssueError{}
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Response = &IssueLocalScannerCertsResponse_Error{Error: v}
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
