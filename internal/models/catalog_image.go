package models

type CatalogImage struct {
	ID        uint `gorm:"primarykey"`
	CatalogId uint
	Url       string `gorm:"type:varchar(2048);not null"`

	Catalog *Catalog
}
