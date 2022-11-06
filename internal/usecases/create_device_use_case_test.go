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

func TestCreateDevice(t *testing.T) {

	t.Run("when the name is empty should return an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("CreateDevice", mock.Anything, mock.Anything).Return(nil)
		createDeviceUseCase := usecases.NewCreateDeviceUseCase(deviceRepositoryMock)

		createDeviceInput := &input.CreateDeviceInput{
			Name:  "",
			Brand: "APPLE",
		}

		output, err := createDeviceUseCase.Execute(context.Background(), createDeviceInput)
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when the brand is empty should return an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("CreateDevice", mock.Anything, mock.Anything).Return(nil)
		createDeviceUseCase := usecases.NewCreateDeviceUseCase(deviceRepositoryMock)

		createDeviceInput := &input.CreateDeviceInput{
			Name:  "Iphone 11",
			Brand: "",
		}

		output, err := createDeviceUseCase.Execute(context.Background(), createDeviceInput)
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when exist device with the same name and brand should return an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("CreateDevice", mock.Anything, mock.Anything).Return(nil)
		createDeviceUseCase := usecases.NewCreateDeviceUseCase(deviceRepositoryMock)

		createDeviceInput := &input.CreateDeviceInput{
			Name:  "Iphone 14",
			Brand: "APPLE",
		}

		deviceRepositoryMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).
			Return(&entities.Device{ID: "1234526789"}, nil)

		output, err := createDeviceUseCase.Execute(context.Background(), createDeviceInput)
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when db store returns an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("CreateDevice", mock.Anything, mock.Anything).Return(errors.New("database failed"))
		deviceRepositoryMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, nil)

		createDeviceUseCase := usecases.NewCreateDeviceUseCase(deviceRepositoryMock)

		createDeviceInput := &input.CreateDeviceInput{
			Name:  "Iphone 14",
			Brand: "APPLE",
		}

		output, err := createDeviceUseCase.Execute(context.Background(), createDeviceInput)
		require.Error(t, err)
		require.Nil(t, output)
	})
}
