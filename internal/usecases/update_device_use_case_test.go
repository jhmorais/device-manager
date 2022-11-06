package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jhmorais/device-manager/internal/domain/entities"
	"github.com/jhmorais/device-manager/internal/usecases"
	"github.com/jhmorais/device-manager/internal/usecases/ports/input"
	"github.com/jhmorais/device-manager/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdateDevice(t *testing.T) {

	t.Run("when the name is empty should return an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("UpdateDevice", mock.Anything, mock.Anything).Return(nil)
		updateDeviceUseCase := usecases.NewUpdateDeviceUseCase(deviceRepositoryMock)

		updateDeviceInput := &input.UpdateDeviceInput{
			Name:  "",
			Brand: "APPLE",
		}

		output, err := updateDeviceUseCase.Execute(context.Background(), updateDeviceInput)
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when the brand is empty should return an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("UpdateDevice", mock.Anything, mock.Anything).Return(nil)
		updateDeviceUseCase := usecases.NewUpdateDeviceUseCase(deviceRepositoryMock)

		updateDeviceInput := &input.UpdateDeviceInput{
			Name:  "Iphone 11",
			Brand: "",
		}

		output, err := updateDeviceUseCase.Execute(context.Background(), updateDeviceInput)
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when exist device with the same name and brand should return an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("UpdateDevice", mock.Anything, mock.Anything).Return(nil)
		updateDeviceUseCase := usecases.NewUpdateDeviceUseCase(deviceRepositoryMock)

		deviceRepositoryMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).
			Return(&entities.Device{ID: "1234526789"}, nil)

		updateDeviceInput := &input.UpdateDeviceInput{
			Name:  "Iphone 14",
			Brand: "APPLE",
		}

		output, err := updateDeviceUseCase.Execute(context.Background(), updateDeviceInput)
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when db store returns an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("UpdateDevice", mock.Anything, mock.Anything).Return(errors.New("database failed"))
		deviceRepositoryMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, nil)

		updateDeviceUseCase := usecases.NewUpdateDeviceUseCase(deviceRepositoryMock)

		updateDeviceInput := &input.UpdateDeviceInput{
			Name:  "Iphone 14",
			Brand: "APPLE",
		}

		output, err := updateDeviceUseCase.Execute(context.Background(), updateDeviceInput)
		require.Error(t, err)
		require.Nil(t, output)
	})
}
