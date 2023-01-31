package usecases

import (
	"api/src/inventory/domain"
	"errors"
)

type UpdateInventoryUseCase struct {
	Repo InventoryRepository
}

func (useCase *UpdateInventoryUseCase) Execute(id *int, req *domain.WriteInventoryBody) (*domain.Inventory, error) {
	inventory, err := useCase.Repo.Update(id, req)
	if err != nil {
		return nil, err
	}

	if inventory == nil {
		return inventory, errors.New("not found")
	}

	return inventory, err
}
