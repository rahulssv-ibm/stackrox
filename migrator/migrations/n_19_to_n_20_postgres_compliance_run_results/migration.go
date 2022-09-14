// Code generated by pg-bindings generator. DO NOT EDIT.
package n19ton20

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	legacy "github.com/stackrox/rox/migrator/migrations/n_19_to_n_20_postgres_compliance_run_results/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_19_to_n_20_postgres_compliance_run_results/postgres"
	"github.com/stackrox/rox/migrator/types"
	pkgMigrations "github.com/stackrox/rox/pkg/migrations"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac"
	"gorm.io/gorm"
)

var (
	migration = types.Migration{
		StartingSeqNum: pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres() + 19,
		VersionAfter:   storage.Version{SeqNum: int32(pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres()) + 20},
		Run: func(databases *types.Databases) error {
			legacyStore, err := legacy.New(databases.PkgRocksDB)
			if err != nil {
				return err
			}
			if err := move(databases.GormDB, databases.PostgresDB, legacyStore); err != nil {
				return errors.Wrap(err,
					"moving compliance_run_results from rocksdb to postgres")
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
	schema    = pkgSchema.ComplianceRunResultsSchema
	log       = loghelper.LogWrapper{}
)

func move(gormDB *gorm.DB, postgresDB *pgxpool.Pool, legacyStore legacy.Store) error {
	ctx := sac.WithAllAccess(context.Background())
	store := pgStore.New(postgresDB)
	// We need to migrate so turn off foreign key constraints
	gormConfig := gormDB.Config
	gormConfig.DisableForeignKeyConstraintWhenMigrating = true
	err := gormDB.Apply(gormConfig)
	if err != nil {
		return errors.Wrap(err, "failed to turn off foreign key constraints")
	}
	pkgSchema.ApplySchemaForTable(context.Background(), gormDB, schema.Table)
	var complianceRunResults []*storage.ComplianceRunResults
	err = walk(ctx, legacyStore, func(obj *storage.ComplianceRunResults) error {
		complianceRunResults = append(complianceRunResults, obj)
		if len(complianceRunResults) == batchSize {
			if err := store.UpsertMany(ctx, complianceRunResults); err != nil {
				log.WriteToStderrf("failed to persist compliance_run_results to store %v", err)
				return err
			}
			complianceRunResults = complianceRunResults[:0]
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(complianceRunResults) > 0 {
		if err = store.UpsertMany(ctx, complianceRunResults); err != nil {
			log.WriteToStderrf("failed to persist compliance_run_results to store %v", err)
			return err
		}
	}
	return nil
}

func walk(ctx context.Context, s legacy.Store, fn func(obj *storage.ComplianceRunResults) error) error {
	return s.Walk(ctx, fn)
}

func init() {
	migrations.MustRegisterMigration(migration)
}
