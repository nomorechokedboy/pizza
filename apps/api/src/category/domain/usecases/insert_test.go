package usecases_test

import (
	"api/src/category/domain"
	"api/src/category/domain/usecases"
	"api/src/category/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

var categoryRepo = repository.CategoryInMemoryRepo{Data: make([]domain.Category, 0), IsErr: false}
var insertUseCase = usecases.InsertCategoryUseCase{Repo: &categoryRepo}
var req = domain.WriteCategoryBody{Name: "Comedy", Description: "Funny stuffs"}

func TestInsertCategoryWithUnknownError(t *testing.T) {
	assert := assert.New(t)
	categoryRepo.IsErr = true
	category, err := insertUseCase.Execute(&req)

	assert.EqualError(err, "unknown error")
	assert.Nil(category)
	categoryRepo.IsErr = false
}

func TestInsertCategoryWithDuplicateError(t *testing.T) {
	assert := assert.New(t)
	category, err := insertUseCase.Execute(&req)

	assert.EqualError(err, "resource exist")
	assert.Nil(category)
}

func TestInsertCategoryHappyCase(t *testing.T) {
	assert := assert.New(t)
	category, err := insertUseCase.Execute(&req)

	assert.Nil(err)
	assert.NotNil(category)
	assert.Equal(req.Description, category.Description)
	assert.Equal(req.Name, category.Name)
}
