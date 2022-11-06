package contracts

import (
	"context"

	"github.com/jhmorais/device-manager/internal/usecases/ports/input"
	"github.com/jhmorais/device-manager/internal/usecases/ports/output"
)

type UpdateDeviceUseCase interface {
	Execute(ctx context.Context, updateDevice *input.UpdateDeviceInput) (*output.CreateDeviceOutput, error)
}
