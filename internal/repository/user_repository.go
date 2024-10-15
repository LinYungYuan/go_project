package repository

import (
	"database/sql"
	"go_project/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	query := `INSERT INTO users (username, email, password, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := r.db.QueryRow(query, user.Username, user.Email, user.Password,
		user.CreatedAt, user.UpdatedAt).Scan(&user.ID)

	return err
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	query := `SELECT id, username, email, password, created_at, updated_at 
              FROM users WHERE username = $1`

	var user model.User
	err := r.db.QueryRow(query, username).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password,
		&user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
