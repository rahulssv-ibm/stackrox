package containerruntime

import (
	"github.com/stackrox/rox/benchmarks/checks/utils"
	"github.com/stackrox/rox/generated/api/v1"
)

type appArmorBenchmark struct{}

func (c *appArmorBenchmark) Definition() utils.Definition {
	return utils.Definition{
		BenchmarkCheckDefinition: v1.BenchmarkCheckDefinition{
			Name:        "CIS Docker v1.1.0 - 5.1",
			Description: "Ensure AppArmor Profile is Enabled",
		}, Dependencies: []utils.Dependency{utils.InitContainers},
	}
}

func (c *appArmorBenchmark) Run() (result v1.BenchmarkCheckResult) {
	utils.Pass(&result)
	for _, container := range utils.ContainersRunning {
		if container.AppArmorProfile == "unconfined" {
			utils.Warn(&result)
			utils.AddNotef(&result, "Container '%v' (%v) does not have app armor configured", container.ID, container.Name)
		}
	}
	return
}

// NewAppArmorBenchmark implements CIS-5.1
func NewAppArmorBenchmark() utils.Check {
	return &appArmorBenchmark{}
}
