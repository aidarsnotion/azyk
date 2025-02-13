package repository

import (
	"azyk/internal/domain/models"
	"database/sql"
	"errors"
)

type CorporateRepository interface {
	AddEmployee(orgID int, user *models.User) error
	GetEmployees(orgID int) ([]*models.User, error)
	GetEmployeeCount(orgID int) (int, error)
	GetMaxEmployees(orgID int) (int, error)
}

type corporateRepository struct {
	db *sql.DB
}

func NewCorporateRepository(db *sql.DB) CorporateRepository {
	return &corporateRepository{db: db}
}

// 🔹 Добавление сотрудника
func (r *corporateRepository) AddEmployee(orgID int, user *models.User) error {
	count, _ := r.GetEmployeeCount(orgID)
	max, _ := r.GetMaxEmployees(orgID)

	if count >= max {
		return errors.New("employee limit reached")
	}

	_, err := r.db.Exec("INSERT INTO users (name, email, role, org_id) VALUES (?, ?, ?, ?)",
		user.Name, user.Email, models.RoleCorporate, orgID)
	return err
}

// 🔹 Получение сотрудников организации
func (r *corporateRepository) GetEmployees(orgID int) ([]*models.User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users WHERE org_id = ?", orgID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*models.User
	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Name, &user.Email)
		employees = append(employees, &user)
	}
	return employees, nil
}

// 🔹 Получение текущего количества сотрудников
func (r *corporateRepository) GetEmployeeCount(orgID int) (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM users WHERE org_id = ?", orgID).Scan(&count)
	return count, err
}

// 🔹 Получение максимального лимита сотрудников
func (r *corporateRepository) GetMaxEmployees(orgID int) (int, error) {
	var max int
	err := r.db.QueryRow("SELECT max_employees FROM organizations WHERE id = ?", orgID).Scan(&max)
	return max, err
}
