package main

import (
	"fmt"

	"github.com/stackrox/rox/generated/storage"
)

var typeRegistry = make(map[string]unmarshaler)

type unmarshaler interface {
	UnmarshalVT([]byte) error
}

func init() {
	for _, v := range []unmarshaler{
		&storage.ActiveComponent{},
		&storage.AdministrationEvent{},
		&storage.AuthMachineToMachineConfig{},
		&storage.AuthProvider{},
		&storage.Blob{},
		&storage.CloudSource{},
		&storage.ClusterHealthStatus{},
		&storage.ClusterCVE{},
		&storage.ClusterCVEEdge{},
		&storage.ComplianceConfig{},
		&storage.ComplianceControlResult{},
		&storage.ComplianceDomain{},
		&storage.ComplianceIntegration{},
		&storage.ComplianceOperatorCheckResult{},
		&storage.ComplianceOperatorCheckResultV2{},
		&storage.ComplianceOperatorClusterScanConfigStatus{},
		&storage.ComplianceOperatorProfile{},
		&storage.ComplianceOperatorProfileV2{},
		&storage.ComplianceOperatorRule{},
		&storage.ComplianceOperatorRuleV2{},
		&storage.ComplianceOperatorScan{},
		&storage.ComplianceOperatorScanConfigurationV2{},
		&storage.ComplianceOperatorScanV2{},
		&storage.ComplianceOperatorScanSettingBinding{},
		&storage.ComplianceOperatorSuiteV2{},
		&storage.ComplianceRunMetadata{},
		&storage.ComplianceRunResults{},
		&storage.ComplianceStrings{},
		&storage.ComponentCVEEdge{},
		&storage.Config{},
		&storage.DeclarativeConfigHealth{},
		&storage.DelegatedRegistryConfig{},
		&storage.DiscoveredCluster{},
		&storage.ExternalBackup{},
		&storage.Group{},
		&storage.Hash{},
		&storage.ImageComponent{},
		&storage.ImageComponentEdge{},
		&storage.ImageCVE{},
		&storage.ImageCVEEdge{},
		&storage.ImageIntegration{},
		&storage.IntegrationHealth{},
		&storage.K8SRoleBinding{},
		&storage.K8SRole{},
		&storage.LogImbue{},
		&storage.NamespaceMetadata{},
		&storage.NetworkBaseline{},
		&storage.NetworkEntity{},
		&storage.NetworkFlow{},
		&storage.NetworkGraphConfig{},
		&storage.NetworkPolicyApplicationUndoDeploymentRecord{},
		&storage.NetworkPolicyApplicationUndoRecord{},
		&storage.NodeComponent{},
		&storage.NodeComponentCVEEdge{},
		&storage.NodeComponentEdge{},
		&storage.NodeCVE{},
		&storage.NotificationSchedule{},
		&storage.Notifier{},
		&storage.NotifierEncConfig{},
		&storage.PermissionSet{},
		&storage.Pod{},
		&storage.Policy{},
		&storage.PolicyCategory{},
		&storage.PolicyCategoryEdge{},
		&storage.ProcessBaselineResults{},
		&storage.ProcessBaseline{},
		&storage.ProcessIndicator{},
		&storage.ProcessListeningOnPortStorage{},
		&storage.ReportConfiguration{},
		&storage.ReportSnapshot{},
		&storage.ResourceCollection{},
		&storage.Risk{},
		&storage.Role{},
		&storage.ComplianceOperatorScanSettingBindingV2{},
		&storage.SecuredUnits{},
		&storage.SensorUpgradeConfig{},
		&storage.ServiceIdentity{},
		&storage.SignatureIntegration{},
		&storage.SimpleAccessScope{},
		&storage.SystemInfo{},
		&storage.TelemetryConfiguration{},
		&storage.TokenMetadata{},
		&storage.User{},
	} {
		// All the string representations of the type of the value are pointers, so remove the "*" in front of the
		// those string representations (e.g., "*storage.ActiveComponent" -> "storage.ActiveComponent") for each key
		typeRegistry[fmt.Sprintf("%T", v)[1:]] = v
	}
}
