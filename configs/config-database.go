package configs

import (
	"log"

	"github.com/Gabriel-Macedogmc/echo-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func ConnectToDatabase() (*gorm.DB, error) {
	dsn := "host=postgres-db-golang user=postgres password=postgres dbname=postgres port=5432"

	Instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Connected to Database...")

	return Instance, nil
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed...")
}
