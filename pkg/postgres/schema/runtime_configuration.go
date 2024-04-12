// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
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
	// CreateTableRuntimeConfigurationStmt holds the create statement for table `runtime_configuration`.
	CreateTableRuntimeConfigurationStmt = &postgres.CreateStmts{
		GormModel: (*RuntimeConfiguration)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// RuntimeConfigurationSchema is the go schema for table `runtime_configuration`.
	RuntimeConfigurationSchema = func() *walker.Schema {
		schema := GetSchemaForTable("runtime_configuration")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.RuntimeFilterData)(nil)), "runtime_configuration")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_RUNTIME_CONFIGURATION, "runtimefilterdata", (*storage.RuntimeFilterData)(nil)))
		schema.ScopingResource = resources.Administration
		RegisterTable(schema, CreateTableRuntimeConfigurationStmt)
		mapping.RegisterCategoryToTable(v1.SearchCategory_RUNTIME_CONFIGURATION, schema)
		return schema
	}()
)

const (
	// RuntimeConfigurationTableName specifies the name of the table in postgres.
	RuntimeConfigurationTableName = "runtime_configuration"
)

// RuntimeConfiguration holds the Gorm model for Postgres table `runtime_configuration`.
type RuntimeConfiguration struct {
	ID         string `gorm:"column:id;type:varchar;primaryKey"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
