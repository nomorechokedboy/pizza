package repository

import (
	"api/src/category/domain"
	"api/src/category/infrastructure"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryGormRepo struct {
	DB *gorm.DB
}

func (repo *CategoryGormRepo) Insert(req *domain.WriteCategoryBody) (*domain.Category, error) {
	category := infrastructure.Category{Name: req.Name, Description: &req.Description}
	if result := repo.DB.Create(&category); result.Error != nil {
		return nil, result.Error
	}

	returning := domain.Category{Id: category.ID, Name: category.Name, Description: *category.Description}

	return &returning, nil
}

func (repo *CategoryGormRepo) Update(id *int, req *domain.WriteCategoryBody) (*domain.Category, error) {
	updateReq := infrastructure.Category{Name: req.Name, Description: &req.Description}
	result := repo.DB.Model(&infrastructure.Category{}).Where("id = ?", id).Updates(&updateReq)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	returning := domain.Category{Id: updateReq.ID, Description: *updateReq.Description, Name: updateReq.Name}
	return &returning, nil
}

func (repo *CategoryGormRepo) Delete(req *int) (*domain.Category, error) {
	deleteReq := infrastructure.Category{Model: gorm.Model{ID: uint(*req)}}
	result := repo.DB.Clauses(clause.Returning{}).Delete(&deleteReq)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	returning := domain.Category{Id: deleteReq.ID, Description: *deleteReq.Description, Name: deleteReq.Name}

	return &returning, nil
}

func (repo *CategoryGormRepo) FindOne(id *int) (*domain.Category, error) {
	dbCategory := infrastructure.Category{Model: gorm.Model{ID: uint(*id)}}

	if result := repo.DB.First(&dbCategory); result.Error != nil {
		return nil, result.Error
	}

	category := domain.Category{Id: dbCategory.ID, Description: *dbCategory.Description, Name: dbCategory.Name}

	return &category, nil
}

func (repo *CategoryGormRepo) Find(req *domain.CategoryQuery) (*[]domain.Category, error) {
	var dbCategories []infrastructure.Category
	categories := make([]domain.Category, 0)

	if result := repo.DB.Where("name = ?", req.Q).Or("description = ?", req.Q).Limit(int(req.PageSize)).Offset(int(req.Page)).Find(&dbCategories); result.Error != nil {
		return nil, result.Error
	}

	for _, dbCategory := range dbCategories {
		categories = append(categories, domain.Category{Id: dbCategory.ID, Description: *dbCategory.Description, Name: dbCategory.Name})
	}

	return &categories, nil
}
