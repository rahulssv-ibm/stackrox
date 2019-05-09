// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/group/datastore (interfaces: DataStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
	reflect "reflect"
)

// MockDataStore is a mock of DataStore interface
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockDataStore) Add(arg0 context.Context, arg1 *storage.Group) error {
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockDataStoreMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockDataStore)(nil).Add), arg0, arg1)
}

// Get mocks base method
func (m *MockDataStore) Get(arg0 context.Context, arg1 *storage.GroupProperties) (*storage.Group, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*storage.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDataStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDataStore)(nil).Get), arg0, arg1)
}

// GetAll mocks base method
func (m *MockDataStore) GetAll(arg0 context.Context) ([]*storage.Group, error) {
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]*storage.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockDataStoreMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockDataStore)(nil).GetAll), arg0)
}

// Mutate mocks base method
func (m *MockDataStore) Mutate(arg0 context.Context, arg1, arg2, arg3 []*storage.Group) error {
	ret := m.ctrl.Call(m, "Mutate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mutate indicates an expected call of Mutate
func (mr *MockDataStoreMockRecorder) Mutate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mutate", reflect.TypeOf((*MockDataStore)(nil).Mutate), arg0, arg1, arg2, arg3)
}

// Remove mocks base method
func (m *MockDataStore) Remove(arg0 context.Context, arg1 *storage.GroupProperties) error {
	ret := m.ctrl.Call(m, "Remove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockDataStoreMockRecorder) Remove(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockDataStore)(nil).Remove), arg0, arg1)
}

// Update mocks base method
func (m *MockDataStore) Update(arg0 context.Context, arg1 *storage.Group) error {
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockDataStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDataStore)(nil).Update), arg0, arg1)
}

// Upsert mocks base method
func (m *MockDataStore) Upsert(arg0 context.Context, arg1 *storage.Group) error {
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert
func (mr *MockDataStoreMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockDataStore)(nil).Upsert), arg0, arg1)
}

// Walk mocks base method
func (m *MockDataStore) Walk(arg0 context.Context, arg1 string, arg2 map[string][]string) ([]*storage.Group, error) {
	ret := m.ctrl.Call(m, "Walk", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*storage.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Walk indicates an expected call of Walk
func (mr *MockDataStoreMockRecorder) Walk(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Walk", reflect.TypeOf((*MockDataStore)(nil).Walk), arg0, arg1, arg2)
}
