package models

// Category представляет категорию продукта.
type Category struct {
	ID             int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	NameOfCategory string `json:"name_of_category"` // наименование категории (например, "Мясные")
	RegionsID      int32  `json:"regions_id"`       // внешний ключ на регион
	Region         Region `gorm:"foreignKey:RegionsID" json:"region"`
}
