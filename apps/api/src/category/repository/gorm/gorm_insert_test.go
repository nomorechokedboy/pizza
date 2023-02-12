package gorm_test

import (
	"api/src/category/domain"
)

func (s *RepositoryIntegrationTestSuite) TestInsertCategoryRepository() {
	s.Run("Duplicate record", func() {
		req := domain.WriteCategoryBody{Name: seeds[0].Name, Description: "This must be conflict"}
		category, err := s.Repo.Insert(&req)

		s.Assertions.EqualError(err, "resource exist")
		s.Assertions.Nil(category)
	})

	s.Run("Happy case", func() {
		req := domain.WriteCategoryBody{Name: "Excuse me", Description: "Test connection between units/services"}
		category, err := s.Repo.Insert(&req)

		s.Assertions.NoError(err)
		s.Assertions.NotNil(category)
		s.Assertions.Equal(req.Description, *category.Description)
		s.Assertions.Equal(req.Name, category.Name)
	})
}
