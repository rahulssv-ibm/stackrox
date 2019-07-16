package datastore

import (
	"context"

	"github.com/stackrox/rox/central/compliance"
	"github.com/stackrox/rox/central/compliance/datastore/types"
	"github.com/stackrox/rox/central/role/resources"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/errorhelpers"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/set"
)

var (
	clusterSAC     = sac.ForResource(resources.Cluster)
	deploymentsSAC = sac.ForResource(resources.Deployment)
	nodeSAC        = sac.ForResource(resources.Node)
)

// SacFilter provides the filtering abilities needed by the compliance datastore.
//go:generate mockgen-wrapper SacFilter
type SacFilter interface {
	FilterRunResults(ctx context.Context, results *storage.ComplianceRunResults) (*storage.ComplianceRunResults, error)
	FilterBatchResults(ctx context.Context, results map[compliance.ClusterStandardPair]types.ResultsWithStatus) (map[compliance.ClusterStandardPair]types.ResultsWithStatus, error)
}

// NewSacFilter returns a new instance of a SacFilter using the input deployment datastore.
func NewSacFilter() SacFilter {
	return &sacFilterImpl{}
}

type sacFilterImpl struct{}

// FilterRunResults filters the deployments and nodes contained in a single ComplianceRunResults to only those that
// the input context has access to.
func (ds *sacFilterImpl) FilterRunResults(ctx context.Context, runResults *storage.ComplianceRunResults) (*storage.ComplianceRunResults, error) {
	filteredDomain, filtered, err := ds.filterDomain(ctx, runResults.Domain)
	if err != nil {
		return nil, err
	}
	if !filtered {
		return runResults, nil
	}

	filteredResults := &storage.ComplianceRunResults{
		Domain:      filteredDomain,
		RunMetadata: runResults.GetRunMetadata(),
	}
	if filteredDomain.GetCluster() != nil {
		filteredResults.ClusterResults = runResults.GetClusterResults()
	}
	if len(filteredDomain.GetNodes()) > 0 {
		filteredResults.NodeResults = runResults.GetNodeResults()
	}
	if len(filteredDomain.GetDeployments()) > 0 {
		if len(filteredResults.GetDeploymentResults()) == len(runResults.GetDomain().GetDeployments()) {
			filteredResults.DeploymentResults = runResults.GetDeploymentResults()
		} else {
			filteredResults.DeploymentResults = make(map[string]*storage.ComplianceRunResults_EntityResults)
			for deploymentID := range filteredDomain.GetDeployments() {
				filteredResults.DeploymentResults[deploymentID] = runResults.GetDeploymentResults()[deploymentID]
			}
		}
	}
	return filteredResults, nil
}

// FilterBatchResults returns a new results map, removing results for the cluster, deployments, and nodes that the input
// context does not have access to.
func (ds *sacFilterImpl) FilterBatchResults(ctx context.Context, batchResults map[compliance.ClusterStandardPair]types.ResultsWithStatus) (map[compliance.ClusterStandardPair]types.ResultsWithStatus, error) {
	clusterIDs := set.NewStringSet()
	for pair := range batchResults {
		clusterIDs.Add(pair.ClusterID)
	}
	allowedClusters, err := ds.filterClusters(ctx, clusterIDs)
	if err != nil {
		return nil, err
	}

	// Create a new map with only the allowed results.
	allowedMap := make(map[compliance.ClusterStandardPair]types.ResultsWithStatus, len(batchResults))
	for pair, batchResult := range batchResults {
		if !allowedClusters.Contains(pair.ClusterID) {
			continue
		}

		// Get and filter the results for the pair.
		batchResult.LastSuccessfulResults, err = ds.FilterRunResults(ctx, batchResult.LastSuccessfulResults)
		if err != nil {
			return nil, err
		}

		// Add the results to filtered returned map.
		allowedMap[pair] = batchResult
	}
	return allowedMap, nil
}

// Helper functions that filter objects.
////////////////////////////////////////

func (ds *sacFilterImpl) filterClusters(ctx context.Context, clusters set.StringSet) (set.StringSet, error) {
	resourceScopeChecker := sac.GlobalAccessScopeChecker(ctx).AccessMode(storage.Access_READ_ACCESS).Resource(resources.Compliance)

	// Filter the compliance results by cluster.
	allowed, maybe := ds.tryFilterClusters(resourceScopeChecker, clusters.AsSlice())
	if maybe.Cardinality() > 0 {
		if err := resourceScopeChecker.PerformChecks(ctx); err != nil {
			return set.StringSet{}, err
		}
		extraAllowed, maybe := ds.tryFilterClusters(resourceScopeChecker, maybe.AsSlice())
		if maybe.Cardinality() > 0 {
			errorhelpers.PanicOnDevelopmentf("still %d maybe results after PerformChecks", maybe.Cardinality())
		}
		allowed.Union(extraAllowed)
	}
	return allowed, nil
}

func (ds *sacFilterImpl) tryFilterClusters(resourceScopeChecker sac.ScopeChecker, clusters []string) (set.StringSet, set.StringSet) {
	allowed := set.NewStringSet()
	maybe := set.NewStringSet()
	for _, cluster := range clusters {
		if res := resourceScopeChecker.TryAllowed(sac.ClusterScopeKey(cluster)); res == sac.Allow {
			allowed.Add(cluster)
		} else if res == sac.Unknown {
			maybe.Add(cluster)
		}
	}
	return allowed, maybe
}

func (ds *sacFilterImpl) filterDomain(ctx context.Context, domain *storage.ComplianceDomain) (*storage.ComplianceDomain, bool, error) {
	var filtered bool
	newDomain := &storage.ComplianceDomain{}

	ok, err := clusterSAC.ReadAllowed(ctx, sac.ClusterScopeKey(domain.Cluster.Id))
	if err != nil {
		return nil, false, err
	} else if !ok {
		filtered = true
	} else {
		newDomain.Cluster = domain.Cluster
	}

	ok, err = nodeSAC.ReadAllowed(ctx, sac.ClusterScopeKey(domain.Cluster.Id))
	if err != nil {
		return nil, false, err
	} else if !ok {
		filtered = true
	} else {
		newDomain.Nodes = domain.Nodes
	}

	ok, err = deploymentsSAC.ReadAllowed(ctx, sac.ClusterScopeKey(domain.Cluster.Id))
	if err != nil {
		return nil, false, err
	} else if !ok {
		filtered = true
		newDomain.Deployments, err = ds.filterDeploymentsFromDomain(ctx, domain.Deployments)
		if err != nil {
			return nil, false, err
		}
	} else {
		newDomain.Deployments = domain.Deployments
	}

	return newDomain, filtered, nil
}

func (ds *sacFilterImpl) filterDeploymentsFromDomain(ctx context.Context, deps map[string]*storage.Deployment) (map[string]*storage.Deployment, error) {
	newDeps := make(map[string]*storage.Deployment)
	for depID, dep := range deps {
		if ok, err := deploymentsSAC.ReadAllowed(ctx, sac.KeyForNSScopedObj(dep)...); err != nil {
			return nil, err
		} else if ok {
			newDeps[depID] = dep
		}
	}
	if len(newDeps) == 0 {
		return nil, nil
	}
	return newDeps, nil
}
