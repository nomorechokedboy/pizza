package domain

import cartItemDomain "api/src/cartItem/domain"

type Cart struct {
	Total     uint
	UserId    uint
	CartItems []*cartItemDomain.CartItem
}

type CartRedis struct {
	Total     uint   `redis:"total" js:"total"`
	UserId    uint   `redis:"user_id" js:"userId"`
	CartItems string `redis:"cart_items" js:"cartItems"`
}
