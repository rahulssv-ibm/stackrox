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

type CollectionsStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestCollectionsStore(t *testing.T) {
	suite.Run(t, new(CollectionsStoreSuite))
}

func (s *CollectionsStoreSuite) SetupSuite() {

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *CollectionsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE collections CASCADE")
	s.T().Log("collections", tag)
	s.store = New(s.testDB.DB)
	s.NoError(err)
}

func (s *CollectionsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *CollectionsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	resourceCollection := &storage.ResourceCollection{}
	s.NoError(testutils.FullInit(resourceCollection, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))
	resourceCollection.EmbeddedCollections = nil

	foundResourceCollection, exists, err := store.Get(ctx, resourceCollection.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundResourceCollection)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, resourceCollection))
	foundResourceCollection, exists, err = store.Get(ctx, resourceCollection.GetId())
	s.NoError(err)
	s.True(exists)
	protoassert.Equal(s.T(), resourceCollection, foundResourceCollection)

	resourceCollectionCount, err := store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(1, resourceCollectionCount)
	resourceCollectionCount, err = store.Count(withNoAccessCtx, search.EmptyQuery())
	s.NoError(err)
	s.Zero(resourceCollectionCount)

	resourceCollectionExists, err := store.Exists(ctx, resourceCollection.GetId())
	s.NoError(err)
	s.True(resourceCollectionExists)
	s.NoError(store.Upsert(ctx, resourceCollection))
	s.ErrorIs(store.Upsert(withNoAccessCtx, resourceCollection), sac.ErrResourceAccessDenied)

	s.NoError(store.Delete(ctx, resourceCollection.GetId()))
	foundResourceCollection, exists, err = store.Get(ctx, resourceCollection.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundResourceCollection)
	s.ErrorIs(store.Delete(withNoAccessCtx, resourceCollection.GetId()), sac.ErrResourceAccessDenied)

	var resourceCollections []*storage.ResourceCollection
	var resourceCollectionIDs []string
	for i := 0; i < 200; i++ {
		resourceCollection := &storage.ResourceCollection{}
		s.NoError(testutils.FullInit(resourceCollection, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		resourceCollection.EmbeddedCollections = nil
		resourceCollections = append(resourceCollections, resourceCollection)
		resourceCollectionIDs = append(resourceCollectionIDs, resourceCollection.GetId())
	}

	s.NoError(store.UpsertMany(ctx, resourceCollections))

	resourceCollectionCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(200, resourceCollectionCount)

	s.NoError(store.DeleteMany(ctx, resourceCollectionIDs))

	resourceCollectionCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(0, resourceCollectionCount)
}
