package gorm_test

import (
	"api/src/category/domain"
	"api/src/utils"
	"fmt"
)

func (s *RepositoryIntegrationTestSuite) TestUpdateCategoryRepository() {
	req := domain.WriteCategoryBody{Name: "Updated", Description: "Updated"}

	s.Run("Record not found", func() {
		category, err := s.Repo.Update(utils.GetDataTypeAddress(100), &req)

		s.Assertions.Nil(category)
		s.Assertions.EqualError(err, "not found")
	})

	s.Run("Test unique constraint when update", func() {
		firstUpdate, err := s.Repo.Update(utils.GetDataTypeAddress(3), &req)
		s.Assertions.NoError(err)

		category, err := s.Repo.Update(utils.GetDataTypeAddress(4), &req)
		fmt.Println(category, err, firstUpdate)
		s.Assertions.EqualError(err, "resource exist")
	})

	s.Run("Happy case", func() {
		req := domain.WriteCategoryBody{Name: "Happy case", Description: "Updated"}
		category, err := s.Repo.Update(utils.GetDataTypeAddress(2), &req)

		s.Assertions.NoError(err)
		s.Assertions.Equal(uint(2), category.ID)

		updatedCategory, err := s.Repo.FindOne(utils.GetDataTypeAddress(2))
		s.Assertions.NoError(err)
		s.Assertions.Equal(category, updatedCategory)
	})
}
