package contracts

import (
	"context"

	"github.com/jhmorais/device-manager/internal/usecases/ports/output"
)

type FindDeviceByBrandUseCase interface {
	Execute(ctx context.Context, brand string) (*output.ListDeviceOutput, error)
}
