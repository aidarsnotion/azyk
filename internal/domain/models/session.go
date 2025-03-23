package models

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// Сессия пользователя
type Session struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `json:"user_id"`
	DeviceID  string    `gorm:"size:255" json:"device_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// JWT Claims
type Claims struct {
	UserID int      `json:"user_id"`
	Role   UserRole `json:"role"`
	jwt.RegisteredClaims
}
