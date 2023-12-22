package models

type VariationDetail struct {
	CatalogVariationId uint
	Name               string `gorm:"type:varchar(255);not null"`
	Description        string `gorm:"type:varchar(255);not null"`

	CatalogVariation *CatalogVariation
}
