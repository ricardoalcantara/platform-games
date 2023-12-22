package models

type TableType struct {
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255);not null"`
}

type CatalogTagTypeId uint8

const (
	CatalogTagTypeAction       CatalogTagTypeId = 1
	CatalogTagTypeAdventure    CatalogTagTypeId = 2
	CatalogTagTypeRPG          CatalogTagTypeId = 3
	CatalogTagTypeSimulation   CatalogTagTypeId = 4
	CatalogTagTypeStrategy     CatalogTagTypeId = 5
	CatalogTagTypeSports       CatalogTagTypeId = 6
	CatalogTagTypeFighting     CatalogTagTypeId = 7
	CatalogTagTypeHorror       CatalogTagTypeId = 8
	CatalogTagTypePuzzle       CatalogTagTypeId = 9
	CatalogTagTypeMusicRhythm  CatalogTagTypeId = 10
	CatalogTagTypeMMO          CatalogTagTypeId = 11
	CatalogTagTypeBattleRoyale CatalogTagTypeId = 12
	CatalogTagTypeSurvival     CatalogTagTypeId = 13
	CatalogTagTypeOpenWorld    CatalogTagTypeId = 14
	CatalogTagTypeEducational  CatalogTagTypeId = 15
	CatalogTagTypeSandbox      CatalogTagTypeId = 16
	CatalogTagTypeVisualNovel  CatalogTagTypeId = 17
)

type CatalogTagType struct {
	ID CatalogTagTypeId `gorm:"type:tinyint unsigned;primary_key"`
	TableType
}
