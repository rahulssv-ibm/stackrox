// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package sensor

func (m *GetScrapeConfigRequest) Size() int                      { return m.SizeVT() }
func (m *GetScrapeConfigRequest) Clone() *GetScrapeConfigRequest { return m.CloneVT() }
func (m *GetScrapeConfigRequest) Marshal() ([]byte, error)       { return m.MarshalVT() }
func (m *GetScrapeConfigRequest) Unmarshal(dAtA []byte) error    { return m.UnmarshalVT(dAtA) }

func (m *AuditEvents) Size() int                   { return m.SizeVT() }
func (m *AuditEvents) Clone() *AuditEvents         { return m.CloneVT() }
func (m *AuditEvents) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *AuditEvents) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *MsgFromCompliance) Size() int                   { return m.SizeVT() }
func (m *MsgFromCompliance) Clone() *MsgFromCompliance   { return m.CloneVT() }
func (m *MsgFromCompliance) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *MsgFromCompliance) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *MsgToCompliance) Size() int                   { return m.SizeVT() }
func (m *MsgToCompliance) Clone() *MsgToCompliance     { return m.CloneVT() }
func (m *MsgToCompliance) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *MsgToCompliance) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *MsgToCompliance_ScrapeConfig) Size() int                            { return m.SizeVT() }
func (m *MsgToCompliance_ScrapeConfig) Clone() *MsgToCompliance_ScrapeConfig { return m.CloneVT() }
func (m *MsgToCompliance_ScrapeConfig) Marshal() ([]byte, error)             { return m.MarshalVT() }
func (m *MsgToCompliance_ScrapeConfig) Unmarshal(dAtA []byte) error          { return m.UnmarshalVT(dAtA) }

func (m *MsgToCompliance_TriggerRun) Size() int                          { return m.SizeVT() }
func (m *MsgToCompliance_TriggerRun) Clone() *MsgToCompliance_TriggerRun { return m.CloneVT() }
func (m *MsgToCompliance_TriggerRun) Marshal() ([]byte, error)           { return m.MarshalVT() }
func (m *MsgToCompliance_TriggerRun) Unmarshal(dAtA []byte) error        { return m.UnmarshalVT(dAtA) }

func (m *MsgToCompliance_AuditLogCollectionRequest) Size() int { return m.SizeVT() }
func (m *MsgToCompliance_AuditLogCollectionRequest) Clone() *MsgToCompliance_AuditLogCollectionRequest {
	return m.CloneVT()
}
func (m *MsgToCompliance_AuditLogCollectionRequest) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *MsgToCompliance_AuditLogCollectionRequest) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *MsgToCompliance_AuditLogCollectionRequest_StartRequest) Size() int { return m.SizeVT() }
func (m *MsgToCompliance_AuditLogCollectionRequest_StartRequest) Clone() *MsgToCompliance_AuditLogCollectionRequest_StartRequest {
	return m.CloneVT()
}
func (m *MsgToCompliance_AuditLogCollectionRequest_StartRequest) Marshal() ([]byte, error) {
	return m.MarshalVT()
}
func (m *MsgToCompliance_AuditLogCollectionRequest_StartRequest) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *MsgToCompliance_AuditLogCollectionRequest_StopRequest) Size() int { return m.SizeVT() }
func (m *MsgToCompliance_AuditLogCollectionRequest_StopRequest) Clone() *MsgToCompliance_AuditLogCollectionRequest_StopRequest {
	return m.CloneVT()
}
func (m *MsgToCompliance_AuditLogCollectionRequest_StopRequest) Marshal() ([]byte, error) {
	return m.MarshalVT()
}
func (m *MsgToCompliance_AuditLogCollectionRequest_StopRequest) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *MsgToCompliance_NodeInventoryACK) Size() int { return m.SizeVT() }
func (m *MsgToCompliance_NodeInventoryACK) Clone() *MsgToCompliance_NodeInventoryACK {
	return m.CloneVT()
}
func (m *MsgToCompliance_NodeInventoryACK) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *MsgToCompliance_NodeInventoryACK) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }
