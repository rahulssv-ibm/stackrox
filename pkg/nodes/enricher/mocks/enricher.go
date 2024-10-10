// Code generated by MockGen. DO NOT EDIT.
// Source: enricher.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/enricher.go -source enricher.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	v4 "github.com/stackrox/rox/generated/internalapi/scanner/v4"
	storage "github.com/stackrox/rox/generated/storage"
	types "github.com/stackrox/rox/pkg/scanners/types"
	gomock "go.uber.org/mock/gomock"
)

// MockNodeEnricher is a mock of NodeEnricher interface.
type MockNodeEnricher struct {
	ctrl     *gomock.Controller
	recorder *MockNodeEnricherMockRecorder
}

// MockNodeEnricherMockRecorder is the mock recorder for MockNodeEnricher.
type MockNodeEnricherMockRecorder struct {
	mock *MockNodeEnricher
}

// NewMockNodeEnricher creates a new mock instance.
func NewMockNodeEnricher(ctrl *gomock.Controller) *MockNodeEnricher {
	mock := &MockNodeEnricher{ctrl: ctrl}
	mock.recorder = &MockNodeEnricherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeEnricher) EXPECT() *MockNodeEnricherMockRecorder {
	return m.recorder
}

// CreateNodeScanner mocks base method.
func (m *MockNodeEnricher) CreateNodeScanner(integration *storage.NodeIntegration) (types.NodeScannerWithDataSource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNodeScanner", integration)
	ret0, _ := ret[0].(types.NodeScannerWithDataSource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNodeScanner indicates an expected call of CreateNodeScanner.
func (mr *MockNodeEnricherMockRecorder) CreateNodeScanner(integration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNodeScanner", reflect.TypeOf((*MockNodeEnricher)(nil).CreateNodeScanner), integration)
}

// EnrichNode mocks base method.
func (m *MockNodeEnricher) EnrichNode(node *storage.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnrichNode", node)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnrichNode indicates an expected call of EnrichNode.
func (mr *MockNodeEnricherMockRecorder) EnrichNode(node any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnrichNode", reflect.TypeOf((*MockNodeEnricher)(nil).EnrichNode), node)
}

// EnrichNodeWithInventory mocks base method.
func (m *MockNodeEnricher) EnrichNodeWithInventory(node *storage.Node, nodeInventory *storage.NodeInventory, indexReport *v4.IndexReport) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnrichNodeWithInventory", node, nodeInventory, indexReport)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnrichNodeWithInventory indicates an expected call of EnrichNodeWithInventory.
func (mr *MockNodeEnricherMockRecorder) EnrichNodeWithInventory(node, nodeInventory, indexReport any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnrichNodeWithInventory", reflect.TypeOf((*MockNodeEnricher)(nil).EnrichNodeWithInventory), node, nodeInventory, indexReport)
}

// RemoveNodeIntegration mocks base method.
func (m *MockNodeEnricher) RemoveNodeIntegration(id string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveNodeIntegration", id)
}

// RemoveNodeIntegration indicates an expected call of RemoveNodeIntegration.
func (mr *MockNodeEnricherMockRecorder) RemoveNodeIntegration(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNodeIntegration", reflect.TypeOf((*MockNodeEnricher)(nil).RemoveNodeIntegration), id)
}

// UpsertNodeIntegration mocks base method.
func (m *MockNodeEnricher) UpsertNodeIntegration(integration *storage.NodeIntegration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertNodeIntegration", integration)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertNodeIntegration indicates an expected call of UpsertNodeIntegration.
func (mr *MockNodeEnricherMockRecorder) UpsertNodeIntegration(integration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertNodeIntegration", reflect.TypeOf((*MockNodeEnricher)(nil).UpsertNodeIntegration), integration)
}

// MockCVESuppressor is a mock of CVESuppressor interface.
type MockCVESuppressor struct {
	ctrl     *gomock.Controller
	recorder *MockCVESuppressorMockRecorder
}

// MockCVESuppressorMockRecorder is the mock recorder for MockCVESuppressor.
type MockCVESuppressorMockRecorder struct {
	mock *MockCVESuppressor
}

// NewMockCVESuppressor creates a new mock instance.
func NewMockCVESuppressor(ctrl *gomock.Controller) *MockCVESuppressor {
	mock := &MockCVESuppressor{ctrl: ctrl}
	mock.recorder = &MockCVESuppressorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCVESuppressor) EXPECT() *MockCVESuppressorMockRecorder {
	return m.recorder
}

// EnrichNodeWithSuppressedCVEs mocks base method.
func (m *MockCVESuppressor) EnrichNodeWithSuppressedCVEs(image *storage.Node) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnrichNodeWithSuppressedCVEs", image)
}

// EnrichNodeWithSuppressedCVEs indicates an expected call of EnrichNodeWithSuppressedCVEs.
func (mr *MockCVESuppressorMockRecorder) EnrichNodeWithSuppressedCVEs(image any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnrichNodeWithSuppressedCVEs", reflect.TypeOf((*MockCVESuppressor)(nil).EnrichNodeWithSuppressedCVEs), image)
}
