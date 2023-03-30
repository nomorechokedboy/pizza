package entities

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID          uint           `json:"id"`
	UserID      uint           `json:"-"`
	User        User           `json:"user"`
	ParentID    *uint          `json:"-"`
	Parent      *Post          `json:"parent"`
	Title       string         `gorm:"size:250" json:"title"`
	Slug        string         `json:"slug"`
	Content     string         `gorm:"size:5000" json:"content"`
	Published   bool           `json:"published"`
	PublishedAt *time.Time     `json:"published_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

type PostReq struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	ParentID  *uint  `json:"parent_id"`
	Published bool   `json:"published"`
}

type PostPaginationResponse struct {
	Posts    []Post `json:"posts"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Total    int    `json:"total"`
}
