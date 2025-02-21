package repository

import (
	"azyk/internal/domain/models"

	"gorm.io/gorm"
)

type SessionRepository interface {
	CreateSession(userID int, deviceID string) error
	GetActiveSession(userID int) (*models.Session, error)
	DeleteSession(userID int) error
}

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db: db}
}

// CreateSession создает новую сессию для пользователя.
// Если активная сессия уже существует, она удаляется (при условии, что пользователь может иметь только одно устройство).
func (r *sessionRepository) CreateSession(userID int, deviceID string) error {
	var existingSession models.Session
	err := r.db.Where("user_id = ?", userID).First(&existingSession).Error
	if err == nil {
		// Если активная сессия найдена, удаляем её.
		if err := r.db.Where("user_id = ?", userID).Delete(&models.Session{}).Error; err != nil {
			return err
		}
	} else if err != gorm.ErrRecordNotFound {
		// Если произошла другая ошибка, возвращаем её.
		return err
	}

	// Создаем новую сессию.
	session := models.Session{
		UserID:   userID,
		DeviceID: deviceID,
	}
	return r.db.Create(&session).Error
}

// GetActiveSession возвращает активную сессию пользователя или nil, если сессия не найдена.
func (r *sessionRepository) GetActiveSession(userID int) (*models.Session, error) {
	var session models.Session
	if err := r.db.Where("user_id = ?", userID).First(&session).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &session, nil
}

// DeleteSession удаляет сессию пользователя по его ID.
func (r *sessionRepository) DeleteSession(userID int) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Session{}).Error
}
