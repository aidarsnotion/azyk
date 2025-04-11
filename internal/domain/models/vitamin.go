package models

// VitaminComposition с переводами
type VitaminComposition struct {
	ID           int32                `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID    int32                `gorm:"index;not null" json:"product_id"`
	Product      Product              `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	VitaminName  string               `gorm:"size:255;not null" json:"vitamin_name"`
	VitaminGroup string               `gorm:"size:255" json:"vitamin_group"`
	Translations []VitaminTranslation `gorm:"foreignKey:VitaminID" json:"translations"`
	Quantity     float64              `json:"quantity"`
	Error        float64              `json:"error"`
	UnitID       int32                `json:"unit_id"`
	Unit         UnitModel            `gorm:"foreignKey:UnitID" json:"unit,omitempty"`
	BaseModel
}

type VitaminTranslation struct {
	ID        int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	VitaminID int32  `gorm:"index;not null" json:"vitamin_id"`
	Lang      string `gorm:"size:2;not null" json:"lang"`
	Name      string `gorm:"size:255;not null" json:"name"`
	BaseModel
}
