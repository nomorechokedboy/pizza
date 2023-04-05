package repository

import (
	"api-blog/pkg/common"
	"api-blog/pkg/entities"
)

type PostRepository interface {
	GetAllPosts(query *entities.PostQuery) (common.BasePaginationResponse[entities.Post], error)
	GetPostBySlug(slug string) (*entities.Post, error)
	CreatePost(post *entities.Post) (uint, error)
	UpdatePost(post *entities.Post) error
	DeletePost(id uint) error
}
