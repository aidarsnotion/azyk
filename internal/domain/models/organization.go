package models

// Организация с переводами
type Organization struct {
	ID            int32                     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string                    `gorm:"size:255;not null" json:"name"`
	Description   string                    `json:"description"`
	Translations  []OrganizationTranslation `gorm:"foreignKey:OrgID" json:"translations"`
	Address       string                    `json:"address"`
	Logo          string                    `json:"logo"`
	EmployeeLimit int                       `json:"employee_limit"`
	Employees     []User                    `gorm:"foreignKey:OrgID" json:"employees,omitempty"`
	BaseModel
}

type OrganizationTranslation struct {
	ID    int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	OrgID int32  `gorm:"index;not null" json:"org_id"`
	Lang  string `gorm:"size:2;not null" json:"lang"`
	Name  string `gorm:"size:255" json:"name"`
	Desc  string `gorm:"type:text" json:"description"`
	BaseModel
}
