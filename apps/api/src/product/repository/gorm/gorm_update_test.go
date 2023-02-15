package gorm_test

import (
	"api/src/product/domain"
	"api/src/utils"
)

func (s *ProductIntegrationTestSuite) TestUpdateProductNotFound() {
	req := domain.ProductReq{Name: "Updated", Description: nil, SKU: "Lmao not gonna work", Price: 12345, CategoryId: 1, Quantity: 1}
	product, err := s.Repo.Update(100, req)
	s.Assertions.Nil(product)
	s.Assertions.Nil(err)
}

func (s *ProductIntegrationTestSuite) TestUpdateProductUniqueConstraint() {
	var products []*domain.Product
	s.Repo.Conn.Find(&products)
	req := domain.ProductReq{Name: "Updated", Description: nil, SKU: "Le gorm suck", Price: 12345, CategoryId: 1, Quantity: 1}

	_, err := s.Repo.Update(products[0].ID, req)
	s.Assertions.NoError(err)

	_, err = s.Repo.Update(products[1].ID, req)
	s.Assertions.EqualError(err, "resource exist")
}

func (s *ProductIntegrationTestSuite) TestUpdateProductHappyCase() {
	var products []*domain.Product
	s.Repo.Conn.Find(&products)
	req := domain.ProductReq{Name: "Happy case", Description: utils.GetDataTypeAddress("Updated"), SKU: "New sku lmao", Price: 123, CategoryId: 2, Quantity: 100}
	product, err := s.Repo.Update(products[0].ID, req)

	s.Assertions.NoError(err)
	s.Assertions.Equal(products[0].ID, product.ID)

	updatedProduct, err := s.Repo.FindOne(products[0].ID)
	s.Assertions.NoError(err)
	s.Assertions.Equal(product.ID, updatedProduct.ID)
	s.Assertions.Equal(product.Name, updatedProduct.Name)
	s.Assertions.Equal(product.Description, updatedProduct.Description)
	s.Assertions.Equal(product.SKU, updatedProduct.SKU)
	s.Assertions.Equal(product.Price, updatedProduct.Price)
	s.Assertions.Equal(req.Quantity, updatedProduct.Inventory.Quantity)
	s.Assertions.Equal(product.CategoryID, updatedProduct.CategoryID)
	s.Assertions.Equal(product.Slug, updatedProduct.Slug)
}
