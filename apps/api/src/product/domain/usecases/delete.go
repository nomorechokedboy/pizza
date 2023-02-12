package usecases

import (
	"api/src/common"
	"api/src/product/domain"
	"errors"
)

type DeleteProductUseCase struct {
	Repo      ProductRepository
	Validator *common.Validate
}

func (useCase *DeleteProductUseCase) Execute(id uint) (*domain.Product, error) {
	deletedProduct, err := useCase.Repo.Delete(id)
	if deletedProduct == nil && err == nil {
		return deletedProduct, errors.New("not found")
	}

	return deletedProduct, err
}
