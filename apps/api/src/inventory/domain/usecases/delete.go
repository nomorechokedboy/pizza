package usecases

import (
	"api/src/inventory/domain"
	"errors"
)

type DeleteInventoryUseCase struct {
	Repo InventoryRepository
}

func (useCase *DeleteInventoryUseCase) Execute(req *int) (*domain.Inventory, error) {
	deletedInventory, err := useCase.Repo.Delete(req)
	if deletedInventory == nil && err == nil {
		return deletedInventory, errors.New("not found")
	}

	return deletedInventory, err
}
