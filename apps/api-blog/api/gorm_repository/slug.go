package gorm_repository

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/repository"

	"gorm.io/gorm"
)

type SlugGormRepo struct {
	db *gorm.DB
}

func NewSlugGormRepository(db *gorm.DB) repository.SlugRepository {
	return &SlugGormRepo{
		db: db,
	}
}

func (r *SlugGormRepo) GetSlugCount(slug string) (int64, error) {
	var slugs []entities.Slug

	foundSlugs := r.db.Find(&slugs, "slug LIKE ?", slug+"%")

	return foundSlugs.RowsAffected, foundSlugs.Error
}

func (r *SlugGormRepo) CreateSlug(slug *entities.Slug) error {
	return r.db.Create(&slug).Error
}
