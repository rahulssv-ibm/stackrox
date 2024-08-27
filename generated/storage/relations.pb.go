// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: storage/relations.proto

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

type ImageComponentEdge struct {
	// / Layer that contains this component
	//
	// Types that are assignable to HasLayerIndex:
	//
	//	*ImageComponentEdge_LayerIndex
	HasLayerIndex isImageComponentEdge_HasLayerIndex `protobuf_oneof:"has_layer_index"`
	state         protoimpl.MessageState

	// id is base 64 encoded Image:Component ids.
	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,id"`                                                                                                   // @gotags: sql:"pk,id"
	Location         string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty" search:"Component Location,store,hidden"`                                                          // @gotags: search:"Component Location,store,hidden"
	ImageId          string `protobuf:"bytes,4,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty" sql:"fk(Image:id),index=hash"`                                                        // @gotags: sql:"fk(Image:id),index=hash"
	ImageComponentId string `protobuf:"bytes,5,opt,name=image_component_id,json=imageComponentId,proto3" json:"image_component_id,omitempty" sql:"fk(ImageComponent:id),no-fk-constraint,index=hash"` // @gotags: sql:"fk(ImageComponent:id),no-fk-constraint,index=hash"
	unknownFields    protoimpl.UnknownFields

	sizeCache protoimpl.SizeCache
}

func (x *ImageComponentEdge) Reset() {
	*x = ImageComponentEdge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_relations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageComponentEdge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageComponentEdge) ProtoMessage() {}

func (x *ImageComponentEdge) ProtoReflect() protoreflect.Message {
	mi := &file_storage_relations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageComponentEdge.ProtoReflect.Descriptor instead.
func (*ImageComponentEdge) Descriptor() ([]byte, []int) {
	return file_storage_relations_proto_rawDescGZIP(), []int{0}
}

func (x *ImageComponentEdge) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *ImageComponentEdge) GetHasLayerIndex() isImageComponentEdge_HasLayerIndex {
	if m != nil {
		return m.HasLayerIndex
	}
	return nil
}

func (x *ImageComponentEdge) GetLayerIndex() int32 {
	if x, ok := x.GetHasLayerIndex().(*ImageComponentEdge_LayerIndex); ok {
		return x.LayerIndex
	}
	return 0
}

func (x *ImageComponentEdge) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *ImageComponentEdge) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *ImageComponentEdge) GetImageComponentId() string {
	if x != nil {
		return x.ImageComponentId
	}
	return ""
}

type isImageComponentEdge_HasLayerIndex interface {
	isImageComponentEdge_HasLayerIndex()
}

type ImageComponentEdge_LayerIndex struct {
	LayerIndex int32 `protobuf:"varint,2,opt,name=layer_index,json=layerIndex,proto3,oneof"`
}

func (*ImageComponentEdge_LayerIndex) isImageComponentEdge_HasLayerIndex() {}

type ComponentCVEEdge struct {
	// Whether there is a version the CVE is fixed in the component.
	//
	// Types that are assignable to HasFixedBy:
	//
	//	*ComponentCVEEdge_FixedBy
	HasFixedBy isComponentCVEEdge_HasFixedBy `protobuf_oneof:"has_fixed_by"`
	state      protoimpl.MessageState

	// base 64 encoded Component:CVE ids.
	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,id"`                                                                                  // @gotags: sql:"pk,id"
	ImageComponentId string `protobuf:"bytes,4,opt,name=image_component_id,json=imageComponentId,proto3" json:"image_component_id,omitempty" sql:"fk(ImageComponent:id),index=hash"` // @gotags: sql:"fk(ImageComponent:id),index=hash"
	ImageCveId       string `protobuf:"bytes,5,opt,name=image_cve_id,json=imageCveId,proto3" json:"image_cve_id,omitempty" sql:"fk(ImageCVE:id),no-fk-constraint,index=hash"`        // @gotags: sql:"fk(ImageCVE:id),no-fk-constraint,index=hash"
	unknownFields    protoimpl.UnknownFields

	sizeCache protoimpl.SizeCache
	IsFixable bool `protobuf:"varint,2,opt,name=is_fixable,json=isFixable,proto3" json:"is_fixable,omitempty" search:"Fixable,store"` // @gotags: search:"Fixable,store"
}

func (x *ComponentCVEEdge) Reset() {
	*x = ComponentCVEEdge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_relations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentCVEEdge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentCVEEdge) ProtoMessage() {}

