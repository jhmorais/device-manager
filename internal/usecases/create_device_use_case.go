package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jhmorais/device-manager/internal/domain/entities"
	"github.com/jhmorais/device-manager/internal/repositories"
	"github.com/jhmorais/device-manager/internal/usecases/contracts"
	"github.com/jhmorais/device-manager/internal/usecases/ports/input"
	"github.com/jhmorais/device-manager/internal/usecases/ports/output"
)

type createDeviceUseCase struct {
	deviceRepository repositories.DeviceRepository
}

func NewCreateDeviceUseCase(deviceRepository repositories.DeviceRepository) contracts.CreateDeviceUseCase {

	return &createDeviceUseCase{
		deviceRepository: deviceRepository,
	}
}

func (c *createDeviceUseCase) Execute(ctx context.Context, createDevice *input.CreateDeviceInput) (*output.CreateDeviceOutput, error) {

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("cannot generate a new device ID: %v", err)
	}

	//max 250
	if len(createDevice.Brand) > 250 {
		//will discard the rest
		createDevice.Brand = createDevice.Brand[:250]
	}

	if createDevice.Name == "" {
		return nil, fmt.Errorf("cannot create a device without name")
	}

	if createDevice.Brand == "" {
		return nil, fmt.Errorf("cannot create a device without brand")
	}

	deviceEntity := &entities.Device{
		ID:        id.String(),
		Name:      createDevice.Name,
		Brand:     createDevice.Brand,
		CreatedAt: time.Now(),
	}

	err = c.deviceRepository.CreateDevice(ctx, deviceEntity)

	if err != nil {
		return nil, fmt.Errorf("cannot save device at database: %v", err)
	}

	createDeviceOutput := &output.CreateDeviceOutput{
		DeviceID: deviceEntity.ID,
	}

	return createDeviceOutput, nil
}
