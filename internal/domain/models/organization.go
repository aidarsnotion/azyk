package models

import "time"

// Organization представляет организацию.
type Organization struct {
	ID            int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string    `gorm:"size:255" json:"name"`
	Description   string    `json:"description"`
	Address       string    `json:"address"`
	Logo          string    `json:"logo"`
	EmployeeLimit int       `json:"employee_limit"` // лимит сотрудников
	Employees     []User    `gorm:"foreignKey:OrgID" json:"employees"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