func (x *ComponentCVEEdge) ProtoReflect() protoreflect.Message {
	mi := &file_storage_relations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentCVEEdge.ProtoReflect.Descriptor instead.
func (*ComponentCVEEdge) Descriptor() ([]byte, []int) {
	return file_storage_relations_proto_rawDescGZIP(), []int{1}
}

func (x *ComponentCVEEdge) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ComponentCVEEdge) GetIsFixable() bool {
	if x != nil {
		return x.IsFixable
	}
	return false
}

func (m *ComponentCVEEdge) GetHasFixedBy() isComponentCVEEdge_HasFixedBy {
	if m != nil {
		return m.HasFixedBy
	}
	return nil
}

func (x *ComponentCVEEdge) GetFixedBy() string {
	if x, ok := x.GetHasFixedBy().(*ComponentCVEEdge_FixedBy); ok {
		return x.FixedBy
	}
	return ""
}

func (x *ComponentCVEEdge) GetImageComponentId() string {
	if x != nil {
		return x.ImageComponentId
	}
	return ""
}

func (x *ComponentCVEEdge) GetImageCveId() string {
	if x != nil {
		return x.ImageCveId
	}
	return ""
}

type isComponentCVEEdge_HasFixedBy interface {
	isComponentCVEEdge_HasFixedBy()
}

type ComponentCVEEdge_FixedBy struct {
	FixedBy string `protobuf:"bytes,3,opt,name=fixed_by,json=fixedBy,proto3,oneof" search:"Fixed By,store,hidden"` // @gotags: search:"Fixed By,store,hidden"
}

func (*ComponentCVEEdge_FixedBy) isComponentCVEEdge_HasFixedBy() {}

type ImageCVEEdge struct {
	state                protoimpl.MessageState
	FirstImageOccurrence *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=first_image_occurrence,json=firstImageOccurrence,proto3" json:"first_image_occurrence,omitempty" search:"First Image Occurrence Timestamp,hidden"` // @gotags: search:"First Image Occurrence Timestamp,hidden"

	// base 64 encoded Image:CVE ids.
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,id"`                                                                           // @gotags: sql:"pk,id"
	ImageId       string `protobuf:"bytes,4,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty" sql:"fk(Image:id),index=hash"`                                // @gotags: sql:"fk(Image:id),index=hash"
	ImageCveId    string `protobuf:"bytes,5,opt,name=image_cve_id,json=imageCveId,proto3" json:"image_cve_id,omitempty" sql:"fk(ImageCVE:id),no-fk-constraint,index=hash"` // @gotags: sql:"fk(ImageCVE:id),no-fk-constraint,index=hash"
	unknownFields protoimpl.UnknownFields

	sizeCache protoimpl.SizeCache
	State     VulnerabilityState `protobuf:"varint,3,opt,name=state,proto3,enum=storage.VulnerabilityState" json:"state,omitempty" search:"Vulnerability State"` // @gotags: search:"Vulnerability State"
}

func (x *ImageCVEEdge) Reset() {
	*x = ImageCVEEdge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_relations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageCVEEdge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageCVEEdge) ProtoMessage() {}

func (x *ImageCVEEdge) ProtoReflect() protoreflect.Message {
	mi := &file_storage_relations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageCVEEdge.ProtoReflect.Descriptor instead.
func (*ImageCVEEdge) Descriptor() ([]byte, []int) {
	return file_storage_relations_proto_rawDescGZIP(), []int{2}
}

func (x *ImageCVEEdge) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ImageCVEEdge) GetFirstImageOccurrence() *timestamppb.Timestamp {
	if x != nil {
		return x.FirstImageOccurrence
	}
	return nil
}

func (x *ImageCVEEdge) GetState() VulnerabilityState {
	if x != nil {
		return x.State
	}
	return VulnerabilityState_OBSERVED
}

func (x *ImageCVEEdge) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *ImageCVEEdge) GetImageCveId() string {
	if x != nil {
		return x.ImageCveId
	}
	return ""
}

type NodeComponentEdge struct {
	state protoimpl.MessageState

	// base 64 encoded Node:Component ids.
	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,id"`                                                                                               // @gotags: sql:"pk,id"
	NodeId          string `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty" sql:"fk(Node:id),index=hash,type(uuid)"`                                             // @gotags: sql:"fk(Node:id),index=hash,type(uuid)"
	NodeComponentId string `protobuf:"bytes,3,opt,name=node_component_id,json=nodeComponentId,proto3" json:"node_component_id,omitempty" sql:"fk(NodeComponent:id),no-fk-constraint,index=hash"` // @gotags: sql:"fk(NodeComponent:id),no-fk-constraint,index=hash"
	unknownFields   protoimpl.UnknownFields

	sizeCache protoimpl.SizeCache
}

