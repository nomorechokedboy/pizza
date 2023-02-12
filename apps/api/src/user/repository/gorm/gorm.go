package repository

import (
	"api/src/user/domain"
	"time"

	"gorm.io/gorm"
)

type UserGormRepo struct {
	DB *gorm.DB
}

func (repo *UserGormRepo) Insert(req *domain.CreateUserReq) (*domain.User, error) {
	user := domain.User{Identifier: req.Identifier, CreatedAt: time.Now(), UpdatedAt: time.Now(), Gender: req.Gender, FullName: req.FullName, BirthDate: req.BirthDate, Email: req.Email, Password: req.Password, PhoneNumber: req.PhoneNumber}
	if result := repo.DB.Create(&user); result.Error != nil {
		return nil, result.Error
	}

	returning := domain.User{Id: user.Id, Identifier: req.Identifier, CreatedAt: time.Now(), UpdatedAt: time.Now(), Gender: req.Gender, FullName: req.FullName, BirthDate: req.BirthDate, Email: req.Email, Password: req.Password, PhoneNumber: req.PhoneNumber}

	return &returning, nil
}
