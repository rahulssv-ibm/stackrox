package schema

import (
	"fmt"
	"reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableImageComponentEdgesStmt holds the create statement for table `image_component_edges`.
	CreateTableImageComponentEdgesStmt = &postgres.CreateStmts{
		GormModel: (*ImageComponentEdges)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ImageComponentEdgesSchema is the go schema for table `image_component_edges`.
	ImageComponentEdgesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ImageComponentEdge)(nil)), "image_component_edges")
		referencedSchemas := map[string]*walker.Schema{
			"storage.Image":          ImagesSchema,
			"storage.ImageComponent": ImageComponentsSchema,
		}

		schema.ResolveReferences(func(messageTypeName string) *walker.Schema {
			return referencedSchemas[fmt.Sprintf("storage.%s", messageTypeName)]
		})
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_IMAGE_COMPONENT_EDGE, "imagecomponentedge", (*storage.ImageComponentEdge)(nil)))
		schema.SetSearchScope([]v1.SearchCategory{
			v1.SearchCategory_IMAGE_VULNERABILITIES,
			v1.SearchCategory_COMPONENT_VULN_EDGE,
			v1.SearchCategory_IMAGE_COMPONENTS,
			v1.SearchCategory_IMAGE_COMPONENT_EDGE,
			v1.SearchCategory_IMAGE_VULN_EDGE,
			v1.SearchCategory_IMAGES,
			v1.SearchCategory_DEPLOYMENTS,
			v1.SearchCategory_NAMESPACES,
			v1.SearchCategory_CLUSTERS,
		}...)
		return schema
	}()
)

const (
	ImageComponentEdgesTableName = "image_component_edges"
)

// ImageComponentEdges holds the Gorm model for Postgres table `image_component_edges`.
type ImageComponentEdges struct {
	Id               string `gorm:"column:id;type:varchar;primaryKey"`
	Location         string `gorm:"column:location;type:varchar"`
	ImageId          string `gorm:"column:imageid;type:varchar;index:imagecomponentedges_imageid,type:hash"`
	ImageComponentId string `gorm:"column:imagecomponentid;type:varchar;index:imagecomponentedges_imagecomponentid,type:hash"`
	Serialized       []byte `gorm:"column:serialized;type:bytea"`
	ImagesRef        Images `gorm:"foreignKey:imageid;references:id;belongsTo;constraint:OnDelete:CASCADE"`
}
