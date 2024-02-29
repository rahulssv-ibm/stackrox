package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableSensorUpgradeConfigsStmt holds the create statement for table `sensor_upgrade_configs`.
	CreateTableSensorUpgradeConfigsStmt = &postgres.CreateStmts{
		GormModel: (*SensorUpgradeConfigs)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// SensorUpgradeConfigsSchema is the go schema for table `sensor_upgrade_configs`.
	SensorUpgradeConfigsSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.SensorUpgradeConfig)(nil)), "sensor_upgrade_configs")
		return schema
	}()
)

const (
	SensorUpgradeConfigsTableName = "sensor_upgrade_configs"
)

// SensorUpgradeConfigs holds the Gorm model for Postgres table `sensor_upgrade_configs`.
type SensorUpgradeConfigs struct {
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
