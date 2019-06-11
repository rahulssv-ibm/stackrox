// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/processwhitelistresults/datastore/internal/store (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
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

// DeleteWhitelistResults mocks base method
func (m *MockStore) DeleteWhitelistResults(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteWhitelistResults", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWhitelistResults indicates an expected call of DeleteWhitelistResults
func (mr *MockStoreMockRecorder) DeleteWhitelistResults(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWhitelistResults", reflect.TypeOf((*MockStore)(nil).DeleteWhitelistResults), arg0)
}

// GetWhitelistResults mocks base method
func (m *MockStore) GetWhitelistResults(arg0 string) (*storage.ProcessWhitelistResults, error) {
	ret := m.ctrl.Call(m, "GetWhitelistResults", arg0)
	ret0, _ := ret[0].(*storage.ProcessWhitelistResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWhitelistResults indicates an expected call of GetWhitelistResults
func (mr *MockStoreMockRecorder) GetWhitelistResults(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWhitelistResults", reflect.TypeOf((*MockStore)(nil).GetWhitelistResults), arg0)
}

// UpsertWhitelistResults mocks base method
func (m *MockStore) UpsertWhitelistResults(arg0 *storage.ProcessWhitelistResults) error {
	ret := m.ctrl.Call(m, "UpsertWhitelistResults", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertWhitelistResults indicates an expected call of UpsertWhitelistResults
func (mr *MockStoreMockRecorder) UpsertWhitelistResults(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertWhitelistResults", reflect.TypeOf((*MockStore)(nil).UpsertWhitelistResults), arg0)
}
