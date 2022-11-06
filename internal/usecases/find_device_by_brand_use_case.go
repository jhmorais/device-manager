package usecases

import (
	"context"
	"fmt"

	"github.com/jhmorais/device-manager/internal/repositories"
	"github.com/jhmorais/device-manager/internal/usecases/contracts"
	"github.com/jhmorais/device-manager/internal/usecases/ports/output"
)

type findDeviceByBrandUseCase struct {
	deviceRepository repositories.DeviceRepository
}

func NewFindDeviceByBrandUseCase(deviceRepository repositories.DeviceRepository) contracts.FindDeviceByBrandUseCase {

	return &findDeviceByBrandUseCase{
		deviceRepository: deviceRepository,
	}
}

func (c *findDeviceByBrandUseCase) Execute(ctx context.Context, brand string) (*output.ListDeviceOutput, error) {

	deviceEntity, err := c.deviceRepository.FindDeviceByBrand(ctx, brand)
	if err != nil {
		return nil, fmt.Errorf("erro to find device with brand: '%s' at database: '%v'", brand, err)
	}

	if len(deviceEntity) == 0 {
		return nil, fmt.Errorf("device not found")
	}

	output := &output.ListDeviceOutput{
		Devices: deviceEntity,
	}

	return output, nil
}
