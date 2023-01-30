package domain_test

import (
	"api/src/product"
	"api/src/product/domain"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProductUseCaseHappyCase(t *testing.T) {
	assert := assert.New(t)
	productMemRepo := product.ProductInMemoryRepo{ProductList: make([]domain.Product, 0), IsErr: false}
	createProductUseCase := domain.CreateProductUseCase{Repo: &productMemRepo}
	req := domain.ProductReq{Description: "Test Description", Name: "Should ok", SKU: "Success", Price: 10000}
	product, err := createProductUseCase.Execute(&req)

	assert.Nil(err)
	t.Logf("Type of product is %v\n", reflect.TypeOf(product))
}

func TestCreateProductUseCaseWithUnknownError(t *testing.T) {
	assert := assert.New(t)
	productMemRepo := product.ProductInMemoryRepo{ProductList: make([]domain.Product, 0), IsErr: false}
	productMemRepo.IsErr = true
	createProductUseCase := domain.CreateProductUseCase{Repo: &productMemRepo}
	req := domain.ProductReq{Description: "Test Description", Name: "Should ok", SKU: "Success", Price: 10000}
	product, err := createProductUseCase.Execute(&req)

	assert.EqualError(err, "connection error")
	assert.Nil(product)
}

func TestCreateProductUseCaseWithDuplicateError(t *testing.T) {
	assert := assert.New(t)
	productMemRepo := product.ProductInMemoryRepo{ProductList: make([]domain.Product, 0), IsErr: false}
	productMemRepo.Insert(&domain.ProductReq{Description: "Another description", Name: "Lmao", SKU: "The duplicate SKU", Price: 101010})
	createProductUseCase := domain.CreateProductUseCase{Repo: &productMemRepo}

	req := domain.ProductReq{Description: "Test Description", Name: "Should ok", SKU: "The duplicate SKU", Price: 10000}
	product, err := createProductUseCase.Execute(&req)

	assert.EqualError(err, "resource exist")
	assert.Nil(product)
}
