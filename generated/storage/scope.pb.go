// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: storage/scope.proto

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

type Scope struct {
	state protoimpl.MessageState
	Label *Scope_Label `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`

	Cluster       string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace     string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	unknownFields protoimpl.UnknownFields

	sizeCache protoimpl.SizeCache
}

func (x *Scope) Reset() {
	*x = Scope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_scope_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scope) ProtoMessage() {}

func (x *Scope) ProtoReflect() protoreflect.Message {
	mi := &file_storage_scope_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scope.ProtoReflect.Descriptor instead.
func (*Scope) Descriptor() ([]byte, []int) {
	return file_storage_scope_proto_rawDescGZIP(), []int{0}
}

func (x *Scope) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *Scope) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Scope) GetLabel() *Scope_Label {
	if x != nil {
		return x.Label
	}
	return nil
}

type Scope_Label struct {
	state protoimpl.MessageState

	Key           string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields

	sizeCache protoimpl.SizeCache
}

func (x *Scope_Label) Reset() {
	*x = Scope_Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_scope_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scope_Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scope_Label) ProtoMessage() {}

func (x *Scope_Label) ProtoReflect() protoreflect.Message {
	mi := &file_storage_scope_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scope_Label.ProtoReflect.Descriptor instead.
func (*Scope_Label) Descriptor() ([]byte, []int) {
	return file_storage_scope_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Scope_Label) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Scope_Label) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_storage_scope_proto protoreflect.FileDescriptor

var file_storage_scope_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0x9c,
	0x01, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x2a, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x1a, 0x2f, 0x0a, 0x05,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2e, 0x0a,
	0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_scope_proto_rawDescOnce sync.Once
	file_storage_scope_proto_rawDescData = file_storage_scope_proto_rawDesc
)

func file_storage_scope_proto_rawDescGZIP() []byte {
	file_storage_scope_proto_rawDescOnce.Do(func() {
		file_storage_scope_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_scope_proto_rawDescData)
	})
	return file_storage_scope_proto_rawDescData
}

var file_storage_scope_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_storage_scope_proto_goTypes = []any{
	(*Scope)(nil),       // 0: storage.Scope
	(*Scope_Label)(nil), // 1: storage.Scope.Label
}
var file_storage_scope_proto_depIdxs = []int32{
	1, // 0: storage.Scope.label:type_name -> storage.Scope.Label
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_storage_scope_proto_init() }
func file_storage_scope_proto_init() {
	if File_storage_scope_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_scope_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Scope); i {
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
		file_storage_scope_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Scope_Label); i {
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
			RawDescriptor: file_storage_scope_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_scope_proto_goTypes,
		DependencyIndexes: file_storage_scope_proto_depIdxs,
		MessageInfos:      file_storage_scope_proto_msgTypes,
	}.Build()
	File_storage_scope_proto = out.File
	file_storage_scope_proto_rawDesc = nil
	file_storage_scope_proto_goTypes = nil
	file_storage_scope_proto_depIdxs = nil
}
