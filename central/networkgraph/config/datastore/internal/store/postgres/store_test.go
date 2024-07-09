// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/protoassert"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/suite"
)

type NetworkGraphConfigsStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestNetworkGraphConfigsStore(t *testing.T) {
	suite.Run(t, new(NetworkGraphConfigsStoreSuite))
}

func (s *NetworkGraphConfigsStoreSuite) SetupSuite() {

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *NetworkGraphConfigsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE network_graph_configs CASCADE")
	s.T().Log("network_graph_configs", tag)
	s.store = New(s.testDB.DB)
	s.NoError(err)
}

func (s *NetworkGraphConfigsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *NetworkGraphConfigsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	networkGraphConfig := &storage.NetworkGraphConfig{}
	s.NoError(testutils.FullInit(networkGraphConfig, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundNetworkGraphConfig, exists, err := store.Get(ctx, networkGraphConfig.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNetworkGraphConfig)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, networkGraphConfig))
	foundNetworkGraphConfig, exists, err = store.Get(ctx, networkGraphConfig.GetId())
	s.NoError(err)
	s.True(exists)
	protoassert.Equal(s.T(), networkGraphConfig, foundNetworkGraphConfig)

	networkGraphConfigCount, err := store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(1, networkGraphConfigCount)
	networkGraphConfigCount, err = store.Count(withNoAccessCtx, search.EmptyQuery())
	s.NoError(err)
	s.Zero(networkGraphConfigCount)

	networkGraphConfigExists, err := store.Exists(ctx, networkGraphConfig.GetId())
	s.NoError(err)
	s.True(networkGraphConfigExists)
	s.NoError(store.Upsert(ctx, networkGraphConfig))
	s.ErrorIs(store.Upsert(withNoAccessCtx, networkGraphConfig), sac.ErrResourceAccessDenied)

	s.NoError(store.Delete(ctx, networkGraphConfig.GetId()))
	foundNetworkGraphConfig, exists, err = store.Get(ctx, networkGraphConfig.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNetworkGraphConfig)
	s.ErrorIs(store.Delete(withNoAccessCtx, networkGraphConfig.GetId()), sac.ErrResourceAccessDenied)

	var networkGraphConfigs []*storage.NetworkGraphConfig
	var networkGraphConfigIDs []string
	for i := 0; i < 200; i++ {
		networkGraphConfig := &storage.NetworkGraphConfig{}
		s.NoError(testutils.FullInit(networkGraphConfig, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		networkGraphConfigs = append(networkGraphConfigs, networkGraphConfig)
		networkGraphConfigIDs = append(networkGraphConfigIDs, networkGraphConfig.GetId())
	}

	s.NoError(store.UpsertMany(ctx, networkGraphConfigs))

	networkGraphConfigCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(200, networkGraphConfigCount)

	s.NoError(store.DeleteMany(ctx, networkGraphConfigIDs))

	networkGraphConfigCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(0, networkGraphConfigCount)
}
