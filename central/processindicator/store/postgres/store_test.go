// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	storage "github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ProcessIndicatorsStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	store       Store
	pool        *pgxpool.Pool
}

func TestProcessIndicatorsStore(t *testing.T) {
	suite.Run(t, new(ProcessIndicatorsStoreSuite))
}

func (s *ProcessIndicatorsStoreSuite) SetupTest() {
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

func (s *ProcessIndicatorsStoreSuite) TearDownTest() {
	if s.pool != nil {
		s.pool.Close()
	}
	s.envIsolator.RestoreAll()
}

func (s *ProcessIndicatorsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	processIndicator := &storage.ProcessIndicator{}
	s.NoError(testutils.FullInit(processIndicator, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundProcessIndicator, exists, err := store.Get(ctx, processIndicator.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundProcessIndicator)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, processIndicator))
	foundProcessIndicator, exists, err = store.Get(ctx, processIndicator.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(processIndicator, foundProcessIndicator)

	processIndicatorCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(1, processIndicatorCount)

	processIndicatorExists, err := store.Exists(ctx, processIndicator.GetId())
	s.NoError(err)
	s.True(processIndicatorExists)
	s.NoError(store.Upsert(ctx, processIndicator))
	s.ErrorIs(store.Upsert(withNoAccessCtx, processIndicator), sac.ErrResourceAccessDenied)

	foundProcessIndicator, exists, err = store.Get(ctx, processIndicator.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(processIndicator, foundProcessIndicator)

	s.NoError(store.Delete(ctx, processIndicator.GetId()))
	foundProcessIndicator, exists, err = store.Get(ctx, processIndicator.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundProcessIndicator)

	var processIndicators []*storage.ProcessIndicator
	for i := 0; i < 200; i++ {
		processIndicator := &storage.ProcessIndicator{}
		s.NoError(testutils.FullInit(processIndicator, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		processIndicators = append(processIndicators, processIndicator)
	}

	s.NoError(store.UpsertMany(ctx, processIndicators))

	processIndicatorCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(200, processIndicatorCount)
}

func (s *ProcessIndicatorsStoreSuite) TestSACUpsert() {
	obj := &storage.ProcessIndicator{}
	s.NoError(testutils.FullInit(obj, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	ctxs := getSACContexts(obj)
	for name, expectedErr := range map[string]error{
		withAllAccess:           nil,
		withNoAccess:            sac.ErrResourceAccessDenied,
		withNoAccessToCluster:   sac.ErrResourceAccessDenied,
		withAccessToDifferentNs: sac.ErrResourceAccessDenied,
		withAccess:              nil,
		withAccessToCluster:     nil,
	} {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			assert.ErrorIs(t, s.store.Upsert(ctxs[name], obj), expectedErr)
		})
	}
}

func (s *ProcessIndicatorsStoreSuite) TestSACUpsertMany() {
	obj := &storage.ProcessIndicator{}
	s.NoError(testutils.FullInit(obj, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	ctxs := getSACContexts(obj)
	for name, expectedErr := range map[string]error{
		withAllAccess:           nil,
		withNoAccess:            sac.ErrResourceAccessDenied,
		withNoAccessToCluster:   sac.ErrResourceAccessDenied,
		withAccessToDifferentNs: sac.ErrResourceAccessDenied,
		withAccess:              nil,
		withAccessToCluster:     nil,
	} {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			assert.ErrorIs(t, s.store.UpsertMany(ctxs[name], []*storage.ProcessIndicator{obj}), expectedErr)
		})
	}
}

const (
	withAllAccess           = "AllAccess"
	withNoAccess            = "NoAccess"
	withAccessToDifferentNs = "AccessToDifferentNs"
	withAccess              = "Access"
	withAccessToCluster     = "AccessToCluster"
	withNoAccessToCluster   = "NoAccessToCluster"
)

func getSACContexts(obj *storage.ProcessIndicator) map[string]context.Context {
	return map[string]context.Context{
		withAllAccess: sac.WithAllAccess(context.Background()),
		withNoAccess:  sac.WithNoAccess(context.Background()),
		withAccessToDifferentNs: sac.WithGlobalAccessScopeChecker(context.Background(),
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(targetResource),
				sac.ClusterScopeKeys(obj.GetClusterId()),
				sac.NamespaceScopeKeys("unknown ns"),
			)),
		withAccess: sac.WithGlobalAccessScopeChecker(context.Background(),
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(targetResource),
				sac.ClusterScopeKeys(obj.GetClusterId()),
				sac.NamespaceScopeKeys(obj.GetNamespace()),
			)),
		withAccessToCluster: sac.WithGlobalAccessScopeChecker(context.Background(),
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(targetResource),
				sac.ClusterScopeKeys(obj.GetClusterId()),
			)),
		withNoAccessToCluster: sac.WithGlobalAccessScopeChecker(context.Background(),
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(targetResource),
				sac.ClusterScopeKeys("unknown cluster"),
			)),
	}
}
