package usecases

import (
	"context"
	"fmt"

	"github.com/jhmorais/device-manager/internal/repositories"
	"github.com/jhmorais/device-manager/internal/usecases/contracts"
	"github.com/jhmorais/device-manager/internal/usecases/ports/output"
	"github.com/jhmorais/device-manager/internal/usecases/validator"
)

type findDeviceByIDUseCase struct {
	deviceRepository repositories.DeviceRepository
}

func NewFindDeviceByIDUseCase(deviceRepository repositories.DeviceRepository) contracts.FindDeviceByIDUseCase {

	return &findDeviceByIDUseCase{
		deviceRepository: deviceRepository,
	}
}

func (c *findDeviceByIDUseCase) Execute(ctx context.Context, deviceID string) (*output.FindDeviceOutput, error) {

	if err := validator.ValidateUUId(deviceID, true, "deviceId"); err != nil {
		return nil, err
	}

	deviceEntity, err := c.deviceRepository.FindDeviceByID(ctx, deviceID)
	if err != nil {
		return nil, fmt.Errorf("erro to find device '%s' at database: '%v'", deviceID, err)
	}

	if deviceEntity == nil || deviceEntity.ID == "" {
		return nil, fmt.Errorf("deviceID not found")
	}

	output := &output.FindDeviceOutput{
		Device: deviceEntity,
	}

	return output, nil
}
