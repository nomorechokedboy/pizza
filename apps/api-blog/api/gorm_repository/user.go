package gorm_repository

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/repository"
	"errors"

	"gorm.io/gorm"
)

type UserGormRepo struct {
	db *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) repository.UserRepository {
	return &UserGormRepo{
		db: db,
	}
}

func (r *UserGormRepo) Create(user *entities.User) error {
	return r.db.Create(&user).Error
}

func (r *UserGormRepo) GetUserById(id uint) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserGormRepo) GetUserByUsername(username string) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, "username = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserGormRepo) GetUserByIdentifier(identifier string) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, "identifier = ?", identifier).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}
