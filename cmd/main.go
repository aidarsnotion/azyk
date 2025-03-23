package main

import (
	"azyk/config"
	"azyk/util"
	"log"

	"azyk/internal/delivery/http"
	"azyk/internal/domain/models"
	"azyk/internal/domain/repository"
	"azyk/internal/usecase"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}
	log.Printf("Конфигурация: %+v", cfg)

	// Подключение к базе данных
	db := util.InitDB(cfg)

	// Автоматическая миграция моделей (при необходимости)
	db.DB.AutoMigrate(&models.User{}, &models.Session{} /*, остальные модели... */)

	// Инициализация слоёв: репозиторий, usecase и HTTP-обработчики
	userRepo := repository.NewUserRepository(db.DB)
	sessionRepo := repository.NewSessionRepository(db.DB)

	// Создаём use-case: бизнес-логику для работы с пользователями и аутентификацией
	userUC := usecase.NewUserUsecase(userRepo, sessionRepo)
	authUC := usecase.NewAuthUseCase(userRepo, sessionRepo)

	// Инициализация HTTP-обработчика для рецептуры
	// Создаём роутер с зарегистрированными маршрутами (HTTP-обработчики)
	router := http.NewRouter(userUC, authUC)

	// Запуск сервера
	// Запускаем HTTP-сервер
	log.Println("Сервер запущен на :8080")
	http.StartServer(router, "8080")
}
