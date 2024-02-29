package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableInstallationInfosStmt holds the create statement for table `installation_infos`.
	CreateTableInstallationInfosStmt = &postgres.CreateStmts{
		GormModel: (*InstallationInfos)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// InstallationInfosSchema is the go schema for table `installation_infos`.
	InstallationInfosSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.InstallationInfo)(nil)), "installation_infos")
		return schema
	}()
)

const (
	InstallationInfosTableName = "installation_infos"
)

// InstallationInfos holds the Gorm model for Postgres table `installation_infos`.
type InstallationInfos struct {
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
