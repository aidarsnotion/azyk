package util

import (
	"azyk/internal/domain/models"
	"gorm.io/gorm"
	"log"
)

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
		log.Fatalf("Ошибка при миграции моделей: %v", err)
	}
	log.Println("Миграции успешно выполнены.")
}
