package usecase

import (
	"ecommerce/user-service/internal/domain"
)

type UserUsecase interface {
	Create(user *domain.User) error
	GetByID(id string) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id string) error
	List() ([]domain.User, error)
}

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) Create(user *domain.User) error {
	return u.repo.Create(user)
}
func (u *userUsecase) GetByID(id string) (*domain.User, error) {
	return u.repo.GetByID(id)
}
func (u *userUsecase) GetByUsername(username string) (*domain.User, error) {
	return u.repo.GetByUsername(username)
}
func (u *userUsecase) Update(user *domain.User) error {
	return u.repo.Update(user)
}
func (u *userUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}
func (u *userUsecase) List() ([]domain.User, error) {
	return u.repo.List()
}
