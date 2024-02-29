package schema

import (
	"reflect"
	"time"

	"github.com/lib/pq"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableImagesStmt holds the create statement for table `images`.
	CreateTableImagesStmt = &postgres.CreateStmts{
		GormModel: (*Images)(nil),
		Children: []*postgres.CreateStmts{
			{
				GormModel: (*ImagesLayers)(nil),
				Children:  []*postgres.CreateStmts{},
			},
		},
	}

	// ImagesSchema is the go schema for table `images`.
	ImagesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.Image)(nil)), "images")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_IMAGES, "image", (*storage.Image)(nil)))
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
	ImagesTableName       = "images"
	ImagesLayersTableName = "images_layers"
)

// Images holds the Gorm model for Postgres table `images`.
type Images struct {
	Id                   string            `gorm:"column:id;type:varchar;primaryKey"`
	NameRegistry         string            `gorm:"column:name_registry;type:varchar"`
	NameRemote           string            `gorm:"column:name_remote;type:varchar"`
	NameTag              string            `gorm:"column:name_tag;type:varchar"`
	NameFullName         string            `gorm:"column:name_fullname;type:varchar"`
	MetadataV1Created    *time.Time        `gorm:"column:metadata_v1_created;type:timestamp"`
	MetadataV1User       string            `gorm:"column:metadata_v1_user;type:varchar"`
	MetadataV1Command    *pq.StringArray   `gorm:"column:metadata_v1_command;type:text[]"`
	MetadataV1Entrypoint *pq.StringArray   `gorm:"column:metadata_v1_entrypoint;type:text[]"`
	MetadataV1Volumes    *pq.StringArray   `gorm:"column:metadata_v1_volumes;type:text[]"`
	MetadataV1Labels     map[string]string `gorm:"column:metadata_v1_labels;type:jsonb"`
	ScanScanTime         *time.Time        `gorm:"column:scan_scantime;type:timestamp"`
	ScanOperatingSystem  string            `gorm:"column:scan_operatingsystem;type:varchar"`
	SignatureFetched     *time.Time        `gorm:"column:signature_fetched;type:timestamp"`
	Components           int32             `gorm:"column:components;type:integer"`
	Cves                 int32             `gorm:"column:cves;type:integer"`
	FixableCves          int32             `gorm:"column:fixablecves;type:integer"`
	LastUpdated          *time.Time        `gorm:"column:lastupdated;type:timestamp"`
	Priority             int64             `gorm:"column:priority;type:bigint"`
	RiskScore            float32           `gorm:"column:riskscore;type:numeric"`
	TopCvss              float32           `gorm:"column:topcvss;type:numeric"`
	Serialized           []byte            `gorm:"column:serialized;type:bytea"`
}

// ImagesLayers holds the Gorm model for Postgres table `images_layers`.
type ImagesLayers struct {
	ImagesId    string `gorm:"column:images_id;type:varchar;primaryKey"`
	Idx         int    `gorm:"column:idx;type:integer;primaryKey;index:imageslayers_idx,type:btree"`
	Instruction string `gorm:"column:instruction;type:varchar"`
	Value       string `gorm:"column:value;type:varchar"`
	ImagesRef   Images `gorm:"foreignKey:images_id;references:id;belongsTo;constraint:OnDelete:CASCADE"`
}
