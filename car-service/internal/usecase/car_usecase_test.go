package usecase

import (
	"ecommerce/car-service/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarUsecase_Create(t *testing.T) {
	mockRepo := new(MockCarRepository)
	usecase := NewCarUsecase(mockRepo)

	car := &domain.Car{ID: "1", Name: "BMW"}
	mockRepo.On("Create", car).Return(nil)

	err := usecase.Create(car)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Create", car)
}

func TestCarUsecase_GetByID(t *testing.T) {
	mockRepo := new(MockCarRepository)
	usecase := NewCarUsecase(mockRepo)

	car := &domain.Car{ID: "1", Name: "BMW"}
	mockRepo.On("GetByID", "1").Return(car, nil)

	result, err := usecase.GetByID("1")
	assert.NoError(t, err)
	assert.Equal(t, car, result)
}

func TestCarUsecase_Update(t *testing.T) {
	mockRepo := new(MockCarRepository)
	usecase := NewCarUsecase(mockRepo)

	car := &domain.Car{ID: "1", Name: "Updated BMW"}
	mockRepo.On("Update", car).Return(nil)

	err := usecase.Update(car)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Update", car)
}

func TestCarUsecase_Delete(t *testing.T) {
	mockRepo := new(MockCarRepository)
	usecase := NewCarUsecase(mockRepo)

	mockRepo.On("Delete", "1").Return(nil)

	err := usecase.Delete("1")
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Delete", "1")
}

func TestCarUsecase_List(t *testing.T) {
	mockRepo := new(MockCarRepository)
	usecase := NewCarUsecase(mockRepo)

	cars := []domain.Car{
		{ID: "1", Name: "BMW"},
		{ID: "2", Name: "Audi"},
	}
	mockRepo.On("List", 1, 10).Return(cars, nil)

	result, err := usecase.List(1, 10)
	assert.NoError(t, err)
	assert.Equal(t, cars, result)
}
