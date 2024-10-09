// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: internalapi/central/hello.proto

package central

import (
	fmt "fmt"
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	storage "github.com/stackrox/rox/generated/storage"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *HelmManagedConfigInit) CloneVT() *HelmManagedConfigInit {
	if m == nil {
		return (*HelmManagedConfigInit)(nil)
	}
	r := new(HelmManagedConfigInit)
	r.ClusterName = m.ClusterName
	r.ClusterId = m.ClusterId
	r.ManagedBy = m.ManagedBy
	if rhs := m.ClusterConfig; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface {
			CloneVT() *storage.CompleteClusterConfig
		}); ok {
			r.ClusterConfig = vtpb.CloneVT()
		} else {
			r.ClusterConfig = proto.Clone(rhs).(*storage.CompleteClusterConfig)
		}
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *HelmManagedConfigInit) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *SensorHello) CloneVT() *SensorHello {
	if m == nil {
		return (*SensorHello)(nil)
	}
	r := new(SensorHello)
	r.SensorVersion = m.SensorVersion
	r.HelmManagedConfigInit = m.HelmManagedConfigInit.CloneVT()
	r.PolicyVersion = m.PolicyVersion
	r.SensorState = m.SensorState
	r.RequestDeduperState = m.RequestDeduperState
	if rhs := m.Capabilities; rhs != nil {
		tmpContainer := make([]string, len(rhs))
		copy(tmpContainer, rhs)
		r.Capabilities = tmpContainer
	}
	if rhs := m.DeploymentIdentification; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface {
			CloneVT() *storage.SensorDeploymentIdentification
		}); ok {
			r.DeploymentIdentification = vtpb.CloneVT()
		} else {
			r.DeploymentIdentification = proto.Clone(rhs).(*storage.SensorDeploymentIdentification)
		}
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *SensorHello) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *CentralHello) CloneVT() *CentralHello {
	if m == nil {
		return (*CentralHello)(nil)
	}
	r := new(CentralHello)
	r.ClusterId = m.ClusterId
	r.ManagedCentral = m.ManagedCentral
	r.CentralId = m.CentralId
	r.SendDeduperState = m.SendDeduperState
	if rhs := m.CertBundle; rhs != nil {
		tmpContainer := make(map[string]string, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v
		}
		r.CertBundle = tmpContainer
	}
	if rhs := m.Capabilities; rhs != nil {
		tmpContainer := make([]string, len(rhs))
		copy(tmpContainer, rhs)
		r.Capabilities = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *CentralHello) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *HelmManagedConfigInit) EqualVT(that *HelmManagedConfigInit) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if equal, ok := interface{}(this.ClusterConfig).(interface {
		EqualVT(*storage.CompleteClusterConfig) bool
	}); ok {
		if !equal.EqualVT(that.ClusterConfig) {
			return false
		}
	} else if !proto.Equal(this.ClusterConfig, that.ClusterConfig) {
		return false
	}
	if this.ClusterName != that.ClusterName {
		return false
	}
	if this.ClusterId != that.ClusterId {
		return false
	}
	if this.ManagedBy != that.ManagedBy {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *HelmManagedConfigInit) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*HelmManagedConfigInit)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *SensorHello) EqualVT(that *SensorHello) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.SensorVersion != that.SensorVersion {
		return false
	}
	if len(this.Capabilities) != len(that.Capabilities) {
		return false
	}
	for i, vx := range this.Capabilities {
		vy := that.Capabilities[i]
		if vx != vy {
			return false
		}
	}
	if !this.HelmManagedConfigInit.EqualVT(that.HelmManagedConfigInit) {
		return false
	}
	if this.PolicyVersion != that.PolicyVersion {
		return false
	}
	if equal, ok := interface{}(this.DeploymentIdentification).(interface {
		EqualVT(*storage.SensorDeploymentIdentification) bool
	}); ok {
		if !equal.EqualVT(that.DeploymentIdentification) {
			return false
		}
	} else if !proto.Equal(this.DeploymentIdentification, that.DeploymentIdentification) {
		return false
	}
	if this.SensorState != that.SensorState {
		return false
	}
	if this.RequestDeduperState != that.RequestDeduperState {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *SensorHello) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*SensorHello)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *CentralHello) EqualVT(that *CentralHello) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ClusterId != that.ClusterId {
		return false
	}
	if len(this.CertBundle) != len(that.CertBundle) {
		return false
	}
	for i, vx := range this.CertBundle {
		vy, ok := that.CertBundle[i]
		if !ok {
			return false
		}
		if vx != vy {
			return false
		}
	}
	if this.ManagedCentral != that.ManagedCentral {
		return false
	}
	if this.CentralId != that.CentralId {
		return false
	}
	if len(this.Capabilities) != len(that.Capabilities) {
		return false
	}
	for i, vx := range this.Capabilities {
		vy := that.Capabilities[i]
		if vx != vy {
			return false
		}
	}
	if this.SendDeduperState != that.SendDeduperState {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CentralHello) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*CentralHello)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *HelmManagedConfigInit) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelmManagedConfigInit) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *HelmManagedConfigInit) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.ManagedBy != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.ManagedBy))
		i--
		dAtA[i] = 0x28
	}
	if len(m.ClusterId) > 0 {
		i -= len(m.ClusterId)
		copy(dAtA[i:], m.ClusterId)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ClusterId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ClusterName) > 0 {
		i -= len(m.ClusterName)
		copy(dAtA[i:], m.ClusterName)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ClusterName)))
		i--
		dAtA[i] = 0x12
	}
	if m.ClusterConfig != nil {
		if vtmsg, ok := interface{}(m.ClusterConfig).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.ClusterConfig)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SensorHello) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SensorHello) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *SensorHello) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.RequestDeduperState {
		i--
		if m.RequestDeduperState {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.SensorState != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.SensorState))
		i--
		dAtA[i] = 0x30
	}
	if m.DeploymentIdentification != nil {
		if vtmsg, ok := interface{}(m.DeploymentIdentification).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.DeploymentIdentification)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PolicyVersion) > 0 {
		i -= len(m.PolicyVersion)
		copy(dAtA[i:], m.PolicyVersion)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.PolicyVersion)))
		i--
		dAtA[i] = 0x22
	}
	if m.HelmManagedConfigInit != nil {
		size, err := m.HelmManagedConfigInit.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Capabilities) > 0 {
		for iNdEx := len(m.Capabilities) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Capabilities[iNdEx])
			copy(dAtA[i:], m.Capabilities[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Capabilities[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.SensorVersion) > 0 {
		i -= len(m.SensorVersion)
		copy(dAtA[i:], m.SensorVersion)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.SensorVersion)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CentralHello) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CentralHello) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *CentralHello) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.SendDeduperState {
		i--
		if m.SendDeduperState {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Capabilities) > 0 {
		for iNdEx := len(m.Capabilities) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Capabilities[iNdEx])
			copy(dAtA[i:], m.Capabilities[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Capabilities[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.CentralId) > 0 {
		i -= len(m.CentralId)
		copy(dAtA[i:], m.CentralId)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.CentralId)))
		i--
		dAtA[i] = 0x22
	}
	if m.ManagedCentral {
		i--
		if m.ManagedCentral {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.CertBundle) > 0 {
		for k := range m.CertBundle {
			v := m.CertBundle[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = protohelpers.EncodeVarint(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ClusterId) > 0 {
		i -= len(m.ClusterId)
		copy(dAtA[i:], m.ClusterId)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ClusterId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HelmManagedConfigInit) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClusterConfig != nil {
		if size, ok := interface{}(m.ClusterConfig).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.ClusterConfig)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.ClusterName)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.ClusterId)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.ManagedBy != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.ManagedBy))
	}
	n += len(m.unknownFields)
	return n
}

