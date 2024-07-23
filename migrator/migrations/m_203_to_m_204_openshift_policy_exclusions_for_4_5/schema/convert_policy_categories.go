// Code generated by pg-bindings generator. DO NOT EDIT.
package schema

import (
	"github.com/stackrox/rox/generated/storage"
)

// ConvertPolicyCategoryFromProto converts a `*storage.PolicyCategory` to Gorm model
func ConvertPolicyCategoryFromProto(obj *storage.PolicyCategory) (*PolicyCategories, error) {
	serialized, err := obj.MarshalVT()
	if err != nil {
		return nil, err
	}
	model := &PolicyCategories{
		ID:                 obj.GetId(),
		Name:               obj.GetName(),
		Serialized:         serialized,
	}
	return model, nil
}

// ConvertPolicyCategoryToProto converts Gorm model `PolicyCategories` to its protobuf type object
func ConvertPolicyCategoryToProto(m *PolicyCategories) (*storage.PolicyCategory, error) {
	var msg storage.PolicyCategory
	if err := msg.UnmarshalVT(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}
