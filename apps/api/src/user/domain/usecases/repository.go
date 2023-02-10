package usecases

import "api/src/user/domain"

type UserRepository interface {
	Insert(req *domain.CreateUserReq) (*domain.User, error)
	Update(id *int, req *domain.CreateUserReq) (*domain.User, error)
	Delete(id *int) (*domain.User, error)
	Find(req *domain.UserQuery) (*[]domain.User, error)
}
