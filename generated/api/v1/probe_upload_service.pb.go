// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: api/v1/probe_upload_service.proto

package v1

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

type ProbeUploadManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*ProbeUploadManifest_File `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *ProbeUploadManifest) Reset() {
	*x = ProbeUploadManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_probe_upload_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeUploadManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeUploadManifest) ProtoMessage() {}

func (x *ProbeUploadManifest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_probe_upload_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeUploadManifest.ProtoReflect.Descriptor instead.
func (*ProbeUploadManifest) Descriptor() ([]byte, []int) {
	return file_api_v1_probe_upload_service_proto_rawDescGZIP(), []int{0}
}

func (x *ProbeUploadManifest) GetFiles() []*ProbeUploadManifest_File {
	if x != nil {
		return x.Files
	}
	return nil
}

type GetExistingProbesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilesToCheck []string `protobuf:"bytes,1,rep,name=files_to_check,json=filesToCheck,proto3" json:"files_to_check,omitempty"`
}

func (x *GetExistingProbesRequest) Reset() {
	*x = GetExistingProbesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_probe_upload_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExistingProbesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExistingProbesRequest) ProtoMessage() {}

func (x *GetExistingProbesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_probe_upload_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExistingProbesRequest.ProtoReflect.Descriptor instead.
func (*GetExistingProbesRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_probe_upload_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetExistingProbesRequest) GetFilesToCheck() []string {
	if x != nil {
		return x.FilesToCheck
	}
	return nil
}

type GetExistingProbesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExistingFiles []*ProbeUploadManifest_File `protobuf:"bytes,1,rep,name=existing_files,json=existingFiles,proto3" json:"existing_files,omitempty"`
}

func (x *GetExistingProbesResponse) Reset() {
	*x = GetExistingProbesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_probe_upload_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExistingProbesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExistingProbesResponse) ProtoMessage() {}

func (x *GetExistingProbesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_probe_upload_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExistingProbesResponse.ProtoReflect.Descriptor instead.
func (*GetExistingProbesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_probe_upload_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetExistingProbesResponse) GetExistingFiles() []*ProbeUploadManifest_File {
	if x != nil {
		return x.ExistingFiles
	}
	return nil
}

type ProbeUploadManifest_File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size  int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Crc32 uint32 `protobuf:"varint,3,opt,name=crc32,proto3" json:"crc32,omitempty"`
}

func (x *ProbeUploadManifest_File) Reset() {
	*x = ProbeUploadManifest_File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_probe_upload_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeUploadManifest_File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeUploadManifest_File) ProtoMessage() {}

func (x *ProbeUploadManifest_File) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_probe_upload_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeUploadManifest_File.ProtoReflect.Descriptor instead.
func (*ProbeUploadManifest_File) Descriptor() ([]byte, []int) {
	return file_api_v1_probe_upload_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProbeUploadManifest_File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProbeUploadManifest_File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ProbeUploadManifest_File) GetCrc32() uint32 {
	if x != nil {
		return x.Crc32
	}
	return 0
}

var File_api_v1_probe_upload_service_proto protoreflect.FileDescriptor

var file_api_v1_probe_upload_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x61, 0x6e,
	0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x1a, 0x44, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x72, 0x63, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x63, 0x72, 0x63, 0x33, 0x32, 0x22, 0x40, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x5f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x54, 0x6f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x22, 0x60, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x0d, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x32, 0x8b, 0x01, 0x0a, 0x12,
	0x50, 0x72, 0x6f, 0x62, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x75, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x1b, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x67, 0x65,
	0x74, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x21, 0x0a, 0x18, 0x69, 0x6f, 0x2e,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x58, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_probe_upload_service_proto_rawDescOnce sync.Once
	file_api_v1_probe_upload_service_proto_rawDescData = file_api_v1_probe_upload_service_proto_rawDesc
)

func file_api_v1_probe_upload_service_proto_rawDescGZIP() []byte {
	file_api_v1_probe_upload_service_proto_rawDescOnce.Do(func() {
		file_api_v1_probe_upload_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_probe_upload_service_proto_rawDescData)
	})
	return file_api_v1_probe_upload_service_proto_rawDescData
}

var file_api_v1_probe_upload_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_v1_probe_upload_service_proto_goTypes = []interface{}{
	(*ProbeUploadManifest)(nil),       // 0: v1.ProbeUploadManifest
	(*GetExistingProbesRequest)(nil),  // 1: v1.GetExistingProbesRequest
	(*GetExistingProbesResponse)(nil), // 2: v1.GetExistingProbesResponse
	(*ProbeUploadManifest_File)(nil),  // 3: v1.ProbeUploadManifest.File
}
var file_api_v1_probe_upload_service_proto_depIdxs = []int32{
	3, // 0: v1.ProbeUploadManifest.files:type_name -> v1.ProbeUploadManifest.File
	3, // 1: v1.GetExistingProbesResponse.existing_files:type_name -> v1.ProbeUploadManifest.File
	1, // 2: v1.ProbeUploadService.GetExistingProbes:input_type -> v1.GetExistingProbesRequest
	2, // 3: v1.ProbeUploadService.GetExistingProbes:output_type -> v1.GetExistingProbesResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_probe_upload_service_proto_init() }
func file_api_v1_probe_upload_service_proto_init() {
	if File_api_v1_probe_upload_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_probe_upload_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeUploadManifest); i {
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
		file_api_v1_probe_upload_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExistingProbesRequest); i {
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
		file_api_v1_probe_upload_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExistingProbesResponse); i {
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
		file_api_v1_probe_upload_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeUploadManifest_File); i {
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
			RawDescriptor: file_api_v1_probe_upload_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_probe_upload_service_proto_goTypes,
		DependencyIndexes: file_api_v1_probe_upload_service_proto_depIdxs,
		MessageInfos:      file_api_v1_probe_upload_service_proto_msgTypes,
	}.Build()
	File_api_v1_probe_upload_service_proto = out.File
	file_api_v1_probe_upload_service_proto_rawDesc = nil
	file_api_v1_probe_upload_service_proto_goTypes = nil
	file_api_v1_probe_upload_service_proto_depIdxs = nil
}
