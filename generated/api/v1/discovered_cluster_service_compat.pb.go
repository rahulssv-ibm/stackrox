// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package v1

func (m *DiscoveredCluster) Size() int                   { return m.SizeVT() }
func (m *DiscoveredCluster) Clone() *DiscoveredCluster   { return m.CloneVT() }
func (m *DiscoveredCluster) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *DiscoveredCluster) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *DiscoveredCluster_Metadata) Size() int                          { return m.SizeVT() }
func (m *DiscoveredCluster_Metadata) Clone() *DiscoveredCluster_Metadata { return m.CloneVT() }
func (m *DiscoveredCluster_Metadata) Marshal() ([]byte, error)           { return m.MarshalVT() }
func (m *DiscoveredCluster_Metadata) Unmarshal(dAtA []byte) error        { return m.UnmarshalVT(dAtA) }

func (m *DiscoveredCluster_CloudSource) Size() int                             { return m.SizeVT() }
func (m *DiscoveredCluster_CloudSource) Clone() *DiscoveredCluster_CloudSource { return m.CloneVT() }
func (m *DiscoveredCluster_CloudSource) Marshal() ([]byte, error)              { return m.MarshalVT() }
func (m *DiscoveredCluster_CloudSource) Unmarshal(dAtA []byte) error           { return m.UnmarshalVT(dAtA) }

func (m *DiscoveredClustersFilter) Size() int                        { return m.SizeVT() }
func (m *DiscoveredClustersFilter) Clone() *DiscoveredClustersFilter { return m.CloneVT() }
func (m *DiscoveredClustersFilter) Marshal() ([]byte, error)         { return m.MarshalVT() }
func (m *DiscoveredClustersFilter) Unmarshal(dAtA []byte) error      { return m.UnmarshalVT(dAtA) }

func (m *CountDiscoveredClustersRequest) Size() int                              { return m.SizeVT() }
func (m *CountDiscoveredClustersRequest) Clone() *CountDiscoveredClustersRequest { return m.CloneVT() }
func (m *CountDiscoveredClustersRequest) Marshal() ([]byte, error)               { return m.MarshalVT() }
func (m *CountDiscoveredClustersRequest) Unmarshal(dAtA []byte) error            { return m.UnmarshalVT(dAtA) }

func (m *CountDiscoveredClustersResponse) Size() int { return m.SizeVT() }
func (m *CountDiscoveredClustersResponse) Clone() *CountDiscoveredClustersResponse {
	return m.CloneVT()
}
func (m *CountDiscoveredClustersResponse) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *CountDiscoveredClustersResponse) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *GetDiscoveredClusterRequest) Size() int                           { return m.SizeVT() }
func (m *GetDiscoveredClusterRequest) Clone() *GetDiscoveredClusterRequest { return m.CloneVT() }
func (m *GetDiscoveredClusterRequest) Marshal() ([]byte, error)            { return m.MarshalVT() }
func (m *GetDiscoveredClusterRequest) Unmarshal(dAtA []byte) error         { return m.UnmarshalVT(dAtA) }

func (m *GetDiscoveredClusterResponse) Size() int                            { return m.SizeVT() }
func (m *GetDiscoveredClusterResponse) Clone() *GetDiscoveredClusterResponse { return m.CloneVT() }
func (m *GetDiscoveredClusterResponse) Marshal() ([]byte, error)             { return m.MarshalVT() }
func (m *GetDiscoveredClusterResponse) Unmarshal(dAtA []byte) error          { return m.UnmarshalVT(dAtA) }

func (m *ListDiscoveredClustersRequest) Size() int                             { return m.SizeVT() }
func (m *ListDiscoveredClustersRequest) Clone() *ListDiscoveredClustersRequest { return m.CloneVT() }
func (m *ListDiscoveredClustersRequest) Marshal() ([]byte, error)              { return m.MarshalVT() }
func (m *ListDiscoveredClustersRequest) Unmarshal(dAtA []byte) error           { return m.UnmarshalVT(dAtA) }

func (m *ListDiscoveredClustersResponse) Size() int                              { return m.SizeVT() }
func (m *ListDiscoveredClustersResponse) Clone() *ListDiscoveredClustersResponse { return m.CloneVT() }
func (m *ListDiscoveredClustersResponse) Marshal() ([]byte, error)               { return m.MarshalVT() }
func (m *ListDiscoveredClustersResponse) Unmarshal(dAtA []byte) error            { return m.UnmarshalVT(dAtA) }
