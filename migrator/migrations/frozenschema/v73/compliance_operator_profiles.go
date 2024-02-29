package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableComplianceOperatorProfilesStmt holds the create statement for table `compliance_operator_profiles`.
	CreateTableComplianceOperatorProfilesStmt = &postgres.CreateStmts{
		GormModel: (*ComplianceOperatorProfiles)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ComplianceOperatorProfilesSchema is the go schema for table `compliance_operator_profiles`.
	ComplianceOperatorProfilesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ComplianceOperatorProfile)(nil)), "compliance_operator_profiles")
		return schema
	}()
)

const (
	ComplianceOperatorProfilesTableName = "compliance_operator_profiles"
)

// ComplianceOperatorProfiles holds the Gorm model for Postgres table `compliance_operator_profiles`.
type ComplianceOperatorProfiles struct {
	Id         string `gorm:"column:id;type:varchar;primaryKey"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
