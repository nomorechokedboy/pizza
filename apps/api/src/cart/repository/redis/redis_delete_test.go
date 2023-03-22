package redis_test

import (
	"api/src/cartItem/domain"
	"api/src/cartItem/domain/usecases"
	"api/src/common"
)

func (s *CartIntegrationTestSuite) TestDeleteHappyCase() {
	insert := usecases.InsertCartItemUseCase{Repo: &s.R, Validator: &common.ValidatorAdapter}
	cart, err := insert.Execute(11, domain.WriteCartItemRequest{ProductID: 1, Quantity: 1})
	s.Assertions.NoError(err)

	deletedCart, err := s.R.Delete(11, 1)
	s.Assertions.NoError(err)
	s.Assertions.Equal(uint(0), deletedCart.Total)
	s.Assertions.Equal(cart.UserId, deletedCart.UserId)
	s.Assertions.Equal(0, len(deletedCart.CartItems))
}

func (s *CartIntegrationTestSuite) TestDeleteNotFound() {
	deletedCart, err := s.R.Delete(10, 11)

	s.Assertions.Nil(deletedCart)
	s.Assertions.Nil(err)
}
