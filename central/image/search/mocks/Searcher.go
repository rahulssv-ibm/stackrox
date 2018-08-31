// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

import v1 "github.com/stackrox/rox/generated/api/v1"

// Searcher is an autogenerated mock type for the Searcher type
type Searcher struct {
	mock.Mock
}

// SearchImages provides a mock function with given fields: q
func (_m *Searcher) SearchImages(q *v1.Query) ([]*v1.SearchResult, error) {
	ret := _m.Called(q)

	var r0 []*v1.SearchResult
	if rf, ok := ret.Get(0).(func(*v1.Query) []*v1.SearchResult); ok {
		r0 = rf(q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.SearchResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.Query) error); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchListImages provides a mock function with given fields: q
func (_m *Searcher) SearchListImages(q *v1.Query) ([]*v1.ListImage, error) {
	ret := _m.Called(q)

	var r0 []*v1.ListImage
	if rf, ok := ret.Get(0).(func(*v1.Query) []*v1.ListImage); ok {
		r0 = rf(q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.ListImage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.Query) error); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchRawImages provides a mock function with given fields: q
func (_m *Searcher) SearchRawImages(q *v1.Query) ([]*v1.Image, error) {
	ret := _m.Called(q)

	var r0 []*v1.Image
	if rf, ok := ret.Get(0).(func(*v1.Query) []*v1.Image); ok {
		r0 = rf(q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.Query) error); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
