// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

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

// Exists mocks base method
func (m *MockDataStore) Exists(ctx context.Context, name string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, name)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (mr *MockDataStoreMockRecorder) Exists(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockDataStore)(nil).Exists), ctx, name)
}

// UpsertWatchedImage mocks base method
func (m *MockDataStore) UpsertWatchedImage(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertWatchedImage", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertWatchedImage indicates an expected call of UpsertWatchedImage
func (mr *MockDataStoreMockRecorder) UpsertWatchedImage(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertWatchedImage", reflect.TypeOf((*MockDataStore)(nil).UpsertWatchedImage), ctx, name)
}

// GetAllWatchedImages mocks base method
func (m *MockDataStore) GetAllWatchedImages(ctx context.Context) ([]*storage.WatchedImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllWatchedImages", ctx)
	ret0, _ := ret[0].([]*storage.WatchedImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllWatchedImages indicates an expected call of GetAllWatchedImages
func (mr *MockDataStoreMockRecorder) GetAllWatchedImages(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllWatchedImages", reflect.TypeOf((*MockDataStore)(nil).GetAllWatchedImages), ctx)
}

// UnwatchImage mocks base method
func (m *MockDataStore) UnwatchImage(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnwatchImage", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnwatchImage indicates an expected call of UnwatchImage
func (mr *MockDataStoreMockRecorder) UnwatchImage(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnwatchImage", reflect.TypeOf((*MockDataStore)(nil).UnwatchImage), ctx, name)
}
