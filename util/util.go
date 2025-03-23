package util

import (
	"fmt"
	"log"

	"azyk/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

// InitDB устанавливает соединение с MS SQL.
func InitDB(cfg config.Config) *Database {
	// Если DATABASE_SSLMODE не задан, используем "disable"
	sslMode := cfg.Database.SSLMode
	if sslMode == "" {
		sslMode = "disable"
	}

	// Формируем строку подключения для MS SQL.
	// Формат строки подключения:
	// sqlserver://username:password@host:port?database=dbname&encrypt=disable
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		sslMode,
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	log.Println("Подключение к базе данных успешно!")
	return &Database{DB: db}
}
