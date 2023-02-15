package usecases

import (
	"api/src/product/domain"
	"errors"
)

type FindOneProductUseCase struct {
	Repo ProductRepository
}

func (useCase *FindOneProductUseCase) Execute(id uint) (*domain.Product, error) {
	product, err := useCase.Repo.FindOne(id)
	if product == nil && err == nil {
		return product, errors.New("not found")
	}

	return product, err
}
