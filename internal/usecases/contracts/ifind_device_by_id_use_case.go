package contracts

import (
	"context"

	"github.com/jhmorais/device-manager/internal/usecases/ports/output"
)

type FindDeviceByIDUseCase interface {
	Execute(ctx context.Context, deviceID string) (*output.FindDeviceOutput, error)
}
