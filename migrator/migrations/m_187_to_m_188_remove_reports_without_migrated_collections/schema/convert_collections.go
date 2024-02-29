package schema

import (
	"github.com/stackrox/rox/generated/storage"
)

// ConvertResourceCollectionFromProto converts a `*storage.ResourceCollection` to Gorm model
func ConvertResourceCollectionFromProto(obj *storage.ResourceCollection) (*Collections, error) {
	serialized, err := obj.Marshal()
	if err != nil {
		return nil, err
	}
	model := &Collections{
		ID:            obj.GetId(),
		Name:          obj.GetName(),
		CreatedByName: obj.GetCreatedBy().GetName(),
		UpdatedByName: obj.GetUpdatedBy().GetName(),
		Serialized:    serialized,
	}
	return model, nil
}

// ConvertResourceCollection_EmbeddedResourceCollectionFromProto converts a `*storage.ResourceCollection_EmbeddedResourceCollection` to Gorm model
func ConvertResourceCollection_EmbeddedResourceCollectionFromProto(obj *storage.ResourceCollection_EmbeddedResourceCollection, idx int, collectionID string) (*CollectionsEmbeddedCollections, error) {
	model := &CollectionsEmbeddedCollections{
		CollectionsID: collectionID,
		Idx:           idx,
		ID:            obj.GetId(),
	}
	return model, nil
}

// ConvertResourceCollectionToProto converts Gorm model `Collections` to its protobuf type object
func ConvertResourceCollectionToProto(m *Collections) (*storage.ResourceCollection, error) {
	var msg storage.ResourceCollection
	if err := msg.Unmarshal(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}
