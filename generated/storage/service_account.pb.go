// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/service_account.proto

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

// Any properties of an individual service account.
// (regardless of time, scope, or context)
// ////////////////////////////////////////
type ServiceAccount struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,type(uuid)"`                                                                                                           // @gotags: sql:"pk,type(uuid)"
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" search:"Service Account,store"`                                                                                                       // @gotags: search:"Service Account,store"
	Namespace            string            `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty" search:"Namespace,store"`                                                                                             // @gotags: search:"Namespace,store"
	ClusterName          string            `protobuf:"bytes,4,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty" search:"Cluster,store"`                                                                      // @gotags: search:"Cluster,store"
	ClusterId            string            `protobuf:"bytes,5,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" search:"Cluster ID,store,hidden" sql:"type(uuid)"`                                                                            // @gotags: search:"Cluster ID,store,hidden" sql:"type(uuid)"
	Labels               map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" search:"Service Account Label"`           // @gotags: search:"Service Account Label"
	Annotations          map[string]string `protobuf:"bytes,7,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" search:"Service Account Annotation"` // @gotags: search:"Service Account Annotation"
	CreatedAt            *types.Timestamp  `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	AutomountToken       bool              `protobuf:"varint,9,opt,name=automount_token,json=automountToken,proto3" json:"automount_token,omitempty"`
	Secrets              []string          `protobuf:"bytes,10,rep,name=secrets,proto3" json:"secrets,omitempty"`
	ImagePullSecrets     []string          `protobuf:"bytes,11,rep,name=image_pull_secrets,json=imagePullSecrets,proto3" json:"image_pull_secrets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ServiceAccount) Reset()         { *m = ServiceAccount{} }
func (m *ServiceAccount) String() string { return proto.CompactTextString(m) }
func (*ServiceAccount) ProtoMessage()    {}
func (*ServiceAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a2dd3acd1a2d3c1, []int{0}
}
func (m *ServiceAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ServiceAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ServiceAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceAccount.Merge(m, src)
}
func (m *ServiceAccount) XXX_Size() int {
	return m.Size()
}
func (m *ServiceAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceAccount proto.InternalMessageInfo

func (m *ServiceAccount) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceAccount) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ServiceAccount) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ServiceAccount) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *ServiceAccount) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ServiceAccount) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *ServiceAccount) GetCreatedAt() *types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ServiceAccount) GetAutomountToken() bool {
	if m != nil {
		return m.AutomountToken
	}
	return false
}

func (m *ServiceAccount) GetSecrets() []string {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *ServiceAccount) GetImagePullSecrets() []string {
	if m != nil {
		return m.ImagePullSecrets
	}
	return nil
}

func (m *ServiceAccount) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ServiceAccount) Clone() *ServiceAccount {
	if m == nil {
		return nil
	}
	cloned := new(ServiceAccount)
	*cloned = *m

	if m.Labels != nil {
		cloned.Labels = make(map[string]string, len(m.Labels))
		for k, v := range m.Labels {
			cloned.Labels[k] = v
		}
	}
	if m.Annotations != nil {
		cloned.Annotations = make(map[string]string, len(m.Annotations))
		for k, v := range m.Annotations {
			cloned.Annotations[k] = v
		}
	}
	cloned.CreatedAt = m.CreatedAt.Clone()
	if m.Secrets != nil {
		cloned.Secrets = make([]string, len(m.Secrets))
		copy(cloned.Secrets, m.Secrets)
	}
	if m.ImagePullSecrets != nil {
		cloned.ImagePullSecrets = make([]string, len(m.ImagePullSecrets))
		copy(cloned.ImagePullSecrets, m.ImagePullSecrets)
	}
	return cloned
}

