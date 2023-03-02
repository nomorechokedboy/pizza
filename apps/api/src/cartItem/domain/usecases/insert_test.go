package usecases_test

import (
	CartDomain "api/src/cart/domain"
	"api/src/cartItem/domain"
	"api/src/cartItem/domain/usecases"
	"api/src/common"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type InsertCartItemTestSuite struct {
	suite.Suite
	U usecases.InsertCartItemUseCase
	R MockCartItemRepository
}

func (s *InsertCartItemTestSuite) SetupTest() {
	s.R = MockCartItemRepository{}
	s.U = usecases.InsertCartItemUseCase{Repo: &s.R, Validator: &common.ValidatorAdapter}
}

var insertReq = domain.WriteCartItemRequest{
	Quantity:  1,
	ProductID: 1,
}

var uid = uint(1)

func (s *InsertCartItemTestSuite) TestInsertWithErrors() {
	table := []struct {
		Description string
		Error       string
	}{
		{
			Description: "Unknown error",
			Error:       "unknown error",
		},
	}

	for _, c := range table {
		s.Run(c.Description, func() {
			s.R.On("Insert", uid, insertReq).Return(nil, errors.New(c.Error)).Once()
			cartItem, err := s.U.Execute(uid, insertReq)

			s.Assertions.EqualError(err, c.Error)
			s.Assertions.Nil(cartItem)

		})
	}
}

func (s *InsertCartItemTestSuite) TestInsertHappyCase() {
	table := []struct {
		Description string
		Expected    *CartDomain.Cart
	}{
		{
			Description: "Happy case",
			Expected:    &CartDomain.Cart{Total: 1, UserId: uid, CartItems: []*domain.CartItem{{ProductID: insertReq.ProductID, Quantity: insertReq.Quantity}}},
		},
		{
			Description: "Duplicate",
			Expected:    &CartDomain.Cart{Total: 1, UserId: uid, CartItems: []*domain.CartItem{{ProductID: insertReq.ProductID, Quantity: insertReq.Quantity * 2}}},
		},
	}

	for _, c := range table {
		s.Run(c.Description, func() {
			s.R.On("Insert", uid, insertReq).Return(c.Expected, nil).Once()
			cartItem, err := s.U.Execute(uid, insertReq)

			s.Assertions.NoError(err)
			s.Assertions.Equal(c.Expected, cartItem)
		})
	}
}

func TestInsertCartItemUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(InsertCartItemTestSuite))
}
