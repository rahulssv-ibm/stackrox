// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/pod_service.proto

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

type PodsResponse struct {
	Pods                 []*storage.Pod `protobuf:"bytes,1,rep,name=pods,proto3" json:"pods,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PodsResponse) Reset()         { *m = PodsResponse{} }
func (m *PodsResponse) String() string { return proto.CompactTextString(m) }
func (*PodsResponse) ProtoMessage()    {}
func (*PodsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5092c07627247a99, []int{0}
}
func (m *PodsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PodsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PodsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PodsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodsResponse.Merge(m, src)
}
func (m *PodsResponse) XXX_Size() int {
	return m.Size()
}
func (m *PodsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PodsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PodsResponse proto.InternalMessageInfo

func (m *PodsResponse) GetPods() []*storage.Pod {
	if m != nil {
		return m.Pods
	}
	return nil
}

func (m *PodsResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *PodsResponse) Clone() *PodsResponse {
	if m == nil {
		return nil
	}
	cloned := new(PodsResponse)
	*cloned = *m

	if m.Pods != nil {
		cloned.Pods = make([]*storage.Pod, len(m.Pods))
		for idx, v := range m.Pods {
			cloned.Pods[idx] = v.Clone()
		}
	}
	return cloned
}

func init() {
	proto.RegisterType((*PodsResponse)(nil), "v1.PodsResponse")
}

func init() { proto.RegisterFile("api/v1/pod_service.proto", fileDescriptor_5092c07627247a99) }

var fileDescriptor_5092c07627247a99 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xb1, 0x4e, 0xc4, 0x30,
	0x0c, 0x86, 0xe9, 0x81, 0x00, 0x85, 0x0e, 0xa7, 0x4c, 0x55, 0x41, 0xd5, 0xe9, 0x26, 0xa6, 0x94,
	0x1c, 0x23, 0x1b, 0x0b, 0x62, 0x2b, 0x65, 0x41, 0x2c, 0x28, 0x34, 0x56, 0x89, 0x38, 0xe2, 0x28,
	0x0e, 0x81, 0x5b, 0x79, 0x05, 0x16, 0x1e, 0x89, 0x11, 0x89, 0x17, 0x40, 0x85, 0x07, 0x41, 0xbd,
	0x76, 0x80, 0xd1, 0xfe, 0x7e, 0x7f, 0xb6, 0x59, 0xa6, 0x9c, 0x29, 0xa3, 0x2c, 0x1d, 0xea, 0x1b,
	0x02, 0x1f, 0x4d, 0x03, 0xc2, 0x79, 0x0c, 0xc8, 0x27, 0x51, 0xe6, 0xfb, 0x23, 0x25, 0x50, 0xbe,
	0xb9, 0xfb, 0x1f, 0xc8, 0x0f, 0x5a, 0xc4, 0x76, 0x09, 0x65, 0x9f, 0x51, 0xd6, 0x62, 0x50, 0xc1,
	0xa0, 0xa5, 0x91, 0x66, 0x14, 0xd0, 0xab, 0x16, 0x4a, 0x0d, 0x6e, 0x89, 0xab, 0x07, 0xb0, 0x61,
	0x20, 0xf3, 0x23, 0x96, 0x56, 0xa8, 0xa9, 0x06, 0x72, 0x68, 0x09, 0xf8, 0x8c, 0x6d, 0x39, 0xd4,
	0x94, 0x25, 0xb3, 0xcd, 0xc3, 0xbd, 0x45, 0x2a, 0xc6, 0x41, 0x51, 0xa1, 0xae, 0xd7, 0x64, 0x71,
	0xce, 0x58, 0x85, 0xfa, 0x72, 0xd8, 0xce, 0x4f, 0xd8, 0xce, 0x19, 0x84, 0x5e, 0xc1, 0x53, 0x11,
	0xa5, 0xa8, 0xd5, 0xd3, 0xc5, 0x23, 0xf8, 0x55, 0x3e, 0xed, 0xab, 0xbf, 0xea, 0xf9, 0xf4, 0xe5,
	0xf3, 0xe7, 0x75, 0xc2, 0xf8, 0xee, 0xf8, 0x22, 0x9d, 0x8a, 0xf7, 0xae, 0x48, 0x3e, 0xba, 0x22,
	0xf9, 0xea, 0x8a, 0xe4, 0xed, 0xbb, 0xd8, 0x60, 0x99, 0x41, 0x41, 0x41, 0x35, 0xf7, 0x1e, 0x9f,
	0x87, 0x0b, 0x85, 0x72, 0x46, 0x44, 0x79, 0x3d, 0x89, 0xf2, 0x2a, 0xb9, 0xdd, 0x5e, 0xf7, 0x8e,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xa0, 0x9b, 0x08, 0x2a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PodServiceClient is the client API for PodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
type PodServiceClient interface {
	// GetPods returns the pods.
	GetPods(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*PodsResponse, error)
}

type podServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPodServiceClient(cc grpc.ClientConnInterface) PodServiceClient {
	return &podServiceClient{cc}
}

func (c *podServiceClient) GetPods(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*PodsResponse, error) {
	out := new(PodsResponse)
	err := c.cc.Invoke(ctx, "/v1.PodService/GetPods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PodServiceServer is the server API for PodService service.
type PodServiceServer interface {
	// GetPods returns the pods.
	GetPods(context.Context, *RawQuery) (*PodsResponse, error)
}

// UnimplementedPodServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPodServiceServer struct {
}

func (*UnimplementedPodServiceServer) GetPods(ctx context.Context, req *RawQuery) (*PodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPods not implemented")
}

func RegisterPodServiceServer(s *grpc.Server, srv PodServiceServer) {
	s.RegisterService(&_PodService_serviceDesc, srv)
}

func _PodService_GetPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodServiceServer).GetPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PodService/GetPods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodServiceServer).GetPods(ctx, req.(*RawQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _PodService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PodService",
	HandlerType: (*PodServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPods",
			Handler:    _PodService_GetPods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/pod_service.proto",
}

func (m *PodsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PodsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Pods) > 0 {
		for iNdEx := len(m.Pods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPodService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPodService(dAtA []byte, offset int, v uint64) int {
	offset -= sovPodService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PodsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pods) > 0 {
		for _, e := range m.Pods {
			l = e.Size()
			n += 1 + l + sovPodService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPodService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPodService(x uint64) (n int) {
	return sovPodService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PodsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPodService
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
			return fmt.Errorf("proto: PodsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPodService
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
				return ErrInvalidLengthPodService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPodService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pods = append(m.Pods, &storage.Pod{})
			if err := m.Pods[len(m.Pods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPodService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPodService
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
func skipPodService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPodService
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
					return 0, ErrIntOverflowPodService
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
					return 0, ErrIntOverflowPodService
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
				return 0, ErrInvalidLengthPodService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPodService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPodService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPodService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPodService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPodService = fmt.Errorf("proto: unexpected end of group")
)
