package usecases

import (
	"api/src/common"
	"api/src/product/domain"
	"errors"
)

type InsertProductUseCase struct {
	Repo      ProductRepository
	Validator common.Validate
}

func (useCase *InsertProductUseCase) Execute(req domain.ProductReq) (*domain.Product, error) {
	err := useCase.Validator.Exec(req)
	if err != nil {
		return nil, errors.New("invalid data")
	}

	return useCase.Repo.Insert(&req)
}
