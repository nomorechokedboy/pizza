package domain

import (
	category "api/src/category/domain"
	"api/src/common"
	inventory "api/src/inventory/domain"
	"time"
)

type Product struct {
	Id          int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Slug        string
	Description string
	Name        string
	SKU         string
	Price       float32
	Category    category.Category
	Inventory   inventory.Inventory
}

type ProductReq struct {
	Description string  `js:"description"`
	Name        string  `js:"name"`
	SKU         string  `js:"sku"`
	Price       float32 `js:"price"`
	CategoryId  int     `js:"categoryId"`
	InventoryId int     `js:"inventoryId"`
}

type ProductQuery struct {
	Base        common.BaseQuery
	InventoryId *int `query:"inventoryId"`
}
