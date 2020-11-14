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

// MockEntityDataStore is a mock of EntityDataStore interface
type MockEntityDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockEntityDataStoreMockRecorder
}

// MockEntityDataStoreMockRecorder is the mock recorder for MockEntityDataStore
type MockEntityDataStoreMockRecorder struct {
	mock *MockEntityDataStore
}

// NewMockEntityDataStore creates a new mock instance
func NewMockEntityDataStore(ctrl *gomock.Controller) *MockEntityDataStore {
	mock := &MockEntityDataStore{ctrl: ctrl}
	mock.recorder = &MockEntityDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEntityDataStore) EXPECT() *MockEntityDataStoreMockRecorder {
	return m.recorder
}

// Exists mocks base method
func (m *MockEntityDataStore) Exists(ctx context.Context, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (mr *MockEntityDataStoreMockRecorder) Exists(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockEntityDataStore)(nil).Exists), ctx, id)
}

// GetIDs mocks base method
func (m *MockEntityDataStore) GetIDs(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIDs", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIDs indicates an expected call of GetIDs
func (mr *MockEntityDataStoreMockRecorder) GetIDs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIDs", reflect.TypeOf((*MockEntityDataStore)(nil).GetIDs), ctx)
}

// GetEntity mocks base method
func (m *MockEntityDataStore) GetEntity(ctx context.Context, id string) (*storage.NetworkEntity, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntity", ctx, id)
	ret0, _ := ret[0].(*storage.NetworkEntity)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetEntity indicates an expected call of GetEntity
func (mr *MockEntityDataStoreMockRecorder) GetEntity(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntity", reflect.TypeOf((*MockEntityDataStore)(nil).GetEntity), ctx, id)
}

// GetAllEntitiesForCluster mocks base method
func (m *MockEntityDataStore) GetAllEntitiesForCluster(ctx context.Context, clusterID string) ([]*storage.NetworkEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEntitiesForCluster", ctx, clusterID)
	ret0, _ := ret[0].([]*storage.NetworkEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEntitiesForCluster indicates an expected call of GetAllEntitiesForCluster
func (mr *MockEntityDataStoreMockRecorder) GetAllEntitiesForCluster(ctx, clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEntitiesForCluster", reflect.TypeOf((*MockEntityDataStore)(nil).GetAllEntitiesForCluster), ctx, clusterID)
}

// GetAllEntities mocks base method
func (m *MockEntityDataStore) GetAllEntities(ctx context.Context) ([]*storage.NetworkEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEntities", ctx)
	ret0, _ := ret[0].([]*storage.NetworkEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEntities indicates an expected call of GetAllEntities
func (mr *MockEntityDataStoreMockRecorder) GetAllEntities(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEntities", reflect.TypeOf((*MockEntityDataStore)(nil).GetAllEntities), ctx)
}

// GetAllMatchingEntities mocks base method
func (m *MockEntityDataStore) GetAllMatchingEntities(ctx context.Context, pred func(*storage.NetworkEntity) bool) ([]*storage.NetworkEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMatchingEntities", ctx, pred)
	ret0, _ := ret[0].([]*storage.NetworkEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMatchingEntities indicates an expected call of GetAllMatchingEntities
func (mr *MockEntityDataStoreMockRecorder) GetAllMatchingEntities(ctx, pred interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMatchingEntities", reflect.TypeOf((*MockEntityDataStore)(nil).GetAllMatchingEntities), ctx, pred)
}

// CreateExternalNetworkEntity mocks base method
func (m *MockEntityDataStore) CreateExternalNetworkEntity(ctx context.Context, entity *storage.NetworkEntity, skipPush bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExternalNetworkEntity", ctx, entity, skipPush)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateExternalNetworkEntity indicates an expected call of CreateExternalNetworkEntity
func (mr *MockEntityDataStoreMockRecorder) CreateExternalNetworkEntity(ctx, entity, skipPush interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExternalNetworkEntity", reflect.TypeOf((*MockEntityDataStore)(nil).CreateExternalNetworkEntity), ctx, entity, skipPush)
}

// UpdateExternalNetworkEntity mocks base method
func (m *MockEntityDataStore) UpdateExternalNetworkEntity(ctx context.Context, entity *storage.NetworkEntity, skipPush bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExternalNetworkEntity", ctx, entity, skipPush)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExternalNetworkEntity indicates an expected call of UpdateExternalNetworkEntity
func (mr *MockEntityDataStoreMockRecorder) UpdateExternalNetworkEntity(ctx, entity, skipPush interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExternalNetworkEntity", reflect.TypeOf((*MockEntityDataStore)(nil).UpdateExternalNetworkEntity), ctx, entity, skipPush)
}

// DeleteExternalNetworkEntity mocks base method
func (m *MockEntityDataStore) DeleteExternalNetworkEntity(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExternalNetworkEntity", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExternalNetworkEntity indicates an expected call of DeleteExternalNetworkEntity
func (mr *MockEntityDataStoreMockRecorder) DeleteExternalNetworkEntity(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExternalNetworkEntity", reflect.TypeOf((*MockEntityDataStore)(nil).DeleteExternalNetworkEntity), ctx, id)
}

// DeleteExternalNetworkEntitiesForCluster mocks base method
func (m *MockEntityDataStore) DeleteExternalNetworkEntitiesForCluster(ctx context.Context, clusterID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExternalNetworkEntitiesForCluster", ctx, clusterID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExternalNetworkEntitiesForCluster indicates an expected call of DeleteExternalNetworkEntitiesForCluster
func (mr *MockEntityDataStoreMockRecorder) DeleteExternalNetworkEntitiesForCluster(ctx, clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExternalNetworkEntitiesForCluster", reflect.TypeOf((*MockEntityDataStore)(nil).DeleteExternalNetworkEntitiesForCluster), ctx, clusterID)
}

// RegisterCluster mocks base method
func (m *MockEntityDataStore) RegisterCluster(clusterID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterCluster", clusterID)
}

// RegisterCluster indicates an expected call of RegisterCluster
func (mr *MockEntityDataStoreMockRecorder) RegisterCluster(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCluster", reflect.TypeOf((*MockEntityDataStore)(nil).RegisterCluster), clusterID)
}
