package postgres

import (
	"database/sql"
	"ecommerce/car-service/internal/domain"
	"errors"
)

type carBrandRepo struct {
	db *sql.DB
}

func NewCarBrandRepo(db *sql.DB) domain.CarBrandRepository {
	return &carBrandRepo{db: db}
}

func (r *carBrandRepo) Create(brand *domain.CarBrand) error {
	query := `
	INSERT INTO carbrands (name, description) 
	VALUES ($1, $2)
	RETURNING id`
	err := r.db.QueryRow(query, brand.Name, brand.Description).Scan(&brand.ID)

	return err
}

func (r *carBrandRepo) GetByID(id string) (*domain.CarBrand, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM carbrands WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var brand domain.CarBrand
	err := row.Scan(&brand.ID, &brand.Name, &brand.Description, &brand.CreatedAt, &brand.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("car brand not found")
		}
		return nil, err
	}
	return &brand, nil
}

func (r *carBrandRepo) Update(brand *domain.CarBrand) error {
	query := `UPDATE carbrands SET name=$1, description=$2, updated_at=CURRENT_TIMESTAMP WHERE id=$3`
	result, err := r.db.Exec(query, brand.Name, brand.Description, brand.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows updated, car brand may not exist")
	}
	return nil
}

func (r *carBrandRepo) Delete(id string) error {
	query := `DELETE FROM carbrands WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows deleted, car brand may not exist")
	}
	return nil
}

func (r *carBrandRepo) List() ([]domain.CarBrand, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM carbrands ORDER BY created_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brands []domain.CarBrand
	for rows.Next() {
		var brand domain.CarBrand
		if err := rows.Scan(&brand.ID, &brand.Name, &brand.Description, &brand.CreatedAt, &brand.UpdatedAt); err != nil {
			return nil, err
		}
		brands = append(brands, brand)
	}
	return brands, nil
}
