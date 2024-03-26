// Code generated by protoc-gen-go-gogo. DO NOT EDIT.

package storage

func (m *Deployment) Size() int                   { return m.SizeVT() }
func (m *Deployment) Clone() *Deployment          { return m.CloneVT() }
func (m *Deployment) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *Deployment) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ContainerImage) Size() int                   { return m.SizeVT() }
func (m *ContainerImage) Clone() *ContainerImage      { return m.CloneVT() }
func (m *ContainerImage) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ContainerImage) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *Container) Size() int                   { return m.SizeVT() }
func (m *Container) Clone() *Container           { return m.CloneVT() }
func (m *Container) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *Container) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *Resources) Size() int                   { return m.SizeVT() }
func (m *Resources) Clone() *Resources           { return m.CloneVT() }
func (m *Resources) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *Resources) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *Volume) Size() int                   { return m.SizeVT() }
func (m *Volume) Clone() *Volume              { return m.CloneVT() }
func (m *Volume) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *Volume) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *LivenessProbe) Size() int                   { return m.SizeVT() }
func (m *LivenessProbe) Clone() *LivenessProbe       { return m.CloneVT() }
func (m *LivenessProbe) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *LivenessProbe) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ReadinessProbe) Size() int                   { return m.SizeVT() }
func (m *ReadinessProbe) Clone() *ReadinessProbe      { return m.CloneVT() }
func (m *ReadinessProbe) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ReadinessProbe) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *Pod) Size() int                   { return m.SizeVT() }
func (m *Pod) Clone() *Pod                 { return m.CloneVT() }
func (m *Pod) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *Pod) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *Pod_ContainerInstanceList) Size() int                         { return m.SizeVT() }
func (m *Pod_ContainerInstanceList) Clone() *Pod_ContainerInstanceList { return m.CloneVT() }
func (m *Pod_ContainerInstanceList) Marshal() ([]byte, error)          { return m.MarshalVT() }
func (m *Pod_ContainerInstanceList) Unmarshal(dAtA []byte) error       { return m.UnmarshalVT(dAtA) }

func (m *ContainerInstance) Size() int                   { return m.SizeVT() }
func (m *ContainerInstance) Clone() *ContainerInstance   { return m.CloneVT() }
func (m *ContainerInstance) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ContainerInstance) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ContainerInstanceID) Size() int                   { return m.SizeVT() }
func (m *ContainerInstanceID) Clone() *ContainerInstanceID { return m.CloneVT() }
func (m *ContainerInstanceID) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ContainerInstanceID) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *EmbeddedSecret) Size() int                   { return m.SizeVT() }
func (m *EmbeddedSecret) Clone() *EmbeddedSecret      { return m.CloneVT() }
func (m *EmbeddedSecret) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *EmbeddedSecret) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *PortConfig) Size() int                   { return m.SizeVT() }
func (m *PortConfig) Clone() *PortConfig          { return m.CloneVT() }
func (m *PortConfig) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *PortConfig) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *PortConfig_ExposureInfo) Size() int                       { return m.SizeVT() }
func (m *PortConfig_ExposureInfo) Clone() *PortConfig_ExposureInfo { return m.CloneVT() }
func (m *PortConfig_ExposureInfo) Marshal() ([]byte, error)        { return m.MarshalVT() }
func (m *PortConfig_ExposureInfo) Unmarshal(dAtA []byte) error     { return m.UnmarshalVT(dAtA) }

func (m *ContainerConfig) Size() int                   { return m.SizeVT() }
func (m *ContainerConfig) Clone() *ContainerConfig     { return m.CloneVT() }
func (m *ContainerConfig) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ContainerConfig) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ContainerConfig_EnvironmentConfig) Size() int { return m.SizeVT() }
func (m *ContainerConfig_EnvironmentConfig) Clone() *ContainerConfig_EnvironmentConfig {
	return m.CloneVT()
}
func (m *ContainerConfig_EnvironmentConfig) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ContainerConfig_EnvironmentConfig) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *SecurityContext) Size() int                   { return m.SizeVT() }
func (m *SecurityContext) Clone() *SecurityContext     { return m.CloneVT() }
func (m *SecurityContext) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *SecurityContext) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *SecurityContext_SELinux) Size() int                       { return m.SizeVT() }
func (m *SecurityContext_SELinux) Clone() *SecurityContext_SELinux { return m.CloneVT() }
func (m *SecurityContext_SELinux) Marshal() ([]byte, error)        { return m.MarshalVT() }
func (m *SecurityContext_SELinux) Unmarshal(dAtA []byte) error     { return m.UnmarshalVT(dAtA) }

func (m *SecurityContext_SeccompProfile) Size() int                              { return m.SizeVT() }
func (m *SecurityContext_SeccompProfile) Clone() *SecurityContext_SeccompProfile { return m.CloneVT() }
func (m *SecurityContext_SeccompProfile) Marshal() ([]byte, error)               { return m.MarshalVT() }
func (m *SecurityContext_SeccompProfile) Unmarshal(dAtA []byte) error            { return m.UnmarshalVT(dAtA) }

func (m *ListDeployment) Size() int                   { return m.SizeVT() }
func (m *ListDeployment) Clone() *ListDeployment      { return m.CloneVT() }
func (m *ListDeployment) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ListDeployment) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }
