package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableServiceIdentitiesStmt holds the create statement for table `service_identities`.
	CreateTableServiceIdentitiesStmt = &postgres.CreateStmts{
		GormModel: (*ServiceIdentities)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ServiceIdentitiesSchema is the go schema for table `service_identities`.
	ServiceIdentitiesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.ServiceIdentity)(nil)), "service_identities")
		return schema
	}()
)

const (
	ServiceIdentitiesTableName = "service_identities"
)

// ServiceIdentities holds the Gorm model for Postgres table `service_identities`.
type ServiceIdentities struct {
	SerialStr  string `gorm:"column:serialstr;type:varchar;primaryKey"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
