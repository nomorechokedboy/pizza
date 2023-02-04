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
	s.repo = repository.ProductInMemoryRepo{DataStore: []domain.Product{{Id: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), Slug: "slug", Description: "Desc", Name: "Lmao", SKU: "123ABC", Price: 1000, Category: category.Category{}, Inventory: inventory.Inventory{}}}, IsErr: false}
	s.useCase = usecases.UpdateProductUseCase{Repo: &s.repo}
}

func (s *UpdateProductTestSuite) TearDownTest() {
	s.useCase.Repo = &repository.ProductInMemoryRepo{DataStore: make([]domain.Product, 0), IsErr: false}
}

func (s *UpdateProductTestSuite) TestUpdateUnknownError() {
	s.repo.IsErr = true
	id := 1
	product, err := s.useCase.Execute(&id, nil)

	s.Assertions.Nil(product)
	s.Assertions.EqualError(err, "unknown error")
}

func (s *UpdateProductTestSuite) TestUpdateDuplicateError() {
	id := 1
	req := domain.ProductReq{Description: "Lalalala", Name: "Updated name", SKU: "Updated SKU", Price: 100, CategoryId: 1, InventoryId: 2}
	product, err := s.useCase.Execute(&id, &req)

	s.Assertions.Nil(product)
	s.Assertions.EqualError(err, "resource exist")
}

func (s *UpdateProductTestSuite) TestUpdateNotFoundError() {
	id := 2
	req := domain.ProductReq{Description: "Lalalala", Name: "Updated name", SKU: "Updated SKU", Price: 100, CategoryId: 1, InventoryId: 2}
	product, err := s.useCase.Execute(&id, &req)

	s.Assertions.Nil(product)
	s.Assertions.EqualError(err, "not found")
}

func (s *UpdateProductTestSuite) TestUpdateHappyCase() {
	id := 1
	req := domain.ProductReq{Description: "Lalalala", Name: "Updated name", SKU: "Updated SKU", Price: 100, CategoryId: 1, InventoryId: 2}
	product, err := s.useCase.Execute(&id, &req)

	s.Assertions.Nil(err)
	s.Assertions.Equal(product.Id, id)
	s.Assertions.Equal(product.Description, req.Description)
	s.Assertions.Equal(product.Category.Id, req.CategoryId)
	s.Assertions.Equal(product.Inventory.Id, req.InventoryId)
	s.Assertions.Equal(product.Name, req.Name)
	s.Assertions.Equal(product.Price, req.Price)
	s.Assertions.Equal(product.SKU, req.SKU)
	s.Assertions.Equal(product, s.repo.DataStore[0])
}

func TestUpdateProductTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateProductTestSuite))
}
