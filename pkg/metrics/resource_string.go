// Code generated by "stringer -type=Resource"; DO NOT EDIT.

package metrics

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Alert-0]
	_ = x[Deployment-1]
	_ = x[ProcessIndicator-2]
	_ = x[ProcessListeningOnPort-3]
	_ = x[Image-4]
	_ = x[Secret-5]
	_ = x[Namespace-6]
	_ = x[NetworkPolicy-7]
	_ = x[Node-8]
	_ = x[NodeInventory-9]
	_ = x[ProviderMetadata-10]
	_ = x[ComplianceReturn-11]
	_ = x[ImageIntegration-12]
	_ = x[ServiceAccount-13]
	_ = x[PermissionSet-14]
	_ = x[Role-15]
	_ = x[RoleBinding-16]
	_ = x[DeploymentReprocess-17]
	_ = x[Pod-18]
	_ = x[ComplianceOperatorCheckResult-19]
	_ = x[ComplianceOperatorProfile-20]
	_ = x[ComplianceOperatorScanSettingBinding-21]
	_ = x[ComplianceOperatorRule-22]
	_ = x[ComplianceOperatorScan-23]
	_ = x[ComplianceOperatorInfo-24]
	_ = x[ComplianceOperatorCheckResultV2-25]
	_ = x[ComplianceOperatorRuleV2-26]
	_ = x[ComplianceOperatorProfileV2-27]
}

const _Resource_name = "AlertDeploymentProcessIndicatorProcessListeningOnPortImageSecretNamespaceNetworkPolicyNodeNodeInventoryProviderMetadataComplianceReturnImageIntegrationServiceAccountPermissionSetRoleRoleBindingDeploymentReprocessPodComplianceOperatorCheckResultComplianceOperatorProfileComplianceOperatorScanSettingBindingComplianceOperatorRuleComplianceOperatorScanComplianceOperatorInfoComplianceOperatorCheckResultV2ComplianceOperatorRuleV2ComplianceOperatorProfileV2"

var _Resource_index = [...]uint16{0, 5, 15, 31, 53, 58, 64, 73, 86, 90, 103, 119, 135, 151, 165, 178, 182, 193, 212, 215, 244, 269, 305, 327, 349, 371, 402, 426, 453}

func (i Resource) String() string {
	if i < 0 || i >= Resource(len(_Resource_index)-1) {
		return "Resource(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Resource_name[_Resource_index[i]:_Resource_index[i+1]]
}
