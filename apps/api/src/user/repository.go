package user

import (
	"api/src/user/domain"
	"errors"
	"time"
)

type UserInMemoryRepo struct {
	UserList []domain.User
	IsErr    bool
}

func (repo *UserInMemoryRepo) Insert(req *domain.CreateUserReq) (*domain.User, error) {
	if repo.IsErr {
		return nil, errors.New("connection error")
	}

	for _, user := range repo.UserList {
		if user.UserName == req.UserName {
			return nil, errors.New("user email already exist")
		}
	}

	Id := len(repo.UserList) + 1
	newUser := domain.User{Id: int32(Id), CreatedAt: time.Now(), UpdatedAt: time.Now(), FirstName: req.FirstName, LastName: req.LastName, UserName: req.UserName, Email: req.Email, Password: req.Password}
	repo.UserList = append(repo.UserList, newUser)
	return &newUser, nil
}
