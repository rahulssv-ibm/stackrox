// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: internalapi/scanner/v4/nodeindexer_service.proto

package v4

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateNodeIndexReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateNodeIndexReportRequest) Reset() {
	*x = CreateNodeIndexReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_scanner_v4_nodeindexer_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeIndexReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeIndexReportRequest) ProtoMessage() {}

func (x *CreateNodeIndexReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_scanner_v4_nodeindexer_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeIndexReportRequest.ProtoReflect.Descriptor instead.
func (*CreateNodeIndexReportRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_scanner_v4_nodeindexer_service_proto_rawDescGZIP(), []int{0}
}

var File_internalapi_scanner_v4_nodeindexer_service_proto protoreflect.FileDescriptor

var file_internalapi_scanner_v4_nodeindexer_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x1a, 0x29,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x1c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x69, 0x0a, 0x0b, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x12, 0x5a, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x28, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x34,
	0x3b, 0x76, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_scanner_v4_nodeindexer_service_proto_rawDescOnce sync.Once
	file_internalapi_scanner_v4_nodeindexer_service_proto_rawDescData = file_internalapi_scanner_v4_nodeindexer_service_proto_rawDesc
)

func file_internalapi_scanner_v4_nodeindexer_service_proto_rawDescGZIP() []byte {
	file_internalapi_scanner_v4_nodeindexer_service_proto_rawDescOnce.Do(func() {
		file_internalapi_scanner_v4_nodeindexer_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_scanner_v4_nodeindexer_service_proto_rawDescData)
	})
	return file_internalapi_scanner_v4_nodeindexer_service_proto_rawDescData
}

var file_internalapi_scanner_v4_nodeindexer_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internalapi_scanner_v4_nodeindexer_service_proto_goTypes = []any{
	(*CreateNodeIndexReportRequest)(nil), // 0: scanner.v4.CreateNodeIndexReportRequest
	(*IndexReport)(nil),                  // 1: scanner.v4.IndexReport
}
var file_internalapi_scanner_v4_nodeindexer_service_proto_depIdxs = []int32{
	0, // 0: scanner.v4.NodeIndexer.CreateNodeIndexReport:input_type -> scanner.v4.CreateNodeIndexReportRequest
	1, // 1: scanner.v4.NodeIndexer.CreateNodeIndexReport:output_type -> scanner.v4.IndexReport
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internalapi_scanner_v4_nodeindexer_service_proto_init() }
func file_internalapi_scanner_v4_nodeindexer_service_proto_init() {
	if File_internalapi_scanner_v4_nodeindexer_service_proto != nil {
		return
	}
	file_internalapi_scanner_v4_index_report_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internalapi_scanner_v4_nodeindexer_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateNodeIndexReportRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_scanner_v4_nodeindexer_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internalapi_scanner_v4_nodeindexer_service_proto_goTypes,
		DependencyIndexes: file_internalapi_scanner_v4_nodeindexer_service_proto_depIdxs,
		MessageInfos:      file_internalapi_scanner_v4_nodeindexer_service_proto_msgTypes,
	}.Build()
	File_internalapi_scanner_v4_nodeindexer_service_proto = out.File
	file_internalapi_scanner_v4_nodeindexer_service_proto_rawDesc = nil
	file_internalapi_scanner_v4_nodeindexer_service_proto_goTypes = nil
	file_internalapi_scanner_v4_nodeindexer_service_proto_depIdxs = nil
}
