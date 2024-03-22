// Code generated by pg-bindings generator. DO NOT EDIT.
package schema

import (
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/protocompat"
)

// ConvertNotifierFromProto converts a `*storage.Notifier` to Gorm model
func ConvertNotifierFromProto(obj *storage.Notifier) (*Notifiers, error) {
	serialized, err := protocompat.Marshal(obj)
	if err != nil {
		return nil, err
	}
	model := &Notifiers{
		ID:         obj.GetId(),
		Name:       obj.GetName(),
		Serialized: serialized,
	}
	return model, nil
}

// ConvertNotifierToProto converts Gorm model `Notifiers` to its protobuf type object
func ConvertNotifierToProto(m *Notifiers) (*storage.Notifier, error) {
	var msg storage.Notifier
	if err := msg.Unmarshal(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}
