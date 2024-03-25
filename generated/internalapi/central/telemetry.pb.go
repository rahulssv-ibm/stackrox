// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: internalapi/central/telemetry.proto

package central

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

type PullTelemetryDataRequest_TelemetryDataType int32

const (
	PullTelemetryDataRequest_UNKNOWN         PullTelemetryDataRequest_TelemetryDataType = 0
	PullTelemetryDataRequest_KUBERNETES_INFO PullTelemetryDataRequest_TelemetryDataType = 1
	PullTelemetryDataRequest_CLUSTER_INFO    PullTelemetryDataRequest_TelemetryDataType = 2
	PullTelemetryDataRequest_METRICS         PullTelemetryDataRequest_TelemetryDataType = 3
)

// Enum value maps for PullTelemetryDataRequest_TelemetryDataType.
var (
	PullTelemetryDataRequest_TelemetryDataType_name = map[int32]string{
		0: "UNKNOWN",
		1: "KUBERNETES_INFO",
		2: "CLUSTER_INFO",
		3: "METRICS",
	}
	PullTelemetryDataRequest_TelemetryDataType_value = map[string]int32{
		"UNKNOWN":         0,
		"KUBERNETES_INFO": 1,
		"CLUSTER_INFO":    2,
		"METRICS":         3,
	}
)

func (x PullTelemetryDataRequest_TelemetryDataType) Enum() *PullTelemetryDataRequest_TelemetryDataType {
	p := new(PullTelemetryDataRequest_TelemetryDataType)
	*p = x
	return p
}

func (x PullTelemetryDataRequest_TelemetryDataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PullTelemetryDataRequest_TelemetryDataType) Descriptor() protoreflect.EnumDescriptor {
	return file_internalapi_central_telemetry_proto_enumTypes[0].Descriptor()
}

func (PullTelemetryDataRequest_TelemetryDataType) Type() protoreflect.EnumType {
	return &file_internalapi_central_telemetry_proto_enumTypes[0]
}

func (x PullTelemetryDataRequest_TelemetryDataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PullTelemetryDataRequest_TelemetryDataType.Descriptor instead.
func (PullTelemetryDataRequest_TelemetryDataType) EnumDescriptor() ([]byte, []int) {
	return file_internalapi_central_telemetry_proto_rawDescGZIP(), []int{1, 0}
}

type CancelPullTelemetryDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *CancelPullTelemetryDataRequest) Reset() {
	*x = CancelPullTelemetryDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_central_telemetry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelPullTelemetryDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelPullTelemetryDataRequest) ProtoMessage() {}

