// Code generated by pg-bindings generator. DO NOT EDIT.
package schema

import (
	"github.com/lib/pq"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/protocompat"
)

// ConvertImageFromProto converts a `*storage.Image` to Gorm model
func ConvertImageFromProto(obj *storage.Image) (*Images, error) {
	serialized, err := obj.Marshal()
	if err != nil {
		return nil, err
	}
	model := &Images{
		ID:                   obj.GetId(),
		NameRegistry:         obj.GetName().GetRegistry(),
		NameRemote:           obj.GetName().GetRemote(),
		NameTag:              obj.GetName().GetTag(),
		NameFullName:         obj.GetName().GetFullName(),
		MetadataV1Created:    protocompat.NilOrTime(obj.GetMetadata().GetV1().GetCreated()),
		MetadataV1User:       obj.GetMetadata().GetV1().GetUser(),
		MetadataV1Command:    pq.Array(obj.GetMetadata().GetV1().GetCommand()).(*pq.StringArray),
		MetadataV1Entrypoint: pq.Array(obj.GetMetadata().GetV1().GetEntrypoint()).(*pq.StringArray),
		MetadataV1Volumes:    pq.Array(obj.GetMetadata().GetV1().GetVolumes()).(*pq.StringArray),
		MetadataV1Labels:     obj.GetMetadata().GetV1().GetLabels(),
		ScanScanTime:         protocompat.NilOrTime(obj.GetScan().GetScanTime()),
		ScanOperatingSystem:  obj.GetScan().GetOperatingSystem(),
		SignatureFetched:     protocompat.NilOrTime(obj.GetSignature().GetFetched()),
		Components:           obj.GetComponents(),
		Cves:                 obj.GetCves(),
		FixableCves:          obj.GetFixableCves(),
		LastUpdated:          protocompat.NilOrTime(obj.GetLastUpdated()),
		Priority:             obj.GetPriority(),
		RiskScore:            obj.GetRiskScore(),
		TopCvss:              obj.GetTopCvss(),
		Serialized:           serialized,
	}
	return model, nil
}

// ConvertImageLayerFromProto converts a `*storage.ImageLayer` to Gorm model
func ConvertImageLayerFromProto(obj *storage.ImageLayer, idx int, imageID string) (*ImagesLayers, error) {
	model := &ImagesLayers{
		ImagesID:    imageID,
		Idx:         idx,
		Instruction: obj.GetInstruction(),
		Value:       obj.GetValue(),
	}
	return model, nil
}

// ConvertImageToProto converts Gorm model `Images` to its protobuf type object
func ConvertImageToProto(m *Images) (*storage.Image, error) {
	var msg storage.Image
	if err := msg.Unmarshal(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}
