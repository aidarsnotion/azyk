package util

import (
	"fmt"
	"log"

	"azyk/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

// InitDB устанавливает соединение с PostgreSQL.
func InitDB() *Database {
	cfg := config.GetConfig()
	sslMode := cfg.GetDatabaseSSLMode()
	if sslMode == "" {
		sslMode = "disable"
	}

	// Формируем строку подключения для PostgreSQL.
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.GetDatabaseHost(),
		cfg.GetDatabasePort(),
		cfg.GetDatabaseUser(),
		cfg.GetDatabasePassword(),
		cfg.GetDatabaseName(),
		sslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Connection to the database successful!")
	return &Database{DB: db}
}
