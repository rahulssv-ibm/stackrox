package schema

import (
	"github.com/stackrox/rox/generated/storage"
)

// ConvertReportConfigurationFromProto converts a `*storage.ReportConfiguration` to Gorm model
func ConvertReportConfigurationFromProto(obj *storage.ReportConfiguration) (*ReportConfigurations, error) {
	serialized, err := obj.Marshal()
	if err != nil {
		return nil, err
	}
	model := &ReportConfigurations{
		ID:                        obj.GetId(),
		Name:                      obj.GetName(),
		Type:                      obj.GetType(),
		ScopeID:                   obj.GetScopeId(),
		ResourceScopeCollectionID: obj.GetResourceScope().GetCollectionId(),
		CreatorName:               obj.GetCreator().GetName(),
		Serialized:                serialized,
	}
	return model, nil
}

// ConvertNotifierConfigurationFromProto converts a `*storage.NotifierConfiguration` to Gorm model
func ConvertNotifierConfigurationFromProto(obj *storage.NotifierConfiguration, idx int, reportConfigurationID string) (*ReportConfigurationsNotifiers, error) {
	model := &ReportConfigurationsNotifiers{
		ReportConfigurationsID: reportConfigurationID,
		Idx:                    idx,
		ID:                     obj.GetId(),
	}
	return model, nil
}

// ConvertReportConfigurationToProto converts Gorm model `ReportConfigurations` to its protobuf type object
func ConvertReportConfigurationToProto(m *ReportConfigurations) (*storage.ReportConfiguration, error) {
	var msg storage.ReportConfiguration
	if err := msg.Unmarshal(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}
