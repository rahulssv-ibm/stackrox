// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internalapi/sensor/deployment_iservice.proto

package sensor

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

type GetDeploymentForPodRequest struct {
	PodName              string   `protobuf:"bytes,1,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeploymentForPodRequest) Reset()         { *m = GetDeploymentForPodRequest{} }
func (m *GetDeploymentForPodRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeploymentForPodRequest) ProtoMessage()    {}
func (*GetDeploymentForPodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_79e5898dd478d161, []int{0}
}
func (m *GetDeploymentForPodRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetDeploymentForPodRequest) XXX_MarshalVT(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetDeploymentForPodRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetDeploymentForPodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeploymentForPodRequest.Merge(m, src)
}
func (m *GetDeploymentForPodRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetDeploymentForPodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeploymentForPodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeploymentForPodRequest proto.InternalMessageInfo

func (m *GetDeploymentForPodRequest) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *GetDeploymentForPodRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetDeploymentForPodRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GetDeploymentForPodRequest) Clone() *GetDeploymentForPodRequest {
	if m == nil {
		return nil
	}
	cloned := new(GetDeploymentForPodRequest)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterType((*GetDeploymentForPodRequest)(nil), "sensor.GetDeploymentForPodRequest")
}

func init() {
	proto.RegisterFile("internalapi/sensor/deployment_iservice.proto", fileDescriptor_79e5898dd478d161)
}

var fileDescriptor_79e5898dd478d161 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc9, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0x49, 0x2c, 0xc8, 0xd4, 0x2f, 0x4e, 0xcd, 0x2b, 0xce, 0x2f, 0xd2, 0x4f,
	0x49, 0x2d, 0xc8, 0xc9, 0xaf, 0xcc, 0x4d, 0xcd, 0x2b, 0x89, 0xcf, 0x2c, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xa8, 0x90, 0x92, 0x28, 0x2e,
	0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0x45, 0x52, 0x0a, 0x51, 0xa1, 0x14, 0xca, 0x25, 0xe5, 0x9e, 0x5a,
	0xe2, 0x02, 0x17, 0x76, 0xcb, 0x2f, 0x0a, 0xc8, 0x4f, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x11, 0x92, 0xe4, 0xe2, 0x28, 0xc8, 0x4f, 0x89, 0xcf, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x62, 0x2f, 0xc8, 0x4f, 0xf1, 0x4b, 0xcc, 0x4d, 0x15, 0x92, 0xe1, 0xe2, 0x04,
	0x09, 0x17, 0x17, 0x24, 0x26, 0xa7, 0x4a, 0x30, 0x81, 0xe5, 0x10, 0x02, 0x46, 0xc9, 0x5c, 0x82,
	0x08, 0x33, 0x83, 0x21, 0x6e, 0x12, 0xf2, 0xe3, 0x12, 0xc6, 0x62, 0x97, 0x90, 0x92, 0x1e, 0xc4,
	0x95, 0x7a, 0xb8, 0x1d, 0x22, 0x25, 0xac, 0x07, 0xf5, 0x81, 0x1e, 0x42, 0x85, 0x93, 0xc4, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x43,
	0x14, 0xd4, 0xbf, 0x49, 0x6c, 0x60, 0xcf, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x49, 0x3c,
	0x37, 0x2c, 0x2e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeploymentServiceClient is the client API for DeploymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
type DeploymentServiceClient interface {
	GetDeploymentForPod(ctx context.Context, in *GetDeploymentForPodRequest, opts ...grpc.CallOption) (*storage.Deployment, error)
}

type deploymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeploymentServiceClient(cc grpc.ClientConnInterface) DeploymentServiceClient {
	return &deploymentServiceClient{cc}
}

func (c *deploymentServiceClient) GetDeploymentForPod(ctx context.Context, in *GetDeploymentForPodRequest, opts ...grpc.CallOption) (*storage.Deployment, error) {
	out := new(storage.Deployment)
	err := c.cc.Invoke(ctx, "/sensor.DeploymentService/GetDeploymentForPod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeploymentServiceServer is the server API for DeploymentService service.
type DeploymentServiceServer interface {
	GetDeploymentForPod(context.Context, *GetDeploymentForPodRequest) (*storage.Deployment, error)
}

// UnimplementedDeploymentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeploymentServiceServer struct {
}

func (*UnimplementedDeploymentServiceServer) GetDeploymentForPod(ctx context.Context, req *GetDeploymentForPodRequest) (*storage.Deployment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeploymentForPod not implemented")
}

func RegisterDeploymentServiceServer(s *grpc.Server, srv DeploymentServiceServer) {
	s.RegisterService(&_DeploymentService_serviceDesc, srv)
}

func _DeploymentService_GetDeploymentForPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeploymentForPodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServiceServer).GetDeploymentForPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sensor.DeploymentService/GetDeploymentForPod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServiceServer).GetDeploymentForPod(ctx, req.(*GetDeploymentForPodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeploymentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sensor.DeploymentService",
	HandlerType: (*DeploymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeploymentForPod",
			Handler:    _DeploymentService_GetDeploymentForPod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internalapi/sensor/deployment_iservice.proto",
}

func (m *GetDeploymentForPodRequest) MarshalVT() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetDeploymentForPodRequest) MarshalVTTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetDeploymentForPodRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintDeploymentIservice(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PodName) > 0 {
		i -= len(m.PodName)
		copy(dAtA[i:], m.PodName)
		i = encodeVarintDeploymentIservice(dAtA, i, uint64(len(m.PodName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeploymentIservice(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeploymentIservice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetDeploymentForPodRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PodName)
	if l > 0 {
		n += 1 + l + sovDeploymentIservice(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovDeploymentIservice(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDeploymentIservice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeploymentIservice(x uint64) (n int) {
	return sovDeploymentIservice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetDeploymentForPodRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeploymentIservice
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
			return fmt.Errorf("proto: GetDeploymentForPodRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetDeploymentForPodRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploymentIservice
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
				return ErrInvalidLengthDeploymentIservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeploymentIservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PodName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploymentIservice
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
				return ErrInvalidLengthDeploymentIservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeploymentIservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeploymentIservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeploymentIservice
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
func skipDeploymentIservice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeploymentIservice
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
					return 0, ErrIntOverflowDeploymentIservice
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
					return 0, ErrIntOverflowDeploymentIservice
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
				return 0, ErrInvalidLengthDeploymentIservice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeploymentIservice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeploymentIservice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeploymentIservice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeploymentIservice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeploymentIservice = fmt.Errorf("proto: unexpected end of group")
)
