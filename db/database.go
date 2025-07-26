package db

import (
	"log"
	"scd-service/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("jobs.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	err = DB.AutoMigrate(&models.Job{})
	if err != nil {
		log.Fatalf("failed to migrate db: %v", err)
	}
}
