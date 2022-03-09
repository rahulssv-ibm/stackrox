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
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type ReportconfigsStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
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
}

func (s *ReportconfigsStoreSuite) TearDownTest() {
	s.envIsolator.RestoreAll()
}

func (s *ReportconfigsStoreSuite) TestStore() {
	source := pgtest.GetConnectionString(s.T())
	config, err := pgxpool.ParseConfig(source)
	s.Require().NoError(err)
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	s.NoError(err)
	defer pool.Close()

	Destroy(pool)
	store := New(pool)

	reportConfiguration := &storage.ReportConfiguration{}
	s.NoError(testutils.FullInit(reportConfiguration, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundReportConfiguration, exists, err := store.Get(reportConfiguration.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundReportConfiguration)

	s.NoError(store.Upsert(reportConfiguration))
	foundReportConfiguration, exists, err = store.Get(reportConfiguration.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(reportConfiguration, foundReportConfiguration)

	reportConfigurationCount, err := store.Count()
	s.NoError(err)
	s.Equal(reportConfigurationCount, 1)

	reportConfigurationExists, err := store.Exists(reportConfiguration.GetId())
	s.NoError(err)
	s.True(reportConfigurationExists)
	s.NoError(store.Upsert(reportConfiguration))

	foundReportConfiguration, exists, err = store.Get(reportConfiguration.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(reportConfiguration, foundReportConfiguration)

	s.NoError(store.Delete(reportConfiguration.GetId()))
	foundReportConfiguration, exists, err = store.Get(reportConfiguration.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundReportConfiguration)
}
