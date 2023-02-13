package usecases_test

import (
	"testing"

	"api/src/inventory/domain"
	"api/src/inventory/domain/usecases"
	"api/src/inventory/repository"

	"github.com/stretchr/testify/assert"
)

var inventoryRepo = repository.InventoryInMemoryRepo{InventoryList: make([]domain.Inventory, 0), IsErr: false}
var insertInventoryUseCase = usecases.InsertInventoryUseCase{Repo: &inventoryRepo}

func TestInsertInventoryUseCaseWithUnknownError(t *testing.T) {
	assert := assert.New(t)
	inventoryRepo.IsErr = true
	req := domain.WriteInventoryBody{Quantity: 4}
	inventory, err := insertInventoryUseCase.Execute(&req)

	assert.EqualError(err, "unknown error")
	assert.Nil(inventory)
	inventoryRepo.IsErr = false
}

func TestInsertInventoryUseCaseHappyCase(t *testing.T) {
	assert := assert.New(t)
	req := domain.WriteInventoryBody{Quantity: 5}
	inventory, err := insertInventoryUseCase.Execute(&req)

	assert.Nil(err)
	assert.NotNil(inventory)
	assert.Equal(uint(5), inventory.Quantity)
}
