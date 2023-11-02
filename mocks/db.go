// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Db is an autogenerated mock type for the Db type
type Db struct {
	mock.Mock
}

// Insert provides a mock function with given fields: _a0
func (_m *Db) Insert(_a0 interface{}) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(interface{}) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewDb creates a new instance of Db. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDb(t interface {
	mock.TestingT
	Cleanup(func())
}) *Db {
	mock := &Db{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}