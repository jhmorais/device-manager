// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	output "github.com/jhmorais/device-manager/internal/usecases/ports/output"
)

// FindDeviceByIDUseCase is an autogenerated mock type for the FindDeviceByIDUseCase type
type FindDeviceByIDUseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, deviceID
func (_m *FindDeviceByIDUseCase) Execute(ctx context.Context, deviceID string) (*output.FindDeviceOutput, error) {
	ret := _m.Called(ctx, deviceID)

	var r0 *output.FindDeviceOutput
	if rf, ok := ret.Get(0).(func(context.Context, string) *output.FindDeviceOutput); ok {
		r0 = rf(ctx, deviceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*output.FindDeviceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, deviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFindDeviceByIDUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewFindDeviceByIDUseCase creates a new instance of FindDeviceByIDUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFindDeviceByIDUseCase(t mockConstructorTestingTNewFindDeviceByIDUseCase) *FindDeviceByIDUseCase {
	mock := &FindDeviceByIDUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
