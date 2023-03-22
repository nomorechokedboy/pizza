package usecases_test

import (
	CartDomain "api/src/cart/domain"
	"api/src/cartItem/domain"
	"api/src/cartItem/domain/usecases"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type UpdateCartItemTestSuite struct {
	suite.Suite
	U *usecases.UpdateCartItemUseCase
	R *MockCartItemRepository
}

func (s *UpdateCartItemTestSuite) SetupTest() {
	s.R = &MockCartItemRepository{}
	s.U = &usecases.UpdateCartItemUseCase{Repo: s.R}
}

var updateReq = domain.WriteCartItemRequest{ProductID: 1, Quantity: 1}

func (s *UpdateCartItemTestSuite) TestUpdateWithErrors() {
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
			s.R.On("Update", uint(1), updateReq).Return(nil, c.Error).Once()
			cart, err := s.U.Execute(1, updateReq)

			s.Assertions.Nil(cart)
			s.Assertions.EqualError(err, c.Expected)
		})
	}
}

func (s *UpdateCartItemTestSuite) TestUpdateHappyCase() {
	expected := &CartDomain.Cart{Total: 0, UserId: 1, CartItems: []*domain.CartItem{{
		ProductID: 1,
		Quantity:  1,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}}}
	s.R.On("Update", uint(1), updateReq).Return(expected, nil)
	cart, err := s.U.Execute(1, updateReq)

	s.Assertions.NoError(err)
	s.Assertions.Equal(expected, cart)
}

func TestUpdateCartItemUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateCartItemTestSuite))
}
