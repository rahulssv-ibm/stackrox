// Code generated by MockGen. DO NOT EDIT.
// Source: matcher.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
	booleanpolicy "github.com/stackrox/rox/pkg/booleanpolicy"
	reflect "reflect"
)

// MockMatcher is a mock of Matcher interface
type MockMatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMatcherMockRecorder
}

// MockMatcherMockRecorder is the mock recorder for MockMatcher
type MockMatcherMockRecorder struct {
	mock *MockMatcher
}

// NewMockMatcher creates a new mock instance
func NewMockMatcher(ctrl *gomock.Controller) *MockMatcher {
	mock := &MockMatcher{ctrl: ctrl}
	mock.recorder = &MockMatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMatcher) EXPECT() *MockMatcherMockRecorder {
	return m.recorder
}

// MatchOne mocks base method
func (m *MockMatcher) MatchOne(ctx context.Context, deployment *storage.Deployment, images []*storage.Image, pi *storage.ProcessIndicator) (booleanpolicy.Violations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MatchOne", ctx, deployment, images, pi)
	ret0, _ := ret[0].(booleanpolicy.Violations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MatchOne indicates an expected call of MatchOne
func (mr *MockMatcherMockRecorder) MatchOne(ctx, deployment, images, pi interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchOne", reflect.TypeOf((*MockMatcher)(nil).MatchOne), ctx, deployment, images, pi)
}
