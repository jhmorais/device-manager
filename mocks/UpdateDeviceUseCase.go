// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	context "context"

	input "github.com/jhmorais/device-manager/internal/usecases/ports/input"

	mock "github.com/stretchr/testify/mock"

	output "github.com/jhmorais/device-manager/internal/usecases/ports/output"
)

// UpdateDeviceUseCase is an autogenerated mock type for the UpdateDeviceUseCase type
type UpdateDeviceUseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, updateDevice
func (_m *UpdateDeviceUseCase) Execute(ctx context.Context, updateDevice *input.UpdateDeviceInput) (*output.CreateDeviceOutput, error) {
	ret := _m.Called(ctx, updateDevice)

	var r0 *output.CreateDeviceOutput
	if rf, ok := ret.Get(0).(func(context.Context, *input.UpdateDeviceInput) *output.CreateDeviceOutput); ok {
		r0 = rf(ctx, updateDevice)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*output.CreateDeviceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *input.UpdateDeviceInput) error); ok {
		r1 = rf(ctx, updateDevice)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUpdateDeviceUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUpdateDeviceUseCase creates a new instance of UpdateDeviceUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUpdateDeviceUseCase(t mockConstructorTestingTNewUpdateDeviceUseCase) *UpdateDeviceUseCase {
	mock := &UpdateDeviceUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}