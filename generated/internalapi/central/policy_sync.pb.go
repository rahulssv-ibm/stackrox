// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: internalapi/central/policy_sync.proto

package central

import (
	storage "github.com/stackrox/rox/generated/storage"
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

type PolicySync struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policies []*storage.Policy `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
}

func (x *PolicySync) Reset() {
	*x = PolicySync{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_central_policy_sync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicySync) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicySync) ProtoMessage() {}

func (x *PolicySync) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_policy_sync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicySync.ProtoReflect.Descriptor instead.
func (*PolicySync) Descriptor() ([]byte, []int) {
	return file_internalapi_central_policy_sync_proto_rawDescGZIP(), []int{0}
}

func (x *PolicySync) GetPolicies() []*storage.Policy {
	if x != nil {
		return x.Policies
	}
	return nil
}

var File_internalapi_central_policy_sync_proto protoreflect.FileDescriptor

var file_internalapi_central_policy_sync_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x79, 0x6e,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c,
	0x1a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x53, 0x79, 0x6e, 0x63, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_central_policy_sync_proto_rawDescOnce sync.Once
	file_internalapi_central_policy_sync_proto_rawDescData = file_internalapi_central_policy_sync_proto_rawDesc
)

func file_internalapi_central_policy_sync_proto_rawDescGZIP() []byte {
	file_internalapi_central_policy_sync_proto_rawDescOnce.Do(func() {
		file_internalapi_central_policy_sync_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_central_policy_sync_proto_rawDescData)
	})
	return file_internalapi_central_policy_sync_proto_rawDescData
}

var file_internalapi_central_policy_sync_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internalapi_central_policy_sync_proto_goTypes = []interface{}{
	(*PolicySync)(nil),     // 0: central.PolicySync
	(*storage.Policy)(nil), // 1: storage.Policy
}
var file_internalapi_central_policy_sync_proto_depIdxs = []int32{
	1, // 0: central.PolicySync.policies:type_name -> storage.Policy
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internalapi_central_policy_sync_proto_init() }
func file_internalapi_central_policy_sync_proto_init() {
	if File_internalapi_central_policy_sync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internalapi_central_policy_sync_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicySync); i {
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
			RawDescriptor: file_internalapi_central_policy_sync_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalapi_central_policy_sync_proto_goTypes,
		DependencyIndexes: file_internalapi_central_policy_sync_proto_depIdxs,
		MessageInfos:      file_internalapi_central_policy_sync_proto_msgTypes,
	}.Build()
	File_internalapi_central_policy_sync_proto = out.File
	file_internalapi_central_policy_sync_proto_rawDesc = nil
	file_internalapi_central_policy_sync_proto_goTypes = nil
	file_internalapi_central_policy_sync_proto_depIdxs = nil
}
