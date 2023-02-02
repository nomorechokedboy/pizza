package usecases

import "api/src/product/domain"

type ProductRepository interface {
	Insert(req *domain.ProductReq) (*domain.Product, error)
}

type CreateProductUseCase struct {
	Repo ProductRepository
}

func (useCase *CreateProductUseCase) Execute(req *domain.ProductReq) (*domain.Product, error) {
	return useCase.Repo.Insert(req)
}
