package usecase

import (
	"ecommerce/order-service/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Create(order *domain.Order) error {
	args := m.Called(order)
	return args.Error(0)
}
func (m *MockOrderRepository) GetByID(id string) (*domain.Order, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Order), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *MockOrderRepository) Update(order *domain.Order) error {
	args := m.Called(order)
	return args.Error(0)
}
func (m *MockOrderRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *MockOrderRepository) List() ([]domain.Order, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]domain.Order), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *MockOrderRepository) ListByUser(userID string) ([]domain.Order, error) {
	args := m.Called(userID)
	if args.Get(0) != nil {
		return args.Get(0).([]domain.Order), args.Error(1)
	}
	return nil, args.Error(1)
}
