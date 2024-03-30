// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SingleResult is an autogenerated mock type for the SingleResult type
type SingleResult struct {
	mock.Mock
}

// Decode provides a mock function with given fields: _a0
func (_m *SingleResult) Decode(_a0 interface{}) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Decode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSingleResult creates a new instance of SingleResult. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSingleResult(t interface {
	mock.TestingT
	Cleanup(func())
}) *SingleResult {
	mock := &SingleResult{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
