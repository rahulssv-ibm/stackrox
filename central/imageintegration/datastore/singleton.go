package datastore

import (
	"context"

	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/imageintegration/search"
	"github.com/stackrox/rox/central/imageintegration/store"
	pgStore "github.com/stackrox/rox/central/imageintegration/store/postgres"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/pkg/utils"
)

var (
	once sync.Once

	dataStore DataStore
)

func initializeIntegrations(storage store.Store) {
	ctx := sac.WithGlobalAccessScopeChecker(context.Background(), sac.AllowAllAccessScopeChecker())
	iis, err := storage.GetAll(ctx)
	utils.CrashOnError(err)
	// If we are starting from scratch in online-mode, add the default image integrations.
	if !env.OfflineModeEnv.BooleanSetting() && len(iis) == 0 {
		// Add default integrations
		for _, ii := range store.DefaultImageIntegrations {
			utils.Should(storage.Upsert(ctx, ii))
		}
	}

	// If the feature flag is disabled, remove all "sourced" autogenerated registries from the datastore.
	if !features.SourcedAutogeneratedIntegrations.Enabled() {
		if len(iis) > 0 {
			log.Infof("[STARTUP] Starting deletion of 'sourced' image integrations")
		}

		var attempted, deleted int
		for _, ii := range iis {
			if ii.GetAutogenerated() && ii.GetSource() != nil {
				attempted++
				// Use Should so release versions do not panic.
				if err := utils.ShouldErr(storage.Delete(ctx, ii.GetId())); err == nil {
					deleted++
				}
			}
		}
		if attempted > 0 {
			log.Infof("Successfully deleted %d out of %d image integration(s)", deleted, attempted)
		}

		log.Info("Completed deletion of 'sourced' image integrations")
	}
}

func initialize() {
	// Create underlying store and datastore.
	storage := pgStore.New(globaldb.GetPostgres())

	initializeIntegrations(storage)
	searcher := search.New(storage, pgStore.NewIndexer(globaldb.GetPostgres()))
	dataStore = New(storage, searcher)
}

// Singleton provides the interface for non-service external interaction.
func Singleton() DataStore {
	once.Do(initialize)
	return dataStore
}
