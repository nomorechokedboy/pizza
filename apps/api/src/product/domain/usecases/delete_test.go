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

type DeleteProductTestSuite struct {
	suite.Suite
	UseCase *usecases.DeleteProductUseCase
	Repo    *repository.ProductInMemoryRepo
}

func (s *DeleteProductTestSuite) SetupTest() {
	s.Repo = &repository.ProductInMemoryRepo{DataStore: []domain.Product{{Id: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), Slug: "slug", Description: "Desc", Name: "Lmao", SKU: "123ABC", Price: 1000, Category: category.Category{}, Inventory: inventory.Inventory{}}}, IsErr: false}
	s.UseCase = &usecases.DeleteProductUseCase{Repo: s.Repo}
}

func (s *DeleteProductTestSuite) TearDownTest() {
	s.Repo = &repository.ProductInMemoryRepo{DataStore: []domain.Product{}, IsErr: false}
}

func (s *DeleteProductTestSuite) TestDeleteUnknownError() {
	s.Repo.IsErr = true
	id := 1
	product, err := s.UseCase.Execute(&id)

	s.Assertions.Nil(product)
	s.Assertions.EqualError(err, "unknown error")
}

func (s *DeleteProductTestSuite) TestDeleteNotFoundError() {
	id := 1
	product, err := s.UseCase.Execute(&id)

	s.Assertions.Nil(product)
	s.Assertions.EqualError(err, "not found")
}

func (s *DeleteProductTestSuite) TestDeleteHappyCase() {
	id := 1
	product, err := s.UseCase.Execute(&id)

	s.Assertions.Nil(err)
	s.Assertions.Equal(product.Id, id)
	s.Assertions.Equal(0, len(s.Repo.DataStore))
}

func TestDeleteProductTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteProductTestSuite))
}
