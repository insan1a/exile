// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/insan1a/exile/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// PersonSaver is an autogenerated mock type for the PersonSaver type
type PersonSaver struct {
	mock.Mock
}

// Save provides a mock function with given fields: ctx, p
func (_m *PersonSaver) Save(ctx context.Context, p models.Person) error {
	ret := _m.Called(ctx, p)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Person) error); ok {
		r0 = rf(ctx, p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPersonSaver interface {
	mock.TestingT
	Cleanup(func())
}

// NewPersonSaver creates a new instance of PersonSaver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPersonSaver(t mockConstructorTestingTNewPersonSaver) *PersonSaver {
	mock := &PersonSaver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
