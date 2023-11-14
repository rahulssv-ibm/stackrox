// Code generated by MockGen. DO NOT EDIT.
// Source: flow.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/flow.go -source flow.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	types "google.golang.org/protobuf/types/known/timestamppb"
	storage "github.com/stackrox/rox/generated/storage"
	timestamp "github.com/stackrox/rox/pkg/timestamp"
	gomock "go.uber.org/mock/gomock"
)

// MockFlowDataStore is a mock of FlowDataStore interface.
type MockFlowDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockFlowDataStoreMockRecorder
}

// MockFlowDataStoreMockRecorder is the mock recorder for MockFlowDataStore.
type MockFlowDataStoreMockRecorder struct {
	mock *MockFlowDataStore
}

// NewMockFlowDataStore creates a new mock instance.
func NewMockFlowDataStore(ctrl *gomock.Controller) *MockFlowDataStore {
	mock := &MockFlowDataStore{ctrl: ctrl}
	mock.recorder = &MockFlowDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlowDataStore) EXPECT() *MockFlowDataStoreMockRecorder {
	return m.recorder
}

// GetAllFlows mocks base method.
func (m *MockFlowDataStore) GetAllFlows(ctx context.Context, since *types.Timestamp) ([]*storage.NetworkFlow, *types.Timestamp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFlows", ctx, since)
	ret0, _ := ret[0].([]*storage.NetworkFlow)
	ret1, _ := ret[1].(*types.Timestamp)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllFlows indicates an expected call of GetAllFlows.
func (mr *MockFlowDataStoreMockRecorder) GetAllFlows(ctx, since any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFlows", reflect.TypeOf((*MockFlowDataStore)(nil).GetAllFlows), ctx, since)
}

// GetFlowsForDeployment mocks base method.
func (m *MockFlowDataStore) GetFlowsForDeployment(ctx context.Context, deploymentID string, adjustForGraph bool) ([]*storage.NetworkFlow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlowsForDeployment", ctx, deploymentID, adjustForGraph)
	ret0, _ := ret[0].([]*storage.NetworkFlow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlowsForDeployment indicates an expected call of GetFlowsForDeployment.
func (mr *MockFlowDataStoreMockRecorder) GetFlowsForDeployment(ctx, deploymentID, adjustForGraph any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlowsForDeployment", reflect.TypeOf((*MockFlowDataStore)(nil).GetFlowsForDeployment), ctx, deploymentID, adjustForGraph)
}

// GetMatchingFlows mocks base method.
func (m *MockFlowDataStore) GetMatchingFlows(ctx context.Context, pred func(*storage.NetworkFlowProperties) bool, since *types.Timestamp) ([]*storage.NetworkFlow, *types.Timestamp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchingFlows", ctx, pred, since)
	ret0, _ := ret[0].([]*storage.NetworkFlow)
	ret1, _ := ret[1].(*types.Timestamp)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMatchingFlows indicates an expected call of GetMatchingFlows.
func (mr *MockFlowDataStoreMockRecorder) GetMatchingFlows(ctx, pred, since any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchingFlows", reflect.TypeOf((*MockFlowDataStore)(nil).GetMatchingFlows), ctx, pred, since)
}

// RemoveFlowsForDeployment mocks base method.
func (m *MockFlowDataStore) RemoveFlowsForDeployment(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFlowsForDeployment", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFlowsForDeployment indicates an expected call of RemoveFlowsForDeployment.
func (mr *MockFlowDataStoreMockRecorder) RemoveFlowsForDeployment(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFlowsForDeployment", reflect.TypeOf((*MockFlowDataStore)(nil).RemoveFlowsForDeployment), ctx, id)
}

// RemoveOrphanedFlows mocks base method.
func (m *MockFlowDataStore) RemoveOrphanedFlows(ctx context.Context, orphanWindow *time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveOrphanedFlows", ctx, orphanWindow)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveOrphanedFlows indicates an expected call of RemoveOrphanedFlows.
func (mr *MockFlowDataStoreMockRecorder) RemoveOrphanedFlows(ctx, orphanWindow any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveOrphanedFlows", reflect.TypeOf((*MockFlowDataStore)(nil).RemoveOrphanedFlows), ctx, orphanWindow)
}

// RemoveStaleFlows mocks base method.
func (m *MockFlowDataStore) RemoveStaleFlows(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveStaleFlows", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveStaleFlows indicates an expected call of RemoveStaleFlows.
func (mr *MockFlowDataStoreMockRecorder) RemoveStaleFlows(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveStaleFlows", reflect.TypeOf((*MockFlowDataStore)(nil).RemoveStaleFlows), ctx)
}

// UpsertFlows mocks base method.
func (m *MockFlowDataStore) UpsertFlows(ctx context.Context, flows []*storage.NetworkFlow, lastUpdateTS timestamp.MicroTS) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertFlows", ctx, flows, lastUpdateTS)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertFlows indicates an expected call of UpsertFlows.
func (mr *MockFlowDataStoreMockRecorder) UpsertFlows(ctx, flows, lastUpdateTS any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertFlows", reflect.TypeOf((*MockFlowDataStore)(nil).UpsertFlows), ctx, flows, lastUpdateTS)
}
