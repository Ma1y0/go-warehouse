package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnecToDatabase() {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to conncet to database")
	}

	// Migrate schemas
	db.AutoMigrate(&ItemModel{})

	DB = db
}