func (x *NodeComponentEdge) Reset() {
	*x = NodeComponentEdge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_relations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeComponentEdge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeComponentEdge) ProtoMessage() {}

func (x *NodeComponentEdge) ProtoReflect() protoreflect.Message {
	mi := &file_storage_relations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeComponentEdge.ProtoReflect.Descriptor instead.
func (*NodeComponentEdge) Descriptor() ([]byte, []int) {
	return file_storage_relations_proto_rawDescGZIP(), []int{3}
}

func (x *NodeComponentEdge) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NodeComponentEdge) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *NodeComponentEdge) GetNodeComponentId() string {
	if x != nil {
		return x.NodeComponentId
	}
	return ""
}

type NodeComponentCVEEdge struct {
	// Whether there is a version the CVE is fixed in the component.
	//
	// Types that are assignable to HasFixedBy:
	//
	//	*NodeComponentCVEEdge_FixedBy
	HasFixedBy isNodeComponentCVEEdge_HasFixedBy `protobuf_oneof:"has_fixed_by"`
	state      protoimpl.MessageState

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,id"`                                                                              // @gotags: sql:"pk,id"
	NodeComponentId string `protobuf:"bytes,4,opt,name=node_component_id,json=nodeComponentId,proto3" json:"node_component_id,omitempty" sql:"fk(NodeComponent:id),index=hash"` // @gotags: sql:"fk(NodeComponent:id),index=hash"
	NodeCveId       string `protobuf:"bytes,5,opt,name=node_cve_id,json=nodeCveId,proto3" json:"node_cve_id,omitempty" sql:"fk(NodeCVE:id),no-fk-constraint,index=hash"`        // @gotags: sql:"fk(NodeCVE:id),no-fk-constraint,index=hash"
	unknownFields   protoimpl.UnknownFields

	sizeCache protoimpl.SizeCache
	IsFixable bool `protobuf:"varint,2,opt,name=is_fixable,json=isFixable,proto3" json:"is_fixable,omitempty" search:"Fixable,store"` // @gotags: search:"Fixable,store"
}

func (x *NodeComponentCVEEdge) Reset() {
	*x = NodeComponentCVEEdge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_relations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeComponentCVEEdge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeComponentCVEEdge) ProtoMessage() {}

func (x *NodeComponentCVEEdge) ProtoReflect() protoreflect.Message {
	mi := &file_storage_relations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeComponentCVEEdge.ProtoReflect.Descriptor instead.
func (*NodeComponentCVEEdge) Descriptor() ([]byte, []int) {
	return file_storage_relations_proto_rawDescGZIP(), []int{4}
}

func (x *NodeComponentCVEEdge) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NodeComponentCVEEdge) GetIsFixable() bool {
	if x != nil {
		return x.IsFixable
	}
	return false
}

func (m *NodeComponentCVEEdge) GetHasFixedBy() isNodeComponentCVEEdge_HasFixedBy {
	if m != nil {
		return m.HasFixedBy
	}
	return nil
}

func (x *NodeComponentCVEEdge) GetFixedBy() string {
	if x, ok := x.GetHasFixedBy().(*NodeComponentCVEEdge_FixedBy); ok {
		return x.FixedBy
	}
	return ""
}

func (x *NodeComponentCVEEdge) GetNodeComponentId() string {
	if x != nil {
		return x.NodeComponentId
	}
	return ""
}

func (x *NodeComponentCVEEdge) GetNodeCveId() string {
	if x != nil {
		return x.NodeCveId
	}
	return ""
}

type isNodeComponentCVEEdge_HasFixedBy interface {
	isNodeComponentCVEEdge_HasFixedBy()
}

type NodeComponentCVEEdge_FixedBy struct {
	FixedBy string `protobuf:"bytes,3,opt,name=fixed_by,json=fixedBy,proto3,oneof" search:"Fixed By,store,hidden"` // @gotags: search:"Fixed By,store,hidden"
}

func (*NodeComponentCVEEdge_FixedBy) isNodeComponentCVEEdge_HasFixedBy() {}

