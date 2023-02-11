package gorm

import (
	"api/src/category/domain"
	apiUtils "api/src/utils"
	"errors"

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

	returning := domain.Category{ID: category.ID, Name: category.Name, Description: category.Description}

	return &returning, nil
}

func (repo *CategoryGormRepo) Update(id *int, req *domain.WriteCategoryBody) (*domain.Category, error) {
	updateReq := domain.Category{Name: req.Name, Description: &req.Description}
	result := repo.DB.Model(&domain.Category{}).Where("id = ?", id).Updates(&updateReq)
	if result.Error != nil {
		return nil, errors.New("unknown error")
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	returning := domain.Category{ID: updateReq.ID, Description: updateReq.Description, Name: updateReq.Name}
	return &returning, nil
}

func (repo *CategoryGormRepo) Delete(req *int) (*domain.Category, error) {
	deleteReq := domain.Category{ID: uint(*req)}
	result := repo.DB.Clauses(clause.Returning{}).Delete(&deleteReq)
	if result.Error != nil {
		return nil, errors.New("unknown error")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	returning := domain.Category{ID: deleteReq.ID, Description: deleteReq.Description, Name: deleteReq.Name}

	return &returning, nil
}

func (repo *CategoryGormRepo) FindOne(id *int) (*domain.Category, error) {
	dbCategory := domain.Category{ID: uint(*id)}

	if result := repo.DB.First(&dbCategory); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		}

		return nil, errors.New("unknown error")
	}

	category := domain.Category{ID: dbCategory.ID, Description: dbCategory.Description, Name: dbCategory.Name}

	return &category, nil
}

func (repo *CategoryGormRepo) Find(req *domain.CategoryQuery) (*[]domain.Category, error) {
	var categories []domain.Category
	queryBuilder := repo.DB.Limit(int(req.PageSize)).Offset(int(req.Page))

	if req.Q != nil {
		queryBuilder = queryBuilder.Where("name LIKE ?", apiUtils.EscapeLike("%", "%", *req.Q)).Or("description LIKE ?", apiUtils.EscapeLike("%", "%", *req.Q))
	}

	if result := queryBuilder.Find(&categories); result.Error != nil {
		return nil, errors.New("unknown error")
	}

	return &categories, nil
}
