package repository

import (
	"time"

	"azyk/internal/domain/models"
	"azyk/util/logger"
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

	// Ищем существующую сессию по user_id.
	err := r.db.Where("user_id = ?", userID).First(&existingSession).Error
	if err == nil {
		// Если найдена активная сессия, удаляем её (если политика требует одного устройства).
		logger.Log.WithField("userID", userID).Info("Найдена активная сессия, удаляем её")
		if delErr := r.db.Where("user_id = ?", userID).Delete(&models.Session{}).Error; delErr != nil {
			logger.Log.WithError(delErr).WithField("userID", userID).Error("Ошибка при удалении активной сессии")
			return delErr
		}
	} else if err != gorm.ErrRecordNotFound {
		logger.Log.WithError(err).WithField("userID", userID).Error("Ошибка поиска активной сессии")
		return err
	}

	// Создаем новую сессию.
	session := &models.Session{
		UserID:    userID,
		DeviceID:  deviceID,
		CreatedAt: time.Now(),
	}
	if err := r.db.Create(session).Error; err != nil {
		logger.Log.WithError(err).WithField("userID", userID).Error("Ошибка создания новой сессии")
		return err
	}

	logger.Log.WithField("userID", userID).Info("Новая сессия успешно создана")
	return nil
}

func (r *sessionRepository) GetActiveSession(userID int) (*models.Session, error) {
	var session models.Session
	err := r.db.Where("user_id = ?", userID).First(&session).Error
	if err == gorm.ErrRecordNotFound {
		logger.Log.WithField("userID", userID).Warn("Активная сессия не найдена")
		return nil, nil
	} else if err != nil {
		logger.Log.WithError(err).WithField("userID", userID).Error("Ошибка получения активной сессии")
		return nil, err
	}

	logger.Log.WithField("userID", userID).Info("Активная сессия успешно получена")
	return &session, nil
}

func (r *sessionRepository) DeleteSession(userID int) error {
	if err := r.db.Where("user_id = ?", userID).Delete(&models.Session{}).Error; err != nil {
		logger.Log.WithError(err).WithField("userID", userID).Error("Ошибка удаления сессии")
		return err
	}
	logger.Log.WithField("userID", userID).Info("Сессия успешно удалена")
	return nil
}
