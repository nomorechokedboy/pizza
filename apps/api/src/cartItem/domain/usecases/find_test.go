package usecases_test

import (
	"api/src/cart/domain"
	"api/src/cartItem/domain/usecases"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FindCartTestSuite struct {
	suite.Suite
	U usecases.FindCartItemsUseCase
	R MockCartItemRepository
}

func (s *FindCartTestSuite) SetupTest() {
	s.R = MockCartItemRepository{}
	s.U = usecases.FindCartItemsUseCase{Repo: &s.R}
}

func (s *FindCartTestSuite) TestFindWithErrors() {
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
			s.R.On("Find", uid).Return(nil, errors.New(c.Error)).Once()
			cart, err := s.U.Execute(uid)

			s.Assertions.EqualError(err, c.Error)
			s.Assertions.Nil(cart)

		})
	}
}

func (s *FindCartTestSuite) TestFindHappyCase() {
	expected := &domain.Cart{Total: 10, UserId: uid}
	s.R.On("Find", uid).Return(expected, nil)
	cart, err := s.U.Execute(uid)

	s.Assertions.NoError(err)
	s.Assertions.Equal(expected, cart)
}

func TestFindCartUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(FindCartTestSuite))
}
