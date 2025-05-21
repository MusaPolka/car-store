package postgres

import (
	"database/sql"
	"ecommerce/car-service/internal/domain"
	"errors"
)

type carRepo struct {
	db *sql.DB
}

func NewCarRepo(db *sql.DB) domain.CarRepository {
	return &carRepo{db: db}
}

func (r *carRepo) Create(car *domain.Car) error {
	query := `
		INSERT INTO cars (id, name, description, brand_id, price, stock, year, color, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`
	_, err := r.db.Exec(query, car.ID, car.Name, car.Description, car.BrandID, car.Price, car.Stock, car.Year, car.Color)
	return err
}

func (r *carRepo) GetByID(id string) (*domain.Car, error) {
	query := `
		SELECT id, name, description, brand_id, price, stock, year, color, created_at, updated_at
		FROM cars WHERE id=$1`
	row := r.db.QueryRow(query, id)

	car := &domain.Car{}
	if err := row.Scan(&car.ID, &car.Name, &car.Description, &car.BrandID, &car.Price, &car.Stock, &car.Year, &car.Color, &car.CreatedAt, &car.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("car not found")
		}
		return nil, err
	}
	return car, nil
}

func (r *carRepo) Update(car *domain.Car) error {
	query := `
		UPDATE cars
		SET name=$1, description=$2, brand_id=$3, price=$4, stock=$5, year=$6, color=$7, updated_at=CURRENT_TIMESTAMP
		WHERE id=$8`
	result, err := r.db.Exec(query, car.Name, car.Description, car.BrandID, car.Price, car.Stock, car.Year, car.Color, car.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows updated, car might not exist")
	}
	return nil
}

func (r *carRepo) Delete(id string) error {
	query := `DELETE FROM cars WHERE id=$1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows deleted, car might not exist")
	}
	return nil
}

func (r *carRepo) List(page, limit int) ([]domain.Car, error) {
	offset := (page - 1) * limit
	query := `
		SELECT id, name, description, brand_id, price, stock, year, color, created_at, updated_at
		FROM cars ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []domain.Car
	for rows.Next() {
		var car domain.Car
		if err := rows.Scan(&car.ID, &car.Name, &car.Description, &car.BrandID, &car.Price, &car.Stock, &car.Year, &car.Color, &car.CreatedAt, &car.UpdatedAt); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}
