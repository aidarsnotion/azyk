package models

import "time"

type Product struct {
	ID           uint
	Category     string
	ResearchDate time.Time
	RegionID     uint
	// Другие поля, необходимые для продукта
}

type ProductTranslation struct {
	ID          uint
	ProductID   uint
	Language    string
	ProductName string
	Description string
}

// Аналогично определяются Region, RegionTranslation и прочие сущности...
