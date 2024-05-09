// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internalapi/central/cluster_status.proto

package central

import (
	fmt "fmt"
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

type DeploymentEnvironmentUpdate struct {
	Environments         []string `protobuf:"bytes,1,rep,name=environments,proto3" json:"environments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeploymentEnvironmentUpdate) Reset()         { *m = DeploymentEnvironmentUpdate{} }
func (m *DeploymentEnvironmentUpdate) String() string { return proto.CompactTextString(m) }
func (*DeploymentEnvironmentUpdate) ProtoMessage()    {}
func (*DeploymentEnvironmentUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc7c40b05e2d7c6f, []int{0}
}
func (m *DeploymentEnvironmentUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeploymentEnvironmentUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeploymentEnvironmentUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeploymentEnvironmentUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentEnvironmentUpdate.Merge(m, src)
}
func (m *DeploymentEnvironmentUpdate) XXX_Size() int {
	return m.Size()
}
func (m *DeploymentEnvironmentUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentEnvironmentUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentEnvironmentUpdate proto.InternalMessageInfo

func (m *DeploymentEnvironmentUpdate) GetEnvironments() []string {
	if m != nil {
		return m.Environments
	}
	return nil
}

func (m *DeploymentEnvironmentUpdate) MessageClone() proto.Message {
	return m.Clone()
}
func (m *DeploymentEnvironmentUpdate) Clone() *DeploymentEnvironmentUpdate {
	if m == nil {
		return nil
	}
	cloned := new(DeploymentEnvironmentUpdate)
	*cloned = *m

	if m.Environments != nil {
		cloned.Environments = make([]string, len(m.Environments))
		copy(cloned.Environments, m.Environments)
	}
	return cloned
}

type ClusterStatusUpdate struct {
	// Making it a oneof for future proofing.
	//
	// Types that are valid to be assigned to Msg:
	//	*ClusterStatusUpdate_Status
	//	*ClusterStatusUpdate_DeploymentEnvUpdate
	Msg                  isClusterStatusUpdate_Msg `protobuf_oneof:"msg"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ClusterStatusUpdate) Reset()         { *m = ClusterStatusUpdate{} }
func (m *ClusterStatusUpdate) String() string { return proto.CompactTextString(m) }
func (*ClusterStatusUpdate) ProtoMessage()    {}
func (*ClusterStatusUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc7c40b05e2d7c6f, []int{1}
}
func (m *ClusterStatusUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterStatusUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterStatusUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterStatusUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterStatusUpdate.Merge(m, src)
}
func (m *ClusterStatusUpdate) XXX_Size() int {
	return m.Size()
}
func (m *ClusterStatusUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterStatusUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterStatusUpdate proto.InternalMessageInfo

type isClusterStatusUpdate_Msg interface {
	isClusterStatusUpdate_Msg()
	MarshalTo([]byte) (int, error)
	Size() int
	Clone() isClusterStatusUpdate_Msg
}

type ClusterStatusUpdate_Status struct {
	Status *storage.ClusterStatus `protobuf:"bytes,1,opt,name=status,proto3,oneof" json:"status,omitempty"`
}
type ClusterStatusUpdate_DeploymentEnvUpdate struct {
	DeploymentEnvUpdate *DeploymentEnvironmentUpdate `protobuf:"bytes,2,opt,name=deployment_env_update,json=deploymentEnvUpdate,proto3,oneof" json:"deployment_env_update,omitempty"`
}

func (*ClusterStatusUpdate_Status) isClusterStatusUpdate_Msg() {}
func (m *ClusterStatusUpdate_Status) Clone() isClusterStatusUpdate_Msg {
	if m == nil {
		return nil
	}
	cloned := new(ClusterStatusUpdate_Status)
	*cloned = *m

	cloned.Status = m.Status.Clone()
	return cloned
}
func (*ClusterStatusUpdate_DeploymentEnvUpdate) isClusterStatusUpdate_Msg() {}
func (m *ClusterStatusUpdate_DeploymentEnvUpdate) Clone() isClusterStatusUpdate_Msg {
	if m == nil {
		return nil
	}
	cloned := new(ClusterStatusUpdate_DeploymentEnvUpdate)
	*cloned = *m

	cloned.DeploymentEnvUpdate = m.DeploymentEnvUpdate.Clone()
	return cloned
}

func (m *ClusterStatusUpdate) GetMsg() isClusterStatusUpdate_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *ClusterStatusUpdate) GetStatus() *storage.ClusterStatus {
	if x, ok := m.GetMsg().(*ClusterStatusUpdate_Status); ok {
		return x.Status
	}
	return nil
}

func (m *ClusterStatusUpdate) GetDeploymentEnvUpdate() *DeploymentEnvironmentUpdate {
	if x, ok := m.GetMsg().(*ClusterStatusUpdate_DeploymentEnvUpdate); ok {
		return x.DeploymentEnvUpdate
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ClusterStatusUpdate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ClusterStatusUpdate_Status)(nil),
		(*ClusterStatusUpdate_DeploymentEnvUpdate)(nil),
	}
}

