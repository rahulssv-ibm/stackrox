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
	// CreateTableNodeComponentsCvesEdgesStmt holds the create statement for table `node_components_cves_edges`.
	CreateTableNodeComponentsCvesEdgesStmt = &postgres.CreateStmts{
		GormModel: (*NodeComponentsCvesEdges)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// NodeComponentsCvesEdgesSchema is the go schema for table `node_components_cves_edges`.
	NodeComponentsCvesEdgesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.NodeComponentCVEEdge)(nil)), "node_components_cves_edges")
		referencedSchemas := map[string]*walker.Schema{
			"storage.NodeComponent": NodeComponentsSchema,
			"storage.NodeCVE":       NodeCvesSchema,
		}

		schema.ResolveReferences(func(messageTypeName string) *walker.Schema {
			return referencedSchemas[fmt.Sprintf("storage.%s", messageTypeName)]
		})
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_NODE_COMPONENT_CVE_EDGE, "nodecomponentcveedge", (*storage.NodeComponentCVEEdge)(nil)))
		schema.SetSearchScope([]v1.SearchCategory{
			v1.SearchCategory_NODE_VULNERABILITIES,
			v1.SearchCategory_NODE_COMPONENT_CVE_EDGE,
			v1.SearchCategory_NODE_COMPONENTS,
			v1.SearchCategory_NODE_COMPONENT_EDGE,
			v1.SearchCategory_NODES,
			v1.SearchCategory_CLUSTERS,
		}...)
		return schema
	}()
)

const (
	NodeComponentsCvesEdgesTableName = "node_components_cves_edges"
)

// NodeComponentsCvesEdges holds the Gorm model for Postgres table `node_components_cves_edges`.
type NodeComponentsCvesEdges struct {
	Id                string         `gorm:"column:id;type:varchar;primaryKey"`
	IsFixable         bool           `gorm:"column:isfixable;type:bool"`
	FixedBy           string         `gorm:"column:fixedby;type:varchar"`
	NodeComponentId   string         `gorm:"column:nodecomponentid;type:varchar;index:nodecomponentscvesedges_nodecomponentid,type:hash"`
	NodeCveId         string         `gorm:"column:nodecveid;type:varchar;index:nodecomponentscvesedges_nodecveid,type:hash"`
	Serialized        []byte         `gorm:"column:serialized;type:bytea"`
	NodeComponentsRef NodeComponents `gorm:"foreignKey:nodecomponentid;references:id;belongsTo;constraint:OnDelete:CASCADE"`
}
