package usecases

import "api/src/product/domain"

type ProductRepository interface {
	Insert(req *domain.ProductReq) (*domain.Product, error)
}

type InsertProductUseCase struct {
	Repo ProductRepository
}

func (useCase *InsertProductUseCase) Execute(req *domain.ProductReq) (*domain.Product, error) {
	return useCase.Repo.Insert(req)
}
