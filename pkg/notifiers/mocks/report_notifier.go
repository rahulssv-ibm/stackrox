// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/pkg/notifiers (interfaces: ReportNotifier)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/report_notifier.go github.com/stackrox/rox/pkg/notifiers ReportNotifier
//

// Package mocks is a generated GoMock package.
package mocks

import (
	bytes "bytes"
	context "context"
	reflect "reflect"

	storage "github.com/stackrox/rox/generated/storage"
	notifiers "github.com/stackrox/rox/pkg/notifiers"
	gomock "go.uber.org/mock/gomock"
)

// MockReportNotifier is a mock of ReportNotifier interface.
type MockReportNotifier struct {
	ctrl     *gomock.Controller
	recorder *MockReportNotifierMockRecorder
}

// MockReportNotifierMockRecorder is the mock recorder for MockReportNotifier.
type MockReportNotifierMockRecorder struct {
	mock *MockReportNotifier
}

// NewMockReportNotifier creates a new mock instance.
func NewMockReportNotifier(ctrl *gomock.Controller) *MockReportNotifier {
	mock := &MockReportNotifier{ctrl: ctrl}
	mock.recorder = &MockReportNotifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReportNotifier) EXPECT() *MockReportNotifierMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockReportNotifier) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockReportNotifierMockRecorder) Close(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReportNotifier)(nil).Close), arg0)
}

// ProtoNotifier mocks base method.
func (m *MockReportNotifier) ProtoNotifier() *storage.Notifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProtoNotifier")
	ret0, _ := ret[0].(*storage.Notifier)
	return ret0
}

// ProtoNotifier indicates an expected call of ProtoNotifier.
func (mr *MockReportNotifierMockRecorder) ProtoNotifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtoNotifier", reflect.TypeOf((*MockReportNotifier)(nil).ProtoNotifier))
}

// ReportNotify mocks base method.
func (m *MockReportNotifier) ReportNotify(arg0 context.Context, arg1 *bytes.Buffer, arg2 []string, arg3, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportNotify", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReportNotify indicates an expected call of ReportNotify.
func (mr *MockReportNotifierMockRecorder) ReportNotify(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportNotify", reflect.TypeOf((*MockReportNotifier)(nil).ReportNotify), arg0, arg1, arg2, arg3, arg4)
}

// Test mocks base method.
func (m *MockReportNotifier) Test(arg0 context.Context) *notifiers.NotifierError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Test", arg0)
	ret0, _ := ret[0].(*notifiers.NotifierError)
	return ret0
}

// Test indicates an expected call of Test.
func (mr *MockReportNotifierMockRecorder) Test(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Test", reflect.TypeOf((*MockReportNotifier)(nil).Test), arg0)
}
