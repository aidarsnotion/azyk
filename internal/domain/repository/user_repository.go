package repository

import (
	"azyk/internal/domain/models"
	"azyk/util/logger"
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

func (r *userRepository) CreateUser(user *models.User) error {
	if err := r.db.Create(user).Error; err != nil {
		logger.Log.WithError(err).WithField("email", user.Email).Error("Ошибка создания пользователя")
		return err
	}
	logger.Log.WithField("userID", user.ID).Info("Пользователь успешно создан")
	return nil
}

func (r *userRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Log.WithField("userID", id).Warn("Пользователь не найден")
			return nil, errors.New("user not found")
		}
		logger.Log.WithError(err).WithField("userID", id).Error("Ошибка получения пользователя по ID")
		return nil, err
	}
	logger.Log.WithField("userID", id).Info("Пользователь успешно получен")
	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Log.WithField("email", email).Warn("Пользователь не найден по email")
			return nil, errors.New("user not found")
		}
		logger.Log.WithError(err).WithField("email", email).Error("Ошибка получения пользователя по email")
		return nil, err
	}
	logger.Log.WithField("email", email).Info("Пользователь успешно получен по email")
	return &user, nil
}

func (r *userRepository) UpdateUser(user *models.User) error {
	if err := r.db.Save(user).Error; err != nil {
		logger.Log.WithError(err).WithField("userID", user.ID).Error("Ошибка обновления пользователя")
		return err
	}
	logger.Log.WithField("userID", user.ID).Info("Пользователь успешно обновлен")
	return nil
}

func (r *userRepository) DeleteUser(id int) error {
	res := r.db.Delete(&models.User{}, id)
	if res.Error != nil {
		logger.Log.WithError(res.Error).WithField("userID", id).Error("Ошибка удаления пользователя")
		return res.Error
	}
	if res.RowsAffected == 0 {
		logger.Log.WithField("userID", id).Warn("Пользователь не найден для удаления")
		return errors.New("user not found")
	}
	logger.Log.WithField("userID", id).Info("Пользователь успешно удален")
	return nil
}
