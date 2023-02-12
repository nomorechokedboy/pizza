package usecases

import "api/src/product/domain"

type ProductRepository interface {
	Delete(id uint) (*domain.Product, error)
	FindOne(id uint) (*domain.Product, error)
	Find(queries *domain.ProductQuery) ([]*domain.Product, error)
	Insert(req domain.ProductReq) (*domain.Product, error)
	Update(id uint, req domain.ProductReq) (*domain.Product, error)
}
