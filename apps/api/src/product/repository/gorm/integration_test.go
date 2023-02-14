package gorm_test

import (
	category "api/src/category/domain"
	CategoryGorm "api/src/category/repository/gorm"
	inventory "api/src/inventory/domain"
	"api/src/product/domain"
	ProductGorm "api/src/product/repository/gorm"
	"api/src/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

var seeds = []domain.Product{
	{
		Slug:        utils.Slug("Hehe boi"),
		Description: utils.GetDataTypeAddress("Gorm is so bruh"),
		Name:        "Pizza mozarella",
		SKU:         "lmao123",
		Price:       123,
		Category:    category.Category{Name: "Test Category", Description: utils.GetDataTypeAddress("Lmao help me")},
		Inventory:   inventory.Inventory{Quantity: 5},
	},
	{
		Slug:        utils.Slug("Sorry brah"),
		Description: utils.GetDataTypeAddress("You're excused"),
		Name:        "I am not your brah",
		SKU:         "Brah456",
		Price:       789123,
		Category:    category.Category{Name: "New Category lmao", Description: utils.GetDataTypeAddress("Oh who is shee")},
		Inventory:   inventory.Inventory{Quantity: 500},
	},
	{
		Slug:        utils.Slug("Sadge"),
		Description: utils.GetDataTypeAddress("I don't known what to say"),
		Name:        "Zoolander",
		SKU:         "Pog789",
		Price:       567,
		Category:    category.Category{Name: "And another one", Description: nil},
		Inventory:   inventory.Inventory{Quantity: 25},
	},
}

type ProductIntegrationTestSuite struct {
	suite.Suite
	Repo         ProductGorm.ProductGormRepo
	CategoryRepo CategoryGorm.CategoryGormRepo
}

func (s *ProductIntegrationTestSuite) SetupSuite() {
	db := utils.SetupGormIntegrationTest(&s.Suite, domain.Product{})

	for _, seed := range seeds {
		db.Create(&seed)
		db.Save(&seed)
	}

	s.Repo = ProductGorm.ProductGormRepo{Conn: db}
	s.CategoryRepo = CategoryGorm.CategoryGormRepo{DB: db}
}

func (s *ProductIntegrationTestSuite) TearDownSuite() {
	s.Repo.Conn.Exec("DROP TABLE products")
	err := s.Repo.Conn.Migrator().DropTable(&domain.Product{})
	s.Assertions.NoError(err)
}

func TestProductRepository(t *testing.T) {
	suite.Run(t, new(ProductIntegrationTestSuite))
}
