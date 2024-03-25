// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: storage/process_listening_on_port.proto

package storage

import (
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

// The API returns an array of these
type ProcessListeningOnPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint           *ProcessListeningOnPort_Endpoint `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	DeploymentId       string                           `protobuf:"bytes,2,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	ContainerName      string                           `protobuf:"bytes,3,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	PodId              string                           `protobuf:"bytes,4,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
	PodUid             string                           `protobuf:"bytes,5,opt,name=pod_uid,json=podUid,proto3" json:"pod_uid,omitempty"`
	Signal             *ProcessSignal                   `protobuf:"bytes,6,opt,name=signal,proto3" json:"signal,omitempty"`
	ClusterId          string                           `protobuf:"bytes,7,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	Namespace          string                           `protobuf:"bytes,8,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ContainerStartTime *timestamppb.Timestamp           `protobuf:"bytes,9,opt,name=container_start_time,json=containerStartTime,proto3" json:"container_start_time,omitempty"`
	ImageId            string                           `protobuf:"bytes,10,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
}

func (x *ProcessListeningOnPort) Reset() {
	*x = ProcessListeningOnPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_process_listening_on_port_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessListeningOnPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessListeningOnPort) ProtoMessage() {}

func (x *ProcessListeningOnPort) ProtoReflect() protoreflect.Message {
	mi := &file_storage_process_listening_on_port_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessListeningOnPort.ProtoReflect.Descriptor instead.
func (*ProcessListeningOnPort) Descriptor() ([]byte, []int) {
	return file_storage_process_listening_on_port_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessListeningOnPort) GetEndpoint() *ProcessListeningOnPort_Endpoint {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *ProcessListeningOnPort) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *ProcessListeningOnPort) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *ProcessListeningOnPort) GetPodId() string {
	if x != nil {
		return x.PodId
	}
	return ""
}

func (x *ProcessListeningOnPort) GetPodUid() string {
	if x != nil {
		return x.PodUid
	}
	return ""
}

func (x *ProcessListeningOnPort) GetSignal() *ProcessSignal {
	if x != nil {
		return x.Signal
	}
	return nil
}

func (x *ProcessListeningOnPort) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ProcessListeningOnPort) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ProcessListeningOnPort) GetContainerStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ContainerStartTime
	}
	return nil
}

func (x *ProcessListeningOnPort) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

// This is what sensor sends to central
type ProcessListeningOnPortFromSensor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port           uint32                     `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Protocol       L4Protocol                 `protobuf:"varint,2,opt,name=protocol,proto3,enum=storage.L4Protocol" json:"protocol,omitempty"`
	Process        *ProcessIndicatorUniqueKey `protobuf:"bytes,3,opt,name=process,proto3" json:"process,omitempty"`
	CloseTimestamp *timestamppb.Timestamp     `protobuf:"bytes,4,opt,name=close_timestamp,json=closeTimestamp,proto3" json:"close_timestamp,omitempty"`
	ClusterId      string                     `protobuf:"bytes,6,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	DeploymentId   string                     `protobuf:"bytes,7,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	PodUid         string                     `protobuf:"bytes,8,opt,name=pod_uid,json=podUid,proto3" json:"pod_uid,omitempty"`
	Namespace      string                     `protobuf:"bytes,9,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ProcessListeningOnPortFromSensor) Reset() {
	*x = ProcessListeningOnPortFromSensor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_process_listening_on_port_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessListeningOnPortFromSensor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessListeningOnPortFromSensor) ProtoMessage() {}

func (x *ProcessListeningOnPortFromSensor) ProtoReflect() protoreflect.Message {
	mi := &file_storage_process_listening_on_port_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessListeningOnPortFromSensor.ProtoReflect.Descriptor instead.
func (*ProcessListeningOnPortFromSensor) Descriptor() ([]byte, []int) {
	return file_storage_process_listening_on_port_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessListeningOnPortFromSensor) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ProcessListeningOnPortFromSensor) GetProtocol() L4Protocol {
	if x != nil {
		return x.Protocol
	}
	return L4Protocol_L4_PROTOCOL_UNKNOWN
}

func (x *ProcessListeningOnPortFromSensor) GetProcess() *ProcessIndicatorUniqueKey {
	if x != nil {
		return x.Process
	}
	return nil
}

func (x *ProcessListeningOnPortFromSensor) GetCloseTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.CloseTimestamp
	}
	return nil
}

func (x *ProcessListeningOnPortFromSensor) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ProcessListeningOnPortFromSensor) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *ProcessListeningOnPortFromSensor) GetPodUid() string {
	if x != nil {
		return x.PodUid
	}
	return ""
}

func (x *ProcessListeningOnPortFromSensor) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// This is what is stored in the database
type ProcessListeningOnPortStorage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ideally it has to be GENERATED ALWAYS AS IDENTITY, which will make it a
	// bigint with a sequence. Unfortunately at the moment some bits of store
	// generator assume an id has to be a string.
	Id                 string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,type(uuid)"`                                                             // @gotags: sql:"pk,type(uuid)"
	Port               uint32                 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty" search:"Port,store"`                                                        // @gotags: search:"Port,store"
	Protocol           L4Protocol             `protobuf:"varint,3,opt,name=protocol,proto3,enum=storage.L4Protocol" json:"protocol,omitempty" search:"Port Protocol,store"`                        // @gotags: search:"Port Protocol,store"
	CloseTimestamp     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=close_timestamp,json=closeTimestamp,proto3" json:"close_timestamp,omitempty" search:"Closed Time,hidden"`               // @gotags: search:"Closed Time,hidden"
	ProcessIndicatorId string                 `protobuf:"bytes,5,opt,name=process_indicator_id,json=processIndicatorId,proto3" json:"process_indicator_id,omitempty" search:"Process ID,store" sql:"fk(ProcessIndicator:id),no-fk-constraint,index=btree,type(uuid)"` // @gotags: search:"Process ID,store" sql:"fk(ProcessIndicator:id),no-fk-constraint,index=btree,type(uuid)"
	// XXX: Make it a partial index on only active, not closed, PLOP
	Closed bool `protobuf:"varint,6,opt,name=closed,proto3" json:"closed,omitempty" search:"Closed,store" sql:"index=btree"` // @gotags: search:"Closed,store" sql:"index=btree"
	// ProcessIndicator will be not empty only for those cases when we were not
	// able to find references process in the database
	Process      *ProcessIndicatorUniqueKey `protobuf:"bytes,7,opt,name=process,proto3" json:"process,omitempty"`
	DeploymentId string                     `protobuf:"bytes,8,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty" search:"Deployment ID,store" sql:"fk(Deployment:id),no-fk-constraint,index=btree,type(uuid)"` // @gotags: search:"Deployment ID,store" sql:"fk(Deployment:id),no-fk-constraint,index=btree,type(uuid)"
	PodUid       string                     `protobuf:"bytes,9,opt,name=pod_uid,json=podUid,proto3" json:"pod_uid,omitempty" search:"Pod UID,hidden" sql:"fk(Pod:id),no-fk-constraint,index=hash,type(uuid)"`                   // @gotags: search:"Pod UID,hidden" sql:"fk(Pod:id),no-fk-constraint,index=hash,type(uuid)"
	ClusterId    string                     `protobuf:"bytes,10,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" search:"Cluster ID" sql:"type(uuid)"`         // @gotags: search:"Cluster ID" sql:"type(uuid)"
	Namespace    string                     `protobuf:"bytes,11,opt,name=namespace,proto3" json:"namespace,omitempty" search:"Namespace"`                          // @gotags: search:"Namespace"
}

