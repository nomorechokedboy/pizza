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

type InsertProductTestSuite struct {
	suite.Suite
	UseCase  usecases.InsertProductUseCase
	mockRepo MockRepository
}

func (s *InsertProductTestSuite) SetupTest() {
	s.mockRepo = MockRepository{}
	s.UseCase = usecases.InsertProductUseCase{Repo: &s.mockRepo}
}

var insertReq = domain.ProductReq{Description: utils.GetDataTypeAddress("Test Description"), Name: "Should ok", SKU: "Success", Price: 10000}

func (s *InsertProductTestSuite) TestInsertWithErrors() {
	table := []struct {
		Description string
		Error       string
	}{
		{
			Description: "Unknown error",
			Error:       "unknown error",
		},
		{
			Description: "Duplicate error",
			Error:       "resource exist",
		},
	}

	for _, c := range table {
		s.Run(c.Description, func() {
			s.mockRepo.On("Insert", insertReq).Return(nil, errors.New(c.Error)).Once()
			product, err := s.UseCase.Execute(insertReq)

			s.Assertions.EqualError(err, c.Error)
			s.Assertions.Nil(product)

		})
	}
}

func (s *InsertProductTestSuite) TestInsertWithDuplicateError() {
	s.mockRepo.On("Insert", insertReq).Return(nil, errors.New("resource exist"))
	product, err := s.UseCase.Execute(insertReq)

	s.Assertions.EqualError(err, "resource exist")
	s.Assertions.Nil(product)
}

func (s *InsertProductTestSuite) TestInsertHappyCase() {
	s.mockRepo.On("Insert", insertReq).Return(&domain.Product{
		ID:          1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Slug:        "slug",
		Description: insertReq.Description,
		Name:        insertReq.Name,
		SKU:         insertReq.SKU,
		Price:       insertReq.Price,
		Category:    category.Category{ID: 1, Name: "Test", Description: nil},
		Inventory:   inventory.Inventory{ID: 1, Quantity: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}, nil)
	product, err := s.UseCase.Execute(insertReq)

	s.Assertions.NoError(err)
	s.Assertions.NotNil(product)
	s.Assertions.Equal(product.Description, insertReq.Description)
	s.Assertions.Equal(product.Name, insertReq.Name)
	s.Assertions.Equal(product.Price, insertReq.Price)
	s.Assertions.Equal(product.SKU, insertReq.SKU)
}

func TestInsertProductUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(InsertProductTestSuite))
}
