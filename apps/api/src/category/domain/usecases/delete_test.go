package usecases_test

import (
	"api/src/category/domain"
	"api/src/category/domain/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

var deleteUsecase = usecases.DeleteCategoryUseCase{Repo: &categoryRepo}

func TestDeleteCategoryUseCaseWithUnknownError(t *testing.T) {
	assert := assert.New(t)
	categoryRepo.IsErr = true
	id := 1
	deletedCategory, err := deleteUsecase.Execute(&id)

	assert.Nil(deletedCategory)
	assert.EqualError(err, "unknown error")
	categoryRepo.IsErr = false
}

func TestDeleteCategoryUseCaseWithNotFoundError(t *testing.T) {
	assert := assert.New(t)
	id := 10
	deletedCategory, err := deleteUsecase.Execute(&id)

	assert.Nil(deletedCategory)
	assert.EqualError(err, "not found")
}

func TestDeleteCategoryHappyCase(t *testing.T) {
	assert := assert.New(t)
	id := 1
	categoryRepo.Data = append(categoryRepo.Data, domain.Category{Name: "Shounen", Description: "Blah blah, bloh bloh description", Id: 1})
	deletedCategory, err := deleteUsecase.Execute(&id)

	assert.Nil(err)
	assert.NotNil(deletedCategory)
	assert.Equal(id, deletedCategory.Id)
	assert.Equal(len(categoryRepo.Data), 0)
}
