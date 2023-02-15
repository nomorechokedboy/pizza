package usecases_test

import (
	"api/src/inventory/domain"
	"api/src/inventory/domain/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

var updateInventoryUseCase = usecases.UpdateInventoryUseCase{Repo: &inventoryRepo}

func TestUpdateInventoryUseCaseWithNotFoundError(t *testing.T) {
	assert := assert.New(t)
	req := domain.WriteInventoryBody{Quantity: 3}
	id := 10
	updatedInventory, err := updateInventoryUseCase.Execute(&id, &req)

	assert.EqualError(err, "not found")
	assert.Nil(updatedInventory)
}

func TestUpdateInventoryUseCaseWhenUnexpectedError(t *testing.T) {
	assert := assert.New(t)
	inventoryRepo.IsErr = true
	req := domain.WriteInventoryBody{Quantity: 3}
	id := 1
	updatedInventory, err := updateInventoryUseCase.Execute(&id, &req)

	assert.EqualError(err, "unknown error")
	assert.Nil(updatedInventory)
	inventoryRepo.IsErr = false
}

func TestUpdateInventoryUseCaseHappyCase(t *testing.T) {
	assert := assert.New(t)
	req := domain.WriteInventoryBody{Quantity: 5}
	id := 1
	updatedInventory, err := updateInventoryUseCase.Execute(&id, &req)

	assert.Nil(err)
	assert.Equal(updatedInventory.Quantity, req.Quantity)
}
