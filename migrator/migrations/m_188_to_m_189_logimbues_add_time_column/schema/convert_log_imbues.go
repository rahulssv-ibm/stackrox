// Code originally generated by pg-bindings generator.

package schema

import (
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/protocompat"
)

// ConvertLogImbueFromProto converts a `*storage.LogImbue` to Gorm model
func ConvertLogImbueFromProto(obj *storage.LogImbue) (*LogImbues, error) {
	serialized, err := obj.Marshal()
	if err != nil {
		return nil, err
	}
	model := &LogImbues{
		ID:         obj.GetId(),
		Timestamp:  protocompat.NilOrNow(obj.GetTimestamp()),
		Serialized: serialized,
	}
	return model, nil
}

// ConvertLogImbueToProto converts Gorm model `LogImbues` to its protobuf type object
func ConvertLogImbueToProto(m *LogImbues) (*storage.LogImbue, error) {
	var msg storage.LogImbue
	if err := msg.Unmarshal(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}
