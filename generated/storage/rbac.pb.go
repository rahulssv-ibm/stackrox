// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: storage/rbac.proto

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

type SubjectKind int32

const (
	SubjectKind_UNSET_KIND      SubjectKind = 0
	SubjectKind_SERVICE_ACCOUNT SubjectKind = 1
	SubjectKind_USER            SubjectKind = 2
	SubjectKind_GROUP           SubjectKind = 3
)

// Enum value maps for SubjectKind.
var (
	SubjectKind_name = map[int32]string{
		0: "UNSET_KIND",
		1: "SERVICE_ACCOUNT",
		2: "USER",
		3: "GROUP",
	}
	SubjectKind_value = map[string]int32{
		"UNSET_KIND":      0,
		"SERVICE_ACCOUNT": 1,
		"USER":            2,
		"GROUP":           3,
	}
)

func (x SubjectKind) Enum() *SubjectKind {
	p := new(SubjectKind)
	*p = x
	return p
}

func (x SubjectKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubjectKind) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_rbac_proto_enumTypes[0].Descriptor()
}

func (SubjectKind) Type() protoreflect.EnumType {
	return &file_storage_rbac_proto_enumTypes[0]
}

func (x SubjectKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubjectKind.Descriptor instead.
func (SubjectKind) EnumDescriptor() ([]byte, []int) {
	return file_storage_rbac_proto_rawDescGZIP(), []int{0}
}

// For any update to PermissionLevel, also update:
// - pkg/searchbasedpolicies/builders/k8s_rbac.go
// - ui/src/messages/common.js
type PermissionLevel int32

const (
	PermissionLevel_UNSET                 PermissionLevel = 0
	PermissionLevel_NONE                  PermissionLevel = 1
	PermissionLevel_DEFAULT               PermissionLevel = 2
	PermissionLevel_ELEVATED_IN_NAMESPACE PermissionLevel = 3
	PermissionLevel_ELEVATED_CLUSTER_WIDE PermissionLevel = 4
	PermissionLevel_CLUSTER_ADMIN         PermissionLevel = 5
)

// Enum value maps for PermissionLevel.
var (
	PermissionLevel_name = map[int32]string{
		0: "UNSET",
		1: "NONE",
		2: "DEFAULT",
		3: "ELEVATED_IN_NAMESPACE",
		4: "ELEVATED_CLUSTER_WIDE",
		5: "CLUSTER_ADMIN",
	}
	PermissionLevel_value = map[string]int32{
		"UNSET":                 0,
		"NONE":                  1,
		"DEFAULT":               2,
		"ELEVATED_IN_NAMESPACE": 3,
		"ELEVATED_CLUSTER_WIDE": 4,
		"CLUSTER_ADMIN":         5,
	}
)

func (x PermissionLevel) Enum() *PermissionLevel {
	p := new(PermissionLevel)
	*p = x
	return p
}

func (x PermissionLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PermissionLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_rbac_proto_enumTypes[1].Descriptor()
}

func (PermissionLevel) Type() protoreflect.EnumType {
	return &file_storage_rbac_proto_enumTypes[1]
}

func (x PermissionLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PermissionLevel.Descriptor instead.
func (PermissionLevel) EnumDescriptor() ([]byte, []int) {
	return file_storage_rbac_proto_rawDescGZIP(), []int{1}
}

// Properties of an individual k8s Role or ClusterRole.
// ////////////////////////////////////////
type K8SRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" search:"Role ID,hidden" sql:"pk,type(uuid)"`                                                                                                           // @gotags: search:"Role ID,hidden" sql:"pk,type(uuid)"
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" search:"Role"`                                                                                                       // @gotags: search:"Role"
	Namespace   string                 `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty" search:"Namespace,store"`                                                                                             // @gotags: search:"Namespace,store"
	ClusterId   string                 `protobuf:"bytes,4,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" search:"Cluster ID,store,hidden" sql:"type(uuid)"`                                                                            // @gotags: search:"Cluster ID,store,hidden" sql:"type(uuid)"
	ClusterName string                 `protobuf:"bytes,5,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty" search:"Cluster"`                                                                      // @gotags: search:"Cluster"
	ClusterRole bool                   `protobuf:"varint,6,opt,name=cluster_role,json=clusterRole,proto3" json:"cluster_role,omitempty" search:"Cluster Role"`                                                                     // @gotags: search:"Cluster Role"
	Labels      map[string]string      `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" search:"Role Label"`           // @gotags: search:"Role Label"
	Annotations map[string]string      `protobuf:"bytes,8,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" search:"Role Annotation"` // @gotags: search:"Role Annotation"
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Rules       []*PolicyRule          `protobuf:"bytes,10,rep,name=rules,proto3" json:"rules,omitempty" sensorhash:"set"` // @gotags: sensorhash:"set"
}

func (x *K8SRole) Reset() {
	*x = K8SRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_rbac_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8SRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SRole) ProtoMessage() {}

func (x *K8SRole) ProtoReflect() protoreflect.Message {
	mi := &file_storage_rbac_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SRole.ProtoReflect.Descriptor instead.
func (*K8SRole) Descriptor() ([]byte, []int) {
	return file_storage_rbac_proto_rawDescGZIP(), []int{0}
}

func (x *K8SRole) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *K8SRole) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *K8SRole) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *K8SRole) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *K8SRole) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *K8SRole) GetClusterRole() bool {
	if x != nil {
		return x.ClusterRole
	}
	return false
}

