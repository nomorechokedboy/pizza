package domain

import "time"

type Product struct {
	Id          int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Slug        string
	Description string
	Name        string
	SKU         string
	Price       float32
}

type ProductReq struct {
	Description string  `js:"description"`
	Name        string  `js:"name"`
	SKU         string  `js:"sku"`
	Price       float32 `js:"price"`
}
