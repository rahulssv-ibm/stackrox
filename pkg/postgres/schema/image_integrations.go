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
	// CreateTableImageIntegrationsStmt holds the create statement for table `image_integrations`.
	CreateTableImageIntegrationsStmt = &postgres.CreateStmts{
		GormModel: (*ImageIntegrations)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ImageIntegrationsSchema is the go schema for table `image_integrations`.
	ImageIntegrationsSchema = func() *walker.Schema {
		schema := GetSchemaForTable("image_integrations")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.ImageIntegration)(nil)), "image_integrations")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_IMAGE_INTEGRATIONS, "imageintegration", (*storage.ImageIntegration)(nil)))
		schema.ScopingResource = resources.Integration
		RegisterTable(schema, CreateTableImageIntegrationsStmt)
		mapping.RegisterCategoryToTable(v1.SearchCategory_IMAGE_INTEGRATIONS, schema)
		return schema
	}()
)

const (
	// ImageIntegrationsTableName specifies the name of the table in postgres.
	ImageIntegrationsTableName = "image_integrations"
)

// ImageIntegrations holds the Gorm model for Postgres table `image_integrations`.
type ImageIntegrations struct {
	ID         string `gorm:"column:id;type:uuid;primaryKey"`
	Name       string `gorm:"column:name;type:varchar;unique"`
	ClusterID  string `gorm:"column:clusterid;type:uuid;index:imageintegrations_sac_filter,type:btree"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
