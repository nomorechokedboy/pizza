package gorm_test

import (
	categoryDomain "api/src/category/domain"
	"api/src/common"
	"api/src/product/domain"
	"api/src/utils"
)

func (s *ProductIntegrationTestSuite) TestFindOneProductNotFound() {
	product, err := s.Repo.FindOne(100)

	s.Assertions.Nil(err)
	s.Assertions.Nil(product)
}

func (s *ProductIntegrationTestSuite) TestFindOneHappyCase() {
	categories, err := s.CategoryRepo.Find(&categoryDomain.CategoryQuery{BaseQuery: common.BaseQuery{Page: utils.GetDataTypeAddress(uint(0))}})
	s.Assertions.NoError(err)
	req := domain.ProductReq{
		Description: seeds[0].Description,
		Name:        seeds[0].Name,
		SKU:         "I am done with gorm",
		Price:       seeds[0].Price,
		CategoryId:  (*categories)[0].ID,
		Quantity:    500,
	}
	s.Assertions.NoError(err)
	createdProduct, err := s.Repo.Insert(&req)
	s.Assertions.NoError(err)

	product, err := s.Repo.FindOne(createdProduct.ID)

	s.Assertions.NoError(err)
	s.Assertions.Equal(createdProduct.ID, product.ID)
	s.Assertions.Equal((*categories)[0], product.Category)
	s.Assertions.Equal(createdProduct.Inventory.Quantity, req.Quantity)
	s.Assertions.Equal(createdProduct.Description, product.Description)
	s.Assertions.Equal(createdProduct.Name, product.Name)
	s.Assertions.Equal(createdProduct.Price, product.Price)
	s.Assertions.Equal(createdProduct.SKU, product.SKU)
	s.Assertions.Equal(createdProduct.Slug, product.Slug)
}
