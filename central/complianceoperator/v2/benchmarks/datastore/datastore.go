package datastore

import (
	"context"
	"testing"

	pgStore "github.com/stackrox/rox/central/complianceoperator/v2/benchmarks/store/postgres"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
)

// DataStore is the entry point for storing/retrieving compliance operator scan objects.
//
//go:generate mockgen-wrapper
type DataStore interface {
	// GetBenchmark retrieves the benchmark object from the database
	GetBenchmark(ctx context.Context, id string) (*storage.ComplianceOperatorBenchmarkV2, bool, error)

	// UpsertBenchmark adds the benchmark object to the database
	UpsertBenchmark(ctx context.Context, result *storage.ComplianceOperatorBenchmarkV2) error

	// DeleteBenchmark removes a benchmark object from the database
	DeleteBenchmark(ctx context.Context, id string) error
	GetControlByRuleName(ctx context.Context, ruleNames []string) ([]*ControlResult, error)
}

// New returns an instance of DataStore.
func New(benchmarkStorage pgStore.Store) DataStore {
	return &datastoreImpl{
		store: benchmarkStorage,
	}
}

// GetTestPostgresDataStore provides a datastore connected to postgres for testing purposes.
func GetTestPostgresDataStore(_ *testing.T, pool postgres.DB) DataStore {
	store := pgStore.New(pool)
	return New(store)
}
