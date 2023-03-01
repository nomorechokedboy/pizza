package usecases

import (
	"api/src/cartItem/domain"
	"api/src/common"
	"errors"
)

type InsertCartItemUseCase struct {
	Repo      CartItemRepository
	Validator common.Validate
}

func (u *InsertCartItemUseCase) Execute(uid uint, req domain.WriteCartItemRequest) (*domain.CartItem, error) {
	err := u.Validator.Exec(req)
	if err != nil {
		return nil, errors.New("invalid data")
	}

	return u.Repo.Insert(uid, req)
}
