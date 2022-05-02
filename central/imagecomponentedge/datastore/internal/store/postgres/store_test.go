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

type ImageComponentRelationsStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	store       Store
	pool        *pgxpool.Pool
}

func TestImageComponentRelationsStore(t *testing.T) {
	suite.Run(t, new(ImageComponentRelationsStoreSuite))
}

func (s *ImageComponentRelationsStoreSuite) SetupTest() {
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

func (s *ImageComponentRelationsStoreSuite) TearDownTest() {
	if s.pool != nil {
		s.pool.Close()
	}
	s.envIsolator.RestoreAll()
}

func (s *ImageComponentRelationsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	imageComponentEdge := &storage.ImageComponentEdge{}
	s.NoError(testutils.FullInit(imageComponentEdge, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundImageComponentEdge, exists, err := store.Get(ctx, imageComponentEdge.GetId(), imageComponentEdge.GetImageId(), imageComponentEdge.GetImageComponentId(), imageComponentEdge.GetImageComponentName(), imageComponentEdge.GetImageComponentVersion(), imageComponentEdge.GetImageComponentOperatingSystem())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundImageComponentEdge)

}
