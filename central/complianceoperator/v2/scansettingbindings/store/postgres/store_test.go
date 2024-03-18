// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ComplianceOperatorScanSettingBindingV2StoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestComplianceOperatorScanSettingBindingV2Store(t *testing.T) {
	suite.Run(t, new(ComplianceOperatorScanSettingBindingV2StoreSuite))
}

func (s *ComplianceOperatorScanSettingBindingV2StoreSuite) SetupSuite() {

	s.T().Setenv(features.ComplianceEnhancements.EnvVar(), "true")
	if !features.ComplianceEnhancements.Enabled() {
		s.T().Skip("Skip postgres store tests because feature flag is off")
		s.T().SkipNow()
	}

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *ComplianceOperatorScanSettingBindingV2StoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE compliance_operator_scan_setting_binding_v2 CASCADE")
	s.T().Log("compliance_operator_scan_setting_binding_v2", tag)
	s.store = New(s.testDB.DB)
	s.NoError(err)
}

func (s *ComplianceOperatorScanSettingBindingV2StoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *ComplianceOperatorScanSettingBindingV2StoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	complianceOperatorScanSettingBindingV2 := &storage.ComplianceOperatorScanSettingBindingV2{}
	s.NoError(testutils.FullInit(complianceOperatorScanSettingBindingV2, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundComplianceOperatorScanSettingBindingV2, exists, err := store.Get(ctx, complianceOperatorScanSettingBindingV2.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComplianceOperatorScanSettingBindingV2)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, complianceOperatorScanSettingBindingV2))
	foundComplianceOperatorScanSettingBindingV2, exists, err = store.Get(ctx, complianceOperatorScanSettingBindingV2.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(complianceOperatorScanSettingBindingV2, foundComplianceOperatorScanSettingBindingV2)

	complianceOperatorScanSettingBindingV2Count, err := store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(1, complianceOperatorScanSettingBindingV2Count)
	complianceOperatorScanSettingBindingV2Count, err = store.Count(withNoAccessCtx, search.EmptyQuery())
	s.NoError(err)
	s.Zero(complianceOperatorScanSettingBindingV2Count)

	complianceOperatorScanSettingBindingV2Exists, err := store.Exists(ctx, complianceOperatorScanSettingBindingV2.GetId())
	s.NoError(err)
	s.True(complianceOperatorScanSettingBindingV2Exists)
	s.NoError(store.Upsert(ctx, complianceOperatorScanSettingBindingV2))
	s.ErrorIs(store.Upsert(withNoAccessCtx, complianceOperatorScanSettingBindingV2), sac.ErrResourceAccessDenied)

	foundComplianceOperatorScanSettingBindingV2, exists, err = store.Get(ctx, complianceOperatorScanSettingBindingV2.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(complianceOperatorScanSettingBindingV2, foundComplianceOperatorScanSettingBindingV2)

	s.NoError(store.Delete(ctx, complianceOperatorScanSettingBindingV2.GetId()))
	foundComplianceOperatorScanSettingBindingV2, exists, err = store.Get(ctx, complianceOperatorScanSettingBindingV2.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComplianceOperatorScanSettingBindingV2)
	s.NoError(store.Delete(withNoAccessCtx, complianceOperatorScanSettingBindingV2.GetId()))

	var complianceOperatorScanSettingBindingV2s []*storage.ComplianceOperatorScanSettingBindingV2
	var complianceOperatorScanSettingBindingV2IDs []string
	for i := 0; i < 200; i++ {
		complianceOperatorScanSettingBindingV2 := &storage.ComplianceOperatorScanSettingBindingV2{}
		s.NoError(testutils.FullInit(complianceOperatorScanSettingBindingV2, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		complianceOperatorScanSettingBindingV2s = append(complianceOperatorScanSettingBindingV2s, complianceOperatorScanSettingBindingV2)
		complianceOperatorScanSettingBindingV2IDs = append(complianceOperatorScanSettingBindingV2IDs, complianceOperatorScanSettingBindingV2.GetId())
	}

	s.NoError(store.UpsertMany(ctx, complianceOperatorScanSettingBindingV2s))

	complianceOperatorScanSettingBindingV2Count, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(200, complianceOperatorScanSettingBindingV2Count)

	s.NoError(store.DeleteMany(ctx, complianceOperatorScanSettingBindingV2IDs))

	complianceOperatorScanSettingBindingV2Count, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(0, complianceOperatorScanSettingBindingV2Count)
}

const (
	withAllAccess                = "AllAccess"
	withNoAccess                 = "NoAccess"
	withAccess                   = "Access"
	withAccessToCluster          = "AccessToCluster"
	withNoAccessToCluster        = "NoAccessToCluster"
	withAccessToDifferentCluster = "AccessToDifferentCluster"
	withAccessToDifferentNs      = "AccessToDifferentNs"
)

var (
	withAllAccessCtx = sac.WithAllAccess(context.Background())
)

type testCase struct {
	context                context.Context
	expectedObjIDs         []string
	expectedIdentifiers    []string
	expectedMissingIndices []int
	expectedObjects        []*storage.ComplianceOperatorScanSettingBindingV2
	expectedWriteError     error
}

func (s *ComplianceOperatorScanSettingBindingV2StoreSuite) getTestData(access ...storage.Access) (*storage.ComplianceOperatorScanSettingBindingV2, *storage.ComplianceOperatorScanSettingBindingV2, map[string]testCase) {
	objA := &storage.ComplianceOperatorScanSettingBindingV2{}
	s.NoError(testutils.FullInit(objA, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))

	objB := &storage.ComplianceOperatorScanSettingBindingV2{}
	s.NoError(testutils.FullInit(objB, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))

	testCases := map[string]testCase{
		withAllAccess: {
			context:                sac.WithAllAccess(context.Background()),
			expectedObjIDs:         []string{objA.GetId(), objB.GetId()},
			expectedIdentifiers:    []string{objA.GetId(), objB.GetId()},
			expectedMissingIndices: []int{},
			expectedObjects:        []*storage.ComplianceOperatorScanSettingBindingV2{objA, objB},
			expectedWriteError:     nil,
		},
		withNoAccess: {
			context:                sac.WithNoAccess(context.Background()),
			expectedObjIDs:         []string{},
			expectedIdentifiers:    []string{},
			expectedMissingIndices: []int{0, 1},
			expectedObjects:        []*storage.ComplianceOperatorScanSettingBindingV2{},
			expectedWriteError:     sac.ErrResourceAccessDenied,
		},
		withNoAccessToCluster: {
			context: sac.WithGlobalAccessScopeChecker(context.Background(),
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(access...),
					sac.ResourceScopeKeys(targetResource),
					sac.ClusterScopeKeys(uuid.Nil.String()),
				)),
			expectedObjIDs:         []string{},
			expectedIdentifiers:    []string{},
			expectedMissingIndices: []int{0, 1},
			expectedObjects:        []*storage.ComplianceOperatorScanSettingBindingV2{},
			expectedWriteError:     sac.ErrResourceAccessDenied,
		},
		withAccess: {
			context: sac.WithGlobalAccessScopeChecker(context.Background(),
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(access...),
					sac.ResourceScopeKeys(targetResource),
					sac.ClusterScopeKeys(objA.GetClusterId()),
				)),
			expectedObjIDs:         []string{objA.GetId()},
			expectedIdentifiers:    []string{objA.GetId()},
			expectedMissingIndices: []int{1},
			expectedObjects:        []*storage.ComplianceOperatorScanSettingBindingV2{objA},
			expectedWriteError:     nil,
		},
		withAccessToCluster: {
			context: sac.WithGlobalAccessScopeChecker(context.Background(),
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(access...),
					sac.ResourceScopeKeys(targetResource),
					sac.ClusterScopeKeys(objA.GetClusterId()),
				)),
			expectedObjIDs:         []string{objA.GetId()},
			expectedIdentifiers:    []string{objA.GetId()},
			expectedMissingIndices: []int{1},
			expectedObjects:        []*storage.ComplianceOperatorScanSettingBindingV2{objA},
			expectedWriteError:     nil,
		},
		withAccessToDifferentCluster: {
			context: sac.WithGlobalAccessScopeChecker(context.Background(),
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(access...),
					sac.ResourceScopeKeys(targetResource),
					sac.ClusterScopeKeys("caaaaaaa-bbbb-4011-0000-111111111111"),
				)),
			expectedObjIDs:         []string{},
			expectedIdentifiers:    []string{},
			expectedMissingIndices: []int{0, 1},
			expectedObjects:        []*storage.ComplianceOperatorScanSettingBindingV2{},
			expectedWriteError:     sac.ErrResourceAccessDenied,
		},
		withAccessToDifferentNs: {
			context: sac.WithGlobalAccessScopeChecker(context.Background(),
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(access...),
					sac.ResourceScopeKeys(targetResource),
					sac.ClusterScopeKeys(objA.GetClusterId()),
					sac.NamespaceScopeKeys("unknown ns"),
				)),
			expectedObjIDs:         []string{objA.GetId()},
			expectedIdentifiers:    []string{objA.GetId()},
			expectedMissingIndices: []int{1},
			expectedObjects:        []*storage.ComplianceOperatorScanSettingBindingV2{objA},
			expectedWriteError:     nil,
		},
	}

	return objA, objB, testCases
}

