package database

import (
	"log"
	"time"

	"github.com/evertonbzr/api-golang/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(uri string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database", "error", err)
	}

	configSQLDriver, err := db.DB()

	if err != nil {
		log.Fatal("Error getting SQL driver", "error", err)
	}

	configSQLDriver.SetMaxIdleConns(10)
	configSQLDriver.SetMaxOpenConns(50)
	configSQLDriver.SetConnMaxIdleTime(30 * time.Minute)
	configSQLDriver.SetConnMaxLifetime(time.Hour)

	if config.IsDevelopment() {
		db = db.Debug()
	}

	return db
}
