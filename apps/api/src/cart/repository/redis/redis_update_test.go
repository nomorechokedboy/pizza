package redis_test

import (
	"api/src/cartItem/domain"
	"api/src/cartItem/domain/usecases"
	"api/src/common"
)

func (s *CartIntegrationTestSuite) TestUpdateHappyCase() {
	insert := usecases.InsertCartItemUseCase{Repo: &s.R, Validator: &common.ValidatorAdapter}
	cart, err := insert.Execute(11, domain.WriteCartItemRequest{ProductID: 1, Quantity: 1})
	s.Assertions.NoError(err)

	updateReq := domain.WriteCartItemRequest{ProductID: 1, Quantity: 100}
	updatedCart, err := s.R.Update(11, updateReq)
	s.Assertions.NoError(err)
	s.Assertions.Equal(cart.Total, updatedCart.Total)
	s.Assertions.Equal(cart.UserId, updatedCart.UserId)
	s.Assertions.Equal(updateReq.Quantity, updatedCart.CartItems[0].Quantity)
	s.Assertions.Equal(updateReq.ProductID, updatedCart.CartItems[0].ProductID)
}

func (s *CartIntegrationTestSuite) TestUpdateNotFound() {
	updatedCart, err := s.R.Update(10, domain.WriteCartItemRequest{ProductID: 1, Quantity: 100})

	s.Assertions.Nil(updatedCart)
	s.Assertions.Nil(err)
}
