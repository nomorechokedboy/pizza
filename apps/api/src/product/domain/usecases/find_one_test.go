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

type FindOneProductTestSuite struct {
	suite.Suite
	UseCase *usecases.FindOneProductUseCase
	Repo    *repository.ProductInMemoryRepo
}

func (s *FindOneProductTestSuite) SetupTest() {
	s.Repo = &repository.ProductInMemoryRepo{DataStore: []domain.Product{{Id: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), Slug: "slug", Description: "Desc", Name: "Lmao", SKU: "123ABC", Price: 1000, Category: category.Category{}, Inventory: inventory.Inventory{}}}, IsErr: false}
	s.UseCase = &usecases.FindOneProductUseCase{Repo: s.Repo}
}

func (s *FindOneProductTestSuite) TearDownTest() {
	s.Repo = &repository.ProductInMemoryRepo{DataStore: []domain.Product{}, IsErr: false}
}

func (s *FindOneProductTestSuite) TestFindOneUnknownError() {
	s.Repo.IsErr = true
	id := 1
	product, err := s.UseCase.Execute(&id)

	s.Assertions.Nil(product)
	s.Assertions.EqualError(err, "unknown error")
}

func (s *FindOneProductTestSuite) TestFindOneNotFoundError() {
	id := 1
	product, err := s.UseCase.Execute(&id)

	s.Assertions.Nil(product)
	s.Assertions.EqualError(err, "not found")
}

func (s *FindOneProductTestSuite) TestFindOneHappyCase() {
	id := 1
	product, err := s.UseCase.Execute(&id)

	s.Assertions.Nil(err)
	s.Assertions.Equal(id, product.Id)
	s.Assertions.Equal(s.Repo.DataStore[0], product)
}

func TestFindOneProductTestSuite(t *testing.T) {
	suite.Run(t, new(FindOneProductTestSuite))
}
