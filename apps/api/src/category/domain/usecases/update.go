package usecases

import (
	"api/src/category/domain"
	"errors"
)

type UpdateCategoryUseCase struct {
	Repo CategoryRepository
}

func (useCase *UpdateCategoryUseCase) Execute(id *int, req *domain.WriteCategoryBody) (*domain.Category, error) {
	updatedCategory, err := useCase.Repo.Update(id, req)
	if updatedCategory == nil && err == nil {
		return updatedCategory, errors.New("not found")
	}

	return updatedCategory, err
}
