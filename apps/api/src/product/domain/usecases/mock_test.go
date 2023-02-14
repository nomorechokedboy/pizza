package usecases_test

import (
	"api/src/product/domain"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Insert(req domain.ProductReq) (*domain.Product, error) {
	args := m.Called(req)

	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	return args.Get(0).(*domain.Product), nil
}

func (m *MockRepository) Update(id uint, req domain.ProductReq) (*domain.Product, error) {
	args := m.Called(id, req)

	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	if args.Get(0) == nil && args.Get(1) == nil {
		return nil, nil
	}

	return args.Get(0).(*domain.Product), nil
}

func (m *MockRepository) Delete(id uint) (*domain.Product, error) {
	args := m.Called(id)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	if args.Get(0) == nil && args.Get(1) == nil {
		return nil, nil
	}

	return args.Get(0).(*domain.Product), nil
}

func (m *MockRepository) FindOne(id uint) (*domain.Product, error) {
	args := m.Called(id)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	if args.Get(0) == nil && args.Get(1) == nil {
		return nil, nil
	}

	return args.Get(0).(*domain.Product), nil
}

func (m *MockRepository) Find(queries *domain.ProductQuery) ([]*domain.Product, error) {
	args := m.Called(queries)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	return args.Get(0).([]*domain.Product), nil
}
