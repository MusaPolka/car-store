package usecase

import (
	"ecommerce/order-service/internal/domain"
)

type OrderUsecase interface {
	Create(order *domain.Order) error
	GetByID(id string) (*domain.Order, error)
	Update(order *domain.Order) error
	Delete(id string) error
	List() ([]domain.Order, error)
	ListByUser(userID string) ([]domain.Order, error)
}

type orderUsecase struct {
	repo domain.OrderRepository
}

func NewOrderUsecase(repo domain.OrderRepository) OrderUsecase {
	return &orderUsecase{repo: repo}
}

func (u *orderUsecase) Create(order *domain.Order) error {
	return u.repo.Create(order)
}
func (u *orderUsecase) GetByID(id string) (*domain.Order, error) {
	return u.repo.GetByID(id)
}
func (u *orderUsecase) Update(order *domain.Order) error {
	return u.repo.Update(order)
}
func (u *orderUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}
func (u *orderUsecase) List() ([]domain.Order, error) {
	return u.repo.List()
}
func (u *orderUsecase) ListByUser(userID string) ([]domain.Order, error) {
	return u.repo.ListByUser(userID)
}
