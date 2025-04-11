package models

// MineralComposition представляет состав минералов в продукте.
type MineralComposition struct {
	ID          int32           `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID   int32           `gorm:"index;not null" json:"product_id"`
	Product     Product         `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	ResearchID  int32           `gorm:"index" json:"research_id"`
	Research    ResearchProject `gorm:"foreignKey:ResearchID" json:"research,omitempty"`
	MineralName string          `gorm:"size:255;not null" json:"mineral_name"`
	Quantity    float64         `json:"quantity"`
	Error       float64         `json:"error"`
	UnitID      int32           `json:"unit_id"`
	Unit        UnitModel       `gorm:"foreignKey:UnitID" json:"unit,omitempty"`
	BaseModel
}
