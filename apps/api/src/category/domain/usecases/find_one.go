package usecases

import (
	"api/src/category/domain"
	"errors"
)

type FindOneCategoryUseCase struct {
	Repo CategoryRepository
}

func (useCase *FindOneCategoryUseCase) Execute(id *int) (*domain.Category, error) {
	inventory, err := useCase.Repo.FindOne(id)
	if inventory == nil && err == nil {
		return inventory, errors.New("not found")
	}

	return inventory, err
}
