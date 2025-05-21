package domain

import "time"

type Car struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	BrandID     string    `json:"brand_id"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Year        int       `json:"year"`
	Color       string    `json:"color"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type CarRepository interface {
	Create(car *Car) error
	GetByID(id string) (*Car, error)
	Update(car *Car) error
	Delete(id string) error
	List(page, limit int) ([]Car, error)
}
