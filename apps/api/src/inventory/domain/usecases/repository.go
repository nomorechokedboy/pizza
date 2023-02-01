package usecases

import "api/src/inventory/domain"

type InventoryRepository interface {
	Insert(req *domain.WriteInventoryBody) (*domain.Inventory, error)
	Update(id *int, req *domain.WriteInventoryBody) (*domain.Inventory, error)
	Delete(req *int) (*domain.Inventory, error)
	FindOne(req *int) (*domain.Inventory, error)
}
