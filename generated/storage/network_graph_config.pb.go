// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: storage/network_graph_config.proto

package storage

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

type NetworkGraphConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" sql:"pk"` // @gotags: sql:"pk"
	HideDefaultExternalSrcs bool   `protobuf:"varint,1,opt,name=hide_default_external_srcs,json=hideDefaultExternalSrcs,proto3" json:"hide_default_external_srcs,omitempty"`
}

func (x *NetworkGraphConfig) Reset() {
	*x = NetworkGraphConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_network_graph_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkGraphConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkGraphConfig) ProtoMessage() {}

func (x *NetworkGraphConfig) ProtoReflect() protoreflect.Message {
	mi := &file_storage_network_graph_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkGraphConfig.ProtoReflect.Descriptor instead.
func (*NetworkGraphConfig) Descriptor() ([]byte, []int) {
	return file_storage_network_graph_config_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkGraphConfig) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NetworkGraphConfig) GetHideDefaultExternalSrcs() bool {
	if x != nil {
		return x.HideDefaultExternalSrcs
	}
	return false
}

var File_storage_network_graph_config_proto protoreflect.FileDescriptor

var file_storage_network_graph_config_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0x61, 0x0a,
	0x12, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x72, 0x61, 0x70, 0x68, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x1a, 0x68, 0x69, 0x64, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x72, 0x63,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x68, 0x69, 0x64, 0x65, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x72, 0x63, 0x73,
	0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_network_graph_config_proto_rawDescOnce sync.Once
	file_storage_network_graph_config_proto_rawDescData = file_storage_network_graph_config_proto_rawDesc
)

func file_storage_network_graph_config_proto_rawDescGZIP() []byte {
	file_storage_network_graph_config_proto_rawDescOnce.Do(func() {
		file_storage_network_graph_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_network_graph_config_proto_rawDescData)
	})
	return file_storage_network_graph_config_proto_rawDescData
}

var file_storage_network_graph_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storage_network_graph_config_proto_goTypes = []any{
	(*NetworkGraphConfig)(nil), // 0: storage.NetworkGraphConfig
}
var file_storage_network_graph_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_storage_network_graph_config_proto_init() }
func file_storage_network_graph_config_proto_init() {
	if File_storage_network_graph_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_network_graph_config_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NetworkGraphConfig); i {
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
			RawDescriptor: file_storage_network_graph_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_network_graph_config_proto_goTypes,
		DependencyIndexes: file_storage_network_graph_config_proto_depIdxs,
		MessageInfos:      file_storage_network_graph_config_proto_msgTypes,
	}.Build()
	File_storage_network_graph_config_proto = out.File
	file_storage_network_graph_config_proto_rawDesc = nil
	file_storage_network_graph_config_proto_goTypes = nil
	file_storage_network_graph_config_proto_depIdxs = nil
}
