package repository

import (
	"api/src/user/domain"
	"errors"
	"strings"
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
		// if userUpdate.Identifier == req.Identifier {
		// 	errors.New("Identifier already exist")
		// }
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

func (repo *UserInMemoryRepo) Delete(req *int) (*domain.User, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	var res domain.User
	pos := -1
	filteredList := make([]domain.User, 0)

	for i, user := range repo.UserList {
		if user.Id == *req {
			res = user
			pos = i
			continue
		}
		filteredList = append(filteredList, user)
	}
	repo.UserList = filteredList

	if pos < 0 {
		return nil, nil

	}
	return &res, nil
}

func (repo *UserInMemoryRepo) Find(req *domain.UserQuery) (*[]domain.User, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	res := repo.UserList
	if req.Q != nil {
		res = make([]domain.User, 0)
		for _, user := range repo.UserList {
			q := strings.ToLower(*req.Q)
			entityContainsQ := strings.Contains(strings.ToLower(user.Email), q) || strings.Contains(strings.ToLower(user.FullName), q) || strings.Contains(strings.ToLower(user.Identifier), q) || strings.Contains(strings.ToLower(user.PhoneNumber), q) || strings.Contains(strings.ToLower(user.Password), q) || strings.Contains(strings.ToLower(user.BirthDate.Format("2006-01-02 15:04:05")), q)

			if entityContainsQ {
				res = append(res, user)
			}

		}
	}

	start := uint(req.GetPage() * req.GetPageSize())
	end := uint(start + req.GetPageSize())

	res = res[min(start, uint(len(res))):min(end, uint(len(res)))]

	return &res, nil
}

func min(x, y uint) uint {
	if y < x {
		return y
	}
	return x
}

func (repo *UserInMemoryRepo) FindOne(id *int) (*domain.User, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	var res *domain.User
	for _, user := range repo.UserList {
		if user.Id == int(*id) {
			res = &user
			break
		}
	}

	return res, nil
}
