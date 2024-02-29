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
	// CreateTableNodeCvesStmt holds the create statement for table `node_cves`.
	CreateTableNodeCvesStmt = &postgres.CreateStmts{
		GormModel: (*NodeCves)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// NodeCvesSchema is the go schema for table `node_cves`.
	NodeCvesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.NodeCVE)(nil)), "node_cves")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_NODE_VULNERABILITIES, "nodecve", (*storage.NodeCVE)(nil)))
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
	NodeCvesTableName = "node_cves"
)

// NodeCves holds the Gorm model for Postgres table `node_cves`.
type NodeCves struct {
	Id                     string                        `gorm:"column:id;type:varchar;primaryKey"`
	CveBaseInfoCve         string                        `gorm:"column:cvebaseinfo_cve;type:varchar"`
	CveBaseInfoPublishedOn *time.Time                    `gorm:"column:cvebaseinfo_publishedon;type:timestamp"`
	CveBaseInfoCreatedAt   *time.Time                    `gorm:"column:cvebaseinfo_createdat;type:timestamp"`
	OperatingSystem        string                        `gorm:"column:operatingsystem;type:varchar"`
	Cvss                   float32                       `gorm:"column:cvss;type:numeric"`
	Severity               storage.VulnerabilitySeverity `gorm:"column:severity;type:integer"`
	ImpactScore            float32                       `gorm:"column:impactscore;type:numeric"`
	Snoozed                bool                          `gorm:"column:snoozed;type:bool"`
	SnoozeExpiry           *time.Time                    `gorm:"column:snoozeexpiry;type:timestamp"`
	Serialized             []byte                        `gorm:"column:serialized;type:bytea"`
}
