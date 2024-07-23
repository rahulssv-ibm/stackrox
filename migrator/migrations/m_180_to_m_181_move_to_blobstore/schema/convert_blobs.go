// Code originally generated by pg-bindings generator.

package schema

import (
	"github.com/stackrox/rox/generated/storage"
)

// ConvertBlobFromProto converts a `*storage.Blob` to Gorm model
func ConvertBlobFromProto(obj *storage.Blob) (*Blobs, error) {
	serialized, err := obj.MarshalVT()
	if err != nil {
		return nil, err
	}
	model := &Blobs{
		Name:       obj.GetName(),
		Serialized: serialized,
	}
	return model, nil
}

// ConvertBlobToProto converts Gorm model `Blobs` to its protobuf type object
func ConvertBlobToProto(m *Blobs) (*storage.Blob, error) {
	var msg storage.Blob
	if err := msg.UnmarshalVT(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}
