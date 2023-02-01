package domain

import "time"

//SKU - stock keeping unit
type User struct {
	Id        int32
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Password  string
}

type CreateUserReq struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	UserName  string `json:"UserName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
}
