package usecases

import (
	"api/src/common"
	"api/src/product/domain"
)

type DeleteProductUseCase struct {
	Repo      ProductRepository
	Validator *common.Validate
}

func (useCase *DeleteProductUseCase) Execute(id *int) (*domain.Product, error) {
	return nil, nil
}
