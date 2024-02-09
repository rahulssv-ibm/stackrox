package gatherers

import (
	"context"
	stdErrors "errors"
	"strings"

	"github.com/pkg/errors"
	depDatastore "github.com/stackrox/rox/central/deployment/datastore"
	nsDatastore "github.com/stackrox/rox/central/namespace/datastore"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/telemetry"
	"github.com/stackrox/rox/pkg/telemetry/data"
)

type namespaceGatherer struct {
	namespaceDatastore  nsDatastore.DataStore
	deploymentDatastore depDatastore.DataStore
}

func newNamespaceGatherer(namespaceDatastore nsDatastore.DataStore, deploymentDatastore depDatastore.DataStore) *namespaceGatherer {
	return &namespaceGatherer{
		namespaceDatastore:  namespaceDatastore,
		deploymentDatastore: deploymentDatastore,
	}
}

// Gather returns a list of stats about all namespaces in a cluster
func (n *namespaceGatherer) Gather(ctx context.Context, clusterID string) ([]*data.NamespaceInfo, []string) {
	namespaces, err := n.namespaceDatastore.SearchNamespaces(ctx, search.NewQueryBuilder().AddExactMatches(search.ClusterID, clusterID).ProtoQuery())
	if err != nil {
		return nil, []string{errors.Wrapf(err, "unable to load namespaces for cluster %s", clusterID).Error()}
	}
	namespaceList := make([]*data.NamespaceInfo, 0, len(namespaces))
	var searchErrs error
	for _, namespace := range namespaces {
		name := namespace.GetName()
		if !telemetry.WellKnownNamespaces.Contains(name) {
			name = ""
		}
		deployments, err := n.deploymentDatastore.Search(ctx, search.NewQueryBuilder().AddExactMatches(search.NamespaceID, namespace.GetId()).ProtoQuery())
		if err != nil {
			searchErrs = stdErrors.Join(searchErrs, errors.Wrapf(err, "unable to load deployments for namespace %s", namespace.GetName()))
			continue
		}
		namespaceList = append(namespaceList, &data.NamespaceInfo{
			ID:             namespace.GetId(),
			Name:           name,
			NumDeployments: len(deployments),
			// TODO: Fill out churn metrics once they are implemented
		})
	}
	var errList []string
	if searchErrs != nil {
		errList = strings.Split(searchErrs.Error(), "\n")
	}
	return namespaceList, errList
}
