package lib

import (
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	ORM *gorm.DB
}

func NewDatabase(config Config) Database {
	pconfig := postgres.Config{
		DSN: config.Database.DSN(),
	}

	db, err := gorm.Open(postgres.New(pconfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("[Database] Error open database[%s]: %s", pconfig.DSN, err)
	}

	return Database{
		ORM: db,
	}
}
