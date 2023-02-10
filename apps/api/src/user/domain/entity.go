package domain

import "time"

type User struct {
	Id          int
	Identifier  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	FullName    string
	Gender      bool
	BirthDate   time.Time
	PhoneNumber string
	Email       string
	Password    string
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
	Page     int    `query:"page"`
	PageSize int    `query:"pageSize"`
	Q        string `query:"q"`
}
