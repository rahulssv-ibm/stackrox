// Code generated by /home/janisz/go/src/github.com/stackrox/stackrox/.gotools/bin/protoc-gen-go-gogo. DO NOT EDIT.

package v1

func (m *NetworkBaselinePeerEntity) Size() int                         { return m.SizeVT() }
func (m *NetworkBaselinePeerEntity) Clone() *NetworkBaselinePeerEntity { return m.CloneVT() }
func (m *NetworkBaselinePeerEntity) Marshal() ([]byte, error)          { return m.MarshalVT() }
func (m *NetworkBaselinePeerEntity) Unmarshal(dAtA []byte) error       { return m.UnmarshalVT(dAtA) }

func (m *NetworkBaselineStatusPeer) Size() int                         { return m.SizeVT() }
func (m *NetworkBaselineStatusPeer) Clone() *NetworkBaselineStatusPeer { return m.CloneVT() }
func (m *NetworkBaselineStatusPeer) Marshal() ([]byte, error)          { return m.MarshalVT() }
func (m *NetworkBaselineStatusPeer) Unmarshal(dAtA []byte) error       { return m.UnmarshalVT(dAtA) }

func (m *NetworkBaselinePeerStatus) Size() int                         { return m.SizeVT() }
func (m *NetworkBaselinePeerStatus) Clone() *NetworkBaselinePeerStatus { return m.CloneVT() }
func (m *NetworkBaselinePeerStatus) Marshal() ([]byte, error)          { return m.MarshalVT() }
func (m *NetworkBaselinePeerStatus) Unmarshal(dAtA []byte) error       { return m.UnmarshalVT(dAtA) }

func (m *NetworkBaselineStatusRequest) Size() int                            { return m.SizeVT() }
func (m *NetworkBaselineStatusRequest) Clone() *NetworkBaselineStatusRequest { return m.CloneVT() }
func (m *NetworkBaselineStatusRequest) Marshal() ([]byte, error)             { return m.MarshalVT() }
func (m *NetworkBaselineStatusRequest) Unmarshal(dAtA []byte) error          { return m.UnmarshalVT(dAtA) }

func (m *NetworkBaselineStatusResponse) Size() int                             { return m.SizeVT() }
func (m *NetworkBaselineStatusResponse) Clone() *NetworkBaselineStatusResponse { return m.CloneVT() }
func (m *NetworkBaselineStatusResponse) Marshal() ([]byte, error)              { return m.MarshalVT() }
func (m *NetworkBaselineStatusResponse) Unmarshal(dAtA []byte) error           { return m.UnmarshalVT(dAtA) }

func (m *ModifyBaselineStatusForPeersRequest) Size() int { return m.SizeVT() }
func (m *ModifyBaselineStatusForPeersRequest) Clone() *ModifyBaselineStatusForPeersRequest {
	return m.CloneVT()
}
func (m *ModifyBaselineStatusForPeersRequest) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *ModifyBaselineStatusForPeersRequest) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}
