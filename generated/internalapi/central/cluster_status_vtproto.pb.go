// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: internalapi/central/cluster_status.proto

package central

import (
	fmt "fmt"
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	storage "github.com/stackrox/rox/generated/storage"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *DeploymentEnvironmentUpdate) CloneVT() *DeploymentEnvironmentUpdate {
	if m == nil {
		return (*DeploymentEnvironmentUpdate)(nil)
	}
	r := new(DeploymentEnvironmentUpdate)
	if rhs := m.Environments; rhs != nil {
		tmpContainer := make([]string, len(rhs))
		copy(tmpContainer, rhs)
		r.Environments = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DeploymentEnvironmentUpdate) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ClusterStatusUpdate) CloneVT() *ClusterStatusUpdate {
	if m == nil {
		return (*ClusterStatusUpdate)(nil)
	}
	r := new(ClusterStatusUpdate)
	if m.Msg != nil {
		r.Msg = m.Msg.(interface {
			CloneVT() isClusterStatusUpdate_Msg
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ClusterStatusUpdate) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ClusterStatusUpdate_Status) CloneVT() isClusterStatusUpdate_Msg {
	if m == nil {
		return (*ClusterStatusUpdate_Status)(nil)
	}
	r := new(ClusterStatusUpdate_Status)
	if rhs := m.Status; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface{ CloneVT() *storage.ClusterStatus }); ok {
			r.Status = vtpb.CloneVT()
		} else {
			r.Status = proto.Clone(rhs).(*storage.ClusterStatus)
		}
	}
	return r
}

func (m *ClusterStatusUpdate_DeploymentEnvUpdate) CloneVT() isClusterStatusUpdate_Msg {
	if m == nil {
		return (*ClusterStatusUpdate_DeploymentEnvUpdate)(nil)
	}
	r := new(ClusterStatusUpdate_DeploymentEnvUpdate)
	r.DeploymentEnvUpdate = m.DeploymentEnvUpdate.CloneVT()
	return r
}

func (m *RawClusterHealthInfo) CloneVT() *RawClusterHealthInfo {
	if m == nil {
		return (*RawClusterHealthInfo)(nil)
	}
	r := new(RawClusterHealthInfo)
	if rhs := m.CollectorHealthInfo; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface {
			CloneVT() *storage.CollectorHealthInfo
		}); ok {
			r.CollectorHealthInfo = vtpb.CloneVT()
		} else {
			r.CollectorHealthInfo = proto.Clone(rhs).(*storage.CollectorHealthInfo)
		}
	}
	if rhs := m.AdmissionControlHealthInfo; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface {
			CloneVT() *storage.AdmissionControlHealthInfo
		}); ok {
			r.AdmissionControlHealthInfo = vtpb.CloneVT()
		} else {
			r.AdmissionControlHealthInfo = proto.Clone(rhs).(*storage.AdmissionControlHealthInfo)
		}
	}
	if rhs := m.ScannerHealthInfo; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface {
			CloneVT() *storage.ScannerHealthInfo
		}); ok {
			r.ScannerHealthInfo = vtpb.CloneVT()
		} else {
			r.ScannerHealthInfo = proto.Clone(rhs).(*storage.ScannerHealthInfo)
		}
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *RawClusterHealthInfo) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ClusterHealthResponse) CloneVT() *ClusterHealthResponse {
	if m == nil {
		return (*ClusterHealthResponse)(nil)
	}
	r := new(ClusterHealthResponse)
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ClusterHealthResponse) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *DeploymentEnvironmentUpdate) EqualVT(that *DeploymentEnvironmentUpdate) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Environments) != len(that.Environments) {
		return false
	}
	for i, vx := range this.Environments {
		vy := that.Environments[i]
		if vx != vy {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeploymentEnvironmentUpdate) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeploymentEnvironmentUpdate)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ClusterStatusUpdate) EqualVT(that *ClusterStatusUpdate) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Msg == nil && that.Msg != nil {
		return false
	} else if this.Msg != nil {
		if that.Msg == nil {
			return false
		}
		if !this.Msg.(interface {
			EqualVT(isClusterStatusUpdate_Msg) bool
		}).EqualVT(that.Msg) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ClusterStatusUpdate) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ClusterStatusUpdate)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ClusterStatusUpdate_Status) EqualVT(thatIface isClusterStatusUpdate_Msg) bool {
	that, ok := thatIface.(*ClusterStatusUpdate_Status)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Status, that.Status; p != q {
		if p == nil {
			p = &storage.ClusterStatus{}
		}
		if q == nil {
			q = &storage.ClusterStatus{}
		}
		if equal, ok := interface{}(p).(interface {
			EqualVT(*storage.ClusterStatus) bool
		}); ok {
			if !equal.EqualVT(q) {
				return false
			}
		} else if !proto.Equal(p, q) {
			return false
		}
	}
	return true
}

