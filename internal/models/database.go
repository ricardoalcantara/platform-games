package models

import (
	"errors"
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	db_url := os.Getenv("DB_URL")
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               db_url,
		DefaultStringSize: 256,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("connection error")
	} else {
		log.Debug().Msg("Db Connected")
	}

	if value, ok := os.LookupEnv("AUTO_MIGRATE"); ok && value == "true" {
		migrate()
		createTypes()
	}
}

func migrate() {
	DB.AutoMigrate(&CatalogTagType{})
	DB.AutoMigrate(&Catalog{})
	DB.AutoMigrate(&CatalogImage{})
	DB.AutoMigrate(&CatalogVariation{})
	DB.AutoMigrate(&VariationDetail{})
}

func createTypes() {
	if err := DB.Take(&CatalogTagType{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		var err error

		err = DB.Transaction(func(tx *gorm.DB) error {
			for _, vt := range []CatalogTagType{
				{ID: CatalogTagTypeAction, TableType: TableType{Name: "Action", Description: "Games that focus on physical challenges, including hand-eye coordination and reaction time."}},
				{ID: CatalogTagTypeAdventure, TableType: TableType{Name: "Adventure", Description: "Story-driven games where players explore virtual worlds, solve puzzles, and interact with characters."}},
				{ID: CatalogTagTypeRPG, TableType: TableType{Name: "Role-Playing Game (RPG)", Description: "Games where players take on the roles of characters in a fictional setting, often involving character development, decision-making, and narrative."}},
				{ID: CatalogTagTypeSimulation, TableType: TableType{Name: "Simulation", Description: "Games that replicate real-world activities, such as driving, flying, or life simulation."}},
				{ID: CatalogTagTypeStrategy, TableType: TableType{Name: "Strategy", Description: "Games that require planning, resource management, and tactical decision-making."}},
				{ID: CatalogTagTypeSports, TableType: TableType{Name: "Sports", Description: "Simulations of real-world sports, including soccer, basketball, and racing."}},
				{ID: CatalogTagTypeFighting, TableType: TableType{Name: "Fighting", Description: "Games where players engage in hand-to-hand combat with opponents."}},
				{ID: CatalogTagTypeHorror, TableType: TableType{Name: "Horror", Description: "Games designed to create a sense of fear and suspense."}},
				{ID: CatalogTagTypePuzzle, TableType: TableType{Name: "Puzzle", Description: "Games that challenge players with logic and problem-solving tasks."}},
				{ID: CatalogTagTypeMusicRhythm, TableType: TableType{Name: "Music/Rhythm", Description: "Games where players interact with music or rhythm-based challenges."}},
				{ID: CatalogTagTypeMMO, TableType: TableType{Name: "Massively Multiplayer Online (MMO)", Description: "Online games that support large numbers of players interacting in a persistent virtual world."}},
				{ID: CatalogTagTypeBattleRoyale, TableType: TableType{Name: "Battle Royale", Description: "Games where a large number of players compete to be the last person or team standing."}},
				{ID: CatalogTagTypeSurvival, TableType: TableType{Name: "Survival", Description: "Games where players must survive in a hostile environment, often with limited resources."}},
				{ID: CatalogTagTypeOpenWorld, TableType: TableType{Name: "Open World", Description: "Games with a vast, seamless game world that players can explore freely."}},
				{ID: CatalogTagTypeEducational, TableType: TableType{Name: "Educational", Description: "Games designed to teach or reinforce educational concepts."}},
				{ID: CatalogTagTypeSandbox, TableType: TableType{Name: "Sandbox", Description: "Games that provide a virtual sandbox for players to create, modify, or interact with the game world."}},
				{ID: CatalogTagTypeVisualNovel, TableType: TableType{Name: "Visual Novel", Description: "Interactive narratives often with static images and text."}},
			} {
				if err = DB.Create(&vt).Error; err != nil {
					return err
				}
			}

			return nil
		})

		if err != nil {
			panic(err)
		}
	}

	// if err := DB.Take(&VpsAddonsType{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
	// 	var err error

	// 	err = DB.Transaction(func(tx *gorm.DB) error {
	// 		for _, at := range []VpsAddonsType{
	// 			{ID: VpsAddonsTypeExtraDiskSpace, TableType: TableType{Name: "Espaço em disco adicional", Description: "Adicione espaço adicional no disco atual da sua VPS"}},
	// 		} {
	// 			if err = DB.Create(&at).Error; err != nil {
	// 				panic(err)
	// 			}
	// 		}

	// 		return nil
	// 	})

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// if err := DB.Take(&UserVpsStatus{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
	// 	var err error

	// 	err = DB.Transaction(func(tx *gorm.DB) error {
	// 		for ID, Name := range UserVpsStatusMap {
	// 			if err = DB.Create(&UserVpsStatus{ID, Name}).Error; err != nil {
	// 				panic(err)
	// 			}
	// 		}

	// 		return nil
	// 	})

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// if err := DB.Take(&VpsPlanType{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
	// 	var err error

	// 	err = DB.Transaction(func(tx *gorm.DB) error {
	// 		for _, vt := range []VpsPlanType{
	// 			{ID: VpsPlanTypePrePaid, TableType: TableType{Name: "Pré pago", Description: "Plano que você consome os créditos enquanto mantém sua VPS criada"}},
	// 			{ID: VpsPlanTypePostPaid, TableType: TableType{Name: "Pós pago", Description: "Plano que você paga no final do mês o que foi consumido pela sua VPS criada"}},
	// 		} {
	// 			if err = DB.Create(&vt).Error; err != nil {
	// 				return err
	// 			}
	// 		}

	// 		return nil
	// 	})

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// if err := DB.Take(&DataCenterLoginType{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
	// 	var err error

	// 	err = DB.Transaction(func(tx *gorm.DB) error {
	// 		for _, vt := range []DataCenterLoginType{
	// 			{ID: DataCenterLoginTypeUserAndPassword, TableType: TableType{Name: "UserAndPassword", Description: "UserAndPassword"}},
	// 			{ID: DataCenterLoginTypeTokenAndSecret, TableType: TableType{Name: "TokenAndSecret", Description: "TokenAndSecret"}},
	// 		} {
	// 			if err = DB.Create(&vt).Error; err != nil {
	// 				return err
	// 			}
	// 		}

	// 		return nil
	// 	})

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
}
