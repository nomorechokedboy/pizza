package gorm

import (
	categoryGorm "api/src/category/repository/gorm"
	"api/src/common"
	inventory "api/src/inventory/domain"
	"api/src/scopes"
	"fmt"
	"strings"

	"api/src/product/domain"
	"api/src/utils"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductGormRepo struct {
	Conn *gorm.DB
}

func (r *ProductGormRepo) Insert(req *domain.ProductReq) (*domain.Product, error) {
	categoryRepo := categoryGorm.CategoryGormRepo{DB: r.Conn}
	category, err := categoryRepo.FindOne(utils.GetDataTypeAddress(int(req.CategoryId)))
	if err != nil {
		return nil, err
	}

	product := domain.Product{
		CategoryID:  req.CategoryId,
		Inventory:   inventory.Inventory{Quantity: req.Quantity},
		Category:    *category,
		Slug:        utils.Slug(req.Name),
		Description: req.Description,
		Name:        req.Name,
		SKU:         req.SKU,
		Price:       req.Price,
	}
	if result := r.Conn.Joins(clause.Associations).Create(&product); result.Error != nil {
		if result.Error.(*pgconn.PgError).Code == "23505" {
			return nil, errors.New("resource exist")
		}
		if result.Error.(*pgconn.PgError).Code == "23503" {
			return nil, errors.New("resource conflict")
		}

		return nil, errors.New("unknown error")
	}

	return &product, nil
}

func (r *ProductGormRepo) Delete(id uint) (*domain.Product, error) {
	deleteReq := domain.Product{ID: id}
	result := r.Conn.Clauses(clause.Returning{}).Delete(&deleteReq)
	if result.Error != nil {
		return nil, errors.New("unknown error")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	return &deleteReq, nil
}

func (r *ProductGormRepo) Update(id uint, req domain.ProductReq) (*domain.Product, error) {
	categoryRepo := categoryGorm.CategoryGormRepo{DB: r.Conn}
	category, err := categoryRepo.FindOne(utils.GetDataTypeAddress(int(req.CategoryId)))
	if err != nil {
		return nil, err
	}
	updateProduct, err := r.FindOne(id)
	if err == nil && updateProduct == nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	updateProduct.Category = *category
	updateProduct.Inventory.Quantity = req.Quantity
	updateProduct.Name = req.Name
	updateProduct.Description = req.Description
	updateProduct.SKU = req.SKU
	updateProduct.Price = req.Price
	updateProduct.Slug = utils.Slug(req.Name)

	result := r.Conn.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&updateProduct)
	if result.Error != nil {
		if result.Error.(*pgconn.PgError).Code == "23505" {
			return nil, errors.New("resource exist")
		}
		if result.Error.(*pgconn.PgError).Code == "23503" {
			return nil, errors.New("resource conflict")
		}

		return nil, errors.New("unknown error")
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return updateProduct, nil
}

func (r *ProductGormRepo) FindOne(id uint) (*domain.Product, error) {
	product := domain.Product{ID: id}

	if result := r.Conn.Joins("Inventory").Joins("Category").First(&product); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New("unknown error")
	}

	return &product, nil
}

func (r *ProductGormRepo) Find(queries *domain.ProductQuery) (common.BasePaginationResponse[domain.Product], error) {
	var products []*domain.Product
	baseRes := &common.BasePaginationResponse[domain.Product]{}
	queryBuilder := r.Conn.Scopes(scopes.Pagination(products, &queries.BaseQuery, baseRes, r.Conn))

	if queries.CategoryId != nil {
		queryBuilder.Where("category_id = ?", queries.CategoryId)
	}

	if queries.Q != nil {
		queryBuilder.Where("products.name ILIKE ?", utils.EscapeLike("%", "%", strings.ToLower(*queries.Q))).Or("products.description ILIKE ?", utils.EscapeLike("%", "%", strings.ToLower(*queries.Q))).Or("products.sku ILIKE ?", utils.EscapeLike("%", "%", strings.ToLower(*queries.Q)))
	}

	if result := queryBuilder.Joins("Category").Joins("Inventory").Find(&products); result.Error != nil {
		return common.BasePaginationResponse[domain.Product]{}, errors.New("unknown error")
	}
	baseRes.Items = products
	fmt.Println("DEBUT find product: ", products)

	return *baseRes, nil
}
