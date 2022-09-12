package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name      string    `gorm:"type:varchar(255);NOT NULL" json:"name" binding:"required"`
	Email     string    `gorm:"type:varchar(255)" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime"  json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
