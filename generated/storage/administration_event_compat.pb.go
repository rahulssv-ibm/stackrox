// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package storage

func (m *AdministrationEvent) Size() int                   { return m.SizeVT() }
func (m *AdministrationEvent) Clone() *AdministrationEvent { return m.CloneVT() }
func (m *AdministrationEvent) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *AdministrationEvent) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *AdministrationEvent_Resource) Size() int                            { return m.SizeVT() }
func (m *AdministrationEvent_Resource) Clone() *AdministrationEvent_Resource { return m.CloneVT() }
func (m *AdministrationEvent_Resource) Marshal() ([]byte, error)             { return m.MarshalVT() }
func (m *AdministrationEvent_Resource) Unmarshal(dAtA []byte) error          { return m.UnmarshalVT(dAtA) }
