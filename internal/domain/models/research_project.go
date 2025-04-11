package models

import "time"

// ResearchProject с переводами
type ResearchProject struct {
	ID           int32                 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string                `gorm:"size:255;not null" json:"name"`
	Description  string                `json:"description"`
	Translations []ResearchTranslation `gorm:"foreignKey:ResearchID" json:"translations"`
	StartDate    time.Time             `json:"start_date"`
	EndDate      time.Time             `json:"end_date"`
	CreatedByID  int32                 `json:"created_by_id"`
	OrgID        int32                 `json:"org_id"`
	Products     []Product             `gorm:"many2many:product_research_projects" json:"products"`
	BaseModel
}

type ResearchTranslation struct {
	ID         int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	ResearchID int32  `gorm:"index;not null" json:"research_id"`
	Lang       string `gorm:"size:2;not null" json:"lang"`
	Name       string `gorm:"size:255" json:"name"`
	Desc       string `gorm:"type:text" json:"description"`
	BaseModel
}
