package usecases

import (
	"api/src/category/domain"
	"errors"
)

type DeleteCategoryUseCase struct {
	Repo CategoryRepository
}

func (useCase *DeleteCategoryUseCase) Execute(id *int) (*domain.Category, error) {
	deletedCategory, err := useCase.Repo.Delete(id)
	if deletedCategory == nil && err == nil {
		return deletedCategory, errors.New("not found")
	}

	return deletedCategory, err
}
