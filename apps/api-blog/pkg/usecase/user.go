package usecase

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(password, username, fullName, phone, email, avatar, identifier string) error
	GetUserById(id uint) (*entities.User, error)
	GetUserByUsername(username string) (*entities.User, error)
	GetUserByIdentifier(indentifier string) (*entities.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (usecase *userUsecase) CreateUser(password, username, fullName, phone, email, avatar, identifier string) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user := &entities.User{
		Identifier:  identifier,
		Username:    username,
		Password:    string(hashPassword),
		Fullname:    fullName,
		PhoneNumber: phone,
		Email:       email,
		Avatar:      avatar,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}
	return usecase.repo.Create(user)

}

func (usecase *userUsecase) GetUserById(id uint) (*entities.User, error) {
	return usecase.repo.GetUserById(id)
}

func (usecase *userUsecase) GetUserByUsername(username string) (*entities.User, error) {
	return usecase.repo.GetUserByUsername(username)
}

func (usecase *userUsecase) GetUserByIdentifier(identifier string) (*entities.User, error) {
	return usecase.repo.GetUserByIdentifier(identifier)
}
