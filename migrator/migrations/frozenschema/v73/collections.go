package schema

import (
	"reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableCollectionsStmt holds the create statement for table `collections`.
	CreateTableCollectionsStmt = &postgres.CreateStmts{
		GormModel: (*Collections)(nil),
		Children: []*postgres.CreateStmts{
			{
				GormModel: (*CollectionsEmbeddedCollections)(nil),
				Children:  []*postgres.CreateStmts{},
			},
		},
	}

	// CollectionsSchema is the go schema for table `collections`.
	CollectionsSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ResourceCollection)(nil)), "collections")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_COLLECTIONS, "resourcecollection", (*storage.ResourceCollection)(nil)))
		return schema
	}()
)

const (
	CollectionsTableName                    = "collections"
	CollectionsEmbeddedCollectionsTableName = "collections_embedded_collections"
)

// Collections holds the Gorm model for Postgres table `collections`.
type Collections struct {
	Id            string `gorm:"column:id;type:varchar;primaryKey"`
	Name          string `gorm:"column:name;type:varchar;unique"`
	CreatedByName string `gorm:"column:createdby_name;type:varchar"`
	UpdatedByName string `gorm:"column:updatedby_name;type:varchar"`
	Serialized    []byte `gorm:"column:serialized;type:bytea"`
}

// CollectionsEmbeddedCollections holds the Gorm model for Postgres table `collections_embedded_collections`.
type CollectionsEmbeddedCollections struct {
	CollectionsId       string      `gorm:"column:collections_id;type:varchar;primaryKey"`
	Idx                 int         `gorm:"column:idx;type:integer;primaryKey;index:collectionsembeddedcollections_idx,type:btree"`
	Id                  string      `gorm:"column:id;type:varchar"`
	CollectionsRef      Collections `gorm:"foreignKey:collections_id;references:id;belongsTo;constraint:OnDelete:CASCADE"`
	CollectionsCycleRef Collections `gorm:"foreignKey:id;references:id;belongsTo;constraint:OnDelete:RESTRICT"`
}
