package models

// FattyType — справочник типов жирных кислот с переводами
type FattyType struct {
	ID           int32                  `gorm:"primaryKey;autoIncrement" json:"id"`
	Code         string                 `gorm:"unique;not null" json:"code"` // saturated, monounsaturated...
	Translations []FattyTypeTranslation `gorm:"foreignKey:FattyTypeID" json:"translations"`
	BaseModel
}

type FattyTypeTranslation struct {
	ID          int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	FattyTypeID int32  `gorm:"index;not null" json:"fatty_type_id"`
	Lang        string `gorm:"size:2;not null" json:"lang"`
	Name        string `gorm:"size:255;not null" json:"name"`
	BaseModel
}

// FattyAcidComposition представляет состав жирных кислот в продукте.
type FattyAcidComposition struct {
	ID            int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID     int32     `gorm:"index;not null" json:"product_id"`
	Product       Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	FattyAcidName string    `gorm:"size:255;not null" json:"fatty_acid_name"`
	FattyTypeID   int32     `json:"fatty_type_id"`
	FattyType     FattyType `gorm:"foreignKey:FattyTypeID" json:"fatty_type"`
	Quantity      float64   `json:"quantity"`
	Error         float64   `json:"error"`
	UnitID        int32     `json:"unit_id"`
	Unit          UnitModel `gorm:"foreignKey:UnitID" json:"unit,omitempty"`
	BaseModel
}
