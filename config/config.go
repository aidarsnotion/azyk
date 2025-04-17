package config

import (
	"azyk/util/logger"
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
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

	Authorization struct {
		MaxLoginAttempts   int
		LoginAttemptWindow time.Duration
	}
}

// LoadConfig загружает конфигурацию напрямую из переменных окружения.
func LoadConfig() (Config, error) {
	var cfg Config

	// Server
	cfg.Server.Port = os.Getenv("SERVER_PORT")
	if cfg.Server.Port == "" {
		logger.Log.WithFields(map[string]interface{}{
			"field": "SERVER_PORT",
		}).Error("Missing environment variable")
		return cfg, errors.New("SERVER_PORT is empty")
	}

	// Database
	cfg.Database.Host = os.Getenv("DATABASE_HOST")
	if cfg.Database.Host == "" {
		logger.Log.WithFields(map[string]interface{}{
			"field": "DATABASE_HOST",
		}).Error("Missing environment variable")
		return cfg, errors.New("DATABASE_HOST is empty")
	}

	dbPortStr := os.Getenv("DATABASE_PORT")
	if dbPortStr == "" {
		logger.Log.WithFields(map[string]interface{}{
			"field": "DATABASE_PORT",
		}).Error("Missing environment variable")
		return cfg, errors.New("DATABASE_PORT is empty")
	}
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"field": "DATABASE_PORT",
			"error": err.Error(),
		}).Error("Failed to parse DATABASE_PORT")
		return cfg, fmt.Errorf("DATABASE_PORT is not a valid integer: %w", err)
	}
	cfg.Database.Port = dbPort

	cfg.Database.User = os.Getenv("DATABASE_USER")
	if cfg.Database.User == "" {
		logger.Log.WithFields(map[string]interface{}{
			"field": "DATABASE_USER",
		}).Error("Missing environment variable")
		return cfg, errors.New("DATABASE_USER is empty")
	}

	cfg.Database.Password = os.Getenv("DATABASE_PASSWORD")
	if cfg.Database.Password == "" {
		logger.Log.WithFields(map[string]interface{}{
			"field": "DATABASE_PASSWORD",
		}).Error("Missing environment variable")
		return cfg, errors.New("DATABASE_PASSWORD is empty")
	}

	cfg.Database.Name = os.Getenv("DATABASE_NAME")
	if cfg.Database.Name == "" {
		logger.Log.WithFields(map[string]interface{}{
			"field": "DATABASE_NAME",
		}).Error("Missing environment variable")
		return cfg, errors.New("DATABASE_NAME is empty")
	}

	cfg.Database.SSLMode = os.Getenv("DATABASE_SSLMODE")

	// Authorization
	maxAttemptsStr := os.Getenv("MAX_LOGIN_ATTEMPTS")
	if maxAttemptsStr == "" {
		logger.Log.WithFields(map[string]interface{}{
			"field": "MAX_LOGIN_ATTEMPTS",
		}).Error("Missing environment variable")
		return cfg, errors.New("MAX_LOGIN_ATTEMPTS is empty")
	}
	maxAttempts, err := strconv.Atoi(maxAttemptsStr)
	if err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"field": "MAX_LOGIN_ATTEMPTS",
			"error": err.Error(),
		}).Error("Failed to parse MAX_LOGIN_ATTEMPTS")
		return cfg, fmt.Errorf("MAX_LOGIN_ATTEMPTS is not a valid integer: %w", err)
	}
	cfg.Authorization.MaxLoginAttempts = maxAttempts

	windowStr := os.Getenv("LOGIN_ATTEMPT_WINDOW")
	if windowStr == "" {
		logger.Log.WithFields(map[string]interface{}{
			"field": "LOGIN_ATTEMPT_WINDOW",
		}).Error("Missing environment variable")
		return cfg, errors.New("LOGIN_ATTEMPT_WINDOW is empty")
	}
	window, err := time.ParseDuration(windowStr)
	if err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"field": "LOGIN_ATTEMPT_WINDOW",
			"error": err.Error(),
		}).Error("Failed to parse LOGIN_ATTEMPT_WINDOW")
		return cfg, fmt.Errorf("LOGIN_ATTEMPT_WINDOW is not a valid duration: %w", err)
	}
	cfg.Authorization.LoginAttemptWindow = window

	return cfg, nil
}

// ConfigProvider — интерфейс для доступа к конфигурации.
type ConfigProvider interface {
	GetServerPort() string
	GetDatabaseHost() string
	GetDatabasePort() int
	GetDatabaseUser() string
	GetDatabasePassword() string
	GetDatabaseName() string
	GetDatabaseSSLMode() string
	GetAuthorizationMaxLoginAttempts() int
	GetAuthorizationMaxLoginWindow() time.Duration
}

// Реализуем методы интерфейса для структуры Config.
func (c *Config) GetServerPort() string                 { return c.Server.Port }
func (c *Config) GetDatabaseHost() string               { return c.Database.Host }
func (c *Config) GetDatabasePort() int                  { return c.Database.Port }
func (c *Config) GetDatabaseUser() string               { return c.Database.User }
func (c *Config) GetDatabasePassword() string           { return c.Database.Password }
func (c *Config) GetDatabaseName() string               { return c.Database.Name }
func (c *Config) GetDatabaseSSLMode() string            { return c.Database.SSLMode }
func (c *Config) GetAuthorizationMaxLoginAttempts() int { return c.Authorization.MaxLoginAttempts }
func (c *Config) GetAuthorizationMaxLoginWindow() time.Duration {
	return c.Authorization.LoginAttemptWindow
}

// --- Глобальный синглтон конфигурации ---
var (
	instance ConfigProvider
	once     sync.Once
)

// InitConfig загружает конфигурацию и инициализирует глобальный экземпляр.
func InitConfig() error {
	var err error
	once.Do(func() {
		cfg, loadErr := LoadConfig()
		if loadErr != nil {
			err = loadErr
			return
		}
		instance = &cfg
	})
	return err
}

// GetConfig возвращает глобальный экземпляр конфигурации.
func GetConfig() ConfigProvider {
	if instance == nil {
		panic("config not initialized. Call InitConfig() first")
	}
	return instance
}
