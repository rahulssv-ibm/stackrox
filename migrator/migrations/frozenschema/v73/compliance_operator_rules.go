package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableComplianceOperatorRulesStmt holds the create statement for table `compliance_operator_rules`.
	CreateTableComplianceOperatorRulesStmt = &postgres.CreateStmts{
		GormModel: (*ComplianceOperatorRules)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ComplianceOperatorRulesSchema is the go schema for table `compliance_operator_rules`.
	ComplianceOperatorRulesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ComplianceOperatorRule)(nil)), "compliance_operator_rules")
		return schema
	}()
)

const (
	ComplianceOperatorRulesTableName = "compliance_operator_rules"
)

// ComplianceOperatorRules holds the Gorm model for Postgres table `compliance_operator_rules`.
type ComplianceOperatorRules struct {
	Id         string `gorm:"column:id;type:varchar;primaryKey"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
