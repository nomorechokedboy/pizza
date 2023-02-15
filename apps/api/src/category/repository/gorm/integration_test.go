package gorm_test

import (
	"api/src/category/domain"
	GormRepo "api/src/category/repository/gorm"
	"api/src/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

var seeds = []domain.Category{
	{
		Name:        "Food",
		Description: nil,
	},
	{
		Name:        "Drink",
		Description: utils.GetDataTypeAddress("So fresh"),
	},
	{
		Name:        "Dessert",
		Description: utils.GetDataTypeAddress("Sweet and sour"),
	},
	{
		Name:        "Best seller",
		Description: utils.GetDataTypeAddress("People always looking for this"),
	},
	{
		Name:        "New dishes",
		Description: nil,
	},
}

type RepositoryIntegrationTestSuite struct {
	suite.Suite
	Repo GormRepo.CategoryGormRepo
}

func (s *RepositoryIntegrationTestSuite) SetupTest() {
	db := utils.SetupGormIntegrationTest(&s.Suite, domain.Category{})

	for _, seed := range seeds {
		db.Create(&seed)
	}
	s.Repo = GormRepo.CategoryGormRepo{DB: db}
}

func (s *RepositoryIntegrationTestSuite) TearDownTest() {
	s.Repo.DB.Exec("DELETE FROM categories")
	err := s.Repo.DB.Migrator().DropTable(&domain.Category{})
	s.Assertions.NoError(err)
}

func TestCategoryRepository(t *testing.T) {
	suite.Run(t, new(RepositoryIntegrationTestSuite))
}
