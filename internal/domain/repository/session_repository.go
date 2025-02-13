package repository

import (
	"azyk/internal/domain/models"
	"database/sql"
)

type SessionRepository interface {
	CreateSession(userID int, deviceID string) (string, error)
	GetActiveSession(userID int) (*models.Session, error)
	DeleteSession(userID int) error
}

type sessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) sessionRepository {
	return &sessionRepository{db: db}
}

// Создание новой сессии (если роль требует 1 устройства — удаляем старую)
func (r *sessionRepository) CreateSession(userID int, deviceID string) error {
	//Проверяем активную сессию
	var existinSession models.Session
	err := r.db.QueryRow("SELECT id FROM session WHERE user_id = ?", userID).Scan(&existinSession.ID)
	if err == nil {
		//Если есть активная сессия, удалим её (если роль требует 1 устройство)
		_, err = r.db.Exec("DELETE FROM session WHERE user_id = ?", userID)
		if err != nil {
			return err
		}
	}

	_, err = r.db.Exec("INSERT INTO session (user_id, device_id, created_at) VALUES (?, ?, NOW())", userID, deviceID)
	return err
}

// Удаление сессии пользователя
func (r *sessionRepository) DeleteSession(userID int) error {
	_, err := r.db.Exec("DELETE FROM session WHERE user_id = ?", userID)
	return err
}
