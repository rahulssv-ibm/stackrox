// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/namespace_metadata.proto

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

type NamespaceMetadata struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" search:"Namespace ID" sql:"pk,type(uuid)"`                                                                                                 // @gotags: search:"Namespace ID" sql:"pk,type(uuid)"
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" search:"Namespace,store"`                                                                                             // @gotags: search:"Namespace,store"
	ClusterId            string            `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" search:"Cluster ID,hidden,store" sql:"fk(Cluster:id),no-fk-constraint,type(uuid)"`                                                                  // @gotags: search:"Cluster ID,hidden,store" sql:"fk(Cluster:id),no-fk-constraint,type(uuid)"
	ClusterName          string            `protobuf:"bytes,4,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty" search:"Cluster"`                                                            // @gotags: search:"Cluster"
	Labels               map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" search:"Namespace Label"` // @gotags: search:"Namespace Label"
	CreationTime         *types.Timestamp  `protobuf:"bytes,6,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	Priority             int64             `protobuf:"varint,7,opt,name=priority,proto3" json:"priority,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,8,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" search:"Namespace Annotation"` // @gotags: search:"Namespace Annotation"
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NamespaceMetadata) Reset()         { *m = NamespaceMetadata{} }
func (m *NamespaceMetadata) String() string { return proto.CompactTextString(m) }
func (*NamespaceMetadata) ProtoMessage()    {}
func (*NamespaceMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cbfd70053cb23bf, []int{0}
}
func (m *NamespaceMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NamespaceMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NamespaceMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NamespaceMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespaceMetadata.Merge(m, src)
}
func (m *NamespaceMetadata) XXX_Size() int {
	return m.Size()
}
func (m *NamespaceMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespaceMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_NamespaceMetadata proto.InternalMessageInfo

func (m *NamespaceMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NamespaceMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamespaceMetadata) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *NamespaceMetadata) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *NamespaceMetadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *NamespaceMetadata) GetCreationTime() *types.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *NamespaceMetadata) GetPriority() int64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *NamespaceMetadata) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *NamespaceMetadata) MessageClone() proto.Message {
	return m.Clone()
}
func (m *NamespaceMetadata) Clone() *NamespaceMetadata {
	if m == nil {
		return nil
	}
	cloned := new(NamespaceMetadata)
	*cloned = *m

	if m.Labels != nil {
		cloned.Labels = make(map[string]string, len(m.Labels))
		for k, v := range m.Labels {
			cloned.Labels[k] = v
		}
	}
	cloned.CreationTime = m.CreationTime.Clone()
	if m.Annotations != nil {
		cloned.Annotations = make(map[string]string, len(m.Annotations))
		for k, v := range m.Annotations {
			cloned.Annotations[k] = v
		}
	}
	return cloned
}

func init() {
	proto.RegisterType((*NamespaceMetadata)(nil), "storage.NamespaceMetadata")
	proto.RegisterMapType((map[string]string)(nil), "storage.NamespaceMetadata.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "storage.NamespaceMetadata.LabelsEntry")
}

func init() { proto.RegisterFile("storage/namespace_metadata.proto", fileDescriptor_5cbfd70053cb23bf) }

var fileDescriptor_5cbfd70053cb23bf = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x50, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xed, 0x4a, 0xfe, 0x1c, 0xb9, 0xc5, 0x5e, 0x7a, 0x50, 0x05, 0x55, 0xd5, 0x1e, 0x8a, 0xa0,
	0xb0, 0x02, 0xf7, 0xd2, 0x0f, 0x70, 0x69, 0x20, 0x87, 0x40, 0x9c, 0x83, 0xc8, 0x29, 0x17, 0xb3,
	0x96, 0x36, 0x66, 0xb1, 0xa4, 0x15, 0xab, 0x75, 0x88, 0xff, 0x49, 0x7e, 0x52, 0x8e, 0xf9, 0x09,
	0xc6, 0xf9, 0x23, 0x41, 0xab, 0x55, 0x62, 0x12, 0x08, 0xe4, 0xb4, 0xb3, 0x6f, 0xde, 0x9b, 0x79,
	0x6f, 0x20, 0xa8, 0x94, 0x90, 0x74, 0xc5, 0xa2, 0x82, 0xe6, 0xac, 0x2a, 0x69, 0xc2, 0x16, 0x39,
	0x53, 0x34, 0xa5, 0x8a, 0x92, 0x52, 0x0a, 0x25, 0x70, 0xdf, 0x30, 0xbc, 0x2f, 0x2b, 0x21, 0x56,
	0x19, 0x8b, 0x34, 0xbc, 0xdc, 0x5c, 0x46, 0x8a, 0xe7, 0xac, 0x52, 0x34, 0x2f, 0x1b, 0xe6, 0xb7,
	0x9d, 0x0d, 0x93, 0xb3, 0x76, 0xcc, 0xdc, 0x4c, 0xc1, 0x1f, 0xc0, 0xe2, 0xa9, 0x8b, 0x02, 0x14,
	0x0e, 0x63, 0x8b, 0xa7, 0x18, 0x43, 0xa7, 0xde, 0xe5, 0x5a, 0x1a, 0xd1, 0x35, 0xfe, 0x0c, 0x90,
	0x64, 0x9b, 0x4a, 0x31, 0xb9, 0xe0, 0xa9, 0x6b, 0xeb, 0xce, 0xd0, 0x20, 0x27, 0x29, 0xfe, 0x0a,
	0xa3, 0xb6, 0xad, 0xa5, 0x1d, 0x4d, 0x70, 0x0c, 0x56, 0xaf, 0xc4, 0x33, 0xe8, 0x65, 0x74, 0xc9,
	0xb2, 0xca, 0xed, 0x06, 0x76, 0xe8, 0x4c, 0xbf, 0x13, 0x63, 0x9b, 0xbc, 0x70, 0x44, 0x4e, 0x35,
	0xf1, 0xb8, 0x50, 0x72, 0x1b, 0x1b, 0x15, 0xfe, 0x07, 0xef, 0x13, 0xc9, 0xa8, 0xe2, 0xa2, 0x58,
	0xd4, 0xb9, 0xdc, 0x5e, 0x80, 0x42, 0x67, 0xea, 0x91, 0x26, 0x34, 0x69, 0x43, 0x93, 0xf3, 0x36,
	0x74, 0x3c, 0x6a, 0x05, 0x35, 0x84, 0x3d, 0x18, 0x94, 0x92, 0x0b, 0xc9, 0xd5, 0xd6, 0xed, 0x07,
	0x28, 0xb4, 0xe3, 0xc7, 0x3f, 0x9e, 0x83, 0x43, 0x8b, 0x42, 0x28, 0xcd, 0xae, 0xdc, 0x81, 0x76,
	0xf8, 0xe3, 0x15, 0x87, 0xff, 0x9f, 0xd8, 0x8d, 0xcd, 0x43, 0xbd, 0xf7, 0x1b, 0x9c, 0x83, 0x08,
	0x78, 0x0c, 0xf6, 0x9a, 0x6d, 0xcd, 0x85, 0xeb, 0x12, 0x7f, 0x84, 0xee, 0x15, 0xcd, 0x36, 0xed,
	0x8d, 0x9b, 0xcf, 0x1f, 0xeb, 0x17, 0xf2, 0x66, 0x30, 0x7e, 0x3e, 0xfb, 0x2d, 0xfa, 0xa3, 0xd9,
	0xed, 0xde, 0x47, 0x77, 0x7b, 0x1f, 0xed, 0xf6, 0x3e, 0xba, 0xb9, 0xf7, 0xdf, 0xc1, 0x27, 0x2e,
	0x48, 0xa5, 0x68, 0xb2, 0x96, 0xe2, 0xba, 0xb9, 0x52, 0x9b, 0xeb, 0x62, 0x42, 0x22, 0x53, 0xfe,
	0x35, 0xef, 0xb2, 0xa7, 0x19, 0x3f, 0x1f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x04, 0x36, 0xbd,
	0x77, 0x02, 0x00, 0x00,
}

func (m *NamespaceMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NamespaceMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Annotations) > 0 {
		for k := range m.Annotations {
			v := m.Annotations[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintNamespaceMetadata(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintNamespaceMetadata(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintNamespaceMetadata(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x42
		}
	}
	if m.Priority != 0 {
		i = encodeVarintNamespaceMetadata(dAtA, i, uint64(m.Priority))
		i--
		dAtA[i] = 0x38
	}
	if m.CreationTime != nil {
		{
			size, err := m.CreationTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNamespaceMetadata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.Labels) > 0 {
		for k := range m.Labels {
			v := m.Labels[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintNamespaceMetadata(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintNamespaceMetadata(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintNamespaceMetadata(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ClusterName) > 0 {
		i -= len(m.ClusterName)
		copy(dAtA[i:], m.ClusterName)
		i = encodeVarintNamespaceMetadata(dAtA, i, uint64(len(m.ClusterName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ClusterId) > 0 {
		i -= len(m.ClusterId)
		copy(dAtA[i:], m.ClusterId)
		i = encodeVarintNamespaceMetadata(dAtA, i, uint64(len(m.ClusterId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintNamespaceMetadata(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintNamespaceMetadata(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNamespaceMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovNamespaceMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NamespaceMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovNamespaceMetadata(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNamespaceMetadata(uint64(l))
	}
	l = len(m.ClusterId)
	if l > 0 {
		n += 1 + l + sovNamespaceMetadata(uint64(l))
	}
	l = len(m.ClusterName)
	if l > 0 {
		n += 1 + l + sovNamespaceMetadata(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovNamespaceMetadata(uint64(len(k))) + 1 + len(v) + sovNamespaceMetadata(uint64(len(v)))
			n += mapEntrySize + 1 + sovNamespaceMetadata(uint64(mapEntrySize))
		}
	}
	if m.CreationTime != nil {
		l = m.CreationTime.Size()
		n += 1 + l + sovNamespaceMetadata(uint64(l))
	}
	if m.Priority != 0 {
		n += 1 + sovNamespaceMetadata(uint64(m.Priority))
	}
	if len(m.Annotations) > 0 {
		for k, v := range m.Annotations {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovNamespaceMetadata(uint64(len(k))) + 1 + len(v) + sovNamespaceMetadata(uint64(len(v)))
			n += mapEntrySize + 1 + sovNamespaceMetadata(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNamespaceMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNamespaceMetadata(x uint64) (n int) {
	return sovNamespaceMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NamespaceMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespaceMetadata
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
			return fmt.Errorf("proto: NamespaceMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespaceMetadata
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
				return ErrInvalidLengthNamespaceMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNamespaceMetadata
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
					return ErrIntOverflowNamespaceMetadata
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
				return ErrInvalidLengthNamespaceMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNamespaceMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespaceMetadata
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
				return ErrInvalidLengthNamespaceMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNamespaceMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespaceMetadata
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
				return ErrInvalidLengthNamespaceMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNamespaceMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespaceMetadata
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
				return ErrInvalidLengthNamespaceMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNamespaceMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNamespaceMetadata
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNamespaceMetadata
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthNamespaceMetadata
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthNamespaceMetadata
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNamespaceMetadata
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthNamespaceMetadata
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthNamespaceMetadata
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipNamespaceMetadata(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthNamespaceMetadata
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespaceMetadata
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
				return ErrInvalidLengthNamespaceMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNamespaceMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreationTime == nil {
				m.CreationTime = &types.Timestamp{}
			}
			if err := m.CreationTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespaceMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Annotations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespaceMetadata
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
				return ErrInvalidLengthNamespaceMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNamespaceMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Annotations == nil {
				m.Annotations = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNamespaceMetadata
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNamespaceMetadata
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthNamespaceMetadata
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthNamespaceMetadata
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNamespaceMetadata
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthNamespaceMetadata
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthNamespaceMetadata
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipNamespaceMetadata(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthNamespaceMetadata
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Annotations[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespaceMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNamespaceMetadata
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
func skipNamespaceMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNamespaceMetadata
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
					return 0, ErrIntOverflowNamespaceMetadata
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
					return 0, ErrIntOverflowNamespaceMetadata
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
				return 0, ErrInvalidLengthNamespaceMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNamespaceMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNamespaceMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNamespaceMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNamespaceMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNamespaceMetadata = fmt.Errorf("proto: unexpected end of group")
)
