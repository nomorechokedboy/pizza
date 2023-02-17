package usecases

import (
	"api/src/category/domain"
	"api/src/common"
)

type FindCategoryUseCase struct {
	Repo CategoryRepository
}

func (useCase *FindCategoryUseCase) Execute(req *domain.CategoryQuery) (common.BasePaginationResponse[domain.Category], error) {
	return useCase.Repo.Find(req)
}
