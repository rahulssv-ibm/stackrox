package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableConfigsStmt holds the create statement for table `configs`.
	CreateTableConfigsStmt = &postgres.CreateStmts{
		GormModel: (*Configs)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ConfigsSchema is the go schema for table `configs`.
	ConfigsSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.Config)(nil)), "configs")
		return schema
	}()
)

const (
	ConfigsTableName = "configs"
)

// Configs holds the Gorm model for Postgres table `configs`.
type Configs struct {
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
