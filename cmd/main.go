package main

import (
	"azyk/config"
	"azyk/internal/usecase/auth"
	"azyk/internal/usecase/throttling"
	"azyk/util"
	"azyk/util/logger"
	"github.com/sirupsen/logrus"
	"os"

	httpDelivery "azyk/internal/delivery/http"
	"azyk/internal/domain/repository"
	"azyk/internal/usecase"
)

func initializeHandlers(db *util.Database) (*httpDelivery.AminoAcidCompositionHandler,
	*httpDelivery.ChemicalCompositionHandler,
	*httpDelivery.FattyAcidCompositionHandler,
	*httpDelivery.MineralCompositionHandler,
	*httpDelivery.ProductHandler,
	*httpDelivery.VitaminCompositionHandler,
	*httpDelivery.CategoryHandler,
	*httpDelivery.RegionHandler) {

	// Репозитории
	aminoRepo := repository.NewAminoAcidCompositionRepository(db.DB)
	chemicalRepo := repository.NewChemicalCompositionRepository(db.DB)
	fattyRepo := repository.NewFattyAcidCompositionRepository(db.DB)
	mineralRepo := repository.NewMineralCompositionRepository(db.DB)
	productRepo := repository.NewProductRepository(db.DB)
	vitaminRepo := repository.NewVitaminCompositionRepository(db.DB)
	categoryRepo := repository.NewCategoryRepository(db.DB)
	regionRepo := repository.NewRegionRepository(db.DB)

	// UseCases (сервисы бизнес-логики)
	aminoSvc := usecase.NewAminoAcidCompositionService(aminoRepo)
	chemicalSvc := usecase.NewChemicalCompositionService(chemicalRepo)
	fattySvc := usecase.NewFattyAcidCompositionService(fattyRepo)
	mineralSvc := usecase.NewMineralCompositionService(mineralRepo)
	productSvc := usecase.NewProductService(productRepo)
	vitaminSvc := usecase.NewVitaminCompositionService(vitaminRepo)
	categorySvc := usecase.NewCategoryUsecase(categoryRepo)
	regionSvc := usecase.NewRegionUsecase(regionRepo)

	// Хендлеры (обработчики HTTP)
	aminoHandler := httpDelivery.NewAminoAcidCompositionHandler(aminoSvc)
	chemicalHandler := httpDelivery.NewChemicalCompositionHandler(chemicalSvc)
	fattyHandler := httpDelivery.NewFattyAcidCompositionHandler(fattySvc)
	mineralHandler := httpDelivery.NewMineralCompositionHandler(mineralSvc)
	productHandler := httpDelivery.NewProductHandler(productSvc)
	vitaminHandler := httpDelivery.NewVitaminCompositionHandler(vitaminSvc)
	categoryHandler := httpDelivery.NewCategoryHandler(categorySvc)
	regionHandler := httpDelivery.NewRegionHandler(regionSvc)

	return aminoHandler, chemicalHandler, fattyHandler, mineralHandler, productHandler, vitaminHandler, categoryHandler, regionHandler
}

func main() {
	logger.Log.WithFields(logrus.Fields{
		"info": "azyk - 1.0.0",
	}).Info("Starting application")

	// Загрузка конфигурации
	if err := config.InitConfig(); err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"error": err.Error(),
		}).Fatal("Failed to load configuration")
		os.Exit(1)
	}

	cfg := config.GetConfig()
	// Создаём throttler на основе in-memory реализации.
	throttler := throttling.NewInMemoryThrottler(cfg.GetAuthorizationMaxLoginAttempts(), cfg.GetAuthorizationMaxLoginWindow())

	// Подключение к базе данных
	db := util.InitDB()

	//Выполнеие миграций
	util.RunMigrations(db.DB)

	// Инициализация слоёв: репозиторий, usecase и HTTP-обработчики
	userRepo := repository.NewUserRepository(db.DB)
	sessionRepo := repository.NewSessionRepository(db.DB)

	// Создаём use-case: бизнес-логику для работы с пользователями и аутентификацией
	userUC := usecase.NewUserUsecase(userRepo, sessionRepo)
	authUC := auth.NewAuthUseCase(userRepo, sessionRepo, throttler)

	// Инициализация дополнительных usecase и handler'ов
	aminoHandler, chemicalHandler, fattyHandler, mineralHandler, productHandler, vitaminHandler, categoryHandler, regionHandler := initializeHandlers(db)

	// Создаём роутер с зарегистрированными маршрутами (HTTP-обработчики)
	router := httpDelivery.NewRouter(userUC, authUC, aminoHandler, chemicalHandler, mineralHandler, fattyHandler, vitaminHandler, productHandler, categoryHandler, regionHandler)

	// Запуск сервера
	httpDelivery.StartServer(router, ":8080")
}
