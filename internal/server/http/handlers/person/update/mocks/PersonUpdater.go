// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/insan1a/exile/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// PersonUpdater is an autogenerated mock type for the PersonUpdater type
type PersonUpdater struct {
	mock.Mock
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *PersonUpdater) Update(_a0 context.Context, _a1 *models.Person) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Person) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPersonUpdater interface {
	mock.TestingT
	Cleanup(func())
}

// NewPersonUpdater creates a new instance of PersonUpdater. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPersonUpdater(t mockConstructorTestingTNewPersonUpdater) *PersonUpdater {
	mock := &PersonUpdater{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
