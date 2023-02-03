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
		if user.Identifier == req.Identifier {
			return nil, errors.New("user account has already exist")
		}
	}

	Id := len(repo.UserList) + 1
	newUser := domain.User{Id: int(Id), CreatedAt: time.Now(), UpdatedAt: time.Now(), FullName: req.FullName, Email: req.Email, Password: req.Password, BirthDate: req.BirthDate, Gender: req.Gender, PhoneNumber: req.PhoneNumber, Identifier: req.Identifier}
	repo.UserList = append(repo.UserList, newUser)
	return &newUser, nil
}

func (repo *UserInMemoryRepo) Update(id *int, req *domain.CreateUserReq) (*domain.User, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	for i := range repo.UserList {
		userUpdate := &repo.UserList[i]
		if userUpdate.Id == *id {
			userUpdate.FullName = req.FullName
			userUpdate.BirthDate = req.BirthDate
			userUpdate.Gender = req.Gender
			userUpdate.Identifier = req.Identifier
			userUpdate.PhoneNumber = req.PhoneNumber
			userUpdate.Password = req.Password
			return userUpdate, nil
		}
	}
	return nil, nil
}
