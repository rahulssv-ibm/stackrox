// Code generated by pg-bindings generator. DO NOT EDIT.
package n17ton18

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	frozenSchema "github.com/stackrox/rox/migrator/migrations/frozenschema/v73"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	legacy "github.com/stackrox/rox/migrator/migrations/n_17_to_n_18_postgres_compliance_operator_scans/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_17_to_n_18_postgres_compliance_operator_scans/postgres"
	"github.com/stackrox/rox/migrator/types"
	pkgMigrations "github.com/stackrox/rox/pkg/migrations"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/sac"
	"gorm.io/gorm"
)

var (
	migration = types.Migration{
		StartingSeqNum: pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres() + 17,
		VersionAfter:   &storage.Version{SeqNum: int32(pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres()) + 18},
		Run: func(databases *types.Databases) error {
			legacyStore, err := legacy.New(databases.PkgRocksDB)
			if err != nil {
				return err
			}
			if err := move(databases.GormDB, databases.PostgresDB, legacyStore); err != nil {
				return errors.Wrap(err,
					"moving compliance_operator_scans from rocksdb to postgres")
			}
			return nil
		},
		LegacyToPostgres: true,
	}
	batchSize = 10000
	schema    = frozenSchema.ComplianceOperatorScansSchema
	log       = loghelper.LogWrapper{}
)

func move(gormDB *gorm.DB, postgresDB *pgxpool.Pool, legacyStore legacy.Store) error {
	ctx := sac.WithAllAccess(context.Background())
	store := pgStore.New(postgresDB)
	pgutils.CreateTableFromModel(context.Background(), gormDB, frozenSchema.CreateTableComplianceOperatorScansStmt)
	var complianceOperatorScans []*storage.ComplianceOperatorScan
	err := walk(ctx, legacyStore, func(obj *storage.ComplianceOperatorScan) error {
		complianceOperatorScans = append(complianceOperatorScans, obj)
		if len(complianceOperatorScans) == batchSize {
			if err := store.UpsertMany(ctx, complianceOperatorScans); err != nil {
				log.WriteToStderrf("failed to persist compliance_operator_scans to store %v", err)
				return err
			}
			complianceOperatorScans = complianceOperatorScans[:0]
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(complianceOperatorScans) > 0 {
		if err = store.UpsertMany(ctx, complianceOperatorScans); err != nil {
			log.WriteToStderrf("failed to persist compliance_operator_scans to store %v", err)
			return err
		}
	}
	return nil
}

func walk(ctx context.Context, s legacy.Store, fn func(obj *storage.ComplianceOperatorScan) error) error {
	return s.Walk(ctx, fn)
}

func init() {
	migrations.MustRegisterMigration(migration)
}
