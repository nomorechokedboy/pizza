package usecases_test

import (
	"api/src/category/domain"
	"api/src/category/domain/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

var findUseCase = usecases.FindCategoryUseCase{Repo: &categoryRepo}

func TestFindCategoryUseCaseWithUnknownError(t *testing.T) {
	assert := assert.New(t)
	category, err := findUseCase.Execute(nil)

	assert.Nil(category)
	assert.EqualError(err, "unknown error")
}

func TestFindCategoryUseCaseHappyCase(t *testing.T) {
	assert := assert.New(t)
	expected := []domain.Category{{Id: 1, Name: "Test", Description: "123"}, {Id: 2, Name: "Test123", Description: "Traitor's requime"}}
	categoryRepo.Data = append(categoryRepo.Data, expected[0], expected[1])
	categories, err := findUseCase.Execute(nil)

	assert.Nil(err)
	assert.Equal(categories, expected)
}

func TestFindCategoryUseCasePagination(t *testing.T) {
	assert := assert.New(t)
	expected := []domain.Category{{Id: 1, Name: "Test", Description: "123"}, {Id: 2, Name: "Test123", Description: "Traitor's requime"}}
	categoryRepo.Data = append(categoryRepo.Data, expected[0], expected[1])
	categories, err := findUseCase.Execute(&domain.CategoryQuery{Page: 0, PageSize: 1})

	assert.Nil(err)
	assert.Equal(categories, []domain.Category{expected[0]})

	categories, err = findUseCase.Execute(&domain.CategoryQuery{Page: 0, PageSize: 1})

	assert.Nil(err)
	assert.Equal(categories, []domain.Category{expected[1]})
}

func TestFindCategoryUseCaseWithSearchQuery(t *testing.T) {
	assert := assert.New(t)
	expected := []domain.Category{{Id: 1, Name: "Test", Description: "123"}, {Id: 2, Name: "Test123", Description: "Traitor's requiem"}}
	categoryRepo.Data = append(categoryRepo.Data, expected[0], expected[1])
	categories, err := findUseCase.Execute(&domain.CategoryQuery{Q: "requiem"})

	assert.Nil(err)
	assert.Equal(categories, []domain.Category{expected[1]})
}
