package repository

import (
	"api-blog/pkg/entities"
)

type PostRepository interface {
	GetAllPosts() ([]entities.Post, error)
	GetAllPostsByUserID(userID uint) ([]entities.Post, error)
	GetAllPostsByParentID(parentID uint) ([]entities.Post, error)
	GetPostBySlug(slug string) (*entities.Post, error)
	CreatePost(post *entities.Post) (uint, error)
	UpdatePost(post *entities.Post) error
	DeletePost(id uint) error
}
