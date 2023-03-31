package entities

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint           `json:"id"`
	UserID    uint           `json:"user_id"`
	User      User           `json:"-"`
	PostID    uint           `json:"post_id"`
	Post      Post           `json:"-"`
	ParentID  *uint          `json:"parent_id"`
	Parent    *Comment       `json:"-"`
	Content   string         `gorm:"size:1000" json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type CommentRes struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Parent    *Comment  `json:"parent"`
	PostID    uint      `json:"post_id"`
	Content   string    `gorm:"size:1000" json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (comment *Comment) ToResponse() CommentRes {
	return CommentRes{
		ID:        comment.ID,
		PostID:    comment.PostID,
		User:      comment.User,
		Parent:    comment.Parent,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}

type CommentReq struct {
	PostID   uint   `json:"post_id"`
	ParentID *uint  `json:"parent_id"`
	Content  string `json:"content"`
}

type CommentPaginationResponse struct {
	Comments []CommentRes `json:"comments"`
	Page     int          `json:"page"`
	PageSize int          `json:"page_size"`
	Total    int          `json:"total"`
}
