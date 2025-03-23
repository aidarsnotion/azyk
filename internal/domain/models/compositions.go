package models

type AminoAcidComposition struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	ProductID  uint    `json:"product_id"`
	ResearchID uint    `json:"research_id"`
	Name       float64 `json:"name"`
}

type ChemicalComposition struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	ProductID  uint    `json:"product_id"`
	ResearchID uint    `json:"research_id"`
	Name       float64 `json:"name"`
}

type MineralComposition struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	ProductID  uint    `json:"product_id"`
	ResearchID uint    `json:"research_id"`
	Name       float64 `json:"name"`
}
