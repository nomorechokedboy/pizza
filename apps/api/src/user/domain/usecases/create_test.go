package usecases_test

import (
	"api/src/user"
	"api/src/user/domain"
	"api/src/user/domain/usecases"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserUseCaseHappyCase(t *testing.T) {
	assert := assert.New(t)
	userMemRepo := user.UserInMemoryRepo{UserList: make([]domain.User, 0), IsErr: false}
	createUserUseCase := usecases.CreateUserUseCase{Repo: &userMemRepo}
	req := domain.CreateUserReq{Identifier: "079201017970", FullName: "Tommy Xiaomi", Gender: true, BirthDate: time.Date(2001, 4, 12, 0, 0, 0, 0, time.UTC), PhoneNumber: "0589168067", Email: "kientin123@gmail.com", Password: "aloalo123123"}
	user, err := createUserUseCase.Execute(&req)
	assert.Nil(err)
	t.Logf("Type of User is %v\n", reflect.TypeOf(user))
}

func TestCreateUserUseCaseUnknown(t *testing.T) {
	assert := assert.New(t)
	userMemRepo := user.UserInMemoryRepo{UserList: make([]domain.User, 0), IsErr: false}
	userMemRepo.IsErr = true
	createUserUseCase := usecases.CreateUserUseCase{Repo: &userMemRepo}
	req := domain.CreateUserReq{Identifier: "079201017970", FullName: "Tommy Xiaomi", Gender: true, BirthDate: time.Date(2001, 4, 12, 0, 0, 0, 0, time.UTC), PhoneNumber: "0589168067", Email: "kientin123@gmail.com", Password: "aloalo123123"}
	user, err := createUserUseCase.Execute(&req)
	assert.EqualError(err, "connection error")
	assert.Nil(user)
}

func TestCreateUserUseCaseWithDuplicateError(t *testing.T) {
	assert := assert.New(t)
	userMemRepo := user.UserInMemoryRepo{UserList: make([]domain.User, 0), IsErr: false}
	userMemRepo.Insert(&domain.CreateUserReq{Identifier: "079201017973", FullName: "Shelby", Gender: true, BirthDate: time.Date(2001, 4, 12, 0, 0, 0, 0, time.UTC), PhoneNumber: "0561013932", Email: "shelby@gmail.com", Password: "shelby123123"})
	createUserUseCase := usecases.CreateUserUseCase{Repo: &userMemRepo}
	req := domain.CreateUserReq{Identifier: "079201017970", FullName: "Shelby", Gender: true, BirthDate: time.Date(2001, 4, 12, 0, 0, 0, 0, time.UTC), PhoneNumber: "0561013932", Email: "shelby@gmail.com", Password: "shelby123123"}
	user, err := createUserUseCase.Execute(&req)
	assert.EqualError(err, "user email already exist")
	assert.Nil(user)
}
