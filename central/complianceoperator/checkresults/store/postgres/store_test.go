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

type ComplianceOperatorCheckResultsStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestComplianceOperatorCheckResultsStore(t *testing.T) {
	suite.Run(t, new(ComplianceOperatorCheckResultsStoreSuite))
}

func (s *ComplianceOperatorCheckResultsStoreSuite) SetupSuite() {

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *ComplianceOperatorCheckResultsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE compliance_operator_check_results CASCADE")
	s.T().Log("compliance_operator_check_results", tag)
	s.store = New(s.testDB.DB)
	s.NoError(err)
}

func (s *ComplianceOperatorCheckResultsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *ComplianceOperatorCheckResultsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	complianceOperatorCheckResult := &storage.ComplianceOperatorCheckResult{}
	s.NoError(testutils.FullInit(complianceOperatorCheckResult, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundComplianceOperatorCheckResult, exists, err := store.Get(ctx, complianceOperatorCheckResult.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComplianceOperatorCheckResult)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, complianceOperatorCheckResult))
	foundComplianceOperatorCheckResult, exists, err = store.Get(ctx, complianceOperatorCheckResult.GetId())
	s.NoError(err)
	s.True(exists)
	protoassert.Equal(s.T(), complianceOperatorCheckResult, foundComplianceOperatorCheckResult)

	complianceOperatorCheckResultCount, err := store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(1, complianceOperatorCheckResultCount)
	complianceOperatorCheckResultCount, err = store.Count(withNoAccessCtx, search.EmptyQuery())
	s.NoError(err)
	s.Zero(complianceOperatorCheckResultCount)

	complianceOperatorCheckResultExists, err := store.Exists(ctx, complianceOperatorCheckResult.GetId())
	s.NoError(err)
	s.True(complianceOperatorCheckResultExists)
	s.NoError(store.Upsert(ctx, complianceOperatorCheckResult))
	s.ErrorIs(store.Upsert(withNoAccessCtx, complianceOperatorCheckResult), sac.ErrResourceAccessDenied)

	foundComplianceOperatorCheckResult, exists, err = store.Get(ctx, complianceOperatorCheckResult.GetId())
	s.NoError(err)
	s.True(exists)
	protoassert.Equal(s.T(), complianceOperatorCheckResult, foundComplianceOperatorCheckResult)

	s.NoError(store.Delete(ctx, complianceOperatorCheckResult.GetId()))
	foundComplianceOperatorCheckResult, exists, err = store.Get(ctx, complianceOperatorCheckResult.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComplianceOperatorCheckResult)
	s.ErrorIs(store.Delete(withNoAccessCtx, complianceOperatorCheckResult.GetId()), sac.ErrResourceAccessDenied)

	var complianceOperatorCheckResults []*storage.ComplianceOperatorCheckResult
	var complianceOperatorCheckResultIDs []string
	for i := 0; i < 200; i++ {
		complianceOperatorCheckResult := &storage.ComplianceOperatorCheckResult{}
		s.NoError(testutils.FullInit(complianceOperatorCheckResult, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		complianceOperatorCheckResults = append(complianceOperatorCheckResults, complianceOperatorCheckResult)
		complianceOperatorCheckResultIDs = append(complianceOperatorCheckResultIDs, complianceOperatorCheckResult.GetId())
	}

	s.NoError(store.UpsertMany(ctx, complianceOperatorCheckResults))

	complianceOperatorCheckResultCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(200, complianceOperatorCheckResultCount)

	s.NoError(store.DeleteMany(ctx, complianceOperatorCheckResultIDs))

	complianceOperatorCheckResultCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(0, complianceOperatorCheckResultCount)
}
