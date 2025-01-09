package main

import (
	"azyk/config"
	"azyk/util"
	"log"

	"azyk/internal/delivery/http"
	"azyk/internal/domain/models"
	"azyk/internal/repository"
	"azyk/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}

	// Подключение к базе данных
	db := util.InitDB(cfg)

	// Автоматическая миграция моделей (при необходимости)
	db.AutoMigrate(&models.Product{}, &models.ProductTranslation{} /*, остальные модели... */)

	// Инициализация слоёв: репозиторий, usecase и HTTP-обработчики
	productRepo := repository.NewProductRepository(db)
	productUC := usecase.NewProductUsecase(productRepo)

	router := gin.Default()
	http.NewProductHandler(router, productUC)

	// Запуск сервера
	router.Run(":8080")
}
