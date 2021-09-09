// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetPolicy mocks base method
func (m *MockStore) GetPolicy(id string) (*storage.Policy, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicy", id)
	ret0, _ := ret[0].(*storage.Policy)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPolicy indicates an expected call of GetPolicy
func (mr *MockStoreMockRecorder) GetPolicy(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicy", reflect.TypeOf((*MockStore)(nil).GetPolicy), id)
}

// GetAllPolicies mocks base method
func (m *MockStore) GetAllPolicies() ([]*storage.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPolicies")
	ret0, _ := ret[0].([]*storage.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPolicies indicates an expected call of GetAllPolicies
func (mr *MockStoreMockRecorder) GetAllPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPolicies", reflect.TypeOf((*MockStore)(nil).GetAllPolicies))
}

// GetPolicies mocks base method
func (m *MockStore) GetPolicies(ids ...string) ([]*storage.Policy, []int, []error, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range ids {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPolicies", varargs...)
	ret0, _ := ret[0].([]*storage.Policy)
	ret1, _ := ret[1].([]int)
	ret2, _ := ret[2].([]error)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetPolicies indicates an expected call of GetPolicies
func (mr *MockStoreMockRecorder) GetPolicies(ids ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicies", reflect.TypeOf((*MockStore)(nil).GetPolicies), ids...)
}

// AddPolicy mocks base method
func (m *MockStore) AddPolicy(policy *storage.Policy, removePolicyTombstone bool) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPolicy", policy, removePolicyTombstone)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPolicy indicates an expected call of AddPolicy
func (mr *MockStoreMockRecorder) AddPolicy(policy, removePolicyTombstone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPolicy", reflect.TypeOf((*MockStore)(nil).AddPolicy), policy, removePolicyTombstone)
}

// UpdatePolicy mocks base method
func (m *MockStore) UpdatePolicy(arg0 *storage.Policy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePolicy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePolicy indicates an expected call of UpdatePolicy
func (mr *MockStoreMockRecorder) UpdatePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePolicy", reflect.TypeOf((*MockStore)(nil).UpdatePolicy), arg0)
}

// RemovePolicy mocks base method
func (m *MockStore) RemovePolicy(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePolicy", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePolicy indicates an expected call of RemovePolicy
func (mr *MockStoreMockRecorder) RemovePolicy(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePolicy", reflect.TypeOf((*MockStore)(nil).RemovePolicy), id)
}

// RenamePolicyCategory mocks base method
func (m *MockStore) RenamePolicyCategory(request *v1.RenamePolicyCategoryRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenamePolicyCategory", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenamePolicyCategory indicates an expected call of RenamePolicyCategory
func (mr *MockStoreMockRecorder) RenamePolicyCategory(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenamePolicyCategory", reflect.TypeOf((*MockStore)(nil).RenamePolicyCategory), request)
}

// DeletePolicyCategory mocks base method
func (m *MockStore) DeletePolicyCategory(request *v1.DeletePolicyCategoryRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePolicyCategory", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePolicyCategory indicates an expected call of DeletePolicyCategory
func (mr *MockStoreMockRecorder) DeletePolicyCategory(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePolicyCategory", reflect.TypeOf((*MockStore)(nil).DeletePolicyCategory), request)
}
