package datastore

import (
	"context"

	"github.com/stackrox/rox/central/complianceoperator/v2/suites/store/postgres"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/errox"
	"github.com/stackrox/rox/pkg/search"
)

type datastoreImpl struct {
	store postgres.Store
}

// GetSuite returns the suite with the name
func (d *datastoreImpl) GetSuite(ctx context.Context, name string) (*storage.ComplianceOperatorSuite, error) {
	return nil, errox.NotImplemented
}

// UpsertSuite adds the suite to the database
func (d *datastoreImpl) UpsertSuite(_ context.Context, _ *storage.ComplianceOperatorSuite) error {
	return errox.NotImplemented
}

// UpsertSuites adds the suites to the database
func (d *datastoreImpl) UpsertSuites(_ context.Context, _ []*storage.ComplianceOperatorSuite) error {
	return errox.NotImplemented
}

// DeleteSuite removes a suite from the database
func (d *datastoreImpl) DeleteSuite(_ context.Context, _ string) error {
	return errox.NotImplemented
}

// GetSuitesByCluster retrieve rules by the cluster
func (d *datastoreImpl) GetSuitesByCluster(ctx context.Context, clusterID string) ([]*storage.ComplianceOperatorSuite, error) {
	return d.store.GetByQuery(ctx, search.NewQueryBuilder().
		AddExactMatches(search.ClusterID, clusterID).ProtoQuery())
}
