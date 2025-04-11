package config

import (
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

	// Загрузка строки, затем преобразование в int для MAX_LOGIN_ATTEMPTS
	maxAttemptsStr := os.Getenv("MAX_LOGIN_ATTEMPTS")
	if maxAttemptsStr == "" {
		return cfg, errors.New("MAX_LOGIN_ATTEMPTS is empty")
	}
	maxAttempts, err := strconv.Atoi(maxAttemptsStr)
	if err != nil {
		return cfg, fmt.Errorf("MAX_LOGIN_ATTEMPTS is not a valid integer: %w", err)
	}
	cfg.Authorization.MaxLoginAttempts = maxAttempts

	// Пример для LoginAttemptWindow, если задан как строка длительности (например, "15m")
	windowStr := os.Getenv("LOGIN_ATTEMPT_WINDOW")
	if windowStr == "" {
		return cfg, errors.New("LOGIN_ATTEMPT_WINDOW is empty")
	}
	window, err := time.ParseDuration(windowStr)
	if err != nil {
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
func (c *Config) GetServerPort() string {
	return c.Server.Port
}

func (c *Config) GetDatabaseHost() string {
	return c.Database.Host
}

func (c *Config) GetDatabasePort() int {
	return c.Database.Port
}

func (c *Config) GetDatabaseUser() string {
	return c.Database.User
}

func (c *Config) GetDatabasePassword() string {
	return c.Database.Password
}

func (c *Config) GetDatabaseName() string {
	return c.Database.Name
}

func (c *Config) GetDatabaseSSLMode() string {
	return c.Database.SSLMode
}

func (c *Config) GetAuthorizationMaxLoginAttempts() int {
	return c.Authorization.MaxLoginAttempts
}

func (c *Config) GetAuthorizationMaxLoginWindow() time.Duration {
	return c.Authorization.LoginAttemptWindow
}

// --- Глобальный синглтон конфигурации ---

var (
	instance ConfigProvider
	once     sync.Once
)

// InitConfig загружает конфигурацию и инициализирует глобальный экземпляр.
// Вызывайте эту функцию один раз при старте приложения.
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
// Если InitConfig не был вызван, функция паникует.
func GetConfig() ConfigProvider {
	if instance == nil {
		panic("config not initialized. Call InitConfig() first")
	}
	return instance
}
