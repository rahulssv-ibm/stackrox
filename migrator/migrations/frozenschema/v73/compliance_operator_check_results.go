package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableComplianceOperatorCheckResultsStmt holds the create statement for table `compliance_operator_check_results`.
	CreateTableComplianceOperatorCheckResultsStmt = &postgres.CreateStmts{
		GormModel: (*ComplianceOperatorCheckResults)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ComplianceOperatorCheckResultsSchema is the go schema for table `compliance_operator_check_results`.
	ComplianceOperatorCheckResultsSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ComplianceOperatorCheckResult)(nil)), "compliance_operator_check_results")
		return schema
	}()
)

const (
	ComplianceOperatorCheckResultsTableName = "compliance_operator_check_results"
)

// ComplianceOperatorCheckResults holds the Gorm model for Postgres table `compliance_operator_check_results`.
type ComplianceOperatorCheckResults struct {
	Id         string `gorm:"column:id;type:varchar;primaryKey"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