func (s *ComplianceOperatorScanSettingBindingV2StoreSuite) TestSACUpsert() {
	obj, _, testCases := s.getTestData(storage.Access_READ_WRITE_ACCESS)
	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			assert.ErrorIs(t, s.store.Upsert(testCase.context, obj), testCase.expectedWriteError)
		})
	}
}

func (s *ComplianceOperatorScanSettingBindingV2StoreSuite) TestSACUpsertMany() {
	obj, _, testCases := s.getTestData(storage.Access_READ_WRITE_ACCESS)
	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			assert.ErrorIs(t, s.store.UpsertMany(testCase.context, []*storage.ComplianceOperatorScanSettingBindingV2{obj}), testCase.expectedWriteError)
		})
	}
}

func (s *ComplianceOperatorScanSettingBindingV2StoreSuite) TestSACCount() {
	objA, objB, testCases := s.getTestData(storage.Access_READ_ACCESS)
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objB))

	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			expectedCount := len(testCase.expectedObjects)
			count, err := s.store.Count(testCase.context, search.EmptyQuery())
			assert.NoError(t, err)
			assert.Equal(t, expectedCount, count)
		})
	}
}

func (s *ComplianceOperatorScanSettingBindingV2StoreSuite) TestSACWalk() {
	objA, objB, testCases := s.getTestData(storage.Access_READ_ACCESS)
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objB))

	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			identifiers := []string{}
			getIDs := func(obj *storage.ComplianceOperatorScanSettingBindingV2) error {
				identifiers = append(identifiers, obj.GetId())
				return nil
			}
			err := s.store.Walk(testCase.context, getIDs)
			assert.NoError(t, err)
			assert.ElementsMatch(t, testCase.expectedIdentifiers, identifiers)
		})
	}
}