func (x *K8SRole) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *K8SRole) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *K8SRole) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *K8SRole) GetRules() []*PolicyRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

// Properties of an individual rules that grant permissions to resources.
// ////////////////////////////////////////
type PolicyRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Verbs           []string `protobuf:"bytes,1,rep,name=verbs,proto3" json:"verbs,omitempty"`
	ApiGroups       []string `protobuf:"bytes,2,rep,name=api_groups,json=apiGroups,proto3" json:"api_groups,omitempty"`
	Resources       []string `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources,omitempty"`
	NonResourceUrls []string `protobuf:"bytes,4,rep,name=non_resource_urls,json=nonResourceUrls,proto3" json:"non_resource_urls,omitempty"`
	ResourceNames   []string `protobuf:"bytes,5,rep,name=resource_names,json=resourceNames,proto3" json:"resource_names,omitempty"`
}

func (x *PolicyRule) Reset() {
	*x = PolicyRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_rbac_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyRule) ProtoMessage() {}

func (x *PolicyRule) ProtoReflect() protoreflect.Message {
	mi := &file_storage_rbac_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyRule.ProtoReflect.Descriptor instead.
func (*PolicyRule) Descriptor() ([]byte, []int) {
	return file_storage_rbac_proto_rawDescGZIP(), []int{1}
}

func (x *PolicyRule) GetVerbs() []string {
	if x != nil {
		return x.Verbs
	}
	return nil
}

func (x *PolicyRule) GetApiGroups() []string {
	if x != nil {
		return x.ApiGroups
	}
	return nil
}

func (x *PolicyRule) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *PolicyRule) GetNonResourceUrls() []string {
	if x != nil {
		return x.NonResourceUrls
	}
	return nil
}

func (x *PolicyRule) GetResourceNames() []string {
	if x != nil {
		return x.ResourceNames
	}
	return nil
}

