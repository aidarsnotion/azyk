package repository

import (
	"time"

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

// CreateSession создает новую сессию; если уже существует активная сессия, удаляет её.
func (r *sessionRepository) CreateSession(userID int, deviceID string) error {
	var existingSession models.Session
	// Ищем существующую сессию по user_id
	err := r.db.Where("user_id = ?", userID).First(&existingSession).Error
	if err == nil {
		// Если найдена активная сессия, удаляем её (если политика требует одного устройства)
		if err := r.db.Where("user_id = ?", userID).Delete(&models.Session{}).Error; err != nil {
			return err
		}
	}
	// Создаем новую сессию
	session := &models.Session{
		UserId:    userID,
		DeviceId:  deviceID,
		CreatedAt: time.Now(),
	}
	return r.db.Create(session).Error
}

func (r *sessionRepository) GetActiveSession(userID int) (*models.Session, error) {
	var session models.Session
	err := r.db.Where("user_id = ?", userID).First(&session).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &session, err
}

func (r *sessionRepository) DeleteSession(userID int) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Session{}).Error
}
