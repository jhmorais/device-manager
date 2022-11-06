package usecases

import (
	"context"
	"fmt"

	"github.com/jhmorais/device-manager/internal/domain/entities"
	"github.com/jhmorais/device-manager/internal/repositories"
	"github.com/jhmorais/device-manager/internal/usecases/contracts"
	"github.com/jhmorais/device-manager/internal/usecases/ports/output"
)

type listDeviceUseCase struct {
	deviceRepository repositories.DeviceRepository
}

func NewListDeviceUseCase(deviceRepository repositories.DeviceRepository) contracts.ListDeviceUseCase {

	return &listDeviceUseCase{
		deviceRepository: deviceRepository,
	}
}

func (l *listDeviceUseCase) Execute(ctx context.Context) (*output.ListDeviceOutput, error) {
	var err error
	output := &output.ListDeviceOutput{Devices: []*entities.Device{}}

	output.Devices, err = l.deviceRepository.ListDevice(ctx)
	if err != nil {
		return nil, fmt.Errorf("error when list devices on database: %v", err)
	}

	return output, nil
}
