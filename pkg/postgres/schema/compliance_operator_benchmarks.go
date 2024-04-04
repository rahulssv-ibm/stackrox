// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/sac/resources"
)

var (
	// CreateTableComplianceOperatorBenchmarksStmt holds the create statement for table `compliance_operator_benchmarks`.
	CreateTableComplianceOperatorBenchmarksStmt = &postgres.CreateStmts{
		GormModel: (*ComplianceOperatorBenchmarks)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ComplianceOperatorBenchmarksSchema is the go schema for table `compliance_operator_benchmarks`.
	ComplianceOperatorBenchmarksSchema = func() *walker.Schema {
		schema := GetSchemaForTable("compliance_operator_benchmarks")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.ComplianceOperatorBenchmarkV2)(nil)), "compliance_operator_benchmarks")
		schema.ScopingResource = resources.Compliance
		RegisterTable(schema, CreateTableComplianceOperatorBenchmarksStmt)
		return schema
	}()
)

const (
	// ComplianceOperatorBenchmarksTableName specifies the name of the table in postgres.
	ComplianceOperatorBenchmarksTableName = "compliance_operator_benchmarks"
)

// ComplianceOperatorBenchmarks holds the Gorm model for Postgres table `compliance_operator_benchmarks`.
type ComplianceOperatorBenchmarks struct {
	ID         string `gorm:"column:id;type:uuid;primaryKey"`
	Name       string `gorm:"column:name;type:varchar"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
