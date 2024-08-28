package storagetov2

import (
	"github.com/stackrox/rox/central/complianceoperator/v2/checkresults/datastore"
	compRule "github.com/stackrox/rox/central/complianceoperator/v2/rules/datastore"
	v2 "github.com/stackrox/rox/generated/api/v2"
	"github.com/stackrox/rox/generated/storage"
	types "github.com/stackrox/rox/pkg/protocompat"
)

type checkResultKey struct {
	scanConfigName string
	scanConfigID   string
	profileName    string
	checkName      string
}

type scanResultKey struct {
	scanConfigName string
	scanConfigID   string
	profileName    string
}

// ComplianceV2CheckResult converts a storage check result to a v2 check result
func ComplianceV2CheckResult(incoming *storage.ComplianceOperatorCheckResultV2, lastScanTime *types.Timestamp, ruleName string, controlResults []*compRule.ControlResult) *v2.ComplianceClusterCheckStatus {
	converted := &v2.ComplianceClusterCheckStatus{
		CheckId:   incoming.GetCheckId(),
		CheckName: incoming.GetCheckName(),
		Clusters: []*v2.ClusterCheckStatus{
			clusterStatus(incoming, lastScanTime),
		},
		Description:  incoming.GetDescription(),
		Instructions: incoming.GetInstructions(),
		Rationale:    incoming.GetRationale(),
		ValuesUsed:   incoming.GetValuesUsed(),
		Warnings:     incoming.GetWarnings(),
		Labels:       incoming.GetLabels(),
		Annotations:  incoming.GetAnnotations(),
		Controls:     GetControls(ruleName, controlResults),
	}

	return converted
}

// ComplianceV2SpecificCheckResult converts a storage check result to a v2 check result
func ComplianceV2SpecificCheckResult(incoming []*storage.ComplianceOperatorCheckResultV2, checkName string, controls []*v2.ComplianceControl) *v2.ComplianceClusterCheckStatus {
	var converted *v2.ComplianceClusterCheckStatus
	for _, result := range incoming {
		if result.GetCheckName() != checkName {
			continue
		}

		if converted == nil {
			converted = &v2.ComplianceClusterCheckStatus{
				CheckId:   result.GetCheckId(),
				CheckName: result.GetCheckName(),
				Clusters: []*v2.ClusterCheckStatus{
					clusterStatus(result, nil),
				},
				Description:  result.GetDescription(),
				Instructions: result.GetInstructions(),
				Rationale:    result.GetRationale(),
				ValuesUsed:   result.GetValuesUsed(),
				Warnings:     result.GetWarnings(),
				Labels:       result.GetLabels(),
				Annotations:  result.GetAnnotations(),
				Controls:     controls,
			}
		} else {
			converted.Clusters = append(converted.Clusters, clusterStatus(result, nil))
		}
	}

	return converted
}

// ComplianceV2ProfileResults converts the counts to the v2 stats
func ComplianceV2ProfileResults(resultCounts []*datastore.ResourceResultsByProfile, controlResults []*compRule.ControlResult) []*v2.ComplianceCheckResultStatusCount {
	var profileResults []*v2.ComplianceCheckResultStatusCount

	for _, resultCount := range resultCounts {
		controls := GetControls(resultCount.RuleName, controlResults)

		profileResults = append(profileResults, &v2.ComplianceCheckResultStatusCount{
			CheckName: resultCount.CheckName,
			Rationale: resultCount.CheckRationale,
			RuleName:  resultCount.RuleName,
			Controls:  controls,
			CheckStats: []*v2.ComplianceCheckStatusCount{
				{
					Count:  int32(resultCount.FailCount),
					Status: v2.ComplianceCheckStatus_FAIL,
				},
				{
					Count:  int32(resultCount.InfoCount),
					Status: v2.ComplianceCheckStatus_INFO,
				},
				{
					Count:  int32(resultCount.PassCount),
					Status: v2.ComplianceCheckStatus_PASS,
				},
				{
					Count:  int32(resultCount.ErrorCount),
					Status: v2.ComplianceCheckStatus_ERROR,
				},
				{
					Count:  int32(resultCount.ManualCount),
					Status: v2.ComplianceCheckStatus_MANUAL,
				},
				{
					Count:  int32(resultCount.InconsistentCount),
					Status: v2.ComplianceCheckStatus_INCONSISTENT,
				},
				{
					Count:  int32(resultCount.NotApplicableCount),
					Status: v2.ComplianceCheckStatus_NOT_APPLICABLE,
				},
			},
		})
	}

	return profileResults
}

