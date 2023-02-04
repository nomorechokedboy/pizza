package usecases

import "api/src/product/domain"

type FindProductUseCase struct {
	Repo ProductRepository
}

func (useCase *FindProductUseCase) Execute(queries *domain.ProductQuery) (*domain.Product, error) {
	return nil, nil
}
