package usecases_test

import (
	category "api/src/category/domain"
	inventory "api/src/inventory/domain"
	"api/src/product/domain"
	"api/src/product/domain/usecases"
	"api/src/utils"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type FindOneProductTestSuite struct {
	suite.Suite
	UseCase *usecases.FindOneProductUseCase
	Repo    *MockProductRepository
}

func (s *FindOneProductTestSuite) SetupTest() {
	s.Repo = &MockProductRepository{}
	s.UseCase = &usecases.FindOneProductUseCase{Repo: s.Repo}
}

func (s *FindOneProductTestSuite) TestFindOneUnknownError() {
	id := uint(1)
	table := []struct {
		Description string
		Error       error
		Expected    string
	}{
		{
			Description: "Unknown error",
			Error:       errors.New("unknown error"),
			Expected:    "unknown error",
		},
		{
			Description: "Not found error",
			Error:       nil,
			Expected:    "not found",
		},
	}

	for _, c := range table {
		s.Run(c.Description, func() {
			s.Repo.On("FindOne", id).Return(nil, c.Error).Once()
			product, err := s.UseCase.Execute(id)

			s.Assertions.Nil(product)
			s.Assertions.EqualError(err, c.Expected)
		})
	}
}

func (s *FindOneProductTestSuite) TestFindOneHappyCase() {
	id := uint(1)
	s.Repo.On("FindOne", id).Return(&domain.Product{
		ID:          1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Slug:        "slug",
		Description: utils.GetDataTypeAddress("Description"),
		Name:        "Name",
		SKU:         "Sku",
		Price:       1000,
		Category:    category.Category{ID: 1, Name: "Test", Description: nil},
		Inventory:   inventory.Inventory{ID: 1, Quantity: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}, nil)
	product, err := s.UseCase.Execute(id)

	s.Assertions.NoError(err)
	s.Assertions.Equal(id, uint(product.ID))
}

func TestFindOneProductUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(FindOneProductTestSuite))
}
