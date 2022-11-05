package contracts

import (
	"context"

	"github.com/jhmorais/device-manager/internal/usecases/ports/output"
)

type DeleteDeviceUseCase interface {
	Execute(ctx context.Context, deviceID string) (*output.DeleteDeviceOutput, error)
}
