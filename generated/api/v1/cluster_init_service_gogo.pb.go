// Code generated by protoc-gen-go-gogo. DO NOT EDIT.

package v1

func (m *InitBundleMeta) Size() int                   { return m.SizeVT() }
func (m *InitBundleMeta) Clone() *InitBundleMeta      { return m.CloneVT() }
func (m *InitBundleMeta) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *InitBundleMeta) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *InitBundleMeta_ImpactedCluster) Size() int                              { return m.SizeVT() }
func (m *InitBundleMeta_ImpactedCluster) Clone() *InitBundleMeta_ImpactedCluster { return m.CloneVT() }
func (m *InitBundleMeta_ImpactedCluster) Marshal() ([]byte, error)               { return m.MarshalVT() }
func (m *InitBundleMeta_ImpactedCluster) Unmarshal(dAtA []byte) error            { return m.UnmarshalVT(dAtA) }

func (m *InitBundleGenResponse) Size() int                     { return m.SizeVT() }
func (m *InitBundleGenResponse) Clone() *InitBundleGenResponse { return m.CloneVT() }
func (m *InitBundleGenResponse) Marshal() ([]byte, error)      { return m.MarshalVT() }
func (m *InitBundleGenResponse) Unmarshal(dAtA []byte) error   { return m.UnmarshalVT(dAtA) }

func (m *GetCAConfigResponse) Size() int                   { return m.SizeVT() }
func (m *GetCAConfigResponse) Clone() *GetCAConfigResponse { return m.CloneVT() }
func (m *GetCAConfigResponse) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *GetCAConfigResponse) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *InitBundleMetasResponse) Size() int                       { return m.SizeVT() }
func (m *InitBundleMetasResponse) Clone() *InitBundleMetasResponse { return m.CloneVT() }
func (m *InitBundleMetasResponse) Marshal() ([]byte, error)        { return m.MarshalVT() }
func (m *InitBundleMetasResponse) Unmarshal(dAtA []byte) error     { return m.UnmarshalVT(dAtA) }

func (m *InitBundleGenRequest) Size() int                    { return m.SizeVT() }
func (m *InitBundleGenRequest) Clone() *InitBundleGenRequest { return m.CloneVT() }
func (m *InitBundleGenRequest) Marshal() ([]byte, error)     { return m.MarshalVT() }
func (m *InitBundleGenRequest) Unmarshal(dAtA []byte) error  { return m.UnmarshalVT(dAtA) }

func (m *InitBundleRevokeRequest) Size() int                       { return m.SizeVT() }
func (m *InitBundleRevokeRequest) Clone() *InitBundleRevokeRequest { return m.CloneVT() }
func (m *InitBundleRevokeRequest) Marshal() ([]byte, error)        { return m.MarshalVT() }
func (m *InitBundleRevokeRequest) Unmarshal(dAtA []byte) error     { return m.UnmarshalVT(dAtA) }

func (m *InitBundleRevokeResponse) Size() int                        { return m.SizeVT() }
func (m *InitBundleRevokeResponse) Clone() *InitBundleRevokeResponse { return m.CloneVT() }
func (m *InitBundleRevokeResponse) Marshal() ([]byte, error)         { return m.MarshalVT() }
func (m *InitBundleRevokeResponse) Unmarshal(dAtA []byte) error      { return m.UnmarshalVT(dAtA) }

func (m *InitBundleRevokeResponse_InitBundleRevocationError) Size() int { return m.SizeVT() }
func (m *InitBundleRevokeResponse_InitBundleRevocationError) Clone() *InitBundleRevokeResponse_InitBundleRevocationError {
	return m.CloneVT()
}
func (m *InitBundleRevokeResponse_InitBundleRevocationError) Marshal() ([]byte, error) {
	return m.MarshalVT()
}
func (m *InitBundleRevokeResponse_InitBundleRevocationError) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}
