package repository

import (
	"api-blog/pkg/entities"
)

type CommentRepository interface {
	GetAllComments(userID uint, postID uint, parentID uint, page int, pageSize int) ([]entities.Comment, error)
	CreateComment(comment *entities.Comment) error
	UpdateComment(comment *entities.Comment) error
	DeleteComment(id uint) error
}
