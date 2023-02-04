package repository

import (
	"api/src/product/domain"
	"errors"
	"time"
)

type ProductInMemoryRepo struct {
	DataStore []domain.Product
	IsErr     bool
}

func (repo *ProductInMemoryRepo) Insert(req *domain.ProductReq) (*domain.Product, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	for _, product := range repo.DataStore {
		if product.SKU == req.SKU {
			return nil, errors.New("resource exist")
		}
	}

	Id := len(repo.DataStore) + 1
	newProduct := domain.Product{Id: int32(Id), CreatedAt: time.Now(), UpdatedAt: time.Now(), Slug: "Not implemented", Description: req.Description, Name: req.Name, SKU: req.SKU, Price: req.Price}
	repo.DataStore = append(repo.DataStore, newProduct)

	return &newProduct, nil
}
