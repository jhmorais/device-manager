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

func TestFindByBrandDevice(t *testing.T) {

	t.Run("when brand is empty should return an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("FindDeviceByBrand", mock.Anything, mock.Anything).Return(nil, nil)
		findDeviceUseCase := usecases.NewFindDeviceByBrandUseCase(deviceRepositoryMock)

		output, err := findDeviceUseCase.Execute(context.Background(), "")
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when db store returns an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("FindDeviceByBrand", mock.Anything, mock.Anything).Return(nil, errors.New("brand cannot be empty"))
		findDeviceUseCase := usecases.NewFindDeviceByBrandUseCase(deviceRepositoryMock)

		output, err := findDeviceUseCase.Execute(context.Background(), "")
		require.Error(t, err)
		require.Nil(t, output)
	})
}
