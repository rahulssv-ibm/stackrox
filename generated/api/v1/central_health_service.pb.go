// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/central_health_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetUpgradeStatusResponse struct {
	UpgradeStatus        *CentralUpgradeStatus `protobuf:"bytes,1,opt,name=upgradeStatus,proto3" json:"upgradeStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetUpgradeStatusResponse) Reset()         { *m = GetUpgradeStatusResponse{} }
func (m *GetUpgradeStatusResponse) String() string { return proto.CompactTextString(m) }
func (*GetUpgradeStatusResponse) ProtoMessage()    {}
func (*GetUpgradeStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e4bd547a0084fea, []int{0}
}
func (m *GetUpgradeStatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetUpgradeStatusResponse) XXX_MarshalVT(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetUpgradeStatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetUpgradeStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUpgradeStatusResponse.Merge(m, src)
}
func (m *GetUpgradeStatusResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetUpgradeStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUpgradeStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUpgradeStatusResponse proto.InternalMessageInfo

func (m *GetUpgradeStatusResponse) GetUpgradeStatus() *CentralUpgradeStatus {
	if m != nil {
		return m.UpgradeStatus
	}
	return nil
}

func (m *GetUpgradeStatusResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GetUpgradeStatusResponse) Clone() *GetUpgradeStatusResponse {
	if m == nil {
		return nil
	}
	cloned := new(GetUpgradeStatusResponse)
	*cloned = *m

	cloned.UpgradeStatus = m.UpgradeStatus.Clone()
	return cloned
}

type CentralUpgradeStatus struct {
	// Current Central Version
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The version of previous clone in Central. This is the version we can force rollback to.
	ForceRollbackTo string `protobuf:"bytes,2,opt,name=force_rollback_to,json=forceRollbackTo,proto3" json:"force_rollback_to,omitempty"`
	// If true, we can rollback to the current version if an upgrade failed.
	CanRollbackAfterUpgrade bool `protobuf:"varint,3,opt,name=can_rollback_after_upgrade,json=canRollbackAfterUpgrade,proto3" json:"can_rollback_after_upgrade,omitempty"`
	// Current disk space stats for upgrade
	SpaceRequiredForRollbackAfterUpgrade  int64    `protobuf:"varint,4,opt,name=space_required_for_rollback_after_upgrade,json=spaceRequiredForRollbackAfterUpgrade,proto3" json:"space_required_for_rollback_after_upgrade,omitempty"`    // Deprecated: Do not use.
	SpaceAvailableForRollbackAfterUpgrade int64    `protobuf:"varint,5,opt,name=space_available_for_rollback_after_upgrade,json=spaceAvailableForRollbackAfterUpgrade,proto3" json:"space_available_for_rollback_after_upgrade,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral                  struct{} `json:"-"`
	XXX_unrecognized                      []byte   `json:"-"`
	XXX_sizecache                         int32    `json:"-"`
}

func (m *CentralUpgradeStatus) Reset()         { *m = CentralUpgradeStatus{} }
func (m *CentralUpgradeStatus) String() string { return proto.CompactTextString(m) }
func (*CentralUpgradeStatus) ProtoMessage()    {}
func (*CentralUpgradeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e4bd547a0084fea, []int{1}
}
func (m *CentralUpgradeStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CentralUpgradeStatus) XXX_MarshalVT(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CentralUpgradeStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CentralUpgradeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CentralUpgradeStatus.Merge(m, src)
}
func (m *CentralUpgradeStatus) XXX_Size() int {
	return m.Size()
}
func (m *CentralUpgradeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_CentralUpgradeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_CentralUpgradeStatus proto.InternalMessageInfo

func (m *CentralUpgradeStatus) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CentralUpgradeStatus) GetForceRollbackTo() string {
	if m != nil {
		return m.ForceRollbackTo
	}
	return ""
}

