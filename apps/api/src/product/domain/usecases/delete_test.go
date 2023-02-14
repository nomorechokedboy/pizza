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

type DeleteProductTestSuite struct {
	suite.Suite
	UseCase *usecases.DeleteProductUseCase
	Repo    *MockRepository
}

func (s *DeleteProductTestSuite) SetupTest() {
	s.Repo = &MockRepository{}
	s.UseCase = &usecases.DeleteProductUseCase{Repo: s.Repo}
}

func (s *DeleteProductTestSuite) TestDeleteUnknownError() {
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
			s.Repo.On("Delete", id).Return(nil, c.Error).Once()
			product, err := s.UseCase.Execute(id)

			s.Assertions.Nil(product)
			s.Assertions.EqualError(err, c.Expected)
		})
	}
}

func (s *DeleteProductTestSuite) TestDeleteHappyCase() {
	id := uint(1)
	s.Repo.On("Delete", id).Return(&domain.Product{
		Id:          1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Slug:        "slug",
		Description: "Description",
		Name:        "Name",
		SKU:         "Sku",
		Price:       1000,
		Category:    category.Category{ID: 1, Name: "Test", Description: nil},
		Inventory:   inventory.Inventory{ID: 1, Quantity: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}, nil)
	product, err := s.UseCase.Execute(id)

	s.Assertions.NoError(err)
	s.Assertions.Equal(product.Id, int32(id))
}

func TestDeleteProductUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteProductTestSuite))
}
