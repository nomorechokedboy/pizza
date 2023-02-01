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
