// Code generated by pg-bindings generator. DO NOT EDIT.
package n28ton29

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	frozenSchema "github.com/stackrox/rox/migrator/migrations/frozenschema/v73"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	legacy "github.com/stackrox/rox/migrator/migrations/n_28_to_n_29_postgres_network_baselines/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_28_to_n_29_postgres_network_baselines/postgres"
	"github.com/stackrox/rox/migrator/types"
	pkgMigrations "github.com/stackrox/rox/pkg/migrations"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/sac"
	"gorm.io/gorm"
)

var (
	migration = types.Migration{
		StartingSeqNum: pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres() + 28,
		VersionAfter:   &storage.Version{SeqNum: int32(pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres()) + 29},
		Run: func(databases *types.Databases) error {
			legacyStore, err := legacy.New(databases.PkgRocksDB)
			if err != nil {
				return err
			}
			if err := move(databases.GormDB, databases.PostgresDB, legacyStore); err != nil {
				return errors.Wrap(err,
					"moving network_baselines from rocksdb to postgres")
			}
			return nil
		},
		LegacyToPostgres: true,
	}
	batchSize = 10000
	schema    = frozenSchema.NetworkBaselinesSchema
	log       = loghelper.LogWrapper{}
)

func move(gormDB *gorm.DB, postgresDB *pgxpool.Pool, legacyStore legacy.Store) error {
	ctx := sac.WithAllAccess(context.Background())
	store := pgStore.New(postgresDB)
	pgutils.CreateTableFromModel(context.Background(), gormDB, frozenSchema.CreateTableNetworkBaselinesStmt)
	var networkBaselines []*storage.NetworkBaseline
	err := walk(ctx, legacyStore, func(obj *storage.NetworkBaseline) error {
		networkBaselines = append(networkBaselines, obj)
		if len(networkBaselines) == batchSize {
			if err := store.UpsertMany(ctx, networkBaselines); err != nil {
				log.WriteToStderrf("failed to persist network_baselines to store %v", err)
				return err
			}
			networkBaselines = networkBaselines[:0]
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(networkBaselines) > 0 {
		if err = store.UpsertMany(ctx, networkBaselines); err != nil {
			log.WriteToStderrf("failed to persist network_baselines to store %v", err)
			return err
		}
	}
	return nil
}

func walk(ctx context.Context, s legacy.Store, fn func(obj *storage.NetworkBaseline) error) error {
	return s.Walk(ctx, fn)
}

func init() {
	migrations.MustRegisterMigration(migration)
}
