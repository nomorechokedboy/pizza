package gorm_test

import (
	"api/src/inventory/domain"
	InventoryGorm "api/src/inventory/repository/gorm"
	"api/src/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

var seeds = []domain.Inventory{
	{
		Quantity: 1,
	},
}

type InventoryIntegrationTestSuite struct {
	suite.Suite
	Repo InventoryGorm.InventoryGormRepo
}

func (s *InventoryIntegrationTestSuite) SetupTest() {
	db := utils.SetupGormIntegrationTest(&s.Suite, domain.Inventory{})
	db.Create(&seeds[0])

	s.Repo = InventoryGorm.InventoryGormRepo{Conn: db}
}

func (s *InventoryIntegrationTestSuite) TearDownTest() {
	err := s.Repo.Conn.Migrator().DropTable(&domain.Inventory{})
	s.Assertions.NoError(err)
}

func TestInventoryRepository(t *testing.T) {
	suite.Run(t, new(InventoryIntegrationTestSuite))
}
