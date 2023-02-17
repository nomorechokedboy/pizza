package gorm

import (
	"api/src/scopes"
	"api/src/user/domain"
	apiUtils "api/src/utils"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserGormRepo struct {
	DB *gorm.DB
}

func (repo *UserGormRepo) Insert(req *domain.CreateUserReq) (*domain.User, error) {
	user := domain.User{Identifier: req.Identifier, CreatedAt: time.Now(), UpdatedAt: time.Now(), Gender: req.Gender, FullName: req.FullName, BirthDate: req.BirthDate, Email: req.Email, Password: req.Password, PhoneNumber: req.PhoneNumber}
	if result := repo.DB.Create(&user); result.Error != nil {
		if result.Error.(*pgconn.PgError).Code == "23505" {
			return nil, errors.New("resource exist")
		}
		return nil, errors.New("unknown error")
	}

	return &user, nil
}

func (repo *UserGormRepo) Update(id *int, req *domain.CreateUserReq) (*domain.User, error) {
	updateUserReq := domain.User{Id: *id}
	result := repo.DB.Model(&updateUserReq).Debug().Clauses(clause.Returning{}).Updates(&domain.User{Identifier: req.Identifier, UpdatedAt: time.Now(), Gender: req.Gender, FullName: req.FullName, BirthDate: req.BirthDate, Email: req.Email, Password: req.Password, PhoneNumber: req.PhoneNumber})
	if result.Error != nil {
		if result.Error.(*pgconn.PgError).Code == "23505" {
			return nil, errors.New("resource exist")
		}
		return nil, errors.New("unknown error")
	}

	if result.RowsAffected == 0 { //LIMIT 1
		return nil, errors.New("not found")
	}

	return &updateUserReq, nil
}

func (repo *UserGormRepo) Delete(req *int) (*domain.User, error) {
	deleteUserReq := domain.User{Id: *req}
	result := repo.DB.Clauses(clause.Returning{}).Delete(&deleteUserReq)
	if result.Error != nil {
		return nil, errors.New("unknown error")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return &deleteUserReq, nil

}

func (repo *UserGormRepo) FindOne(id *int) (*domain.User, error) {
	dbUser := domain.User{Id: *id}

	if result := repo.DB.First(&dbUser); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		}

		return nil, errors.New("unknown error")
	}

	return &dbUser, nil
}

func (repo *UserGormRepo) Find(req *domain.UserQuery) (*[]domain.User, error) {
	var users []domain.User
	queryBuilder := repo.DB.Scopes(scopes.Pagination(&req.BaseQuery))

	if req.Q != nil {
		queryBuilder = queryBuilder.Where("Identifier ILIKE ?", apiUtils.EscapeLike("%", "%", strings.ToLower(*req.Q))).Or("FullName ILIKE ?", apiUtils.EscapeLike("%", "%", strings.ToLower(*req.Q))).Or("PhoneNumber ILIKE ?", apiUtils.EscapeLike("%", "%", strings.ToLower(*req.Q))).Or("FullName ILIKE ?", apiUtils.EscapeLike("%", "%", strings.ToLower(*req.Q))).Or("PassWord ILIKE ?", apiUtils.EscapeLike("%", "%", strings.ToLower(*req.Q))).Or("Gender ILIKE ?", apiUtils.EscapeLike("%", "%", strings.ToLower(*req.Q)))
	}

	if result := queryBuilder.Find(&users); result.Error != nil {
		return nil, errors.New("unknown error")
	}
	return &users, nil
}