func (this *ClusterStatusUpdate_DeploymentEnvUpdate) EqualVT(thatIface isClusterStatusUpdate_Msg) bool {
	that, ok := thatIface.(*ClusterStatusUpdate_DeploymentEnvUpdate)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.DeploymentEnvUpdate, that.DeploymentEnvUpdate; p != q {
		if p == nil {
			p = &DeploymentEnvironmentUpdate{}
		}
		if q == nil {
			q = &DeploymentEnvironmentUpdate{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *RawClusterHealthInfo) EqualVT(that *RawClusterHealthInfo) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if equal, ok := interface{}(this.CollectorHealthInfo).(interface {
		EqualVT(*storage.CollectorHealthInfo) bool
	}); ok {
		if !equal.EqualVT(that.CollectorHealthInfo) {
			return false
		}
	} else if !proto.Equal(this.CollectorHealthInfo, that.CollectorHealthInfo) {
		return false
	}
	if equal, ok := interface{}(this.AdmissionControlHealthInfo).(interface {
		EqualVT(*storage.AdmissionControlHealthInfo) bool
	}); ok {
		if !equal.EqualVT(that.AdmissionControlHealthInfo) {
			return false
		}
	} else if !proto.Equal(this.AdmissionControlHealthInfo, that.AdmissionControlHealthInfo) {
		return false
	}
	if equal, ok := interface{}(this.ScannerHealthInfo).(interface {
		EqualVT(*storage.ScannerHealthInfo) bool
	}); ok {
		if !equal.EqualVT(that.ScannerHealthInfo) {
			return false
		}
	} else if !proto.Equal(this.ScannerHealthInfo, that.ScannerHealthInfo) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *RawClusterHealthInfo) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*RawClusterHealthInfo)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ClusterHealthResponse) EqualVT(that *ClusterHealthResponse) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ClusterHealthResponse) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ClusterHealthResponse)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *DeploymentEnvironmentUpdate) MarshalVT() (dAtA []byte, err error) {
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

func (m *DeploymentEnvironmentUpdate) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *DeploymentEnvironmentUpdate) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.Environments) > 0 {
		for iNdEx := len(m.Environments) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Environments[iNdEx])
			copy(dAtA[i:], m.Environments[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Environments[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ClusterStatusUpdate) MarshalVT() (dAtA []byte, err error) {
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

func (m *ClusterStatusUpdate) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ClusterStatusUpdate) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if vtmsg, ok := m.Msg.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *ClusterStatusUpdate_Status) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ClusterStatusUpdate_Status) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Status != nil {
		if vtmsg, ok := interface{}(m.Status).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Status)
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
func (m *ClusterStatusUpdate_DeploymentEnvUpdate) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ClusterStatusUpdate_DeploymentEnvUpdate) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DeploymentEnvUpdate != nil {
		size, err := m.DeploymentEnvUpdate.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *RawClusterHealthInfo) MarshalVT() (dAtA []byte, err error) {
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

func (m *RawClusterHealthInfo) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *RawClusterHealthInfo) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.ScannerHealthInfo != nil {
		if vtmsg, ok := interface{}(m.ScannerHealthInfo).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.ScannerHealthInfo)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.AdmissionControlHealthInfo != nil {
		if vtmsg, ok := interface{}(m.AdmissionControlHealthInfo).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.AdmissionControlHealthInfo)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.CollectorHealthInfo != nil {
		if vtmsg, ok := interface{}(m.CollectorHealthInfo).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.CollectorHealthInfo)
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

