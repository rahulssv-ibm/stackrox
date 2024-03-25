// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: internalapi/sensor/admission_control.proto

package sensor

import (
	central "github.com/stackrox/rox/generated/internalapi/central"
	storage "github.com/stackrox/rox/generated/storage"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AdmissionControlSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterConfig              *storage.DynamicClusterConfig `protobuf:"bytes,1,opt,name=cluster_config,json=clusterConfig,proto3" json:"cluster_config,omitempty"`
	EnforcedDeployTimePolicies *storage.PolicyList           `protobuf:"bytes,2,opt,name=enforced_deploy_time_policies,json=enforcedDeployTimePolicies,proto3" json:"enforced_deploy_time_policies,omitempty"`
	Timestamp                  *timestamppb.Timestamp        `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	CacheVersion               string                        `protobuf:"bytes,4,opt,name=cache_version,json=cacheVersion,proto3" json:"cache_version,omitempty"`
	CentralEndpoint            string                        `protobuf:"bytes,5,opt,name=central_endpoint,json=centralEndpoint,proto3" json:"central_endpoint,omitempty"`
	ClusterId                  string                        `protobuf:"bytes,6,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	RuntimePolicies            *storage.PolicyList           `protobuf:"bytes,7,opt,name=runtime_policies,json=runtimePolicies,proto3" json:"runtime_policies,omitempty"`
}

func (x *AdmissionControlSettings) Reset() {
	*x = AdmissionControlSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_sensor_admission_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdmissionControlSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdmissionControlSettings) ProtoMessage() {}

func (x *AdmissionControlSettings) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_admission_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdmissionControlSettings.ProtoReflect.Descriptor instead.
func (*AdmissionControlSettings) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_admission_control_proto_rawDescGZIP(), []int{0}
}

func (x *AdmissionControlSettings) GetClusterConfig() *storage.DynamicClusterConfig {
	if x != nil {
		return x.ClusterConfig
	}
	return nil
}

func (x *AdmissionControlSettings) GetEnforcedDeployTimePolicies() *storage.PolicyList {
	if x != nil {
		return x.EnforcedDeployTimePolicies
	}
	return nil
}

func (x *AdmissionControlSettings) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *AdmissionControlSettings) GetCacheVersion() string {
	if x != nil {
		return x.CacheVersion
	}
	return ""
}

func (x *AdmissionControlSettings) GetCentralEndpoint() string {
	if x != nil {
		return x.CentralEndpoint
	}
	return ""
}

func (x *AdmissionControlSettings) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *AdmissionControlSettings) GetRuntimePolicies() *storage.PolicyList {
	if x != nil {
		return x.RuntimePolicies
	}
	return nil
}

type AdmissionControlAlerts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlertResults []*central.AlertResults `protobuf:"bytes,1,rep,name=alert_results,json=alertResults,proto3" json:"alert_results,omitempty"`
}

func (x *AdmissionControlAlerts) Reset() {
	*x = AdmissionControlAlerts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_sensor_admission_control_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdmissionControlAlerts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdmissionControlAlerts) ProtoMessage() {}

func (x *AdmissionControlAlerts) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_admission_control_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdmissionControlAlerts.ProtoReflect.Descriptor instead.
func (*AdmissionControlAlerts) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_admission_control_proto_rawDescGZIP(), []int{1}
}

func (x *AdmissionControlAlerts) GetAlertResults() []*central.AlertResults {
	if x != nil {
		return x.AlertResults
	}
	return nil
}

type AdmCtrlUpdateResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action central.ResourceAction `protobuf:"varint,1,opt,name=action,proto3,enum=central.ResourceAction" json:"action,omitempty"`
	// Types that are assignable to Resource:
	//
	//	*AdmCtrlUpdateResourceRequest_Deployment
	//	*AdmCtrlUpdateResourceRequest_Pod
	//	*AdmCtrlUpdateResourceRequest_Namespace
	//	*AdmCtrlUpdateResourceRequest_Synced
	Resource isAdmCtrlUpdateResourceRequest_Resource `protobuf_oneof:"resource"`
}

func (x *AdmCtrlUpdateResourceRequest) Reset() {
	*x = AdmCtrlUpdateResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_sensor_admission_control_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdmCtrlUpdateResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdmCtrlUpdateResourceRequest) ProtoMessage() {}

