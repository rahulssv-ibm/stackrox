package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableNetworkpolicyapplicationundorecordsStmt holds the create statement for table `networkpolicyapplicationundorecords`.
	CreateTableNetworkpolicyapplicationundorecordsStmt = &postgres.CreateStmts{
		GormModel: (*Networkpolicyapplicationundorecords)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// NetworkpolicyapplicationundorecordsSchema is the go schema for table `networkpolicyapplicationundorecords`.
	NetworkpolicyapplicationundorecordsSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.NetworkPolicyApplicationUndoRecord)(nil)), "networkpolicyapplicationundorecords")
		return schema
	}()
)

const (
	NetworkpolicyapplicationundorecordsTableName = "networkpolicyapplicationundorecords"
)

// Networkpolicyapplicationundorecords holds the Gorm model for Postgres table `networkpolicyapplicationundorecords`.
type Networkpolicyapplicationundorecords struct {
	ClusterId  string `gorm:"column:clusterid;type:uuid;primaryKey"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
