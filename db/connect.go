package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

func Connect(dbFile string, config *gorm.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbFile), config)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Could not connect to database")
		panic(err)
	}
	return db
}
