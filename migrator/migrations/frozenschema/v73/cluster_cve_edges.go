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
	// CreateTableClusterCveEdgesStmt holds the create statement for table `cluster_cve_edges`.
	CreateTableClusterCveEdgesStmt = &postgres.CreateStmts{
		GormModel: (*ClusterCveEdges)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ClusterCveEdgesSchema is the go schema for table `cluster_cve_edges`.
	ClusterCveEdgesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ClusterCVEEdge)(nil)), "cluster_cve_edges")
		referencedSchemas := map[string]*walker.Schema{
			"storage.Cluster":    ClustersSchema,
			"storage.ClusterCVE": ClusterCvesSchema,
		}

		schema.ResolveReferences(func(messageTypeName string) *walker.Schema {
			return referencedSchemas[fmt.Sprintf("storage.%s", messageTypeName)]
		})
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_CLUSTER_VULN_EDGE, "clustercveedge", (*storage.ClusterCVEEdge)(nil)))
		schema.SetSearchScope([]v1.SearchCategory{
			v1.SearchCategory_CLUSTER_VULNERABILITIES,
			v1.SearchCategory_CLUSTER_VULN_EDGE,
			v1.SearchCategory_CLUSTERS,
		}...)
		return schema
	}()
)

const (
	ClusterCveEdgesTableName = "cluster_cve_edges"
)

// ClusterCveEdges holds the Gorm model for Postgres table `cluster_cve_edges`.
type ClusterCveEdges struct {
	Id          string   `gorm:"column:id;type:varchar;primaryKey"`
	IsFixable   bool     `gorm:"column:isfixable;type:bool"`
	FixedBy     string   `gorm:"column:fixedby;type:varchar"`
	ClusterId   string   `gorm:"column:clusterid;type:uuid"`
	CveId       string   `gorm:"column:cveid;type:varchar;index:clustercveedges_cveid,type:hash"`
	Serialized  []byte   `gorm:"column:serialized;type:bytea"`
	ClustersRef Clusters `gorm:"foreignKey:clusterid;references:id;belongsTo;constraint:OnDelete:CASCADE"`
}
