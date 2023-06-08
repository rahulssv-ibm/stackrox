// Code generated by MockGen. DO NOT EDIT.
// Source: indexer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	search "github.com/stackrox/rox/pkg/search"
	blevesearch "github.com/stackrox/rox/pkg/search/blevesearch"
)

// MockStandardIndexer is a mock of StandardIndexer interface.
type MockStandardIndexer struct {
	ctrl     *gomock.Controller
	recorder *MockStandardIndexerMockRecorder
}

// MockStandardIndexerMockRecorder is the mock recorder for MockStandardIndexer.
type MockStandardIndexerMockRecorder struct {
	mock *MockStandardIndexer
}

// NewMockStandardIndexer creates a new mock instance.
func NewMockStandardIndexer(ctrl *gomock.Controller) *MockStandardIndexer {
	mock := &MockStandardIndexer{ctrl: ctrl}
	mock.recorder = &MockStandardIndexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStandardIndexer) EXPECT() *MockStandardIndexerMockRecorder {
	return m.recorder
}

// Search mocks base method.
func (m *MockStandardIndexer) Search(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, q}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Search", varargs...)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockStandardIndexerMockRecorder) Search(ctx, q interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, q}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockStandardIndexer)(nil).Search), varargs...)
}

// MockControlIndexer is a mock of ControlIndexer interface.
type MockControlIndexer struct {
	ctrl     *gomock.Controller
	recorder *MockControlIndexerMockRecorder
}

// MockControlIndexerMockRecorder is the mock recorder for MockControlIndexer.
type MockControlIndexerMockRecorder struct {
	mock *MockControlIndexer
}

// NewMockControlIndexer creates a new mock instance.
func NewMockControlIndexer(ctrl *gomock.Controller) *MockControlIndexer {
	mock := &MockControlIndexer{ctrl: ctrl}
	mock.recorder = &MockControlIndexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControlIndexer) EXPECT() *MockControlIndexerMockRecorder {
	return m.recorder
}

// Search mocks base method.
func (m *MockControlIndexer) Search(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, q}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Search", varargs...)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockControlIndexerMockRecorder) Search(ctx, q interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, q}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockControlIndexer)(nil).Search), varargs...)
}