func (s *ComplianceOperatorScanSettingBindingV2StoreSuite) TestSACGetIDs() {
	objA, objB, testCases := s.getTestData(storage.Access_READ_ACCESS)
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objB))

	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			identifiers, err := s.store.GetIDs(testCase.context)
			assert.NoError(t, err)
			assert.ElementsMatch(t, testCase.expectedObjIDs, identifiers)
		})
	}
}

func (s *ComplianceOperatorScanSettingBindingV2StoreSuite) TestSACExists() {
	objA, _, testCases := s.getTestData(storage.Access_READ_ACCESS)
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))

	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			exists, err := s.store.Exists(testCase.context, objA.GetId())
			assert.NoError(t, err)

			// Assumption from the test case structure: objA is always in the visible list
			// in the first position.
			expectedFound := len(testCase.expectedObjects) > 0
			assert.Equal(t, expectedFound, exists)
		})
	}
}

func (s *ComplianceOperatorScanSettingBindingV2StoreSuite) TestSACGet() {
	objA, _, testCases := s.getTestData(storage.Access_READ_ACCESS)
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))

	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			actual, exists, err := s.store.Get(testCase.context, objA.GetId())
			assert.NoError(t, err)

			// Assumption from the test case structure: objA is always in the visible list
			// in the first position.
			expectedFound := len(testCase.expectedObjects) > 0
			assert.Equal(t, expectedFound, exists)
			if expectedFound {
				assert.Equal(t, objA, actual)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}

func (s *ComplianceOperatorScanSettingBindingV2StoreSuite) TestSACDelete() {
	objA, objB, testCases := s.getTestData(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS)

	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			s.SetupTest()

			s.NoError(s.store.Upsert(withAllAccessCtx, objA))
			s.NoError(s.store.Upsert(withAllAccessCtx, objB))

			assert.NoError(t, s.store.Delete(testCase.context, objA.GetId()))
			assert.NoError(t, s.store.Delete(testCase.context, objB.GetId()))

			count, err := s.store.Count(withAllAccessCtx, search.EmptyQuery())
			assert.NoError(t, err)
			assert.Equal(t, 2-len(testCase.expectedObjects), count)

			// Ensure objects allowed by test scope were actually deleted
			for _, obj := range testCase.expectedObjects {
				found, err := s.store.Exists(withAllAccessCtx, obj.GetId())
				assert.NoError(t, err)
				assert.False(t, found)
			}
		})
	}
}

func (s *ComplianceOperatorScanSettingBindingV2StoreSuite) TestSACDeleteMany() {
	objA, objB, testCases := s.getTestData(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS)
	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			s.SetupTest()

			s.NoError(s.store.Upsert(withAllAccessCtx, objA))
			s.NoError(s.store.Upsert(withAllAccessCtx, objB))

			assert.NoError(t, s.store.DeleteMany(testCase.context, []string{
				objA.GetId(),
				objB.GetId(),
			}))

			count, err := s.store.Count(withAllAccessCtx, search.EmptyQuery())
			assert.NoError(t, err)
			assert.Equal(t, 2-len(testCase.expectedObjects), count)

			// Ensure objects allowed by test scope were actually deleted
			for _, obj := range testCase.expectedObjects {
				found, err := s.store.Exists(withAllAccessCtx, obj.GetId())
				assert.NoError(t, err)
				assert.False(t, found)
			}
		})
	}
}

func (s *ComplianceOperatorScanSettingBindingV2StoreSuite) TestSACGetMany() {
	objA, objB, testCases := s.getTestData(storage.Access_READ_ACCESS)
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objB))

	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			actual, missingIndices, err := s.store.GetMany(testCase.context, []string{objA.GetId(), objB.GetId()})
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedObjects, actual)
			assert.Equal(t, testCase.expectedMissingIndices, missingIndices)
		})
	}

	s.T().Run("with no identifiers", func(t *testing.T) {
		actual, missingIndices, err := s.store.GetMany(withAllAccessCtx, []string{})
		assert.Nil(t, err)
		assert.Nil(t, actual)
		assert.Nil(t, missingIndices)
	})
}
