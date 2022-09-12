package repositories

import (
	"github.com/Gabriel-Macedogmc/echo-golang/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	Create(models.User) (models.User, error)
	GetAll() ([]models.User, error)
	FindByEmail(email string) (models.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

func (u userRepository) FindByEmail(email string) (user models.User, err error) {
	err = u.DB.Table("users").Where("email = ?", email).Error
	return user, err
}

func (u userRepository) Create(data models.User) (models.User, error) {
	data.ID = uuid.NewV4()
	err := u.DB.Create(&data).Error
	return data, err
}

func (u userRepository) GetAll() (users []models.User, err error) {
	err = u.DB.Find(&users).Error
	return users, err
}
