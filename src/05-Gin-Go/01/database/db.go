package database

import (
	"gin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	ConectString := "host=localhost user=root password=root dbname=root sslmode=disable"
	DB, err := gorm.Open(postgres.Open(ConectString), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to database!")
	}

	DB.AutoMigrate(&models.Student{})
}