func (m *ClusterStatusUpdate) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ClusterStatusUpdate) Clone() *ClusterStatusUpdate {
	if m == nil {
		return nil
	}
	cloned := new(ClusterStatusUpdate)
	*cloned = *m

	if m.Msg != nil {
		cloned.Msg = m.Msg.Clone()
	}
	return cloned
}

type RawClusterHealthInfo struct {
	CollectorHealthInfo        *storage.CollectorHealthInfo        `protobuf:"bytes,1,opt,name=collector_health_info,json=collectorHealthInfo,proto3" json:"collector_health_info,omitempty"`
	AdmissionControlHealthInfo *storage.AdmissionControlHealthInfo `protobuf:"bytes,2,opt,name=admission_control_health_info,json=admissionControlHealthInfo,proto3" json:"admission_control_health_info,omitempty"`
	ScannerHealthInfo          *storage.ScannerHealthInfo          `protobuf:"bytes,3,opt,name=scanner_health_info,json=scannerHealthInfo,proto3" json:"scanner_health_info,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                            `json:"-"`
	XXX_unrecognized           []byte                              `json:"-"`
	XXX_sizecache              int32                               `json:"-"`
}

func (m *RawClusterHealthInfo) Reset()         { *m = RawClusterHealthInfo{} }
func (m *RawClusterHealthInfo) String() string { return proto.CompactTextString(m) }
func (*RawClusterHealthInfo) ProtoMessage()    {}
func (*RawClusterHealthInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc7c40b05e2d7c6f, []int{2}
}
func (m *RawClusterHealthInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RawClusterHealthInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RawClusterHealthInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RawClusterHealthInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawClusterHealthInfo.Merge(m, src)
}
func (m *RawClusterHealthInfo) XXX_Size() int {
	return m.Size()
}
func (m *RawClusterHealthInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RawClusterHealthInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RawClusterHealthInfo proto.InternalMessageInfo

func (m *RawClusterHealthInfo) GetCollectorHealthInfo() *storage.CollectorHealthInfo {
	if m != nil {
		return m.CollectorHealthInfo
	}
	return nil
}

func (m *RawClusterHealthInfo) GetAdmissionControlHealthInfo() *storage.AdmissionControlHealthInfo {
	if m != nil {
		return m.AdmissionControlHealthInfo
	}
	return nil
}

func (m *RawClusterHealthInfo) GetScannerHealthInfo() *storage.ScannerHealthInfo {
	if m != nil {
		return m.ScannerHealthInfo
	}
	return nil
}

func (m *RawClusterHealthInfo) MessageClone() proto.Message {
	return m.Clone()
}
func (m *RawClusterHealthInfo) Clone() *RawClusterHealthInfo {
	if m == nil {
		return nil
	}
	cloned := new(RawClusterHealthInfo)
	*cloned = *m

	cloned.CollectorHealthInfo = m.CollectorHealthInfo.Clone()
	cloned.AdmissionControlHealthInfo = m.AdmissionControlHealthInfo.Clone()
	cloned.ScannerHealthInfo = m.ScannerHealthInfo.Clone()
	return cloned
}

type ClusterHealthResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterHealthResponse) Reset()         { *m = ClusterHealthResponse{} }
func (m *ClusterHealthResponse) String() string { return proto.CompactTextString(m) }
func (*ClusterHealthResponse) ProtoMessage()    {}
func (*ClusterHealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc7c40b05e2d7c6f, []int{3}
}
func (m *ClusterHealthResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterHealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterHealthResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterHealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterHealthResponse.Merge(m, src)
}
func (m *ClusterHealthResponse) XXX_Size() int {
	return m.Size()
}
func (m *ClusterHealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterHealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterHealthResponse proto.InternalMessageInfo

func (m *ClusterHealthResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ClusterHealthResponse) Clone() *ClusterHealthResponse {
	if m == nil {
		return nil
	}
	cloned := new(ClusterHealthResponse)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterType((*DeploymentEnvironmentUpdate)(nil), "central.DeploymentEnvironmentUpdate")
	proto.RegisterType((*ClusterStatusUpdate)(nil), "central.ClusterStatusUpdate")
	proto.RegisterType((*RawClusterHealthInfo)(nil), "central.RawClusterHealthInfo")
	proto.RegisterType((*ClusterHealthResponse)(nil), "central.ClusterHealthResponse")
}

func init() {
	proto.RegisterFile("internalapi/central/cluster_status.proto", fileDescriptor_fc7c40b05e2d7c6f)
}

var fileDescriptor_fc7c40b05e2d7c6f = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x3d, 0x4e, 0xc3, 0x30,
	0x18, 0x86, 0x9b, 0x56, 0x14, 0x61, 0x58, 0x48, 0x08, 0x54, 0x85, 0x46, 0x55, 0x60, 0xe8, 0x94,
	0x20, 0x18, 0x99, 0xda, 0x82, 0x54, 0x98, 0x50, 0x2a, 0x96, 0x2e, 0x91, 0x49, 0xbe, 0xb6, 0x91,
	0xdc, 0xcf, 0x91, 0xed, 0x16, 0x71, 0x13, 0x06, 0x4e, 0xc0, 0x49, 0x18, 0x39, 0x02, 0x2a, 0x17,
	0x41, 0x24, 0xee, 0x4f, 0xf8, 0x9b, 0x12, 0xbd, 0xdf, 0xeb, 0xc7, 0x8f, 0x2d, 0x93, 0x56, 0x82,
	0x0a, 0x04, 0x52, 0x46, 0xd3, 0xc4, 0x8f, 0x00, 0x95, 0xa0, 0xcc, 0x8f, 0xd8, 0x54, 0x2a, 0x10,
	0xa1, 0x54, 0x54, 0x4d, 0xa5, 0x97, 0x0a, 0xae, 0xb8, 0xb9, 0xa9, 0xa7, 0x75, 0x5b, 0x2a, 0x2e,
	0xe8, 0x08, 0x16, 0xb5, 0x7c, 0xee, 0xb6, 0xc9, 0xe1, 0x25, 0xa4, 0x8c, 0x3f, 0x4e, 0x00, 0xd5,
	0x15, 0xce, 0x12, 0xc1, 0xf1, 0xeb, 0xf7, 0x2e, 0x8d, 0xa9, 0x02, 0xd3, 0x25, 0x3b, 0xb0, 0x0a,
	0x65, 0xcd, 0x68, 0x56, 0x5a, 0x5b, 0x41, 0x21, 0x73, 0x5f, 0x0c, 0x62, 0x75, 0x73, 0x68, 0x3f,
	0xdb, 0x5a, 0xaf, 0x3d, 0x25, 0xd5, 0x5c, 0xa5, 0x66, 0x34, 0x8d, 0xd6, 0xf6, 0xd9, 0xbe, 0xa7,
	0x15, 0xbc, 0x42, 0xbb, 0x57, 0x0a, 0x74, 0xcf, 0x1c, 0x10, 0x3b, 0x5e, 0xca, 0x84, 0x80, 0xb3,
	0x70, 0x9a, 0xa1, 0x6a, 0xe5, 0x0c, 0x70, 0xe2, 0xe9, 0xc3, 0x78, 0xff, 0x28, 0xf7, 0x4a, 0x81,
	0x15, 0xaf, 0x8f, 0xf3, 0xb8, 0xb3, 0x41, 0x2a, 0x13, 0x39, 0x72, 0x9f, 0xcb, 0x64, 0x2f, 0xa0,
	0x0f, 0xda, 0xa0, 0x07, 0x94, 0xa9, 0xf1, 0x35, 0x0e, 0xb9, 0x79, 0x4b, 0xec, 0x88, 0x33, 0x06,
	0x91, 0xe2, 0x22, 0x1c, 0x67, 0x79, 0x98, 0xe0, 0x90, 0x6b, 0xf9, 0xa3, 0x95, 0xfc, 0xa2, 0xb5,
	0x5a, 0x1c, 0x58, 0xd1, 0xcf, 0xd0, 0x1c, 0x92, 0x06, 0x8d, 0x27, 0x89, 0x94, 0x09, 0xc7, 0x30,
	0xe2, 0xa8, 0x04, 0x67, 0x05, 0x72, 0x7e, 0xaa, 0xe3, 0x25, 0xb9, 0xbd, 0x68, 0x77, 0xf3, 0xf2,
	0xda, 0x06, 0x75, 0xfa, 0xe7, 0xcc, 0xbc, 0x21, 0x96, 0x8c, 0x28, 0x22, 0x14, 0xbd, 0x2b, 0x19,
	0xbd, 0xbe, 0xa4, 0xf7, 0xf3, 0xce, 0x1a, 0x74, 0x57, 0x7e, 0x8f, 0xdc, 0x03, 0x62, 0x17, 0xae,
	0x26, 0x00, 0x99, 0x72, 0x94, 0xd0, 0xf1, 0x5f, 0xe7, 0x8e, 0xf1, 0x36, 0x77, 0x8c, 0xf7, 0xb9,
	0x63, 0x3c, 0x7d, 0x38, 0xa5, 0x41, 0xc3, 0xf3, 0x7f, 0x79, 0x85, 0x17, 0xfa, 0x7b, 0x5f, 0xcd,
	0xde, 0xd7, 0xf9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0xb8, 0xf0, 0x70, 0xab, 0x02, 0x00,
	0x00,
}

func (m *DeploymentEnvironmentUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeploymentEnvironmentUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeploymentEnvironmentUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Environments) > 0 {
		for iNdEx := len(m.Environments) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Environments[iNdEx])
			copy(dAtA[i:], m.Environments[iNdEx])
			i = encodeVarintClusterStatus(dAtA, i, uint64(len(m.Environments[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ClusterStatusUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterStatusUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterStatusUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Msg != nil {
		{
			size := m.Msg.Size()
			i -= size
			if _, err := m.Msg.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *ClusterStatusUpdate_Status) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterStatusUpdate_Status) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Status != nil {
		{
			size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClusterStatus(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *ClusterStatusUpdate_DeploymentEnvUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterStatusUpdate_DeploymentEnvUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DeploymentEnvUpdate != nil {
		{
			size, err := m.DeploymentEnvUpdate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClusterStatus(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *RawClusterHealthInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RawClusterHealthInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RawClusterHealthInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ScannerHealthInfo != nil {
		{
			size, err := m.ScannerHealthInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClusterStatus(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.AdmissionControlHealthInfo != nil {
		{
			size, err := m.AdmissionControlHealthInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClusterStatus(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.CollectorHealthInfo != nil {
		{
			size, err := m.CollectorHealthInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClusterStatus(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClusterHealthResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterHealthResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterHealthResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintClusterStatus(dAtA []byte, offset int, v uint64) int {
	offset -= sovClusterStatus(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeploymentEnvironmentUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Environments) > 0 {
		for _, s := range m.Environments {
			l = len(s)
			n += 1 + l + sovClusterStatus(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ClusterStatusUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Msg != nil {
		n += m.Msg.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ClusterStatusUpdate_Status) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovClusterStatus(uint64(l))
	}
	return n
}
func (m *ClusterStatusUpdate_DeploymentEnvUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DeploymentEnvUpdate != nil {
		l = m.DeploymentEnvUpdate.Size()
		n += 1 + l + sovClusterStatus(uint64(l))
	}
	return n
}
func (m *RawClusterHealthInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CollectorHealthInfo != nil {
		l = m.CollectorHealthInfo.Size()
		n += 1 + l + sovClusterStatus(uint64(l))
	}
	if m.AdmissionControlHealthInfo != nil {
		l = m.AdmissionControlHealthInfo.Size()
		n += 1 + l + sovClusterStatus(uint64(l))
	}
	if m.ScannerHealthInfo != nil {
		l = m.ScannerHealthInfo.Size()
		n += 1 + l + sovClusterStatus(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ClusterHealthResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovClusterStatus(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClusterStatus(x uint64) (n int) {
	return sovClusterStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeploymentEnvironmentUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClusterStatus
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
			return fmt.Errorf("proto: DeploymentEnvironmentUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeploymentEnvironmentUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Environments", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterStatus
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
				return ErrInvalidLengthClusterStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClusterStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Environments = append(m.Environments, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClusterStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClusterStatus
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
func (m *ClusterStatusUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClusterStatus
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
			return fmt.Errorf("proto: ClusterStatusUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterStatusUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterStatus
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
				return ErrInvalidLengthClusterStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClusterStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &storage.ClusterStatus{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Msg = &ClusterStatusUpdate_Status{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeploymentEnvUpdate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterStatus
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
				return ErrInvalidLengthClusterStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClusterStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DeploymentEnvironmentUpdate{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Msg = &ClusterStatusUpdate_DeploymentEnvUpdate{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClusterStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClusterStatus
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
func (m *RawClusterHealthInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClusterStatus
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
			return fmt.Errorf("proto: RawClusterHealthInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RawClusterHealthInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectorHealthInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterStatus
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
				return ErrInvalidLengthClusterStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClusterStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CollectorHealthInfo == nil {
				m.CollectorHealthInfo = &storage.CollectorHealthInfo{}
			}
			if err := m.CollectorHealthInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdmissionControlHealthInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterStatus
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
				return ErrInvalidLengthClusterStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClusterStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AdmissionControlHealthInfo == nil {
				m.AdmissionControlHealthInfo = &storage.AdmissionControlHealthInfo{}
			}
			if err := m.AdmissionControlHealthInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScannerHealthInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterStatus
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
				return ErrInvalidLengthClusterStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClusterStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ScannerHealthInfo == nil {
				m.ScannerHealthInfo = &storage.ScannerHealthInfo{}
			}
			if err := m.ScannerHealthInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClusterStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClusterStatus
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
func (m *ClusterHealthResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClusterStatus
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
			return fmt.Errorf("proto: ClusterHealthResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterHealthResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipClusterStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClusterStatus
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
func skipClusterStatus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClusterStatus
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
					return 0, ErrIntOverflowClusterStatus
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
					return 0, ErrIntOverflowClusterStatus
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
				return 0, ErrInvalidLengthClusterStatus
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClusterStatus
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClusterStatus
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClusterStatus        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClusterStatus          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClusterStatus = fmt.Errorf("proto: unexpected end of group")
)
