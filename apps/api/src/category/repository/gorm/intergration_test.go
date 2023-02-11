package gorm_test

import (
	"api/src/category/domain"
	GormRepo "api/src/category/repository/gorm"
	"api/src/config"
	"api/src/utils"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	config, err := config.LoadEnv()
	s.Require().NoError(err)
	db, err := gorm.Open(postgres.Open(utils.GetDbURI(config)), &gorm.Config{})
	s.Require().NoError(err)
	err = db.AutoMigrate(&domain.Category{})
	s.Require().NoError(err)

	db.Create(&seeds)
	s.Repo = GormRepo.CategoryGormRepo{DB: db}
}

func (s *RepositoryIntegrationTestSuite) TearDownTest() {
	s.Repo.DB.Exec("DELETE FROM categories")
}

func TestCategoryRepository(t *testing.T) {
	suite.Run(t, new(RepositoryIntegrationTestSuite))
}
