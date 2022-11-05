package contracts

import (
	"context"

	"github.com/jhmorais/device-manager/internal/usecases/ports/input"
	"github.com/jhmorais/device-manager/internal/usecases/ports/output"
)

type CreateDeviceUseCase interface {
	Execute(ctx context.Context, createDevice *input.CreateDeviceInput) (*output.CreateDeviceOutput, error)
}
