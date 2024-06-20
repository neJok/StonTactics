// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	mongo "github.com/neJok/StonTactics/mongo"
	mock "github.com/stretchr/testify/mock"
)

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

// Client provides a mock function with given fields:
func (_m *Database) Client() mongo.Client {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Client")
	}

	var r0 mongo.Client
	if rf, ok := ret.Get(0).(func() mongo.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo.Client)
		}
	}

	return r0
}

// Collection provides a mock function with given fields: _a0
func (_m *Database) Collection(_a0 string) mongo.Collection {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Collection")
	}

	var r0 mongo.Collection
	if rf, ok := ret.Get(0).(func(string) mongo.Collection); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo.Collection)
		}
	}

	return r0
}

// NewDatabase creates a new instance of Database. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDatabase(t interface {
	mock.TestingT
	Cleanup(func())
}) *Database {
	mock := &Database{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
