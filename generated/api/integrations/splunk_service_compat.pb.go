// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package integrations

func (m *SplunkViolationsResponse) Size() int                        { return m.SizeVT() }
func (m *SplunkViolationsResponse) Clone() *SplunkViolationsResponse { return m.CloneVT() }
func (m *SplunkViolationsResponse) Marshal() ([]byte, error)         { return m.MarshalVT() }
func (m *SplunkViolationsResponse) Unmarshal(dAtA []byte) error      { return m.UnmarshalVT(dAtA) }

func (m *SplunkViolation) Size() int                   { return m.SizeVT() }
func (m *SplunkViolation) Clone() *SplunkViolation     { return m.CloneVT() }
func (m *SplunkViolation) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *SplunkViolation) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *SplunkViolation_ViolationInfo) Size() int                             { return m.SizeVT() }
func (m *SplunkViolation_ViolationInfo) Clone() *SplunkViolation_ViolationInfo { return m.CloneVT() }
func (m *SplunkViolation_ViolationInfo) Marshal() ([]byte, error)              { return m.MarshalVT() }
func (m *SplunkViolation_ViolationInfo) Unmarshal(dAtA []byte) error           { return m.UnmarshalVT(dAtA) }

func (m *SplunkViolation_AlertInfo) Size() int                         { return m.SizeVT() }
func (m *SplunkViolation_AlertInfo) Clone() *SplunkViolation_AlertInfo { return m.CloneVT() }
func (m *SplunkViolation_AlertInfo) Marshal() ([]byte, error)          { return m.MarshalVT() }
func (m *SplunkViolation_AlertInfo) Unmarshal(dAtA []byte) error       { return m.UnmarshalVT(dAtA) }

func (m *SplunkViolation_ProcessInfo) Size() int                           { return m.SizeVT() }
func (m *SplunkViolation_ProcessInfo) Clone() *SplunkViolation_ProcessInfo { return m.CloneVT() }
func (m *SplunkViolation_ProcessInfo) Marshal() ([]byte, error)            { return m.MarshalVT() }
func (m *SplunkViolation_ProcessInfo) Unmarshal(dAtA []byte) error         { return m.UnmarshalVT(dAtA) }

func (m *SplunkViolation_DeploymentInfo) Size() int                              { return m.SizeVT() }
func (m *SplunkViolation_DeploymentInfo) Clone() *SplunkViolation_DeploymentInfo { return m.CloneVT() }
func (m *SplunkViolation_DeploymentInfo) Marshal() ([]byte, error)               { return m.MarshalVT() }
func (m *SplunkViolation_DeploymentInfo) Unmarshal(dAtA []byte) error            { return m.UnmarshalVT(dAtA) }

func (m *SplunkViolation_ResourceInfo) Size() int                            { return m.SizeVT() }
func (m *SplunkViolation_ResourceInfo) Clone() *SplunkViolation_ResourceInfo { return m.CloneVT() }
func (m *SplunkViolation_ResourceInfo) Marshal() ([]byte, error)             { return m.MarshalVT() }
func (m *SplunkViolation_ResourceInfo) Unmarshal(dAtA []byte) error          { return m.UnmarshalVT(dAtA) }

func (m *SplunkViolation_PolicyInfo) Size() int                          { return m.SizeVT() }
func (m *SplunkViolation_PolicyInfo) Clone() *SplunkViolation_PolicyInfo { return m.CloneVT() }
func (m *SplunkViolation_PolicyInfo) Marshal() ([]byte, error)           { return m.MarshalVT() }
func (m *SplunkViolation_PolicyInfo) Unmarshal(dAtA []byte) error        { return m.UnmarshalVT(dAtA) }
