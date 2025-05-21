package postgres

import (
	"database/sql"
	"ecommerce/order-service/internal/domain"
	"errors"
)

type orderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) domain.OrderRepository {
	return &orderRepo{db: db}
}

func (r *orderRepo) Create(order *domain.Order) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// Insert order
	orderQuery := `
		INSERT INTO orders (id, user_id, total_price, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`
	_, err = tx.Exec(orderQuery, order.ID, order.UserID, order.TotalPrice, order.Status)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert order items
	itemQuery := `
		INSERT INTO order_items (id, order_id, car_id, quantity, unit_price, subtotal)
		VALUES ($1, $2, $3, $4, $5, $6)`
	for _, item := range order.Items {
		_, err := tx.Exec(itemQuery, item.ID, order.ID, item.CarID, item.Quantity, item.UnitPrice, item.Subtotal)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *orderRepo) GetByID(id string) (*domain.Order, error) {
	// Get order
	orderQuery := `SELECT id, user_id, total_price, status, created_at, updated_at FROM orders WHERE id=$1`
	row := r.db.QueryRow(orderQuery, id)

	var order domain.Order
	err := row.Scan(&order.ID, &order.UserID, &order.TotalPrice, &order.Status, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	// Get order items
	itemQuery := `SELECT id, order_id, car_id, quantity, unit_price, subtotal FROM order_items WHERE order_id=$1`
	rows, err := r.db.Query(itemQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.OrderItem
	for rows.Next() {
		var item domain.OrderItem
		if err := rows.Scan(&item.ID, &item.OrderID, &item.CarID, &item.Quantity, &item.UnitPrice, &item.Subtotal); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	order.Items = items

	return &order, nil
}

func (r *orderRepo) Update(order *domain.Order) error {
	// For simplicity, only update status and total_price
	query := `UPDATE orders SET total_price=$1, status=$2, updated_at=CURRENT_TIMESTAMP WHERE id=$3`
	result, err := r.db.Exec(query, order.TotalPrice, order.Status, order.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows updated, order might not exist")
	}
	return nil
}

func (r *orderRepo) Delete(id string) error {
	// Remove order items first for FK constraint
	_, err := r.db.Exec(`DELETE FROM order_items WHERE order_id=$1`, id)
	if err != nil {
		return err
	}
	result, err := r.db.Exec(`DELETE FROM orders WHERE id=$1`, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows deleted, order might not exist")
	}
	return nil
}

func (r *orderRepo) List() ([]domain.Order, error) {
	query := `SELECT id, user_id, total_price, status, created_at, updated_at FROM orders ORDER BY created_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.TotalPrice, &order.Status, &order.CreatedAt, &order.UpdatedAt); err != nil {
			return nil, err
		}
		// You may want to fetch items here as well if needed
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *orderRepo) ListByUser(userID string) ([]domain.Order, error) {
	query := `SELECT id, user_id, total_price, status, created_at, updated_at FROM orders WHERE user_id=$1 ORDER BY created_at DESC`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.TotalPrice, &order.Status, &order.CreatedAt, &order.UpdatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
