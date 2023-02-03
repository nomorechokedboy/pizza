package usecases

import "api/src/inventory/domain"

type InsertInventoryUseCase struct {
	Repo InventoryRepository
}

func (useCase *InsertInventoryUseCase) Execute(req *domain.WriteInventoryBody) (*domain.Inventory, error) {
	return useCase.Repo.Insert(req)
}
