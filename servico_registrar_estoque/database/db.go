package database

import (
	"fmt"

	"github.com/TalesPalma/kafka_consume/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	sqliteDriver = "sqlite3"
	sqliteSource = "database.db"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	DB, err = gorm.Open(sqlite.Open(sqliteSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(&models.Product{})

	fmt.Println("Connected to database")
}
