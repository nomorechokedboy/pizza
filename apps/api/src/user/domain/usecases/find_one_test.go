package usecases_test

import (
	"api/src/user/domain"
	"api/src/user/domain/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

var findOneUseCase = usecases.FindOneUserUseCase{Repo: &userMemRepo}

func TestFindOneUserUseCaseHappyCase(t *testing.T) {
	assert := assert.New(t)
	id := 10
	userMemRepo.UserList = append(userMemRepo.UserList, domain.User{Id: 10})
	user, err := findOneUseCase.Execute(&id)

	assert.Nil(err)
	assert.NotNil(user)
	assert.Equal(user.Id, id)
}

func TestFindOneUserWithUnknownError(t *testing.T) {
	assert := assert.New(t)
	id := 1
	userMemRepo.IsErr = true
	user, err := findOneUseCase.Execute(&id)

	assert.Nil(user)
	assert.EqualError(err, "unknown error")
	userMemRepo.IsErr = false
}

func TestFindOneUserWithNotFoundError(t *testing.T) {
	assert := assert.New(t)
	id := 1
	user, err := findOneUseCase.Execute(&id)

	assert.Nil(user)
	assert.EqualError(err, "not found")

}
