// Code generated by MockGen. DO NOT EDIT.
// Source: client.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/client.go -source client.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	discoveredclusters "github.com/stackrox/rox/pkg/cloudsources/discoveredclusters"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetDiscoveredClusters mocks base method.
func (m *MockClient) GetDiscoveredClusters(ctx context.Context) ([]*discoveredclusters.DiscoveredCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiscoveredClusters", ctx)
	ret0, _ := ret[0].([]*discoveredclusters.DiscoveredCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiscoveredClusters indicates an expected call of GetDiscoveredClusters.
func (mr *MockClientMockRecorder) GetDiscoveredClusters(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiscoveredClusters", reflect.TypeOf((*MockClient)(nil).GetDiscoveredClusters), ctx)
}