// Properties of an individual k8s RoleBinding or ClusterRoleBinding.
// ////////////////////////////////////////
type K8SRoleBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" search:"Role Binding ID,hidden" sql:"pk,type(uuid)"`                                      // @gotags: search:"Role Binding ID,hidden" sql:"pk,type(uuid)"
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" search:"Role Binding"`                                  // @gotags: search:"Role Binding"
	Namespace   string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty" search:"Namespace,store"`                        // @gotags: search:"Namespace,store"
	ClusterId   string `protobuf:"bytes,4,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" search:"Cluster ID,store,hidden" sql:"type(uuid)"`       // @gotags: search:"Cluster ID,store,hidden" sql:"type(uuid)"
	ClusterName string `protobuf:"bytes,5,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty" search:"Cluster"` // @gotags: search:"Cluster"
	// ClusterRole specifies whether the binding binds a cluster role. However, it cannot be used to determine whether
	// the binding is a cluster role binding. This can be done in conjunction with the namespace. If the namespace is
	// empty and cluster role is true, the binding is a cluster role binding.
	ClusterRole bool                   `protobuf:"varint,6,opt,name=cluster_role,json=clusterRole,proto3" json:"cluster_role,omitempty" search:"Cluster Role"`                                                                     // @gotags: search:"Cluster Role"
	Labels      map[string]string      `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" search:"Role Binding Label"`           // @gotags: search:"Role Binding Label"
	Annotations map[string]string      `protobuf:"bytes,8,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" search:"Role Binding Annotation"` // @gotags: search:"Role Binding Annotation"
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Subjects    []*Subject             `protobuf:"bytes,10,rep,name=subjects,proto3" json:"subjects,omitempty"`
	RoleId      string                 `protobuf:"bytes,11,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty" search:"Role ID" sql:"type(uuid)"` // @gotags: search:"Role ID" sql:"type(uuid)"
}

func (x *K8SRoleBinding) Reset() {
	*x = K8SRoleBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_rbac_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8SRoleBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SRoleBinding) ProtoMessage() {}

func (x *K8SRoleBinding) ProtoReflect() protoreflect.Message {
	mi := &file_storage_rbac_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SRoleBinding.ProtoReflect.Descriptor instead.
func (*K8SRoleBinding) Descriptor() ([]byte, []int) {
	return file_storage_rbac_proto_rawDescGZIP(), []int{2}
}

func (x *K8SRoleBinding) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *K8SRoleBinding) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *K8SRoleBinding) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *K8SRoleBinding) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *K8SRoleBinding) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *K8SRoleBinding) GetClusterRole() bool {
	if x != nil {
		return x.ClusterRole
	}
	return false
}

func (x *K8SRoleBinding) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *K8SRoleBinding) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *K8SRoleBinding) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *K8SRoleBinding) GetSubjects() []*Subject {
	if x != nil {
		return x.Subjects
	}
	return nil
}

func (x *K8SRoleBinding) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

// Properties of an individual subjects who are granted roles via role bindings.
// ////////////////////////////////////////
type Subject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string      `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`                               // ID is derived from base64 of cluster id and name
	Kind        SubjectKind `protobuf:"varint,1,opt,name=kind,proto3,enum=storage.SubjectKind" json:"kind,omitempty" search:"Subject Kind"` // @gotags: search:"Subject Kind"
	Name        string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" search:"Subject"`                           // @gotags: search:"Subject"
	Namespace   string      `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ClusterId   string      `protobuf:"bytes,5,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	ClusterName string      `protobuf:"bytes,6,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
}

func (x *Subject) Reset() {
	*x = Subject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_rbac_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subject) ProtoMessage() {}

func (x *Subject) ProtoReflect() protoreflect.Message {
	mi := &file_storage_rbac_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subject.ProtoReflect.Descriptor instead.
func (*Subject) Descriptor() ([]byte, []int) {
	return file_storage_rbac_proto_rawDescGZIP(), []int{3}
}

func (x *Subject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Subject) GetKind() SubjectKind {
	if x != nil {
		return x.Kind
	}
	return SubjectKind_UNSET_KIND
}

func (x *Subject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Subject) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Subject) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *Subject) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

var File_storage_rbac_proto protoreflect.FileDescriptor

