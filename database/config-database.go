package database

import (
	"log"

	"github.com/Gabriel-Macedogmc/echo-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func ConnectToDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432"

	Instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Connected to Database...")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed...")
}
