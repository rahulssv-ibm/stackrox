package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableNetworkpoliciesundodeploymentsStmt holds the create statement for table `networkpoliciesundodeployments`.
	CreateTableNetworkpoliciesundodeploymentsStmt = &postgres.CreateStmts{
		GormModel: (*Networkpoliciesundodeployments)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// NetworkpoliciesundodeploymentsSchema is the go schema for table `networkpoliciesundodeployments`.
	NetworkpoliciesundodeploymentsSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.NetworkPolicyApplicationUndoDeploymentRecord)(nil)), "networkpoliciesundodeployments")
		return schema
	}()
)

const (
	NetworkpoliciesundodeploymentsTableName = "networkpoliciesundodeployments"
)

// Networkpoliciesundodeployments holds the Gorm model for Postgres table `networkpoliciesundodeployments`.
type Networkpoliciesundodeployments struct {
	DeploymentId string `gorm:"column:deploymentid;type:uuid;primaryKey"`
	Serialized   []byte `gorm:"column:serialized;type:bytea"`
}
