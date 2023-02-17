package usecases

import (
	"api/src/common"
	"api/src/product/domain"
)

type FindProductUseCase struct {
	Repo ProductRepository
}

func (useCase *FindProductUseCase) Execute(queries *domain.ProductQuery) (common.BasePaginationResponse[domain.Product], error) {
	return useCase.Repo.Find(queries)
}