func (x *CancelPullTelemetryDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_telemetry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelPullTelemetryDataRequest.ProtoReflect.Descriptor instead.
func (*CancelPullTelemetryDataRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_central_telemetry_proto_rawDescGZIP(), []int{0}
}

func (x *CancelPullTelemetryDataRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type PullTelemetryDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string                                     `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	DataType  PullTelemetryDataRequest_TelemetryDataType `protobuf:"varint,2,opt,name=data_type,json=dataType,proto3,enum=central.PullTelemetryDataRequest_TelemetryDataType" json:"data_type,omitempty"`
	TimeoutMs int64                                      `protobuf:"varint,3,opt,name=timeout_ms,json=timeoutMs,proto3" json:"timeout_ms,omitempty"`
	Since     *timestamppb.Timestamp                     `protobuf:"bytes,4,opt,name=since,proto3" json:"since,omitempty"`
}

func (x *PullTelemetryDataRequest) Reset() {
	*x = PullTelemetryDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_central_telemetry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullTelemetryDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullTelemetryDataRequest) ProtoMessage() {}

func (x *PullTelemetryDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_telemetry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullTelemetryDataRequest.ProtoReflect.Descriptor instead.
func (*PullTelemetryDataRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_central_telemetry_proto_rawDescGZIP(), []int{1}
}

func (x *PullTelemetryDataRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *PullTelemetryDataRequest) GetDataType() PullTelemetryDataRequest_TelemetryDataType {
	if x != nil {
		return x.DataType
	}
	return PullTelemetryDataRequest_UNKNOWN
}

func (x *PullTelemetryDataRequest) GetTimeoutMs() int64 {
	if x != nil {
		return x.TimeoutMs
	}
	return 0
}

func (x *PullTelemetryDataRequest) GetSince() *timestamppb.Timestamp {
	if x != nil {
		return x.Since
	}
	return nil
}

type TelemetryResponsePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*TelemetryResponsePayload_EndOfStream_
	//	*TelemetryResponsePayload_KubernetesInfo_
	//	*TelemetryResponsePayload_ClusterInfo_
	//	*TelemetryResponsePayload_MetricsInfo
	Payload isTelemetryResponsePayload_Payload `protobuf_oneof:"payload"`
}

func (x *TelemetryResponsePayload) Reset() {
	*x = TelemetryResponsePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_central_telemetry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryResponsePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryResponsePayload) ProtoMessage() {}

func (x *TelemetryResponsePayload) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_telemetry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryResponsePayload.ProtoReflect.Descriptor instead.
func (*TelemetryResponsePayload) Descriptor() ([]byte, []int) {
	return file_internalapi_central_telemetry_proto_rawDescGZIP(), []int{2}
}

func (m *TelemetryResponsePayload) GetPayload() isTelemetryResponsePayload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *TelemetryResponsePayload) GetEndOfStream() *TelemetryResponsePayload_EndOfStream {
	if x, ok := x.GetPayload().(*TelemetryResponsePayload_EndOfStream_); ok {
		return x.EndOfStream
	}
	return nil
}

func (x *TelemetryResponsePayload) GetKubernetesInfo() *TelemetryResponsePayload_KubernetesInfo {
	if x, ok := x.GetPayload().(*TelemetryResponsePayload_KubernetesInfo_); ok {
		return x.KubernetesInfo
	}
	return nil
}

func (x *TelemetryResponsePayload) GetClusterInfo() *TelemetryResponsePayload_ClusterInfo {
	if x, ok := x.GetPayload().(*TelemetryResponsePayload_ClusterInfo_); ok {
		return x.ClusterInfo
	}
	return nil
}

func (x *TelemetryResponsePayload) GetMetricsInfo() *TelemetryResponsePayload_KubernetesInfo {
	if x, ok := x.GetPayload().(*TelemetryResponsePayload_MetricsInfo); ok {
		return x.MetricsInfo
	}
	return nil
}

type isTelemetryResponsePayload_Payload interface {
	isTelemetryResponsePayload_Payload()
}

type TelemetryResponsePayload_EndOfStream_ struct {
	EndOfStream *TelemetryResponsePayload_EndOfStream `protobuf:"bytes,1,opt,name=end_of_stream,json=endOfStream,proto3,oneof"`
}

type TelemetryResponsePayload_KubernetesInfo_ struct {
	KubernetesInfo *TelemetryResponsePayload_KubernetesInfo `protobuf:"bytes,2,opt,name=kubernetes_info,json=kubernetesInfo,proto3,oneof"`
}

type TelemetryResponsePayload_ClusterInfo_ struct {
	ClusterInfo *TelemetryResponsePayload_ClusterInfo `protobuf:"bytes,3,opt,name=cluster_info,json=clusterInfo,proto3,oneof"`
}

type TelemetryResponsePayload_MetricsInfo struct {
	MetricsInfo *TelemetryResponsePayload_KubernetesInfo `protobuf:"bytes,4,opt,name=metrics_info,json=metricsInfo,proto3,oneof"`
}

func (*TelemetryResponsePayload_EndOfStream_) isTelemetryResponsePayload_Payload() {}

func (*TelemetryResponsePayload_KubernetesInfo_) isTelemetryResponsePayload_Payload() {}

func (*TelemetryResponsePayload_ClusterInfo_) isTelemetryResponsePayload_Payload() {}

func (*TelemetryResponsePayload_MetricsInfo) isTelemetryResponsePayload_Payload() {}

type PullTelemetryDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string                    `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Payload   *TelemetryResponsePayload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *PullTelemetryDataResponse) Reset() {
	*x = PullTelemetryDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_central_telemetry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullTelemetryDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullTelemetryDataResponse) ProtoMessage() {}

func (x *PullTelemetryDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_telemetry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullTelemetryDataResponse.ProtoReflect.Descriptor instead.
func (*PullTelemetryDataResponse) Descriptor() ([]byte, []int) {
	return file_internalapi_central_telemetry_proto_rawDescGZIP(), []int{3}
}

func (x *PullTelemetryDataResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *PullTelemetryDataResponse) GetPayload() *TelemetryResponsePayload {
	if x != nil {
		return x.Payload
	}
	return nil
}

type TelemetryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// API user ID hash:
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// An empty endpoint means using default endpoint:
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Telemetry storage service key:
	StorageKeyV1 string `protobuf:"bytes,3,opt,name=storage_key_v1,json=storageKeyV1,proto3" json:"storage_key_v1,omitempty"`
}

func (x *TelemetryConfig) Reset() {
	*x = TelemetryConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_central_telemetry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryConfig) ProtoMessage() {}

func (x *TelemetryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_telemetry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryConfig.ProtoReflect.Descriptor instead.
func (*TelemetryConfig) Descriptor() ([]byte, []int) {
	return file_internalapi_central_telemetry_proto_rawDescGZIP(), []int{4}
}

func (x *TelemetryConfig) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TelemetryConfig) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *TelemetryConfig) GetStorageKeyV1() string {
	if x != nil {
		return x.StorageKeyV1
	}
	return ""
}

