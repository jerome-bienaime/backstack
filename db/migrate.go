package db

import (
	"web/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Warn(err error, tableName string) {
	log.WithFields(log.Fields{
		"error": err,
		"table": tableName,
	}).Warn("Could not migrate table")
}

func Info(model interface{}, tableName string) {
	log.WithFields(log.Fields{
		"table": tableName,
		"model": model,
	}).Info("Migrated table")
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		Warn(err, "users")
	}
	Info(&models.User{}, "users")
}
