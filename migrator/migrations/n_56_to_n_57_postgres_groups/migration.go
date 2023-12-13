// Code originally generated by pg-bindings generator.

package n56ton57

import (
	"context"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	frozenSchema "github.com/stackrox/rox/migrator/migrations/frozenschema/v73"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	legacy "github.com/stackrox/rox/migrator/migrations/n_56_to_n_57_postgres_groups/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_56_to_n_57_postgres_groups/postgres"
	"github.com/stackrox/rox/migrator/types"
	pkgMigrations "github.com/stackrox/rox/pkg/migrations"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"gorm.io/gorm"
)

var (
	startingSeqNum = pkgMigrations.BasePostgresDBVersionSeqNum() + 56 // 167

	migration = types.Migration{
		StartingSeqNum: startingSeqNum,
		VersionAfter:   &storage.Version{SeqNum: int32(startingSeqNum + 1)}, // 168
		Run: func(databases *types.Databases) error {
			legacyStore := legacy.New(databases.BoltDB)
			if err := move(databases.DBCtx, databases.GormDB, databases.PostgresDB, legacyStore); err != nil {
				return errors.Wrap(err,
					"moving groups from rocksdb to postgres")
			}
			return nil
		},
	}
	log = loghelper.LogWrapper{}
)

func move(ctx context.Context, gormDB *gorm.DB, postgresDB postgres.DB, legacyStore legacy.Store) error {
	store := pgStore.New(postgresDB)
	pgutils.CreateTableFromModel(context.Background(), gormDB, frozenSchema.CreateTableGroupsStmt)

	groups, err := legacyStore.GetAll(ctx)
	if err != nil {
		return err
	}
	if len(groups) > 0 {
		if err = store.UpsertMany(ctx, groups); err != nil {
			log.WriteToStderrf("failed to persist groups to store %v", err)
			return err
		}
	}
	return nil
}

func init() {
	migrations.MustRegisterMigration(migration)
}
