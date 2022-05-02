// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	storage "github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type ReportconfigsStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	store       Store
	pool        *pgxpool.Pool
}

func TestReportconfigsStore(t *testing.T) {
	suite.Run(t, new(ReportconfigsStoreSuite))
}

func (s *ReportconfigsStoreSuite) SetupTest() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	ctx := sac.WithAllAccess(context.Background())

	source := pgtest.GetConnectionString(s.T())
	config, err := pgxpool.ParseConfig(source)
	s.Require().NoError(err)
	pool, err := pgxpool.ConnectConfig(ctx, config)
	s.Require().NoError(err)

	Destroy(ctx, pool)

	s.pool = pool
	s.store = New(ctx, pool)
}

func (s *ReportconfigsStoreSuite) TearDownTest() {
	if s.pool != nil {
		s.pool.Close()
	}
	s.envIsolator.RestoreAll()
}

func (s *ReportconfigsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	reportConfiguration := &storage.ReportConfiguration{}
	s.NoError(testutils.FullInit(reportConfiguration, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundReportConfiguration, exists, err := store.Get(ctx, reportConfiguration.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundReportConfiguration)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, reportConfiguration))
	foundReportConfiguration, exists, err = store.Get(ctx, reportConfiguration.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(reportConfiguration, foundReportConfiguration)

	reportConfigurationCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(1, reportConfigurationCount)
	reportConfigurationCount, err = store.Count(withNoAccessCtx)
	s.NoError(err)
	s.Zero(reportConfigurationCount)

	reportConfigurationExists, err := store.Exists(ctx, reportConfiguration.GetId())
	s.NoError(err)
	s.True(reportConfigurationExists)
	s.NoError(store.Upsert(ctx, reportConfiguration))
	s.ErrorIs(store.Upsert(withNoAccessCtx, reportConfiguration), sac.ErrResourceAccessDenied)

	foundReportConfiguration, exists, err = store.Get(ctx, reportConfiguration.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(reportConfiguration, foundReportConfiguration)

	s.NoError(store.Delete(ctx, reportConfiguration.GetId()))
	foundReportConfiguration, exists, err = store.Get(ctx, reportConfiguration.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundReportConfiguration)
	s.ErrorIs(store.Delete(withNoAccessCtx, reportConfiguration.GetId()), sac.ErrResourceAccessDenied)

	var reportConfigurations []*storage.ReportConfiguration
	for i := 0; i < 200; i++ {
		reportConfiguration := &storage.ReportConfiguration{}
		s.NoError(testutils.FullInit(reportConfiguration, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		reportConfigurations = append(reportConfigurations, reportConfiguration)
	}

	s.NoError(store.UpsertMany(ctx, reportConfigurations))

	reportConfigurationCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(200, reportConfigurationCount)
}
