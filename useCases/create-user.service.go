package useCases

import (
	"github.com/Gabriel-Macedogmc/echo-golang/database/repositories"
	"github.com/Gabriel-Macedogmc/echo-golang/models"
)

type UserService interface {
	Save(models.User) (models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) UserService {
	return userService{
		userRepository: r,
	}
}

func (u userService) Save(user models.User) (models.User, error) {

	return u.userRepository.Create(user)
}
