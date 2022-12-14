// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	output "github.com/jhmorais/device-manager/internal/usecases/ports/output"
)

// ListDeviceUseCase is an autogenerated mock type for the ListDeviceUseCase type
type ListDeviceUseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx
func (_m *ListDeviceUseCase) Execute(ctx context.Context) (*output.ListDeviceOutput, error) {
	ret := _m.Called(ctx)

	var r0 *output.ListDeviceOutput
	if rf, ok := ret.Get(0).(func(context.Context) *output.ListDeviceOutput); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*output.ListDeviceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewListDeviceUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewListDeviceUseCase creates a new instance of ListDeviceUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewListDeviceUseCase(t mockConstructorTestingTNewListDeviceUseCase) *ListDeviceUseCase {
	mock := &ListDeviceUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
