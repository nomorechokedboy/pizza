package domain_test

import (
	"api/src/user"
	"api/src/user/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserUseCaseUnknown(t *testing.T) {
	assert := assert.New(t)
	userMemRepo := user.UserInMemoryRepo{UserList: make([]domain.User, 0), IsErr: false}
	userMemRepo.IsErr = true
	createUserUseCase := domain.CreateUserUseCase{Repo: &userMemRepo}
	req := domain.CreateUserReq{Identifier: "079201017970", FullName: "Tommy Xiaomi", Gender: true, BirthDate: time.Date(2001, 4, 12, 0, 0, 0, 0, time.UTC), PhoneNumber: "0589168067", Email: "kientin123@gmail.com", Password: "aloalo123123"}
	user, err := createUserUseCase.Execute(&req)
	assert.EqualError(err, "connection error")
	assert.Nil(user)
}
