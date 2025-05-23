package usecase

import (
	"ecommerce/order-service/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderUsecase_Create(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	usecase := NewOrderUsecase(mockRepo)

	order := &domain.Order{ID: "1", UserID: "u1"}
	mockRepo.On("Create", order).Return(nil)

	err := usecase.Create(order)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Create", order)
}

func TestOrderUsecase_GetByID(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	usecase := NewOrderUsecase(mockRepo)

	order := &domain.Order{ID: "1", UserID: "u1"}
	mockRepo.On("GetByID", "1").Return(order, nil)

	result, err := usecase.GetByID("1")
	assert.NoError(t, err)
	assert.Equal(t, order, result)
}

func TestOrderUsecase_Update(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	usecase := NewOrderUsecase(mockRepo)

	order := &domain.Order{ID: "1", UserID: "u1"}
	mockRepo.On("Update", order).Return(nil)

	err := usecase.Update(order)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Update", order)
}

func TestOrderUsecase_Delete(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	usecase := NewOrderUsecase(mockRepo)

	mockRepo.On("Delete", "1").Return(nil)

	err := usecase.Delete("1")
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Delete", "1")
}

func TestOrderUsecase_List(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	usecase := NewOrderUsecase(mockRepo)

	orders := []domain.Order{{ID: "1", UserID: "u1"}, {ID: "2", UserID: "u2"}}
	mockRepo.On("List").Return(orders, nil)

	result, err := usecase.List()
	assert.NoError(t, err)
	assert.Equal(t, orders, result)
}

func TestOrderUsecase_ListByUser(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	usecase := NewOrderUsecase(mockRepo)

	orders := []domain.Order{{ID: "1", UserID: "u1"}}
	mockRepo.On("ListByUser", "u1").Return(orders, nil)

	result, err := usecase.ListByUser("u1")
	assert.NoError(t, err)
	assert.Equal(t, orders, result)
}
