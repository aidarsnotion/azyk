package models

// UnitModel с переводами
type UnitModel struct {
	ID           int               `gorm:"primaryKey;autoIncrement" json:"id"`
	Code         string            `gorm:"unique;not null" json:"code"`
	Name         string            `gorm:"size:255;not null" json:"name"`
	Translations []UnitTranslation `gorm:"foreignKey:UnitID" json:"translations"`
}

type UnitTranslation struct {
	ID     int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	UnitID int32  `gorm:"index;not null" json:"unit_id"`
	Lang   string `gorm:"size:2;not null" json:"lang"`
	Name   string `gorm:"size:255" json:"name"`
	BaseModel
}
