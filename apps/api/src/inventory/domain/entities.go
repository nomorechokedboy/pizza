package domain

import "time"

type Inventory struct {
	ID        uint `gorm:"primaryKey"`
	Quantity  uint `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WriteInventoryBody struct {
	Quantity uint `validate:"required,numeric,min=1"`
}
