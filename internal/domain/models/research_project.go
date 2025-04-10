package models

import "time"

// ResearchProject представляет исследовательский проект.
type ResearchProject struct {
	ID          int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"size:255" json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedByID int32     `json:"created_by_id"`
	OrgID       int32     `json:"org_id"` // внешний ключ на организацию
	Products    []Product `gorm:"many2many:product_research_projects;" json:"products"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
