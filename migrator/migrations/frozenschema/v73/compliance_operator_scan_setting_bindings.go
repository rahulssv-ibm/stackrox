package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableComplianceOperatorScanSettingBindingsStmt holds the create statement for table `compliance_operator_scan_setting_bindings`.
	CreateTableComplianceOperatorScanSettingBindingsStmt = &postgres.CreateStmts{
		GormModel: (*ComplianceOperatorScanSettingBindings)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ComplianceOperatorScanSettingBindingsSchema is the go schema for table `compliance_operator_scan_setting_bindings`.
	ComplianceOperatorScanSettingBindingsSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ComplianceOperatorScanSettingBinding)(nil)), "compliance_operator_scan_setting_bindings")
		return schema
	}()
)

const (
	ComplianceOperatorScanSettingBindingsTableName = "compliance_operator_scan_setting_bindings"
)

// ComplianceOperatorScanSettingBindings holds the Gorm model for Postgres table `compliance_operator_scan_setting_bindings`.
type ComplianceOperatorScanSettingBindings struct {
	Id         string `gorm:"column:id;type:varchar;primaryKey"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
