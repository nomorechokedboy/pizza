package usecases_test

import (
	category "api/src/category/domain"
	inventory "api/src/inventory/domain"
	"api/src/product/domain"
	"api/src/product/domain/usecases"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type FindOneProductTestSuite struct {
	suite.Suite
	UseCase *usecases.FindOneProductUseCase
	Repo    *MockRepository
}

func (s *FindOneProductTestSuite) SetupTest() {
	s.Repo = &MockRepository{}
	s.UseCase = &usecases.FindOneProductUseCase{Repo: s.Repo}
}

func (s *FindOneProductTestSuite) TestFindOneUnknownError() {
	id := uint(1)
	table := []struct {
		Description string
		Error       string
	}{
		{
			Description: "Unknown error",
			Error:       "unknown error",
		},
		{
			Description: "Not found error",
			Error:       "not found",
		},
	}

	for _, c := range table {
		s.Run(c.Description, func() {
			s.Repo.On("FindOne", id).Return(nil, errors.New(c.Error))
			product, err := s.UseCase.Execute(id)

			s.Assertions.Nil(product)
			s.Assertions.EqualError(err, "unknown error")
		})
	}
}

func (s *FindOneProductTestSuite) TestFindOneHappyCase() {
	id := uint(1)
	s.Repo.On("FindOne", id).Return(&domain.Product{
		Id:          1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Slug:        "slug",
		Description: "Description",
		Name:        "Name",
		SKU:         "Sku",
		Price:       1000,
		Category:    category.Category{ID: 1, Name: "Test", Description: nil},
		Inventory:   inventory.Inventory{Id: 1, Quantity: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}, nil)
	product, err := s.UseCase.Execute(id)

	s.Assertions.NoError(err)
	s.Assertions.Equal(id, uint(product.Id))
}

func TestFindOneProductUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(FindOneProductTestSuite))
}
