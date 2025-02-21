package repository

import (
	"azyk/internal/domain/models"
	"azyk/util"
	"errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// CreateUser создает нового пользователя
func (r *userRepository) CreateUser(user *models.User) error {
	// Хешируем пароль перед сохранением
	pass, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = pass

	// Создаем пользователя в базе данных
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

// GetUserByID возвращает пользователя по ID
func (r *userRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail возвращает пользователя по email
func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUser обновляет данные пользователя
func (r *userRepository) UpdateUser(user *models.User) error {
	// Обновляем только указанные поля
	updates := map[string]interface{}{
		"name":       user.Name,
		"email":      user.Email,
		"updated_at": gorm.Expr("NOW()"),
	}

	// Если пароль был изменен, хешируем его
	if user.Password != "" {
		pass, err := util.HashPassword(user.Password)
		if err != nil {
			return err
		}

		user.Password = pass
		updates["password"] = user.Password
	}

	// Обновляем пользователя по ID
	if err := r.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

// DeleteUser удаляет пользователя по ID
func (r *userRepository) DeleteUser(id int) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}
	return nil
}
