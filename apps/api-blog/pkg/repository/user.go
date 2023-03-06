package repository

import (
	"api-blog/pkg/entities"
)

type UserRepository interface {
	GetUserById(id uint) (*entities.User, error)
	Create(user *entities.User) error
	GetUserByUsername(username string) (*entities.User, error)
	GetUserByIdentifier(identifier string) (*entities.User, error)
	UpdateUserInfo(user *entities.User) error
}
