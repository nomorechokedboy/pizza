package usecases

import (
	"api/src/user/domain"
	"errors"
)

type DeleteUserUseCase struct {
	Repo UserRepository
}

func (useCase *DeleteUserUseCase) Execute(id *int) (*domain.User, error) {
	deletedUser, err := useCase.Repo.Delete(id)
	if deletedUser == nil && err == nil {
		return deletedUser, errors.New("not found")
	}

	return deletedUser, err
}
