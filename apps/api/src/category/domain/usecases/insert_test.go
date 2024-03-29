package usecases_test

import (
	"api/src/category/domain"
	"api/src/category/domain/usecases"
	"api/src/category/repository"
	"api/src/common"
	"api/src/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var categoryRepo = repository.CategoryInMemoryRepo{Data: make([]*domain.Category, 0), IsErr: false}
var insertUseCase = usecases.InsertCategoryUseCase{Repo: &categoryRepo, Validator: &common.ValidatorAdapter}

func TestInsertCategoryUseCaseWithUnknownError(t *testing.T) {
	assert := assert.New(t)
	categoryRepo.IsErr = true
	req := domain.WriteCategoryBody{Name: "Comedy", Description: "Funny stuffs"}
	category, err := insertUseCase.Execute(&req)

	assert.EqualError(err, "unknown error")
	assert.Nil(category)
	categoryRepo.IsErr = false
}

func TestInsertCategoryUseCaseWithDuplicateError(t *testing.T) {
	assert := assert.New(t)
	categoryRepo.Data = append(categoryRepo.Data, &domain.Category{ID: 1, Name: "Comedy", Description: utils.GetDataTypeAddress("Another description")})
	req := domain.WriteCategoryBody{Name: "Comedy", Description: "Funny stuffs"}
	category, err := insertUseCase.Execute(&req)

	assert.EqualError(err, "resource exist")
	assert.Nil(category)
}

func TestInsertCategoryUseCaseHappyCase(t *testing.T) {
	assert := assert.New(t)
	req := domain.WriteCategoryBody{Name: "Happy case", Description: "Funny stuffs"}
	category, err := insertUseCase.Execute(&req)

	assert.Nil(err)
	assert.NotNil(category)
	assert.Equal(req.Description, *category.Description)
	assert.Equal(req.Name, category.Name)
}

func TestInsertCategoryUseCaseWithInvalidData(t *testing.T) {
	assert := assert.New(t)
	req := domain.WriteCategoryBody{Name: "A", Description: "Funny stuffs"}
	category, err := insertUseCase.Execute(&req)

	assert.EqualError(err, "invalid data")
	assert.Nil(category)
}
