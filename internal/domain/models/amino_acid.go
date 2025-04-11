package models

// AminoAcidComposition + перевод названия аминокислоты
type AminoAcidComposition struct {
	ID            int32                  `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID     int32                  `gorm:"index;not null" json:"product_id"`
	Product       Product                `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"product,omitempty"`
	ResearchID    int32                  `gorm:"index" json:"research_id"`
	Research      ResearchProject        `gorm:"foreignKey:ResearchID" json:"research,omitempty"`
	AminoAcidName string                 `gorm:"size:255;not null" json:"amino_acid_name"`
	Translations  []AminoAcidTranslation `gorm:"foreignKey:AminoAcidID" json:"translations"`
	Quantity      float64                `json:"quantity"`
	Error         float64                `json:"error"`
	UnitID        int32                  `json:"unit_id"`
	Unit          UnitModel              `gorm:"foreignKey:UnitID" json:"unit,omitempty"`
	BaseModel
}

type AminoAcidTranslation struct {
	ID          int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	AminoAcidID int32  `gorm:"index;not null" json:"amino_acid_id"`
	Lang        string `gorm:"size:2;not null" json:"lang"`
	Name        string `gorm:"size:255;not null" json:"name"`
	BaseModel
}
