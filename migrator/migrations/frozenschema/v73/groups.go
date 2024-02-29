package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableGroupsStmt holds the create statement for table `groups`.
	CreateTableGroupsStmt = &postgres.CreateStmts{
		GormModel: (*Groups)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// GroupsSchema is the go schema for table `groups`.
	GroupsSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.Group)(nil)), "groups")
		return schema
	}()
)

const (
	GroupsTableName = "groups"
)

// Groups holds the Gorm model for Postgres table `groups`.
type Groups struct {
	PropsId    string `gorm:"column:props_id;type:varchar;primaryKey"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
