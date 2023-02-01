package usecases_test

import (
	"api/src/inventory/domain"
	"api/src/inventory/domain/usecases"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var findOneInventoryUseCase = usecases.FindOneInventoryUseCase{Repo: &inventoryRepo}

func TestFindOneInventoryUseCaseWithUnknownError(t *testing.T) {
	assert := assert.New(t)
	id := 1
	inventoryRepo.IsErr = true
	inventory, err := findOneInventoryUseCase.Execute(&id)

	assert.EqualError(err, "unknown error")
	assert.Nil(inventory)
	inventoryRepo.IsErr = false
}

func TestFindOneInventoryUseCaseWithNotFoundError(t *testing.T) {
	assert := assert.New(t)
	id := 1
	inventory, err := findOneInventoryUseCase.Execute(&id)

	assert.EqualError(err, "not found")
	assert.Nil(inventory)
}

func TestFindOneInventoryUseCaseHappyCase(t *testing.T) {
	assert := assert.New(t)
	id := 1
	inventoryRepo.InventoryList = append(inventoryRepo.InventoryList, domain.Inventory{Id: id, Quantity: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	inventory, err := findOneInventoryUseCase.Execute(&id)

	assert.NotNil(inventory)
	assert.Equal(id, inventory.Id)
	assert.Nil(err)
}
