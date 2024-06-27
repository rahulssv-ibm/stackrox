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

type SignatureIntegrationsStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestSignatureIntegrationsStore(t *testing.T) {
	suite.Run(t, new(SignatureIntegrationsStoreSuite))
}

func (s *SignatureIntegrationsStoreSuite) SetupSuite() {

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *SignatureIntegrationsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE signature_integrations CASCADE")
	s.T().Log("signature_integrations", tag)
	s.store = New(s.testDB.DB)
	s.NoError(err)
}

func (s *SignatureIntegrationsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *SignatureIntegrationsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	signatureIntegration := &storage.SignatureIntegration{}
	s.NoError(testutils.FullInit(signatureIntegration, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundSignatureIntegration, exists, err := store.Get(ctx, signatureIntegration.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundSignatureIntegration)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, signatureIntegration))
	foundSignatureIntegration, exists, err = store.Get(ctx, signatureIntegration.GetId())
	s.NoError(err)
	s.True(exists)
	protoassert.Equal(s.T(), signatureIntegration, foundSignatureIntegration)

	signatureIntegrationCount, err := store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(1, signatureIntegrationCount)
	signatureIntegrationCount, err = store.Count(withNoAccessCtx, search.EmptyQuery())
	s.NoError(err)
	s.Zero(signatureIntegrationCount)

	signatureIntegrationExists, err := store.Exists(ctx, signatureIntegration.GetId())
	s.NoError(err)
	s.True(signatureIntegrationExists)
	s.NoError(store.Upsert(ctx, signatureIntegration))
	s.ErrorIs(store.Upsert(withNoAccessCtx, signatureIntegration), sac.ErrResourceAccessDenied)

	foundSignatureIntegration, exists, err = store.Get(ctx, signatureIntegration.GetId())
	s.NoError(err)
	s.True(exists)
	protoassert.Equal(s.T(), signatureIntegration, foundSignatureIntegration)

	s.NoError(store.Delete(ctx, signatureIntegration.GetId()))
	foundSignatureIntegration, exists, err = store.Get(ctx, signatureIntegration.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundSignatureIntegration)
	s.ErrorIs(store.Delete(withNoAccessCtx, signatureIntegration.GetId()), sac.ErrResourceAccessDenied)

	var signatureIntegrations []*storage.SignatureIntegration
	var signatureIntegrationIDs []string
	for i := 0; i < 200; i++ {
		signatureIntegration := &storage.SignatureIntegration{}
		s.NoError(testutils.FullInit(signatureIntegration, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		signatureIntegrations = append(signatureIntegrations, signatureIntegration)
		signatureIntegrationIDs = append(signatureIntegrationIDs, signatureIntegration.GetId())
	}

	s.NoError(store.UpsertMany(ctx, signatureIntegrations))

	signatureIntegrationCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(200, signatureIntegrationCount)

	s.NoError(store.DeleteMany(ctx, signatureIntegrationIDs))

	signatureIntegrationCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(0, signatureIntegrationCount)
}
