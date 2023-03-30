package repository

import (
	"api-blog/pkg/entities"
)

type PostRepository interface {
	GetAllPosts(userID uint, parentID uint, page int, pageSize int) ([]entities.Post, error)
	GetPostBySlug(slug string) (*entities.Post, error)
	CreatePost(post *entities.Post) (uint, error)
	UpdatePost(post *entities.Post) error
	DeletePost(id uint) error
}
