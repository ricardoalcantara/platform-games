package models

type CatalogVariation struct {
	ID         uint `gorm:"primarykey"`
	CatalogId  uint
	Name       string  `gorm:"type:varchar(255);not null"`
	Price      float64 `gorm:"type:decimal(15,2);not null"`
	Active     bool    `gorm:"not null;"`
	Memory     int     `gorm:"type:int;not null"`
	DiskSize   int     `gorm:"type:int;not null"`
	MaxPlayers int     `gorm:"type:int;not null"`

	Catalog         *Catalog
	VariationDetail []VariationDetail
}
