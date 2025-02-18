package repository

import (
	"azyk/internal/domain/models"
	"database/sql"
)

type SessionRepository interface {
	CreateSession(userID int, deviceID string) error
	GetActiveSession(userID int) (*models.Session, error)
	DeleteSession(userID int) error
}

type sessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) SessionRepository {
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

func (r *sessionRepository) GetActiveSession(userID int) (*models.Session, error) {
	var session models.Session
	err := r.db.QueryRow("SELECT id, user_id, device_id, created_at FROM sessions WHERE user_id = ?",
		userID).Scan(&session.ID, &session.UserId, &session.UserId, &session.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &session, err
}

// Удаление сессии пользователя
func (r *sessionRepository) DeleteSession(userID int) error {
	_, err := r.db.Exec("DELETE FROM session WHERE user_id = ?", userID)
	return err
}
