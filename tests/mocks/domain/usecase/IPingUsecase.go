// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IPingUsecase is an autogenerated mock type for the IPingUsecase type
type IPingUsecase struct {
	mock.Mock
}

// PingRepository provides a mock function with given fields:
func (_m *IPingUsecase) PingRepository() (bool, error) {
	ret := _m.Called()

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func() (bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PingUsecase provides a mock function with given fields:
func (_m *IPingUsecase) PingUsecase() (bool, error) {
	ret := _m.Called()

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func() (bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIPingUsecase creates a new instance of IPingUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIPingUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IPingUsecase {
	mock := &IPingUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
