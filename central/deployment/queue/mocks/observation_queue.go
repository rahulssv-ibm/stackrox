// Code generated by MockGen. DO NOT EDIT.
// Source: observation_queue.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/observation_queue.go -source observation_queue.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	queue "github.com/stackrox/rox/central/deployment/queue"
	gomock "go.uber.org/mock/gomock"
)

// MockDeploymentObservationQueue is a mock of DeploymentObservationQueue interface.
type MockDeploymentObservationQueue struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentObservationQueueMockRecorder
}

// MockDeploymentObservationQueueMockRecorder is the mock recorder for MockDeploymentObservationQueue.
type MockDeploymentObservationQueueMockRecorder struct {
	mock *MockDeploymentObservationQueue
}

// NewMockDeploymentObservationQueue creates a new mock instance.
func NewMockDeploymentObservationQueue(ctrl *gomock.Controller) *MockDeploymentObservationQueue {
	mock := &MockDeploymentObservationQueue{ctrl: ctrl}
	mock.recorder = &MockDeploymentObservationQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentObservationQueue) EXPECT() *MockDeploymentObservationQueueMockRecorder {
	return m.recorder
}

// GetObservationDetails mocks base method.
func (m *MockDeploymentObservationQueue) GetObservationDetails(deploymentID string) *queue.DeploymentObservation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObservationDetails", deploymentID)
	ret0, _ := ret[0].(*queue.DeploymentObservation)
	return ret0
}

// GetObservationDetails indicates an expected call of GetObservationDetails.
func (mr *MockDeploymentObservationQueueMockRecorder) GetObservationDetails(deploymentID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObservationDetails", reflect.TypeOf((*MockDeploymentObservationQueue)(nil).GetObservationDetails), deploymentID)
}

// InObservation mocks base method.
func (m *MockDeploymentObservationQueue) InObservation(deploymentID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InObservation", deploymentID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// InObservation indicates an expected call of InObservation.
func (mr *MockDeploymentObservationQueueMockRecorder) InObservation(deploymentID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InObservation", reflect.TypeOf((*MockDeploymentObservationQueue)(nil).InObservation), deploymentID)
}

// Peek mocks base method.
func (m *MockDeploymentObservationQueue) Peek() *queue.DeploymentObservation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peek")
	ret0, _ := ret[0].(*queue.DeploymentObservation)
	return ret0
}

// Peek indicates an expected call of Peek.
func (mr *MockDeploymentObservationQueueMockRecorder) Peek() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peek", reflect.TypeOf((*MockDeploymentObservationQueue)(nil).Peek))
}

// Pull mocks base method.
func (m *MockDeploymentObservationQueue) Pull() *queue.DeploymentObservation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pull")
	ret0, _ := ret[0].(*queue.DeploymentObservation)
	return ret0
}

// Pull indicates an expected call of Pull.
func (mr *MockDeploymentObservationQueueMockRecorder) Pull() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockDeploymentObservationQueue)(nil).Pull))
}

// Push mocks base method.
func (m *MockDeploymentObservationQueue) Push(observation *queue.DeploymentObservation) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Push", observation)
}

// Push indicates an expected call of Push.
func (mr *MockDeploymentObservationQueueMockRecorder) Push(observation any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockDeploymentObservationQueue)(nil).Push), observation)
}

// PutBackInObservation mocks base method.
func (m *MockDeploymentObservationQueue) PutBackInObservation(observation *queue.DeploymentObservation) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutBackInObservation", observation)
}

// PutBackInObservation indicates an expected call of PutBackInObservation.
func (mr *MockDeploymentObservationQueueMockRecorder) PutBackInObservation(observation any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutBackInObservation", reflect.TypeOf((*MockDeploymentObservationQueue)(nil).PutBackInObservation), observation)
}

// RemoveDeployment mocks base method.
func (m *MockDeploymentObservationQueue) RemoveDeployment(deploymentID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveDeployment", deploymentID)
}

// RemoveDeployment indicates an expected call of RemoveDeployment.
func (mr *MockDeploymentObservationQueueMockRecorder) RemoveDeployment(deploymentID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDeployment", reflect.TypeOf((*MockDeploymentObservationQueue)(nil).RemoveDeployment), deploymentID)
}

// RemoveFromObservation mocks base method.
func (m *MockDeploymentObservationQueue) RemoveFromObservation(deploymentID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveFromObservation", deploymentID)
}

// RemoveFromObservation indicates an expected call of RemoveFromObservation.
func (mr *MockDeploymentObservationQueueMockRecorder) RemoveFromObservation(deploymentID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromObservation", reflect.TypeOf((*MockDeploymentObservationQueue)(nil).RemoveFromObservation), deploymentID)
}
