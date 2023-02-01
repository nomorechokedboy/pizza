package domain

type UserRepository interface {
	Insert(req *CreateUserReq) (*User, error)
}

type CreateUserUseCase struct {
	Repo UserRepository
}

func (useCase *CreateUserUseCase) Execute(req *CreateUserReq) (*User, error) {
	return useCase.Repo.Insert(req)
}
