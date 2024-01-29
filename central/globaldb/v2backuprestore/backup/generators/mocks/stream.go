// Code generated by MockGen. DO NOT EDIT.
// Source: stream.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/stream.go -source stream.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockStreamGenerator is a mock of StreamGenerator interface.
type MockStreamGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockStreamGeneratorMockRecorder
}

// MockStreamGeneratorMockRecorder is the mock recorder for MockStreamGenerator.
type MockStreamGeneratorMockRecorder struct {
	mock *MockStreamGenerator
}

// NewMockStreamGenerator creates a new mock instance.
func NewMockStreamGenerator(ctrl *gomock.Controller) *MockStreamGenerator {
	mock := &MockStreamGenerator{ctrl: ctrl}
	mock.recorder = &MockStreamGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamGenerator) EXPECT() *MockStreamGeneratorMockRecorder {
	return m.recorder
}

// WriteTo mocks base method.
func (m *MockStreamGenerator) WriteTo(ctx context.Context, writer io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTo", ctx, writer)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTo indicates an expected call of WriteTo.
func (mr *MockStreamGeneratorMockRecorder) WriteTo(ctx, writer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTo", reflect.TypeOf((*MockStreamGenerator)(nil).WriteTo), ctx, writer)
}
