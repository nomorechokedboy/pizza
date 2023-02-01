package domain

import "time"

type Inventory struct {
	Id        int
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WriteInventoryBody struct {
	Quantity int `validate:"required,numeric,min=1"`
}
