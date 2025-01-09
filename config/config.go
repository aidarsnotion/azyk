package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config — структура для хранения настроек приложения.
type Config struct {
	Server struct {
		Port string `mapstructure:"SERVER_PORT"`
	}

	Database struct {
		Host     string `mapstructure:"DATABASE_HOST"`
		Port     int    `mapstructure:"DATABASE_PORT"`
		User     string `mapstructure:"DATABASE_USER"`
		Password string `mapstructure:"DATABASE_PASSWORD"`
		Name     string `mapstructure:"DATABASE_NAME"`
		SSLMode  string `mapstructure:"DATABASE_SSLMODE"`
	}
}

// LoadConfig загружает конфигурацию из .env и переменных окружения.
func LoadConfig() (config Config, err error) {
	// Загрузка переменных из .env файла
	if err = godotenv.Load(); err != nil {
		log.Println("Не удалось загрузить .env файл, продолжаем с существующими переменными окружения")
	}

	// Говорим Viper автоматически считывать все переменные окружения
	viper.AutomaticEnv()

	// Распаковка настроек в структуру Config
	err = viper.Unmarshal(&config)
	return
}
