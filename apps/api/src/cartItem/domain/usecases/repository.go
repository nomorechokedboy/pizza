package usecases

import (
	"api/src/cart/domain"
	CartItemDomain "api/src/cartItem/domain"
)

type CartRepository interface {
	Insert(uid uint, req CartItemDomain.WriteCartItemRequest) (*domain.Cart, error)
	Find(uid uint) (*domain.Cart, error)
	Delete(uid, productID uint) (*domain.Cart, error)
	Update(uid uint, req CartItemDomain.WriteCartItemRequest) (*domain.Cart, error)
}
