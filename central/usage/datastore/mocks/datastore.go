// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "github.com/gogo/protobuf/types"
	central "github.com/stackrox/rox/generated/internalapi/central"
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

// CutMetrics mocks base method.
func (m *MockDataStore) CutMetrics(ctx context.Context) (*storage.Usage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CutMetrics", ctx)
	ret0, _ := ret[0].(*storage.Usage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CutMetrics indicates an expected call of CutMetrics.
func (mr *MockDataStoreMockRecorder) CutMetrics(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CutMetrics", reflect.TypeOf((*MockDataStore)(nil).CutMetrics), ctx)
}

// Get mocks base method.
func (m *MockDataStore) Get(ctx context.Context, from, to *types.Timestamp) ([]*storage.Usage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, from, to)
	ret0, _ := ret[0].([]*storage.Usage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDataStoreMockRecorder) Get(ctx, from, to interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDataStore)(nil).Get), ctx, from, to)
}

// GetCurrent mocks base method.
func (m *MockDataStore) GetCurrent(ctx context.Context) (*storage.Usage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrent", ctx)
	ret0, _ := ret[0].(*storage.Usage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrent indicates an expected call of GetCurrent.
func (mr *MockDataStoreMockRecorder) GetCurrent(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrent", reflect.TypeOf((*MockDataStore)(nil).GetCurrent), ctx)
}

// Insert mocks base method.
func (m *MockDataStore) Insert(ctx context.Context, metrics *storage.Usage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, metrics)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockDataStoreMockRecorder) Insert(ctx, metrics interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockDataStore)(nil).Insert), ctx, metrics)
}

// UpdateUsage mocks base method.
func (m *MockDataStore) UpdateUsage(clusterID string, metrics *central.ClusterMetrics) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUsage", clusterID, metrics)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUsage indicates an expected call of UpdateUsage.
func (mr *MockDataStoreMockRecorder) UpdateUsage(clusterID, metrics interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUsage", reflect.TypeOf((*MockDataStore)(nil).UpdateUsage), clusterID, metrics)
}
