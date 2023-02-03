package usecases_test

import (
	"api/src/user/domain"
	"api/src/user/domain/usecases"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var updateUserMemRepo = usecases.UpdateUserUseCase{Repo: &userMemRepo}
var req = domain.CreateUserReq{Identifier: "test update Iden", FullName: "lmao", Email: "test@example.com", Password: "test123", PhoneNumber: "0909091231", Gender: true, BirthDate: time.Date(2001, 4, 12, 0, 0, 0, 0, time.UTC)}
var id = 0

func TestUpdateUserHappyCase(t *testing.T) {
	assert := assert.New(t)
	userMemRepo.UserList = append(userMemRepo.UserList, domain.User{Identifier: "test update Iden", FullName: "lmao", Email: "test@example.com", Password: "test123", PhoneNumber: "0909091231", Gender: true, BirthDate: time.Date(2001, 4, 12, 0, 0, 0, 0, time.UTC)})
	updateUser, err := updateUserMemRepo.Execute(&id, &req)

	assert.Nil(err)
	assert.Equal(updateUser.Identifier, req.Identifier)
	assert.Equal(updateUser.FullName, req.FullName)
	assert.Equal(updateUser.Email, req.Email)
	assert.Equal(updateUser.PhoneNumber, req.PhoneNumber)
	assert.Equal(updateUser.Password, req.Password)
	assert.Equal(updateUser.BirthDate, req.BirthDate)
	assert.Equal(updateUser.Gender, req.Gender)
}

func TestUpdateUserWithUnknownError(t *testing.T) {
	assert := assert.New(t)
	userMemRepo.IsErr = true
	updateUser, err := updateUserMemRepo.Execute(&id, &req)
	assert.Nil(updateUser)
	assert.EqualError(err, "unknown error")
	userMemRepo.IsErr = false
}
