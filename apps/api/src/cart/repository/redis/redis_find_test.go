package redis_test

import (
	cartDomain "api/src/cart/domain"
	"api/src/cartItem/domain"
	"api/src/cartItem/domain/usecases"
	"api/src/common"
)

func (s *CartIntegrationTestSuite) TestFindHappyCase() {
	insert := usecases.InsertCartItemUseCase{Repo: &s.R, Validator: &common.ValidatorAdapter}
	cart, err := insert.Execute(11, domain.WriteCartItemRequest{ProductID: 1, Quantity: 1})
	s.Assertions.NoError(err)

	table := []struct {
		Uid      uint
		Expected *cartDomain.Cart
	}{
		{
			Uid: 10,
			Expected: &cartDomain.Cart{
				CartItems: []*domain.CartItem{},
				UserId:    10,
				Total:     0,
			},
		},
		{
			Uid:      11,
			Expected: cart,
		},
	}

	for _, c := range table {
		cart, err := s.R.Find(c.Uid)

		s.Assertions.NoError(err)
		s.Assertions.Equal(c.Expected.Total, cart.Total)
		s.Assertions.Equal(c.Expected.UserId, cart.UserId)
		if len(cart.CartItems) > 0 {
			s.Assertions.Equal(c.Expected.CartItems[0].ProductID, cart.CartItems[0].ProductID)
			s.Assertions.Equal(c.Expected.CartItems[0].Quantity, cart.CartItems[0].Quantity)
		}
	}
}
