// Code generated by pg-bindings generator. DO NOT EDIT.
package n25ton26

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	frozenSchema "github.com/stackrox/rox/migrator/migrations/frozenschema/v73"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	legacy "github.com/stackrox/rox/migrator/migrations/n_25_to_n_26_postgres_integration_healths/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_25_to_n_26_postgres_integration_healths/postgres"
	"github.com/stackrox/rox/migrator/types"
	pkgMigrations "github.com/stackrox/rox/pkg/migrations"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/sac"
	"gorm.io/gorm"
)

var (
	migration = types.Migration{
		StartingSeqNum: pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres() + 25,
		VersionAfter:   &storage.Version{SeqNum: int32(pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres()) + 26},
		Run: func(databases *types.Databases) error {
			legacyStore, err := legacy.New(databases.PkgRocksDB)
			if err != nil {
				return err
			}
			if err := move(databases.GormDB, databases.PostgresDB, legacyStore); err != nil {
				return errors.Wrap(err,
					"moving integration_healths from rocksdb to postgres")
			}
			return nil
		},
		LegacyToPostgres: true,
	}
	batchSize = 10000
	schema    = frozenSchema.IntegrationHealthsSchema
	log       = loghelper.LogWrapper{}
)

func move(gormDB *gorm.DB, postgresDB *pgxpool.Pool, legacyStore legacy.Store) error {
	ctx := sac.WithAllAccess(context.Background())
	store := pgStore.New(postgresDB)
	pgutils.CreateTableFromModel(context.Background(), gormDB, frozenSchema.CreateTableIntegrationHealthsStmt)
	var integrationHealths []*storage.IntegrationHealth
	err := walk(ctx, legacyStore, func(obj *storage.IntegrationHealth) error {
		integrationHealths = append(integrationHealths, obj)
		if len(integrationHealths) == batchSize {
			if err := store.UpsertMany(ctx, integrationHealths); err != nil {
				log.WriteToStderrf("failed to persist integration_healths to store %v", err)
				return err
			}
			integrationHealths = integrationHealths[:0]
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(integrationHealths) > 0 {
		if err = store.UpsertMany(ctx, integrationHealths); err != nil {
			log.WriteToStderrf("failed to persist integration_healths to store %v", err)
			return err
		}
	}
	return nil
}

func walk(ctx context.Context, s legacy.Store, fn func(obj *storage.IntegrationHealth) error) error {
	return s.Walk(ctx, fn)
}

func init() {
	migrations.MustRegisterMigration(migration)
}
