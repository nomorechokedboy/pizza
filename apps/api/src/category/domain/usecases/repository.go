package usecases

import "api/src/category/domain"

type CategoryRepository interface {
	Insert(req *domain.WriteCategoryBody) (*domain.Category, error)
}
