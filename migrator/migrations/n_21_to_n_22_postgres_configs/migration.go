// Code generated by pg-bindings generator. DO NOT EDIT.
package n21ton22

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	legacy "github.com/stackrox/rox/migrator/migrations/n_21_to_n_22_postgres_configs/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_21_to_n_22_postgres_configs/postgres"
	"github.com/stackrox/rox/migrator/types"
	pkgMigrations "github.com/stackrox/rox/pkg/migrations"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac"
	"gorm.io/gorm"
)

var (
	migration = types.Migration{
		StartingSeqNum: pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres() + 21,
		VersionAfter:   storage.Version{SeqNum: int32(pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres()) + 22},
		Run: func(databases *types.Databases) error {
			legacyStore := legacy.New(databases.BoltDB)
			if err := move(databases.GormDB, databases.PostgresDB, legacyStore); err != nil {
				return errors.Wrap(err,
					"moving configs from rocksdb to postgres")
			}
			// Now that migrations are complete, turn the constraints back on
			gormConfig := databases.GormDB.Config
			gormConfig.DisableForeignKeyConstraintWhenMigrating = false
			err := databases.GormDB.Apply(gormConfig)
			if err != nil {
				return errors.Wrap(err, "failed to turn on foreign key constraints")
			}
			pkgSchema.ApplySchemaForTable(context.Background(), databases.GormDB, schema.Table)
			return nil
		},
	}
	batchSize = 10000
	schema    = pkgSchema.ConfigsSchema
	log       = loghelper.LogWrapper{}
)

func move(gormDB *gorm.DB, postgresDB *pgxpool.Pool, legacyStore legacy.Store) error {
	ctx := sac.WithAllAccess(context.Background())
	store := pgStore.New(ctx, postgresDB)
	// We need to migrate so turn off foreign key constraints
	gormConfig := gormDB.Config
	gormConfig.DisableForeignKeyConstraintWhenMigrating = true
	err := gormDB.Apply(gormConfig)
	if err != nil {
		return errors.Wrap(err, "failed to turn off foreign key constraints")
	}
	pkgSchema.ApplySchemaForTable(context.Background(), gormDB, schema.Table)
	obj, found, err := legacyStore.Get(ctx)
	if err != nil {
		log.WriteToStderr("failed to fetch config")
		return err
	}
	if !found {
		return nil
	}
	if err = store.Upsert(ctx, obj); err != nil {
		log.WriteToStderrf("failed to persist object to store %v", err)
		return err
	}
	return nil
}

func init() {
	migrations.MustRegisterMigration(migration)
}
