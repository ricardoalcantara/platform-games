package models

import (
	"gorm.io/gorm"
)

type CatalogImage struct {
	ID        uint `gorm:"primarykey"`
	CatalogId uint
	Url       string

	Catalog *Catalog
}

type Catalog struct {
	gorm.Model

	Name        string
	Description string
	Tags        []string

	CatalogVariation []CatalogVariation
	CatalogImage     []CatalogImage
}

type CatalogVariation struct {
	ID        uint `gorm:"primarykey"`
	CatalogId uint
	Name      string
	Price     float64 `gorm:"type:decimal(15,2);not null"`
	Active    bool    `gorm:"not null;"`

	Memory     int `gorm:"type:int;not null"`
	DiskSize   int `gorm:"type:int;not null"`
	MaxPlayers int `gorm:"type:int;not null"`

	Catalog         *Catalog
	VariationDetail []VariationDetail
}

type VariationDetail struct {
	CatalogVariationId uint

	Key    string
	Value  string
	Global bool

	CatalogVariation *CatalogVariation
}

func (p *Catalog) Save(tx *gorm.DB) error {
	if tx == nil {
		tx = DB
	}

	err := tx.Create(&p).Error
	if err != nil {
		return err
	}

	return nil
}

func ListActiveVpsCatalog() ([]Catalog, error) {
	var err error
	var p []Catalog

	// if filter == nil {
	// 	filter = &VpsCatalogFilter{}
	// }

	err = DB.
		// Scopes(filter.GetScope).
		Where("active = ?", true).
		Order("vps_catalog_type_id, price").
		Find(&p).Error

	if err != nil {
		return nil, err
	}

	return p, nil
}

func GetVpsCatalog(id uint) (*Catalog, error) {
	var err error
	p := Catalog{}
	err = DB.Take(&p, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &p, nil
}
