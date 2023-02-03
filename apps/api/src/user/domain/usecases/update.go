package usecases

import (
	"api/src/user/domain"
	"errors"
)

type UpdateUserUseCase struct {
	Repo UserRepository
}

func (useCase *UpdateUserUseCase) Execute(id *int, req *domain.CreateUserReq) (*domain.User, error) {
	updateUser, err := useCase.Repo.Update(id, req)
	if updateUser == nil && err == nil {
		return updateUser, errors.New("not found")
	}

	return updateUser, err
}
