package usecases_test

import (
	"api/src/common"
	"api/src/product/domain"

	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Insert(req *domain.ProductReq) (*domain.Product, error) {
	args := m.Called(*req)

	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	return args.Get(0).(*domain.Product), nil
}

func (m *MockProductRepository) Update(id uint, req domain.ProductReq) (*domain.Product, error) {
	args := m.Called(id, req)

	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	if args.Get(0) == nil && args.Get(1) == nil {
		return nil, nil
	}

	return args.Get(0).(*domain.Product), nil
}

func (m *MockProductRepository) Delete(id uint) (*domain.Product, error) {
	args := m.Called(id)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	if args.Get(0) == nil && args.Get(1) == nil {
		return nil, nil
	}

	return args.Get(0).(*domain.Product), nil
}

func (m *MockProductRepository) FindOne(id uint) (*domain.Product, error) {
	args := m.Called(id)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	if args.Get(0) == nil && args.Get(1) == nil {
		return nil, nil
	}

	return args.Get(0).(*domain.Product), nil
}

func (m *MockProductRepository) Find(queries *domain.ProductQuery) (common.BasePaginationResponse[domain.Product], error) {
	args := m.Called(queries)
	if args.Get(1) != nil {
		return common.BasePaginationResponse[domain.Product]{}, args.Get(1).(error)
	}

	return args.Get(0).(common.BasePaginationResponse[domain.Product]), nil
}
