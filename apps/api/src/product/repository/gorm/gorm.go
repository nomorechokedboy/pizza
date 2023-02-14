package gorm

import (
	inventory "api/src/inventory/domain"

	"api/src/product/domain"
	"api/src/utils"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type ProductGormRepo struct {
	Conn *gorm.DB
}

func (r *ProductGormRepo) Insert(req *domain.ProductReq) (*domain.Product, error) {
	product := domain.Product{
		CategoryID:  req.CategoryId,
		Inventory:   inventory.Inventory{Quantity: req.Quantity},
		Slug:        utils.Slug(req.Name),
		Description: req.Description,
		Name:        req.Name,
		SKU:         req.SKU,
		Price:       req.Price,
	}
	if result := r.Conn.Create(&product); result.Error != nil {
		if result.Error.(*pgconn.PgError).Code == "23505" {
			return nil, errors.New("resource exist")
		}

		return nil, errors.New("unknown error")
	}

	return &product, nil
}

func (r *ProductGormRepo) Delete(id uint) (*domain.Product, error) {
	return nil, nil
}

func (r *ProductGormRepo) Update(id uint, req domain.ProductReq) (*domain.Product, error) {
	return nil, nil
}

func (r *ProductGormRepo) FindOne(id uint) (*domain.Product, error) {
	return nil, nil
}

func (r *ProductGormRepo) Find(queries *domain.ProductQuery) ([]*domain.Product, error) {
	return nil, nil
}
