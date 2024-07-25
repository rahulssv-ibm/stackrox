// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: storage/service_identity.proto

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

// Next available tag: 16
type ServiceType int32

const (
	ServiceType_UNKNOWN_SERVICE            ServiceType = 0
	ServiceType_SENSOR_SERVICE             ServiceType = 1
	ServiceType_CENTRAL_SERVICE            ServiceType = 2
	ServiceType_CENTRAL_DB_SERVICE         ServiceType = 12
	ServiceType_REMOTE_SERVICE             ServiceType = 3
	ServiceType_COLLECTOR_SERVICE          ServiceType = 4
	ServiceType_MONITORING_UI_SERVICE      ServiceType = 5
	ServiceType_MONITORING_DB_SERVICE      ServiceType = 6
	ServiceType_MONITORING_CLIENT_SERVICE  ServiceType = 7
	ServiceType_BENCHMARK_SERVICE          ServiceType = 8
	ServiceType_SCANNER_SERVICE            ServiceType = 9
	ServiceType_SCANNER_DB_SERVICE         ServiceType = 10
	ServiceType_ADMISSION_CONTROL_SERVICE  ServiceType = 11
	ServiceType_SCANNER_V4_INDEXER_SERVICE ServiceType = 13
	ServiceType_SCANNER_V4_MATCHER_SERVICE ServiceType = 14
	ServiceType_SCANNER_V4_DB_SERVICE      ServiceType = 15
)

// Enum value maps for ServiceType.
var (
	ServiceType_name = map[int32]string{
		0:  "UNKNOWN_SERVICE",
		1:  "SENSOR_SERVICE",
		2:  "CENTRAL_SERVICE",
		12: "CENTRAL_DB_SERVICE",
		3:  "REMOTE_SERVICE",
		4:  "COLLECTOR_SERVICE",
		5:  "MONITORING_UI_SERVICE",
		6:  "MONITORING_DB_SERVICE",
		7:  "MONITORING_CLIENT_SERVICE",
		8:  "BENCHMARK_SERVICE",
		9:  "SCANNER_SERVICE",
		10: "SCANNER_DB_SERVICE",
		11: "ADMISSION_CONTROL_SERVICE",
		13: "SCANNER_V4_INDEXER_SERVICE",
		14: "SCANNER_V4_MATCHER_SERVICE",
		15: "SCANNER_V4_DB_SERVICE",
	}
	ServiceType_value = map[string]int32{
		"UNKNOWN_SERVICE":            0,
		"SENSOR_SERVICE":             1,
		"CENTRAL_SERVICE":            2,
		"CENTRAL_DB_SERVICE":         12,
		"REMOTE_SERVICE":             3,
		"COLLECTOR_SERVICE":          4,
		"MONITORING_UI_SERVICE":      5,
		"MONITORING_DB_SERVICE":      6,
		"MONITORING_CLIENT_SERVICE":  7,
		"BENCHMARK_SERVICE":          8,
		"SCANNER_SERVICE":            9,
		"SCANNER_DB_SERVICE":         10,
		"ADMISSION_CONTROL_SERVICE":  11,
		"SCANNER_V4_INDEXER_SERVICE": 13,
		"SCANNER_V4_MATCHER_SERVICE": 14,
		"SCANNER_V4_DB_SERVICE":      15,
	}
)

func (x ServiceType) Enum() *ServiceType {
	p := new(ServiceType)
	*p = x
	return p
}

func (x ServiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_service_identity_proto_enumTypes[0].Descriptor()
}

func (ServiceType) Type() protoreflect.EnumType {
	return &file_storage_service_identity_proto_enumTypes[0]
}

