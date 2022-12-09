// Code generated by pg-bindings generator. DO NOT EDIT.
package n33ton34

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	frozenSchema "github.com/stackrox/rox/migrator/migrations/frozenschema/v73"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	legacy "github.com/stackrox/rox/migrator/migrations/n_33_to_n_34_postgres_networkpoliciesundodeployments/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_33_to_n_34_postgres_networkpoliciesundodeployments/postgres"
	"github.com/stackrox/rox/migrator/types"
	pkgMigrations "github.com/stackrox/rox/pkg/migrations"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/sac"
	"gorm.io/gorm"
)

var (
	migration = types.Migration{
		StartingSeqNum: pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres() + 33,
		VersionAfter:   &storage.Version{SeqNum: int32(pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres()) + 34},
		Run: func(databases *types.Databases) error {
			if databases.PkgRocksDB == nil {
				log.WriteToStderr("RocksDB is nil.  Skipping migration n33ton34")
				return nil
			}
			legacyStore, err := legacy.New(databases.PkgRocksDB)
			if err != nil {
				return err
			}
			if err := move(databases.GormDB, databases.PostgresDB, legacyStore); err != nil {
				return errors.Wrap(err,
					"moving networkpoliciesundodeployments from rocksdb to postgres")
			}
			return nil
		},
	}
	batchSize = 10000
	schema    = frozenSchema.NetworkpoliciesundodeploymentsSchema
	log       = loghelper.LogWrapper{}
)

func move(gormDB *gorm.DB, postgresDB *pgxpool.Pool, legacyStore legacy.Store) error {
	ctx := sac.WithAllAccess(context.Background())
	store := pgStore.New(postgresDB)
	pgutils.CreateTableFromModel(context.Background(), gormDB, frozenSchema.CreateTableNetworkpoliciesundodeploymentsStmt)
	var networkpoliciesundodeployments []*storage.NetworkPolicyApplicationUndoDeploymentRecord
	err := walk(ctx, legacyStore, func(obj *storage.NetworkPolicyApplicationUndoDeploymentRecord) error {
		networkpoliciesundodeployments = append(networkpoliciesundodeployments, obj)
		if len(networkpoliciesundodeployments) == batchSize {
			if err := store.UpsertMany(ctx, networkpoliciesundodeployments); err != nil {
				log.WriteToStderrf("failed to persist networkpoliciesundodeployments to store %v", err)
				return err
			}
			networkpoliciesundodeployments = networkpoliciesundodeployments[:0]
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(networkpoliciesundodeployments) > 0 {
		if err = store.UpsertMany(ctx, networkpoliciesundodeployments); err != nil {
			log.WriteToStderrf("failed to persist networkpoliciesundodeployments to store %v", err)
			return err
		}
	}
	return nil
}

func walk(ctx context.Context, s legacy.Store, fn func(obj *storage.NetworkPolicyApplicationUndoDeploymentRecord) error) error {
	return s.Walk(ctx, fn)
}

func init() {
	migrations.MustRegisterMigration(migration)
}
