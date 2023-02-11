package gorm_test

import "api/src/utils"

func (s *RepositoryIntegrationTestSuite) TestFindOneCategoryRepository() {
	s.Run("Record not found", func() {
		category, err := s.Repo.FindOne(utils.GetDataTypeAddress(100))

		s.Assertions.Error(err, "not found")
		s.Assertions.Nil(category)
	})

	s.Run("Happy case", func() {
		category, err := s.Repo.FindOne(utils.GetDataTypeAddress(int(seeds[0].ID)))

		s.Assertions.Nil(err)
		s.Assertions.Equal(seeds[0].ID, category.ID)
		s.Assertions.Equal(seeds[0].Description, category.Description)
		s.Assertions.Equal(seeds[0].Name, category.Name)
	})
}
