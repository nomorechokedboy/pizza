package repository

import (
	"api-blog/pkg/common"
	"api-blog/pkg/entities"
)

type CommentRepository interface {
	GetAllComments(query *entities.CommentQuery) (common.BasePaginationResponse[entities.Comment], error)
	CreateComment(comment *entities.Comment) error
	UpdateComment(comment *entities.Comment) error
	DeleteComment(id uint) error
}
