package usecases

import "api/src/user/domain"

type CreateUserUseCase struct {
	Repo UserRepository
}

func (useCase *CreateUserUseCase) Execute(req *domain.CreateUserReq) (*domain.User, error) {
	return useCase.Repo.Insert(req)
}