type TelemetryResponsePayload_EndOfStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorMessage string `protobuf:"bytes,1,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"` // empty indicates success
}

func (x *TelemetryResponsePayload_EndOfStream) Reset() {
	*x = TelemetryResponsePayload_EndOfStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_central_telemetry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryResponsePayload_EndOfStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryResponsePayload_EndOfStream) ProtoMessage() {}

func (x *TelemetryResponsePayload_EndOfStream) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_telemetry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryResponsePayload_EndOfStream.ProtoReflect.Descriptor instead.
func (*TelemetryResponsePayload_EndOfStream) Descriptor() ([]byte, []int) {
	return file_internalapi_central_telemetry_proto_rawDescGZIP(), []int{2, 0}
}

func (x *TelemetryResponsePayload_EndOfStream) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type TelemetryResponsePayload_KubernetesInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*TelemetryResponsePayload_KubernetesInfo_File `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *TelemetryResponsePayload_KubernetesInfo) Reset() {
	*x = TelemetryResponsePayload_KubernetesInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_central_telemetry_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryResponsePayload_KubernetesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryResponsePayload_KubernetesInfo) ProtoMessage() {}

func (x *TelemetryResponsePayload_KubernetesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_telemetry_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryResponsePayload_KubernetesInfo.ProtoReflect.Descriptor instead.
func (*TelemetryResponsePayload_KubernetesInfo) Descriptor() ([]byte, []int) {
	return file_internalapi_central_telemetry_proto_rawDescGZIP(), []int{2, 1}
}

func (x *TelemetryResponsePayload_KubernetesInfo) GetFiles() []*TelemetryResponsePayload_KubernetesInfo_File {
	if x != nil {
		return x.Files
	}
	return nil
}

type TelemetryResponsePayload_ClusterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *TelemetryResponsePayload_ClusterInfo) Reset() {
	*x = TelemetryResponsePayload_ClusterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_central_telemetry_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryResponsePayload_ClusterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryResponsePayload_ClusterInfo) ProtoMessage() {}

func (x *TelemetryResponsePayload_ClusterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_telemetry_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryResponsePayload_ClusterInfo.ProtoReflect.Descriptor instead.
func (*TelemetryResponsePayload_ClusterInfo) Descriptor() ([]byte, []int) {
	return file_internalapi_central_telemetry_proto_rawDescGZIP(), []int{2, 2}
}

func (x *TelemetryResponsePayload_ClusterInfo) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type TelemetryResponsePayload_KubernetesInfo_File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Contents []byte `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"` // compression (if any) is handled at the gRPC level
}

