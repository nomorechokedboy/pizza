package gorm_test

import (
	"api/src/category/domain"
	"api/src/utils"
)

func (s *RepositoryIntegrationTestSuite) TestUpdateCategoryRepository() {
	req := domain.WriteCategoryBody{Name: "Updated", Description: "Updated"}

	s.Run("Record not found", func() {
		category, err := s.Repo.Update(utils.GetDataTypeAddress(100), &req)

		s.Assertions.Nil(category)
		s.Assertions.Error(err, "not found")
	})

	s.Run("Happy case", func() {
		category, err := s.Repo.Update(utils.GetDataTypeAddress(2), &req)

		s.Assertions.NoError(err)
		s.Assertions.Equal(uint(2), category.ID)

		updatedCategory, err := s.Repo.FindOne(utils.GetDataTypeAddress(2))
		s.Assertions.NoError(err)
		s.Assertions.Equal(category, updatedCategory)
	})
}
