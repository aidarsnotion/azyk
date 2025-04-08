package models

import (
	"time"
)

// Region представляет регион или страну.
type Region struct {
	ID         int32      `gorm:"primaryKey;autoIncrement" json:"id"`
	ParentID   *int32     `json:"parent_id"` // указатель, так как может быть NULL
	Parent     *Region    `gorm:"foreignKey:ParentID" json:"parent"`
	Children   []Region   `gorm:"foreignKey:ParentID" json:"children"`
	Code       string     `json:"code"`
	Name       string     `json:"name"`
	Level      int        `json:"level"` // уровень иерархии (0 – страна, 1 – регион и т.д.)
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	Categories []Category `gorm:"foreignKey:RegionsID" json:"categories"`
}