func init() {
	proto.RegisterType((*ServiceAccount)(nil), "storage.ServiceAccount")
	proto.RegisterMapType((map[string]string)(nil), "storage.ServiceAccount.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "storage.ServiceAccount.LabelsEntry")
}

func init() { proto.RegisterFile("storage/service_account.proto", fileDescriptor_7a2dd3acd1a2d3c1) }

var fileDescriptor_7a2dd3acd1a2d3c1 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0x92, 0x26, 0xf5, 0x18, 0x85, 0xb0, 0xe2, 0xb0, 0x44, 0x34, 0x18, 0x38, 0xe0,
	0x03, 0xda, 0x48, 0xe5, 0x42, 0xa9, 0x54, 0x29, 0x48, 0x1c, 0x40, 0x08, 0x21, 0xb7, 0x27, 0x2e,
	0xd6, 0xc6, 0x1e, 0x22, 0x2b, 0x6b, 0x6f, 0xe4, 0x1d, 0x57, 0xf4, 0x4d, 0x78, 0x20, 0x0e, 0x1c,
	0x79, 0x04, 0x14, 0x5e, 0x04, 0x79, 0x77, 0x4d, 0x0b, 0x12, 0x07, 0x4e, 0x3b, 0xfb, 0xcf, 0x37,
	0xbf, 0xbd, 0xff, 0xc0, 0x91, 0x21, 0xdd, 0xc8, 0x0d, 0x2e, 0x0d, 0x36, 0x97, 0x65, 0x8e, 0x99,
	0xcc, 0x73, 0xdd, 0xd6, 0x24, 0x76, 0x8d, 0x26, 0xcd, 0x26, 0xbe, 0x3d, 0x7f, 0xb8, 0xd1, 0x7a,
	0xa3, 0x70, 0x69, 0xe5, 0x75, 0xfb, 0x69, 0x49, 0x65, 0x85, 0x86, 0x64, 0xb5, 0x73, 0xe4, 0xe3,
	0xaf, 0x23, 0x98, 0x9e, 0x3b, 0x8f, 0x95, 0xb3, 0x60, 0x53, 0x18, 0x94, 0x05, 0x0f, 0xe2, 0x20,
	0x09, 0xd3, 0x41, 0x59, 0x30, 0x06, 0xa3, 0x5a, 0x56, 0xc8, 0x07, 0x56, 0xb1, 0x35, 0x7b, 0x00,
	0x61, 0x77, 0x9a, 0x9d, 0xcc, 0x91, 0x0f, 0x6d, 0xe3, 0x5a, 0x60, 0x8f, 0xe0, 0x76, 0xae, 0x5a,
	0x43, 0xd8, 0x64, 0x76, 0x72, 0x64, 0x81, 0xc8, 0x6b, 0xef, 0x3b, 0x83, 0x23, 0x80, 0x1e, 0x29,
	0x0b, 0x7e, 0xe0, 0x1c, 0xbc, 0xf2, 0xa6, 0x60, 0xa7, 0x30, 0x56, 0x72, 0x8d, 0xca, 0xf0, 0x71,
	0x3c, 0x4c, 0xa2, 0xe3, 0x27, 0xc2, 0xbf, 0x48, 0xfc, 0xf9, 0xb3, 0xe2, 0x9d, 0xa5, 0x5e, 0xd7,
	0xd4, 0x5c, 0xa5, 0x7e, 0x84, 0xbd, 0x85, 0x48, 0xd6, 0xb5, 0x26, 0x49, 0xa5, 0xae, 0x0d, 0x9f,
	0x58, 0x87, 0xe4, 0x5f, 0x0e, 0xab, 0x6b, 0xd4, 0xd9, 0xdc, 0x1c, 0x66, 0x27, 0x00, 0x79, 0x83,
	0x92, 0xb0, 0xc8, 0x24, 0xf1, 0xc3, 0x38, 0x48, 0xa2, 0xe3, 0xb9, 0x70, 0xa9, 0x8a, 0x3e, 0x55,
	0x71, 0xd1, 0xa7, 0x9a, 0x86, 0x9e, 0x5e, 0x11, 0x7b, 0x0a, 0x77, 0x64, 0x4b, 0xba, 0xea, 0xbe,
	0x92, 0x91, 0xde, 0x62, 0xcd, 0xc3, 0x38, 0x48, 0x0e, 0xd3, 0xe9, 0x6f, 0xf9, 0xa2, 0x53, 0x19,
	0x87, 0x89, 0xc1, 0xbc, 0x41, 0x32, 0x1c, 0xe2, 0x61, 0x12, 0xa6, 0xfd, 0x95, 0x3d, 0x03, 0x56,
	0x56, 0x72, 0x83, 0xd9, 0xae, 0x55, 0x2a, 0xeb, 0xa1, 0xc8, 0x42, 0x33, 0xdb, 0xf9, 0xd0, 0x2a,
	0x75, 0xee, 0xf4, 0xf9, 0x09, 0x44, 0x37, 0xe2, 0x60, 0x33, 0x18, 0x6e, 0xf1, 0xca, 0x2f, 0xb2,
	0x2b, 0xd9, 0x3d, 0x38, 0xb8, 0x94, 0xaa, 0xed, 0x57, 0xe9, 0x2e, 0x2f, 0x07, 0x2f, 0x82, 0xf9,
	0x19, 0xcc, 0xfe, 0xce, 0xe1, 0x7f, 0xe6, 0x5f, 0x9d, 0x7d, 0xdb, 0x2f, 0x82, 0xef, 0xfb, 0x45,
	0xf0, 0x63, 0xbf, 0x08, 0xbe, 0xfc, 0x5c, 0xdc, 0x82, 0xfb, 0xa5, 0x16, 0x86, 0x64, 0xbe, 0x6d,
	0xf4, 0x67, 0x17, 0x54, 0xbf, 0x80, 0x8f, 0x77, 0xc5, 0xd2, 0x97, 0xa7, 0xfe, 0x5c, 0x8f, 0x2d,
	0xf1, 0xfc, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x64, 0x7c, 0x79, 0x5a, 0xd8, 0x02, 0x00, 0x00,
}

