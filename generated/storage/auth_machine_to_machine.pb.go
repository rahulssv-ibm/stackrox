// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: storage/auth_machine_to_machine.proto

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

type AuthMachineToMachineConfig_Type int32

const (
	AuthMachineToMachineConfig_GENERIC        AuthMachineToMachineConfig_Type = 0
	AuthMachineToMachineConfig_GITHUB_ACTIONS AuthMachineToMachineConfig_Type = 1
)

// Enum value maps for AuthMachineToMachineConfig_Type.
var (
	AuthMachineToMachineConfig_Type_name = map[int32]string{
		0: "GENERIC",
		1: "GITHUB_ACTIONS",
	}
	AuthMachineToMachineConfig_Type_value = map[string]int32{
		"GENERIC":        0,
		"GITHUB_ACTIONS": 1,
	}
)

func (x AuthMachineToMachineConfig_Type) Enum() *AuthMachineToMachineConfig_Type {
	p := new(AuthMachineToMachineConfig_Type)
	*p = x
	return p
}

func (x AuthMachineToMachineConfig_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthMachineToMachineConfig_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_auth_machine_to_machine_proto_enumTypes[0].Descriptor()
}

func (AuthMachineToMachineConfig_Type) Type() protoreflect.EnumType {
	return &file_storage_auth_machine_to_machine_proto_enumTypes[0]
}

func (x AuthMachineToMachineConfig_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthMachineToMachineConfig_Type.Descriptor instead.
func (AuthMachineToMachineConfig_Type) EnumDescriptor() ([]byte, []int) {
	return file_storage_auth_machine_to_machine_proto_rawDescGZIP(), []int{0, 0}
}

// AuthMachineToMachineConfig is the storage representation of auth machine to machine configs in Central.
//
// Refer to v1.AuthMachineToMachineConfig for a more detailed doc.
// Next tag: 6.
type AuthMachineToMachineConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                      string                                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,type(uuid)"` // @gotags: sql:"pk,type(uuid)"
	Type                    AuthMachineToMachineConfig_Type       `protobuf:"varint,2,opt,name=type,proto3,enum=storage.AuthMachineToMachineConfig_Type" json:"type,omitempty"`
	TokenExpirationDuration string                                `protobuf:"bytes,3,opt,name=token_expiration_duration,json=tokenExpirationDuration,proto3" json:"token_expiration_duration,omitempty"`
	Mappings                []*AuthMachineToMachineConfig_Mapping `protobuf:"bytes,4,rep,name=mappings,proto3" json:"mappings,omitempty"`
	// The issuer is related to an ID token's issuer.
	// Spec: https://openid.net/specs/openid-connect-core-1_0.html#IDToken.
	Issuer string `protobuf:"bytes,5,opt,name=issuer,proto3" json:"issuer,omitempty" sql:"unique"` // @gotags: sql:"unique"
}

func (x *AuthMachineToMachineConfig) Reset() {
	*x = AuthMachineToMachineConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_auth_machine_to_machine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthMachineToMachineConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthMachineToMachineConfig) ProtoMessage() {}

func (x *AuthMachineToMachineConfig) ProtoReflect() protoreflect.Message {
	mi := &file_storage_auth_machine_to_machine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthMachineToMachineConfig.ProtoReflect.Descriptor instead.
func (*AuthMachineToMachineConfig) Descriptor() ([]byte, []int) {
	return file_storage_auth_machine_to_machine_proto_rawDescGZIP(), []int{0}
}

func (x *AuthMachineToMachineConfig) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthMachineToMachineConfig) GetType() AuthMachineToMachineConfig_Type {
	if x != nil {
		return x.Type
	}
	return AuthMachineToMachineConfig_GENERIC
}

func (x *AuthMachineToMachineConfig) GetTokenExpirationDuration() string {
	if x != nil {
		return x.TokenExpirationDuration
	}
	return ""
}

func (x *AuthMachineToMachineConfig) GetMappings() []*AuthMachineToMachineConfig_Mapping {
	if x != nil {
		return x.Mappings
	}
	return nil
}

func (x *AuthMachineToMachineConfig) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

