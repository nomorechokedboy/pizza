package product

import (
	"api/src/product/domain"
	"errors"
	"time"
)

type ProductInMemoryRepo struct {
	ProductList []domain.Product
	IsErr       bool
}

func (repo *ProductInMemoryRepo) Insert(req *domain.ProductReq) (*domain.Product, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	for _, product := range repo.ProductList {
		if product.SKU == req.SKU {
			return nil, errors.New("resource exist")
		}
	}

	Id := len(repo.ProductList) + 1
	newProduct := domain.Product{Id: int32(Id), CreatedAt: time.Now(), UpdatedAt: time.Now(), Slug: "Not implemented", Description: req.Description, Name: req.Name, SKU: req.SKU, Price: req.Price}
	repo.ProductList = append(repo.ProductList, newProduct)

	return &newProduct, nil
}
