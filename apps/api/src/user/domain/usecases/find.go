package usecases

import "api/src/user/domain"

type FindUserUseCase struct {
	Repo UserRepository
}

func (useCase *FindUserUseCase) Execute(req *domain.UserQuery) (*[]domain.User, error) {
	return useCase.Repo.Find(req)
}
