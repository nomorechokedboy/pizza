package domain

import (
	category "api/src/category/domain"
	"api/src/common"
	inventory "api/src/inventory/domain"
	"time"
)

type Product struct {
	CategoryID  uint
	InventoryID uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ID          uint                `gorm:"primaryKey"`
	Slug        string              `gorm:"size:20;not null"`
	Description *string             `gorm:"size:1000"`
	Name        string              `gorm:"size:20; not null"`
	SKU         string              `gorm:"unique;not null;size:20"`
	Price       float32             `gorm:"check:price > 0; not null"`
	Category    category.Category   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Inventory   inventory.Inventory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ProductReq struct {
	Description *string `js:"description" validate:"required,max=1000"`
	Name        string  `js:"name" validate:"required,min=3,max=50"`
	SKU         string  `js:"sku" validate:"required,min=3,max=20"`
	Price       float32 `js:"price" validate:"required,min=1"`
	CategoryId  uint    `js:"categoryId" validate:"required,min=1"`
	Quantity    uint    `js:"quantity" validate:"required,min=1"`
}

type ProductQuery struct {
	common.BaseQuery
	CategoryId *uint `query:"categoryId" validate:"min=0"`
}
