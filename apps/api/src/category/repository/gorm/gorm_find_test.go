package gorm_test

import (
	"api/src/category/domain"
	"api/src/common"
	"api/src/utils"
)

func (s *RepositoryIntegrationTestSuite) TestFindCategoryRepository() {
	s.Run("Happy case", func() {
		categories, err := s.Repo.Find(&domain.CategoryQuery{BaseQuery: common.BaseQuery{Page: utils.GetDataTypeAddress(uint(0)), PageSize: utils.GetDataTypeAddress(uint(10))}})

		s.Assertions.NoError(err)
		s.Assertions.Equal(5, len(*categories))
	})

	s.Run("Pagination", func() {
		tables := []struct {
			Page     uint
			PageSize uint
			Expected int
		}{
			{
				Page:     1,
				PageSize: 2,
				Expected: 2,
			},
			{
				Page:     0,
				PageSize: 2,
				Expected: 2,
			},
			{
				Page:     2,
				PageSize: 2,
				Expected: 1,
			},
			{
				Page:     3,
				PageSize: 2,
				Expected: 0,
			},
		}

		for _, c := range tables {
			categories, err := s.Repo.Find(&domain.CategoryQuery{BaseQuery: common.BaseQuery{Page: &c.Page, PageSize: &c.PageSize}})

			s.Assertions.NoError(err)
			s.Assertions.Equal(c.Expected, len(*categories))
		}
	})

	s.Run("Search", func() {
		table := []struct {
			Page     uint
			PageSize uint
			Q        *string
			Expected int
		}{
			{
				Page:     0,
				Q:        utils.GetDataTypeAddress("e"),
				Expected: 4,
			},
			{
				Page:     0,
				PageSize: 3,
				Q:        utils.GetDataTypeAddress("e"),
				Expected: 3,
			},
			{
				Page:     1,
				PageSize: 3,
				Q:        utils.GetDataTypeAddress("e"),
				Expected: 1,
			},
			{
				Q:        utils.GetDataTypeAddress("drink"),
				Expected: 1,
			},
		}

		for _, c := range table {
			categories, err := s.Repo.Find(&domain.CategoryQuery{BaseQuery: common.BaseQuery{Page: &c.Page, PageSize: &c.PageSize, Q: c.Q}})

			s.Assertions.NoError(err)
			s.Assertions.Equal(c.Expected, len(*categories))
		}
	})
}
