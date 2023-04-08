package entities

import (
	"api-blog/pkg/common"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID          uint           `json:"id"`
	UserID      uint           `json:"userId"`
	User        User           `json:"-"`
	ParentID    *uint          `json:"parentId"`
	Parent      *Post          `json:"-"`
	Image       *string        `json:"image"`
	Comments    []Comment      `json:"comments"`
	Title       string         `gorm:"size:250" json:"title"`
	Slug        string         `json:"slug"`
	Content     string         `gorm:"size:5000" json:"content"`
	Published   bool           `json:"published"`
	PublishedAt *time.Time     `json:"publishedAt"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

type PostRes struct {
	ID          uint       `json:"id"`
	User        User       `json:"user"`
	Parent      *Post      `json:"parent"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Image       *string    `json:"image"`
	Comments    []Comment  `json:"comments"`
	Content     string     `json:"content"`
	Published   bool       `json:"published"`
	PublishedAt *time.Time `json:"publishedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func (post *Post) ToResponse() PostRes {
	return PostRes{
		ID:          post.ID,
		Slug:        post.Slug,
		User:        post.User,
		Title:       post.Title,
		Image:       post.Image,
		Parent:      post.Parent,
		Content:     post.Content,
		Comments:    post.Comments,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Published:   post.Published,
		PublishedAt: post.PublishedAt,
	}
}

type PostReq struct {
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	Image     *string `json:"image"`
	ParentID  *uint   `json:"parentId"`
	Published bool    `json:"published"`
}

type PostQuery struct {
	common.BaseQuery
	UserID   uint
	ParentID uint
}
