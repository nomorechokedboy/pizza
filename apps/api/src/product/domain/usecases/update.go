package usecases

import (
	"api/src/common"
	"api/src/product/domain"
	"errors"
)

type UpdateProductUseCase struct {
	Repo      ProductRepository
	Validator *common.Validate
}

func (useCase *UpdateProductUseCase) Execute(id uint, req domain.ProductReq) (*domain.Product, error) {
	updatedProduct, err := useCase.Repo.Update(id, req)
	if updatedProduct == nil && err == nil {
		return updatedProduct, errors.New("not found")
	}

	return updatedProduct, err
}
