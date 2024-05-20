package service

import (
	"github.com/stackrox/rox/central/complianceoperator/v2/compliancemanager"
	reportManager "github.com/stackrox/rox/central/complianceoperator/v2/report/manager"
	scanSettingsDS "github.com/stackrox/rox/central/complianceoperator/v2/scanconfigurations/datastore"
	scanSettingBindingsDS "github.com/stackrox/rox/central/complianceoperator/v2/scansettingbindings/datastore"
	suiteDS "github.com/stackrox/rox/central/complianceoperator/v2/suites/datastore"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	serviceInstance     Service
	serviceInstanceInit sync.Once
)

// Singleton returns the singleton instance of the compliance service.
func Singleton() Service {
	if !features.ComplianceEnhancements.Enabled() {
		return nil
	}

	if features.ComplianceReporting.Enabled() {
		go reportManager.Singleton().Start()
	}

	serviceInstanceInit.Do(func() {
		serviceInstance = New(scanSettingsDS.Singleton(), scanSettingBindingsDS.Singleton(), suiteDS.Singleton(), compliancemanager.Singleton())
	})
	return serviceInstance
}
