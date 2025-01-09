package util

import (
	"fmt"
	"log"

	"azyk/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB устанавливает соединение с PostgreSQL используя настройки из конфигурации.
func InitDB(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	return db
}
