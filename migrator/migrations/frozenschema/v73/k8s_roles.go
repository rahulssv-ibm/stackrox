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
	// CreateTableK8sRolesStmt holds the create statement for table `k8s_roles`.
	CreateTableK8sRolesStmt = &postgres.CreateStmts{
		GormModel: (*K8sRoles)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// K8sRolesSchema is the go schema for table `k8s_roles`.
	K8sRolesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.K8SRole)(nil)), "k8s_roles")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_ROLES, "k8srole", (*storage.K8SRole)(nil)))
		return schema
	}()
)

const (
	K8sRolesTableName = "k8s_roles"
)

// K8sRoles holds the Gorm model for Postgres table `k8s_roles`.
type K8sRoles struct {
	Id          string            `gorm:"column:id;type:uuid;primaryKey"`
	Name        string            `gorm:"column:name;type:varchar"`
	Namespace   string            `gorm:"column:namespace;type:varchar;index:k8sroles_sac_filter,type:btree"`
	ClusterId   string            `gorm:"column:clusterid;type:uuid;index:k8sroles_sac_filter,type:btree"`
	ClusterName string            `gorm:"column:clustername;type:varchar"`
	ClusterRole bool              `gorm:"column:clusterrole;type:bool"`
	Labels      map[string]string `gorm:"column:labels;type:jsonb"`
	Annotations map[string]string `gorm:"column:annotations;type:jsonb"`
	Serialized  []byte            `gorm:"column:serialized;type:bytea"`
}
