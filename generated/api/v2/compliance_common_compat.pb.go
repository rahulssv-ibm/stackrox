// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package v2

func (m *ComplianceRule) Size() int                   { return m.SizeVT() }
func (m *ComplianceRule) Clone() *ComplianceRule      { return m.CloneVT() }
func (m *ComplianceRule) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ComplianceRule) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ComplianceRule_Fix) Size() int                   { return m.SizeVT() }
func (m *ComplianceRule_Fix) Clone() *ComplianceRule_Fix  { return m.CloneVT() }
func (m *ComplianceRule_Fix) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ComplianceRule_Fix) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ComplianceScanCluster) Size() int                     { return m.SizeVT() }
func (m *ComplianceScanCluster) Clone() *ComplianceScanCluster { return m.CloneVT() }
func (m *ComplianceScanCluster) Marshal() ([]byte, error)      { return m.MarshalVT() }
func (m *ComplianceScanCluster) Unmarshal(dAtA []byte) error   { return m.UnmarshalVT(dAtA) }

func (m *ComplianceCheckStatusCount) Size() int                          { return m.SizeVT() }
func (m *ComplianceCheckStatusCount) Clone() *ComplianceCheckStatusCount { return m.CloneVT() }
func (m *ComplianceCheckStatusCount) Marshal() ([]byte, error)           { return m.MarshalVT() }
func (m *ComplianceCheckStatusCount) Unmarshal(dAtA []byte) error        { return m.UnmarshalVT(dAtA) }

func (m *ComplianceCheckResultStatusCount) Size() int { return m.SizeVT() }
func (m *ComplianceCheckResultStatusCount) Clone() *ComplianceCheckResultStatusCount {
	return m.CloneVT()
}
func (m *ComplianceCheckResultStatusCount) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ComplianceCheckResultStatusCount) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ComplianceControl) Size() int                   { return m.SizeVT() }
func (m *ComplianceControl) Clone() *ComplianceControl   { return m.CloneVT() }
func (m *ComplianceControl) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ComplianceControl) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ComplianceBenchmark) Size() int                   { return m.SizeVT() }
func (m *ComplianceBenchmark) Clone() *ComplianceBenchmark { return m.CloneVT() }
func (m *ComplianceBenchmark) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ComplianceBenchmark) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ListComplianceProfileResults) Size() int                            { return m.SizeVT() }
func (m *ListComplianceProfileResults) Clone() *ListComplianceProfileResults { return m.CloneVT() }
func (m *ListComplianceProfileResults) Marshal() ([]byte, error)             { return m.MarshalVT() }
func (m *ListComplianceProfileResults) Unmarshal(dAtA []byte) error          { return m.UnmarshalVT(dAtA) }

func (m *ComplianceClusterOverallStats) Size() int                             { return m.SizeVT() }
func (m *ComplianceClusterOverallStats) Clone() *ComplianceClusterOverallStats { return m.CloneVT() }
func (m *ComplianceClusterOverallStats) Marshal() ([]byte, error)              { return m.MarshalVT() }
func (m *ComplianceClusterOverallStats) Unmarshal(dAtA []byte) error           { return m.UnmarshalVT(dAtA) }

func (m *ListComplianceClusterOverallStatsResponse) Size() int { return m.SizeVT() }
func (m *ListComplianceClusterOverallStatsResponse) Clone() *ListComplianceClusterOverallStatsResponse {
	return m.CloneVT()
}
func (m *ListComplianceClusterOverallStatsResponse) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *ListComplianceClusterOverallStatsResponse) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *ComplianceProfileResultsRequest) Size() int { return m.SizeVT() }
func (m *ComplianceProfileResultsRequest) Clone() *ComplianceProfileResultsRequest {
	return m.CloneVT()
}
func (m *ComplianceProfileResultsRequest) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ComplianceProfileResultsRequest) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ComplianceProfileCheckRequest) Size() int                             { return m.SizeVT() }
func (m *ComplianceProfileCheckRequest) Clone() *ComplianceProfileCheckRequest { return m.CloneVT() }
func (m *ComplianceProfileCheckRequest) Marshal() ([]byte, error)              { return m.MarshalVT() }
func (m *ComplianceProfileCheckRequest) Unmarshal(dAtA []byte) error           { return m.UnmarshalVT(dAtA) }

func (m *ComplianceProfileSummary) Size() int                        { return m.SizeVT() }
func (m *ComplianceProfileSummary) Clone() *ComplianceProfileSummary { return m.CloneVT() }
func (m *ComplianceProfileSummary) Marshal() ([]byte, error)         { return m.MarshalVT() }
func (m *ComplianceProfileSummary) Unmarshal(dAtA []byte) error      { return m.UnmarshalVT(dAtA) }
