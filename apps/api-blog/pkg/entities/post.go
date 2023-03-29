package entities

import (
	"time"

	"gorm.io/gorm"
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
	DeletedAt   gorm.DeletedAt
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

type PostPaginationResponse struct {
	Posts    []PostRes `json:"posts"`
	Page     int       `json:"page"`
	PageSize int       `json:"page_size"`
	Total    int       `json:"total"`
}
