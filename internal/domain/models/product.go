package models

import "time"

type Product struct {
	ID              uint            `gorm:"primaryKey" json:"id"`
	Category        string          `gorm:"size:255" json:"category"`
	ResearchDate    time.Time       `json:"research_date"`
	RegionID        uint            `json:"region_id"`
	ResearchId      uint            `json:"research_id"`
	ResearchProject ResearchProject `gorm:"foreignKey:ResearchId;references:ID" json:"research_project"`
}

type ProductTranslation struct {
	ID          uint
	ProductID   uint
	Language    string
	ProductName string
	Description string
}
