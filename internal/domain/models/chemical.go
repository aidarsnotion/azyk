package models

// ChemicalComposition + перевод названия соединения
type ChemicalComposition struct {
	ID           int32                 `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID    int32                 `gorm:"index;not null" json:"product_id"`
	Product      Product               `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	ResearchID   int32                 `gorm:"index" json:"research_id"`
	Research     ResearchProject       `gorm:"foreignKey:ResearchID" json:"research,omitempty"`
	CompoundName string                `gorm:"size:255;not null" json:"compound_name"`
	Translations []ChemicalTranslation `gorm:"foreignKey:ChemicalID" json:"translations"`
	Quantity     float64               `json:"quantity"`
	Error        float64               `json:"error"`
	UnitID       int32                 `json:"unit_id"`
	Unit         UnitModel             `gorm:"foreignKey:UnitID" json:"unit,omitempty"`
	BaseModel
}

type ChemicalTranslation struct {
	ID         int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	ChemicalID int32  `gorm:"index;not null" json:"chemical_id"`
	Lang       string `gorm:"size:2;not null" json:"lang"`
	Name       string `gorm:"size:255;not null" json:"name"`
	BaseModel
}
