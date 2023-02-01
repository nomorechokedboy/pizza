package usecases

import "api/src/inventory/domain"

type FindOneInventoryUseCase struct {
	Repo InventoryRepository
}

func (useCase *FindOneInventoryUseCase) Execute(id *int) (*domain.Inventory, error) {
	return nil, nil
}
