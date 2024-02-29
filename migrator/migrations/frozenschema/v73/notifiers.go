package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableNotifiersStmt holds the create statement for table `notifiers`.
	CreateTableNotifiersStmt = &postgres.CreateStmts{
		GormModel: (*Notifiers)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// NotifiersSchema is the go schema for table `notifiers`.
	NotifiersSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.Notifier)(nil)), "notifiers")
		return schema
	}()
)

const (
	NotifiersTableName = "notifiers"
)

// Notifiers holds the Gorm model for Postgres table `notifiers`.
type Notifiers struct {
	Id         string `gorm:"column:id;type:varchar;primaryKey"`
	Name       string `gorm:"column:name;type:varchar;unique"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
