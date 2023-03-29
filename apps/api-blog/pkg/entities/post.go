package entities

import (
	"time"
)

type Post struct {
	ID          uint
	UserID      uint
	User        User `json:"-"`
	ParentID    *uint
	Parent      *Post  `json:"-"`
	Title       string `gorm:"size:250"`
	Slug        string
	Content     string `gorm:"size:5000"`
	Published   bool   `gorm:"default:false"`
	PublishedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type PostReq struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	ParentID  *uint  `json:"parent_id"`
	Published bool   `json:"published"`
}

type PostRes struct {
	ID          uint       `json:"id"`
	User        *User      `json:"user"`
	Parent      *Post      `json:"parent"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Content     string     `json:"content"`
	Published   bool       `json:"published"`
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
	DeletedAt *time.Time
}