func (x ServiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceType.Descriptor instead.
func (ServiceType) EnumDescriptor() ([]byte, []int) {
	return file_storage_service_identity_proto_rawDescGZIP(), []int{0}
}

type ServiceIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialStr string `protobuf:"bytes,4,opt,name=serial_str,json=serialStr,proto3" json:"serial_str,omitempty" sql:"pk"` // The serial number in decimal representation. // @gotags: sql:"pk"
	// Types that are assignable to Srl:
	//
	//	*ServiceIdentity_Serial
	Srl          isServiceIdentity_Srl `protobuf_oneof:"srl"`
	Id           string                `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Type         ServiceType           `protobuf:"varint,3,opt,name=type,proto3,enum=storage.ServiceType" json:"type,omitempty"`
	InitBundleId string                `protobuf:"bytes,5,opt,name=init_bundle_id,json=initBundleId,proto3" json:"init_bundle_id,omitempty"`
}

func (x *ServiceIdentity) Reset() {
	*x = ServiceIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_service_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceIdentity) ProtoMessage() {}

func (x *ServiceIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_storage_service_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceIdentity.ProtoReflect.Descriptor instead.
func (*ServiceIdentity) Descriptor() ([]byte, []int) {
	return file_storage_service_identity_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceIdentity) GetSerialStr() string {
	if x != nil {
		return x.SerialStr
	}
	return ""
}

func (m *ServiceIdentity) GetSrl() isServiceIdentity_Srl {
	if m != nil {
		return m.Srl
	}
	return nil
}

// Deprecated: Marked as deprecated in storage/service_identity.proto.
func (x *ServiceIdentity) GetSerial() int64 {
	if x, ok := x.GetSrl().(*ServiceIdentity_Serial); ok {
		return x.Serial
	}
	return 0
}

func (x *ServiceIdentity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ServiceIdentity) GetType() ServiceType {
	if x != nil {
		return x.Type
	}
	return ServiceType_UNKNOWN_SERVICE
}

func (x *ServiceIdentity) GetInitBundleId() string {
	if x != nil {
		return x.InitBundleId
	}
	return ""
}

type isServiceIdentity_Srl interface {
	isServiceIdentity_Srl()
}

type ServiceIdentity_Serial struct {
	// Deprecated: Marked as deprecated in storage/service_identity.proto.
	Serial int64 `protobuf:"varint,1,opt,name=serial,proto3,oneof"`
}

func (*ServiceIdentity_Serial) isServiceIdentity_Srl() {}

type ServiceCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CertPem []byte `protobuf:"bytes,1,opt,name=cert_pem,json=certPem,proto3" json:"cert_pem,omitempty"`
	KeyPem  []byte `protobuf:"bytes,2,opt,name=key_pem,json=keyPem,proto3" json:"key_pem,omitempty"`
}

func (x *ServiceCertificate) Reset() {
	*x = ServiceCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_service_identity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceCertificate) ProtoMessage() {}

func (x *ServiceCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_storage_service_identity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceCertificate.ProtoReflect.Descriptor instead.
func (*ServiceCertificate) Descriptor() ([]byte, []int) {
	return file_storage_service_identity_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceCertificate) GetCertPem() []byte {
	if x != nil {
		return x.CertPem
	}
	return nil
}

func (x *ServiceCertificate) GetKeyPem() []byte {
	if x != nil {
		return x.KeyPem
	}
	return nil
}

type TypedServiceCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceType ServiceType         `protobuf:"varint,1,opt,name=service_type,json=serviceType,proto3,enum=storage.ServiceType" json:"service_type,omitempty"`
	Cert        *ServiceCertificate `protobuf:"bytes,2,opt,name=cert,proto3" json:"cert,omitempty"`
}

func (x *TypedServiceCertificate) Reset() {
	*x = TypedServiceCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_service_identity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypedServiceCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypedServiceCertificate) ProtoMessage() {}

func (x *TypedServiceCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_storage_service_identity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypedServiceCertificate.ProtoReflect.Descriptor instead.
func (*TypedServiceCertificate) Descriptor() ([]byte, []int) {
	return file_storage_service_identity_proto_rawDescGZIP(), []int{2}
}

func (x *TypedServiceCertificate) GetServiceType() ServiceType {
	if x != nil {
		return x.ServiceType
	}
	return ServiceType_UNKNOWN_SERVICE
}

func (x *TypedServiceCertificate) GetCert() *ServiceCertificate {
	if x != nil {
		return x.Cert
	}
	return nil
}

type TypedServiceCertificateSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaPem        []byte                     `protobuf:"bytes,1,opt,name=ca_pem,json=caPem,proto3" json:"ca_pem,omitempty"`
	ServiceCerts []*TypedServiceCertificate `protobuf:"bytes,2,rep,name=service_certs,json=serviceCerts,proto3" json:"service_certs,omitempty"`
}

func (x *TypedServiceCertificateSet) Reset() {
	*x = TypedServiceCertificateSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_service_identity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypedServiceCertificateSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypedServiceCertificateSet) ProtoMessage() {}

func (x *TypedServiceCertificateSet) ProtoReflect() protoreflect.Message {
	mi := &file_storage_service_identity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypedServiceCertificateSet.ProtoReflect.Descriptor instead.
func (*TypedServiceCertificateSet) Descriptor() ([]byte, []int) {
	return file_storage_service_identity_proto_rawDescGZIP(), []int{3}
}

func (x *TypedServiceCertificateSet) GetCaPem() []byte {
	if x != nil {
		return x.CaPem
	}
	return nil
}

func (x *TypedServiceCertificateSet) GetServiceCerts() []*TypedServiceCertificate {
	if x != nil {
		return x.ServiceCerts
	}
	return nil
}

var File_storage_service_identity_proto protoreflect.FileDescriptor

var file_storage_service_identity_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x0f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x12, 0x1c, 0x0a, 0x06,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x18, 0x01,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x62, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e,
	0x69, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x49, 0x64, 0x42, 0x05, 0x0a, 0x03, 0x73, 0x72,
	0x6c, 0x22, 0x48, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x65, 0x72, 0x74, 0x5f,
	0x70, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x65, 0x72, 0x74, 0x50,
	0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x65, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x6b, 0x65, 0x79, 0x50, 0x65, 0x6d, 0x22, 0x83, 0x01, 0x0a, 0x17,
	0x54, 0x79, 0x70, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2f, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x04, 0x63, 0x65, 0x72,
	0x74, 0x22, 0x7a, 0x0a, 0x1a, 0x54, 0x79, 0x70, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x63, 0x61, 0x5f, 0x70, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x63, 0x61, 0x50, 0x65, 0x6d, 0x12, 0x45, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x65, 0x72, 0x74, 0x73, 0x2a, 0xa1, 0x03,
	0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a,
	0x0f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x4e, 0x53, 0x4f, 0x52, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x45, 0x4e, 0x54, 0x52, 0x41,
	0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x43,
	0x45, 0x4e, 0x54, 0x52, 0x41, 0x4c, 0x5f, 0x44, 0x42, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x10, 0x0c, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4c, 0x4c, 0x45,
	0x43, 0x54, 0x4f, 0x52, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x04, 0x12, 0x19,
	0x0a, 0x15, 0x4d, 0x4f, 0x4e, 0x49, 0x54, 0x4f, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x49, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f, 0x4e,
	0x49, 0x54, 0x4f, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x42, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x10, 0x06, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x4f, 0x4e, 0x49, 0x54, 0x4f, 0x52, 0x49,
	0x4e, 0x47, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x45, 0x4e, 0x43, 0x48, 0x4d, 0x41, 0x52, 0x4b,
	0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x43,
	0x41, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x09, 0x12,
	0x16, 0x0a, 0x12, 0x53, 0x43, 0x41, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x44, 0x42, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x0a, 0x12, 0x1d, 0x0a, 0x19, 0x41, 0x44, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x10, 0x0b, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x43, 0x41, 0x4e, 0x4e, 0x45,
	0x52, 0x5f, 0x56, 0x34, 0x5f, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x10, 0x0d, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x43, 0x41, 0x4e, 0x4e, 0x45,
	0x52, 0x5f, 0x56, 0x34, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x10, 0x0e, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x43, 0x41, 0x4e, 0x4e, 0x45,
	0x52, 0x5f, 0x56, 0x34, 0x5f, 0x44, 0x42, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10,
	0x0f, 0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x11,
	0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_service_identity_proto_rawDescOnce sync.Once
	file_storage_service_identity_proto_rawDescData = file_storage_service_identity_proto_rawDesc
)

func file_storage_service_identity_proto_rawDescGZIP() []byte {
	file_storage_service_identity_proto_rawDescOnce.Do(func() {
		file_storage_service_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_service_identity_proto_rawDescData)
	})
	return file_storage_service_identity_proto_rawDescData
}

var file_storage_service_identity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_storage_service_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_storage_service_identity_proto_goTypes = []any{
	(ServiceType)(0),                   // 0: storage.ServiceType
	(*ServiceIdentity)(nil),            // 1: storage.ServiceIdentity
	(*ServiceCertificate)(nil),         // 2: storage.ServiceCertificate
	(*TypedServiceCertificate)(nil),    // 3: storage.TypedServiceCertificate
	(*TypedServiceCertificateSet)(nil), // 4: storage.TypedServiceCertificateSet
}
var file_storage_service_identity_proto_depIdxs = []int32{
	0, // 0: storage.ServiceIdentity.type:type_name -> storage.ServiceType
	0, // 1: storage.TypedServiceCertificate.service_type:type_name -> storage.ServiceType
	2, // 2: storage.TypedServiceCertificate.cert:type_name -> storage.ServiceCertificate
	3, // 3: storage.TypedServiceCertificateSet.service_certs:type_name -> storage.TypedServiceCertificate
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_storage_service_identity_proto_init() }
func file_storage_service_identity_proto_init() {
	if File_storage_service_identity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_service_identity_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ServiceIdentity); i {
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
		file_storage_service_identity_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ServiceCertificate); i {
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
		file_storage_service_identity_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TypedServiceCertificate); i {
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
		file_storage_service_identity_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TypedServiceCertificateSet); i {
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
	file_storage_service_identity_proto_msgTypes[0].OneofWrappers = []any{
		(*ServiceIdentity_Serial)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_service_identity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_service_identity_proto_goTypes,
		DependencyIndexes: file_storage_service_identity_proto_depIdxs,
		EnumInfos:         file_storage_service_identity_proto_enumTypes,
		MessageInfos:      file_storage_service_identity_proto_msgTypes,
	}.Build()
	File_storage_service_identity_proto = out.File
	file_storage_service_identity_proto_rawDesc = nil
	file_storage_service_identity_proto_goTypes = nil
	file_storage_service_identity_proto_depIdxs = nil
}
