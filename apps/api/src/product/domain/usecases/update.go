package usecases

import (
	"api/src/common"
	"api/src/product/domain"
)

type UpdateProductUseCase struct {
	Repo      ProductRepository
	Validator *common.Validate
}

func (useCase *UpdateProductUseCase) Execute(id *int, req *domain.ProductReq) (*domain.Product, error) {
	return nil, nil
}
