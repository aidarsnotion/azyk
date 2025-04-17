package util

import (
	"azyk/config"
	"azyk/util/logger"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

// InitDB устанавливает соединение с PostgreSQL и настраивает пул соединений.
func InitDB() *Database {
	cfg := config.GetConfig()
	sslMode := cfg.GetDatabaseSSLMode()
	if sslMode == "" {
		sslMode = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.GetDatabaseHost(),
		cfg.GetDatabasePort(),
		cfg.GetDatabaseUser(),
		cfg.GetDatabasePassword(),
		cfg.GetDatabaseName(),
		sslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.WithField("error", err.Error()).Fatal("Failed to connect to the database")
		os.Exit(1)
	}

	// Настройка пула соединений
	sqlDB, err := db.DB()
	if err != nil {
		logger.Log.WithField("error", err.Error()).Fatal("Failed to get database object from GORM")
		os.Exit(1)
	}

	sqlDB.SetMaxIdleConns(10)                  // Максимум 10 простаивающих подключений
	sqlDB.SetMaxOpenConns(50)                  // Максимум 50 открытых подключений к БД
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Подключение живет максимум 30 минут
	sqlDB.SetConnMaxIdleTime(10 * time.Minute) // Подключение может быть неактивным максимум 10 минут

	logger.Log.WithFields(map[string]interface{}{
		"host": cfg.GetDatabaseHost(),
		"port": cfg.GetDatabasePort(),
		"user": cfg.GetDatabaseUser(),
		"db":   cfg.GetDatabaseName(),
	}).Info("Connection to the database successful and pool configured!")

	return &Database{DB: db}
}
