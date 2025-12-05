package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDatabase() {
	stringDB := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringDB))
	if err != nil {
		log.Fatal("failed to connect database")
	}

	log.Println("Database connection established")
}
