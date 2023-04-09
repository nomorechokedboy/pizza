package entities

import (
	"time"
)

type User struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	Identifier  string    `gorm:"unique;size:50; not null" json:"identifier"`
	Password    string    `gorm:"size:250; not null" json:"-"`
	Username    string    `gorm:"unique; size:50; default:null" json:"userName"`
	Fullname    string    `gorm:"size:50; not null" json:"fullName"`
	PhoneNumber string    `gorm:"size:50; not null" json:"phoneNumber"`
	Email       string    `gorm:"unique; size:50; not null" json:"email"`
	Avatar      string    `gorm:"size:250; not null" json:"avatar"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdateAt    time.Time `json:"updatedAt"`
}

type UserReq struct {
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

type UpdatePassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
