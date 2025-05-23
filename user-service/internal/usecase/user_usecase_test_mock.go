package usecase

import (
	"ecommerce/user-service/internal/domain"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository implements domain.UserRepository for tests
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}
func (m *MockUserRepository) GetByID(id string) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *MockUserRepository) GetByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *MockUserRepository) Update(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}
func (m *MockUserRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *MockUserRepository) List() ([]domain.User, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}
