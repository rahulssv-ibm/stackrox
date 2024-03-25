// Code generated by pg-bindings generator. DO NOT EDIT.
package schema

import (
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/protocompat"
)

// ConvertTestParent2FromProto converts a `*storage.TestParent2` to Gorm model
func ConvertTestParent2FromProto(obj *storage.TestParent2) (*TestParent2, error) {
	serialized, err := obj.Marshal()
	if err != nil {
		return nil, err
	}
	model := &TestParent2{
		ID:         obj.GetId(),
		ParentID:   obj.GetParentId(),
		Val:        obj.GetVal(),
		Serialized: serialized,
	}
	return model, nil
}

// ConvertTestParent2ToProto converts Gorm model `TestParent2` to its protobuf type object
func ConvertTestParent2ToProto(m *TestParent2) (*storage.TestParent2, error) {
	var msg storage.TestParent2
	if err := protocompat.Unmarshal(m.Serialized, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
