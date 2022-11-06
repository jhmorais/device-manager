package sample

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jhmorais/device-manager/internal/domain/entities"
	"github.com/jhmorais/device-manager/internal/repositories"
	"gorm.io/gorm"
)

func DBSeed(db *gorm.DB) error {
	deviceRepository := repositories.NewDeviceRepository(db)

	err := createDevice(deviceRepository, "Iphone 14", "APPLE")
	if err != nil {
		return err
	}

	err = createDevice(deviceRepository, "Iphone 13", "APPLE")
	if err != nil {
		return err
	}

	err = createDevice(deviceRepository, "Galaxy S21", "SAMSUNG")
	if err != nil {
		return err
	}

	err = createDevice(deviceRepository, "Mi", "XIAOMI")
	if err != nil {
		return err
	}

	return nil
}

func createDevice(deviceRepository repositories.DeviceRepository, name string, brand string) error {
	ctx := context.Background()
	device, err := deviceRepository.FindDevice(ctx, brand, name)
	if err != nil {
		return err
	}

	_, err = deviceRepository.ListDevice(ctx)
	if err != nil {
		return err
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	if device.ID == "" {
		device = &entities.Device{
			ID:        id.String(),
			Name:      name,
			Brand:     brand,
			CreatedAt: time.Now(),
		}
		err := deviceRepository.CreateDevice(context.Background(), device)
		if err != nil {
			return err
		}
	}

	return nil
}
