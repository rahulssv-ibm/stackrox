// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/process_listening_on_port_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	storage "github.com/stackrox/rox/generated/storage"
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

type GetProcessesListeningOnPortsRequest struct {
	DeploymentId         string   `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProcessesListeningOnPortsRequest) Reset()         { *m = GetProcessesListeningOnPortsRequest{} }
func (m *GetProcessesListeningOnPortsRequest) String() string { return proto.CompactTextString(m) }
func (*GetProcessesListeningOnPortsRequest) ProtoMessage()    {}
func (*GetProcessesListeningOnPortsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8bf40985da37317, []int{0}
}
func (m *GetProcessesListeningOnPortsRequest) XXX_UnmarshalVT(b []byte) error {
	return m.UnmarshalVT(b)
}
func (m *GetProcessesListeningOnPortsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetProcessesListeningOnPortsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetProcessesListeningOnPortsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProcessesListeningOnPortsRequest.Merge(m, src)
}
func (m *GetProcessesListeningOnPortsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetProcessesListeningOnPortsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProcessesListeningOnPortsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProcessesListeningOnPortsRequest proto.InternalMessageInfo

func (m *GetProcessesListeningOnPortsRequest) GetDeploymentId() string {
	if m != nil {
		return m.DeploymentId
	}
	return ""
}

func (m *GetProcessesListeningOnPortsRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GetProcessesListeningOnPortsRequest) Clone() *GetProcessesListeningOnPortsRequest {
	if m == nil {
		return nil
	}
	cloned := new(GetProcessesListeningOnPortsRequest)
	*cloned = *m

	return cloned
}

type GetProcessesListeningOnPortsResponse struct {
	ListeningEndpoints   []*storage.ProcessListeningOnPort `protobuf:"bytes,1,rep,name=listening_endpoints,json=listeningEndpoints,proto3" json:"listening_endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *GetProcessesListeningOnPortsResponse) Reset()         { *m = GetProcessesListeningOnPortsResponse{} }
func (m *GetProcessesListeningOnPortsResponse) String() string { return proto.CompactTextString(m) }
func (*GetProcessesListeningOnPortsResponse) ProtoMessage()    {}
func (*GetProcessesListeningOnPortsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8bf40985da37317, []int{1}
}
func (m *GetProcessesListeningOnPortsResponse) XXX_UnmarshalVT(b []byte) error {
	return m.UnmarshalVT(b)
}
func (m *GetProcessesListeningOnPortsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetProcessesListeningOnPortsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetProcessesListeningOnPortsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProcessesListeningOnPortsResponse.Merge(m, src)
}
func (m *GetProcessesListeningOnPortsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetProcessesListeningOnPortsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProcessesListeningOnPortsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProcessesListeningOnPortsResponse proto.InternalMessageInfo

func (m *GetProcessesListeningOnPortsResponse) GetListeningEndpoints() []*storage.ProcessListeningOnPort {
	if m != nil {
		return m.ListeningEndpoints
	}
	return nil
}

func (m *GetProcessesListeningOnPortsResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GetProcessesListeningOnPortsResponse) Clone() *GetProcessesListeningOnPortsResponse {
	if m == nil {
		return nil
	}
	cloned := new(GetProcessesListeningOnPortsResponse)
	*cloned = *m

	if m.ListeningEndpoints != nil {
		cloned.ListeningEndpoints = make([]*storage.ProcessListeningOnPort, len(m.ListeningEndpoints))
		for idx, v := range m.ListeningEndpoints {
			cloned.ListeningEndpoints[idx] = v.Clone()
		}
	}
	return cloned
}

func init() {
	proto.RegisterType((*GetProcessesListeningOnPortsRequest)(nil), "v1.GetProcessesListeningOnPortsRequest")
	proto.RegisterType((*GetProcessesListeningOnPortsResponse)(nil), "v1.GetProcessesListeningOnPortsResponse")
}

func init() {
	proto.RegisterFile("api/v1/process_listening_on_port_service.proto", fileDescriptor_d8bf40985da37317)
}

var fileDescriptor_d8bf40985da37317 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbd, 0x4a, 0x03, 0x41,
	0x14, 0x85, 0x33, 0x11, 0x04, 0x47, 0x6d, 0x46, 0x84, 0x18, 0x64, 0x0d, 0x89, 0x90, 0x54, 0xb3,
	0x6c, 0xb4, 0xb2, 0x14, 0x24, 0x28, 0x01, 0x43, 0x6c, 0xc4, 0x66, 0x59, 0xb3, 0x97, 0x65, 0x30,
	0xce, 0x1d, 0x67, 0xae, 0x4b, 0x44, 0x6c, 0x7c, 0x05, 0x1b, 0x9f, 0xc0, 0x27, 0xb1, 0xb0, 0x14,
	0x7c, 0x01, 0x89, 0x3e, 0x88, 0xc4, 0x4d, 0x5c, 0x34, 0xfe, 0xb5, 0x87, 0xef, 0x1c, 0xe6, 0x9c,
	0x3b, 0x5c, 0x46, 0x46, 0xf9, 0x69, 0xe0, 0x1b, 0x8b, 0x3d, 0x70, 0x2e, 0xec, 0x2b, 0x47, 0xa0,
	0x95, 0x4e, 0x42, 0xd4, 0xa1, 0x41, 0x4b, 0xa1, 0x03, 0x9b, 0xaa, 0x1e, 0x48, 0x63, 0x91, 0x50,
	0x14, 0xd3, 0xa0, 0xbc, 0x9a, 0x20, 0x26, 0x7d, 0xf0, 0x47, 0xd6, 0x48, 0x6b, 0xa4, 0x88, 0x14,
	0x6a, 0x97, 0x11, 0xe5, 0xba, 0x23, 0xb4, 0x51, 0x02, 0x3f, 0x47, 0x66, 0x60, 0x75, 0x8f, 0xd7,
	0x5a, 0x40, 0x9d, 0x8c, 0x02, 0xd7, 0x9e, 0x60, 0xfb, 0xba, 0x83, 0x96, 0x5c, 0x17, 0xce, 0xce,
	0xc1, 0x91, 0xa8, 0xf1, 0xc5, 0x18, 0x4c, 0x1f, 0x2f, 0x4e, 0x41, 0x53, 0xa8, 0xe2, 0x12, 0xab,
	0xb0, 0xc6, 0x5c, 0x77, 0x21, 0x17, 0x77, 0xe3, 0xea, 0x80, 0xaf, 0xff, 0x9e, 0xe5, 0x0c, 0x6a,
	0x07, 0xa2, 0xc3, 0x97, 0xf2, 0xe7, 0x80, 0x8e, 0x0d, 0x2a, 0x4d, 0xae, 0xc4, 0x2a, 0x33, 0x8d,
	0xf9, 0xe6, 0x9a, 0x1c, 0x3f, 0x5d, 0x8e, 0x83, 0xbe, 0xc4, 0x74, 0xc5, 0x87, 0x77, 0x67, 0x62,
	0x6d, 0xde, 0x33, 0xbe, 0xd2, 0x9e, 0x92, 0x0f, 0xb2, 0xd1, 0xc4, 0x1d, 0xe3, 0xcb, 0x2d, 0xa0,
	0x69, 0x40, 0xd4, 0x65, 0x1a, 0xc8, 0x7f, 0xf4, 0x2f, 0x37, 0xfe, 0x06, 0xb3, 0x72, 0xd5, 0xad,
	0xeb, 0xa7, 0xd7, 0x9b, 0xe2, 0xa6, 0x68, 0x8e, 0x0e, 0xfa, 0x4d, 0x4d, 0x3f, 0x9f, 0xcc, 0xbf,
	0xfc, 0xb4, 0xe9, 0xd5, 0xb6, 0x7c, 0x18, 0x7a, 0xec, 0x71, 0xe8, 0xb1, 0xe7, 0xa1, 0xc7, 0x6e,
	0x5f, 0xbc, 0x02, 0x2f, 0x29, 0x94, 0x8e, 0xa2, 0xde, 0x89, 0xc5, 0x41, 0x76, 0xb1, 0xd1, 0x5f,
	0x91, 0x69, 0x70, 0x54, 0x4c, 0x83, 0xc3, 0xc2, 0xf1, 0xec, 0xbb, 0xb6, 0xf1, 0x16, 0x00, 0x00,
	0xff, 0xff, 0x42, 0x49, 0x2e, 0xa5, 0x42, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ListeningEndpointsServiceClient is the client API for ListeningEndpointsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
type ListeningEndpointsServiceClient interface {
	// GetListeningEndpoints returns the listening endpoints and the processes that opened them for a given deployment
	GetListeningEndpoints(ctx context.Context, in *GetProcessesListeningOnPortsRequest, opts ...grpc.CallOption) (*GetProcessesListeningOnPortsResponse, error)
}

type listeningEndpointsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewListeningEndpointsServiceClient(cc grpc.ClientConnInterface) ListeningEndpointsServiceClient {
	return &listeningEndpointsServiceClient{cc}
}

func (c *listeningEndpointsServiceClient) GetListeningEndpoints(ctx context.Context, in *GetProcessesListeningOnPortsRequest, opts ...grpc.CallOption) (*GetProcessesListeningOnPortsResponse, error) {
	out := new(GetProcessesListeningOnPortsResponse)
	err := c.cc.Invoke(ctx, "/v1.ListeningEndpointsService/GetListeningEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListeningEndpointsServiceServer is the server API for ListeningEndpointsService service.
type ListeningEndpointsServiceServer interface {
	// GetListeningEndpoints returns the listening endpoints and the processes that opened them for a given deployment
	GetListeningEndpoints(context.Context, *GetProcessesListeningOnPortsRequest) (*GetProcessesListeningOnPortsResponse, error)
}

// UnimplementedListeningEndpointsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedListeningEndpointsServiceServer struct {
}

func (*UnimplementedListeningEndpointsServiceServer) GetListeningEndpoints(ctx context.Context, req *GetProcessesListeningOnPortsRequest) (*GetProcessesListeningOnPortsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListeningEndpoints not implemented")
}

func RegisterListeningEndpointsServiceServer(s *grpc.Server, srv ListeningEndpointsServiceServer) {
	s.RegisterService(&_ListeningEndpointsService_serviceDesc, srv)
}

func _ListeningEndpointsService_GetListeningEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProcessesListeningOnPortsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListeningEndpointsServiceServer).GetListeningEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ListeningEndpointsService/GetListeningEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListeningEndpointsServiceServer).GetListeningEndpoints(ctx, req.(*GetProcessesListeningOnPortsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ListeningEndpointsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ListeningEndpointsService",
	HandlerType: (*ListeningEndpointsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetListeningEndpoints",
			Handler:    _ListeningEndpointsService_GetListeningEndpoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/process_listening_on_port_service.proto",
}

func (m *GetProcessesListeningOnPortsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetProcessesListeningOnPortsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetProcessesListeningOnPortsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.DeploymentId) > 0 {
		i -= len(m.DeploymentId)
		copy(dAtA[i:], m.DeploymentId)
		i = encodeVarintProcessListeningOnPortService(dAtA, i, uint64(len(m.DeploymentId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetProcessesListeningOnPortsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetProcessesListeningOnPortsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetProcessesListeningOnPortsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ListeningEndpoints) > 0 {
		for iNdEx := len(m.ListeningEndpoints) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ListeningEndpoints[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProcessListeningOnPortService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintProcessListeningOnPortService(dAtA []byte, offset int, v uint64) int {
	offset -= sovProcessListeningOnPortService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetProcessesListeningOnPortsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DeploymentId)
	if l > 0 {
		n += 1 + l + sovProcessListeningOnPortService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetProcessesListeningOnPortsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ListeningEndpoints) > 0 {
		for _, e := range m.ListeningEndpoints {
			l = e.Size()
			n += 1 + l + sovProcessListeningOnPortService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProcessListeningOnPortService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProcessListeningOnPortService(x uint64) (n int) {
	return sovProcessListeningOnPortService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetProcessesListeningOnPortsRequest) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProcessListeningOnPortService
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
			return fmt.Errorf("proto: GetProcessesListeningOnPortsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetProcessesListeningOnPortsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeploymentId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessListeningOnPortService
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
				return ErrInvalidLengthProcessListeningOnPortService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProcessListeningOnPortService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeploymentId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProcessListeningOnPortService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProcessListeningOnPortService
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
func (m *GetProcessesListeningOnPortsResponse) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProcessListeningOnPortService
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
			return fmt.Errorf("proto: GetProcessesListeningOnPortsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetProcessesListeningOnPortsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListeningEndpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessListeningOnPortService
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
				return ErrInvalidLengthProcessListeningOnPortService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProcessListeningOnPortService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListeningEndpoints = append(m.ListeningEndpoints, &storage.ProcessListeningOnPort{})
			if err := m.ListeningEndpoints[len(m.ListeningEndpoints)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProcessListeningOnPortService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProcessListeningOnPortService
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
func skipProcessListeningOnPortService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProcessListeningOnPortService
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
					return 0, ErrIntOverflowProcessListeningOnPortService
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
					return 0, ErrIntOverflowProcessListeningOnPortService
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
				return 0, ErrInvalidLengthProcessListeningOnPortService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProcessListeningOnPortService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProcessListeningOnPortService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProcessListeningOnPortService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProcessListeningOnPortService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProcessListeningOnPortService = fmt.Errorf("proto: unexpected end of group")
)
