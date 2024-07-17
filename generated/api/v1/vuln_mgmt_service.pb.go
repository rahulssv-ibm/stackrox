// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/vuln_mgmt_service.proto

package v1

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

type VulnMgmtExportWorkloadsRequest struct {
	// Request timeout in seconds.
	Timeout int32 `protobuf:"varint,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Query to constrain the deployments for which vulnerability data is returned.
	// The queries contain pairs of `Search Option:Value` separated by `+` signs.
	// For HTTP requests the query should be quoted. For example
	// > curl "$ROX_ENDPOINT/v1/export/vuln-mgmt/workloads?query=Deployment%3Ascanner%2BNamespace%3Astackrox"
	// queries vulnerability data for all scanner deployments in the stackrox namespace.
	// See https://docs.openshift.com/acs/operating/search-filter.html for more information.
	Query                string   `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VulnMgmtExportWorkloadsRequest) Reset()         { *m = VulnMgmtExportWorkloadsRequest{} }
func (m *VulnMgmtExportWorkloadsRequest) String() string { return proto.CompactTextString(m) }
func (*VulnMgmtExportWorkloadsRequest) ProtoMessage()    {}
func (*VulnMgmtExportWorkloadsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4673920823fe3266, []int{0}
}
func (m *VulnMgmtExportWorkloadsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VulnMgmtExportWorkloadsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VulnMgmtExportWorkloadsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VulnMgmtExportWorkloadsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VulnMgmtExportWorkloadsRequest.Merge(m, src)
}
func (m *VulnMgmtExportWorkloadsRequest) XXX_Size() int {
	return m.Size()
}
func (m *VulnMgmtExportWorkloadsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VulnMgmtExportWorkloadsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VulnMgmtExportWorkloadsRequest proto.InternalMessageInfo

func (m *VulnMgmtExportWorkloadsRequest) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *VulnMgmtExportWorkloadsRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *VulnMgmtExportWorkloadsRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *VulnMgmtExportWorkloadsRequest) Clone() *VulnMgmtExportWorkloadsRequest {
	if m == nil {
		return nil
	}
	cloned := new(VulnMgmtExportWorkloadsRequest)
	*cloned = *m

	return cloned
}

