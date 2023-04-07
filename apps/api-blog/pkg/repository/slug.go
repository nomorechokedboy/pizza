package repository

import (
	"api-blog/pkg/entities"
)

type SlugRepository interface {
	GetSlugCount(slug string) (int64, error)
	CreateSlug(slug *entities.Slug) error
}
