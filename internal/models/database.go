package models

import (
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
	// DB.AutoMigrate(&DataCenterIp{})
}

func createTypes() {
	// if err := DB.Take(&VpsCatalogType{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
	// 	var err error

	// 	err = DB.Transaction(func(tx *gorm.DB) error {
	// 		for _, vt := range []VpsCatalogType{
	// 			{ID: VpsCatalogTypeShared, TableType: TableType{Name: "Shared", Description: "VPS com vCPU compartilhado"}},
	// 			{ID: VpsCatalogTypeCPUOptimized, TableType: TableType{Name: "CPU Otimizado", Description: "VPS com vCPU dedicado para uso maior de processamento"}},
	// 			{ID: VpsCatalogTypeMemoryOptimized, TableType: TableType{Name: "CPU Optimized", Description: "VPS com Memória otimizada e vCPU dedicada para uso maior de memória"}},
	// 			{ID: VpsCatalogTypeGeneral, TableType: TableType{Name: "Recurso Dedicado", Description: "VPS com vCPU dedicada e mais memória para uso geral"}},
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
