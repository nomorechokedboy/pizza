package usecases

import "api/src/inventory/domain"

type InventoryRepository interface {
	Insert(req *domain.WriteInventoryBody) (*domain.Inventory, error)
	Update(id *int, req *domain.WriteInventoryBody) (*domain.Inventory, error)
	Delete(req *string) (*domain.Inventory, error)
	Find() (*[]domain.Inventory, error)
	FindOne(req *string) (*domain.Inventory, error)
}
