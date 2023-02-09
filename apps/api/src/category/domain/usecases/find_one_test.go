package usecases_test

import (
	"api/src/category/domain"
	"api/src/category/domain/usecases"
	"api/src/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var findOneUseCase = usecases.FindOneCategoryUseCase{Repo: &categoryRepo}

func TestFindOneCategoryUseCaseWithUnknownError(t *testing.T) {
	assert := assert.New(t)
	id := 1
	categoryRepo.IsErr = true
	category, err := findOneUseCase.Execute(&id)

	assert.Nil(category)
	assert.EqualError(err, "unknown error")
	categoryRepo.IsErr = false
}

func TestFindOneCategoryUseCaseWithNotFoundError(t *testing.T) {
	assert := assert.New(t)
	id := 1
	category, err := findOneUseCase.Execute(&id)

	assert.Nil(category)
	assert.EqualError(err, "not found")
}

func TestFindOneCategoryUseCaseHappyCase(t *testing.T) {
	assert := assert.New(t)
	id := 11
	categoryRepo.Data = append(categoryRepo.Data, domain.Category{ID: 11, Name: "Test Category", Description: utils.GetDataTypeAddress("Lalalalalalalala")})
	category, err := findOneUseCase.Execute(&id)

	assert.Nil(err)
	assert.NotNil(category)
	assert.Equal(category.ID, uint(id))
}
