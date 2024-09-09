package database

import (
	"log"

	"github.com/TalesPalma/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DB_NAME = "database.db"
)

var (
	DB *gorm.DB
)

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while connecting to database", err)
	}
	db.AutoMigrate(&models.Product{})

	DB = db

	log.Println("Database connected")
}
