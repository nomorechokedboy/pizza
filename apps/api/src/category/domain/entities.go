package domain

import "api/src/common"

type Category struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"unique;not null;size:20"`
	Description *string `gorm:"size:1000"`
}

type CategoryQuery struct {
	common.BaseQuery
}

type WriteCategoryBody struct {
	Name        string `validate:"required,min=3,max=20"`
	Description string `validate:"required,max=1000"`
}
