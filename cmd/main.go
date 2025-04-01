package main

import (
	"azyk/config"
	"azyk/internal/usecase/auth"
	"azyk/internal/usecase/throttling"
	"azyk/util"
	"log"

	"azyk/internal/delivery/http"
	"azyk/internal/domain/models"
	"azyk/internal/domain/repository"
	"azyk/internal/usecase"
)

func main() {
	// Загрузка конфигурации
	if err := config.InitConfig(); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	cfg := config.GetConfig()
	// Создаём throttler на основе in-memory реализации.
	throttler := throttling.NewInMemoryThrottler(cfg.GetAuthorizationMaxLoginAttempts(), cfg.GetAuthorizationMaxLoginWindow())

	// Подключение к базе данных
	db := util.InitDB()

	// Автоматическая миграция моделей (при необходимости)
	db.DB.AutoMigrate(&models.User{}, &models.Session{} /*, остальные модели... */)

	// Инициализация слоёв: репозиторий, usecase и HTTP-обработчики
	userRepo := repository.NewUserRepository(db.DB)
	sessionRepo := repository.NewSessionRepository(db.DB)

	// Создаём use-case: бизнес-логику для работы с пользователями и аутентификацией
	userUC := usecase.NewUserUsecase(userRepo, sessionRepo)
	authUC := auth.NewAuthUseCase(userRepo, sessionRepo, throttler)

	// Инициализация HTTP-обработчика для рецептуры
	// Создаём роутер с зарегистрированными маршрутами (HTTP-обработчики)
	router := http.NewRouter(userUC, authUC)

	// Запуск сервера
	// Запускаем HTTP-сервер
	http.StartServer(router, ":8080")
}
