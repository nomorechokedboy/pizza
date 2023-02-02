package usecases

import "api/src/user/domain"

type UserRepository interface {
	Insert(req *domain.CreateUserReq) (*domain.User, error)
	Update(id *int32, req *domain.CreateUserReq) (*domain.User, error)
}
