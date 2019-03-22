package standards

import (
	"github.com/stackrox/rox/central/compliance/framework"
	"github.com/stackrox/rox/central/compliance/standards/index"
	"github.com/stackrox/rox/central/compliance/standards/metadata"
	"github.com/stackrox/rox/central/globalindex"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/pkg/utils"
)

var (
	registryInstance     *Registry
	registryInstanceInit sync.Once
)

// RegistrySingleton returns the singleton instance of the compliance standards Registry.
func RegistrySingleton() *Registry {
	registryInstanceInit.Do(func() {
		indexer := index.New(globalindex.GetGlobalIndex())
		registryInstance = NewRegistry(indexer, framework.RegistrySingleton())
		utils.Must(registryInstance.RegisterStandards(metadata.AllStandards...))
	})
	return registryInstance
}
