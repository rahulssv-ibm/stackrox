package schema

import (
	"reflect"
	"time"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableReportSnapshotsStmt holds the create statement for table `report_snapshots`.
	CreateTableReportSnapshotsStmt = &postgres.CreateStmts{
		GormModel: (*ReportSnapshots)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ReportSnapshotsSchema is the go schema for table `report_snapshots`.
	ReportSnapshotsSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ReportSnapshot)(nil)), "report_snapshots")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_REPORT_SNAPSHOT, "reportsnapshot", (*storage.ReportSnapshot)(nil)))
		schema.ScopingResource = resources.WorkflowAdministration
		return schema
	}()
)

const (
	// ReportSnapshotsTableName specifies the name of the table in postgres.
	ReportSnapshotsTableName = "report_snapshots"
)

// ReportSnapshots holds the Gorm model for Postgres table `report_snapshots`.
type ReportSnapshots struct {
	ReportID                             string                                  `gorm:"column:reportid;type:uuid;primaryKey"`
	ReportConfigurationID                string                                  `gorm:"column:reportconfigurationid;type:varchar"`
	Name                                 string                                  `gorm:"column:name;type:varchar"`
	ReportStatusRunState                 storage.ReportStatus_RunState           `gorm:"column:reportstatus_runstate;type:integer"`
	ReportStatusQueuedAt                 *time.Time                              `gorm:"column:reportstatus_queuedat;type:timestamp"`
	ReportStatusCompletedAt              *time.Time                              `gorm:"column:reportstatus_completedat;type:timestamp"`
	ReportStatusReportRequestType        storage.ReportStatus_RunMethod          `gorm:"column:reportstatus_reportrequesttype;type:integer"`
	ReportStatusReportNotificationMethod storage.ReportStatus_NotificationMethod `gorm:"column:reportstatus_reportnotificationmethod;type:integer"`
	RequesterID                          string                                  `gorm:"column:requester_id;type:varchar"`
	RequesterName                        string                                  `gorm:"column:requester_name;type:varchar"`
	Serialized                           []byte                                  `gorm:"column:serialized;type:bytea"`
	ReportConfigurationsRef              ReportConfigurations                    `gorm:"foreignKey:reportconfigurationid;references:id;belongsTo;constraint:OnDelete:CASCADE"`
}
