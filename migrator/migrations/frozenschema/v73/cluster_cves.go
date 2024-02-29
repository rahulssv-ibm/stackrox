package schema

import (
	"reflect"
	"time"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableClusterCvesStmt holds the create statement for table `cluster_cves`.
	CreateTableClusterCvesStmt = &postgres.CreateStmts{
		GormModel: (*ClusterCves)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ClusterCvesSchema is the go schema for table `cluster_cves`.
	ClusterCvesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ClusterCVE)(nil)), "cluster_cves")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_CLUSTER_VULNERABILITIES, "clustercve", (*storage.ClusterCVE)(nil)))
		schema.SetSearchScope([]v1.SearchCategory{
			v1.SearchCategory_CLUSTER_VULNERABILITIES,
			v1.SearchCategory_CLUSTER_VULN_EDGE,
			v1.SearchCategory_CLUSTERS,
		}...)
		return schema
	}()
)

const (
	ClusterCvesTableName = "cluster_cves"
)

// ClusterCves holds the Gorm model for Postgres table `cluster_cves`.
type ClusterCves struct {
	Id                     string                        `gorm:"column:id;type:varchar;primaryKey"`
	CveBaseInfoCve         string                        `gorm:"column:cvebaseinfo_cve;type:varchar"`
	CveBaseInfoPublishedOn *time.Time                    `gorm:"column:cvebaseinfo_publishedon;type:timestamp"`
	CveBaseInfoCreatedAt   *time.Time                    `gorm:"column:cvebaseinfo_createdat;type:timestamp"`
	Cvss                   float32                       `gorm:"column:cvss;type:numeric"`
	Severity               storage.VulnerabilitySeverity `gorm:"column:severity;type:integer"`
	ImpactScore            float32                       `gorm:"column:impactscore;type:numeric"`
	Snoozed                bool                          `gorm:"column:snoozed;type:bool"`
	SnoozeExpiry           *time.Time                    `gorm:"column:snoozeexpiry;type:timestamp"`
	Type                   storage.CVE_CVEType           `gorm:"column:type;type:integer"`
	Serialized             []byte                        `gorm:"column:serialized;type:bytea"`
}
