// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"

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
	// CreateTableComplianceOperatorBenchmarkV2Stmt holds the create statement for table `compliance_operator_benchmark_v2`.
	CreateTableComplianceOperatorBenchmarkV2Stmt = &postgres.CreateStmts{
		GormModel: (*ComplianceOperatorBenchmarkV2)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ComplianceOperatorBenchmarkV2Schema is the go schema for table `compliance_operator_benchmark_v2`.
	ComplianceOperatorBenchmarkV2Schema = func() *walker.Schema {
		schema := GetSchemaForTable("compliance_operator_benchmark_v2")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.ComplianceOperatorBenchmarkV2)(nil)), "compliance_operator_benchmark_v2")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_COMPLIANCE_BENCHMARK, "complianceoperatorbenchmarkv2", (*storage.ComplianceOperatorBenchmarkV2)(nil)))
		schema.ScopingResource = resources.ComplianceOperator
		RegisterTable(schema, CreateTableComplianceOperatorBenchmarkV2Stmt, features.ComplianceEnhancements.Enabled)
		mapping.RegisterCategoryToTable(v1.SearchCategory_COMPLIANCE_BENCHMARK, schema)
		return schema
	}()
)

const (
	// ComplianceOperatorBenchmarkV2TableName specifies the name of the table in postgres.
	ComplianceOperatorBenchmarkV2TableName = "compliance_operator_benchmark_v2"
)

// ComplianceOperatorBenchmarkV2 holds the Gorm model for Postgres table `compliance_operator_benchmark_v2`.
type ComplianceOperatorBenchmarkV2 struct {
	ID         string `gorm:"column:id;type:uuid;primaryKey"`
	Name       string `gorm:"column:name;type:varchar"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