func (m *CentralUpgradeStatus) GetCanRollbackAfterUpgrade() bool {
	if m != nil {
		return m.CanRollbackAfterUpgrade
	}
	return false
}

// Deprecated: Do not use.
func (m *CentralUpgradeStatus) GetSpaceRequiredForRollbackAfterUpgrade() int64 {
	if m != nil {
		return m.SpaceRequiredForRollbackAfterUpgrade
	}
	return 0
}

// Deprecated: Do not use.
func (m *CentralUpgradeStatus) GetSpaceAvailableForRollbackAfterUpgrade() int64 {
	if m != nil {
		return m.SpaceAvailableForRollbackAfterUpgrade
	}
	return 0
}

func (m *CentralUpgradeStatus) MessageClone() proto.Message {
	return m.Clone()
}
func (m *CentralUpgradeStatus) Clone() *CentralUpgradeStatus {
	if m == nil {
		return nil
	}
	cloned := new(CentralUpgradeStatus)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterType((*GetUpgradeStatusResponse)(nil), "v1.GetUpgradeStatusResponse")
	proto.RegisterType((*CentralUpgradeStatus)(nil), "v1.CentralUpgradeStatus")
}

func init() {
	proto.RegisterFile("api/v1/central_health_service.proto", fileDescriptor_6e4bd547a0084fea)
}

