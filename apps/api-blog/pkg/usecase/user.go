package usecase

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(req entities.SignUpBody) error
	GetUserById(id uint) (*entities.User, error)
	GetUserByUsername(username string) (*entities.User, error)
	GetUserByIdentifier(indentifier string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUserInfo(password, fullName, username, phone, email, avatar string, uId uint) error
	UpdatePasswordById(password string, userId uint) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (usecase *userUsecase) CreateUser(req entities.SignUpBody) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return err
	}
	user := &entities.User{
		Identifier: req.Email,
		Password:   string(hashPassword),
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

func (usecase *userUsecase) UpdateUserInfo(password, fullName, username, phone, email, avatar string, uId uint) error {
	var hashPassword []byte
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	if password == "" {
		hashPassword = []byte(password)
	}
	user := &entities.User{
		Id:          uId,
		Password:    string(hashPassword),
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
	var hashPassword []byte
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user := &entities.User{
		Id:       userId,
		Password: string(hashPassword),
	}
	return usecase.repo.UpdateUserInfo(user)

}
