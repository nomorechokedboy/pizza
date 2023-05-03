package entities

import (
	"time"
)

type Reaction struct {
	UserID        uint      `json:"-" gorm:"primaryKey;autoIncrement:false"`
	User          User      `json:"user"`
	ReactableID   uint      `json:"-" gorm:"primaryKey;autoIncrement:false;"`
	ReactableType string    `json:"-" gorm:"primaryKey"`
	CreatedAt     time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"not null"`
}

type WriteReactionBody struct {
	ReactableID   uint   `json:"reactableId" validate:"required"`
	ReactableType string `json:"reactableType" validate:"required,oneof=posts comments"`
}
