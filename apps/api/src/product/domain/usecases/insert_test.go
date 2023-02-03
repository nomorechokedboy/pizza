package usecases_test

import (
	category "api/src/category/domain"
	inventory "api/src/inventory/domain"
	"api/src/product/domain"
	"api/src/product/domain/usecases"
	"api/src/product/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type InsertProductTestSuite struct {
	suite.Suite
	UseCase usecases.InsertProductUseCase
	Repo    repository.ProductInMemoryRepo
}

func (s *InsertProductTestSuite) SetupTest() {
	s.Repo = repository.ProductInMemoryRepo{DataStore: make([]domain.Product, 0), IsErr: false}
	s.UseCase = usecases.InsertProductUseCase{Repo: &s.Repo}
}

func (s *InsertProductTestSuite) TearDownTest() {
	s.UseCase.Repo = &repository.ProductInMemoryRepo{DataStore: make([]domain.Product, 0), IsErr: false}
}

func (s *InsertProductTestSuite) TestInsertWithErrors() {
	s.Repo.IsErr = true
	req := domain.ProductReq{Description: "Test Description", Name: "Should ok", SKU: "Success", Price: 10000}
	product, err := s.UseCase.Execute(&req)

	s.Assertions.EqualError(err, "unknown error")
	s.Assertions.Nil(product)
}

func (s *InsertProductTestSuite) TestInsertWithDuplicateError() {
	req := domain.ProductReq{Description: "Test Description", Name: "Should ok", SKU: "The duplicate SKU", Price: 10000}
	s.Repo.DataStore = append(s.Repo.DataStore, domain.Product{Id: 1, Description: "Another description", Name: "Lmao", SKU: "The duplicate SKU", Price: 101010, CreatedAt: time.Now(), UpdatedAt: time.Now(), Slug: "lmao", Category: category.Category{}, Inventory: inventory.Inventory{}})
	product, err := s.UseCase.Execute(&req)

	s.Assertions.EqualError(err, "resource exist")
	s.Assertions.Nil(product)
}

func (s *InsertProductTestSuite) TestInsertHappyCase() {
	req := domain.ProductReq{Description: "Test Description", Name: "Should ok", SKU: "Success", Price: 10000}
	product, err := s.UseCase.Execute(&req)

	s.Assertions.Nil(err)
	s.Assertions.NotNil(product)
	s.Assertions.Equal(product.Description, req.Description)
	s.Assertions.Equal(product.Name, req.Name)
	s.Assertions.Equal(product.Price, req.Price)
	s.Assertions.Equal(product.SKU, req.SKU)
}

func TestInsertProductTestSuite(t *testing.T) {
	suite.Run(t, new(InsertProductTestSuite))
}
