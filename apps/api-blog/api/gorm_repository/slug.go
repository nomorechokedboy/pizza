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
	var count int64

	if err := r.db.Where("slug LIKE ?", "%"+slug+"%").Find(&slugs).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (r *SlugGormRepo) CreateSlug(slug *entities.Slug) error {
	return r.db.Create(&slug).Error
}
