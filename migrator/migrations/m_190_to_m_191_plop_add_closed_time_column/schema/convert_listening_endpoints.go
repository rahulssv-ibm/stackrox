package schema

import (
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
)

// ConvertProcessListeningOnPortStorageFromProto converts a `*storage.ProcessListeningOnPortStorage` to Gorm model
func ConvertProcessListeningOnPortStorageFromProto(obj *storage.ProcessListeningOnPortStorage) (*ListeningEndpoints, error) {
	serialized, err := obj.Marshal()
	if err != nil {
		return nil, err
	}
	model := &ListeningEndpoints{
		ID:                 obj.GetId(),
		Port:               obj.GetPort(),
		Protocol:           obj.GetProtocol(),
		CloseTimestamp:     pgutils.NilOrTime(obj.GetCloseTimestamp()),
		ProcessIndicatorID: obj.GetProcessIndicatorId(),
		Closed:             obj.GetClosed(),
		DeploymentID:       obj.GetDeploymentId(),
		Serialized:         serialized,
	}
	return model, nil
}

// ConvertProcessListeningOnPortStorageToProto converts Gorm model `ListeningEndpoints` to its protobuf type object
func ConvertProcessListeningOnPortStorageToProto(m *ListeningEndpoints) (*storage.ProcessListeningOnPortStorage, error) {
	var msg storage.ProcessListeningOnPortStorage
	if err := msg.Unmarshal(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}
