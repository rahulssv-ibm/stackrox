// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.1-0.20240319094008-0393e58bdf10
// source: internalapi/central/process_listening_on_ports_update.proto

package central

import (
	fmt "fmt"
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	timestamppb1 "github.com/planetscale/vtprotobuf/types/known/timestamppb"
	storage "github.com/stackrox/rox/generated/storage"
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

func (m *ProcessListeningOnPortsUpdate) CloneVT() *ProcessListeningOnPortsUpdate {
	if m == nil {
		return (*ProcessListeningOnPortsUpdate)(nil)
	}
	r := new(ProcessListeningOnPortsUpdate)
	r.Time = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.Time).CloneVT())
	if rhs := m.ProcessesListeningOnPorts; rhs != nil {
		tmpContainer := make([]*storage.ProcessListeningOnPortFromSensor, len(rhs))
		for k, v := range rhs {
			if vtpb, ok := interface{}(v).(interface {
				CloneVT() *storage.ProcessListeningOnPortFromSensor
			}); ok {
				tmpContainer[k] = vtpb.CloneVT()
			} else {
				tmpContainer[k] = proto.Clone(v).(*storage.ProcessListeningOnPortFromSensor)
			}
		}
		r.ProcessesListeningOnPorts = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ProcessListeningOnPortsUpdate) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *ProcessListeningOnPortsUpdate) EqualVT(that *ProcessListeningOnPortsUpdate) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.ProcessesListeningOnPorts) != len(that.ProcessesListeningOnPorts) {
		return false
	}
	for i, vx := range this.ProcessesListeningOnPorts {
		vy := that.ProcessesListeningOnPorts[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &storage.ProcessListeningOnPortFromSensor{}
			}
			if q == nil {
				q = &storage.ProcessListeningOnPortFromSensor{}
			}
			if equal, ok := interface{}(p).(interface {
				EqualVT(*storage.ProcessListeningOnPortFromSensor) bool
			}); ok {
				if !equal.EqualVT(q) {
					return false
				}
			} else if !proto.Equal(p, q) {
				return false
			}
		}
	}
	if !(*timestamppb1.Timestamp)(this.Time).EqualVT((*timestamppb1.Timestamp)(that.Time)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ProcessListeningOnPortsUpdate) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ProcessListeningOnPortsUpdate)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *ProcessListeningOnPortsUpdate) MarshalVT() (dAtA []byte, err error) {
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

func (m *ProcessListeningOnPortsUpdate) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ProcessListeningOnPortsUpdate) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.Time != nil {
		size, err := (*timestamppb1.Timestamp)(m.Time).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ProcessesListeningOnPorts) > 0 {
		for iNdEx := len(m.ProcessesListeningOnPorts) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.ProcessesListeningOnPorts[iNdEx]).(interface {
				MarshalToSizedBufferVT([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.ProcessesListeningOnPorts[iNdEx])
				if err != nil {
					return 0, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ProcessListeningOnPortsUpdate) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ProcessesListeningOnPorts) > 0 {
		for _, e := range m.ProcessesListeningOnPorts {
			if size, ok := interface{}(e).(interface {
				SizeVT() int
			}); ok {
				l = size.SizeVT()
			} else {
				l = proto.Size(e)
			}
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.Time != nil {
		l = (*timestamppb1.Timestamp)(m.Time).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ProcessListeningOnPortsUpdate) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: ProcessListeningOnPortsUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessListeningOnPortsUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProcessesListeningOnPorts", wireType)
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
			m.ProcessesListeningOnPorts = append(m.ProcessesListeningOnPorts, &storage.ProcessListeningOnPortFromSensor{})
			if unmarshal, ok := interface{}(m.ProcessesListeningOnPorts[len(m.ProcessesListeningOnPorts)-1]).(interface {
				UnmarshalVT([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.ProcessesListeningOnPorts[len(m.ProcessesListeningOnPorts)-1]); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
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
			if m.Time == nil {
				m.Time = &timestamppb.Timestamp{}
			}
			if err := (*timestamppb1.Timestamp)(m.Time).UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *ProcessListeningOnPortsUpdate) UnmarshalVTUnsafe(dAtA []byte) error {
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
			return fmt.Errorf("proto: ProcessListeningOnPortsUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessListeningOnPortsUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProcessesListeningOnPorts", wireType)
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
			m.ProcessesListeningOnPorts = append(m.ProcessesListeningOnPorts, &storage.ProcessListeningOnPortFromSensor{})
			if unmarshal, ok := interface{}(m.ProcessesListeningOnPorts[len(m.ProcessesListeningOnPorts)-1]).(interface {
				UnmarshalVTUnsafe([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVTUnsafe(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.ProcessesListeningOnPorts[len(m.ProcessesListeningOnPorts)-1]); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
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
			if m.Time == nil {
				m.Time = &timestamppb.Timestamp{}
			}
			if err := (*timestamppb1.Timestamp)(m.Time).UnmarshalVTUnsafe(dAtA[iNdEx:postIndex]); err != nil {
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
