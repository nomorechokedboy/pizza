package gorm_test

import (
	"api/src/category/domain"
)

func (s *RepositoryIntegrationTestSuite) TestInsertCategoryRepository() {
	s.Run("Duplicate record", func() {
		req := domain.WriteCategoryBody{Name: seeds[0].Name, Description: "This must be conflict"}
		category, err := s.Repo.Insert(&req)

		s.Assertions.Error(err, "resource exist")
		s.Assertions.Nil(category)
	})

	s.Run("Happy case", func() {
		req := domain.WriteCategoryBody{Name: "Integration", Description: "Test connection between units/services"}
		category, err := s.Repo.Insert(&req)

		s.Assertions.Nil(err)
		s.Assertions.Equal(req.Description, *category.Description)
		s.Assertions.Equal(req.Name, category.Name)
	})
}
