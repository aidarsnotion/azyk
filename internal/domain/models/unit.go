package models

// UnitModel представляет справочник единиц измерения.
type UnitModel struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Code string `gorm:"unique;not null"` // Например, "g", "mg" и т.п.
	Name string // Полное название (например, "gram", "milligram")
}
