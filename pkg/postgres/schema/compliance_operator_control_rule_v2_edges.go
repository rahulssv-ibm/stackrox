// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"fmt"
	"reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/postgres/mapping"
)

var (
	// CreateTableComplianceOperatorControlRuleV2EdgesStmt holds the create statement for table `compliance_operator_control_rule_v2_edges`.
	CreateTableComplianceOperatorControlRuleV2EdgesStmt = &postgres.CreateStmts{
		GormModel: (*ComplianceOperatorControlRuleV2Edges)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ComplianceOperatorControlRuleV2EdgesSchema is the go schema for table `compliance_operator_control_rule_v2_edges`.
	ComplianceOperatorControlRuleV2EdgesSchema = func() *walker.Schema {
		schema := GetSchemaForTable("compliance_operator_control_rule_v2_edges")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.ComplianceOperatorControlRuleV2Edge)(nil)), "compliance_operator_control_rule_v2_edges")
		referencedSchemas := map[string]*walker.Schema{
			"storage.ComplianceOperatorControlV2": ComplianceOperatorControlV2Schema,
			"storage.ComplianceOperatorRuleV2":    ComplianceOperatorRuleV2Schema,
		}

		schema.ResolveReferences(func(messageTypeName string) *walker.Schema {
			return referencedSchemas[fmt.Sprintf("storage.%s", messageTypeName)]
		})
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_COMPLIANCE_CONTROL_RULE_EDGE, "complianceoperatorcontrolrulev2edge", (*storage.ComplianceOperatorControlRuleV2Edge)(nil)))
		schema.ScopingResource = resources.Compliance
		RegisterTable(schema, CreateTableComplianceOperatorControlRuleV2EdgesStmt)
		mapping.RegisterCategoryToTable(v1.SearchCategory_COMPLIANCE_CONTROL_RULE_EDGE, schema)
		return schema
	}()
)

const (
	// ComplianceOperatorControlRuleV2EdgesTableName specifies the name of the table in postgres.
	ComplianceOperatorControlRuleV2EdgesTableName = "compliance_operator_control_rule_v2_edges"
)

// ComplianceOperatorControlRuleV2Edges holds the Gorm model for Postgres table `compliance_operator_control_rule_v2_edges`.
type ComplianceOperatorControlRuleV2Edges struct {
	ID                             string                      `gorm:"column:id;type:uuid;primaryKey"`
	ControlID                      string                      `gorm:"column:controlid;type:varchar"`
	RuleID                         string                      `gorm:"column:ruleid;type:varchar"`
	Serialized                     []byte                      `gorm:"column:serialized;type:bytea"`
	ComplianceOperatorControlV2Ref ComplianceOperatorControlV2 `gorm:"foreignKey:controlid;references:id;belongsTo;constraint:OnDelete:CASCADE"`
	ComplianceOperatorRuleV2Ref    ComplianceOperatorRuleV2    `gorm:"foreignKey:ruleid;references:id;belongsTo;constraint:OnDelete:CASCADE"`
}