func (x *TelemetryResponsePayload_KubernetesInfo_File) Reset() {
	*x = TelemetryResponsePayload_KubernetesInfo_File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_central_telemetry_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryResponsePayload_KubernetesInfo_File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryResponsePayload_KubernetesInfo_File) ProtoMessage() {}

func (x *TelemetryResponsePayload_KubernetesInfo_File) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_telemetry_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryResponsePayload_KubernetesInfo_File.ProtoReflect.Descriptor instead.
func (*TelemetryResponsePayload_KubernetesInfo_File) Descriptor() ([]byte, []int) {
	return file_internalapi_central_telemetry_proto_rawDescGZIP(), []int{2, 1, 0}
}

func (x *TelemetryResponsePayload_KubernetesInfo_File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *TelemetryResponsePayload_KubernetesInfo_File) GetContents() []byte {
	if x != nil {
		return x.Contents
	}
	return nil
}

var File_internalapi_central_telemetry_proto protoreflect.FileDescriptor

var file_internalapi_central_telemetry_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3f, 0x0a, 0x1e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x75, 0x6c, 0x6c, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x22, 0xb2, 0x02, 0x0a, 0x18, 0x50, 0x75, 0x6c, 0x6c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x50, 0x0a, 0x09,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x33, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x54, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x73, 0x12, 0x30, 0x0a,
	0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x22,
	0x54, 0x0a, 0x11, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53, 0x5f,
	0x49, 0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45,
	0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x54, 0x52,
	0x49, 0x43, 0x53, 0x10, 0x03, 0x22, 0xf3, 0x04, 0x0a, 0x18, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x53, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x45, 0x6e, 0x64,
	0x4f, 0x66, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x00, 0x52, 0x0b, 0x65, 0x6e, 0x64, 0x4f,
	0x66, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x5b, 0x0a, 0x0f, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x48, 0x00, 0x52, 0x0e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x52, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x65, 0x6e,
	0x74, 0x72, 0x61, 0x6c, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x55, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x48, 0x00, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x32, 0x0a, 0x0b, 0x45, 0x6e, 0x64, 0x4f, 0x66, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x95, 0x01, 0x0a, 0x0e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4b, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e,
	0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x1a, 0x36, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x23, 0x0a, 0x0b, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x77, 0x0a, 0x19, 0x50,
	0x75, 0x6c, 0x6c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6c, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x6c, 0x0a, 0x0f, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x31, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4b, 0x65, 0x79,
	0x56, 0x31, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_central_telemetry_proto_rawDescOnce sync.Once
	file_internalapi_central_telemetry_proto_rawDescData = file_internalapi_central_telemetry_proto_rawDesc
)

func file_internalapi_central_telemetry_proto_rawDescGZIP() []byte {
	file_internalapi_central_telemetry_proto_rawDescOnce.Do(func() {
		file_internalapi_central_telemetry_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_central_telemetry_proto_rawDescData)
	})
	return file_internalapi_central_telemetry_proto_rawDescData
}

