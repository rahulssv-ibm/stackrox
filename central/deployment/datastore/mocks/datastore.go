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

	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
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

// Count mocks base method.
func (m *MockDataStore) Count(ctx context.Context, q *v1.Query) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, q)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockDataStoreMockRecorder) Count(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDataStore)(nil).Count), ctx, q)
}

// CountDeployments mocks base method.
func (m *MockDataStore) CountDeployments(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountDeployments", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountDeployments indicates an expected call of CountDeployments.
func (mr *MockDataStoreMockRecorder) CountDeployments(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountDeployments", reflect.TypeOf((*MockDataStore)(nil).CountDeployments), ctx)
}

// GetDeployment mocks base method.
func (m *MockDataStore) GetDeployment(ctx context.Context, id string) (*storage.Deployment, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployment", ctx, id)
	ret0, _ := ret[0].(*storage.Deployment)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDeployment indicates an expected call of GetDeployment.
func (mr *MockDataStoreMockRecorder) GetDeployment(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockDataStore)(nil).GetDeployment), ctx, id)
}

// GetDeploymentIDs mocks base method.
func (m *MockDataStore) GetDeploymentIDs(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentIDs", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeploymentIDs indicates an expected call of GetDeploymentIDs.
func (mr *MockDataStoreMockRecorder) GetDeploymentIDs(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentIDs", reflect.TypeOf((*MockDataStore)(nil).GetDeploymentIDs), ctx)
}

// GetDeployments mocks base method.
func (m *MockDataStore) GetDeployments(ctx context.Context, ids []string) ([]*storage.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployments", ctx, ids)
	ret0, _ := ret[0].([]*storage.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployments indicates an expected call of GetDeployments.
func (mr *MockDataStoreMockRecorder) GetDeployments(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployments", reflect.TypeOf((*MockDataStore)(nil).GetDeployments), ctx, ids)
}

// GetImagesForDeployment mocks base method.
func (m *MockDataStore) GetImagesForDeployment(ctx context.Context, deployment *storage.Deployment) ([]*storage.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImagesForDeployment", ctx, deployment)
	ret0, _ := ret[0].([]*storage.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImagesForDeployment indicates an expected call of GetImagesForDeployment.
func (mr *MockDataStoreMockRecorder) GetImagesForDeployment(ctx, deployment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImagesForDeployment", reflect.TypeOf((*MockDataStore)(nil).GetImagesForDeployment), ctx, deployment)
}

// ListDeployment mocks base method.
func (m *MockDataStore) ListDeployment(ctx context.Context, id string) (*storage.ListDeployment, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeployment", ctx, id)
	ret0, _ := ret[0].(*storage.ListDeployment)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListDeployment indicates an expected call of ListDeployment.
func (mr *MockDataStoreMockRecorder) ListDeployment(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeployment", reflect.TypeOf((*MockDataStore)(nil).ListDeployment), ctx, id)
}

// RemoveDeployment mocks base method.
func (m *MockDataStore) RemoveDeployment(ctx context.Context, clusterID, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDeployment", ctx, clusterID, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDeployment indicates an expected call of RemoveDeployment.
func (mr *MockDataStoreMockRecorder) RemoveDeployment(ctx, clusterID, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDeployment", reflect.TypeOf((*MockDataStore)(nil).RemoveDeployment), ctx, clusterID, id)
}

// Search mocks base method.
func (m *MockDataStore) Search(ctx context.Context, q *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, q)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockDataStoreMockRecorder) Search(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockDataStore)(nil).Search), ctx, q)
}

// SearchDeployments mocks base method.
func (m *MockDataStore) SearchDeployments(ctx context.Context, q *v1.Query) ([]*v1.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchDeployments", ctx, q)
	ret0, _ := ret[0].([]*v1.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchDeployments indicates an expected call of SearchDeployments.
func (mr *MockDataStoreMockRecorder) SearchDeployments(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchDeployments", reflect.TypeOf((*MockDataStore)(nil).SearchDeployments), ctx, q)
}

// SearchListDeployments mocks base method.
func (m *MockDataStore) SearchListDeployments(ctx context.Context, q *v1.Query) ([]*storage.ListDeployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchListDeployments", ctx, q)
	ret0, _ := ret[0].([]*storage.ListDeployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchListDeployments indicates an expected call of SearchListDeployments.
func (mr *MockDataStoreMockRecorder) SearchListDeployments(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchListDeployments", reflect.TypeOf((*MockDataStore)(nil).SearchListDeployments), ctx, q)
}

// SearchRawDeployments mocks base method.
func (m *MockDataStore) SearchRawDeployments(ctx context.Context, q *v1.Query) ([]*storage.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRawDeployments", ctx, q)
	ret0, _ := ret[0].([]*storage.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRawDeployments indicates an expected call of SearchRawDeployments.
func (mr *MockDataStoreMockRecorder) SearchRawDeployments(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRawDeployments", reflect.TypeOf((*MockDataStore)(nil).SearchRawDeployments), ctx, q)
}

// UpsertDeployment mocks base method.
func (m *MockDataStore) UpsertDeployment(ctx context.Context, deployment *storage.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertDeployment", ctx, deployment)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertDeployment indicates an expected call of UpsertDeployment.
func (mr *MockDataStoreMockRecorder) UpsertDeployment(ctx, deployment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertDeployment", reflect.TypeOf((*MockDataStore)(nil).UpsertDeployment), ctx, deployment)
}

// WalkByQuery mocks base method.
func (m *MockDataStore) WalkByQuery(ctx context.Context, query *v1.Query, fn func(*storage.Deployment) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalkByQuery", ctx, query, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalkByQuery indicates an expected call of WalkByQuery.
func (mr *MockDataStoreMockRecorder) WalkByQuery(ctx, query, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalkByQuery", reflect.TypeOf((*MockDataStore)(nil).WalkByQuery), ctx, query, fn)
}
