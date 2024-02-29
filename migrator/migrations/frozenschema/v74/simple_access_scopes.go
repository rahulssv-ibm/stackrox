package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableSimpleAccessScopesStmt holds the create statement for table `simple_access_scopes`.
	CreateTableSimpleAccessScopesStmt = &postgres.CreateStmts{
		GormModel: (*SimpleAccessScopes)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// SimpleAccessScopesSchema is the go schema for table `simple_access_scopes`.
	SimpleAccessScopesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.SimpleAccessScope)(nil)), "simple_access_scopes")
		return schema
	}()
)

const (
	SimpleAccessScopesTableName = "simple_access_scopes"
)

// SimpleAccessScopes holds the Gorm model for Postgres table `simple_access_scopes`.
type SimpleAccessScopes struct {
	Id         string `gorm:"column:id;type:uuid;primaryKey"`
	Name       string `gorm:"column:name;type:varchar;unique"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
