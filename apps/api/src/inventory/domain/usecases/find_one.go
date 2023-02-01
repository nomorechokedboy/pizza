package usecases

import (
	"api/src/inventory/domain"
	"errors"
)

type FindOneInventoryUseCase struct {
	Repo InventoryRepository
}

func (useCase *FindOneInventoryUseCase) Execute(id *int) (*domain.Inventory, error) {
	inventory, err := useCase.Repo.FindOne(id)
	if inventory == nil && err == nil {
		return inventory, errors.New("not found")
	}

	return inventory, err
}
