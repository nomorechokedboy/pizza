package usecases

import (
	cartDomain "api/src/cart/domain"
)

type FindCartItemsUseCase struct {
	Repo CartRepository
}

func (u *FindCartItemsUseCase) Execute(uid uint) (*cartDomain.Cart, error) {
	return u.Repo.Find(uid)
}
