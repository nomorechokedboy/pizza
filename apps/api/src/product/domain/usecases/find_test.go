package usecases_test

import (
	category "api/src/category/domain"
	"api/src/common"
	inventory "api/src/inventory/domain"
	"api/src/product/domain"
	"api/src/product/domain/usecases"
	"api/src/utils"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type FindProductTestSuite struct {
	suite.Suite
	UseCase usecases.FindProductUseCase
	Repo    MockRepository
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
	s.Repo = MockRepository{}
	s.UseCase = usecases.FindProductUseCase{Repo: &s.Repo}
}

func (s *FindProductTestSuite) TestFindUnknownError() {
	queries := &domain.ProductQuery{}
	s.Repo.On("Find", queries).Return(nil, errors.New("unknown error"))
	product, err := s.UseCase.Execute(queries)

	s.Assertions.Nil(product)
	s.Assertions.EqualError(err, "unknown error")
}

func (s *FindProductTestSuite) TestFindHappyCase() {
	ExceedPage := uint(3)
	ExceedPageSize := uint(3)
	Page := uint(0)
	PageSize := uint(1)
	Q := "lmao"
	testCases := []struct {
		Queries     *domain.ProductQuery
		Expected    []*domain.Product
		Description string
	}{
		{
			Queries: &domain.ProductQuery{
				InventoryId: nil,
				BaseQuery: common.BaseQuery{
					Page:     &ExceedPage,
					PageSize: &ExceedPageSize,
				},
			},
			Expected:    []*domain.Product{},
			Description: "Exceed Page Number",
		},
		{
			Queries: &domain.ProductQuery{
				InventoryId: nil,
				BaseQuery: common.BaseQuery{
					Page:     &Page,
					PageSize: &PageSize,
				},
			},
			Expected: []*domain.Product{
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
			Expected: []*domain.Product{
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
				BaseQuery: common.BaseQuery{
					Page:     nil,
					PageSize: nil,
					Q:        &Q,
				},
				InventoryId: nil,
			},
			Expected:    []*domain.Product{},
			Description: "Search Not Found",
		},
		{
			Queries: &domain.ProductQuery{
				BaseQuery: common.BaseQuery{
					Page:     nil,
					PageSize: nil,
					Q:        &Q,
				},
			},
			Expected:    []*domain.Product{},
			Description: "Search Happy Case",
		},
		{
			Queries: &domain.ProductQuery{
				BaseQuery: common.BaseQuery{
					Page:     utils.GetDataTypeAddress(uint(0)),
					PageSize: utils.GetDataTypeAddress(uint(5)),
					Q:        &Q,
				},
			},
			Expected: []*domain.Product{
				&initData[1],
			},
			Description: "Search With All Params",
		},
		{
			Queries: &domain.ProductQuery{
				InventoryId: utils.GetDataTypeAddress(uint(10)),
			},
			Expected:    []*domain.Product{},
			Description: "Query By Inventory Id Not Found",
		},
		{
			Queries: &domain.ProductQuery{
				InventoryId: utils.GetDataTypeAddress(uint(1)),
			},
			Expected:    []*domain.Product{&initData[0], &initData[1]},
			Description: "Query By Inventory Id Happy Case",
		},
		{
			Queries: &domain.ProductQuery{
				BaseQuery: common.BaseQuery{
					Page:     utils.GetDataTypeAddress(uint(1)),
					PageSize: utils.GetDataTypeAddress(uint(2)),
					Q:        utils.GetDataTypeAddress("free"),
				},
				InventoryId: utils.GetDataTypeAddress(uint(2)),
			},
			Expected:    []*domain.Product{&initData[2]},
			Description: "Query By Inventory Id All Params",
		},
	}

	for _, c := range testCases {
		s.Run(c.Description, func() {
			s.Repo.On("Find", c.Queries).Return(c.Expected, nil).Once()
			products, err := s.UseCase.Execute(c.Queries)

			s.Assertions.NoError(err)
			s.Assertions.Equal(c.Expected, products)
		})
	}
}

func TestFindProductUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(FindProductTestSuite))
}
