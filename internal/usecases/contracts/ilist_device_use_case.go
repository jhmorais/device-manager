package contracts

import (
	"context"

	"github.com/jhmorais/device-manager/internal/usecases/ports/output"
)

type ListDeviceUseCase interface {
	Execute(ctx context.Context) (*output.ListDeviceOutput, error)
}
