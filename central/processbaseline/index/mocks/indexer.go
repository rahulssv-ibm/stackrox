// Code generated by MockGen. DO NOT EDIT.
// Source: indexer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
	blevesearch "github.com/stackrox/rox/pkg/search/blevesearch"
)

// MockIndexer is a mock of Indexer interface.
type MockIndexer struct {
	ctrl     *gomock.Controller
	recorder *MockIndexerMockRecorder
}

// MockIndexerMockRecorder is the mock recorder for MockIndexer.
type MockIndexerMockRecorder struct {
	mock *MockIndexer
}

// NewMockIndexer creates a new mock instance.
func NewMockIndexer(ctrl *gomock.Controller) *MockIndexer {
	mock := &MockIndexer{ctrl: ctrl}
	mock.recorder = &MockIndexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndexer) EXPECT() *MockIndexerMockRecorder {
	return m.recorder
}

// AddBaseline mocks base method.
func (m *MockIndexer) AddProcessBaseline(baseline *storage.ProcessBaseline) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProcessBaseline", baseline)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBaseline indicates an expected call of AddBaseline.
func (mr *MockIndexerMockRecorder) AddBaseline(baseline interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProcessBaseline", reflect.TypeOf((*MockIndexer)(nil).AddProcessBaseline), baseline)
}

// AddBaselines mocks base method.
func (m *MockIndexer) AddProcessBaselines(baselines []*storage.ProcessBaseline) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProcessBaselines", baselines)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBaselines indicates an expected call of AddBaselines.
func (mr *MockIndexerMockRecorder) AddBaselines(baselines interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProcessBaselines", reflect.TypeOf((*MockIndexer)(nil).AddProcessBaselines), baselines)
}

// Count mocks base method.
func (m *MockIndexer) Count(q *v1.Query, opts ...blevesearch.SearchOption) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{q}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Count", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockIndexerMockRecorder) Count(q interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{q}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockIndexer)(nil).Count), varargs...)
}

// DeleteBaseline mocks base method.
func (m *MockIndexer) DeleteProcessBaseline(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProcessBaseline", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBaseline indicates an expected call of DeleteBaseline.
func (mr *MockIndexerMockRecorder) DeleteBaseline(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProcessBaseline", reflect.TypeOf((*MockIndexer)(nil).DeleteProcessBaseline), id)
}

// DeleteBaselines mocks base method.
func (m *MockIndexer) DeleteProcessBaselines(ids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProcessBaselines", ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBaselines indicates an expected call of DeleteBaselines.
func (mr *MockIndexerMockRecorder) DeleteBaselines(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProcessBaselines", reflect.TypeOf((*MockIndexer)(nil).DeleteProcessBaselines), ids)
}

// MarkInitialIndexingComplete mocks base method.
func (m *MockIndexer) MarkInitialIndexingComplete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkInitialIndexingComplete")
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkInitialIndexingComplete indicates an expected call of MarkInitialIndexingComplete.
func (mr *MockIndexerMockRecorder) MarkInitialIndexingComplete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkInitialIndexingComplete", reflect.TypeOf((*MockIndexer)(nil).MarkInitialIndexingComplete))
}

// NeedsInitialIndexing mocks base method.
func (m *MockIndexer) NeedsInitialIndexing() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeedsInitialIndexing")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NeedsInitialIndexing indicates an expected call of NeedsInitialIndexing.
func (mr *MockIndexerMockRecorder) NeedsInitialIndexing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeedsInitialIndexing", reflect.TypeOf((*MockIndexer)(nil).NeedsInitialIndexing))
}

// Search mocks base method.
func (m *MockIndexer) Search(q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{q}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Search", varargs...)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockIndexerMockRecorder) Search(q interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{q}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockIndexer)(nil).Search), varargs...)
}
