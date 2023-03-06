package usecases_test

import (
	cartDomain "api/src/cart/domain"
	"api/src/cartItem/domain"

	"github.com/stretchr/testify/mock"
)

type MockCartItemRepository struct {
	mock.Mock
}

func (m *MockCartItemRepository) Insert(uid uint, req domain.WriteCartItemRequest) (*cartDomain.Cart, error) {
	args := m.Called(uid, req)

	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	return args.Get(0).(*cartDomain.Cart), nil
}

func (m *MockCartItemRepository) Update(uid uint, req domain.WriteCartItemRequest) (*cartDomain.Cart, error) {
	args := m.Called(uid, req)

	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	if args.Get(0) == nil && args.Get(1) == nil {
		return nil, nil
	}

	return args.Get(0).(*cartDomain.Cart), nil
}

func (m *MockCartItemRepository) Delete(uid uint, productID uint) (*cartDomain.Cart, error) {
	args := m.Called(uid, productID)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	if args.Get(0) == nil && args.Get(1) == nil {
		return nil, nil
	}

	return args.Get(0).(*cartDomain.Cart), nil
}

func (m *MockCartItemRepository) Find(uid uint) (*cartDomain.Cart, error) {
	args := m.Called(uid)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	return args.Get(0).(*cartDomain.Cart), nil
}
