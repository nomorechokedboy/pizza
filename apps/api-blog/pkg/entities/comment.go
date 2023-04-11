package entities

import (
	"api-blog/pkg/common"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint           `json:"id"`
	UserID    uint           `json:"userId"`
	User      User           `json:"-"`
	PostID    uint           `json:"postId"`
	Post      Post           `json:"-"`
	ParentID  *uint          `json:"parentId"`
	Parent    *Comment       `json:"-"`
	Content   string         `gorm:"size:1000" json:"content"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type CommentResponse struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Parent    *Comment  `json:"parent"`
	PostID    uint      `json:"postId"`
	Content   string    `gorm:"size:1000" json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (comment *Comment) ToResponse() CommentResponse {
	return CommentResponse{
		ID:        comment.ID,
		PostID:    comment.PostID,
		User:      comment.User,
		Parent:    comment.Parent,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}

type CommentRequest struct {
	PostID   uint   `json:"postId"`
	ParentID *uint  `json:"parentId"`
	Content  string `json:"content"`
}

type CommentQuery struct {
	common.BaseQuery
	ParentID uint
	UserID   uint
	PostID   uint
}
