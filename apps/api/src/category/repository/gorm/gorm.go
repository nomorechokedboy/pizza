package gorm

import (
	"api/src/category/domain"
	"api/src/common"
	"api/src/scopes"
	apiUtils "api/src/utils"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryGormRepo struct {
	DB *gorm.DB
}

func (repo *CategoryGormRepo) Insert(req *domain.WriteCategoryBody) (*domain.Category, error) {
	category := domain.Category{Name: req.Name, Description: &req.Description}
	if result := repo.DB.Create(&category); result.Error != nil {
		if result.Error.(*pgconn.PgError).Code == "23505" {
			return nil, errors.New("resource exist")
		}

		return nil, errors.New("unknown error")
	}

	return &category, nil
}

func (repo *CategoryGormRepo) Update(id *int, req *domain.WriteCategoryBody) (*domain.Category, error) {
	updateReq := domain.Category{ID: uint(*id)}
	result := repo.DB.Model(&updateReq).Clauses(clause.Returning{}).Updates(&domain.Category{Name: req.Name, Description: &req.Description})
	if result.Error != nil {
		if result.Error.(*pgconn.PgError).Code == "23505" {
			return nil, errors.New("resource exist")
		}
		return nil, errors.New("unknown error")
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &updateReq, nil
}

func (repo *CategoryGormRepo) Delete(req *int) (*domain.Category, error) {
	deleteReq := domain.Category{ID: uint(*req)}
	result := repo.DB.Preload(clause.Associations).Clauses(clause.Returning{}).Delete(&deleteReq)
	if result.Error != nil {
		return nil, errors.New("unknown error")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	return &deleteReq, nil
}

func (repo *CategoryGormRepo) FindOne(id *int) (*domain.Category, error) {
	dbCategory := domain.Category{ID: uint(*id)}

	if result := repo.DB.First(&dbCategory); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		}

		return nil, errors.New("unknown error")
	}

	return &dbCategory, nil
}

func (repo *CategoryGormRepo) Find(req *domain.CategoryQuery) (common.BasePaginationResponse[domain.Category], error) {
	var categories []*domain.Category
	baseRes := &common.BasePaginationResponse[domain.Category]{}
	queryBuilder := repo.DB.Scopes(scopes.Pagination(categories, &req.BaseQuery, baseRes, repo.DB))

	if req.Q != nil {
		queryBuilder.Where("name ILIKE ?", apiUtils.EscapeLike("%", "%", strings.ToLower(*req.Q))).Or("description ILIKE ?", apiUtils.EscapeLike("%", "%", strings.ToLower(*req.Q)))
	}

	if result := queryBuilder.Find(&categories); result.Error != nil {
		return common.BasePaginationResponse[domain.Category]{}, errors.New("unknown error")
	}
	baseRes.Items = categories

	return *baseRes, nil
}
