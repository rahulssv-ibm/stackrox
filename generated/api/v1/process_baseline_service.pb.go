// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: api/v1/process_baseline_service.proto

package v1

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

type GetProcessBaselineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *storage.ProcessBaselineKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetProcessBaselineRequest) Reset() {
	*x = GetProcessBaselineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_process_baseline_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProcessBaselineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProcessBaselineRequest) ProtoMessage() {}

func (x *GetProcessBaselineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_process_baseline_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProcessBaselineRequest.ProtoReflect.Descriptor instead.
func (*GetProcessBaselineRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_process_baseline_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetProcessBaselineRequest) GetKey() *storage.ProcessBaselineKey {
	if x != nil {
		return x.Key
	}
	return nil
}

type UpdateProcessBaselinesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys           []*storage.ProcessBaselineKey `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	AddElements    []*storage.BaselineItem       `protobuf:"bytes,2,rep,name=add_elements,json=addElements,proto3" json:"add_elements,omitempty"`
	RemoveElements []*storage.BaselineItem       `protobuf:"bytes,3,rep,name=remove_elements,json=removeElements,proto3" json:"remove_elements,omitempty"`
}

func (x *UpdateProcessBaselinesRequest) Reset() {
	*x = UpdateProcessBaselinesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_process_baseline_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProcessBaselinesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProcessBaselinesRequest) ProtoMessage() {}

func (x *UpdateProcessBaselinesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_process_baseline_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProcessBaselinesRequest.ProtoReflect.Descriptor instead.
func (*UpdateProcessBaselinesRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_process_baseline_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateProcessBaselinesRequest) GetKeys() []*storage.ProcessBaselineKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *UpdateProcessBaselinesRequest) GetAddElements() []*storage.BaselineItem {
	if x != nil {
		return x.AddElements
	}
	return nil
}

func (x *UpdateProcessBaselinesRequest) GetRemoveElements() []*storage.BaselineItem {
	if x != nil {
		return x.RemoveElements
	}
	return nil
}

type ProcessBaselinesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Baselines []*storage.ProcessBaseline `protobuf:"bytes,1,rep,name=baselines,proto3" json:"baselines,omitempty"`
}

func (x *ProcessBaselinesResponse) Reset() {
	*x = ProcessBaselinesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_process_baseline_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessBaselinesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessBaselinesResponse) ProtoMessage() {}

func (x *ProcessBaselinesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_process_baseline_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessBaselinesResponse.ProtoReflect.Descriptor instead.
func (*ProcessBaselinesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_process_baseline_service_proto_rawDescGZIP(), []int{2}
}

func (x *ProcessBaselinesResponse) GetBaselines() []*storage.ProcessBaseline {
	if x != nil {
		return x.Baselines
	}
	return nil
}

type ProcessBaselineUpdateError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string                      `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Key   *storage.ProcessBaselineKey `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ProcessBaselineUpdateError) Reset() {
	*x = ProcessBaselineUpdateError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_process_baseline_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessBaselineUpdateError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessBaselineUpdateError) ProtoMessage() {}

func (x *ProcessBaselineUpdateError) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_process_baseline_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessBaselineUpdateError.ProtoReflect.Descriptor instead.
func (*ProcessBaselineUpdateError) Descriptor() ([]byte, []int) {
	return file_api_v1_process_baseline_service_proto_rawDescGZIP(), []int{3}
}

func (x *ProcessBaselineUpdateError) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ProcessBaselineUpdateError) GetKey() *storage.ProcessBaselineKey {
	if x != nil {
		return x.Key
	}
	return nil
}

type UpdateProcessBaselinesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Baselines []*storage.ProcessBaseline    `protobuf:"bytes,1,rep,name=baselines,proto3" json:"baselines,omitempty"`
	Errors    []*ProcessBaselineUpdateError `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *UpdateProcessBaselinesResponse) Reset() {
	*x = UpdateProcessBaselinesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_process_baseline_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProcessBaselinesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProcessBaselinesResponse) ProtoMessage() {}