func (x *ProcessListeningOnPortStorage) Reset() {
	*x = ProcessListeningOnPortStorage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_process_listening_on_port_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessListeningOnPortStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessListeningOnPortStorage) ProtoMessage() {}

func (x *ProcessListeningOnPortStorage) ProtoReflect() protoreflect.Message {
	mi := &file_storage_process_listening_on_port_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessListeningOnPortStorage.ProtoReflect.Descriptor instead.
func (*ProcessListeningOnPortStorage) Descriptor() ([]byte, []int) {
	return file_storage_process_listening_on_port_proto_rawDescGZIP(), []int{2}
}

func (x *ProcessListeningOnPortStorage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProcessListeningOnPortStorage) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ProcessListeningOnPortStorage) GetProtocol() L4Protocol {
	if x != nil {
		return x.Protocol
	}
	return L4Protocol_L4_PROTOCOL_UNKNOWN
}

func (x *ProcessListeningOnPortStorage) GetCloseTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.CloseTimestamp
	}
	return nil
}

func (x *ProcessListeningOnPortStorage) GetProcessIndicatorId() string {
	if x != nil {
		return x.ProcessIndicatorId
	}
	return ""
}

func (x *ProcessListeningOnPortStorage) GetClosed() bool {
	if x != nil {
		return x.Closed
	}
	return false
}

func (x *ProcessListeningOnPortStorage) GetProcess() *ProcessIndicatorUniqueKey {
	if x != nil {
		return x.Process
	}
	return nil
}

func (x *ProcessListeningOnPortStorage) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *ProcessListeningOnPortStorage) GetPodUid() string {
	if x != nil {
		return x.PodUid
	}
	return ""
}

func (x *ProcessListeningOnPortStorage) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ProcessListeningOnPortStorage) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type ProcessListeningOnPort_Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port     uint32     `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Protocol L4Protocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=storage.L4Protocol" json:"protocol,omitempty"`
}

func (x *ProcessListeningOnPort_Endpoint) Reset() {
	*x = ProcessListeningOnPort_Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_process_listening_on_port_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessListeningOnPort_Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessListeningOnPort_Endpoint) ProtoMessage() {}

func (x *ProcessListeningOnPort_Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_storage_process_listening_on_port_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessListeningOnPort_Endpoint.ProtoReflect.Descriptor instead.
func (*ProcessListeningOnPort_Endpoint) Descriptor() ([]byte, []int) {
	return file_storage_process_listening_on_port_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProcessListeningOnPort_Endpoint) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ProcessListeningOnPort_Endpoint) GetProtocol() L4Protocol {
	if x != nil {
		return x.Protocol
	}
	return L4Protocol_L4_PROTOCOL_UNKNOWN
}

