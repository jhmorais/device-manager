package input

import "time"

type CreateDeviceInput struct {
	Name      string
	Brand     string
	CreatedAt time.Time
}
