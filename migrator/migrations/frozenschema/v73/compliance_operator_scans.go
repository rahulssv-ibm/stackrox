package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableComplianceOperatorScansStmt holds the create statement for table `compliance_operator_scans`.
	CreateTableComplianceOperatorScansStmt = &postgres.CreateStmts{
		GormModel: (*ComplianceOperatorScans)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ComplianceOperatorScansSchema is the go schema for table `compliance_operator_scans`.
	ComplianceOperatorScansSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ComplianceOperatorScan)(nil)), "compliance_operator_scans")
		return schema
	}()
)

const (
	ComplianceOperatorScansTableName = "compliance_operator_scans"
)

// ComplianceOperatorScans holds the Gorm model for Postgres table `compliance_operator_scans`.
type ComplianceOperatorScans struct {
	Id         string `gorm:"column:id;type:varchar;primaryKey"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
