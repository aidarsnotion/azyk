package models

// AminoAcidComposition представляет состав аминокислот в продукте.
type AminoAcidComposition struct {
	ID            int32   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID     int32   `json:"product_id"`      // внешний ключ на Product
	ResearchID    int32   `json:"research_id"`     // внешний ключ на исследовательский проект
	AminoAcidName string  `json:"amino_acid_name"` // название аминокислоты
	Quantity      float64 `json:"quantity"`        // количество
	Error         float64 `json:"error"`           // погрешность измерения
	UnitID        int32   `json:"unit_id"`         // внешний ключ на справочник единиц измерения (если используется)
}

// ChemicalComposition представляет состав химических соединений в продукте.
type ChemicalComposition struct {
	ID           int32   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID    int32   `json:"product_id"`
	ResearchID   int32   `json:"research_id"`
	CompoundName string  `json:"compound_name"` // название вещества
	Quantity     float64 `json:"quantity"`      // количество вещества
	Error        float64 `json:"error"`         // погрешность измерения
	UnitID       int32   `json:"unit_id"`       // внешний ключ на таблицу Unit, если необходимо
}

// MineralComposition представляет состав минералов в продукте.
type MineralComposition struct {
	ID          int32   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID   int32   `json:"product_id"`
	ResearchID  int32   `json:"research_id"`
	MineralName string  `json:"mineral_name"` // например, "Calcium", "Iron" и т.д.
	Quantity    float64 `json:"quantity"`
	Error       float64 `json:"error"`
	UnitID      int32   `json:"unit_id"`
}

// FattyAcidComposition представляет состав жирных кислот в продукте.
type FattyAcidComposition struct {
	ID        int32 `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID int32 `gorm:"not null" json:"product_id"`
	// Для отношений можно использовать также поле Product (если требуется загрузка связанных данных)
	// Product         Product `gorm:"foreignKey:ProductID"`
	FattyAcidName   string  `json:"fatty_acid_name"`    // название жирной кислоты
	TypeOfFattyAcid string  `json:"type_of_fatty_acid"` // тип жирной кислоты
	Quantity        float64 `json:"quantity"`
	Error           float64 `json:"error"`
	UnitID          int32   `json:"unit_id"` // внешний ключ на таблицу Unit
}

// VitaminComposition представляет состав витаминов в продукте.
type VitaminComposition struct {
	ID        int32 `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID int32 `json:"product_id"` // заменили ProductsID на ProductID для согласованности
	// При необходимости можно добавить поле Product, связанное по foreignKey
	// Product      Product   `gorm:"foreignKey:ProductID"`
	VitaminName  string  `json:"vitamin_name"`  // название витамина
	VitaminGroup string  `json:"vitamin_group"` // группа витаминов (например, жирорастворимые или водорастворимые)
	Quantity     float64 `json:"quantity"`
	Error        float64 `json:"error"`
	UnitID       int32   `json:"unit_id"` // внешний ключ на справочник Unit
}
