package entities

import (
	"api-blog/pkg/common"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint           `json:"id"`
	UserID    uint           `json:"-"`
	User      User           `json:"user"`
	PostID    uint           `json:"postId"`
	Post      Post           `json:"-"`
	ParentID  *uint          `json:"-"`
	Replies   []Comment      `json:"replies" gorm:"foreignkey:ParentID"`
	Content   string         `gorm:"size:1000" json:"content"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (comment *Comment) New() Comment {
	return Comment{
		ID:        comment.ID,
		UserID:    comment.UserID,
		User:      comment.User,
		PostID:    comment.PostID,
		Post:      comment.Post,
		ParentID:  comment.ParentID,
		Replies:   comment.Replies,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		DeletedAt: comment.DeletedAt,
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
