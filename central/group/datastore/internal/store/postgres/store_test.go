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

type GroupsStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestGroupsStore(t *testing.T) {
	suite.Run(t, new(GroupsStoreSuite))
}

func (s *GroupsStoreSuite) SetupSuite() {

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *GroupsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE groups CASCADE")
	s.T().Log("groups", tag)
	s.store = New(s.testDB.DB)
	s.NoError(err)
}

func (s *GroupsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *GroupsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	group := &storage.Group{}
	s.NoError(testutils.FullInit(group, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundGroup, exists, err := store.Get(ctx, group.GetProps().GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundGroup)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, group))
	foundGroup, exists, err = store.Get(ctx, group.GetProps().GetId())
	s.NoError(err)
	s.True(exists)
	protoassert.Equal(s.T(), group, foundGroup)

	groupCount, err := store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(1, groupCount)
	groupCount, err = store.Count(withNoAccessCtx, search.EmptyQuery())
	s.NoError(err)
	s.Zero(groupCount)

	groupExists, err := store.Exists(ctx, group.GetProps().GetId())
	s.NoError(err)
	s.True(groupExists)
	s.NoError(store.Upsert(ctx, group))
	s.ErrorIs(store.Upsert(withNoAccessCtx, group), sac.ErrResourceAccessDenied)

	s.NoError(store.Delete(ctx, group.GetProps().GetId()))
	foundGroup, exists, err = store.Get(ctx, group.GetProps().GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundGroup)
	s.ErrorIs(store.Delete(withNoAccessCtx, group.GetProps().GetId()), sac.ErrResourceAccessDenied)

	var groups []*storage.Group
	var groupIDs []string
	for i := 0; i < 200; i++ {
		group := &storage.Group{}
		s.NoError(testutils.FullInit(group, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		groups = append(groups, group)
		groupIDs = append(groupIDs, group.GetProps().GetId())
	}

	s.NoError(store.UpsertMany(ctx, groups))
	allGroup, err := store.GetAll(ctx)
	s.NoError(err)
	protoassert.ElementsMatch(s.T(), groups, allGroup)

	groupCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(200, groupCount)

	s.NoError(store.DeleteMany(ctx, groupIDs))

	groupCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(0, groupCount)
}
