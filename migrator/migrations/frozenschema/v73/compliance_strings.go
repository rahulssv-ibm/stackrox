package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableComplianceStringsStmt holds the create statement for table `compliance_strings`.
	CreateTableComplianceStringsStmt = &postgres.CreateStmts{
		GormModel: (*ComplianceStrings)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ComplianceStringsSchema is the go schema for table `compliance_strings`.
	ComplianceStringsSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ComplianceStrings)(nil)), "compliance_strings")
		return schema
	}()
)

const (
	ComplianceStringsTableName = "compliance_strings"
)

// ComplianceStrings holds the Gorm model for Postgres table `compliance_strings`.
type ComplianceStrings struct {
	Id         string `gorm:"column:id;type:varchar;primaryKey"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
