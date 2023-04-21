// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/complianceoperator/manager (interfaces: Manager)

// Package mock_manager is a generated GoMock package.
package mock_manager

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// AddProfile mocks base method.
func (m *MockManager) AddProfile(arg0 *storage.ComplianceOperatorProfile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProfile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProfile indicates an expected call of AddProfile.
func (mr *MockManagerMockRecorder) AddProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProfile", reflect.TypeOf((*MockManager)(nil).AddProfile), arg0)
}

// AddRule mocks base method.
func (m *MockManager) AddRule(arg0 *storage.ComplianceOperatorRule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRule", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRule indicates an expected call of AddRule.
func (mr *MockManagerMockRecorder) AddRule(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRule", reflect.TypeOf((*MockManager)(nil).AddRule), arg0)
}

// AddScan mocks base method.
func (m *MockManager) AddScan(arg0 *storage.ComplianceOperatorScan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddScan", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddScan indicates an expected call of AddScan.
func (mr *MockManagerMockRecorder) AddScan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddScan", reflect.TypeOf((*MockManager)(nil).AddScan), arg0)
}

// DeleteProfile mocks base method.
func (m *MockManager) DeleteProfile(arg0 *storage.ComplianceOperatorProfile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProfile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProfile indicates an expected call of DeleteProfile.
func (mr *MockManagerMockRecorder) DeleteProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProfile", reflect.TypeOf((*MockManager)(nil).DeleteProfile), arg0)
}

// DeleteRule mocks base method.
func (m *MockManager) DeleteRule(arg0 *storage.ComplianceOperatorRule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRule", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRule indicates an expected call of DeleteRule.
func (mr *MockManagerMockRecorder) DeleteRule(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRule", reflect.TypeOf((*MockManager)(nil).DeleteRule), arg0)
}

// DeleteScan mocks base method.
func (m *MockManager) DeleteScan(arg0 *storage.ComplianceOperatorScan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScan", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteScan indicates an expected call of DeleteScan.
func (mr *MockManagerMockRecorder) DeleteScan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScan", reflect.TypeOf((*MockManager)(nil).DeleteScan), arg0)
}

// GetMachineConfigs mocks base method.
func (m *MockManager) GetMachineConfigs(arg0 string) (map[string][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineConfigs", arg0)
	ret0, _ := ret[0].(map[string][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineConfigs indicates an expected call of GetMachineConfigs.
func (mr *MockManagerMockRecorder) GetMachineConfigs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineConfigs", reflect.TypeOf((*MockManager)(nil).GetMachineConfigs), arg0)
}

// IsStandardActive mocks base method.
func (m *MockManager) IsStandardActive(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsStandardActive", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsStandardActive indicates an expected call of IsStandardActive.
func (mr *MockManagerMockRecorder) IsStandardActive(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsStandardActive", reflect.TypeOf((*MockManager)(nil).IsStandardActive), arg0)
}

// IsStandardActiveForCluster mocks base method.
func (m *MockManager) IsStandardActiveForCluster(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsStandardActiveForCluster", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsStandardActiveForCluster indicates an expected call of IsStandardActiveForCluster.
func (mr *MockManagerMockRecorder) IsStandardActiveForCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsStandardActiveForCluster", reflect.TypeOf((*MockManager)(nil).IsStandardActiveForCluster), arg0, arg1)
}
