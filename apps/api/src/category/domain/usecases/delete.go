package usecases

import "api/src/category/domain"

type DeleteCategoryUseCase struct {
	Repo CategoryRepository
}

func (useCase *DeleteCategoryUseCase) Execute(id *int) (*domain.Category, error) {
	return nil, nil
}
