package usecases

import (
	"api/src/category/domain"
	"errors"
)

type FindOneCategoryUseCase struct {
	Repo CategoryRepository
}

func (useCase *FindOneCategoryUseCase) Execute(id *int) (*domain.Category, error) {
	category, err := useCase.Repo.FindOne(id)
	if category == nil && err == nil {
		return category, errors.New("not found")
	}

	return category, err
}
