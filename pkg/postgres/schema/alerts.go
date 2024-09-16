// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"
	"time"

	"github.com/lib/pq"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/postgres/mapping"
)

var (
	// CreateTableAlertsStmt holds the create statement for table `alerts`.
	CreateTableAlertsStmt = &postgres.CreateStmts{
		GormModel: (*Alerts)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// AlertsSchema is the go schema for table `alerts`.
	AlertsSchema = func() *walker.Schema {
		schema := GetSchemaForTable("alerts")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.Alert)(nil)), "alerts")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_ALERTS, "alert", (*storage.Alert)(nil)))
		schema.ScopingResource = resources.Alert
		RegisterTable(schema, CreateTableAlertsStmt)
		mapping.RegisterCategoryToTable(v1.SearchCategory_ALERTS, schema)
		return schema
	}()
)

const (
	// AlertsTableName specifies the name of the table in postgres.
	AlertsTableName = "alerts"
)

// Alerts holds the Gorm model for Postgres table `alerts`.
type Alerts struct {
	ID                       string                              `gorm:"column:id;type:uuid;primaryKey"`
	PolicyID                 string                              `gorm:"column:policy_id;type:varchar;index:alerts_policy_id,type:btree"`
	PolicyName               string                              `gorm:"column:policy_name;type:varchar"`
	PolicyDescription        string                              `gorm:"column:policy_description;type:varchar"`
	PolicyDisabled           bool                                `gorm:"column:policy_disabled;type:bool"`
	PolicyCategories         *pq.StringArray                     `gorm:"column:policy_categories;type:text[]"`
	PolicySeverity           storage.Severity                    `gorm:"column:policy_severity;type:integer"`
	PolicyEnforcementActions *pq.Int32Array                      `gorm:"column:policy_enforcementactions;type:int[]"`
	PolicyLastUpdated        *time.Time                          `gorm:"column:policy_lastupdated;type:timestamp"`
	PolicySORTName           string                              `gorm:"column:policy_sortname;type:varchar"`
	PolicySORTLifecycleStage string                              `gorm:"column:policy_sortlifecyclestage;type:varchar"`
	PolicySORTEnforcement    bool                                `gorm:"column:policy_sortenforcement;type:bool"`
	LifecycleStage           storage.LifecycleStage              `gorm:"column:lifecyclestage;type:integer;index:alerts_lifecyclestage,type:btree"`
	ClusterID                string                              `gorm:"column:clusterid;type:uuid;index:alerts_sac_filter,type:btree"`
	ClusterName              string                              `gorm:"column:clustername;type:varchar"`
	Namespace                string                              `gorm:"column:namespace;type:varchar;index:alerts_sac_filter,type:btree"`
	NamespaceID              string                              `gorm:"column:namespaceid;type:uuid"`
	DeploymentID             string                              `gorm:"column:deployment_id;type:uuid;index:alerts_deployment_id,type:hash"`
	DeploymentName           string                              `gorm:"column:deployment_name;type:varchar"`
	DeploymentInactive       bool                                `gorm:"column:deployment_inactive;type:bool"`
	ImageID                  string                              `gorm:"column:image_id;type:varchar"`
	ImageNameRegistry        string                              `gorm:"column:image_name_registry;type:varchar"`
	ImageNameRemote          string                              `gorm:"column:image_name_remote;type:varchar"`
	ImageNameTag             string                              `gorm:"column:image_name_tag;type:varchar"`
	ImageNameFullName        string                              `gorm:"column:image_name_fullname;type:varchar"`
	ResourceResourceType     storage.Alert_Resource_ResourceType `gorm:"column:resource_resourcetype;type:integer"`
	ResourceName             string                              `gorm:"column:resource_name;type:varchar"`
	EnforcementAction        storage.EnforcementAction           `gorm:"column:enforcement_action;type:integer"`
	Time                     *time.Time                          `gorm:"column:time;type:timestamp"`
	State                    storage.ViolationState              `gorm:"column:state;type:integer;index:alerts_state,type:btree"`
	Serialized               []byte                              `gorm:"column:serialized;type:bytea"`
}
