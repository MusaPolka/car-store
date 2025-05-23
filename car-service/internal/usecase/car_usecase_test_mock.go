package usecase

import (
	"ecommerce/car-service/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockCarRepository struct {
	mock.Mock
}

func (m *MockCarRepository) Create(car *domain.Car) error {
	args := m.Called(car)
	return args.Error(0)
}
func (m *MockCarRepository) GetByID(id string) (*domain.Car, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Car), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *MockCarRepository) Update(car *domain.Car) error {
	args := m.Called(car)
	return args.Error(0)
}
func (m *MockCarRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *MockCarRepository) List(page, limit int) ([]domain.Car, error) {
	args := m.Called(page, limit)
	if args.Get(0) != nil {
		return args.Get(0).([]domain.Car), args.Error(1)
	}
	return nil, args.Error(1)
}