func (x *UpdateProcessBaselinesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_process_baseline_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProcessBaselinesResponse.ProtoReflect.Descriptor instead.
func (*UpdateProcessBaselinesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_process_baseline_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateProcessBaselinesResponse) GetBaselines() []*storage.ProcessBaseline {
	if x != nil {
		return x.Baselines
	}
	return nil
}

func (x *UpdateProcessBaselinesResponse) GetErrors() []*ProcessBaselineUpdateError {
	if x != nil {
		return x.Errors
	}
	return nil
}

type LockProcessBaselinesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys   []*storage.ProcessBaselineKey `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Locked bool                          `protobuf:"varint,2,opt,name=locked,proto3" json:"locked,omitempty"`
}

func (x *LockProcessBaselinesRequest) Reset() {
	*x = LockProcessBaselinesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_process_baseline_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockProcessBaselinesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockProcessBaselinesRequest) ProtoMessage() {}

func (x *LockProcessBaselinesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_process_baseline_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockProcessBaselinesRequest.ProtoReflect.Descriptor instead.
func (*LockProcessBaselinesRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_process_baseline_service_proto_rawDescGZIP(), []int{5}
}

func (x *LockProcessBaselinesRequest) GetKeys() []*storage.ProcessBaselineKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *LockProcessBaselinesRequest) GetLocked() bool {
	if x != nil {
		return x.Locked
	}
	return false
}

type DeleteProcessBaselinesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query   string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Confirm bool   `protobuf:"varint,2,opt,name=confirm,proto3" json:"confirm,omitempty"`
}

func (x *DeleteProcessBaselinesRequest) Reset() {
	*x = DeleteProcessBaselinesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_process_baseline_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProcessBaselinesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProcessBaselinesRequest) ProtoMessage() {}

func (x *DeleteProcessBaselinesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_process_baseline_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProcessBaselinesRequest.ProtoReflect.Descriptor instead.
func (*DeleteProcessBaselinesRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_process_baseline_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteProcessBaselinesRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *DeleteProcessBaselinesRequest) GetConfirm() bool {
	if x != nil {
		return x.Confirm
	}
	return false
}

type DeleteProcessBaselinesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumDeleted int32 `protobuf:"varint,1,opt,name=num_deleted,json=numDeleted,proto3" json:"num_deleted,omitempty"`
	DryRun     bool  `protobuf:"varint,2,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
}

func (x *DeleteProcessBaselinesResponse) Reset() {
	*x = DeleteProcessBaselinesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_process_baseline_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProcessBaselinesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProcessBaselinesResponse) ProtoMessage() {}

func (x *DeleteProcessBaselinesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_process_baseline_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProcessBaselinesResponse.ProtoReflect.Descriptor instead.
func (*DeleteProcessBaselinesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_process_baseline_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteProcessBaselinesResponse) GetNumDeleted() int32 {
	if x != nil {
		return x.NumDeleted
	}
	return 0
}

func (x *DeleteProcessBaselinesResponse) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

var File_api_v1_process_baseline_service_proto protoreflect.FileDescriptor

var file_api_v1_process_baseline_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4b, 0x65, 0x79,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xca, 0x01, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4b,
	0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x38, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x5f,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x3e, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x52, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x09, 0x62, 0x61, 0x73,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x61, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x90, 0x01, 0x0a, 0x1e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09,
	0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x66, 0x0a, 0x1b,
	0x4c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x6b,
	0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x22, 0x4f, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x22, 0x5a, 0x0a, 0x1e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6e, 0x75,
	0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x72, 0x79, 0x5f,
	0x72, 0x75, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x72, 0x79, 0x52, 0x75,
	0x6e, 0x32, 0x8f, 0x04, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x12, 0x80, 0x01,
	0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42,
	0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61,
	0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x1a, 0x14, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73,
	0x12, 0x81, 0x01, 0x0a, 0x14, 0x4c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x1a, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2f,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x7d, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x21,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72,
	0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a,
	0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x58, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_process_baseline_service_proto_rawDescOnce sync.Once
	file_api_v1_process_baseline_service_proto_rawDescData = file_api_v1_process_baseline_service_proto_rawDesc
)

func file_api_v1_process_baseline_service_proto_rawDescGZIP() []byte {
	file_api_v1_process_baseline_service_proto_rawDescOnce.Do(func() {
		file_api_v1_process_baseline_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_process_baseline_service_proto_rawDescData)
	})
	return file_api_v1_process_baseline_service_proto_rawDescData
}

