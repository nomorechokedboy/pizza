package usecases

import "api/src/category/domain"

type CategoryRepository interface {
	Insert(req *domain.WriteCategoryBody) (*domain.Category, error)
	Update(id *int, req *domain.WriteCategoryBody) (*domain.Category, error)
	Delete(id *int) (*domain.Category, error)
	FindOne(id *int) (*domain.Category, error)
	Find(req *domain.CategoryQuery) ([]*domain.Category, error)
}
