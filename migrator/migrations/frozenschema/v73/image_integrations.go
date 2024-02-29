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
	// CreateTableImageIntegrationsStmt holds the create statement for table `image_integrations`.
	CreateTableImageIntegrationsStmt = &postgres.CreateStmts{
		GormModel: (*ImageIntegrations)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ImageIntegrationsSchema is the go schema for table `image_integrations`.
	ImageIntegrationsSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ImageIntegration)(nil)), "image_integrations")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_IMAGE_INTEGRATIONS, "imageintegration", (*storage.ImageIntegration)(nil)))
		return schema
	}()
)

const (
	ImageIntegrationsTableName = "image_integrations"
)

// ImageIntegrations holds the Gorm model for Postgres table `image_integrations`.
type ImageIntegrations struct {
	Id         string `gorm:"column:id;type:uuid;primaryKey"`
	Name       string `gorm:"column:name;type:varchar;unique"`
	ClusterId  string `gorm:"column:clusterid;type:uuid;index:imageintegrations_sac_filter,type:btree"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
