package models

import (
	"gorm.io/gorm"
)

type Catalog struct {
	gorm.Model

	Name             string `gorm:"type:varchar(255);not null"`
	Description      string `gorm:"type:text;not null"`
	Active           bool
	CatalogTagType   []CatalogTagType `gorm:"many2many:catalog_tags;"`
	CatalogVariation []CatalogVariation
	CatalogImage     []CatalogImage
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

func ListCatalog() ([]Catalog, error) {
	var err error
	var p []Catalog

	// if filter == nil {
	// 	filter = &VpsCatalogFilter{}
	// }

	err = DB.
		// Scopes(filter.GetScope).
		Preload("CatalogTagType").
		Preload("CatalogVariation").
		Preload("CatalogVariation.VariationDetail").
		Preload("CatalogImage").
		// Where("active = ?", true).
		Find(&p).Error

	if err != nil {
		return nil, err
	}

	return p, nil
}

func GetCatalog(id uint) (*Catalog, error) {
	var err error
	p := Catalog{}
	err = DB.
		Preload("CatalogTagType").
		Preload("CatalogVariation").
		Preload("CatalogVariation.VariationDetail").
		Preload("CatalogImage").
		Take(&p, "id = ?", id).
		Error

	if err != nil {
		return nil, err
	}

	return &p, nil
}
