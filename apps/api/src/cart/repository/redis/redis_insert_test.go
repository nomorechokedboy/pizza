package redis_test

import "api/src/cartItem/domain"

func (s *CartIntegrationTestSuite) TestInsertHappyCases() {
	req := domain.WriteCartItemRequest{ProductID: 1, Quantity: 1}
	uid := uint(1)
	cart, err := s.R.Insert(uid, req)

	s.Assertions.NoError(err)
	s.Assertions.Equal(req.ProductID, cart.CartItems[0].ProductID)
	s.Assertions.Equal(req.Quantity, cart.CartItems[0].Quantity)

	cart, err = s.R.Insert(uid, req)

	s.Assertions.NoError(err)
	s.Assertions.Equal(req.ProductID, cart.CartItems[0].ProductID)
	s.Assertions.Equal(req.Quantity*2, cart.CartItems[0].Quantity)
}
