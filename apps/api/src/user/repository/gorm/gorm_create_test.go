package gorm_test

import (
	"api/src/user/domain"
	"time"
)

func (s *RepositoryIntegrationTestSuite) TestCreateUserRepository() {
	s.Run("Duplicate record", func() {
		req := domain.CreateUserReq{
			Identifier:  "079201017970",
			FullName:    "Nguyen Van A",
			Email:       "nguyenvana@gmail.com",
			Password:    "nguyenvana123",
			PhoneNumber: "0234567891",
			BirthDate:   time.Date(2001, 4, 12, 0, 0, 0, 0, time.UTC),
			Gender:      true}
		user, err := s.Repo.Insert(&req)

		s.Assertions.EqualError(err, "resource exist")
		s.Assertions.Nil(user)

	})

	s.Run("Happy case", func() {
		req := domain.CreateUserReq{
			Identifier:  "079201017975",
			FullName:    "Nguyen Van T",
			Email:       "nguyenvant@gmail.com",
			Password:    "nguyenvant123",
			PhoneNumber: "0678912345",
			BirthDate:   time.Date(1993, 12, 15, 0, 0, 0, 0, time.UTC),
			Gender:      true,
		}
		user, err := s.Repo.Insert(&req)

		s.Assertions.NoError(err)
		s.Assertions.NotNil(user)
		s.Assertions.Equal(req.Identifier, user.Identifier)
		s.Assertions.Equal(req.FullName, user.FullName)
		s.Assertions.Equal(req.Email, user.Email)
		s.Assertions.Equal(req.BirthDate, user.BirthDate)
		s.Assertions.Equal(req.Password, user.Password)
		s.Assertions.Equal(req.Gender, user.Gender)
		s.Assertions.Equal(req.PhoneNumber, user.PhoneNumber)
	})
}
