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
	// CreateTableImageCvesStmt holds the create statement for table `image_cves`.
	CreateTableImageCvesStmt = &postgres.CreateStmts{
		GormModel: (*ImageCves)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ImageCvesSchema is the go schema for table `image_cves`.
	ImageCvesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ImageCVE)(nil)), "image_cves")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_IMAGE_VULNERABILITIES, "imagecve", (*storage.ImageCVE)(nil)))
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
	ImageCvesTableName = "image_cves"
)

// ImageCves holds the Gorm model for Postgres table `image_cves`.
type ImageCves struct {
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
