package gorm_test

import (
	"api/src/product/domain"
	"api/src/utils"
)

func (s *ProductIntegrationTestSuite) TestDeleteProductNotFound() {
	product, err := s.Repo.Delete(100)

	s.Assertions.EqualError(err, "not found")
	s.Assertions.Nil(product)
}

func (s *ProductIntegrationTestSuite) TestDeleteProductHappyCase() {
	createdProduct, err := s.Repo.Insert(&domain.ProductReq{
		Description: utils.GetDataTypeAddress("Orm suck"),
		Name:        "Hail sqlx",
		SKU:         "sqlx-forever",
		Price:       789,
		CategoryId:  1,
		Quantity:    500,
	})
	s.Assertions.NoError(err)
	product, err := s.Repo.Delete(createdProduct.ID)
	s.Assertions.NoError(err)
	s.Assertions.Equal(createdProduct.ID, product.ID)

	product, err = s.Repo.FindOne(createdProduct.ID)
	s.Assertions.Nil(err)
	s.Assertions.Nil(product)
}
