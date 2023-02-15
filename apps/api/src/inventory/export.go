package inventory

import (
	"api/src/inventory/domain"
	"api/src/inventory/repository"
)

var InventoryMemRepo = repository.InventoryInMemoryRepo{InventoryList: make([]domain.Inventory, 0), IsErr: false}
