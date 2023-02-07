package usecases_test

import (
	"api/src/user/domain"
	"api/src/user/domain/usecases"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var deleteUsecase = usecases.DeleteUserUseCase{Repo: &userMemRepo}

func TestDeleteUserUseCaseWithUnknownError(t *testing.T) {
	assert := assert.New(t)
	userMemRepo.IsErr = true
	id := 1
	deleteUser, err := deleteUsecase.Execute(&id)

	assert.Nil(deleteUser)
	assert.EqualError(err, "unknown error")
	userMemRepo.IsErr = false
}

func TestDeleteUserUseCaseWithNotFoundError(t *testing.T) {
	assert := assert.New(t)
	id := 100
	deletedUser, err := deleteUsecase.Execute(&id)

	assert.Nil(deletedUser)
	assert.EqualError(err, "not found")
}

func TestDeleteUserHappyCase(t *testing.T) {
	assert := assert.New(t)
	id := 1
	userMemRepo.UserList = append(userMemRepo.UserList, domain.User{Id: id, Identifier: "testcaseiden123", CreatedAt: time.Now(), UpdatedAt: time.Now(), FullName: "Kakashi", Gender: true, BirthDate: time.Date(2001, 4, 12, 0, 0, 0, 0, time.UTC), PhoneNumber: "0589168077", Email: "kakashi@gmail.com", Password: "kakashi123"})
	deletedUser, err := deleteUsecase.Execute(&id)

	assert.Nil(err)
	assert.NotNil(deletedUser)
	assert.Equal(id, deletedUser.Id)
	assert.Equal(len(userMemRepo.UserList), 0)
}
