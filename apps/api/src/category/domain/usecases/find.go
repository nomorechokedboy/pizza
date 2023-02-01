package usecases

import "api/src/category/domain"

type FindCategoryUseCase struct {
	Repo CategoryRepository
}

func (useCase *FindCategoryUseCase) Execute(req *domain.CategoryQuery) (*[]domain.Category, error) {
	return useCase.Repo.Find(req)
}
