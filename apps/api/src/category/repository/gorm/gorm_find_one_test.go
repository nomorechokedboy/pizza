package gorm_test

import "api/src/utils"

func (s *RepositoryIntegrationTestSuite) TestFindOneCategoryRepository() {
	s.Run("Record not found", func() {
		category, err := s.Repo.FindOne(utils.GetDataTypeAddress(100))

		s.Assertions.Error(err, "not found")
		s.Assertions.Nil(category)
	})

	s.Run("Happy case", func() {
		category, err := s.Repo.FindOne(utils.GetDataTypeAddress(2))

		s.Assertions.Nil(err)
		s.Assertions.Equal(uint(2), category.ID)
	})
}
