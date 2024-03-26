// Code generated by /home/janisz/go/src/github.com/stackrox/stackrox/.gotools/bin/protoc-gen-go-gogo. DO NOT EDIT.

package v2

func (m *ComplianceRule) Size() int                   { return m.SizeVT() }
func (m *ComplianceRule) Clone() *ComplianceRule      { return m.CloneVT() }
func (m *ComplianceRule) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ComplianceRule) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ComplianceProfile) Size() int                   { return m.SizeVT() }
func (m *ComplianceProfile) Clone() *ComplianceProfile   { return m.CloneVT() }
func (m *ComplianceProfile) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ComplianceProfile) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ComplianceProfileSummary) Size() int                        { return m.SizeVT() }
func (m *ComplianceProfileSummary) Clone() *ComplianceProfileSummary { return m.CloneVT() }
func (m *ComplianceProfileSummary) Marshal() ([]byte, error)         { return m.MarshalVT() }
func (m *ComplianceProfileSummary) Unmarshal(dAtA []byte) error      { return m.UnmarshalVT(dAtA) }

func (m *ListComplianceProfilesResponse) Size() int                              { return m.SizeVT() }
func (m *ListComplianceProfilesResponse) Clone() *ListComplianceProfilesResponse { return m.CloneVT() }
func (m *ListComplianceProfilesResponse) Marshal() ([]byte, error)               { return m.MarshalVT() }
func (m *ListComplianceProfilesResponse) Unmarshal(dAtA []byte) error            { return m.UnmarshalVT(dAtA) }

func (m *ListComplianceProfileSummaryResponse) Size() int { return m.SizeVT() }
func (m *ListComplianceProfileSummaryResponse) Clone() *ListComplianceProfileSummaryResponse {
	return m.CloneVT()
}
func (m *ListComplianceProfileSummaryResponse) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *ListComplianceProfileSummaryResponse) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *CountComplianceProfilesResponse) Size() int { return m.SizeVT() }
func (m *CountComplianceProfilesResponse) Clone() *CountComplianceProfilesResponse {
	return m.CloneVT()
}
func (m *CountComplianceProfilesResponse) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *CountComplianceProfilesResponse) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ProfilesForClusterRequest) Size() int                         { return m.SizeVT() }
func (m *ProfilesForClusterRequest) Clone() *ProfilesForClusterRequest { return m.CloneVT() }
func (m *ProfilesForClusterRequest) Marshal() ([]byte, error)          { return m.MarshalVT() }
func (m *ProfilesForClusterRequest) Unmarshal(dAtA []byte) error       { return m.UnmarshalVT(dAtA) }

func (m *ClustersProfileSummaryRequest) Size() int                             { return m.SizeVT() }
func (m *ClustersProfileSummaryRequest) Clone() *ClustersProfileSummaryRequest { return m.CloneVT() }
func (m *ClustersProfileSummaryRequest) Marshal() ([]byte, error)              { return m.MarshalVT() }
func (m *ClustersProfileSummaryRequest) Unmarshal(dAtA []byte) error           { return m.UnmarshalVT(dAtA) }
