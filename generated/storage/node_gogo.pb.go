// Code generated by protoc-gen-go-gogo. DO NOT EDIT.

package storage

func (m *Node) Size() int                   { return m.SizeVT() }
func (m *Node) Clone() *Node                { return m.CloneVT() }
func (m *Node) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *Node) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *NodeScan) Size() int                   { return m.SizeVT() }
func (m *NodeScan) Clone() *NodeScan            { return m.CloneVT() }
func (m *NodeScan) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *NodeScan) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *NodeInventory) Size() int                   { return m.SizeVT() }
func (m *NodeInventory) Clone() *NodeInventory       { return m.CloneVT() }
func (m *NodeInventory) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *NodeInventory) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *NodeInventory_Components) Size() int                        { return m.SizeVT() }
func (m *NodeInventory_Components) Clone() *NodeInventory_Components { return m.CloneVT() }
func (m *NodeInventory_Components) Marshal() ([]byte, error)         { return m.MarshalVT() }
func (m *NodeInventory_Components) Unmarshal(dAtA []byte) error      { return m.UnmarshalVT(dAtA) }

func (m *NodeInventory_Components_RHELComponent) Size() int { return m.SizeVT() }
func (m *NodeInventory_Components_RHELComponent) Clone() *NodeInventory_Components_RHELComponent {
	return m.CloneVT()
}
func (m *NodeInventory_Components_RHELComponent) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *NodeInventory_Components_RHELComponent) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *NodeInventory_Components_RHELComponent_Executable) Size() int { return m.SizeVT() }
func (m *NodeInventory_Components_RHELComponent_Executable) Clone() *NodeInventory_Components_RHELComponent_Executable {
	return m.CloneVT()
}
func (m *NodeInventory_Components_RHELComponent_Executable) Marshal() ([]byte, error) {
	return m.MarshalVT()
}
func (m *NodeInventory_Components_RHELComponent_Executable) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *NodeInventory_Components_RHELComponent_Executable_FeatureNameVersion) Size() int {
	return m.SizeVT()
}
func (m *NodeInventory_Components_RHELComponent_Executable_FeatureNameVersion) Clone() *NodeInventory_Components_RHELComponent_Executable_FeatureNameVersion {
	return m.CloneVT()
}
func (m *NodeInventory_Components_RHELComponent_Executable_FeatureNameVersion) Marshal() ([]byte, error) {
	return m.MarshalVT()
}
func (m *NodeInventory_Components_RHELComponent_Executable_FeatureNameVersion) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *EmbeddedNodeScanComponent) Size() int                         { return m.SizeVT() }
func (m *EmbeddedNodeScanComponent) Clone() *EmbeddedNodeScanComponent { return m.CloneVT() }
func (m *EmbeddedNodeScanComponent) Marshal() ([]byte, error)          { return m.MarshalVT() }
func (m *EmbeddedNodeScanComponent) Unmarshal(dAtA []byte) error       { return m.UnmarshalVT(dAtA) }
