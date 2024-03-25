// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: storage/orchestrator_integration.proto

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

// Next Tag: 5
type OrchestratorIntegration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Types that are assignable to IntegrationConfig:
	//
	//	*OrchestratorIntegration_Clairify
	IntegrationConfig isOrchestratorIntegration_IntegrationConfig `protobuf_oneof:"IntegrationConfig"`
}

func (x *OrchestratorIntegration) Reset() {
	*x = OrchestratorIntegration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_orchestrator_integration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrchestratorIntegration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrchestratorIntegration) ProtoMessage() {}

func (x *OrchestratorIntegration) ProtoReflect() protoreflect.Message {
	mi := &file_storage_orchestrator_integration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrchestratorIntegration.ProtoReflect.Descriptor instead.
func (*OrchestratorIntegration) Descriptor() ([]byte, []int) {
	return file_storage_orchestrator_integration_proto_rawDescGZIP(), []int{0}
}

func (x *OrchestratorIntegration) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrchestratorIntegration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrchestratorIntegration) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (m *OrchestratorIntegration) GetIntegrationConfig() isOrchestratorIntegration_IntegrationConfig {
	if m != nil {
		return m.IntegrationConfig
	}
	return nil
}

func (x *OrchestratorIntegration) GetClairify() *ClairifyConfig {
	if x, ok := x.GetIntegrationConfig().(*OrchestratorIntegration_Clairify); ok {
		return x.Clairify
	}
	return nil
}

type isOrchestratorIntegration_IntegrationConfig interface {
	isOrchestratorIntegration_IntegrationConfig()
}

type OrchestratorIntegration_Clairify struct {
	Clairify *ClairifyConfig `protobuf:"bytes,4,opt,name=clairify,proto3,oneof"`
}

func (*OrchestratorIntegration_Clairify) isOrchestratorIntegration_IntegrationConfig() {}

var File_storage_orchestrator_integration_proto protoreflect.FileDescriptor

var file_storage_orchestrator_integration_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x1a, 0x1f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x17, 0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x69, 0x72, 0x69,
	0x66, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6c, 0x61, 0x69, 0x72, 0x69, 0x66, 0x79, 0x42, 0x13, 0x0a,
	0x11, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0x27, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a,
	0x0a, 0x2e, 0x2f, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_storage_orchestrator_integration_proto_rawDescOnce sync.Once
	file_storage_orchestrator_integration_proto_rawDescData = file_storage_orchestrator_integration_proto_rawDesc
)

func file_storage_orchestrator_integration_proto_rawDescGZIP() []byte {
	file_storage_orchestrator_integration_proto_rawDescOnce.Do(func() {
		file_storage_orchestrator_integration_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_orchestrator_integration_proto_rawDescData)
	})
	return file_storage_orchestrator_integration_proto_rawDescData
}

var file_storage_orchestrator_integration_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storage_orchestrator_integration_proto_goTypes = []interface{}{
	(*OrchestratorIntegration)(nil), // 0: storage.OrchestratorIntegration
	(*ClairifyConfig)(nil),          // 1: storage.ClairifyConfig
}
var file_storage_orchestrator_integration_proto_depIdxs = []int32{
	1, // 0: storage.OrchestratorIntegration.clairify:type_name -> storage.ClairifyConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_storage_orchestrator_integration_proto_init() }
func file_storage_orchestrator_integration_proto_init() {
	if File_storage_orchestrator_integration_proto != nil {
		return
	}
	file_storage_image_integration_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_storage_orchestrator_integration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrchestratorIntegration); i {
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
	file_storage_orchestrator_integration_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OrchestratorIntegration_Clairify)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_orchestrator_integration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_orchestrator_integration_proto_goTypes,
		DependencyIndexes: file_storage_orchestrator_integration_proto_depIdxs,
		MessageInfos:      file_storage_orchestrator_integration_proto_msgTypes,
	}.Build()
	File_storage_orchestrator_integration_proto = out.File
	file_storage_orchestrator_integration_proto_rawDesc = nil
	file_storage_orchestrator_integration_proto_goTypes = nil
	file_storage_orchestrator_integration_proto_depIdxs = nil
}
