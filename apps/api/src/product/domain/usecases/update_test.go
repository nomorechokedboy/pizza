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

type UpdateProductTestSuite struct {
	suite.Suite
	useCase usecases.UpdateProductUseCase
	repo    repository.ProductInMemoryRepo
}

type UpdateCategoryTestCase struct {
	Input    domain.Product
	Expected []domain.Product
	TestName string
	InitData []domain.Product
}

func (s *UpdateProductTestSuite) SetupTest() {
	s.repo = repository.ProductInMemoryRepo{DataStore: make([]domain.Product, 0), IsErr: false}
	s.useCase = usecases.UpdateProductUseCase{Repo: &s.repo}
}

func (s *UpdateProductTestSuite) TearDownTest() {
	s.useCase.Repo = &repository.ProductInMemoryRepo{DataStore: make([]domain.Product, 0), IsErr: false}
}

func (s *UpdateProductTestSuite) TestUpdateUnknownError() {
	s.repo.IsErr = true
	category, err := s.useCase.Execute(nil)

	s.Assertions.Nil(category)
	s.Assertions.EqualError(err, "unknown error")
}

func (s *UpdateProductTestSuite) TestUpdateDuplicateError() {
	s.repo.DataStore = append(s.repo.DataStore, domain.Product{Id: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), Slug: "slug", Description: "Desc", Name: "Lmao", SKU: "123ABC", Price: 1000, Category: category.Category{}, Inventory: inventory.Inventory{}})
	req := domain.ProductReq{Description: "Lalalala", Name: "Updated name", SKU: "Updated SKU", Price: 100, CategoryId: 1, InventoryId: 2}
	product, err := s.useCase.Execute(&req)

	s.Assertions.Nil(product)
	s.Assertions.EqualError(err, "resource exist")
}

func (s *UpdateProductTestSuite) TestUpdateNotFoundError() {
	req := domain.ProductReq{Description: "Lalalala", Name: "Updated name", SKU: "Updated SKU", Price: 100, CategoryId: 1, InventoryId: 2}
	product, err := s.useCase.Execute(&req)

	s.Assertions.Nil(product)
	s.Assertions.EqualError(err, "not found")
}

func (s *UpdateProductTestSuite) TestUpdateHappyCase() {
	s.repo.DataStore = append(s.repo.DataStore, domain.Product{Id: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), Slug: "slug", Description: "Desc", Name: "Lmao", SKU: "123ABC", Price: 1000, Category: category.Category{}, Inventory: inventory.Inventory{}})

	id := 1
	req := domain.ProductReq{Description: "Lalalala", Name: "Updated name", SKU: "Updated SKU", Price: 100, CategoryId: 1, InventoryId: 2}
	product, err := s.useCase.Execute(&req)

	s.Assertions.Nil(err)
	s.Assertions.Equal(product.Id, id)
	s.Assertions.Equal(product.Description, req.Description)
	s.Assertions.Equal(product.Category.Id, req.CategoryId)
	s.Assertions.Equal(product.Inventory.Id, req.InventoryId)
	s.Assertions.Equal(product.Name, req.Name)
	s.Assertions.Equal(product.Price, req.Price)
	s.Assertions.Equal(product.SKU, req.SKU)
}

func TestUpdateProductTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateProductTestSuite))
}
