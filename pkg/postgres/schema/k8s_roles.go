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
	// CreateTableK8sRolesStmt holds the create statement for table `k8s_roles`.
	CreateTableK8sRolesStmt = &postgres.CreateStmts{
		GormModel: (*K8sRoles)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// K8sRolesSchema is the go schema for table `k8s_roles`.
	K8sRolesSchema = func() *walker.Schema {
		schema := GetSchemaForTable("k8s_roles")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.K8SRole)(nil)), "k8s_roles")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_ROLES, "k8srole", (*storage.K8SRole)(nil)))
		schema.ScopingResource = resources.K8sRole
		RegisterTable(schema, CreateTableK8sRolesStmt)
		mapping.RegisterCategoryToTable(v1.SearchCategory_ROLES, schema)
		return schema
	}()
)

const (
	// K8sRolesTableName specifies the name of the table in postgres.
	K8sRolesTableName = "k8s_roles"
)

// K8sRoles holds the Gorm model for Postgres table `k8s_roles`.
type K8sRoles struct {
	ID          string            `gorm:"column:id;type:uuid;primaryKey"`
	Name        string            `gorm:"column:name;type:varchar"`
	Namespace   string            `gorm:"column:namespace;type:varchar;index:k8sroles_sac_filter,type:btree"`
	ClusterID   string            `gorm:"column:clusterid;type:uuid;index:k8sroles_sac_filter,type:btree"`
	ClusterName string            `gorm:"column:clustername;type:varchar"`
	ClusterRole bool              `gorm:"column:clusterrole;type:bool"`
	Labels      map[string]string `gorm:"column:labels;type:jsonb"`
	Annotations map[string]string `gorm:"column:annotations;type:jsonb"`
	Serialized  []byte            `gorm:"column:serialized;type:bytea"`
}
