// Code generated by /home/janisz/go/src/github.com/stackrox/stackrox/.gotools/bin/protoc-gen-go-gogo. DO NOT EDIT.

package v1

func (m *DeploymentLabelsResponse) Size() int                        { return m.SizeVT() }
func (m *DeploymentLabelsResponse) Clone() *DeploymentLabelsResponse { return m.CloneVT() }
func (m *DeploymentLabelsResponse) Marshal() ([]byte, error)         { return m.MarshalVT() }
func (m *DeploymentLabelsResponse) Unmarshal(dAtA []byte) error      { return m.UnmarshalVT(dAtA) }

func (m *DeploymentLabelsResponse_LabelValues) Size() int { return m.SizeVT() }
func (m *DeploymentLabelsResponse_LabelValues) Clone() *DeploymentLabelsResponse_LabelValues {
	return m.CloneVT()
}
func (m *DeploymentLabelsResponse_LabelValues) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *DeploymentLabelsResponse_LabelValues) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *ListDeploymentsResponse) Size() int                       { return m.SizeVT() }
func (m *ListDeploymentsResponse) Clone() *ListDeploymentsResponse { return m.CloneVT() }
func (m *ListDeploymentsResponse) Marshal() ([]byte, error)        { return m.MarshalVT() }
func (m *ListDeploymentsResponse) Unmarshal(dAtA []byte) error     { return m.UnmarshalVT(dAtA) }

func (m *CountDeploymentsResponse) Size() int                        { return m.SizeVT() }
func (m *CountDeploymentsResponse) Clone() *CountDeploymentsResponse { return m.CloneVT() }
func (m *CountDeploymentsResponse) Marshal() ([]byte, error)         { return m.MarshalVT() }
func (m *CountDeploymentsResponse) Unmarshal(dAtA []byte) error      { return m.UnmarshalVT(dAtA) }

func (m *ListDeploymentsWithProcessInfoResponse) Size() int { return m.SizeVT() }
func (m *ListDeploymentsWithProcessInfoResponse) Clone() *ListDeploymentsWithProcessInfoResponse {
	return m.CloneVT()
}
func (m *ListDeploymentsWithProcessInfoResponse) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *ListDeploymentsWithProcessInfoResponse) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo) Size() int {
	return m.SizeVT()
}
func (m *ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo) Clone() *ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo {
	return m.CloneVT()
}
func (m *ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo) Marshal() ([]byte, error) {
	return m.MarshalVT()
}
func (m *ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *GetDeploymentWithRiskResponse) Size() int                             { return m.SizeVT() }
func (m *GetDeploymentWithRiskResponse) Clone() *GetDeploymentWithRiskResponse { return m.CloneVT() }
func (m *GetDeploymentWithRiskResponse) Marshal() ([]byte, error)              { return m.MarshalVT() }
func (m *GetDeploymentWithRiskResponse) Unmarshal(dAtA []byte) error           { return m.UnmarshalVT(dAtA) }

func (m *ExportDeploymentRequest) Size() int                       { return m.SizeVT() }
func (m *ExportDeploymentRequest) Clone() *ExportDeploymentRequest { return m.CloneVT() }
func (m *ExportDeploymentRequest) Marshal() ([]byte, error)        { return m.MarshalVT() }
func (m *ExportDeploymentRequest) Unmarshal(dAtA []byte) error     { return m.UnmarshalVT(dAtA) }

func (m *ExportDeploymentResponse) Size() int                        { return m.SizeVT() }
func (m *ExportDeploymentResponse) Clone() *ExportDeploymentResponse { return m.CloneVT() }
func (m *ExportDeploymentResponse) Marshal() ([]byte, error)         { return m.MarshalVT() }
func (m *ExportDeploymentResponse) Unmarshal(dAtA []byte) error      { return m.UnmarshalVT(dAtA) }
