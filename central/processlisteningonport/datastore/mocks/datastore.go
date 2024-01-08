// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/datastore.go -source datastore.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	datastore "github.com/stackrox/rox/central/processlisteningonport/datastore"
	storage "github.com/stackrox/rox/generated/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockDataStore is a mock of DataStore interface.
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore.
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance.
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// AddProcessListeningOnPort mocks base method.
func (m *MockDataStore) AddProcessListeningOnPort(arg0 context.Context, arg1 ...*storage.ProcessListeningOnPortFromSensor) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddProcessListeningOnPort", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProcessListeningOnPort indicates an expected call of AddProcessListeningOnPort.
func (mr *MockDataStoreMockRecorder) AddProcessListeningOnPort(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProcessListeningOnPort", reflect.TypeOf((*MockDataStore)(nil).AddProcessListeningOnPort), varargs...)
}

// GetProcessListeningOnPort mocks base method.
func (m *MockDataStore) GetProcessListeningOnPort(ctx context.Context, deployment string) ([]*storage.ProcessListeningOnPort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProcessListeningOnPort", ctx, deployment)
	ret0, _ := ret[0].([]*storage.ProcessListeningOnPort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProcessListeningOnPort indicates an expected call of GetProcessListeningOnPort.
func (mr *MockDataStoreMockRecorder) GetProcessListeningOnPort(ctx, deployment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProcessListeningOnPort", reflect.TypeOf((*MockDataStore)(nil).GetProcessListeningOnPort), ctx, deployment)
}

// Lock mocks base method.
func (m *MockDataStore) Lock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Lock")
}

// Lock indicates an expected call of Lock.
func (mr *MockDataStoreMockRecorder) Lock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockDataStore)(nil).Lock))
}

// RemovePlopsByPod mocks base method.
func (m *MockDataStore) RemovePlopsByPod(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePlopsByPod", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePlopsByPod indicates an expected call of RemovePlopsByPod.
func (mr *MockDataStoreMockRecorder) RemovePlopsByPod(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePlopsByPod", reflect.TypeOf((*MockDataStore)(nil).RemovePlopsByPod), ctx, id)
}

// RemoveProcessListeningOnPort mocks base method.
func (m *MockDataStore) RemoveProcessListeningOnPort(ctx context.Context, ids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProcessListeningOnPort", ctx, ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProcessListeningOnPort indicates an expected call of RemoveProcessListeningOnPort.
func (mr *MockDataStoreMockRecorder) RemoveProcessListeningOnPort(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProcessListeningOnPort", reflect.TypeOf((*MockDataStore)(nil).RemoveProcessListeningOnPort), ctx, ids)
}

// Unlock mocks base method.
func (m *MockDataStore) Unlock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unlock")
}

// Unlock indicates an expected call of Unlock.
func (mr *MockDataStoreMockRecorder) Unlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockDataStore)(nil).Unlock))
}

// WalkAll mocks base method.
func (m *MockDataStore) WalkAll(ctx context.Context, fn datastore.WalkFn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalkAll", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalkAll indicates an expected call of WalkAll.
func (mr *MockDataStoreMockRecorder) WalkAll(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalkAll", reflect.TypeOf((*MockDataStore)(nil).WalkAll), ctx, fn)
}
