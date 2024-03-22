// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internalapi/central/network_flow.proto

package central

import (
	fmt "fmt"
	types "github.com/gogo/protobuf/types"
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

type NetworkFlowUpdate struct {
	// Network flows that were added or removed from the last time state was sent to Central.
	Updated              []*storage.NetworkFlow     `protobuf:"bytes,1,rep,name=updated,proto3" json:"updated,omitempty"`
	UpdatedEndpoints     []*storage.NetworkEndpoint `protobuf:"bytes,3,rep,name=updated_endpoints,json=updatedEndpoints,proto3" json:"updated_endpoints,omitempty"`
	Time                 *types.Timestamp           `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *NetworkFlowUpdate) Reset()         { *m = NetworkFlowUpdate{} }
func (m *NetworkFlowUpdate) String() string { return proto.CompactTextString(m) }
func (*NetworkFlowUpdate) ProtoMessage()    {}
func (*NetworkFlowUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_82ca8b583996f885, []int{0}
}
func (m *NetworkFlowUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkFlowUpdate) XXX_MarshalVT(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkFlowUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkFlowUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkFlowUpdate.Merge(m, src)
}
func (m *NetworkFlowUpdate) XXX_Size() int {
	return m.Size()
}
func (m *NetworkFlowUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkFlowUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkFlowUpdate proto.InternalMessageInfo

func (m *NetworkFlowUpdate) GetUpdated() []*storage.NetworkFlow {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *NetworkFlowUpdate) GetUpdatedEndpoints() []*storage.NetworkEndpoint {
	if m != nil {
		return m.UpdatedEndpoints
	}
	return nil
}

func (m *NetworkFlowUpdate) GetTime() *types.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *NetworkFlowUpdate) MessageClone() proto.Message {
	return m.Clone()
}
func (m *NetworkFlowUpdate) Clone() *NetworkFlowUpdate {
	if m == nil {
		return nil
	}
	cloned := new(NetworkFlowUpdate)
	*cloned = *m

	if m.Updated != nil {
		cloned.Updated = make([]*storage.NetworkFlow, len(m.Updated))
		for idx, v := range m.Updated {
			cloned.Updated[idx] = v.Clone()
		}
	}
	if m.UpdatedEndpoints != nil {
		cloned.UpdatedEndpoints = make([]*storage.NetworkEndpoint, len(m.UpdatedEndpoints))
		for idx, v := range m.UpdatedEndpoints {
			cloned.UpdatedEndpoints[idx] = v.Clone()
		}
	}
	cloned.Time = m.Time.Clone()
	return cloned
}

type PushNetworkEntitiesRequest struct {
	Entities             []*storage.NetworkEntityInfo `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
	SeqID                int64                        `protobuf:"varint,2,opt,name=seqID,proto3" json:"seqID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *PushNetworkEntitiesRequest) Reset()         { *m = PushNetworkEntitiesRequest{} }
func (m *PushNetworkEntitiesRequest) String() string { return proto.CompactTextString(m) }
func (*PushNetworkEntitiesRequest) ProtoMessage()    {}
func (*PushNetworkEntitiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82ca8b583996f885, []int{1}
}
func (m *PushNetworkEntitiesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushNetworkEntitiesRequest) XXX_MarshalVT(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushNetworkEntitiesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushNetworkEntitiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNetworkEntitiesRequest.Merge(m, src)
}
func (m *PushNetworkEntitiesRequest) XXX_Size() int {
	return m.Size()
}
func (m *PushNetworkEntitiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNetworkEntitiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushNetworkEntitiesRequest proto.InternalMessageInfo

func (m *PushNetworkEntitiesRequest) GetEntities() []*storage.NetworkEntityInfo {
	if m != nil {
		return m.Entities
	}
	return nil
}

func (m *PushNetworkEntitiesRequest) GetSeqID() int64 {
	if m != nil {
		return m.SeqID
	}
	return 0
}

func (m *PushNetworkEntitiesRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *PushNetworkEntitiesRequest) Clone() *PushNetworkEntitiesRequest {
	if m == nil {
		return nil
	}
	cloned := new(PushNetworkEntitiesRequest)
	*cloned = *m

	if m.Entities != nil {
		cloned.Entities = make([]*storage.NetworkEntityInfo, len(m.Entities))
		for idx, v := range m.Entities {
			cloned.Entities[idx] = v.Clone()
		}
	}
	return cloned
}

func init() {
	proto.RegisterType((*NetworkFlowUpdate)(nil), "central.NetworkFlowUpdate")
	proto.RegisterType((*PushNetworkEntitiesRequest)(nil), "central.PushNetworkEntitiesRequest")
}

func init() {
	proto.RegisterFile("internalapi/central/network_flow.proto", fileDescriptor_82ca8b583996f885)
}

var fileDescriptor_82ca8b583996f885 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4e, 0xf3, 0x30,
	0x14, 0x84, 0x7f, 0xff, 0x05, 0x8a, 0xdc, 0x0d, 0xb5, 0xba, 0x08, 0x5e, 0x84, 0xaa, 0x0b, 0xd4,
	0x95, 0x23, 0x15, 0x89, 0x03, 0x20, 0x8a, 0xd4, 0x0d, 0x42, 0x16, 0x6c, 0xd8, 0x54, 0x2e, 0x7d,
	0x0d, 0x86, 0xd4, 0x4e, 0xe3, 0x17, 0x45, 0xdc, 0x84, 0xb3, 0x70, 0x02, 0x96, 0x1c, 0x01, 0x85,
	0x8b, 0x20, 0x1c, 0xa7, 0x42, 0xc0, 0xee, 0x65, 0xf2, 0xcd, 0x68, 0x46, 0xa6, 0xc7, 0xda, 0x20,
	0x14, 0x46, 0x65, 0x2a, 0xd7, 0xc9, 0x1d, 0x18, 0x2c, 0x54, 0x96, 0x18, 0xc0, 0xca, 0x16, 0x8f,
	0xf3, 0x55, 0x66, 0x2b, 0x91, 0x17, 0x16, 0x2d, 0xeb, 0x86, 0x7f, 0xfc, 0x28, 0xb5, 0x36, 0xcd,
	0x20, 0xf1, 0xf2, 0xa2, 0x5c, 0x25, 0xa8, 0xd7, 0xe0, 0x50, 0xad, 0xf3, 0x86, 0xe4, 0xdc, 0xa1,
	0x2d, 0x54, 0x0a, 0x7f, 0xa4, 0x8c, 0x5e, 0x08, 0xed, 0x5f, 0x36, 0xf2, 0x45, 0x66, 0xab, 0x9b,
	0x7c, 0xa9, 0x10, 0x98, 0xa0, 0xdd, 0xd2, 0x5f, 0xcb, 0x88, 0x0c, 0x3b, 0xe3, 0xde, 0x64, 0x20,
	0x42, 0x86, 0xf8, 0x06, 0xcb, 0x16, 0x62, 0x53, 0xda, 0x0f, 0xe7, 0x1c, 0xcc, 0x32, 0xb7, 0xda,
	0xa0, 0x8b, 0x3a, 0xde, 0x19, 0xfd, 0x74, 0x4e, 0x03, 0x20, 0x0f, 0x82, 0xa5, 0x15, 0x1c, 0x13,
	0x74, 0xe7, 0xab, 0x7b, 0xf4, 0x7f, 0x48, 0xc6, 0xbd, 0x09, 0x17, 0xcd, 0x30, 0xd1, 0x0e, 0x13,
	0xd7, 0xed, 0x30, 0xe9, 0xb9, 0xd1, 0x03, 0xe5, 0x57, 0xa5, 0xbb, 0xdf, 0x06, 0xa3, 0x46, 0x0d,
	0x4e, 0xc2, 0xa6, 0x04, 0x87, 0xec, 0x94, 0xee, 0x43, 0x90, 0xc2, 0x0a, 0xfe, 0xbb, 0x0b, 0x6a,
	0x7c, 0x9a, 0x99, 0x95, 0x95, 0x5b, 0x96, 0x0d, 0xe8, 0xae, 0x83, 0xcd, 0xec, 0xdc, 0xd7, 0xe8,
	0xc8, 0xe6, 0xe3, 0xec, 0xf0, 0xb5, 0x8e, 0xc9, 0x5b, 0x1d, 0x93, 0xf7, 0x3a, 0x26, 0xcf, 0x1f,
	0xf1, 0xbf, 0xdb, 0xf6, 0x01, 0x16, 0x7b, 0xbe, 0xe0, 0xc9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xa7, 0xaf, 0xf1, 0x40, 0xba, 0x01, 0x00, 0x00,
}

func (m *NetworkFlowUpdate) MarshalVT() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkFlowUpdate) MarshalVTTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkFlowUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.UpdatedEndpoints) > 0 {
		for iNdEx := len(m.UpdatedEndpoints) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UpdatedEndpoints[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNetworkFlow(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Time != nil {
		{
			size, err := m.Time.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNetworkFlow(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Updated) > 0 {
		for iNdEx := len(m.Updated) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Updated[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNetworkFlow(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PushNetworkEntitiesRequest) MarshalVT() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushNetworkEntitiesRequest) MarshalVTTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PushNetworkEntitiesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.SeqID != 0 {
		i = encodeVarintNetworkFlow(dAtA, i, uint64(m.SeqID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Entities) > 0 {
		for iNdEx := len(m.Entities) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entities[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNetworkFlow(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintNetworkFlow(dAtA []byte, offset int, v uint64) int {
	offset -= sovNetworkFlow(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NetworkFlowUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Updated) > 0 {
		for _, e := range m.Updated {
			l = e.Size()
			n += 1 + l + sovNetworkFlow(uint64(l))
		}
	}
	if m.Time != nil {
		l = m.Time.Size()
		n += 1 + l + sovNetworkFlow(uint64(l))
	}
	if len(m.UpdatedEndpoints) > 0 {
		for _, e := range m.UpdatedEndpoints {
			l = e.Size()
			n += 1 + l + sovNetworkFlow(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PushNetworkEntitiesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entities) > 0 {
		for _, e := range m.Entities {
			l = e.Size()
			n += 1 + l + sovNetworkFlow(uint64(l))
		}
	}
	if m.SeqID != 0 {
		n += 1 + sovNetworkFlow(uint64(m.SeqID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNetworkFlow(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNetworkFlow(x uint64) (n int) {
	return sovNetworkFlow(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NetworkFlowUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkFlow
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
			return fmt.Errorf("proto: NetworkFlowUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkFlowUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updated", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkFlow
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
				return ErrInvalidLengthNetworkFlow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkFlow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Updated = append(m.Updated, &storage.NetworkFlow{})
			if err := m.Updated[len(m.Updated)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkFlow
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
				return ErrInvalidLengthNetworkFlow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkFlow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Time == nil {
				m.Time = &types.Timestamp{}
			}
			if err := m.Time.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedEndpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkFlow
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
				return ErrInvalidLengthNetworkFlow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkFlow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedEndpoints = append(m.UpdatedEndpoints, &storage.NetworkEndpoint{})
			if err := m.UpdatedEndpoints[len(m.UpdatedEndpoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkFlow(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNetworkFlow
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
func (m *PushNetworkEntitiesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkFlow
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
			return fmt.Errorf("proto: PushNetworkEntitiesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushNetworkEntitiesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkFlow
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
				return ErrInvalidLengthNetworkFlow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkFlow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entities = append(m.Entities, &storage.NetworkEntityInfo{})
			if err := m.Entities[len(m.Entities)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeqID", wireType)
			}
			m.SeqID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SeqID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkFlow(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNetworkFlow
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
func skipNetworkFlow(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNetworkFlow
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
					return 0, ErrIntOverflowNetworkFlow
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
					return 0, ErrIntOverflowNetworkFlow
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
				return 0, ErrInvalidLengthNetworkFlow
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNetworkFlow
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNetworkFlow
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNetworkFlow        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetworkFlow          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNetworkFlow = fmt.Errorf("proto: unexpected end of group")
)
