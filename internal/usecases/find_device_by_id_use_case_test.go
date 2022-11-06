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

func TestFindByIDDevice(t *testing.T) {

	t.Run("when id is empty should return an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("FindDeviceByID", mock.Anything, mock.Anything).Return(nil, nil)
		findDeviceUseCase := usecases.NewDeleteDeviceUseCase(deviceRepositoryMock)

		output, err := findDeviceUseCase.Execute(context.Background(), "")
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when id has invalid value should return an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("FindDeviceByID", mock.Anything, mock.Anything).Return(nil, nil)
		findDeviceUseCase := usecases.NewDeleteDeviceUseCase(deviceRepositoryMock)

		output, err := findDeviceUseCase.Execute(context.Background(), "a123b456c789")
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when db store returns an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("DeleteDevice", mock.Anything, mock.Anything).Return(nil, errors.New("database failed"))
		findDeviceUseCase := usecases.NewDeleteDeviceUseCase(deviceRepositoryMock)

		output, err := findDeviceUseCase.Execute(context.Background(), "987321654")
		require.Error(t, err)
		require.Nil(t, output)
	})
}
