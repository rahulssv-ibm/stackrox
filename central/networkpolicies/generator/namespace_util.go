package generator

import (
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/namespaces"
)

var allowAllNamespaces = &storage.LabelSelector{}

func createNamespacesByNameMap(namespaces []*storage.NamespaceMetadata) map[string]*storage.NamespaceMetadata {
	result := make(map[string]*storage.NamespaceMetadata, len(namespaces))

	for _, ns := range namespaces {
		result[ns.GetName()] = ns
	}
	return result
}

func labelSelectorForNamespace(ns *storage.NamespaceMetadata) *storage.LabelSelector {
	if ns == nil {
		return allowAllNamespaces
	}

	var matchLabels map[string]string

	nsLabels := ns.GetLabels()
	if nsLabels[namespaces.NamespaceNameLabel] == ns.GetName() {
		matchLabels = map[string]string{
			namespaces.NamespaceNameLabel: ns.GetName(),
		}
	} else {
		matchLabels = nsLabels
	}

	return &storage.LabelSelector{
		MatchLabels: matchLabels,
	}
}
