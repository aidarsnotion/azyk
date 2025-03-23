package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config — структура для хранения настроек приложения.
type Config struct {
	Server struct {
		Port string
	}

	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
		SSLMode  string
	}
}

// LoadConfig загружает конфигурацию напрямую из переменных окружения.
func LoadConfig() (Config, error) {
	var cfg Config

	// Читаем переменные для сервера
	cfg.Server.Port = os.Getenv("SERVER_PORT")
	if cfg.Server.Port == "" {
		return cfg, fmt.Errorf("SERVER_PORT is empty")
	}

	// Читаем переменные для базы данных
	cfg.Database.Host = os.Getenv("DATABASE_HOST")
	if cfg.Database.Host == "" {
		return cfg, fmt.Errorf("DATABASE_HOST is empty")
	}

	dbPortStr := os.Getenv("DATABASE_PORT")
	if dbPortStr == "" {
		return cfg, fmt.Errorf("DATABASE_PORT is empty")
	}
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return cfg, fmt.Errorf("DATABASE_PORT is not a valid integer: %w", err)
	}
	cfg.Database.Port = dbPort

	cfg.Database.User = os.Getenv("DATABASE_USER")
	if cfg.Database.User == "" {
		return cfg, fmt.Errorf("DATABASE_USER is empty")
	}

	cfg.Database.Password = os.Getenv("DATABASE_PASSWORD")
	if cfg.Database.Password == "" {
		return cfg, fmt.Errorf("DATABASE_PASSWORD is empty")
	}

	cfg.Database.Name = os.Getenv("DATABASE_NAME")
	if cfg.Database.Name == "" {
		return cfg, fmt.Errorf("DATABASE_NAME is empty")
	}

	// Для SSLMode можно оставить пустым, если не используется шифрование
	cfg.Database.SSLMode = os.Getenv("DATABASE_SSLMODE")

	return cfg, nil
}