// The workloads response contains the full image details including the
// vulnerability data.
type VulnMgmtExportWorkloadsResponse struct {
	Deployment           *storage.Deployment `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	Images               []*storage.Image    `protobuf:"bytes,2,rep,name=images,proto3" json:"images,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VulnMgmtExportWorkloadsResponse) Reset()         { *m = VulnMgmtExportWorkloadsResponse{} }
func (m *VulnMgmtExportWorkloadsResponse) String() string { return proto.CompactTextString(m) }
func (*VulnMgmtExportWorkloadsResponse) ProtoMessage()    {}
func (*VulnMgmtExportWorkloadsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4673920823fe3266, []int{1}
}
func (m *VulnMgmtExportWorkloadsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VulnMgmtExportWorkloadsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VulnMgmtExportWorkloadsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VulnMgmtExportWorkloadsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VulnMgmtExportWorkloadsResponse.Merge(m, src)
}
func (m *VulnMgmtExportWorkloadsResponse) XXX_Size() int {
	return m.Size()
}
func (m *VulnMgmtExportWorkloadsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VulnMgmtExportWorkloadsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VulnMgmtExportWorkloadsResponse proto.InternalMessageInfo

func (m *VulnMgmtExportWorkloadsResponse) GetDeployment() *storage.Deployment {
	if m != nil {
		return m.Deployment
	}
	return nil
}

func (m *VulnMgmtExportWorkloadsResponse) GetImages() []*storage.Image {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *VulnMgmtExportWorkloadsResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *VulnMgmtExportWorkloadsResponse) Clone() *VulnMgmtExportWorkloadsResponse {
	if m == nil {
		return nil
	}
	cloned := new(VulnMgmtExportWorkloadsResponse)
	*cloned = *m

	cloned.Deployment = m.Deployment.Clone()
	if m.Images != nil {
		cloned.Images = make([]*storage.Image, len(m.Images))
		for idx, v := range m.Images {
			cloned.Images[idx] = v.Clone()
		}
	}
	return cloned
}

func init() {
	proto.RegisterType((*VulnMgmtExportWorkloadsRequest)(nil), "v1.VulnMgmtExportWorkloadsRequest")
	proto.RegisterType((*VulnMgmtExportWorkloadsResponse)(nil), "v1.VulnMgmtExportWorkloadsResponse")
}

func init() { proto.RegisterFile("api/v1/vuln_mgmt_service.proto", fileDescriptor_4673920823fe3266) }

var fileDescriptor_4673920823fe3266 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0x7b, 0x91, 0x56, 0xbc, 0x82, 0xc2, 0x55, 0x30, 0x14, 0x89, 0x21, 0x42, 0xe9, 0xe2,
	0xc5, 0xb4, 0x83, 0x83, 0x9b, 0xe8, 0xe0, 0x20, 0x48, 0x04, 0x15, 0x97, 0x72, 0xb6, 0x47, 0x08,
	0x4d, 0xf2, 0xa6, 0xb9, 0xcb, 0xd9, 0x2e, 0x0e, 0xce, 0x6e, 0x2e, 0x8e, 0x7e, 0x1c, 0x47, 0xc1,
	0x2f, 0x20, 0xd5, 0x0f, 0x22, 0xf9, 0x57, 0x5d, 0xda, 0xf1, 0xbd, 0xe7, 0x77, 0x0f, 0xcf, 0xbd,
	0xcf, 0x61, 0x83, 0xc5, 0xbe, 0xad, 0x1c, 0x5b, 0xa5, 0x41, 0x34, 0x08, 0xbd, 0x50, 0x0e, 0x04,
	0x4f, 0x94, 0x3f, 0xe4, 0x34, 0x4e, 0x40, 0x02, 0xd1, 0x94, 0xd3, 0xde, 0xf5, 0x00, 0xbc, 0x80,
	0xdb, 0x19, 0xca, 0xa2, 0x08, 0x24, 0x93, 0x3e, 0x44, 0xa2, 0x20, 0xda, 0xba, 0x90, 0x90, 0x30,
	0x8f, 0xdb, 0x23, 0x1e, 0x07, 0x30, 0x0b, 0x79, 0x24, 0x4b, 0xa5, 0x55, 0x29, 0x7e, 0xc8, 0xbc,
	0xd2, 0xd0, 0xba, 0xc4, 0xc6, 0x75, 0x1a, 0x44, 0x17, 0x5e, 0x28, 0xcf, 0xa6, 0x31, 0x24, 0xf2,
	0x06, 0x92, 0x71, 0x00, 0x6c, 0x24, 0x5c, 0x3e, 0x49, 0xb9, 0x90, 0x44, 0xc7, 0xeb, 0xd2, 0x0f,
	0x39, 0xa4, 0x52, 0x47, 0x26, 0xea, 0xd6, 0xdd, 0x6a, 0x24, 0xdb, 0xb8, 0x3e, 0x49, 0x79, 0x32,
	0xd3, 0x35, 0x13, 0x75, 0x37, 0xdc, 0x62, 0xb0, 0x1e, 0xf1, 0xde, 0x52, 0x47, 0x11, 0x43, 0x24,
	0x38, 0xe9, 0x63, 0xfc, 0x97, 0x2e, 0x77, 0x6d, 0xf6, 0x5a, 0xb4, 0x8c, 0x47, 0x4f, 0x17, 0x92,
	0xfb, 0x0f, 0x23, 0x1d, 0xdc, 0xc8, 0x83, 0x0b, 0x5d, 0x33, 0xd7, 0xba, 0xcd, 0xde, 0xe6, 0xe2,
	0xc2, 0x79, 0x76, 0xec, 0x96, 0x6a, 0xef, 0x0d, 0xe1, 0xad, 0x2a, 0xc0, 0x55, 0xb1, 0x3c, 0xf2,
	0x8c, 0xf0, 0xce, 0x92, 0x50, 0xc4, 0xa2, 0xca, 0xa1, 0xab, 0x77, 0xd0, 0xde, 0x5f, 0xc9, 0x14,
	0xaf, 0xb2, 0x3a, 0x4f, 0x9f, 0x3f, 0x2f, 0x9a, 0x49, 0x8c, 0xac, 0x40, 0x9e, 0x43, 0x79, 0x8f,
	0x07, 0x59, 0x8f, 0xf6, 0x43, 0xc5, 0x1f, 0xa2, 0x93, 0xa3, 0xf7, 0xb9, 0x81, 0x3e, 0xe6, 0x06,
	0xfa, 0x9a, 0x1b, 0xe8, 0xf5, 0xdb, 0xa8, 0x61, 0xdd, 0x07, 0x2a, 0x24, 0x1b, 0x8e, 0x13, 0x98,
	0x16, 0xcd, 0x50, 0x16, 0xfb, 0x54, 0x39, 0x77, 0x4d, 0x6a, 0x17, 0x7f, 0xe2, 0x58, 0x39, 0xb7,
	0xb5, 0xfb, 0x46, 0x2e, 0xf6, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x81, 0x77, 0xf9, 0xff, 0x29,
	0x02, 0x00, 0x00,
}

func (m *VulnMgmtExportWorkloadsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VulnMgmtExportWorkloadsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VulnMgmtExportWorkloadsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Query) > 0 {
		i -= len(m.Query)
		copy(dAtA[i:], m.Query)
		i = encodeVarintVulnMgmtService(dAtA, i, uint64(len(m.Query)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timeout != 0 {
		i = encodeVarintVulnMgmtService(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *VulnMgmtExportWorkloadsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VulnMgmtExportWorkloadsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VulnMgmtExportWorkloadsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Images) > 0 {
		for iNdEx := len(m.Images) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Images[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVulnMgmtService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Deployment != nil {
		{
			size, err := m.Deployment.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVulnMgmtService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVulnMgmtService(dAtA []byte, offset int, v uint64) int {
	offset -= sovVulnMgmtService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VulnMgmtExportWorkloadsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timeout != 0 {
		n += 1 + sovVulnMgmtService(uint64(m.Timeout))
	}
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovVulnMgmtService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *VulnMgmtExportWorkloadsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Deployment != nil {
		l = m.Deployment.Size()
		n += 1 + l + sovVulnMgmtService(uint64(l))
	}
	if len(m.Images) > 0 {
		for _, e := range m.Images {
			l = e.Size()
			n += 1 + l + sovVulnMgmtService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovVulnMgmtService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVulnMgmtService(x uint64) (n int) {
	return sovVulnMgmtService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VulnMgmtExportWorkloadsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVulnMgmtService
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
			return fmt.Errorf("proto: VulnMgmtExportWorkloadsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VulnMgmtExportWorkloadsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVulnMgmtService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVulnMgmtService
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
				return ErrInvalidLengthVulnMgmtService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVulnMgmtService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVulnMgmtService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVulnMgmtService
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
func (m *VulnMgmtExportWorkloadsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVulnMgmtService
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
			return fmt.Errorf("proto: VulnMgmtExportWorkloadsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VulnMgmtExportWorkloadsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deployment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVulnMgmtService
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
				return ErrInvalidLengthVulnMgmtService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVulnMgmtService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Deployment == nil {
				m.Deployment = &storage.Deployment{}
			}
			if err := m.Deployment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Images", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVulnMgmtService
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
				return ErrInvalidLengthVulnMgmtService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVulnMgmtService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Images = append(m.Images, &storage.Image{})
			if err := m.Images[len(m.Images)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVulnMgmtService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVulnMgmtService
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
func skipVulnMgmtService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVulnMgmtService
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
					return 0, ErrIntOverflowVulnMgmtService
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
					return 0, ErrIntOverflowVulnMgmtService
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
				return 0, ErrInvalidLengthVulnMgmtService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVulnMgmtService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVulnMgmtService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVulnMgmtService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVulnMgmtService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVulnMgmtService = fmt.Errorf("proto: unexpected end of group")
)
