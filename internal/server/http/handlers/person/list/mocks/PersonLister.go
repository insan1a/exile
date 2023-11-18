// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/insan1a/exile/internal/models"
)

// PersonLister is an autogenerated mock type for the PersonLister type
type PersonLister struct {
	mock.Mock
}

// List provides a mock function with given fields: ctx, filter, query
func (_m *PersonLister) List(ctx context.Context, filter *models.Filter, query string) ([]models.Person, error) {
	ret := _m.Called(ctx, filter, query)

	var r0 []models.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Filter, string) ([]models.Person, error)); ok {
		return rf(ctx, filter, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Filter, string) []models.Person); ok {
		r0 = rf(ctx, filter, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Filter, string) error); ok {
		r1 = rf(ctx, filter, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPersonLister interface {
	mock.TestingT
	Cleanup(func())
}

// NewPersonLister creates a new instance of PersonLister. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPersonLister(t mockConstructorTestingTNewPersonLister) *PersonLister {
	mock := &PersonLister{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
