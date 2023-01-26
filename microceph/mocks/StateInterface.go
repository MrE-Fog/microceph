// Package mocks cluster state interface. Generated by mockery with a minor update as mockery confuses import paths
package mocks

import (
	state "github.com/canonical/microcluster/state" // mockery gets confused about import paths here
	mock "github.com/stretchr/testify/mock"
)

// StateInterface is an autogenerated mock type for the StateInterface type
type StateInterface struct {
	mock.Mock
}

// ClusterState provides a mock function with given fields:
func (_m *StateInterface) ClusterState() *state.State {
	ret := _m.Called()

	var r0 *state.State
	if rf, ok := ret.Get(0).(func() *state.State); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.State)
		}
	}

	return r0
}

type mockConstructorTestingTNewStateInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewStateInterface creates a new instance of StateInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStateInterface(t mockConstructorTestingTNewStateInterface) *StateInterface {
	mock := &StateInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
