package repository

import (
	"api-blog/pkg/common"
	"api-blog/pkg/entities"
)

type PostRepository interface {
	GetAllPosts(query *entities.PostQuery) (common.BasePaginationResponse[entities.Post], error)
	GetPostBySlug(slug string) (*entities.Post, error)
	CreatePost(post *entities.Post) (*entities.Post, error)
	UpdatePost(post *entities.Post) (*entities.Post, error)
	DeletePost(id uint) (*entities.Post, error)
}
