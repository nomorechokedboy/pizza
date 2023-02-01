package domain_test

import (
	"api/src/user"
	"api/src/user/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserUseCaseUnknown(t *testing.T) {
	assert := assert.New(t)
	userMemRepo := user.UserInMemoryRepo{UserList: make([]domain.User, 0), IsErr: false}
	userMemRepo.IsErr = true
	createUserUseCase := domain.CreateUserUseCase{Repo: &userMemRepo}
	req := domain.CreateUserReq{FirstName: "Tommy", LastName: "XiaoMi", UserName: "kientin123", Email: "kientin123@gmail.com", Password: "aloalo123123"}
	user, err := createUserUseCase.Execute(&req)
	assert.EqualError(err, "connection error")
	assert.Nil(user)
}
