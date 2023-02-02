package usecases

import (
	"api/src/category/domain"
	"api/src/common"
	"errors"
)

type InsertCategoryUseCase struct {
	Repo      CategoryRepository
	Validator common.Validate
}

func (useCase *InsertCategoryUseCase) Execute(req *domain.WriteCategoryBody) (*domain.Category, error) {
	err := useCase.Validator.Exec(req)
	if err != nil {
		return nil, errors.New("invalid data")
	}

	return useCase.Repo.Insert(req)
}
