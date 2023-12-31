// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Consumer is an autogenerated mock type for the Consumer type
type Consumer struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Consumer) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Consume provides a mock function with given fields: timeout
func (_m *Consumer) Consume(timeout time.Duration) ([]byte, error) {
	ret := _m.Called(timeout)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Duration) ([]byte, error)); ok {
		return rf(timeout)
	}
	if rf, ok := ret.Get(0).(func(time.Duration) []byte); ok {
		r0 = rf(timeout)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(time.Duration) error); ok {
		r1 = rf(timeout)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewConsumer interface {
	mock.TestingT
	Cleanup(func())
}

// NewConsumer creates a new instance of Consumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConsumer(t mockConstructorTestingTNewConsumer) *Consumer {
	mock := &Consumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
