package repositories

import (
	"github.com/Gabriel-Macedogmc/echo-golang/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func Create(data models.User) models.User {

	db.Create(&data)

	return data
}
