package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableLogImbuesStmt holds the create statement for table `log_imbues`.
	CreateTableLogImbuesStmt = &postgres.CreateStmts{
		GormModel: (*LogImbues)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// LogImbuesSchema is the go schema for table `log_imbues`.
	LogImbuesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.LogImbue)(nil)), "log_imbues")
		return schema
	}()
)

const (
	LogImbuesTableName = "log_imbues"
)

// LogImbues holds the Gorm model for Postgres table `log_imbues`.
type LogImbues struct {
	Id         string `gorm:"column:id;type:varchar;primaryKey"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
