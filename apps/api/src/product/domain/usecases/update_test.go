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

type UpdateProductTestSuite struct {
	suite.Suite
	UseCase usecases.UpdateProductUseCase
	Repo    MockRepository
}

type UpdateCategoryTestCase struct {
	Input    domain.Product
	Expected []domain.Product
	TestName string
	InitData []domain.Product
}

var updateReq = domain.ProductReq{Description: "Lalalala", Name: "Updated name", SKU: "Updated SKU", Price: 100, CategoryId: 1, InventoryId: 2}

func (s *UpdateProductTestSuite) SetupTest() {
	s.Repo = MockRepository{}
	s.UseCase = usecases.UpdateProductUseCase{Repo: &s.Repo}
}

func (s *UpdateProductTestSuite) TestUpdateUnknownError() {
	id := uint(1)
	s.Repo.On("Update", id, updateReq).Return(nil, errors.New("unknown error"))
	product, err := s.UseCase.Execute(id, updateReq)

	s.Assertions.Nil(product)
	s.Assertions.EqualError(err, "unknown error")
}

func (s *UpdateProductTestSuite) TestUpdateDuplicateError() {
	id := uint(1)
	s.Repo.On("Update", id, updateReq).Return(nil, errors.New("resource exist"))
	product, err := s.UseCase.Execute(id, updateReq)

	s.Assertions.Nil(product)
	s.Assertions.EqualError(err, "resource exist")
}

func (s *UpdateProductTestSuite) TestUpdateNotFoundError() {
	id := uint(2)
	s.Repo.On("Update", id, updateReq).Return(nil, nil)
	product, err := s.UseCase.Execute(id, updateReq)

	s.Assertions.Nil(product)
	s.Assertions.EqualError(err, "not found")
}

func (s *UpdateProductTestSuite) TestUpdateHappyCase() {
	id := uint(1)
	s.Repo.On("Update", id, updateReq).Return(&domain.Product{
		Id:          1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Slug:        "slug",
		Description: updateReq.Description,
		Name:        updateReq.Name,
		SKU:         updateReq.SKU,
		Price:       updateReq.Price,
		Category:    category.Category{ID: uint(updateReq.CategoryId), Name: "Test", Description: nil},
		Inventory:   inventory.Inventory{ID: uint(updateReq.InventoryId), Quantity: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}, nil)
	product, err := s.UseCase.Execute(id, updateReq)

	s.Assertions.NoError(err)
	s.Assertions.Equal(id, uint(product.Id))
	s.Assertions.Equal(product.Description, updateReq.Description)
	s.Assertions.Equal(uint(updateReq.CategoryId), product.Category.ID)
	s.Assertions.Equal(updateReq.InventoryId, product.Inventory.ID)
	s.Assertions.Equal(updateReq.Name, product.Name)
	s.Assertions.Equal(updateReq.Price, product.Price)
	s.Assertions.Equal(updateReq.SKU, product.SKU)
}

func TestUpdateProductUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateProductTestSuite))
}
