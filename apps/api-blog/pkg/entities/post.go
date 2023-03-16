package entities

import (
	"time"
)

type Post struct {
	ID          uint
	UserID      uint
	User        User
	ParentID    *uint
	Parent      *Post
	Title       string `gorm:"size:250"`
	Content     string `gorm:"size:5000"`
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type PostReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

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
	DeletedAt time.Time
}

type Slug struct {
	Slug   string `gorm:"size:300"`
	PostID uint
	Post   Post
}
