package usecase

import (
	"ecommerce/car-service/internal/domain"
)

type CarUsecase interface {
	Create(car *domain.Car) error
	GetByID(id string) (*domain.Car, error)
	Update(car *domain.Car) error
	Delete(id string) error
	List(page, limit int) ([]domain.Car, error)
}

type carUsecase struct {
	repo domain.CarRepository
}

func NewCarUsecase(repo domain.CarRepository) CarUsecase {
	return &carUsecase{repo: repo}
}

func (u *carUsecase) Create(car *domain.Car) error {
	return u.repo.Create(car)
}
func (u *carUsecase) GetByID(id string) (*domain.Car, error) {
	return u.repo.GetByID(id)
}
func (u *carUsecase) Update(car *domain.Car) error {
	return u.repo.Update(car)
}
func (u *carUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}
func (u *carUsecase) List(page, limit int) ([]domain.Car, error) {
	return u.repo.List(page, limit)
}
