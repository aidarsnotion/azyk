package models

// Category представляет категорию продукта.
type Category struct {
	ID           int32                 `gorm:"primaryKey;autoIncrement" json:"id"`
	RegionsID    int32                 `json:"regions_id"`
	Region       Region                `gorm:"foreignKey:RegionsID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"region"`
	Translations []CategoryTranslation `gorm:"foreignKey:CategoryID" json:"translations"`
	BaseModel
}

// CategoryTranslation — переводы категории
type CategoryTranslation struct {
	ID         int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryID int32  `gorm:"index;not null" json:"category_id"`
	Lang       string `gorm:"size:2;not null" json:"lang"` // 'ru', 'en', 'kg'
	Name       string `gorm:"size:255;not null" json:"name"`
	BaseModel
}
