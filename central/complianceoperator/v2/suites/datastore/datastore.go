package datastore

import (
	"context"
	"testing"

	pgStore "github.com/stackrox/rox/central/complianceoperator/v2/suites/store/postgres"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
)

// DataStore is the entry point for storing/retrieving compliance operator suite.
//
//go:generate mockgen-wrapper
type DataStore interface {
	// GetSuite returns the suite for the given name
	GetSuite(ctx context.Context, name string) (*storage.ComplianceOperatorSuite, error)

	// GetSuitesByCluster retrieve rules by the cluster
	GetSuitesByCluster(ctx context.Context, clusterID string) ([]*storage.ComplianceOperatorSuite, error)

	// UpsertSuite adds the suite to the database
	UpsertSuite(ctx context.Context, suite *storage.ComplianceOperatorSuite) error

	// UpsertSuites adds the suites to the database
	UpsertSuites(ctx context.Context, suites []*storage.ComplianceOperatorSuite) error

	// DeleteSuite removes a suite from the database
	DeleteSuite(ctx context.Context, name string) error
}

// New returns an instance of DataStore.
func New(complianceSuiteStorage pgStore.Store) DataStore {
	ds := &datastoreImpl{
		store: complianceSuiteStorage,
	}
	return ds
}

// NewForTestOnly returns an instance of DataStore only for tests.
func NewForTestOnly(_ *testing.T, complianceSuiteStorage pgStore.Store) DataStore {
	ds := &datastoreImpl{
		store: complianceSuiteStorage,
	}
	return ds
}

// GetTestPostgresDataStore provides a datastore connected to postgres for testing purposes.
func GetTestPostgresDataStore(_ *testing.T, pool postgres.DB) (DataStore, error) {
	store := pgStore.New(pool)
	return New(store), nil
}
