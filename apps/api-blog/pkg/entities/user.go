package entities

import (
	"time"
)

type User struct {
	Id          uint   `gorm:"primaryKey"`
	Identifier  string `gorm:"unique;size:50; not null"`
	Password    string `gorm:"size:250; not null"`
	Username    string `gorm:"unique; size:50; default:null"`
	Fullname    string `gorm:"size:50; not null"`
	PhoneNumber string `gorm:"size:50; not null"`
	Email       string `gorm:"size:50; not null"`
	Avatar      string `gorm:"size:250; not null"`
	CreatedAt   time.Time
	UpdateAt    time.Time
}

type UserReq struct {
	Password    string `json:"password"`
	Username    string `json:"username"`
	Fullname    string `json:"fullname"`
	PhoneNumber string `json:"phonenumber"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
}
type UserLogin struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type ResponseEmail struct {
	Link     string
	Username string
	Sender   string
}

type UserResponse struct {
	Username    string `json:"username"`
	Id          uint   `json:"id"`
	Fullname    string `json:"fullname"`
	PhoneNumber string `json:"phone"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
}

type SignUpBody struct {
	Password string  `json:"password"`
	Username *string `json:"username"`
	Fullname *string `json:"fullname"`
	Email    string  `json:"email"`
}
