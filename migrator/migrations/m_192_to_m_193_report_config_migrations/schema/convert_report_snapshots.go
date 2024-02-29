package schema

import (
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
)

// ConvertReportSnapshotFromProto converts a `*storage.ReportSnapshot` to Gorm model
func ConvertReportSnapshotFromProto(obj *storage.ReportSnapshot) (*ReportSnapshots, error) {
	serialized, err := obj.Marshal()
	if err != nil {
		return nil, err
	}
	model := &ReportSnapshots{
		ReportID:                             obj.GetReportId(),
		ReportConfigurationID:                obj.GetReportConfigurationId(),
		Name:                                 obj.GetName(),
		ReportStatusRunState:                 obj.GetReportStatus().GetRunState(),
		ReportStatusQueuedAt:                 pgutils.NilOrTime(obj.GetReportStatus().GetQueuedAt()),
		ReportStatusCompletedAt:              pgutils.NilOrTime(obj.GetReportStatus().GetCompletedAt()),
		ReportStatusReportRequestType:        obj.GetReportStatus().GetReportRequestType(),
		ReportStatusReportNotificationMethod: obj.GetReportStatus().GetReportNotificationMethod(),
		RequesterID:                          obj.GetRequester().GetId(),
		RequesterName:                        obj.GetRequester().GetName(),
		Serialized:                           serialized,
	}
	return model, nil
}

// ConvertReportSnapshotToProto converts Gorm model `ReportSnapshots` to its protobuf type object
func ConvertReportSnapshotToProto(m *ReportSnapshots) (*storage.ReportSnapshot, error) {
	var msg storage.ReportSnapshot
	if err := msg.Unmarshal(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}
