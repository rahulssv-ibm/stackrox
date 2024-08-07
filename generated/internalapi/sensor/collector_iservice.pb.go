// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: internalapi/sensor/collector_iservice.proto

package sensor

import (
	common "github.com/stackrox/rox/generated/internalapi/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MsgToCollector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//
	//	*MsgToCollector_ConfigWithCluster
	Msg isMsgToCollector_Msg `protobuf_oneof:"msg"`
}

func (x *MsgToCollector) Reset() {
	*x = MsgToCollector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_sensor_collector_iservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgToCollector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgToCollector) ProtoMessage() {}

func (x *MsgToCollector) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_collector_iservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgToCollector.ProtoReflect.Descriptor instead.
func (*MsgToCollector) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_collector_iservice_proto_rawDescGZIP(), []int{0}
}

func (m *MsgToCollector) GetMsg() isMsgToCollector_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *MsgToCollector) GetConfigWithCluster() *common.CollectorRuntimeConfig {
	if x, ok := x.GetMsg().(*MsgToCollector_ConfigWithCluster); ok {
		return x.ConfigWithCluster
	}
	return nil
}

type isMsgToCollector_Msg interface {
	isMsgToCollector_Msg()
}

type MsgToCollector_ConfigWithCluster struct {
	ConfigWithCluster *common.CollectorRuntimeConfig `protobuf:"bytes,1,opt,name=config_with_cluster,json=configWithCluster,proto3,oneof"`
}

func (*MsgToCollector_ConfigWithCluster) isMsgToCollector_Msg() {}

var File_internalapi_sensor_collector_iservice_proto protoreflect.FileDescriptor

var file_internalapi_sensor_collector_iservice_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x31, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0e, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x50, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x57, 0x69,
	0x74, 0x68, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x32, 0x53, 0x0a, 0x10, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x30, 0x01, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x3b, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_sensor_collector_iservice_proto_rawDescOnce sync.Once
	file_internalapi_sensor_collector_iservice_proto_rawDescData = file_internalapi_sensor_collector_iservice_proto_rawDesc
)

func file_internalapi_sensor_collector_iservice_proto_rawDescGZIP() []byte {
	file_internalapi_sensor_collector_iservice_proto_rawDescOnce.Do(func() {
		file_internalapi_sensor_collector_iservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_sensor_collector_iservice_proto_rawDescData)
	})
	return file_internalapi_sensor_collector_iservice_proto_rawDescData
}

var file_internalapi_sensor_collector_iservice_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internalapi_sensor_collector_iservice_proto_goTypes = []any{
	(*MsgToCollector)(nil),                // 0: sensor.MsgToCollector
	(*common.CollectorRuntimeConfig)(nil), // 1: common.CollectorRuntimeConfig
	(*emptypb.Empty)(nil),                 // 2: google.protobuf.Empty
}
var file_internalapi_sensor_collector_iservice_proto_depIdxs = []int32{
	1, // 0: sensor.MsgToCollector.config_with_cluster:type_name -> common.CollectorRuntimeConfig
	2, // 1: sensor.CollectorService.Communicate:input_type -> google.protobuf.Empty
	0, // 2: sensor.CollectorService.Communicate:output_type -> sensor.MsgToCollector
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internalapi_sensor_collector_iservice_proto_init() }
func file_internalapi_sensor_collector_iservice_proto_init() {
	if File_internalapi_sensor_collector_iservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internalapi_sensor_collector_iservice_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MsgToCollector); i {
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
	file_internalapi_sensor_collector_iservice_proto_msgTypes[0].OneofWrappers = []any{
		(*MsgToCollector_ConfigWithCluster)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_sensor_collector_iservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internalapi_sensor_collector_iservice_proto_goTypes,
		DependencyIndexes: file_internalapi_sensor_collector_iservice_proto_depIdxs,
		MessageInfos:      file_internalapi_sensor_collector_iservice_proto_msgTypes,
	}.Build()
	File_internalapi_sensor_collector_iservice_proto = out.File
	file_internalapi_sensor_collector_iservice_proto_rawDesc = nil
	file_internalapi_sensor_collector_iservice_proto_goTypes = nil
	file_internalapi_sensor_collector_iservice_proto_depIdxs = nil
}
