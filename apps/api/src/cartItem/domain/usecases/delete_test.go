package usecases_test

import (
	CartDomain "api/src/cart/domain"
	"api/src/cartItem/domain"
	"api/src/cartItem/domain/usecases"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DeleteCartItemTestSuite struct {
	suite.Suite
	U *usecases.DeleteCartItemUseCase
	R *MockCartItemRepository
}

func (s *DeleteCartItemTestSuite) SetupTest() {
	s.R = &MockCartItemRepository{}
	s.U = &usecases.DeleteCartItemUseCase{Repo: s.R}
}

func (s *DeleteCartItemTestSuite) TestDeleteWithErrors() {
	table := []struct {
		Description string
		Error       error
		Expected    string
	}{
		{
			Description: "Unknown error",
			Error:       errors.New("unknown error"),
			Expected:    "unknown error",
		},
		{
			Description: "Not found",
			Error:       nil,
			Expected:    "not found",
		},
	}

	for _, c := range table {
		s.Run(c.Description, func() {
			s.R.On("Delete", uint(1), uint(10)).Return(nil, c.Error).Once()
			cart, err := s.U.Execute(1, 10)

			s.Assertions.Nil(cart)
			s.Assertions.EqualError(err, c.Expected)
		})
	}
}

func (s *DeleteCartItemTestSuite) TestDeleteHappyCase() {
	expected := &CartDomain.Cart{Total: 0, UserId: 1, CartItems: make([]*domain.CartItem, 0)}
	s.R.On("Delete", uint(1), uint(1)).Return(expected, nil)
	cart, err := s.U.Execute(1, 1)

	s.Assertions.NoError(err)
	s.Assertions.Equal(expected, cart)
}

func TestDeleteCartItemUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteCartItemTestSuite))
}
