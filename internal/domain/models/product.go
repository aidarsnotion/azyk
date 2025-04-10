package models

import (
	"time"
)

// Product представляет продукт.
type Product struct {
	ID                     int32                  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                   string                 `json:"name"`
	RegionsID              int32                  `json:"regions_id"`
	Region                 Region                 `gorm:"foreignKey:RegionsID" json:"region"`
	National               bool                   `json:"national"`
	CategoriesID           int32                  `json:"categories_id"`
	Categories             Category               `gorm:"foreignKey:CategoriesID" json:"categories"`
	Description            string                 `gorm:"type:text" json:"description"`
	ResearchDate           time.Time              `gorm:"column:research_date" json:"research_date"`
	CreatedAt              time.Time              `json:"created_at"`
	UpdatedAt              time.Time              `json:"updated_at"`
	ProductResearchSources []ResearchProject      `gorm:"many2many:product_research_projects;" json:"product_research_sources"`
	AminoAcidCompositions  []AminoAcidComposition `gorm:"foreignKey:ProductID" json:"amino_acid_compositions"`
	MineralCompositions    []MineralComposition   `gorm:"foreignKey:ProductID" json:"mineral_compositions"`
	ChemicalCompositions   []ChemicalComposition  `gorm:"foreignKey:ProductID" json:"chemical_compositions"`
	FattyAcidCompositions  []FattyAcidComposition `gorm:"foreignKey:ProductID" json:"fatty_acid_compositions"`
	VitaminCompositions    []VitaminComposition   `gorm:"foreignKey:ProductID" json:"vitamin_compositions"`
}

// ProductTranslation используется для перевода информации о продукте.
type ProductTranslation struct {
	ID          int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID   int32  `json:"product_id"`
	Language    string `json:"language"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
}
