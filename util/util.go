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
func InitDB() *Database {
	cfg := config.GetConfig()
	// Если DATABASE_SSLMODE не задан, используем "disable"
	sslMode := cfg.GetDatabaseSSLMode()
	if sslMode == "" {
		sslMode = "disable"
	}

	// Формируем строку подключения для MS SQL.
	// Формат строки подключения:
	// sqlserver://username:password@host:port?database=dbname&encrypt=disable
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=%s",
		cfg.GetDatabaseUser(),
		cfg.GetDatabasePassword(),
		cfg.GetDatabaseHost(),
		cfg.GetDatabasePort(),
		cfg.GetDatabaseName(),
		sslMode,
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Connection to the database successful!")
	return &Database{DB: db}
}
