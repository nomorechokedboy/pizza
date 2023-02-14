package repository

import (
	"api/src/inventory/domain"
	"errors"
	"time"
)

type InventoryInMemoryRepo struct {
	InventoryList []domain.Inventory
	IsErr         bool
}

func (repo *InventoryInMemoryRepo) Insert(req *domain.WriteInventoryBody) (*domain.Inventory, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	Id := uint(len(repo.InventoryList)) + 1
	newProduct := domain.Inventory{ID: Id, Quantity: req.Quantity, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	repo.InventoryList = append(repo.InventoryList, newProduct)

	return &newProduct, nil
}

func (repo *InventoryInMemoryRepo) Update(id *int, req *domain.WriteInventoryBody) (*domain.Inventory, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	for i := range repo.InventoryList {
		inventory := &repo.InventoryList[i]
		if inventory.ID == uint(*id) {
			inventory.Quantity = req.Quantity
			inventory.UpdatedAt = time.Now()
			return inventory, nil
		}
	}

	return nil, nil
}

func (repo *InventoryInMemoryRepo) Delete(req *int) (*domain.Inventory, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	var res domain.Inventory
	pos := -1
	filteredList := make([]domain.Inventory, 0)

	for i, inventory := range repo.InventoryList {
		if inventory.ID == uint(*req) {
			res = inventory
			pos = i
			continue
		}

		filteredList = append(filteredList, inventory)
	}
	repo.InventoryList = filteredList

	if pos < 0 {
		return nil, nil
	}

	return &res, nil
}

func (repo *InventoryInMemoryRepo) FindOne(req *int) (*domain.Inventory, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	inventoryListLen := len(repo.InventoryList)

	if *req > inventoryListLen {
		return nil, nil
	}

	var res *domain.Inventory

	for _, inventory := range repo.InventoryList {
		if inventory.ID == uint(*req) {
			res = &inventory
			break
		}
	}

	return res, nil
}
