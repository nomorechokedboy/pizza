package usecases_test

import (
	"api/src/category/domain"
	"api/src/category/domain/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

var usecase = usecases.UpdateCategoryUseCase{Repo: &categoryRepo}
var req = domain.WriteCategoryBody{Name: "Updated category", Description: "Updated description"}
var id = 0

func TestUpdateCategoryUseCaseWithUnknownError(t *testing.T) {
	assert := assert.New(t)
	categoryRepo.IsErr = true
	updatedCategory, err := usecase.Execute(&id, &req)

	assert.Nil(updatedCategory)
	assert.EqualError(err, "unknown error")
	categoryRepo.IsErr = false
}

func TestUpdateCategoryUseCaseWithNotFoundError(t *testing.T) {
	assert := assert.New(t)
	categoryRepo.Data = make([]domain.Category, 0)
	updatedCategory, err := usecase.Execute(&id, &req)

	assert.Nil(updatedCategory)
	assert.EqualError(err, "not found")
}

func TestUpdateCategoryHappyCase(t *testing.T) {
	assert := assert.New(t)
	categoryRepo.Data = append(categoryRepo.Data, domain.Category{Name: "Shounen", Description: "Blah blah, bloh bloh description"})
	updatedCategory, err := usecase.Execute(&id, &req)

	assert.Nil(err)
	assert.Equal(updatedCategory.Description, req.Description)
	assert.Equal(updatedCategory.Name, req.Name)
	assert.Equal(updatedCategory.Id, id)
}
