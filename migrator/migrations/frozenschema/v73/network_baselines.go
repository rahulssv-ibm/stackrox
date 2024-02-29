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
	// CreateTableNetworkBaselinesStmt holds the create statement for table `network_baselines`.
	CreateTableNetworkBaselinesStmt = &postgres.CreateStmts{
		GormModel: (*NetworkBaselines)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// NetworkBaselinesSchema is the go schema for table `network_baselines`.
	NetworkBaselinesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.NetworkBaseline)(nil)), "network_baselines")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_NETWORK_BASELINE, "networkbaseline", (*storage.NetworkBaseline)(nil)))
		return schema
	}()
)

const (
	NetworkBaselinesTableName = "network_baselines"
)

// NetworkBaselines holds the Gorm model for Postgres table `network_baselines`.
type NetworkBaselines struct {
	DeploymentId string `gorm:"column:deploymentid;type:uuid;primaryKey"`
	ClusterId    string `gorm:"column:clusterid;type:uuid;index:networkbaselines_sac_filter,type:btree"`
	Namespace    string `gorm:"column:namespace;type:varchar;index:networkbaselines_sac_filter,type:btree"`
	Serialized   []byte `gorm:"column:serialized;type:bytea"`
}