func (m *ServiceAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServiceAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ImagePullSecrets) > 0 {
		for iNdEx := len(m.ImagePullSecrets) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ImagePullSecrets[iNdEx])
			copy(dAtA[i:], m.ImagePullSecrets[iNdEx])
			i = encodeVarintServiceAccount(dAtA, i, uint64(len(m.ImagePullSecrets[iNdEx])))
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.Secrets) > 0 {
		for iNdEx := len(m.Secrets) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Secrets[iNdEx])
			copy(dAtA[i:], m.Secrets[iNdEx])
			i = encodeVarintServiceAccount(dAtA, i, uint64(len(m.Secrets[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if m.AutomountToken {
		i--
		if m.AutomountToken {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.CreatedAt != nil {
		{
			size, err := m.CreatedAt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintServiceAccount(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.Annotations) > 0 {
		for k := range m.Annotations {
			v := m.Annotations[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintServiceAccount(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintServiceAccount(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintServiceAccount(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Labels) > 0 {
		for k := range m.Labels {
			v := m.Labels[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintServiceAccount(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintServiceAccount(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintServiceAccount(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ClusterId) > 0 {
		i -= len(m.ClusterId)
		copy(dAtA[i:], m.ClusterId)
		i = encodeVarintServiceAccount(dAtA, i, uint64(len(m.ClusterId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ClusterName) > 0 {
		i -= len(m.ClusterName)
		copy(dAtA[i:], m.ClusterName)
		i = encodeVarintServiceAccount(dAtA, i, uint64(len(m.ClusterName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintServiceAccount(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintServiceAccount(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintServiceAccount(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintServiceAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovServiceAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ServiceAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovServiceAccount(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovServiceAccount(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovServiceAccount(uint64(l))
	}
	l = len(m.ClusterName)
	if l > 0 {
		n += 1 + l + sovServiceAccount(uint64(l))
	}
	l = len(m.ClusterId)
	if l > 0 {
		n += 1 + l + sovServiceAccount(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovServiceAccount(uint64(len(k))) + 1 + len(v) + sovServiceAccount(uint64(len(v)))
			n += mapEntrySize + 1 + sovServiceAccount(uint64(mapEntrySize))
		}
	}
	if len(m.Annotations) > 0 {
		for k, v := range m.Annotations {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovServiceAccount(uint64(len(k))) + 1 + len(v) + sovServiceAccount(uint64(len(v)))
			n += mapEntrySize + 1 + sovServiceAccount(uint64(mapEntrySize))
		}
	}
	if m.CreatedAt != nil {
		l = m.CreatedAt.Size()
		n += 1 + l + sovServiceAccount(uint64(l))
	}
	if m.AutomountToken {
		n += 2
	}
	if len(m.Secrets) > 0 {
		for _, s := range m.Secrets {
			l = len(s)
			n += 1 + l + sovServiceAccount(uint64(l))
		}
	}
	if len(m.ImagePullSecrets) > 0 {
		for _, s := range m.ImagePullSecrets {
			l = len(s)
			n += 1 + l + sovServiceAccount(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovServiceAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozServiceAccount(x uint64) (n int) {
	return sovServiceAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ServiceAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceAccount
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
			return fmt.Errorf("proto: ServiceAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceAccount
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
				return ErrInvalidLengthServiceAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceAccount
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
					return ErrIntOverflowServiceAccount
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
				return ErrInvalidLengthServiceAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceAccount
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
				return ErrInvalidLengthServiceAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceAccount
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
				return ErrInvalidLengthServiceAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceAccount
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
				return ErrInvalidLengthServiceAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceAccount
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
				return ErrInvalidLengthServiceAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthServiceAccount
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
						return ErrIntOverflowServiceAccount
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
							return ErrIntOverflowServiceAccount
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
						return ErrInvalidLengthServiceAccount
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthServiceAccount
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
							return ErrIntOverflowServiceAccount
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
						return ErrInvalidLengthServiceAccount
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthServiceAccount
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipServiceAccount(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthServiceAccount
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Annotations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceAccount
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
				return ErrInvalidLengthServiceAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthServiceAccount
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
						return ErrIntOverflowServiceAccount
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
							return ErrIntOverflowServiceAccount
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
						return ErrInvalidLengthServiceAccount
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthServiceAccount
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
							return ErrIntOverflowServiceAccount
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
						return ErrInvalidLengthServiceAccount
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthServiceAccount
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipServiceAccount(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthServiceAccount
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Annotations[mapkey] = mapvalue
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceAccount
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
				return ErrInvalidLengthServiceAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthServiceAccount
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
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AutomountToken", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceAccount
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
			m.AutomountToken = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secrets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceAccount
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
				return ErrInvalidLengthServiceAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secrets = append(m.Secrets, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImagePullSecrets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceAccount
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
				return ErrInvalidLengthServiceAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImagePullSecrets = append(m.ImagePullSecrets, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServiceAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthServiceAccount
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
func skipServiceAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowServiceAccount
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
					return 0, ErrIntOverflowServiceAccount
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
					return 0, ErrIntOverflowServiceAccount
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
				return 0, ErrInvalidLengthServiceAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupServiceAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthServiceAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthServiceAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowServiceAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupServiceAccount = fmt.Errorf("proto: unexpected end of group")
)