var fileDescriptor_6e4bd547a0084fea = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x6b, 0x14, 0x31,
	0x18, 0x36, 0x53, 0xbf, 0x1a, 0x11, 0x35, 0x08, 0x86, 0xa1, 0x8c, 0xeb, 0xaa, 0xb8, 0xf6, 0x30,
	0xc3, 0xd4, 0xa3, 0x20, 0xb4, 0xe2, 0xc7, 0x79, 0xaa, 0x20, 0x15, 0x19, 0xde, 0x9d, 0x7d, 0x77,
	0x1b, 0x3a, 0xe6, 0x1d, 0x93, 0x6c, 0x50, 0xf0, 0xe4, 0x5f, 0xf0, 0xe2, 0x4f, 0xf2, 0x28, 0x88,
	0x77, 0x59, 0xfd, 0x21, 0x32, 0x99, 0xd9, 0xda, 0xca, 0xae, 0xb7, 0xe4, 0x7d, 0x3e, 0xf2, 0x90,
	0xf7, 0xe1, 0xb7, 0xa1, 0x51, 0x99, 0xcf, 0xb3, 0x0a, 0xb5, 0x33, 0x50, 0x97, 0x87, 0x08, 0xb5,
	0x3b, 0x2c, 0x2d, 0x1a, 0xaf, 0x2a, 0x4c, 0x1b, 0x43, 0x8e, 0x44, 0xe4, 0xf3, 0x58, 0xf4, 0x44,
	0x7c, 0xdb, 0xb8, 0x0f, 0xdd, 0x3c, 0xde, 0x9a, 0x11, 0xcd, 0x6a, 0xcc, 0x5a, 0x08, 0xb4, 0x26,
	0x07, 0x4e, 0x91, 0xb6, 0x1d, 0x3a, 0x3c, 0xe0, 0xf2, 0x19, 0xba, 0x97, 0xcd, 0xcc, 0xc0, 0x04,
	0xf7, 0x1d, 0xb8, 0xb9, 0x2d, 0xd0, 0x36, 0xa4, 0x2d, 0x8a, 0x47, 0xfc, 0xf2, 0xfc, 0x24, 0x20,
	0xd9, 0x80, 0x8d, 0x2e, 0xed, 0xc8, 0xd4, 0xe7, 0xe9, 0xe3, 0x2e, 0xca, 0x69, 0xe1, 0x69, 0xfa,
	0xf0, 0x47, 0xc4, 0xaf, 0xaf, 0xe2, 0x09, 0xc9, 0x2f, 0x78, 0x34, 0x56, 0x91, 0x0e, 0x96, 0x9b,
	0xc5, 0xf2, 0x2a, 0xb6, 0xf9, 0xb5, 0x29, 0x99, 0x0a, 0x4b, 0x43, 0x75, 0x3d, 0x86, 0xea, 0xa8,
	0x74, 0x24, 0xa3, 0xc0, 0xb9, 0x12, 0x80, 0xa2, 0x9f, 0xbf, 0x20, 0xf1, 0x90, 0xc7, 0x15, 0xe8,
	0xbf, 0x4c, 0x98, 0x3a, 0x34, 0x65, 0x1f, 0x41, 0x6e, 0x0c, 0xd8, 0xe8, 0x62, 0x71, 0xa3, 0x02,
	0xbd, 0x94, 0xec, 0xb6, 0x78, 0x1f, 0x44, 0xbc, 0xe6, 0xf7, 0x6d, 0x03, 0xed, 0x43, 0xf8, 0x6e,
	0xae, 0x0c, 0x4e, 0xca, 0x29, 0x99, 0x75, 0x5e, 0x67, 0x07, 0x6c, 0xb4, 0xb1, 0x17, 0x49, 0x56,
	0xdc, 0x09, 0xa2, 0xa2, 0xd7, 0x3c, 0x25, 0xb3, 0xd2, 0xfc, 0x0d, 0xdf, 0xee, 0xcc, 0xc1, 0x83,
	0xaa, 0x61, 0x5c, 0xe3, 0xff, 0xdc, 0xcf, 0x1d, 0xbb, 0xdf, 0x0d, 0xaa, 0xdd, 0xa5, 0x68, 0x8d,
	0xfd, 0xce, 0xc7, 0xe3, 0x6f, 0x7d, 0x1e, 0x8a, 0xb0, 0xdf, 0xf5, 0x40, 0x4c, 0xf8, 0xd5, 0x7f,
	0x77, 0x29, 0x36, 0xdb, 0x65, 0x3d, 0x69, 0xeb, 0x10, 0x6f, 0xb5, 0xc7, 0x75, 0xcb, 0x1e, 0xde,
	0xfb, 0xf4, 0xfd, 0xf7, 0xe7, 0xe8, 0x96, 0xb8, 0x79, 0xa2, 0x68, 0x5d, 0xcf, 0xb2, 0x3e, 0xa8,
	0x0d, 0x82, 0xbd, 0xf4, 0xeb, 0x22, 0x61, 0xdf, 0x16, 0x09, 0xfb, 0xb9, 0x48, 0xd8, 0x97, 0x5f,
	0xc9, 0x19, 0x2e, 0x15, 0xa5, 0xd6, 0x41, 0x75, 0x64, 0xe8, 0x7d, 0x57, 0xab, 0x14, 0x1a, 0x95,
	0xfa, 0xfc, 0x20, 0xf2, 0xf9, 0x2b, 0x36, 0x3e, 0x1f, 0x66, 0x0f, 0xfe, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xe3, 0x2c, 0x9d, 0x0b, 0xc7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CentralHealthServiceClient is the client API for CentralHealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
type CentralHealthServiceClient interface {
	GetUpgradeStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetUpgradeStatusResponse, error)
}

type centralHealthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCentralHealthServiceClient(cc grpc.ClientConnInterface) CentralHealthServiceClient {
	return &centralHealthServiceClient{cc}
}

func (c *centralHealthServiceClient) GetUpgradeStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetUpgradeStatusResponse, error) {
	out := new(GetUpgradeStatusResponse)
	err := c.cc.Invoke(ctx, "/v1.CentralHealthService/GetUpgradeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CentralHealthServiceServer is the server API for CentralHealthService service.
type CentralHealthServiceServer interface {
	GetUpgradeStatus(context.Context, *Empty) (*GetUpgradeStatusResponse, error)
}

// UnimplementedCentralHealthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCentralHealthServiceServer struct {
}

func (*UnimplementedCentralHealthServiceServer) GetUpgradeStatus(ctx context.Context, req *Empty) (*GetUpgradeStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUpgradeStatus not implemented")
}

func RegisterCentralHealthServiceServer(s *grpc.Server, srv CentralHealthServiceServer) {
	s.RegisterService(&_CentralHealthService_serviceDesc, srv)
}

func _CentralHealthService_GetUpgradeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentralHealthServiceServer).GetUpgradeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CentralHealthService/GetUpgradeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentralHealthServiceServer).GetUpgradeStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _CentralHealthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.CentralHealthService",
	HandlerType: (*CentralHealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUpgradeStatus",
			Handler:    _CentralHealthService_GetUpgradeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/central_health_service.proto",
}

