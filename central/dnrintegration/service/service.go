package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	clusterDataStore "github.com/stackrox/rox/central/cluster/datastore"
	deploymentDataStore "github.com/stackrox/rox/central/deployment/datastore"
	"github.com/stackrox/rox/central/dnrintegration/datastore"
	"github.com/stackrox/rox/central/enrichment"
	"github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/grpc"
	"github.com/stackrox/rox/pkg/logging"
)

var (
	log = logging.LoggerForModule()
)

// Service provides the interface to the microservice that serves alert data.
type Service interface {
	grpc.APIService

	AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error)

	GetDNRIntegration(ctx context.Context, req *v1.ResourceByID) (*v1.DNRIntegration, error)
	GetDNRIntegrations(ctx context.Context, req *v1.GetDNRIntegrationsRequest) (*v1.GetDNRIntegrationsResponse, error)
	PostDNRIntegration(ctx context.Context, req *v1.DNRIntegration) (*v1.DNRIntegration, error)
	PutDNRIntegration(ctx context.Context, req *v1.DNRIntegration) (*v1.DNRIntegration, error)
	DeleteDNRIntegration(ctx context.Context, req *v1.ResourceByID) (*empty.Empty, error)
	TestDNRIntegration(ctx context.Context, req *v1.DNRIntegration) (*empty.Empty, error)
}

// New returns a new Service instance using the given DataStore.
func New(datastore datastore.DataStore, clusters clusterDataStore.DataStore, deployments deploymentDataStore.DataStore, enricher enrichment.Enricher) Service {
	return &serviceImpl{
		datastore:   datastore,
		deployments: deployments,
		clusters:    clusters,
		enricher:    enricher,
	}
}
