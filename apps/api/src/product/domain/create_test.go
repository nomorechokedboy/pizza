package domain_test

import (
	"api/src/product"
	"api/src/product/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProductUseCaseWithUnknownError(t *testing.T) {
	assert := assert.New(t)
	productMemRepo := product.ProductInMemoryRepo{ProductList: make([]domain.Product, 0), IsErr: false}
	productMemRepo.IsErr = true
	createProductUseCase := domain.CreateProductUseCase{Repo: &productMemRepo}
	req := domain.ProductReq{Description: "Test Description", Name: "Should ok", SKU: "Success", Price: 10000}
	product, err := createProductUseCase.Execute(&req)

	assert.EqualError(err, "unknown error")
	assert.Nil(product)
}

func TestCreateProductUseCaseWithDuplicateError(t *testing.T) {
	assert := assert.New(t)
	productMemRepo := product.ProductInMemoryRepo{ProductList: make([]domain.Product, 0), IsErr: false}
	_, err := productMemRepo.Insert(&domain.ProductReq{Description: "Another description", Name: "Lmao", SKU: "The duplicate SKU", Price: 101010})
	assert.Nil(err)
	createProductUseCase := domain.CreateProductUseCase{Repo: &productMemRepo}

	req := domain.ProductReq{Description: "Test Description", Name: "Should ok", SKU: "The duplicate SKU", Price: 10000}
	product, err := createProductUseCase.Execute(&req)

	assert.EqualError(err, "resource exist")
	assert.Nil(product)
}

func TestCreateProductUseCaseHappyCase(t *testing.T) {
	assert := assert.New(t)
	productMemRepo := product.ProductInMemoryRepo{ProductList: make([]domain.Product, 0), IsErr: false}
	createProductUseCase := domain.CreateProductUseCase{Repo: &productMemRepo}
	req := domain.ProductReq{Description: "Test Description", Name: "Should ok", SKU: "Success", Price: 10000}
	product, err := createProductUseCase.Execute(&req)

	assert.Nil(err)
	assert.NotNil(product)
	assert.Equal(product.Description, req.Description)
	assert.Equal(product.Name, req.Name)
	assert.Equal(product.Price, req.Price)
	assert.Equal(product.SKU, req.SKU)
}
