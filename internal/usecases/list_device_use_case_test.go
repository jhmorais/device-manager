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

func TestListDevice(t *testing.T) {

	t.Run("when db store returns an error", func(t *testing.T) {
		t.Parallel()
		deviceRepositoryMock := &mocks.DeviceRepository{}
		deviceRepositoryMock.On("ListDevice", mock.Anything).Return(nil, errors.New("database failed"))
		ListDeviceUseCase := usecases.NewListDeviceUseCase(deviceRepositoryMock)

		output, err := ListDeviceUseCase.Execute(context.Background())
		require.Error(t, err)
		require.Nil(t, output)
	})
}