var File_storage_process_listening_on_port_proto protoreflect.FileDescriptor

var file_storage_process_listening_on_port_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x6e, 0x5f, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x81, 0x04, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x44, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x2e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a,
	0x06, 0x70, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x6f, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x64, 0x5f, 0x75, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x64, 0x55, 0x69, 0x64, 0x12, 0x2e, 0x0a,
	0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x14, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x1a, 0x4f, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x4c, 0x34, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x22, 0xe5, 0x02, 0x0a, 0x20, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x46,
	0x72, 0x6f, 0x6d, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2f, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x34, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x3c,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x4b, 0x65, 0x79, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x43, 0x0a, 0x0f,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0e, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x64, 0x5f, 0x75, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x64, 0x55, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xbc, 0x03, 0x0a,
	0x1d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e,
	0x67, 0x4f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4c,
	0x34, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x12, 0x43, 0x0a, 0x0f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x64, 0x12, 0x3c, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x64, 0x5f, 0x75, 0x69, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x64, 0x55, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x0f, 0x5a, 0x0a, 0x2e,
	0x2f, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_process_listening_on_port_proto_rawDescOnce sync.Once
	file_storage_process_listening_on_port_proto_rawDescData = file_storage_process_listening_on_port_proto_rawDesc
)

func file_storage_process_listening_on_port_proto_rawDescGZIP() []byte {
	file_storage_process_listening_on_port_proto_rawDescOnce.Do(func() {
		file_storage_process_listening_on_port_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_process_listening_on_port_proto_rawDescData)
	})
	return file_storage_process_listening_on_port_proto_rawDescData
}

var file_storage_process_listening_on_port_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_storage_process_listening_on_port_proto_goTypes = []interface{}{
	(*ProcessListeningOnPort)(nil),           // 0: storage.ProcessListeningOnPort
	(*ProcessListeningOnPortFromSensor)(nil), // 1: storage.ProcessListeningOnPortFromSensor
	(*ProcessListeningOnPortStorage)(nil),    // 2: storage.ProcessListeningOnPortStorage
	(*ProcessListeningOnPort_Endpoint)(nil),  // 3: storage.ProcessListeningOnPort.Endpoint
	(*ProcessSignal)(nil),                    // 4: storage.ProcessSignal
	(*timestamppb.Timestamp)(nil),            // 5: google.protobuf.Timestamp
	(L4Protocol)(0),                          // 6: storage.L4Protocol
	(*ProcessIndicatorUniqueKey)(nil),        // 7: storage.ProcessIndicatorUniqueKey
}
var file_storage_process_listening_on_port_proto_depIdxs = []int32{
	3,  // 0: storage.ProcessListeningOnPort.endpoint:type_name -> storage.ProcessListeningOnPort.Endpoint
	4,  // 1: storage.ProcessListeningOnPort.signal:type_name -> storage.ProcessSignal
	5,  // 2: storage.ProcessListeningOnPort.container_start_time:type_name -> google.protobuf.Timestamp
	6,  // 3: storage.ProcessListeningOnPortFromSensor.protocol:type_name -> storage.L4Protocol
	7,  // 4: storage.ProcessListeningOnPortFromSensor.process:type_name -> storage.ProcessIndicatorUniqueKey
	5,  // 5: storage.ProcessListeningOnPortFromSensor.close_timestamp:type_name -> google.protobuf.Timestamp
	6,  // 6: storage.ProcessListeningOnPortStorage.protocol:type_name -> storage.L4Protocol
	5,  // 7: storage.ProcessListeningOnPortStorage.close_timestamp:type_name -> google.protobuf.Timestamp
	7,  // 8: storage.ProcessListeningOnPortStorage.process:type_name -> storage.ProcessIndicatorUniqueKey
	6,  // 9: storage.ProcessListeningOnPort.Endpoint.protocol:type_name -> storage.L4Protocol
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_storage_process_listening_on_port_proto_init() }
func file_storage_process_listening_on_port_proto_init() {
	if File_storage_process_listening_on_port_proto != nil {
		return
	}
	file_storage_network_flow_proto_init()
	file_storage_process_indicator_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_storage_process_listening_on_port_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessListeningOnPort); i {
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
		file_storage_process_listening_on_port_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessListeningOnPortFromSensor); i {
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
		file_storage_process_listening_on_port_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessListeningOnPortStorage); i {
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
		file_storage_process_listening_on_port_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessListeningOnPort_Endpoint); i {
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
			RawDescriptor: file_storage_process_listening_on_port_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_process_listening_on_port_proto_goTypes,
		DependencyIndexes: file_storage_process_listening_on_port_proto_depIdxs,
		MessageInfos:      file_storage_process_listening_on_port_proto_msgTypes,
	}.Build()
	File_storage_process_listening_on_port_proto = out.File
	file_storage_process_listening_on_port_proto_rawDesc = nil
	file_storage_process_listening_on_port_proto_goTypes = nil
	file_storage_process_listening_on_port_proto_depIdxs = nil
}
