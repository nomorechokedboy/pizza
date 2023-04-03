package entities

import (
	"api-blog/pkg/common"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID          uint           `json:"id"`
	UserID      uint           `json:"user_id"`
	User        User           `json:"-"`
	ParentID    *uint          `json:"parent_id"`
	Parent      *Post          `json:"-"`
	Title       string         `gorm:"size:250" json:"title"`
	Slug        string         `json:"slug"`
	Content     string         `gorm:"size:5000" json:"content"`
	Published   bool           `json:"published"`
	PublishedAt *time.Time     `json:"published_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

type PostRes struct {
	ID          uint       `json:"id"`
	User        User       `json:"user"`
	Parent      *Post      `json:"parent"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Content     string     `json:"content"`
	Published   bool       `json:"published"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (post *Post) ToResponse() PostRes {
	return PostRes{
		ID:          post.ID,
		User:        post.User,
		Parent:      post.Parent,
		Title:       post.Title,
		Slug:        post.Slug,
		Content:     post.Content,
		Published:   post.Published,
		PublishedAt: post.PublishedAt,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}
}

type PostReq struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	ParentID  *uint  `json:"parent_id"`
	Published bool   `json:"published"`
}

type PostQuery struct {
	common.BaseQuery
	UserID   uint
	ParentID uint
}
