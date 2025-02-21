package models

import "time"

// Исследовательский проект
type ResearchProject struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:255" json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedByID int       `json:"created_by_id"` // Кто создал проект
	OrgID       int       `json:"org_id"`        // Ссылка на организацию
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
