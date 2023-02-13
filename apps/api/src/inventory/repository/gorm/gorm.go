package gorm

import (
	"api/src/inventory/domain"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InventoryGormRepo struct {
	Conn *gorm.DB
}

func (r *InventoryGormRepo) Insert(req *domain.WriteInventoryBody) (*domain.Inventory, error) {
	category := domain.Inventory{Quantity: req.Quantity}
	if result := r.Conn.Create(&category); result.Error != nil {
		if result.Error.(*pgconn.PgError).Code == "23505" {
			return nil, errors.New("resource exist")
		}

		return nil, errors.New("unknown error")
	}

	return &category, nil
}

func (r *InventoryGormRepo) Update(id *int, req *domain.WriteInventoryBody) (*domain.Inventory, error) {
	updateReq := domain.Inventory{ID: uint(*id)}
	result := r.Conn.Model(&updateReq).Debug().Clauses(clause.Returning{}).Updates(&domain.Inventory{Quantity: req.Quantity})
	if result.Error != nil {
		if result.Error.(*pgconn.PgError).Code == "23505" {
			return nil, errors.New("resource exist")
		}
		return nil, errors.New("unknown error")
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	return &updateReq, nil
}

func (r *InventoryGormRepo) Delete(id *int) (*domain.Inventory, error) {
	return nil, nil
}

func (r *InventoryGormRepo) FindOne(id *int) (*domain.Inventory, error) {
	return nil, nil
}
