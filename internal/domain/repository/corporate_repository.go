package repository

import (
	"azyk/internal/domain/models"
	"errors"

	"gorm.io/gorm"
)

type CorporateRepository interface {
	AddEmployee(orgID int, user *models.User) error
	GetEmployees(orgID int) ([]*models.User, error)
	GetEmployeeCount(orgID int) (int, error)
	GetMaxEmployees(orgID int) (int, error)
}

type corporateRepository struct {
	db *gorm.DB
}

func NewCorporateRepository(db *gorm.DB) CorporateRepository {
	return &corporateRepository{db: db}
}

// Добавление сотрудника
func (r *corporateRepository) AddEmployee(orgID int, user *models.User) error {
	count, _ := r.GetEmployeeCount(orgID)
	max, _ := r.GetMaxEmployees(orgID)

	if count >= max {
		return errors.New("employee limit reached")
	}

	user.OrgID = &orgID
	user.Role = models.RoleCorporate

	return r.db.Create(&user).Error
}

// Получение сотрудников организации
func (r *corporateRepository) GetEmployees(orgID int) ([]*models.User, error) {
	var employees []*models.User
	err := r.db.Where("org_id = ?", orgID).Find(&employees).Error
	return employees, err
}

// Получение текущего количества сотрудников в организации
func (r *corporateRepository) GetEmployeeCount(orgID int) (int, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("org_id = ?", orgID).Count(&count).Error
	return int(count), err
}

// Получение максимального лимита сотрудников организации
func (r *corporateRepository) GetMaxEmployees(orgID int) (int, error) {
	var maxEmployees int
	err := r.db.Table("organizations").Select("max_employees").Where("id = ?", orgID).Scan(&maxEmployees).Error
	if err != nil {
		return 0, err
	}
	if maxEmployees == 0 {
		return 0, errors.New("organization not found or max_employees not set")
	}
	return maxEmployees, nil
}
