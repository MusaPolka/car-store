package usecase

import (
	"ecommerce/user-service/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserUsecase_Create(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)

	user := &domain.User{ID: "1", Username: "alice"}
	mockRepo.On("Create", user).Return(nil)

	err := usecase.Create(user)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Create", user)
}

func TestUserUsecase_GetByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)

	user := &domain.User{ID: "1", Username: "alice"}
	mockRepo.On("GetByID", "1").Return(user, nil)

	result, err := usecase.GetByID("1")
	assert.NoError(t, err)
	assert.Equal(t, user, result)
}

func TestUserUsecase_GetByUsername(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)

	user := &domain.User{ID: "1", Username: "alice"}
	mockRepo.On("GetByUsername", "alice").Return(user, nil)

	result, err := usecase.GetByUsername("alice")
	assert.NoError(t, err)
	assert.Equal(t, user, result)
}

func TestUserUsecase_Update(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)

	user := &domain.User{ID: "1", Username: "alice"}
	mockRepo.On("Update", user).Return(nil)

	err := usecase.Update(user)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Update", user)
}

func TestUserUsecase_Delete(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)

	mockRepo.On("Delete", "1").Return(nil)

	err := usecase.Delete("1")
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Delete", "1")
}

func TestUserUsecase_List(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)

	users := []domain.User{{ID: "1", Username: "alice"}, {ID: "2", Username: "bob"}}
	mockRepo.On("List").Return(users, nil)

	result, err := usecase.List()
	assert.NoError(t, err)
	assert.Equal(t, users, result)
}