type ClusterCVEEdge struct {
	// Whether there is a version the CVE is fixed in the Cluster.
	//
	// Types that are assignable to HasFixedBy:
	//
	//	*ClusterCVEEdge_FixedBy
	HasFixedBy isClusterCVEEdge_HasFixedBy `protobuf_oneof:"has_fixed_by"`
	state      protoimpl.MessageState

	// base 64 encoded Cluster:CVE ids.
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,id"`                                                            // @gotags: sql:"pk,id"
	ClusterId     string `protobuf:"bytes,4,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" sql:"fk(Cluster:id),type(uuid)"`         // @gotags: sql:"fk(Cluster:id),type(uuid)"
	CveId         string `protobuf:"bytes,5,opt,name=cve_id,json=cveId,proto3" json:"cve_id,omitempty" sql:"fk(ClusterCVE:id),no-fk-constraint,index=hash"` // @gotags: sql:"fk(ClusterCVE:id),no-fk-constraint,index=hash"
	unknownFields protoimpl.UnknownFields

	sizeCache protoimpl.SizeCache
	IsFixable bool `protobuf:"varint,2,opt,name=is_fixable,json=isFixable,proto3" json:"is_fixable,omitempty" search:"Cluster CVE Fixable,store,hidden"` // @gotags: search:"Cluster CVE Fixable,store,hidden"
}

func (x *ClusterCVEEdge) Reset() {
	*x = ClusterCVEEdge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_relations_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterCVEEdge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterCVEEdge) ProtoMessage() {}

func (x *ClusterCVEEdge) ProtoReflect() protoreflect.Message {
	mi := &file_storage_relations_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterCVEEdge.ProtoReflect.Descriptor instead.
func (*ClusterCVEEdge) Descriptor() ([]byte, []int) {
	return file_storage_relations_proto_rawDescGZIP(), []int{5}
}

func (x *ClusterCVEEdge) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClusterCVEEdge) GetIsFixable() bool {
	if x != nil {
		return x.IsFixable
	}
	return false
}

func (m *ClusterCVEEdge) GetHasFixedBy() isClusterCVEEdge_HasFixedBy {
	if m != nil {
		return m.HasFixedBy
	}
	return nil
}

func (x *ClusterCVEEdge) GetFixedBy() string {
	if x, ok := x.GetHasFixedBy().(*ClusterCVEEdge_FixedBy); ok {
		return x.FixedBy
	}
	return ""
}

func (x *ClusterCVEEdge) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ClusterCVEEdge) GetCveId() string {
	if x != nil {
		return x.CveId
	}
	return ""
}

type isClusterCVEEdge_HasFixedBy interface {
	isClusterCVEEdge_HasFixedBy()
}

type ClusterCVEEdge_FixedBy struct {
	FixedBy string `protobuf:"bytes,3,opt,name=fixed_by,json=fixedBy,proto3,oneof" search:"Cluster CVE Fixed By,store,hidden"` // @gotags: search:"Cluster CVE Fixed By,store,hidden"
}

func (*ClusterCVEEdge_FixedBy) isClusterCVEEdge_HasFixedBy() {}

type PolicyCategoryEdge struct {
	state protoimpl.MessageState

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,id"`                                                                     // @gotags: sql:"pk,id"
	PolicyId      string `protobuf:"bytes,2,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty" sql:"fk(Policy:id)" search:"Policy ID,store,hidden"` // @gotags: sql:"fk(Policy:id)" search:"Policy ID,store,hidden"
	CategoryId    string `protobuf:"bytes,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty" sql:"fk(PolicyCategory:id)"`                   // @gotags: sql:"fk(PolicyCategory:id)"
	unknownFields protoimpl.UnknownFields

	sizeCache protoimpl.SizeCache
}

func (x *PolicyCategoryEdge) Reset() {
	*x = PolicyCategoryEdge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_relations_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyCategoryEdge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyCategoryEdge) ProtoMessage() {}

func (x *PolicyCategoryEdge) ProtoReflect() protoreflect.Message {
	mi := &file_storage_relations_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyCategoryEdge.ProtoReflect.Descriptor instead.
func (*PolicyCategoryEdge) Descriptor() ([]byte, []int) {
	return file_storage_relations_proto_rawDescGZIP(), []int{6}
}

func (x *PolicyCategoryEdge) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PolicyCategoryEdge) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

func (x *PolicyCategoryEdge) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

var File_storage_relations_proto protoreflect.FileDescriptor

