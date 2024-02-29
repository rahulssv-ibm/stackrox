package schema

import (
	"reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableReportConfigurationsStmt holds the create statement for table `report_configurations`.
	CreateTableReportConfigurationsStmt = &postgres.CreateStmts{
		GormModel: (*ReportConfigurations)(nil),
		Children: []*postgres.CreateStmts{
			{
				GormModel: (*ReportConfigurationsNotifiers)(nil),
				Children:  []*postgres.CreateStmts{},
			},
		},
	}

	// ReportConfigurationsSchema is the go schema for table `report_configurations`.
	ReportConfigurationsSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ReportConfiguration)(nil)), "report_configurations")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_REPORT_CONFIGURATIONS, "reportconfiguration", (*storage.ReportConfiguration)(nil)))
		schema.SetSearchScope([]v1.SearchCategory{
			v1.SearchCategory_REPORT_SNAPSHOT,
		}...)
		schema.ScopingResource = resources.WorkflowAdministration
		return schema
	}()
)

const (
	// ReportConfigurationsTableName specifies the name of the table in postgres.
	ReportConfigurationsTableName = "report_configurations"
	// ReportConfigurationsNotifiersTableName specifies the name of the table in postgres.
	ReportConfigurationsNotifiersTableName = "report_configurations_notifiers"
)

// ReportConfigurations holds the Gorm model for Postgres table `report_configurations`.
type ReportConfigurations struct {
	ID                        string                                 `gorm:"column:id;type:varchar;primaryKey"`
	Name                      string                                 `gorm:"column:name;type:varchar"`
	Type                      storage.ReportConfiguration_ReportType `gorm:"column:type;type:integer"`
	ScopeID                   string                                 `gorm:"column:scopeid;type:varchar"`
	ResourceScopeCollectionID string                                 `gorm:"column:resourcescope_collectionid;type:varchar"`
	CreatorName               string                                 `gorm:"column:creator_name;type:varchar"`
	Serialized                []byte                                 `gorm:"column:serialized;type:bytea"`
}

// ReportConfigurationsNotifiers holds the Gorm model for Postgres table `report_configurations_notifiers`.
type ReportConfigurationsNotifiers struct {
	ReportConfigurationsID  string               `gorm:"column:report_configurations_id;type:varchar;primaryKey"`
	Idx                     int                  `gorm:"column:idx;type:integer;primaryKey;index:reportconfigurationsnotifiers_idx,type:btree"`
	ID                      string               `gorm:"column:id;type:varchar"`
	ReportConfigurationsRef ReportConfigurations `gorm:"foreignKey:report_configurations_id;references:id;belongsTo;constraint:OnDelete:CASCADE"`
	NotifiersRef            Notifiers            `gorm:"foreignKey:id;references:id;belongsTo;constraint:OnDelete:RESTRICT"`
}