func (m *ClusterHealthResponse) MarshalVT() (dAtA []byte, err error) {
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

func (m *ClusterHealthResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ClusterHealthResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	return len(dAtA) - i, nil
}

func (m *DeploymentEnvironmentUpdate) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Environments) > 0 {
		for _, s := range m.Environments {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *ClusterStatusUpdate) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Msg.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *ClusterStatusUpdate_Status) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != nil {
		if size, ok := interface{}(m.Status).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Status)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	return n
}
func (m *ClusterStatusUpdate_DeploymentEnvUpdate) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DeploymentEnvUpdate != nil {
		l = m.DeploymentEnvUpdate.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	return n
}
func (m *RawClusterHealthInfo) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CollectorHealthInfo != nil {
		if size, ok := interface{}(m.CollectorHealthInfo).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.CollectorHealthInfo)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.AdmissionControlHealthInfo != nil {
		if size, ok := interface{}(m.AdmissionControlHealthInfo).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.AdmissionControlHealthInfo)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.ScannerHealthInfo != nil {
		if size, ok := interface{}(m.ScannerHealthInfo).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.ScannerHealthInfo)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ClusterHealthResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *DeploymentEnvironmentUpdate) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: DeploymentEnvironmentUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeploymentEnvironmentUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Environments", wireType)
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
			m.Environments = append(m.Environments, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
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
func (m *ClusterStatusUpdate) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: ClusterStatusUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterStatusUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
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
			if oneof, ok := m.Msg.(*ClusterStatusUpdate_Status); ok {
				if unmarshal, ok := interface{}(oneof.Status).(interface {
					UnmarshalVT([]byte) error
				}); ok {
					if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
						return err
					}
				} else {
					if err := proto.Unmarshal(dAtA[iNdEx:postIndex], oneof.Status); err != nil {
						return err
					}
				}
			} else {
				v := &storage.ClusterStatus{}
				if unmarshal, ok := interface{}(v).(interface {
					UnmarshalVT([]byte) error
				}); ok {
					if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
						return err
					}
				} else {
					if err := proto.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
						return err
					}
				}
				m.Msg = &ClusterStatusUpdate_Status{Status: v}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeploymentEnvUpdate", wireType)
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
			if oneof, ok := m.Msg.(*ClusterStatusUpdate_DeploymentEnvUpdate); ok {
				if err := oneof.DeploymentEnvUpdate.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &DeploymentEnvironmentUpdate{}
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Msg = &ClusterStatusUpdate_DeploymentEnvUpdate{DeploymentEnvUpdate: v}
			}
			iNdEx = postIndex
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
func (m *RawClusterHealthInfo) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: RawClusterHealthInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RawClusterHealthInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectorHealthInfo", wireType)
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
			if m.CollectorHealthInfo == nil {
				m.CollectorHealthInfo = &storage.CollectorHealthInfo{}
			}
			if unmarshal, ok := interface{}(m.CollectorHealthInfo).(interface {
				UnmarshalVT([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.CollectorHealthInfo); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdmissionControlHealthInfo", wireType)
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
			if m.AdmissionControlHealthInfo == nil {
				m.AdmissionControlHealthInfo = &storage.AdmissionControlHealthInfo{}
			}
			if unmarshal, ok := interface{}(m.AdmissionControlHealthInfo).(interface {
				UnmarshalVT([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.AdmissionControlHealthInfo); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScannerHealthInfo", wireType)
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
			if m.ScannerHealthInfo == nil {
				m.ScannerHealthInfo = &storage.ScannerHealthInfo{}
			}
			if unmarshal, ok := interface{}(m.ScannerHealthInfo).(interface {
				UnmarshalVT([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.ScannerHealthInfo); err != nil {
					return err
				}
			}
			iNdEx = postIndex
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
func (m *ClusterHealthResponse) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: ClusterHealthResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterHealthResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
