package gorm_test

import (
	"api/src/user/domain"
	GormRepo "api/src/user/repository/gorm"
	"api/src/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

var seeds = []domain.User{
	{
		Identifier:  "079201017970",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		FullName:    "Nguyen Van A",
		Email:       "nguyenvana@gmail.com",
		Password:    "nguyenvana123",
		PhoneNumber: "0234567891",
		BirthDate:   time.Date(2001, 4, 12, 0, 0, 0, 0, time.UTC),
		Gender:      true,
	},
	{
		Identifier:  "079201017971",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		FullName:    "Pham Van B",
		Email:       "phamvanb@gmail.com",
		Password:    "phamvanb123",
		PhoneNumber: "0345678910",
		BirthDate:   time.Date(1999, 3, 12, 0, 0, 0, 0, time.UTC),
		Gender:      true,
	},
	{
		Identifier:  "079301017972",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		FullName:    "Le Thi D",
		Email:       "lethid@gmail.com",
		Password:    "lethid123",
		PhoneNumber: "0456789012",
		BirthDate:   time.Date(2003, 7, 20, 0, 0, 0, 0, time.UTC),
		Gender:      false,
	},
	{
		Identifier:  "079301017973",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		FullName:    "Truong Thi E",
		Email:       "truongthie@gmail.com",
		Password:    "truongthie123",
		PhoneNumber: "0567890123",
		BirthDate:   time.Date(2001, 4, 12, 0, 0, 0, 0, time.UTC),
		Gender:      false,
	},
}

type RepositoryIntegrationTestSuite struct {
	suite.Suite
	Repo GormRepo.UserGormRepo
}

func (s *RepositoryIntegrationTestSuite) SetupTest() {
	db := utils.SetupGormIntegrationTest(&s.Suite, domain.User{})

	for _, seed := range seeds {
		db.Create(&seed)
	}
	s.Repo = GormRepo.UserGormRepo{DB: db}
}

func (s *RepositoryIntegrationTestSuite) TearDownTest() {
	s.Repo.DB.Exec("DELETE FROM users")
	err := s.Repo.DB.Migrator().DropTable(&domain.User{})
	s.Assertions.NoError(err)
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(RepositoryIntegrationTestSuite))
}
