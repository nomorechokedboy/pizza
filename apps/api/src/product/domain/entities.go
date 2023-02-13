package domain

import (
	category "api/src/category/domain"
	"api/src/common"
	inventory "api/src/inventory/domain"
	"time"
)

type Product struct {
	ID          uint                `gorm:"primaryKey"`
	Slug        string              `gorm:"size:20;not null"`
	Description string              `gorm:"size:1000"`
	Name        string              `gorm:"size:20; not null"`
	SKU         string              `gorm:"unique;not null;size:20"`
	Price       float32             `gorm:"check:price > 0; not null"`
	Category    category.Category   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Inventory   inventory.Inventory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductReq struct {
	Description string  `js:"description"`
	Name        string  `js:"name"`
	SKU         string  `js:"sku"`
	Price       float32 `js:"price"`
	CategoryId  uint    `js:"categoryId"`
	InventoryId uint    `js:"inventoryId"`
}

type ProductQuery struct {
	common.BaseQuery
	InventoryId *uint `query:"inventoryId"`
}
