package gorm_test

import (
	categoryDomain "api/src/category/domain"
	"api/src/common"
	"api/src/product/domain"
	"api/src/utils"
)

func (s *ProductIntegrationTestSuite) TestInsertDuplicate() {
	categories, err := s.CategoryRepo.Find(&categoryDomain.CategoryQuery{BaseQuery: common.BaseQuery{Page: utils.GetDataTypeAddress(uint(0))}})
	s.Assertions.NoError(err)

	req := domain.ProductReq{
		Description: utils.GetDataTypeAddress("CC"),
		Name:        "Help me",
		SKU:         "Duplicate sku",
		Price:       100,
		CategoryId:  categories.Items[0].ID,
		Quantity:    100,
	}
	product, err := s.Repo.Insert(&req)
	s.NotNil(product)
	s.Assertions.NoError(err)

	product, err = s.Repo.Insert(&req)

	s.Assertions.EqualError(err, "resource exist")
	s.Assertions.Nil(product)
}

func (s *ProductIntegrationTestSuite) TestInsertHappyCase() {
	categories, err := s.CategoryRepo.Find(&categoryDomain.CategoryQuery{BaseQuery: common.BaseQuery{Page: utils.GetDataTypeAddress(uint(0))}})
	s.Assertions.NoError(err)

	req := domain.ProductReq{
		Description: seeds[0].Description,
		Name:        seeds[0].Name,
		SKU:         "Happy sku",
		Price:       seeds[0].Price,
		CategoryId:  categories.Items[0].ID,
		Quantity:    500,
	}
	product, err := s.Repo.Insert(&req)

	s.Assertions.NoError(err)
	s.Assertions.NotNil(product)
	s.Assertions.Equal(product.Description, req.Description)
	s.Assertions.Equal(product.Name, req.Name)
	s.Assertions.Equal(product.SKU, req.SKU)
	s.Assertions.Equal(product.Price, req.Price)
	s.Assertions.Equal(product.CategoryID, req.CategoryId)
	s.Assertions.Equal(product.Inventory.Quantity, req.Quantity)
}
