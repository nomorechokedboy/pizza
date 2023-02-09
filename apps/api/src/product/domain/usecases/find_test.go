package usecases_test

import (
	category "api/src/category/domain"
	"api/src/common"
	inventory "api/src/inventory/domain"
	"api/src/product/domain"
	"api/src/product/domain/usecases"
	"api/src/product/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type FindProductTestSuite struct {
	suite.Suite
	UseCase *usecases.FindProductUseCase
	Repo    *repository.ProductInMemoryRepo
}

var initData = []domain.Product{
	{
		Id:          1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Slug:        "slug",
		Description: "Desc",
		Name:        "Lmao",
		SKU:         "123ABC",
		Price:       1000,
		Category:    category.Category{ID: 1},
		Inventory:   inventory.Inventory{},
	},
	{
		Id:          2,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Slug:        "pizza-morella",
		Description: "morella",
		Name:        "pizza morella",
		SKU:         "p123",
		Price:       1000,
		Category:    category.Category{ID: 1},
		Inventory:   inventory.Inventory{},
	},
	{
		Id:          3,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Slug:        "pizza-free",
		Description: "free",
		Name:        "pizza free",
		SKU:         "1234",
		Price:       1000,
		Category:    category.Category{ID: 2},
		Inventory:   inventory.Inventory{},
	},
}

func (s *FindProductTestSuite) SetupTest() {
	s.Repo = &repository.ProductInMemoryRepo{DataStore: initData, IsErr: false}
	s.UseCase = &usecases.FindProductUseCase{Repo: s.Repo}
}

func (s *FindProductTestSuite) TearDownTest() {
	s.Repo = &repository.ProductInMemoryRepo{DataStore: []domain.Product{}, IsErr: false}
}

func (s *FindProductTestSuite) TestFindUnknownError() {
	s.Repo.IsErr = true
	product, err := s.UseCase.Execute(&domain.ProductQuery{})

	s.Assertions.Nil(product)
	s.Assertions.EqualError(err, "unknown error")
}

func (s *FindProductTestSuite) TestFindUseCases() {
	ExceedPage := 3
	ExceedPageSize := 3
	Page := 0
	PageSize := 1
	Q := "lmao"
	testCases := []struct {
		Queries     *domain.ProductQuery
		Expected    *[]domain.Product
		Description string
	}{
		{
			Queries: &domain.ProductQuery{
				InventoryId: nil,
				Base: common.BaseQuery{
					Page:     &ExceedPage,
					PageSize: &ExceedPageSize,
				},
			},
			Expected:    &[]domain.Product{},
			Description: "Exceed Page Number",
		},
		{
			Queries: &domain.ProductQuery{
				InventoryId: nil,
				Base: common.BaseQuery{
					Page:     &Page,
					PageSize: &PageSize,
				},
			},
			Expected: &[]domain.Product{
				{
					Id:          1,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					Slug:        "slug",
					Description: "Desc",
					Name:        "Lmao",
					SKU:         "123ABC",
					Price:       1000,
					Category: category.Category{
						ID: 1,
					},
					Inventory: inventory.Inventory{},
				},
			},
			Description: "Pagination",
		},
		{
			Queries: &domain.ProductQuery{},
			Expected: &[]domain.Product{
				{
					Id: 2, CreatedAt: time.Now(),
					UpdatedAt:   time.Now(),
					Slug:        "pizza-morella",
					Description: "morella",
					Name:        "pizza morella",
					SKU:         "p123",
					Price:       1000,
					Category: category.Category{
						ID: 1,
					},
					Inventory: inventory.Inventory{},
				},
			},
			Description: "Pagination",
		},
		{
			Queries: &domain.ProductQuery{
				Base: common.BaseQuery{
					Page:     nil,
					PageSize: nil,
					Q:        &Q,
				},
				InventoryId: nil,
			},
			Expected:    &[]domain.Product{},
			Description: "Search Not Found",
		},
		{Queries: &domain.ProductQuery{
			Base: common.BaseQuery{
				Page:     nil,
				PageSize: nil,
				Q:        &Q,
			},
		},
			Expected:    nil,
			Description: "Search Happy Case",
		},
		{
			Queries: &domain.ProductQuery{
				Base: common.BaseQuery{
					Page:     &[]int{0}[0],
					PageSize: &[]int{5}[0],
					Q:        &Q,
				},
			},
			Expected: &[]domain.Product{
				initData[1],
			},
			Description: "Search With All Params",
		},
		{
			Queries: &domain.ProductQuery{
				InventoryId: &[]int{10}[0],
			},
			Expected:    &[]domain.Product{},
			Description: "Query By Inventory Id Not Found",
		},
		{
			Queries: &domain.ProductQuery{
				InventoryId: &[]int{1}[0],
			},
			Expected:    &[]domain.Product{initData[0], initData[1]},
			Description: "Query By Inventory Id Happy Case",
		},
		{
			Queries: &domain.ProductQuery{
				Base: common.BaseQuery{
					Page:     &[]int{1}[0],
					PageSize: &[]int{2}[0],
					Q:        &[]string{"free"}[0],
				},
				InventoryId: &[]int{2}[0],
			},
			Expected:    &[]domain.Product{initData[2]},
			Description: "Query By Inventory Id All Params",
		},
	}

	for _, c := range testCases {
		s.T().Run(c.Description, func(t *testing.T) {
			products, err := s.UseCase.Execute(c.Queries)

			s.Assertions.Nil(err)
			s.Assertions.Equal(c.Expected, *products)
		})
	}
}

func TestFindProductTestSuite(t *testing.T) {
	suite.Run(t, new(FindProductTestSuite))
}
