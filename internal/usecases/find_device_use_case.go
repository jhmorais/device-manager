package usecases

import (
	"context"
	"fmt"

	"github.com/jhmorais/device-manager/internal/repositories"
	"github.com/jhmorais/device-manager/internal/usecases/contracts"
	"github.com/jhmorais/device-manager/internal/usecases/ports/output"
)

type findDeviceUseCase struct {
	deviceRepository repositories.DeviceRepository
}

func NewFindDeviceUseCase(deviceRepository repositories.DeviceRepository) contracts.FindDeviceUseCase {

	return &findDeviceUseCase{
		deviceRepository: deviceRepository,
	}
}

func (c *findDeviceUseCase) Execute(ctx context.Context, brand, name string) (*output.FindDeviceOutput, error) {

	deviceEntity, err := c.deviceRepository.FindDevice(ctx, brand, name)
	if err != nil {
		return nil, fmt.Errorf("erro to find device with name: '%s' and brand: '%s' at database: '%v'", name, brand, err)
	}

	if deviceEntity == nil || deviceEntity.ID == "" {
		return nil, fmt.Errorf("device not found")
	}

	output := &output.FindDeviceOutput{
		Device: deviceEntity,
	}

	return output, nil
}
