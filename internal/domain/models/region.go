package models

// Region с переводами

type Region struct {
	ID           int32               `gorm:"primaryKey;autoIncrement" json:"id"`
	ParentID     *int32              `json:"parent_id"`
	Parent       *Region             `gorm:"foreignKey:ParentID" json:"parent"`
	Children     []Region            `gorm:"foreignKey:ParentID" json:"children"`
	Code         string              `gorm:"size:50" validate:"required" json:"code"`
	Name         string              `gorm:"size:255;not null" validate:"required" json:"name"`
	Translations []RegionTranslation `gorm:"foreignKey:RegionID" json:"translations"`
	Level        int                 `json:"level" validate:"required"`
	Categories   []Category          `gorm:"foreignKey:RegionsID" json:"categories"`
	BaseModel
}

type RegionTranslation struct {
	ID       int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	RegionID int32  `gorm:"index;not null" json:"region_id"`
	Lang     string `gorm:"size:2;not null" json:"lang"`
	Name     string `gorm:"size:255" json:"name"`
	BaseModel
}
