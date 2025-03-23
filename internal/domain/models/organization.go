package models

import "time"

// Организация
type Organization struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"size:255" json:"name"`
	Description   string    `json:"description"`
	Address       string    `json:"address"`
	Logo          string    `json:"logo"`
	EmployeeLimit int       `json:"employee_limit"` // Лимит сотрудников
	Employees     []User    `gorm:"foreignKey:OrgID" json:"employees"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
