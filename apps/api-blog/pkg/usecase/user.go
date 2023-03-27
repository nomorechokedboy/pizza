package usecase

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/repository"
	"time"
)

type UserUsecase interface {
	CreateUser(req entities.SignUpBody) error
	GetUserById(id uint) (*entities.User, error)
	GetUserByUsername(username string) (*entities.User, error)
	GetUserByIdentifier(indentifier string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUserInfo(fullName, username, phone, email, avatar string, uId uint) error
	UpdatePasswordById(password string, userId uint) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (usecase *userUsecase) CreateUser(req entities.SignUpBody) error {
	user := &entities.User{
		Identifier: req.Email,
		Password:   req.Password,
		Username:   *req.Username,
		CreatedAt:  time.Now(),
		UpdateAt:   time.Now(),
	}
	return usecase.repo.Create(user)

}

func (usecase *userUsecase) GetUserById(id uint) (*entities.User, error) {
	return usecase.repo.GetUserById(id)
}

func (usecase *userUsecase) GetUserByEmail(email string) (*entities.User, error) {
	return usecase.repo.GetUserByEmail(email)
}

func (usecase *userUsecase) GetUserByUsername(username string) (*entities.User, error) {
	return usecase.repo.GetUserByUsername(username)
}

func (usecase *userUsecase) GetUserByIdentifier(identifier string) (*entities.User, error) {
	return usecase.repo.GetUserByIdentifier(identifier)
}

func (usecase *userUsecase) UpdateUserInfo(fullName, username, phone, email, avatar string, uId uint) error {
	user := &entities.User{
		Id:          uId,
		PhoneNumber: phone,
		Email:       email,
		Username:    username,
		Avatar:      avatar,
		Fullname:    fullName,
		UpdateAt:    time.Now(),
	}
	return usecase.repo.UpdateUserInfo(user)
}

func (usecase *userUsecase) UpdatePasswordById(password string, userId uint) error {
	user := &entities.User{
		Id:       userId,
		Password: password,
	}
	return usecase.repo.UpdateUserInfo(user)

}
