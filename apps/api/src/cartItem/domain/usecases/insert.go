package usecases

import (
	"api/src/cart/domain"
	CartItemDomain "api/src/cartItem/domain"
	"api/src/common"
	"errors"
)

type InsertCartItemUseCase struct {
	Repo      CartRepository
	Validator common.Validate
}

func (u *InsertCartItemUseCase) Execute(uid uint, req CartItemDomain.WriteCartItemRequest) (*domain.Cart, error) {
	err := u.Validator.Exec(req)
	if err != nil {
		return nil, errors.New("invalid data")
	}

	return u.Repo.Insert(uid, req)
}
