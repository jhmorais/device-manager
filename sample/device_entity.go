package sample

import (
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/jhmorais/device-manager/internal/domain/entities"
)

func init() {
	gofakeit.Seed(time.Now().UnixNano())
}

func NewTaskEntity() *entities.Device {
	return NewDeviceEntityWithUser("", "")
}

func NewDeviceEntityWithUser(brand, name string) *entities.Device {
	task := &entities.Device{
		ID:        RandomID(),
		Name:      name,
		Brand:     brand,
		CreatedAt: time.Now(),
	}
	return task
}
