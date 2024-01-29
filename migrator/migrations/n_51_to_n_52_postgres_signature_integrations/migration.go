// Code originally generated by pg-bindings generator.

package n51ton52

import (
	"context"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	frozenSchema "github.com/stackrox/rox/migrator/migrations/frozenschema/v73"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	legacy "github.com/stackrox/rox/migrator/migrations/n_51_to_n_52_postgres_signature_integrations/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_51_to_n_52_postgres_signature_integrations/postgres"
	"github.com/stackrox/rox/migrator/types"
	pkgMigrations "github.com/stackrox/rox/pkg/migrations"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"gorm.io/gorm"
)

var (
	startingSeqNum = pkgMigrations.BasePostgresDBVersionSeqNum() + 51 // 162

	migration = types.Migration{
		StartingSeqNum: startingSeqNum,
		VersionAfter:   &storage.Version{SeqNum: int32(startingSeqNum + 1)}, // 163
		Run: func(databases *types.Databases) error {
			legacyStore, err := legacy.New(databases.PkgRocksDB)
			if err != nil {
				return err
			}
			if err := move(databases.DBCtx, databases.GormDB, databases.PostgresDB, legacyStore); err != nil {
				return errors.Wrap(err,
					"moving signature_integrations from rocksdb to postgres")
			}
			return nil
		},
	}
	batchSize = 10000
	log       = loghelper.LogWrapper{}
)

func move(ctx context.Context, gormDB *gorm.DB, postgresDB postgres.DB, legacyStore legacy.Store) error {
	store := pgStore.New(postgresDB)
	pgutils.CreateTableFromModel(context.Background(), gormDB, frozenSchema.CreateTableSignatureIntegrationsStmt)

	var signatureIntegrations []*storage.SignatureIntegration
	err := walk(ctx, legacyStore, func(obj *storage.SignatureIntegration) error {
		signatureIntegrations = append(signatureIntegrations, obj)
		if len(signatureIntegrations) == batchSize {
			if err := store.UpsertMany(ctx, signatureIntegrations); err != nil {
				log.WriteToStderrf("failed to persist signature_integrations to store %v", err)
				return err
			}
			signatureIntegrations = signatureIntegrations[:0]
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(signatureIntegrations) > 0 {
		if err = store.UpsertMany(ctx, signatureIntegrations); err != nil {
			log.WriteToStderrf("failed to persist signature_integrations to store %v", err)
			return err
		}
	}
	return nil
}

func walk(ctx context.Context, s legacy.Store, fn func(obj *storage.SignatureIntegration) error) error {
	return s.Walk(ctx, fn)
}

func init() {
	migrations.MustRegisterMigration(migration)
}
