package usecases

import "api/src/product/domain"

type FindOneProductUseCase struct {
	Repo ProductRepository
}

func (useCase *FindOneProductUseCase) Execute(id *int) (*domain.Product, error) {
	return nil, nil
}
