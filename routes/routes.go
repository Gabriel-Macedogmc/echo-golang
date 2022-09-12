package routes

import (
	"github.com/Gabriel-Macedogmc/echo-golang/controllers"
	"github.com/Gabriel-Macedogmc/echo-golang/database/repositories"
	"github.com/Gabriel-Macedogmc/echo-golang/useCases"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, e *echo.Echo) {

	userRepository := repositories.NewUserRepository(db)

	userService := useCases.NewUserService(userRepository)

	userController := controllers.NewUserController(userService)

	e.POST("/users", userController.Create)
	e.GET("/", userController.GetAll)

}
