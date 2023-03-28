package entities

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID          uint
	UserID      uint
	User        User
	ParentID    *uint `gorm:"default:null"`
	Parent      *Post
	Title       string `gorm:"size:250"`
	Content     string `gorm:"size:5000"`
	PublishedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type PostReq struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	ParentID *uint  `json:"parent_id"`
}

type PostRes struct {
	ID          uint       `json:"id"`
	UserID      uint       `json:"user_id"`
	ParentID    *uint      `json:"parent_id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Comment struct {
	ID        uint
	UserID    uint
	User      User
	PostID    uint
	Post      Post
	ParentID  uint
	Parent    *Comment
	Content   string `gorm:"size:1000"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
