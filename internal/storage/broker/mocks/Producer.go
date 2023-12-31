// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Producer is an autogenerated mock type for the Producer type
type Producer struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Producer) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Produce provides a mock function with given fields: message
func (_m *Producer) Produce(message []byte) error {
	ret := _m.Called(message)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewProducer interface {
	mock.TestingT
	Cleanup(func())
}

// NewProducer creates a new instance of Producer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProducer(t mockConstructorTestingTNewProducer) *Producer {
	mock := &Producer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
