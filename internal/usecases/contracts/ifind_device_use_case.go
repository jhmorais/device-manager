package contracts

import (
	"context"

	"github.com/jhmorais/device-manager/internal/usecases/ports/output"
)

type FindDeviceUseCase interface {
	Execute(ctx context.Context, brand, name string) (*output.FindDeviceOutput, error)
}
