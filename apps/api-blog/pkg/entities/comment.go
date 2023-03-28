package entities

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint
	UserID    uint
	User      User
	PostID    uint
	Post      Post
	ParentID  *uint
	Parent    *Comment
	Content   string `gorm:"size:1000"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type CommentReq struct {
	PostID   uint   `json:"post_id"`
	ParentID *uint  `json:"parent_id"`
	Content  string `json:"content"`
}

type CommentRes struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	PostID    uint      `json:"post_id"`
	ParentID  *uint     `json:"parent_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