var file_internalapi_central_telemetry_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internalapi_central_telemetry_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_internalapi_central_telemetry_proto_goTypes = []interface{}{
	(PullTelemetryDataRequest_TelemetryDataType)(0),      // 0: central.PullTelemetryDataRequest.TelemetryDataType
	(*CancelPullTelemetryDataRequest)(nil),               // 1: central.CancelPullTelemetryDataRequest
	(*PullTelemetryDataRequest)(nil),                     // 2: central.PullTelemetryDataRequest
	(*TelemetryResponsePayload)(nil),                     // 3: central.TelemetryResponsePayload
	(*PullTelemetryDataResponse)(nil),                    // 4: central.PullTelemetryDataResponse
	(*TelemetryConfig)(nil),                              // 5: central.TelemetryConfig
	(*TelemetryResponsePayload_EndOfStream)(nil),         // 6: central.TelemetryResponsePayload.EndOfStream
	(*TelemetryResponsePayload_KubernetesInfo)(nil),      // 7: central.TelemetryResponsePayload.KubernetesInfo
	(*TelemetryResponsePayload_ClusterInfo)(nil),         // 8: central.TelemetryResponsePayload.ClusterInfo
	(*TelemetryResponsePayload_KubernetesInfo_File)(nil), // 9: central.TelemetryResponsePayload.KubernetesInfo.File
	(*timestamppb.Timestamp)(nil),                        // 10: google.protobuf.Timestamp
}
var file_internalapi_central_telemetry_proto_depIdxs = []int32{
	0,  // 0: central.PullTelemetryDataRequest.data_type:type_name -> central.PullTelemetryDataRequest.TelemetryDataType
	10, // 1: central.PullTelemetryDataRequest.since:type_name -> google.protobuf.Timestamp
	6,  // 2: central.TelemetryResponsePayload.end_of_stream:type_name -> central.TelemetryResponsePayload.EndOfStream
	7,  // 3: central.TelemetryResponsePayload.kubernetes_info:type_name -> central.TelemetryResponsePayload.KubernetesInfo
	8,  // 4: central.TelemetryResponsePayload.cluster_info:type_name -> central.TelemetryResponsePayload.ClusterInfo
	7,  // 5: central.TelemetryResponsePayload.metrics_info:type_name -> central.TelemetryResponsePayload.KubernetesInfo
	3,  // 6: central.PullTelemetryDataResponse.payload:type_name -> central.TelemetryResponsePayload
	9,  // 7: central.TelemetryResponsePayload.KubernetesInfo.files:type_name -> central.TelemetryResponsePayload.KubernetesInfo.File
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_internalapi_central_telemetry_proto_init() }
func file_internalapi_central_telemetry_proto_init() {
	if File_internalapi_central_telemetry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internalapi_central_telemetry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelPullTelemetryDataRequest); i {
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
		file_internalapi_central_telemetry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullTelemetryDataRequest); i {
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
		file_internalapi_central_telemetry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryResponsePayload); i {
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
		file_internalapi_central_telemetry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullTelemetryDataResponse); i {
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
		file_internalapi_central_telemetry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryConfig); i {
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
		file_internalapi_central_telemetry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryResponsePayload_EndOfStream); i {
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
		file_internalapi_central_telemetry_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryResponsePayload_KubernetesInfo); i {
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
		file_internalapi_central_telemetry_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryResponsePayload_ClusterInfo); i {
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
		file_internalapi_central_telemetry_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryResponsePayload_KubernetesInfo_File); i {
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
	file_internalapi_central_telemetry_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*TelemetryResponsePayload_EndOfStream_)(nil),
		(*TelemetryResponsePayload_KubernetesInfo_)(nil),
		(*TelemetryResponsePayload_ClusterInfo_)(nil),
		(*TelemetryResponsePayload_MetricsInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_central_telemetry_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalapi_central_telemetry_proto_goTypes,
		DependencyIndexes: file_internalapi_central_telemetry_proto_depIdxs,
		EnumInfos:         file_internalapi_central_telemetry_proto_enumTypes,
		MessageInfos:      file_internalapi_central_telemetry_proto_msgTypes,
	}.Build()
	File_internalapi_central_telemetry_proto = out.File
	file_internalapi_central_telemetry_proto_rawDesc = nil
	file_internalapi_central_telemetry_proto_goTypes = nil
	file_internalapi_central_telemetry_proto_depIdxs = nil
}
