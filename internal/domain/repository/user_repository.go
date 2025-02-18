package repository

import (
	"azyk/internal/domain/models"
	"database/sql"
	"errors"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// Create usre
func (r *userRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (name, email, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING id"
	err := r.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	return err
}

// Get user by id
func (r *userRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return user, err
}

// get user by email
func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, name, email, created_at, updated_at FROM users WHERE email = $1"
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return user, err
}

// ubdate user
func (r *userRepository) UpdateUser(user *models.User) error {
	query := "UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3"
	res, err := r.db.Exec(query, user.Name, user.Email, user.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

// delete user
func (r *userRepository) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
