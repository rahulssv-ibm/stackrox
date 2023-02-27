// Code generated by pg-bindings generator. DO NOT EDIT.
package n47ton48

import (
	"context"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	frozenSchema "github.com/stackrox/rox/migrator/migrations/frozenschema/v1"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	legacy "github.com/stackrox/rox/migrator/migrations/n_47_to_n_48_postgres_secrets/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_47_to_n_48_postgres_secrets/postgres"
	"github.com/stackrox/rox/migrator/types"
	pkgMigrations "github.com/stackrox/rox/pkg/migrations"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/sac"
	"gorm.io/gorm"
)

var (
	startingSeqNum = pkgMigrations.BasePostgresDBVersionSeqNum() + 47 // 158

	migration = types.Migration{
		StartingSeqNum: startingSeqNum,
		VersionAfter:   &storage.Version{SeqNum: int32(startingSeqNum + 1)}, // 159
		Run: func(databases *types.Databases) error {
			legacyStore, err := legacy.New(databases.PkgRocksDB)
			if err != nil {
				return err
			}
			if err := move(databases.GormDB, databases.PostgresDB, legacyStore); err != nil {
				return errors.Wrap(err,
					"moving secrets from rocksdb to postgres")
			}
			return nil
		},
	}
	batchSize = 10000
	schema    = frozenSchema.SecretsSchema
	log       = loghelper.LogWrapper{}
)

func move(gormDB *gorm.DB, postgresDB *postgres.DB, legacyStore legacy.Store) error {
	ctx := sac.WithAllAccess(context.Background())
	store := pgStore.New(postgresDB)
	pgutils.CreateTableFromModel(context.Background(), gormDB, frozenSchema.CreateTableSecretsStmt)
	var secrets []*storage.Secret
	err := walk(ctx, legacyStore, func(obj *storage.Secret) error {
		secrets = append(secrets, obj)
		if len(secrets) == batchSize {
			if err := store.UpsertMany(ctx, secrets); err != nil {
				log.WriteToStderrf("failed to persist secrets to store %v", err)
				return err
			}
			secrets = secrets[:0]
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(secrets) > 0 {
		if err = store.UpsertMany(ctx, secrets); err != nil {
			log.WriteToStderrf("failed to persist secrets to store %v", err)
			return err
		}
	}
	return nil
}

func walk(ctx context.Context, s legacy.Store, fn func(obj *storage.Secret) error) error {
	return s.Walk(ctx, fn)
}

func init() {
	migrations.MustRegisterMigration(migration)
}
