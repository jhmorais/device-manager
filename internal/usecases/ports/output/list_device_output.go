package output

import "github.com/jhmorais/device-manager/internal/domain/entities"

type ListDeviceOutput struct {
	Devices []*entities.Device
}
