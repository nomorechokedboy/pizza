package usecases

import (
	"api/src/category/domain"
)

type InsertCategoryUseCase struct {
	Repo CategoryRepository
}

func (useCase *InsertCategoryUseCase) Execute(req *domain.WriteCategoryBody) (*domain.Category, error) {
	return useCase.Repo.Insert(req)
}
