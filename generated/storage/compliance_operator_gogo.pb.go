// Code generated by /home/janisz/go/src/github.com/stackrox/stackrox/.gotools/bin/protoc-gen-go-gogo. DO NOT EDIT.

package storage

func (m *ComplianceOperatorCheckResult) Size() int                             { return m.SizeVT() }
func (m *ComplianceOperatorCheckResult) Clone() *ComplianceOperatorCheckResult { return m.CloneVT() }
func (m *ComplianceOperatorCheckResult) Marshal() ([]byte, error)              { return m.MarshalVT() }
func (m *ComplianceOperatorCheckResult) Unmarshal(dAtA []byte) error           { return m.UnmarshalVT(dAtA) }

func (m *ComplianceOperatorProfile) Size() int                         { return m.SizeVT() }
func (m *ComplianceOperatorProfile) Clone() *ComplianceOperatorProfile { return m.CloneVT() }
func (m *ComplianceOperatorProfile) Marshal() ([]byte, error)          { return m.MarshalVT() }
func (m *ComplianceOperatorProfile) Unmarshal(dAtA []byte) error       { return m.UnmarshalVT(dAtA) }

func (m *ComplianceOperatorProfile_Rule) Size() int                              { return m.SizeVT() }
func (m *ComplianceOperatorProfile_Rule) Clone() *ComplianceOperatorProfile_Rule { return m.CloneVT() }
func (m *ComplianceOperatorProfile_Rule) Marshal() ([]byte, error)               { return m.MarshalVT() }
func (m *ComplianceOperatorProfile_Rule) Unmarshal(dAtA []byte) error            { return m.UnmarshalVT(dAtA) }

func (m *ComplianceOperatorRule) Size() int                      { return m.SizeVT() }
func (m *ComplianceOperatorRule) Clone() *ComplianceOperatorRule { return m.CloneVT() }
func (m *ComplianceOperatorRule) Marshal() ([]byte, error)       { return m.MarshalVT() }
func (m *ComplianceOperatorRule) Unmarshal(dAtA []byte) error    { return m.UnmarshalVT(dAtA) }

func (m *ComplianceOperatorScanSettingBinding) Size() int { return m.SizeVT() }
func (m *ComplianceOperatorScanSettingBinding) Clone() *ComplianceOperatorScanSettingBinding {
	return m.CloneVT()
}
func (m *ComplianceOperatorScanSettingBinding) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *ComplianceOperatorScanSettingBinding) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *ComplianceOperatorScanSettingBinding_Profile) Size() int { return m.SizeVT() }
func (m *ComplianceOperatorScanSettingBinding_Profile) Clone() *ComplianceOperatorScanSettingBinding_Profile {
	return m.CloneVT()
}
func (m *ComplianceOperatorScanSettingBinding_Profile) Marshal() ([]byte, error) {
	return m.MarshalVT()
}
func (m *ComplianceOperatorScanSettingBinding_Profile) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *ComplianceOperatorScan) Size() int                      { return m.SizeVT() }
func (m *ComplianceOperatorScan) Clone() *ComplianceOperatorScan { return m.CloneVT() }
func (m *ComplianceOperatorScan) Marshal() ([]byte, error)       { return m.MarshalVT() }
func (m *ComplianceOperatorScan) Unmarshal(dAtA []byte) error    { return m.UnmarshalVT(dAtA) }