type AuthMachineToMachineConfig_Mapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key             string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ValueExpression string `protobuf:"bytes,2,opt,name=value_expression,json=valueExpression,proto3" json:"value_expression,omitempty"`
	Role            string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty" sql:"fk(Role:name),restrict-delete"` // @gotags: sql:"fk(Role:name),restrict-delete"
}

func (x *AuthMachineToMachineConfig_Mapping) Reset() {
	*x = AuthMachineToMachineConfig_Mapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_auth_machine_to_machine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthMachineToMachineConfig_Mapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthMachineToMachineConfig_Mapping) ProtoMessage() {}

func (x *AuthMachineToMachineConfig_Mapping) ProtoReflect() protoreflect.Message {
	mi := &file_storage_auth_machine_to_machine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthMachineToMachineConfig_Mapping.ProtoReflect.Descriptor instead.
func (*AuthMachineToMachineConfig_Mapping) Descriptor() ([]byte, []int) {
	return file_storage_auth_machine_to_machine_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AuthMachineToMachineConfig_Mapping) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AuthMachineToMachineConfig_Mapping) GetValueExpression() string {
	if x != nil {
		return x.ValueExpression
	}
	return ""
}

func (x *AuthMachineToMachineConfig_Mapping) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_storage_auth_machine_to_machine_proto protoreflect.FileDescriptor

var file_storage_auth_machine_to_machine_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x22, 0x8c, 0x03, 0x0a, 0x1a, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x54, 0x6f, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x54, 0x6f, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a,
	0x19, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x17, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x08, 0x6d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x54, 0x6f, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x1a, 0x5a, 0x0a, 0x07, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x27, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x47,
	0x49, 0x54, 0x48, 0x55, 0x42, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x01, 0x42,
	0x27, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x0a, 0x2e, 0x2f,
	0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_auth_machine_to_machine_proto_rawDescOnce sync.Once
	file_storage_auth_machine_to_machine_proto_rawDescData = file_storage_auth_machine_to_machine_proto_rawDesc
)

func file_storage_auth_machine_to_machine_proto_rawDescGZIP() []byte {
	file_storage_auth_machine_to_machine_proto_rawDescOnce.Do(func() {
		file_storage_auth_machine_to_machine_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_auth_machine_to_machine_proto_rawDescData)
	})
	return file_storage_auth_machine_to_machine_proto_rawDescData
}

var file_storage_auth_machine_to_machine_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_storage_auth_machine_to_machine_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_storage_auth_machine_to_machine_proto_goTypes = []interface{}{
	(AuthMachineToMachineConfig_Type)(0),       // 0: storage.AuthMachineToMachineConfig.Type
	(*AuthMachineToMachineConfig)(nil),         // 1: storage.AuthMachineToMachineConfig
	(*AuthMachineToMachineConfig_Mapping)(nil), // 2: storage.AuthMachineToMachineConfig.Mapping
}
var file_storage_auth_machine_to_machine_proto_depIdxs = []int32{
	0, // 0: storage.AuthMachineToMachineConfig.type:type_name -> storage.AuthMachineToMachineConfig.Type
	2, // 1: storage.AuthMachineToMachineConfig.mappings:type_name -> storage.AuthMachineToMachineConfig.Mapping
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_storage_auth_machine_to_machine_proto_init() }
func file_storage_auth_machine_to_machine_proto_init() {
	if File_storage_auth_machine_to_machine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_auth_machine_to_machine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthMachineToMachineConfig); i {
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
		file_storage_auth_machine_to_machine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthMachineToMachineConfig_Mapping); i {
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
			RawDescriptor: file_storage_auth_machine_to_machine_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_auth_machine_to_machine_proto_goTypes,
		DependencyIndexes: file_storage_auth_machine_to_machine_proto_depIdxs,
		EnumInfos:         file_storage_auth_machine_to_machine_proto_enumTypes,
		MessageInfos:      file_storage_auth_machine_to_machine_proto_msgTypes,
	}.Build()
	File_storage_auth_machine_to_machine_proto = out.File
	file_storage_auth_machine_to_machine_proto_rawDesc = nil
	file_storage_auth_machine_to_machine_proto_goTypes = nil
	file_storage_auth_machine_to_machine_proto_depIdxs = nil
}
