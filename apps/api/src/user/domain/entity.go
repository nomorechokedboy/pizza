package domain

import (
	"api/src/common"
	"time"
)

type User struct {
	Id          int    `gorm:"primaryKey"`
	Identifier  string `gorm:"unique;not null;size:50"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	FullName    string `gorm:"size:20"`
	Gender      bool
	BirthDate   time.Time
	PhoneNumber string `gorm:"size:20"`
	Email       string `gorm:"size:20"`
	Password    string `gorm:"size:20"`
}

type CreateUserReq struct {
	Identifier  string    `json:"Identifier"`
	FullName    string    `json:"FullName"`
	Email       string    `json:"Email"`
	Password    string    `json:"Password"`
	PhoneNumber string    `json:"PhoneNumber"`
	Gender      bool      `json:"Gender"`
	BirthDate   time.Time `json:"BirthDate"`
}

type UserQuery struct {
	common.BaseQuery
}
