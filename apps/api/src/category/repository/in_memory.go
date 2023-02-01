package repository

import (
	"api/src/category/domain"
	"errors"
)

type CategoryInMemoryRepo struct {
	Data  []domain.Category
	IsErr bool
}

func (repo *CategoryInMemoryRepo) Insert(req *domain.WriteCategoryBody) (*domain.Category, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	for _, product := range repo.Data {
		if product.Name == req.Name {
			return nil, errors.New("resource exist")
		}
	}

	Id := len(repo.Data) + 1
	newCategory := domain.Category{Id: Id, Description: req.Description, Name: req.Name}
	repo.Data = append(repo.Data, newCategory)

	return &newCategory, nil
}

func (repo *CategoryInMemoryRepo) Update(id *int, req *domain.WriteCategoryBody) (*domain.Category, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	for _, category := range repo.Data {
		if category.Id == *id {
			category.Description = req.Description
			category.Name = req.Name
			return &category, nil
		}
	}

	return nil, nil
}

func (repo *CategoryInMemoryRepo) Delete(req *int) (*domain.Category, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	inventoryListLen := len(repo.Data)

	if *req > inventoryListLen {
		return nil, nil
	}

	var res domain.Category
	filteredList := make([]domain.Category, 0)

	for _, inventory := range repo.Data {
		if inventory.Id == *req {
			res = inventory
			continue
		}

		filteredList = append(filteredList, inventory)
	}
	repo.Data = filteredList

	return &res, nil
}

func (repo *CategoryInMemoryRepo) FindOne(id *int) (*domain.Category, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	var res *domain.Category

	for _, inventory := range repo.Data {
		if inventory.Id == *id {
			res = &inventory
			break
		}
	}

	return res, nil
}
