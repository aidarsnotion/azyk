package util

import (
	"azyk/internal/domain/models"
	"azyk/util/logger"
	"gorm.io/gorm"
	"os"
)

// RunMigrations выполняет миграцию моделей в базе данных.
func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Organization{},
		&models.OrganizationTranslation{},
		&models.User{},
		&models.Session{},
		&models.Product{},
		&models.ProductTranslation{},
		&models.Category{},
		&models.CategoryTranslation{},
		&models.AminoAcidComposition{},
		&models.AminoAcidTranslation{},
		&models.ChemicalComposition{},
		&models.ChemicalTranslation{},
		&models.MineralComposition{},
		&models.FattyAcidComposition{},
		&models.FattyType{},
		&models.FattyTypeTranslation{},
		&models.VitaminComposition{},
		&models.VitaminTranslation{},
		&models.Region{},
		&models.RegionTranslation{},
		&models.ResearchProject{},
		&models.ResearchTranslation{},
		&models.UnitModel{},
		&models.UnitTranslation{},
	)

	if err != nil {
		logger.Log.WithField("error", err.Error()).Fatal("Ошибка при миграции моделей")
		os.Exit(1) // важно выйти после fatal-ошибки
	}

	logger.Log.Info("automatic migration has been successfully completed.")
}
