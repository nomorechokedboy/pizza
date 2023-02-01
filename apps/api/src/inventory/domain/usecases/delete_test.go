package usecases_test

import (
	"api/src/inventory/domain"
	"api/src/inventory/domain/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

var deleteInventoryUseCase = usecases.DeleteInventoryUseCase{Repo: &inventoryRepo}

func TestDeleteInventoryUseCaseWithUnknownError(t *testing.T) {
	assert := assert.New(t)
	id := 10
	inventoryRepo.IsErr = true
	deletedInventory, err := deleteInventoryUseCase.Execute(&id)

	assert.EqualError(err, "unknown error")
	assert.Nil(deletedInventory)
	inventoryRepo.IsErr = false
}

func TestDeleteInventoryUseCaseWithNotFoundError(t *testing.T) {
	assert := assert.New(t)
	id := 10
	deletedInventory, err := deleteInventoryUseCase.Execute(&id)

	assert.EqualError(err, "not found")
	assert.Nil(deletedInventory)
}

func TestDeleteInventoryUseCaseHappyCase(t *testing.T) {
	assert := assert.New(t)
	id := 1
	_, err := insertInventoryUseCase.Execute(&domain.WriteInventoryBody{Quantity: 5})
	assert.Nil(err)
	deletedInventory, err := deleteInventoryUseCase.Execute(&id)

	assert.NotNil(deletedInventory)
	assert.Nil(err)
	assert.Equal(id, deletedInventory.Id)
	assert.Equal(5, deletedInventory.Quantity)
}