var file_storage_relations_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x76, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x12, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x45, 0x64, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a,
	0x0b, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x68, 0x61, 0x73, 0x5f, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0xbe, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x43, 0x56, 0x45, 0x45, 0x64, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x46, 0x69, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x08,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x42, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x63, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x43, 0x76, 0x65, 0x49, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x68, 0x61, 0x73,
	0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x22, 0xe0, 0x01, 0x0a, 0x0c, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x43, 0x56, 0x45, 0x45, 0x64, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x50, 0x0a, 0x16, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x14, 0x66, 0x69, 0x72, 0x73, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x63, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x76, 0x65, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x11,
	0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x45, 0x64, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xbe, 0x01, 0x0a, 0x14, 0x4e, 0x6f, 0x64, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x43, 0x56, 0x45, 0x45, 0x64, 0x67, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x46, 0x69, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1b,
	0x0a, 0x08, 0x66, 0x69, 0x78, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x42, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x63, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x6f,
	0x64, 0x65, 0x43, 0x76, 0x65, 0x49, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x5f, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x22, 0xa2, 0x01, 0x0a, 0x0e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x56, 0x45, 0x45, 0x64, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73,
	0x5f, 0x66, 0x69, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x46, 0x69, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x76, 0x65, 0x49, 0x64, 0x42, 0x0e, 0x0a, 0x0c,
	0x68, 0x61, 0x73, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x22, 0x62, 0x0a, 0x12,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x45, 0x64,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_relations_proto_rawDescOnce sync.Once
	file_storage_relations_proto_rawDescData = file_storage_relations_proto_rawDesc
)

func file_storage_relations_proto_rawDescGZIP() []byte {
	file_storage_relations_proto_rawDescOnce.Do(func() {
		file_storage_relations_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_relations_proto_rawDescData)
	})
	return file_storage_relations_proto_rawDescData
}

var file_storage_relations_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_storage_relations_proto_goTypes = []any{
	(*ImageComponentEdge)(nil),    // 0: storage.ImageComponentEdge
	(*ComponentCVEEdge)(nil),      // 1: storage.ComponentCVEEdge
	(*ImageCVEEdge)(nil),          // 2: storage.ImageCVEEdge
	(*NodeComponentEdge)(nil),     // 3: storage.NodeComponentEdge
	(*NodeComponentCVEEdge)(nil),  // 4: storage.NodeComponentCVEEdge
	(*ClusterCVEEdge)(nil),        // 5: storage.ClusterCVEEdge
	(*PolicyCategoryEdge)(nil),    // 6: storage.PolicyCategoryEdge
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(VulnerabilityState)(0),       // 8: storage.VulnerabilityState
}
var file_storage_relations_proto_depIdxs = []int32{
	7, // 0: storage.ImageCVEEdge.first_image_occurrence:type_name -> google.protobuf.Timestamp
	8, // 1: storage.ImageCVEEdge.state:type_name -> storage.VulnerabilityState
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_storage_relations_proto_init() }
func file_storage_relations_proto_init() {
	if File_storage_relations_proto != nil {
		return
	}
	file_storage_cve_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_storage_relations_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ImageComponentEdge); i {
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
		file_storage_relations_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ComponentCVEEdge); i {
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
		file_storage_relations_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ImageCVEEdge); i {
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
		file_storage_relations_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*NodeComponentEdge); i {
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
		file_storage_relations_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*NodeComponentCVEEdge); i {
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
		file_storage_relations_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ClusterCVEEdge); i {
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
		file_storage_relations_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*PolicyCategoryEdge); i {
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
	file_storage_relations_proto_msgTypes[0].OneofWrappers = []any{
		(*ImageComponentEdge_LayerIndex)(nil),
	}
	file_storage_relations_proto_msgTypes[1].OneofWrappers = []any{
		(*ComponentCVEEdge_FixedBy)(nil),
	}
	file_storage_relations_proto_msgTypes[4].OneofWrappers = []any{
		(*NodeComponentCVEEdge_FixedBy)(nil),
	}
	file_storage_relations_proto_msgTypes[5].OneofWrappers = []any{
		(*ClusterCVEEdge_FixedBy)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_relations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_relations_proto_goTypes,
		DependencyIndexes: file_storage_relations_proto_depIdxs,
		MessageInfos:      file_storage_relations_proto_msgTypes,
	}.Build()
	File_storage_relations_proto = out.File
	file_storage_relations_proto_rawDesc = nil
	file_storage_relations_proto_goTypes = nil
	file_storage_relations_proto_depIdxs = nil
}
