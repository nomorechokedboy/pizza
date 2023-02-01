package repository

import "api/src/category/domain"

type CategoryInMemoryRepo struct {
	Data  []domain.Category
	IsErr bool
}

func (repo *CategoryInMemoryRepo) Insert(req *domain.WriteCategoryBody) (*domain.Category, error) {
	return nil, nil
}
