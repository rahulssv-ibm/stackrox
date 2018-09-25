// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import v1 "github.com/stackrox/rox/generated/api/v1"

// AlertManager is an autogenerated mock type for the AlertManager type
type AlertManager struct {
	mock.Mock
}

// AlertAndNotify provides a mock function with given fields: previousAlerts, currentAlerts
func (_m *AlertManager) AlertAndNotify(previousAlerts []*v1.Alert, currentAlerts []*v1.Alert) error {
	ret := _m.Called(previousAlerts, currentAlerts)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*v1.Alert, []*v1.Alert) error); ok {
		r0 = rf(previousAlerts, currentAlerts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAlertsByDeployment provides a mock function with given fields: deploymentID
func (_m *AlertManager) GetAlertsByDeployment(deploymentID string) ([]*v1.Alert, error) {
	ret := _m.Called(deploymentID)

	var r0 []*v1.Alert
	if rf, ok := ret.Get(0).(func(string) []*v1.Alert); ok {
		r0 = rf(deploymentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Alert)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(deploymentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAlertsByDeploymentAndPolicy provides a mock function with given fields: deploymentID, policyID
func (_m *AlertManager) GetAlertsByDeploymentAndPolicy(deploymentID string, policyID string) (*v1.Alert, error) {
	ret := _m.Called(deploymentID, policyID)

	var r0 *v1.Alert
	if rf, ok := ret.Get(0).(func(string, string) *v1.Alert); ok {
		r0 = rf(deploymentID, policyID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Alert)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(deploymentID, policyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAlertsByPolicy provides a mock function with given fields: policyID
func (_m *AlertManager) GetAlertsByPolicy(policyID string) ([]*v1.Alert, error) {
	ret := _m.Called(policyID)

	var r0 []*v1.Alert
	if rf, ok := ret.Get(0).(func(string) []*v1.Alert); ok {
		r0 = rf(policyID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Alert)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(policyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