func (m *GetUpgradeStatusResponse) MarshalVT() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetUpgradeStatusResponse) MarshalVTTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetUpgradeStatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.UpgradeStatus != nil {
		{
			size, err := m.UpgradeStatus.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCentralHealthService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CentralUpgradeStatus) MarshalVT() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CentralUpgradeStatus) MarshalVTTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CentralUpgradeStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.SpaceAvailableForRollbackAfterUpgrade != 0 {
		i = encodeVarintCentralHealthService(dAtA, i, uint64(m.SpaceAvailableForRollbackAfterUpgrade))
		i--
		dAtA[i] = 0x28
	}
	if m.SpaceRequiredForRollbackAfterUpgrade != 0 {
		i = encodeVarintCentralHealthService(dAtA, i, uint64(m.SpaceRequiredForRollbackAfterUpgrade))
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
		i = encodeVarintCentralHealthService(dAtA, i, uint64(len(m.ForceRollbackTo)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintCentralHealthService(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCentralHealthService(dAtA []byte, offset int, v uint64) int {
	offset -= sovCentralHealthService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetUpgradeStatusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UpgradeStatus != nil {
		l = m.UpgradeStatus.Size()
		n += 1 + l + sovCentralHealthService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CentralUpgradeStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovCentralHealthService(uint64(l))
	}
	l = len(m.ForceRollbackTo)
	if l > 0 {
		n += 1 + l + sovCentralHealthService(uint64(l))
	}
	if m.CanRollbackAfterUpgrade {
		n += 2
	}
	if m.SpaceRequiredForRollbackAfterUpgrade != 0 {
		n += 1 + sovCentralHealthService(uint64(m.SpaceRequiredForRollbackAfterUpgrade))
	}
	if m.SpaceAvailableForRollbackAfterUpgrade != 0 {
		n += 1 + sovCentralHealthService(uint64(m.SpaceAvailableForRollbackAfterUpgrade))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCentralHealthService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCentralHealthService(x uint64) (n int) {
	return sovCentralHealthService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetUpgradeStatusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCentralHealthService
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
					return ErrIntOverflowCentralHealthService
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
				return ErrInvalidLengthCentralHealthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCentralHealthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpgradeStatus == nil {
				m.UpgradeStatus = &CentralUpgradeStatus{}
			}
			if err := m.UpgradeStatus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCentralHealthService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCentralHealthService
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
func (m *CentralUpgradeStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCentralHealthService
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
					return ErrIntOverflowCentralHealthService
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
				return ErrInvalidLengthCentralHealthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCentralHealthService
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
					return ErrIntOverflowCentralHealthService
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
				return ErrInvalidLengthCentralHealthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCentralHealthService
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
					return ErrIntOverflowCentralHealthService
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
					return ErrIntOverflowCentralHealthService
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
					return ErrIntOverflowCentralHealthService
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
			skippy, err := skipCentralHealthService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCentralHealthService
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
func skipCentralHealthService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCentralHealthService
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
					return 0, ErrIntOverflowCentralHealthService
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
					return 0, ErrIntOverflowCentralHealthService
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
				return 0, ErrInvalidLengthCentralHealthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCentralHealthService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCentralHealthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCentralHealthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCentralHealthService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCentralHealthService = fmt.Errorf("proto: unexpected end of group")
)