func (x *AdmCtrlUpdateResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_admission_control_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdmCtrlUpdateResourceRequest.ProtoReflect.Descriptor instead.
func (*AdmCtrlUpdateResourceRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_admission_control_proto_rawDescGZIP(), []int{2}
}

func (x *AdmCtrlUpdateResourceRequest) GetAction() central.ResourceAction {
	if x != nil {
		return x.Action
	}
	return central.ResourceAction(0)
}

func (m *AdmCtrlUpdateResourceRequest) GetResource() isAdmCtrlUpdateResourceRequest_Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (x *AdmCtrlUpdateResourceRequest) GetDeployment() *storage.Deployment {
	if x, ok := x.GetResource().(*AdmCtrlUpdateResourceRequest_Deployment); ok {
		return x.Deployment
	}
	return nil
}

func (x *AdmCtrlUpdateResourceRequest) GetPod() *storage.Pod {
	if x, ok := x.GetResource().(*AdmCtrlUpdateResourceRequest_Pod); ok {
		return x.Pod
	}
	return nil
}

func (x *AdmCtrlUpdateResourceRequest) GetNamespace() *storage.NamespaceMetadata {
	if x, ok := x.GetResource().(*AdmCtrlUpdateResourceRequest_Namespace); ok {
		return x.Namespace
	}
	return nil
}

func (x *AdmCtrlUpdateResourceRequest) GetSynced() *AdmCtrlUpdateResourceRequest_ResourcesSynced {
	if x, ok := x.GetResource().(*AdmCtrlUpdateResourceRequest_Synced); ok {
		return x.Synced
	}
	return nil
}

type isAdmCtrlUpdateResourceRequest_Resource interface {
	isAdmCtrlUpdateResourceRequest_Resource()
}

type AdmCtrlUpdateResourceRequest_Deployment struct {
	Deployment *storage.Deployment `protobuf:"bytes,2,opt,name=deployment,proto3,oneof"`
}

type AdmCtrlUpdateResourceRequest_Pod struct {
	Pod *storage.Pod `protobuf:"bytes,3,opt,name=pod,proto3,oneof"`
}

type AdmCtrlUpdateResourceRequest_Namespace struct {
	Namespace *storage.NamespaceMetadata `protobuf:"bytes,4,opt,name=namespace,proto3,oneof"`
}

type AdmCtrlUpdateResourceRequest_Synced struct {
	Synced *AdmCtrlUpdateResourceRequest_ResourcesSynced `protobuf:"bytes,5,opt,name=synced,proto3,oneof"`
}

func (*AdmCtrlUpdateResourceRequest_Deployment) isAdmCtrlUpdateResourceRequest_Resource() {}

func (*AdmCtrlUpdateResourceRequest_Pod) isAdmCtrlUpdateResourceRequest_Resource() {}

func (*AdmCtrlUpdateResourceRequest_Namespace) isAdmCtrlUpdateResourceRequest_Resource() {}

func (*AdmCtrlUpdateResourceRequest_Synced) isAdmCtrlUpdateResourceRequest_Resource() {}

type AdmCtrlUpdateResourceRequest_ResourcesSynced struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdmCtrlUpdateResourceRequest_ResourcesSynced) Reset() {
	*x = AdmCtrlUpdateResourceRequest_ResourcesSynced{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_sensor_admission_control_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdmCtrlUpdateResourceRequest_ResourcesSynced) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdmCtrlUpdateResourceRequest_ResourcesSynced) ProtoMessage() {}

func (x *AdmCtrlUpdateResourceRequest_ResourcesSynced) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_admission_control_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdmCtrlUpdateResourceRequest_ResourcesSynced.ProtoReflect.Descriptor instead.
func (*AdmCtrlUpdateResourceRequest_ResourcesSynced) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_admission_control_proto_rawDescGZIP(), []int{2, 0}
}

var File_internalapi_sensor_admission_control_proto protoreflect.FileDescriptor

var file_internalapi_sensor_admission_control_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x03, 0x0a, 0x18, 0x41, 0x64, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x44, 0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x56, 0x0a, 0x1d, 0x65, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x1a, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x64,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x65, 0x6e,
	0x74, 0x72, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x10, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0f, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x22, 0x54, 0x0a, 0x16, 0x41,
	0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x0d, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63,
	0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x52, 0x0c, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x22, 0xd3, 0x02, 0x0a, 0x1c, 0x41, 0x64, 0x6d, 0x43, 0x74, 0x72, 0x6c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0a,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x70, 0x6f,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x50, 0x6f, 0x64, 0x48, 0x00, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x12, 0x3a, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x06, 0x73, 0x79, 0x6e, 0x63,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x2e, 0x41, 0x64, 0x6d, 0x43, 0x74, 0x72, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x06, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x1a, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_sensor_admission_control_proto_rawDescOnce sync.Once
	file_internalapi_sensor_admission_control_proto_rawDescData = file_internalapi_sensor_admission_control_proto_rawDesc
)

