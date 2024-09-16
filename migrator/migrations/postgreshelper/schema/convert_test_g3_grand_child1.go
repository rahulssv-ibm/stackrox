// Code generated by pg-bindings generator. DO NOT EDIT.
package schema

import (
	"github.com/stackrox/rox/generated/storage"
)

// ConvertTestG3GrandChild1FromProto converts a `*storage.TestG3GrandChild1` to Gorm model
func ConvertTestG3GrandChild1FromProto(obj *storage.TestG3GrandChild1) (*TestG3GrandChild1, error) {
	serialized, err := obj.MarshalVT()
	if err != nil {
		return nil, err
	}
	model := &TestG3GrandChild1{
		ID:         obj.GetId(),
		Val:        obj.GetVal(),
		Serialized: serialized,
	}
	return model, nil
}

// ConvertTestG3GrandChild1ToProto converts Gorm model `TestG3GrandChild1` to its protobuf type object
func ConvertTestG3GrandChild1ToProto(m *TestG3GrandChild1) (*storage.TestG3GrandChild1, error) {
	var msg storage.TestG3GrandChild1
	if err := msg.UnmarshalVTUnsafe(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}
