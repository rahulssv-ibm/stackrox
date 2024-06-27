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

type BlobsStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestBlobsStore(t *testing.T) {
	suite.Run(t, new(BlobsStoreSuite))
}

func (s *BlobsStoreSuite) SetupSuite() {

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *BlobsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE blobs CASCADE")
	s.T().Log("blobs", tag)
	s.store = New(s.testDB.DB)
	s.NoError(err)
}

func (s *BlobsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *BlobsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	blob := &storage.Blob{}
	s.NoError(testutils.FullInit(blob, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundBlob, exists, err := store.Get(ctx, blob.GetName())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundBlob)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, blob))
	foundBlob, exists, err = store.Get(ctx, blob.GetName())
	s.NoError(err)
	s.True(exists)
	protoassert.Equal(s.T(), blob, foundBlob)

	blobCount, err := store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(1, blobCount)
	blobCount, err = store.Count(withNoAccessCtx, search.EmptyQuery())
	s.NoError(err)
	s.Zero(blobCount)

	blobExists, err := store.Exists(ctx, blob.GetName())
	s.NoError(err)
	s.True(blobExists)
	s.NoError(store.Upsert(ctx, blob))
	s.ErrorIs(store.Upsert(withNoAccessCtx, blob), sac.ErrResourceAccessDenied)

	foundBlob, exists, err = store.Get(ctx, blob.GetName())
	s.NoError(err)
	s.True(exists)
	protoassert.Equal(s.T(), blob, foundBlob)

	s.NoError(store.Delete(ctx, blob.GetName()))
	foundBlob, exists, err = store.Get(ctx, blob.GetName())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundBlob)
	s.ErrorIs(store.Delete(withNoAccessCtx, blob.GetName()), sac.ErrResourceAccessDenied)

	var blobs []*storage.Blob
	var blobIDs []string
	for i := 0; i < 200; i++ {
		blob := &storage.Blob{}
		s.NoError(testutils.FullInit(blob, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		blobs = append(blobs, blob)
		blobIDs = append(blobIDs, blob.GetName())
	}

	s.NoError(store.UpsertMany(ctx, blobs))

	blobCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(200, blobCount)

	s.NoError(store.DeleteMany(ctx, blobIDs))

	blobCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(0, blobCount)
}
