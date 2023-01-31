package repository

import (
	"api/src/inventory/domain"
	"errors"
	"fmt"
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

	Id := len(repo.InventoryList) + 1
	newProduct := domain.Inventory{Id: Id, Quantity: req.Quantity, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	repo.InventoryList = append(repo.InventoryList, newProduct)

	return &newProduct, nil
}

func (repo *InventoryInMemoryRepo) Update(id *int, req *domain.WriteInventoryBody) (*domain.Inventory, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	for _, inventory := range repo.InventoryList {
		if inventory.Id == *id {
			inventory.Quantity = req.Quantity
			inventory.UpdatedAt = time.Now()
			return &inventory, nil
		}
	}

	return nil, nil
}

func (repo *InventoryInMemoryRepo) Delete(req *int) (*domain.Inventory, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	inventoryListLen := len(repo.InventoryList)

	if *req > inventoryListLen {
		return nil, nil
	}

	var res domain.Inventory = domain.Inventory{Id: 12, Quantity: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	filteredList := make([]domain.Inventory, 0)

	for _, inventory := range repo.InventoryList {
		if inventory.Id == *req {
			res = inventory
			fmt.Println(res)
			continue
		}

		filteredList = append(filteredList, inventory)
	}
	repo.InventoryList = filteredList

	return &res, nil
}

func (repo *InventoryInMemoryRepo) Find() (*[]domain.Inventory, error) {
	return nil, nil
}

func (repo *InventoryInMemoryRepo) FindOne(req *int) (*domain.Inventory, error) {
	return nil, nil
}
