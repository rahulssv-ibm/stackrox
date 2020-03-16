// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	concurrency "github.com/stackrox/rox/pkg/concurrency"
	search "github.com/stackrox/rox/pkg/search"
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

// Search mocks base method
func (m *MockDataStore) Search(ctx context.Context, q *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, q)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockDataStoreMockRecorder) Search(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockDataStore)(nil).Search), ctx, q)
}

// SearchRawProcessIndicators mocks base method
func (m *MockDataStore) SearchRawProcessIndicators(ctx context.Context, q *v1.Query) ([]*storage.ProcessIndicator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRawProcessIndicators", ctx, q)
	ret0, _ := ret[0].([]*storage.ProcessIndicator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRawProcessIndicators indicates an expected call of SearchRawProcessIndicators
func (mr *MockDataStoreMockRecorder) SearchRawProcessIndicators(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRawProcessIndicators", reflect.TypeOf((*MockDataStore)(nil).SearchRawProcessIndicators), ctx, q)
}

// GetProcessIndicator mocks base method
func (m *MockDataStore) GetProcessIndicator(ctx context.Context, id string) (*storage.ProcessIndicator, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProcessIndicator", ctx, id)
	ret0, _ := ret[0].(*storage.ProcessIndicator)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProcessIndicator indicates an expected call of GetProcessIndicator
func (mr *MockDataStoreMockRecorder) GetProcessIndicator(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProcessIndicator", reflect.TypeOf((*MockDataStore)(nil).GetProcessIndicator), ctx, id)
}

// AddProcessIndicators mocks base method
func (m *MockDataStore) AddProcessIndicators(arg0 context.Context, arg1 ...*storage.ProcessIndicator) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddProcessIndicators", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProcessIndicators indicates an expected call of AddProcessIndicators
func (mr *MockDataStoreMockRecorder) AddProcessIndicators(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProcessIndicators", reflect.TypeOf((*MockDataStore)(nil).AddProcessIndicators), varargs...)
}

// RemoveProcessIndicatorsByDeployment mocks base method
func (m *MockDataStore) RemoveProcessIndicatorsByDeployment(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProcessIndicatorsByDeployment", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProcessIndicatorsByDeployment indicates an expected call of RemoveProcessIndicatorsByDeployment
func (mr *MockDataStoreMockRecorder) RemoveProcessIndicatorsByDeployment(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProcessIndicatorsByDeployment", reflect.TypeOf((*MockDataStore)(nil).RemoveProcessIndicatorsByDeployment), ctx, id)
}

// RemoveProcessIndicatorsByPod mocks base method
func (m *MockDataStore) RemoveProcessIndicatorsByPod(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProcessIndicatorsByPod", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProcessIndicatorsByPod indicates an expected call of RemoveProcessIndicatorsByPod
func (mr *MockDataStoreMockRecorder) RemoveProcessIndicatorsByPod(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProcessIndicatorsByPod", reflect.TypeOf((*MockDataStore)(nil).RemoveProcessIndicatorsByPod), ctx, id)
}

// RemoveProcessIndicatorsOfStaleContainers mocks base method
func (m *MockDataStore) RemoveProcessIndicatorsOfStaleContainers(ctx context.Context, deployment *storage.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProcessIndicatorsOfStaleContainers", ctx, deployment)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProcessIndicatorsOfStaleContainers indicates an expected call of RemoveProcessIndicatorsOfStaleContainers
func (mr *MockDataStoreMockRecorder) RemoveProcessIndicatorsOfStaleContainers(ctx, deployment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProcessIndicatorsOfStaleContainers", reflect.TypeOf((*MockDataStore)(nil).RemoveProcessIndicatorsOfStaleContainers), ctx, deployment)
}

// RemoveProcessIndicatorsOfStaleContainersByPod mocks base method
func (m *MockDataStore) RemoveProcessIndicatorsOfStaleContainersByPod(ctx context.Context, pod *storage.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProcessIndicatorsOfStaleContainersByPod", ctx, pod)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProcessIndicatorsOfStaleContainersByPod indicates an expected call of RemoveProcessIndicatorsOfStaleContainersByPod
func (mr *MockDataStoreMockRecorder) RemoveProcessIndicatorsOfStaleContainersByPod(ctx, pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProcessIndicatorsOfStaleContainersByPod", reflect.TypeOf((*MockDataStore)(nil).RemoveProcessIndicatorsOfStaleContainersByPod), ctx, pod)
}

// RemoveProcessIndicators mocks base method
func (m *MockDataStore) RemoveProcessIndicators(ctx context.Context, ids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProcessIndicators", ctx, ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProcessIndicators indicates an expected call of RemoveProcessIndicators
func (mr *MockDataStoreMockRecorder) RemoveProcessIndicators(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProcessIndicators", reflect.TypeOf((*MockDataStore)(nil).RemoveProcessIndicators), ctx, ids)
}

// AddProcessComment mocks base method
func (m *MockDataStore) AddProcessComment(ctx context.Context, processID string, comment *storage.Comment) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProcessComment", ctx, processID, comment)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProcessComment indicates an expected call of AddProcessComment
func (mr *MockDataStoreMockRecorder) AddProcessComment(ctx, processID, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProcessComment", reflect.TypeOf((*MockDataStore)(nil).AddProcessComment), ctx, processID, comment)
}

// UpdateProcessComment mocks base method
func (m *MockDataStore) UpdateProcessComment(ctx context.Context, processID string, comment *storage.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProcessComment", ctx, processID, comment)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProcessComment indicates an expected call of UpdateProcessComment
func (mr *MockDataStoreMockRecorder) UpdateProcessComment(ctx, processID, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProcessComment", reflect.TypeOf((*MockDataStore)(nil).UpdateProcessComment), ctx, processID, comment)
}

// GetCommentsForProcess mocks base method
func (m *MockDataStore) GetCommentsForProcess(ctx context.Context, processID string) ([]*storage.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentsForProcess", ctx, processID)
	ret0, _ := ret[0].([]*storage.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentsForProcess indicates an expected call of GetCommentsForProcess
func (mr *MockDataStoreMockRecorder) GetCommentsForProcess(ctx, processID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentsForProcess", reflect.TypeOf((*MockDataStore)(nil).GetCommentsForProcess), ctx, processID)
}

// RemoveProcessComment mocks base method
func (m *MockDataStore) RemoveProcessComment(ctx context.Context, processID, commentID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProcessComment", ctx, processID, commentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProcessComment indicates an expected call of RemoveProcessComment
func (mr *MockDataStoreMockRecorder) RemoveProcessComment(ctx, processID, commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProcessComment", reflect.TypeOf((*MockDataStore)(nil).RemoveProcessComment), ctx, processID, commentID)
}

// WalkAll mocks base method
func (m *MockDataStore) WalkAll(ctx context.Context, fn func(*storage.ProcessIndicator) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalkAll", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalkAll indicates an expected call of WalkAll
func (mr *MockDataStoreMockRecorder) WalkAll(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalkAll", reflect.TypeOf((*MockDataStore)(nil).WalkAll), ctx, fn)
}

// Stop mocks base method
func (m *MockDataStore) Stop() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockDataStoreMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockDataStore)(nil).Stop))
}

// Wait mocks base method
func (m *MockDataStore) Wait(cancelWhen concurrency.Waitable) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", cancelWhen)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Wait indicates an expected call of Wait
func (mr *MockDataStoreMockRecorder) Wait(cancelWhen interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockDataStore)(nil).Wait), cancelWhen)
}
