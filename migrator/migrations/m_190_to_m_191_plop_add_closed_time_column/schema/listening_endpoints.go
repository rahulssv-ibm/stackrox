package schema

import (
	"reflect"
	"time"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableListeningEndpointsStmt holds the create statement for table `listening_endpoints`.
	CreateTableListeningEndpointsStmt = &postgres.CreateStmts{
		GormModel: (*ListeningEndpoints)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ListeningEndpointsSchema is the go schema for table `listening_endpoints`.
	ListeningEndpointsSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ProcessListeningOnPortStorage)(nil)), "listening_endpoints")

		schema.SetOptionsMap(search.Walk(v1.SearchCategory_PROCESS_LISTENING_ON_PORT, "processlisteningonportstorage", (*storage.ProcessListeningOnPortStorage)(nil)))
		schema.ScopingResource = resources.DeploymentExtension
		return schema
	}()
)

const (
	// ListeningEndpointsTableName specifies the name of the table in postgres.
	ListeningEndpointsTableName = "listening_endpoints"
)

// ListeningEndpoints holds the Gorm model for Postgres table `listening_endpoints`.
type ListeningEndpoints struct {
	ID                 string             `gorm:"column:id;type:uuid;primaryKey"`
	Port               uint32             `gorm:"column:port;type:bigint"`
	Protocol           storage.L4Protocol `gorm:"column:protocol;type:integer"`
	CloseTimestamp     *time.Time         `gorm:"column:closetimestamp;type:timestamp"`
	ProcessIndicatorID string             `gorm:"column:processindicatorid;type:uuid;index:listeningendpoints_processindicatorid,type:btree"`
	Closed             bool               `gorm:"column:closed;type:bool;index:listeningendpoints_closed,type:btree"`
	DeploymentID       string             `gorm:"column:deploymentid;type:uuid;index:listeningendpoints_deploymentid,type:btree"`
	Serialized         []byte             `gorm:"column:serialized;type:bytea"`
}
