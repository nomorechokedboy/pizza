package domain

type ProductRepository interface {
	Insert(req *ProductReq) (*Product, error)
}

type CreateProductUseCase struct {
	Repo ProductRepository
}

func (useCase *CreateProductUseCase) Execute(req *ProductReq) (*Product, error) {
	return useCase.Repo.Insert(req)
}
