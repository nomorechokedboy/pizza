package gorm_test

import (
	"api/src/utils"
)

func (s *RepositoryIntegrationTestSuite) TestDeleteCategoryRepository() {
	s.Run("Record not found", func() {
		category, err := s.Repo.Delete(utils.GetDataTypeAddress(100))

		s.Assertions.EqualError(err, "not found")
		s.Assertions.Nil(category)
	})

	s.Run("Happy case", func() {
		category, err := s.Repo.Delete(utils.GetDataTypeAddress(1))
		s.Assertions.NoError(err)
		s.Assertions.Equal(uint(1), category.ID)

		category, err = s.Repo.FindOne(utils.GetDataTypeAddress(1))
		s.Assertions.EqualError(err, "not found")
		s.Assertions.Nil(category)
	})
}