var file_api_v1_process_baseline_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_v1_process_baseline_service_proto_goTypes = []any{
	(*GetProcessBaselineRequest)(nil),      // 0: v1.GetProcessBaselineRequest
	(*UpdateProcessBaselinesRequest)(nil),  // 1: v1.UpdateProcessBaselinesRequest
	(*ProcessBaselinesResponse)(nil),       // 2: v1.ProcessBaselinesResponse
	(*ProcessBaselineUpdateError)(nil),     // 3: v1.ProcessBaselineUpdateError
	(*UpdateProcessBaselinesResponse)(nil), // 4: v1.UpdateProcessBaselinesResponse
	(*LockProcessBaselinesRequest)(nil),    // 5: v1.LockProcessBaselinesRequest
	(*DeleteProcessBaselinesRequest)(nil),  // 6: v1.DeleteProcessBaselinesRequest
	(*DeleteProcessBaselinesResponse)(nil), // 7: v1.DeleteProcessBaselinesResponse
	(*storage.ProcessBaselineKey)(nil),     // 8: storage.ProcessBaselineKey
	(*storage.BaselineItem)(nil),           // 9: storage.BaselineItem
	(*storage.ProcessBaseline)(nil),        // 10: storage.ProcessBaseline
}
var file_api_v1_process_baseline_service_proto_depIdxs = []int32{
	8,  // 0: v1.GetProcessBaselineRequest.key:type_name -> storage.ProcessBaselineKey
	8,  // 1: v1.UpdateProcessBaselinesRequest.keys:type_name -> storage.ProcessBaselineKey
	9,  // 2: v1.UpdateProcessBaselinesRequest.add_elements:type_name -> storage.BaselineItem
	9,  // 3: v1.UpdateProcessBaselinesRequest.remove_elements:type_name -> storage.BaselineItem
	10, // 4: v1.ProcessBaselinesResponse.baselines:type_name -> storage.ProcessBaseline
	8,  // 5: v1.ProcessBaselineUpdateError.key:type_name -> storage.ProcessBaselineKey
	10, // 6: v1.UpdateProcessBaselinesResponse.baselines:type_name -> storage.ProcessBaseline
	3,  // 7: v1.UpdateProcessBaselinesResponse.errors:type_name -> v1.ProcessBaselineUpdateError
	8,  // 8: v1.LockProcessBaselinesRequest.keys:type_name -> storage.ProcessBaselineKey
	0,  // 9: v1.ProcessBaselineService.GetProcessBaseline:input_type -> v1.GetProcessBaselineRequest
	1,  // 10: v1.ProcessBaselineService.UpdateProcessBaselines:input_type -> v1.UpdateProcessBaselinesRequest
	5,  // 11: v1.ProcessBaselineService.LockProcessBaselines:input_type -> v1.LockProcessBaselinesRequest
	6,  // 12: v1.ProcessBaselineService.DeleteProcessBaselines:input_type -> v1.DeleteProcessBaselinesRequest
	10, // 13: v1.ProcessBaselineService.GetProcessBaseline:output_type -> storage.ProcessBaseline
	4,  // 14: v1.ProcessBaselineService.UpdateProcessBaselines:output_type -> v1.UpdateProcessBaselinesResponse
	4,  // 15: v1.ProcessBaselineService.LockProcessBaselines:output_type -> v1.UpdateProcessBaselinesResponse
	7,  // 16: v1.ProcessBaselineService.DeleteProcessBaselines:output_type -> v1.DeleteProcessBaselinesResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_v1_process_baseline_service_proto_init() }
func file_api_v1_process_baseline_service_proto_init() {
	if File_api_v1_process_baseline_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_process_baseline_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetProcessBaselineRequest); i {
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
		file_api_v1_process_baseline_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateProcessBaselinesRequest); i {
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
		file_api_v1_process_baseline_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ProcessBaselinesResponse); i {
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
		file_api_v1_process_baseline_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ProcessBaselineUpdateError); i {
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
		file_api_v1_process_baseline_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateProcessBaselinesResponse); i {
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
		file_api_v1_process_baseline_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*LockProcessBaselinesRequest); i {
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
		file_api_v1_process_baseline_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteProcessBaselinesRequest); i {
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
		file_api_v1_process_baseline_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteProcessBaselinesResponse); i {
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
			RawDescriptor: file_api_v1_process_baseline_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_process_baseline_service_proto_goTypes,
		DependencyIndexes: file_api_v1_process_baseline_service_proto_depIdxs,
		MessageInfos:      file_api_v1_process_baseline_service_proto_msgTypes,
	}.Build()
	File_api_v1_process_baseline_service_proto = out.File
	file_api_v1_process_baseline_service_proto_rawDesc = nil
	file_api_v1_process_baseline_service_proto_goTypes = nil
	file_api_v1_process_baseline_service_proto_depIdxs = nil
}
