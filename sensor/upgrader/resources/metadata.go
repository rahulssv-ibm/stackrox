package resources

import (
	"strings"

	"github.com/pkg/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
)

// Metadata represents Kubernetes API resource metadata.
type Metadata struct {
	v1.APIResource
}

// GroupVersionKind returns the `schema.GroupVersionKind` of an API resource. The returned value is safe to be used
// in map keys etc.
func (m *Metadata) GroupVersionKind() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   m.Group,
		Version: m.Version,
		Kind:    m.Kind,
	}
}

// GetAvailableResources uses the Kubernetes Discovery API to list all relevant resources on the server.
func GetAvailableResources(client discovery.DiscoveryInterface, relevantGVKs []schema.GroupVersionKind) (map[schema.GroupVersionKind]*Metadata, error) {
	resourceLists, err := client.ServerResources()
	if err != nil {
		return nil, errors.Wrap(err, "retrieving list of server resources")
	}

	relevantGVKSet := make(map[schema.GroupVersionKind]struct{}, len(relevantGVKs))
	for _, gvk := range relevantGVKs {
		relevantGVKSet[gvk] = struct{}{}
	}

	result := make(map[schema.GroupVersionKind]*Metadata)

	for _, resourceList := range resourceLists {
		gv, err := schema.ParseGroupVersion(resourceList.GroupVersion)
		if err != nil {
			return nil, errors.Wrap(err, "parsing group/version of API resource list")
		}
		for _, apiResource := range resourceList.APIResources {
			if strings.ContainsRune(apiResource.Name, '/') {
				continue // ignore sub-resources like `deployments/scale`
			}
			if apiResource.Group == "" {
				apiResource.Group = gv.Group
			}
			if apiResource.Version == "" {
				apiResource.Version = gv.Version
			}

			gvk := schema.GroupVersionKind{
				Group:   apiResource.Group,
				Version: apiResource.Version,
				Kind:    apiResource.Kind,
			}

			if _, relevant := relevantGVKSet[gvk]; !relevant {
				continue
			}

			md := &Metadata{
				APIResource: apiResource,
			}
			result[gvk] = md
		}
	}

	return result, nil
}
