package entities

import "time"

type Device struct {
	ID        string `gorm:"id"`
	Name      string `gorm:"size:250"`
	Brand     string `gorm:"index"`
	CreatedAt time.Time
}
