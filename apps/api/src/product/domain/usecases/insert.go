package usecases

import "api/src/product/domain"

type InsertProductUseCase struct {
	Repo ProductRepository
}

func (useCase *InsertProductUseCase) Execute(req domain.ProductReq) (*domain.Product, error) {
	return useCase.Repo.Insert(req)
}
