package domain

import "time"

type CartItem struct {
	ProductID uint      `js:"productId"`
	Quantity  uint      `js:"quantity"`
	CreatedAt time.Time `js:"createdAt"`
	UpdatedAt time.Time `js:"updatedAt"`
}

type WriteCartItemRequest struct {
	ProductID uint `validate:"required,numeric,min=0"`
	Quantity  uint `validate:"required,numeric,min=1"`
}