var file_storage_rbac_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c,
	0x04, 0x0a, 0x07, 0x4b, 0x38, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x34, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4b, 0x38, 0x73, 0x52,
	0x6f, 0x6c, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x43, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4b, 0x38, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x2e, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a,
	0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb2, 0x01,
	0x0a, 0x0a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x65, 0x72, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x65, 0x72,
	0x62, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x69, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x2a, 0x0a, 0x11, 0x6e, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x22, 0xbd, 0x04, 0x0a, 0x0e, 0x4b, 0x38, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4b, 0x38, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x4a, 0x0a, 0x0b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4b, 0x38, 0x73, 0x52, 0x6f, 0x6c, 0x65,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x2c, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xb7, 0x01, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x69,
	0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x47, 0x0a, 0x0b,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x0a, 0x55,
	0x4e, 0x53, 0x45, 0x54, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52,
	0x4f, 0x55, 0x50, 0x10, 0x03, 0x2a, 0x7c, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e, 0x53, 0x45,
	0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x4c,
	0x45, 0x56, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50,
	0x41, 0x43, 0x45, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x4c, 0x45, 0x56, 0x41, 0x54, 0x45,
	0x44, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x57, 0x49, 0x44, 0x45, 0x10, 0x04,
	0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x41, 0x44, 0x4d, 0x49,
	0x4e, 0x10, 0x05, 0x42, 0x27, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72,
	0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_rbac_proto_rawDescOnce sync.Once
	file_storage_rbac_proto_rawDescData = file_storage_rbac_proto_rawDesc
)

func file_storage_rbac_proto_rawDescGZIP() []byte {
	file_storage_rbac_proto_rawDescOnce.Do(func() {
		file_storage_rbac_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_rbac_proto_rawDescData)
	})
	return file_storage_rbac_proto_rawDescData
}

var file_storage_rbac_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_storage_rbac_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_storage_rbac_proto_goTypes = []interface{}{
	(SubjectKind)(0),              // 0: storage.SubjectKind
	(PermissionLevel)(0),          // 1: storage.PermissionLevel
	(*K8SRole)(nil),               // 2: storage.K8sRole
	(*PolicyRule)(nil),            // 3: storage.PolicyRule
	(*K8SRoleBinding)(nil),        // 4: storage.K8sRoleBinding
	(*Subject)(nil),               // 5: storage.Subject
	nil,                           // 6: storage.K8sRole.LabelsEntry
	nil,                           // 7: storage.K8sRole.AnnotationsEntry
	nil,                           // 8: storage.K8sRoleBinding.LabelsEntry
	nil,                           // 9: storage.K8sRoleBinding.AnnotationsEntry
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_storage_rbac_proto_depIdxs = []int32{
	6,  // 0: storage.K8sRole.labels:type_name -> storage.K8sRole.LabelsEntry
	7,  // 1: storage.K8sRole.annotations:type_name -> storage.K8sRole.AnnotationsEntry
	10, // 2: storage.K8sRole.created_at:type_name -> google.protobuf.Timestamp
	3,  // 3: storage.K8sRole.rules:type_name -> storage.PolicyRule
	8,  // 4: storage.K8sRoleBinding.labels:type_name -> storage.K8sRoleBinding.LabelsEntry
	9,  // 5: storage.K8sRoleBinding.annotations:type_name -> storage.K8sRoleBinding.AnnotationsEntry
	10, // 6: storage.K8sRoleBinding.created_at:type_name -> google.protobuf.Timestamp
	5,  // 7: storage.K8sRoleBinding.subjects:type_name -> storage.Subject
	0,  // 8: storage.Subject.kind:type_name -> storage.SubjectKind
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_storage_rbac_proto_init() }
func file_storage_rbac_proto_init() {
	if File_storage_rbac_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_rbac_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*K8SRole); i {
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
		file_storage_rbac_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyRule); i {
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
		file_storage_rbac_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*K8SRoleBinding); i {
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
		file_storage_rbac_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subject); i {
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
			RawDescriptor: file_storage_rbac_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_rbac_proto_goTypes,
		DependencyIndexes: file_storage_rbac_proto_depIdxs,
		EnumInfos:         file_storage_rbac_proto_enumTypes,
		MessageInfos:      file_storage_rbac_proto_msgTypes,
	}.Build()
	File_storage_rbac_proto = out.File
	file_storage_rbac_proto_rawDesc = nil
	file_storage_rbac_proto_goTypes = nil
	file_storage_rbac_proto_depIdxs = nil
}
