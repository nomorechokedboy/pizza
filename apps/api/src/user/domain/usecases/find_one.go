package usecases

import (
	"api/src/user/domain"
	"errors"
)

type FindOneUserUseCase struct {
	Repo UserRepository
}

func (useCase *FindOneUserUseCase) Execute(id *int) (*domain.User, error) {
	user, err := useCase.Repo.FindOne(id)
	if user == nil && err == nil {
		return user, errors.New("not found")
	}
	return user, err
}
