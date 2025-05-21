package domain

import "time"

type Order struct {
	ID         string      `json:"id"`
	UserID     string      `json:"user_id"`
	TotalPrice float64     `json:"total_price"`
	Status     string      `json:"status"` // pending, paid, canceled, etc.
	Items      []OrderItem `json:"items"`
	CreatedAt  time.Time   `json:"created_at,omitempty"`
	UpdatedAt  time.Time   `json:"updated_at,omitempty"`
}

type OrderItem struct {
	ID        string  `json:"id"`
	OrderID   string  `json:"order_id"`
	CarID     string  `json:"car_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	Subtotal  float64 `json:"subtotal"`
}

type OrderRepository interface {
	Create(order *Order) error
	GetByID(id string) (*Order, error)
	Update(order *Order) error
	Delete(id string) error
	List() ([]Order, error)
	ListByUser(userID string) ([]Order, error)
}