func file_internalapi_sensor_admission_control_proto_rawDescGZIP() []byte {
	file_internalapi_sensor_admission_control_proto_rawDescOnce.Do(func() {
		file_internalapi_sensor_admission_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_sensor_admission_control_proto_rawDescData)
	})
	return file_internalapi_sensor_admission_control_proto_rawDescData
}

var file_internalapi_sensor_admission_control_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internalapi_sensor_admission_control_proto_goTypes = []interface{}{
	(*AdmissionControlSettings)(nil),                     // 0: sensor.AdmissionControlSettings
	(*AdmissionControlAlerts)(nil),                       // 1: sensor.AdmissionControlAlerts
	(*AdmCtrlUpdateResourceRequest)(nil),                 // 2: sensor.AdmCtrlUpdateResourceRequest
	(*AdmCtrlUpdateResourceRequest_ResourcesSynced)(nil), // 3: sensor.AdmCtrlUpdateResourceRequest.ResourcesSynced
	(*storage.DynamicClusterConfig)(nil),                 // 4: storage.DynamicClusterConfig
	(*storage.PolicyList)(nil),                           // 5: storage.PolicyList
	(*timestamppb.Timestamp)(nil),                        // 6: google.protobuf.Timestamp
	(*central.AlertResults)(nil),                         // 7: central.AlertResults
	(central.ResourceAction)(0),                          // 8: central.ResourceAction
	(*storage.Deployment)(nil),                           // 9: storage.Deployment
	(*storage.Pod)(nil),                                  // 10: storage.Pod
	(*storage.NamespaceMetadata)(nil),                    // 11: storage.NamespaceMetadata
}
var file_internalapi_sensor_admission_control_proto_depIdxs = []int32{
	4,  // 0: sensor.AdmissionControlSettings.cluster_config:type_name -> storage.DynamicClusterConfig
	5,  // 1: sensor.AdmissionControlSettings.enforced_deploy_time_policies:type_name -> storage.PolicyList
	6,  // 2: sensor.AdmissionControlSettings.timestamp:type_name -> google.protobuf.Timestamp
	5,  // 3: sensor.AdmissionControlSettings.runtime_policies:type_name -> storage.PolicyList
	7,  // 4: sensor.AdmissionControlAlerts.alert_results:type_name -> central.AlertResults
	8,  // 5: sensor.AdmCtrlUpdateResourceRequest.action:type_name -> central.ResourceAction
	9,  // 6: sensor.AdmCtrlUpdateResourceRequest.deployment:type_name -> storage.Deployment
	10, // 7: sensor.AdmCtrlUpdateResourceRequest.pod:type_name -> storage.Pod
	11, // 8: sensor.AdmCtrlUpdateResourceRequest.namespace:type_name -> storage.NamespaceMetadata
	3,  // 9: sensor.AdmCtrlUpdateResourceRequest.synced:type_name -> sensor.AdmCtrlUpdateResourceRequest.ResourcesSynced
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_internalapi_sensor_admission_control_proto_init() }
func file_internalapi_sensor_admission_control_proto_init() {
	if File_internalapi_sensor_admission_control_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internalapi_sensor_admission_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdmissionControlSettings); i {
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
		file_internalapi_sensor_admission_control_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdmissionControlAlerts); i {
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
		file_internalapi_sensor_admission_control_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdmCtrlUpdateResourceRequest); i {
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
		file_internalapi_sensor_admission_control_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdmCtrlUpdateResourceRequest_ResourcesSynced); i {
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
	file_internalapi_sensor_admission_control_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*AdmCtrlUpdateResourceRequest_Deployment)(nil),
		(*AdmCtrlUpdateResourceRequest_Pod)(nil),
		(*AdmCtrlUpdateResourceRequest_Namespace)(nil),
		(*AdmCtrlUpdateResourceRequest_Synced)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_sensor_admission_control_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalapi_sensor_admission_control_proto_goTypes,
		DependencyIndexes: file_internalapi_sensor_admission_control_proto_depIdxs,
		MessageInfos:      file_internalapi_sensor_admission_control_proto_msgTypes,
	}.Build()
	File_internalapi_sensor_admission_control_proto = out.File
	file_internalapi_sensor_admission_control_proto_rawDesc = nil
	file_internalapi_sensor_admission_control_proto_goTypes = nil
	file_internalapi_sensor_admission_control_proto_depIdxs = nil
}
