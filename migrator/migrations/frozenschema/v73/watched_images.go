package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableWatchedImagesStmt holds the create statement for table `watched_images`.
	CreateTableWatchedImagesStmt = &postgres.CreateStmts{
		GormModel: (*WatchedImages)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// WatchedImagesSchema is the go schema for table `watched_images`.
	WatchedImagesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.WatchedImage)(nil)), "watched_images")
		return schema
	}()
)

const (
	WatchedImagesTableName = "watched_images"
)

// WatchedImages holds the Gorm model for Postgres table `watched_images`.
type WatchedImages struct {
	Name       string `gorm:"column:name;type:varchar;primaryKey"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
