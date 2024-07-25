// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package v1

func (m *PutConfigRequest) Size() int                   { return m.SizeVT() }
func (m *PutConfigRequest) Clone() *PutConfigRequest    { return m.CloneVT() }
func (m *PutConfigRequest) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *PutConfigRequest) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *DayOption) Size() int                   { return m.SizeVT() }
func (m *DayOption) Clone() *DayOption           { return m.CloneVT() }
func (m *DayOption) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *DayOption) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *VulnerabilityExceptionConfig) Size() int                            { return m.SizeVT() }
func (m *VulnerabilityExceptionConfig) Clone() *VulnerabilityExceptionConfig { return m.CloneVT() }
func (m *VulnerabilityExceptionConfig) Marshal() ([]byte, error)             { return m.MarshalVT() }
func (m *VulnerabilityExceptionConfig) Unmarshal(dAtA []byte) error          { return m.UnmarshalVT(dAtA) }

func (m *VulnerabilityExceptionConfig_FixableCVEOptions) Size() int { return m.SizeVT() }
func (m *VulnerabilityExceptionConfig_FixableCVEOptions) Clone() *VulnerabilityExceptionConfig_FixableCVEOptions {
	return m.CloneVT()
}
func (m *VulnerabilityExceptionConfig_FixableCVEOptions) Marshal() ([]byte, error) {
	return m.MarshalVT()
}
func (m *VulnerabilityExceptionConfig_FixableCVEOptions) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *VulnerabilityExceptionConfig_ExpiryOptions) Size() int { return m.SizeVT() }
func (m *VulnerabilityExceptionConfig_ExpiryOptions) Clone() *VulnerabilityExceptionConfig_ExpiryOptions {
	return m.CloneVT()
}
func (m *VulnerabilityExceptionConfig_ExpiryOptions) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *VulnerabilityExceptionConfig_ExpiryOptions) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *GetVulnerabilityExceptionConfigResponse) Size() int { return m.SizeVT() }
func (m *GetVulnerabilityExceptionConfigResponse) Clone() *GetVulnerabilityExceptionConfigResponse {
	return m.CloneVT()
}
func (m *GetVulnerabilityExceptionConfigResponse) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *GetVulnerabilityExceptionConfigResponse) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *UpdateVulnerabilityExceptionConfigRequest) Size() int { return m.SizeVT() }
func (m *UpdateVulnerabilityExceptionConfigRequest) Clone() *UpdateVulnerabilityExceptionConfigRequest {
	return m.CloneVT()
}
func (m *UpdateVulnerabilityExceptionConfigRequest) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *UpdateVulnerabilityExceptionConfigRequest) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *UpdateVulnerabilityExceptionConfigResponse) Size() int { return m.SizeVT() }
func (m *UpdateVulnerabilityExceptionConfigResponse) Clone() *UpdateVulnerabilityExceptionConfigResponse {
	return m.CloneVT()
}
func (m *UpdateVulnerabilityExceptionConfigResponse) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *UpdateVulnerabilityExceptionConfigResponse) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}
