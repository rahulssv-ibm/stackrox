// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"fmt"
	"reflect"
	"time"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/postgres/mapping"
)

var (
	// CreateTableComplianceOperatorScanV2Stmt holds the create statement for table `compliance_operator_scan_v2`.
	CreateTableComplianceOperatorScanV2Stmt = &postgres.CreateStmts{
		GormModel: (*ComplianceOperatorScanV2)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ComplianceOperatorScanV2Schema is the go schema for table `compliance_operator_scan_v2`.
	ComplianceOperatorScanV2Schema = func() *walker.Schema {
		schema := GetSchemaForTable("compliance_operator_scan_v2")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.ComplianceOperatorScanV2)(nil)), "compliance_operator_scan_v2")
		referencedSchemas := map[string]*walker.Schema{
			"storage.ComplianceOperatorProfileV2":           ComplianceOperatorProfileV2Schema,
			"storage.ComplianceOperatorScanConfigurationV2": ComplianceOperatorScanConfigurationV2Schema,
		}

		schema.ResolveReferences(func(messageTypeName string) *walker.Schema {
			return referencedSchemas[fmt.Sprintf("storage.%s", messageTypeName)]
		})
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_COMPLIANCE_SCAN, "complianceoperatorscanv2", (*storage.ComplianceOperatorScanV2)(nil)))
		schema.ScopingResource = resources.Compliance
		RegisterTable(schema, CreateTableComplianceOperatorScanV2Stmt, features.ComplianceEnhancements.Enabled)
		mapping.RegisterCategoryToTable(v1.SearchCategory_COMPLIANCE_SCAN, schema)
		return schema
	}()
)

const (
	// ComplianceOperatorScanV2TableName specifies the name of the table in postgres.
	ComplianceOperatorScanV2TableName = "compliance_operator_scan_v2"
)

// ComplianceOperatorScanV2 holds the Gorm model for Postgres table `compliance_operator_scan_v2`.
type ComplianceOperatorScanV2 struct {
	ID                  string     `gorm:"column:id;type:varchar;primaryKey"`
	ScanConfigName      string     `gorm:"column:scanconfigname;type:varchar"`
	ClusterID           string     `gorm:"column:clusterid;type:uuid;index:complianceoperatorscanv2_sac_filter,type:hash"`
	ProfileProfileRefID string     `gorm:"column:profile_profilerefid;type:uuid"`
	StatusResult        string     `gorm:"column:status_result;type:varchar"`
	LastExecutedTime    *time.Time `gorm:"column:lastexecutedtime;type:timestamp"`
	ScanName            string     `gorm:"column:scanname;type:varchar"`
	ScanRefID           string     `gorm:"column:scanrefid;type:uuid"`
	Serialized          []byte     `gorm:"column:serialized;type:bytea"`
}
