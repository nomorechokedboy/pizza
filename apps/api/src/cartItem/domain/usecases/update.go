package usecases

import (
	"api/src/cart/domain"
	CartItemDomain "api/src/cartItem/domain"
	"api/src/common"
	"errors"
)

type UpdateCartItemUseCase struct {
	Repo      CartRepository
	Validator common.Validate
}

func (u *UpdateCartItemUseCase) Execute(uid uint, req CartItemDomain.WriteCartItemRequest) (*domain.Cart, error) {
	updatedCart, err := u.Repo.Update(uid, req)
	if updatedCart == nil && err == nil {
		return updatedCart, errors.New("not found")
	}

	return updatedCart, err
}
