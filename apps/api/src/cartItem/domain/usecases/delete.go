package usecases

import (
	"api/src/cart/domain"
	"errors"
)

type DeleteCartItemUseCase struct {
	Repo CartRepository
}

func (u *DeleteCartItemUseCase) Execute(uid uint, productID uint) (*domain.Cart, error) {
	cart, err := u.Repo.Delete(uid, productID)
	if cart == nil && err == nil {
		return cart, errors.New("not found")
	}

	return cart, err
}
