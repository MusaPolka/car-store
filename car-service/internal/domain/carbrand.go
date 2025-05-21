package domain

import "time"

type CarBrand struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type CarBrandRepository interface {
	Create(brand *CarBrand) error
	GetByID(id string) (*CarBrand, error)
	Update(brand *CarBrand) error
	Delete(id string) error
	List() ([]CarBrand, error)
}
