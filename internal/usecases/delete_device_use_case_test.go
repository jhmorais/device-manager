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

func TestDeleteDevice(t *testing.T) {

	t.Run("when id is empty should return an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("DeleteDevice", mock.Anything, mock.Anything).Return(nil)
		deleteDeviceUseCase := usecases.NewDeleteDeviceUseCase(deviceRepositoryMock)

		output, err := deleteDeviceUseCase.Execute(context.Background(), "")
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when id doesn't exists should return an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("DeleteDevice", mock.Anything, mock.Anything).Return(nil)
		deviceRepositoryMock.On("FindDeviceByID", mock.Anything, mock.Anything).Return(nil, errors.New("device not found"))
		deleteDeviceUseCase := usecases.NewDeleteDeviceUseCase(deviceRepositoryMock)

		output, err := deleteDeviceUseCase.Execute(context.Background(), "123654")
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when db store returns an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("DeleteDevice", mock.Anything, mock.Anything).Return(errors.New("database failed"))
		deleteDeviceUseCase := usecases.NewDeleteDeviceUseCase(deviceRepositoryMock)

		output, err := deleteDeviceUseCase.Execute(context.Background(), "987321654")
		require.Error(t, err)
		require.Nil(t, output)
	})
}
