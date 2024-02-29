package schema

import (
	"reflect"
	"time"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableComplianceRunMetadataStmt holds the create statement for table `compliance_run_metadata`.
	CreateTableComplianceRunMetadataStmt = &postgres.CreateStmts{
		GormModel: (*ComplianceRunMetadata)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ComplianceRunMetadataSchema is the go schema for table `compliance_run_metadata`.
	ComplianceRunMetadataSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ComplianceRunMetadata)(nil)), "compliance_run_metadata")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_COMPLIANCE_METADATA, "compliancerunmetadata", (*storage.ComplianceRunMetadata)(nil)))
		return schema
	}()
)

const (
	ComplianceRunMetadataTableName = "compliance_run_metadata"
)

// ComplianceRunMetadata holds the Gorm model for Postgres table `compliance_run_metadata`.
type ComplianceRunMetadata struct {
	RunId           string     `gorm:"column:runid;type:varchar;primaryKey"`
	StandardId      string     `gorm:"column:standardid;type:varchar"`
	ClusterId       string     `gorm:"column:clusterid;type:uuid;index:compliancerunmetadata_sac_filter,type:hash"`
	FinishTimestamp *time.Time `gorm:"column:finishtimestamp;type:timestamp"`
	Serialized      []byte     `gorm:"column:serialized;type:bytea"`
}
