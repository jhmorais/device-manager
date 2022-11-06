package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jhmorais/device-manager/internal/usecases"
	"github.com/jhmorais/device-manager/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestFindDevice(t *testing.T) {

	t.Run("when brand is empty should return an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
		findDeviceUseCase := usecases.NewFindDeviceUseCase(deviceRepositoryMock)

		output, err := findDeviceUseCase.Execute(context.Background(), "", "name device")
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when name is empty should return an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
		findDeviceUseCase := usecases.NewFindDeviceUseCase(deviceRepositoryMock)

		output, err := findDeviceUseCase.Execute(context.Background(), "brand device", "")
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when name and brand are empty should return an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
		findDeviceUseCase := usecases.NewFindDeviceUseCase(deviceRepositoryMock)

		output, err := findDeviceUseCase.Execute(context.Background(), "", "")
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when not found some device returns an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
		findDeviceUseCase := usecases.NewFindDeviceUseCase(deviceRepositoryMock)

		output, err := findDeviceUseCase.Execute(context.Background(), "brand device", "name device")
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when db store returns an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("database failed"))
		findDeviceUseCase := usecases.NewFindDeviceUseCase(deviceRepositoryMock)

		output, err := findDeviceUseCase.Execute(context.Background(), "brand device", "name device")
		require.Error(t, err)
		require.Nil(t, output)
	})
}
