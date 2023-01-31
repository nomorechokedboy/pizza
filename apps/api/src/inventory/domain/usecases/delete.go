package usecases

import "api/src/inventory/domain"

type DeleteInventoryUseCase struct {
	Repo InventoryRepository
}

func (useCase *DeleteInventoryUseCase) Execute(req *domain.WriteInventoryBody) (*domain.Inventory, error) {
	return nil, nil
}
