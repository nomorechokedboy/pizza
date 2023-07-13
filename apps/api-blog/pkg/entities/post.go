package entities

import (
	"api-blog/pkg/common"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID            uint           `json:"id"`
	UserID        uint           `json:"userId"`
	User          User           `json:"-"`
	ParentID      *uint          `json:"parentId"`
	Parent        *Post          `json:"-"`
	Image         *string        `json:"image"`
	Comments      []Comment      `json:"-"`
	CommentCount  int            `json:"-"`
	Title         string         `gorm:"size:250" json:"title"`
	Slug          string         `json:"slug"`
	Content       string         `gorm:"size:5000" json:"content"`
	Published     bool           `json:"published"`
	PublishedAt   *time.Time     `json:"publishedAt"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"-"`
	Reactions     []Reaction     `json:"-" gorm:"polymorphic:Reactable;"`
	ReactionCount uint           `json:"reactionCount"`
}

type PostResponse struct {
	ID            uint       `json:"id"`
	User          User       `json:"user"`
	Parent        *Post      `json:"parent"`
	Title         string     `json:"title"`
	Slug          string     `json:"slug"`
	Image         *string    `json:"image"`
	CommentCount  int        `json:"commentCount"`
	Content       string     `json:"content"`
	Published     bool       `json:"published"`
	PublishedAt   *time.Time `json:"publishedAt"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	ReactionCount uint       `json:"reactionCount"`
	Reactions     []Reaction `json:"reactions" gorm:"polymorphic:Reactable;"`
}

func (post *Post) ToResponse() PostResponse {
	return PostResponse{
		ID:            post.ID,
		User:          post.User,
		Parent:        post.Parent,
		Title:         post.Title,
		Slug:          post.Slug,
		Image:         post.Image,
		CommentCount:  post.CommentCount,
		Content:       post.Content,
		Published:     post.Published,
		PublishedAt:   post.PublishedAt,
		CreatedAt:     post.CreatedAt,
		UpdatedAt:     post.UpdatedAt,
		ReactionCount: post.ReactionCount,
		Reactions:     post.Reactions,
	}
}

type PostRequest struct {
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
