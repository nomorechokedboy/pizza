package repository

import (
	"api-blog/pkg/entities"
)

type CommentRepository interface {
	GetAllComments() ([]entities.Comment, error)
	GetAllCommentsByQuery(userID uint, postID uint, parentID uint) ([]entities.Comment, error)
	CreateComment(comment *entities.Comment) error
	UpdateComment(comment *entities.Comment) error
	DeleteComment(id uint) error
}
