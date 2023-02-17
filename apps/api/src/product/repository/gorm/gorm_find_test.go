package gorm_test

import (
	"api/src/product/domain"
)

func (s *ProductIntegrationTestSuite) TestFindProduct() {
	table := []struct {
		Query    domain.ProductQuery
		Expected int
	}{
		// {
		// 	Query:    domain.ProductQuery{CategoryId: utils.GetDataTypeAddress(uint(2))},
		// 	Expected: 1,
		// },
		// {
		// 	Query: domain.ProductQuery{BaseQuery: common.BaseQuery{
		// 		Q: utils.GetDataTypeAddress("brah"),
		// 	}},
		// 	Expected: 1,
		// },
	}

	for _, c := range table {
		products, err := s.Repo.Find(&c.Query)

		s.Assertions.NoError(err)
		s.Assertions.Equal(c.Expected, len(products.Items))
	}
}
