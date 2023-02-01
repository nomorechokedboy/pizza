package usecases

import (
	"api/src/category/domain"
)

type UpdateCategoryUseCase struct {
	Repo CategoryRepository
}

func (useCase *UpdateCategoryUseCase) Execute(req *domain.WriteCategoryBody) (*domain.Category, error) {
	return useCase.Repo.Update(req)
}
