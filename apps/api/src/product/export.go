package product

import (
	"api/src/product/domain"
	"api/src/product/domain/usecases"
	"api/src/product/repository"
)

var ProductMemRepo = repository.ProductInMemoryRepo{DataStore: make([]domain.Product, 0), IsErr: false}

var CreateProductUseCase = usecases.InsertProductUseCase{Repo: &ProductMemRepo}
