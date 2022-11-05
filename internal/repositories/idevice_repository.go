package repositories

import (
	"context"

	"github.com/jhmorais/device-manager/internal/domain/entities"
)

type DeviceRepository interface {
	CreateDevice(ctx context.Context, entity *entities.Device) error
	DeleteDevice(ctx context.Context, entity *entities.Device) error
	FindDeviceByID(ctx context.Context, id string) (*entities.Device, error)
	FindDevice(ctx context.Context, brand, name string) (*entities.Device, error)
	ListDevice(ctx context.Context) ([]*entities.Device, error)
}
