// Code generated by /home/janisz/go/src/github.com/stackrox/stackrox/.gotools/bin/protoc-gen-go-gogo. DO NOT EDIT.

package v2

func (m *ClusterScanStatus) Size() int                   { return m.SizeVT() }
func (m *ClusterScanStatus) Clone() *ClusterScanStatus   { return m.CloneVT() }
func (m *ClusterScanStatus) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ClusterScanStatus) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ClusterScanStatus_SuiteStatus) Size() int                             { return m.SizeVT() }
func (m *ClusterScanStatus_SuiteStatus) Clone() *ClusterScanStatus_SuiteStatus { return m.CloneVT() }
func (m *ClusterScanStatus_SuiteStatus) Marshal() ([]byte, error)              { return m.MarshalVT() }
func (m *ClusterScanStatus_SuiteStatus) Unmarshal(dAtA []byte) error           { return m.UnmarshalVT(dAtA) }

func (m *BaseComplianceScanConfigurationSettings) Size() int { return m.SizeVT() }
func (m *BaseComplianceScanConfigurationSettings) Clone() *BaseComplianceScanConfigurationSettings {
	return m.CloneVT()
}
func (m *BaseComplianceScanConfigurationSettings) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *BaseComplianceScanConfigurationSettings) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *ComplianceScanConfiguration) Size() int                           { return m.SizeVT() }
func (m *ComplianceScanConfiguration) Clone() *ComplianceScanConfiguration { return m.CloneVT() }
func (m *ComplianceScanConfiguration) Marshal() ([]byte, error)            { return m.MarshalVT() }
func (m *ComplianceScanConfiguration) Unmarshal(dAtA []byte) error         { return m.UnmarshalVT(dAtA) }

func (m *ComplianceScanConfigurationStatus) Size() int { return m.SizeVT() }
func (m *ComplianceScanConfigurationStatus) Clone() *ComplianceScanConfigurationStatus {
	return m.CloneVT()
}
func (m *ComplianceScanConfigurationStatus) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ComplianceScanConfigurationStatus) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ComplianceScanConfigurationsCount) Size() int { return m.SizeVT() }
func (m *ComplianceScanConfigurationsCount) Clone() *ComplianceScanConfigurationsCount {
	return m.CloneVT()
}
func (m *ComplianceScanConfigurationsCount) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ComplianceScanConfigurationsCount) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ListComplianceScanConfigurationsResponse) Size() int { return m.SizeVT() }
func (m *ListComplianceScanConfigurationsResponse) Clone() *ListComplianceScanConfigurationsResponse {
	return m.CloneVT()
}
func (m *ListComplianceScanConfigurationsResponse) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *ListComplianceScanConfigurationsResponse) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}
