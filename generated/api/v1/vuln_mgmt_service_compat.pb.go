// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package v1

func (m *VulnMgmtExportWorkloadsRequest) Size() int                              { return m.SizeVT() }
func (m *VulnMgmtExportWorkloadsRequest) Clone() *VulnMgmtExportWorkloadsRequest { return m.CloneVT() }
func (m *VulnMgmtExportWorkloadsRequest) Marshal() ([]byte, error)               { return m.MarshalVT() }
func (m *VulnMgmtExportWorkloadsRequest) Unmarshal(dAtA []byte) error            { return m.UnmarshalVT(dAtA) }

func (m *VulnMgmtExportWorkloadsResponse) Size() int { return m.SizeVT() }
func (m *VulnMgmtExportWorkloadsResponse) Clone() *VulnMgmtExportWorkloadsResponse {
	return m.CloneVT()
}
func (m *VulnMgmtExportWorkloadsResponse) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *VulnMgmtExportWorkloadsResponse) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }
