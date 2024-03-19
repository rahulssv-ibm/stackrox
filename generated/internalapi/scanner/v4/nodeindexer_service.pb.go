// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internalapi/scanner/v4/nodeindexer_service.proto

package v4

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

type CreateNodeIndexReportRequest struct {
	HashId               string   `protobuf:"bytes,1,opt,name=hash_id,json=hashId,proto3" json:"hash_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNodeIndexReportRequest) Reset()         { *m = CreateNodeIndexReportRequest{} }
func (m *CreateNodeIndexReportRequest) String() string { return proto.CompactTextString(m) }
func (*CreateNodeIndexReportRequest) ProtoMessage()    {}
func (*CreateNodeIndexReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_29583f3c47770ebe, []int{0}
}
func (m *CreateNodeIndexReportRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateNodeIndexReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateNodeIndexReportRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateNodeIndexReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNodeIndexReportRequest.Merge(m, src)
}
func (m *CreateNodeIndexReportRequest) XXX_Size() int {
	return m.Size()
}
func (m *CreateNodeIndexReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNodeIndexReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNodeIndexReportRequest proto.InternalMessageInfo

func (m *CreateNodeIndexReportRequest) GetHashId() string {
	if m != nil {
		return m.HashId
	}
	return ""
}

func (m *CreateNodeIndexReportRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *CreateNodeIndexReportRequest) Clone() *CreateNodeIndexReportRequest {
	if m == nil {
		return nil
	}
	cloned := new(CreateNodeIndexReportRequest)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterType((*CreateNodeIndexReportRequest)(nil), "scanner.v4.CreateNodeIndexReportRequest")
}

func init() {
	proto.RegisterFile("internalapi/scanner/v4/nodeindexer_service.proto", fileDescriptor_29583f3c47770ebe)
}

var fileDescriptor_29583f3c47770ebe = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xc8, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0x49, 0x2c, 0xc8, 0xd4, 0x2f, 0x4e, 0x4e, 0xcc, 0xcb, 0x4b, 0x2d, 0xd2,
	0x2f, 0x33, 0xd1, 0xcf, 0xcb, 0x4f, 0x49, 0xcd, 0xcc, 0x4b, 0x49, 0xad, 0x48, 0x2d, 0x8a, 0x2f,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0xaa,
	0xd2, 0x2b, 0x33, 0x91, 0xd2, 0xc4, 0xa1, 0x1b, 0xac, 0x33, 0xbe, 0x28, 0xb5, 0x20, 0xbf, 0xa8,
	0x04, 0xa2, 0x4d, 0xc9, 0x9c, 0x4b, 0xc6, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0xd5, 0x2f, 0x3f, 0x25,
	0xd5, 0x13, 0x24, 0x1f, 0x04, 0x96, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe7,
	0x62, 0xcf, 0x48, 0x2c, 0xce, 0x88, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62,
	0x03, 0x71, 0x3d, 0x53, 0x8c, 0x32, 0xb9, 0xb8, 0xe1, 0x5a, 0x52, 0x8b, 0x84, 0xa2, 0xb8, 0x44,
	0xb1, 0x9a, 0x23, 0xa4, 0xa1, 0x87, 0x70, 0x98, 0x1e, 0x3e, 0xab, 0xa4, 0xc4, 0x91, 0x55, 0x22,
	0xc9, 0x3b, 0x89, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c,
	0x33, 0x1e, 0xcb, 0x31, 0x44, 0x31, 0x95, 0x99, 0x24, 0xb1, 0x81, 0x3d, 0x60, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xed, 0xa6, 0x2b, 0x11, 0x2b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeIndexerClient is the client API for NodeIndexer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
type NodeIndexerClient interface {
	// CreateNodeIndexReport creates an index report for the node the container runs on and returns the report.
	CreateNodeIndexReport(ctx context.Context, in *CreateNodeIndexReportRequest, opts ...grpc.CallOption) (*IndexReport, error)
}

type nodeIndexerClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeIndexerClient(cc grpc.ClientConnInterface) NodeIndexerClient {
	return &nodeIndexerClient{cc}
}

func (c *nodeIndexerClient) CreateNodeIndexReport(ctx context.Context, in *CreateNodeIndexReportRequest, opts ...grpc.CallOption) (*IndexReport, error) {
	out := new(IndexReport)
	err := c.cc.Invoke(ctx, "/scanner.v4.NodeIndexer/CreateNodeIndexReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeIndexerServer is the server API for NodeIndexer service.
type NodeIndexerServer interface {
	// CreateNodeIndexReport creates an index report for the node the container runs on and returns the report.
	CreateNodeIndexReport(context.Context, *CreateNodeIndexReportRequest) (*IndexReport, error)
}

// UnimplementedNodeIndexerServer can be embedded to have forward compatible implementations.
type UnimplementedNodeIndexerServer struct {
}

func (*UnimplementedNodeIndexerServer) CreateNodeIndexReport(ctx context.Context, req *CreateNodeIndexReportRequest) (*IndexReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodeIndexReport not implemented")
}

func RegisterNodeIndexerServer(s *grpc.Server, srv NodeIndexerServer) {
	s.RegisterService(&_NodeIndexer_serviceDesc, srv)
}

func _NodeIndexer_CreateNodeIndexReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeIndexReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeIndexerServer).CreateNodeIndexReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scanner.v4.NodeIndexer/CreateNodeIndexReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeIndexerServer).CreateNodeIndexReport(ctx, req.(*CreateNodeIndexReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeIndexer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scanner.v4.NodeIndexer",
	HandlerType: (*NodeIndexerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNodeIndexReport",
			Handler:    _NodeIndexer_CreateNodeIndexReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internalapi/scanner/v4/nodeindexer_service.proto",
}

func (m *CreateNodeIndexReportRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateNodeIndexReportRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateNodeIndexReportRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.HashId) > 0 {
		i -= len(m.HashId)
		copy(dAtA[i:], m.HashId)
		i = encodeVarintNodeindexerService(dAtA, i, uint64(len(m.HashId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNodeindexerService(dAtA []byte, offset int, v uint64) int {
	offset -= sovNodeindexerService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreateNodeIndexReportRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HashId)
	if l > 0 {
		n += 1 + l + sovNodeindexerService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNodeindexerService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNodeindexerService(x uint64) (n int) {
	return sovNodeindexerService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateNodeIndexReportRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodeindexerService
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
			return fmt.Errorf("proto: CreateNodeIndexReportRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateNodeIndexReportRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeindexerService
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
				return ErrInvalidLengthNodeindexerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeindexerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HashId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNodeindexerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNodeindexerService
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
func skipNodeindexerService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNodeindexerService
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
					return 0, ErrIntOverflowNodeindexerService
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
					return 0, ErrIntOverflowNodeindexerService
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
				return 0, ErrInvalidLengthNodeindexerService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNodeindexerService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNodeindexerService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNodeindexerService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNodeindexerService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNodeindexerService = fmt.Errorf("proto: unexpected end of group")
)
