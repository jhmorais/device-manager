package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/jhmorais/device-manager/internal/domain/entities"
	"github.com/jhmorais/device-manager/internal/repositories"
	"github.com/jhmorais/device-manager/internal/usecases/contracts"
	"github.com/jhmorais/device-manager/internal/usecases/ports/input"
	"github.com/jhmorais/device-manager/internal/usecases/ports/output"
	"github.com/jhmorais/device-manager/internal/usecases/validator"
)

type updateDeviceUseCase struct {
	deviceRepository repositories.DeviceRepository
}

func NewUpdateDeviceUseCase(deviceRepository repositories.DeviceRepository) contracts.UpdateDeviceUseCase {

	return &updateDeviceUseCase{
		deviceRepository: deviceRepository,
	}
}

func (c *updateDeviceUseCase) Execute(ctx context.Context, updateDevice *input.UpdateDeviceInput) (*output.CreateDeviceOutput, error) {

	if err := validator.ValidateUUId(updateDevice.ID, true, "deviceId"); err != nil {
		return nil, err
	}

	//max 250
	if len(updateDevice.Brand) > 250 {
		//will discard the rest
		updateDevice.Brand = updateDevice.Brand[:250]
	}

	deviceEntity := &entities.Device{
		ID:        updateDevice.ID,
		Name:      updateDevice.Name,
		Brand:     updateDevice.Brand,
		CreatedAt: time.Now(),
	}

	errUpdate := c.deviceRepository.UpdateDevice(ctx, deviceEntity)
	if errUpdate != nil {
		return nil, fmt.Errorf("cannot save device at database: %v", errUpdate)
	}

	createDeviceOutput := &output.CreateDeviceOutput{
		DeviceID: deviceEntity.ID,
	}

	return createDeviceOutput, nil
}
