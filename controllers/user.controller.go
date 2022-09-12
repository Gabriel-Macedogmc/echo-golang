package controllers

import (
	"net/http"

	"github.com/Gabriel-Macedogmc/echo-golang/configs"
	"github.com/Gabriel-Macedogmc/echo-golang/models"
	"github.com/Gabriel-Macedogmc/echo-golang/useCases"
	"github.com/labstack/echo/v4"
)

type UserController interface {
	Create(c echo.Context) error
	GetAll(c echo.Context) error
}

type userController struct {
	userService useCases.UserService
}

func NewUserController(u useCases.UserService) UserController {
	return userController{
		userService: u,
	}
}

func (u userController) Create(c echo.Context) error {
	user := models.User{}

	c.Bind(&user)

	userCreated, err := u.userService.Save(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, userCreated)
}

func (u userController) GetAll(c echo.Context) error {
	users := []models.User{}
	configs.Instance.Find(&users)

	return c.JSON(http.StatusOK, users)
}
