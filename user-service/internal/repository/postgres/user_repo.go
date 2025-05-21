package postgres

import (
	"database/sql"
	"ecommerce/user-service/internal/domain"
	"errors"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) domain.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *domain.User) error {
	query := `
		INSERT INTO users (id, username, password, email, full_name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`
	_, err := r.db.Exec(query, user.ID, user.Username, user.Password, user.Email, user.FullName)
	return err
}

func (r *userRepo) GetByID(id string) (*domain.User, error) {
	query := `SELECT id, username, password, email, full_name, created_at, updated_at FROM users WHERE id=$1`
	row := r.db.QueryRow(query, id)

	var user domain.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.FullName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetByUsername(username string) (*domain.User, error) {
	query := `SELECT id, username, password, email, full_name, created_at, updated_at FROM users WHERE username=$1`
	row := r.db.QueryRow(query, username)

	var user domain.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.FullName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Update(user *domain.User) error {
	query := `
		UPDATE users SET username=$1, password=$2, email=$3, full_name=$4, updated_at=CURRENT_TIMESTAMP WHERE id=$5`
	result, err := r.db.Exec(query, user.Username, user.Password, user.Email, user.FullName, user.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows updated, user might not exist")
	}
	return nil
}

func (r *userRepo) Delete(id string) error {
	query := `DELETE FROM users WHERE id=$1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows deleted, user might not exist")
	}
	return nil
}

func (r *userRepo) List() ([]domain.User, error) {
	query := `SELECT id, username, password, email, full_name, created_at, updated_at FROM users ORDER BY created_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.FullName, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
