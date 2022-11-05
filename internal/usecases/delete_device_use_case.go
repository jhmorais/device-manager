package usecases

import (
	"context"
	"fmt"

	"github.com/jhmorais/device-manager/internal/repositories"
	"github.com/jhmorais/device-manager/internal/usecases/contracts"
	"github.com/jhmorais/device-manager/internal/usecases/ports/output"
	"github.com/jhmorais/device-manager/internal/usecases/validator"
)

type deleteDeviceUseCase struct {
	deviceRepository repositories.DeviceRepository
}

func NewDeleteDeviceUseCase(deviceRepository repositories.DeviceRepository) contracts.DeleteDeviceUseCase {

	return &deleteDeviceUseCase{
		deviceRepository: deviceRepository,
	}
}

func (c *deleteDeviceUseCase) Execute(ctx context.Context, deviceID string) (*output.DeleteDeviceOutput, error) {

	if err := validator.ValidateUUId(deviceID, true, "deviceId"); err != nil {
		return nil, err
	}

	deviceEntity, err := c.deviceRepository.FindDevice(ctx, deviceID)
	if err != nil {
		return nil, fmt.Errorf("failed to find device '%s' at database: '%v'", deviceID, err)
	}

	if deviceEntity == nil || deviceEntity.ID == "" {
		return nil, fmt.Errorf("deviceID not found")
	}

	err = c.deviceRepository.DeleteDevice(ctx, deviceEntity)
	if err != nil {
		return nil, fmt.Errorf("failed to delete device '%s'", deviceEntity.ID)
	}

	output := &output.DeleteDeviceOutput{
		DeviceID:   deviceEntity.ID,
		DeviceName: deviceEntity.Name,
	}

	return output, nil
}