// ComplianceV2CheckClusterResults converts the storage check results to v2 scan results
func ComplianceV2CheckClusterResults(incoming []*storage.ComplianceOperatorCheckResultV2, lastTimeMap map[string]*types.Timestamp) []*v2.ClusterCheckStatus {
	clusterResults := make([]*v2.ClusterCheckStatus, 0, len(incoming))
	for _, result := range incoming {
		clusterResults = append(clusterResults, clusterStatus(result, lastTimeMap[result.ClusterId]))
	}

	return clusterResults
}

// ComplianceV2CheckResults converts the storage check results to v2 scan results
func ComplianceV2CheckResults(incoming []*storage.ComplianceOperatorCheckResultV2, ruleMap map[string]string, controlResults []*compRule.ControlResult) []*v2.ComplianceCheckResult {
	clusterResults := make([]*v2.ComplianceCheckResult, 0, len(incoming))
	for _, result := range incoming {
		clusterResults = append(clusterResults, checkResult(result, ruleMap[result.GetRuleRefId()], controlResults))
	}

	return clusterResults
}

func ComplianceV2CheckData(incoming []*storage.ComplianceOperatorCheckResultV2, ruleMap map[string]string, controlMap map[string][]*compRule.ControlResult) []*v2.ComplianceCheckData {
	results := make([]*v2.ComplianceCheckData, 0, len(incoming))
	for _, result := range incoming {
		results = append(results, &v2.ComplianceCheckData{
			ClusterId: result.GetClusterId(),
			ScanName:  result.GetScanConfigName(),
			Result:    checkResult(result, ruleMap[result.GetRuleRefId()], controlMap[result.GetCheckName()]),
		})
	}

	return results
}

func clusterStatus(incoming *storage.ComplianceOperatorCheckResultV2, lastScanTime *types.Timestamp) *v2.ClusterCheckStatus {
	return &v2.ClusterCheckStatus{
		Cluster: &v2.ComplianceScanCluster{
			ClusterId:   incoming.GetClusterId(),
			ClusterName: incoming.GetClusterName(),
		},
		Status:       convertComplianceCheckStatus(incoming.Status),
		CreatedTime:  incoming.GetCreatedTime(),
		CheckUid:     incoming.GetId(),
		LastScanTime: lastScanTime,
	}
}

func checkResult(incoming *storage.ComplianceOperatorCheckResultV2, ruleName string, controlResults []*compRule.ControlResult) *v2.ComplianceCheckResult {
	return &v2.ComplianceCheckResult{
		CheckId:      incoming.GetCheckId(),
		CheckName:    incoming.GetCheckName(),
		CheckUid:     incoming.GetId(),
		Description:  incoming.GetDescription(),
		Instructions: incoming.GetInstructions(),
		Controls:     GetControls(ruleName, controlResults),
		Rationale:    incoming.GetRationale(),
		ValuesUsed:   incoming.GetValuesUsed(),
		Warnings:     incoming.GetWarnings(),
		Status:       convertComplianceCheckStatus(incoming.Status),
		RuleName:     ruleName,
		Labels:       incoming.GetLabels(),
		Annotations:  incoming.GetAnnotations(),
	}
}

func GetControls(ruleName string, controlResults []*compRule.ControlResult) []*v2.ComplianceControl {
	var controls []*v2.ComplianceControl
	for _, controlResult := range controlResults {
		if controlResult.RuleName == ruleName {
			controls = append(controls, &v2.ComplianceControl{
				Standard: controlResult.Standard,
				Control:  controlResult.Control,
			})
		}
	}

	return controls
}