func (m *SensorHello) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SensorVersion)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.Capabilities) > 0 {
		for _, s := range m.Capabilities {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.HelmManagedConfigInit != nil {
		l = m.HelmManagedConfigInit.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.PolicyVersion)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.DeploymentIdentification != nil {
		if size, ok := interface{}(m.DeploymentIdentification).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.DeploymentIdentification)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.SensorState != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.SensorState))
	}
	if m.RequestDeduperState {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *CentralHello) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClusterId)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.CertBundle) > 0 {
		for k, v := range m.CertBundle {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + protohelpers.SizeOfVarint(uint64(len(k))) + 1 + len(v) + protohelpers.SizeOfVarint(uint64(len(v)))
			n += mapEntrySize + 1 + protohelpers.SizeOfVarint(uint64(mapEntrySize))
		}
	}
	if m.ManagedCentral {
		n += 2
	}
	l = len(m.CentralId)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.Capabilities) > 0 {
		for _, s := range m.Capabilities {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.SendDeduperState {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *HelmManagedConfigInit) UnmarshalVTUnsafe(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HelmManagedConfigInit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelmManagedConfigInit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClusterConfig == nil {
				m.ClusterConfig = &storage.CompleteClusterConfig{}
			}
			if unmarshal, ok := interface{}(m.ClusterConfig).(interface {
				UnmarshalVTUnsafe([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVTUnsafe(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.ClusterConfig); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var stringValue string
			if intStringLen > 0 {
				stringValue = unsafe.String(&dAtA[iNdEx], intStringLen)
			}
			m.ClusterName = stringValue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var stringValue string
			if intStringLen > 0 {
				stringValue = unsafe.String(&dAtA[iNdEx], intStringLen)
			}
			m.ClusterId = stringValue
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ManagedBy", wireType)
			}
			m.ManagedBy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ManagedBy |= storage.ManagerType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SensorHello) UnmarshalVTUnsafe(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SensorHello: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SensorHello: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SensorVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var stringValue string
			if intStringLen > 0 {
				stringValue = unsafe.String(&dAtA[iNdEx], intStringLen)
			}
			m.SensorVersion = stringValue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Capabilities", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var stringValue string
			if intStringLen > 0 {
				stringValue = unsafe.String(&dAtA[iNdEx], intStringLen)
			}
			m.Capabilities = append(m.Capabilities, stringValue)
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HelmManagedConfigInit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HelmManagedConfigInit == nil {
				m.HelmManagedConfigInit = &HelmManagedConfigInit{}
			}
			if err := m.HelmManagedConfigInit.UnmarshalVTUnsafe(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PolicyVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var stringValue string
			if intStringLen > 0 {
				stringValue = unsafe.String(&dAtA[iNdEx], intStringLen)
			}
			m.PolicyVersion = stringValue
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeploymentIdentification", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeploymentIdentification == nil {
				m.DeploymentIdentification = &storage.SensorDeploymentIdentification{}
			}
			if unmarshal, ok := interface{}(m.DeploymentIdentification).(interface {
				UnmarshalVTUnsafe([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVTUnsafe(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.DeploymentIdentification); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SensorState", wireType)
			}
			m.SensorState = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SensorState |= SensorHello_SensorState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestDeduperState", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RequestDeduperState = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CentralHello) UnmarshalVTUnsafe(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CentralHello: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CentralHello: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var stringValue string
			if intStringLen > 0 {
				stringValue = unsafe.String(&dAtA[iNdEx], intStringLen)
			}
			m.ClusterId = stringValue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertBundle", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CertBundle == nil {
				m.CertBundle = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protohelpers.ErrIntOverflow
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protohelpers.ErrIntOverflow
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return protohelpers.ErrInvalidLength
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return protohelpers.ErrInvalidLength
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					if intStringLenmapkey == 0 {
						mapkey = ""
					} else {
						mapkey = unsafe.String(&dAtA[iNdEx], intStringLenmapkey)
					}
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protohelpers.ErrIntOverflow
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return protohelpers.ErrInvalidLength
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return protohelpers.ErrInvalidLength
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					if intStringLenmapvalue == 0 {
						mapvalue = ""
					} else {
						mapvalue = unsafe.String(&dAtA[iNdEx], intStringLenmapvalue)
					}
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := protohelpers.Skip(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return protohelpers.ErrInvalidLength
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.CertBundle[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ManagedCentral", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ManagedCentral = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CentralId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var stringValue string
			if intStringLen > 0 {
				stringValue = unsafe.String(&dAtA[iNdEx], intStringLen)
			}
			m.CentralId = stringValue
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Capabilities", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var stringValue string
			if intStringLen > 0 {
				stringValue = unsafe.String(&dAtA[iNdEx], intStringLen)
			}
			m.Capabilities = append(m.Capabilities, stringValue)
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendDeduperState", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SendDeduperState = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
