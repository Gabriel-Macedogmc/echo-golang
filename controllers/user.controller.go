package controllers

import (
	"net/http"

	"github.com/Gabriel-Macedogmc/echo-golang/database"
	"github.com/Gabriel-Macedogmc/echo-golang/database/repositories"
	"github.com/Gabriel-Macedogmc/echo-golang/models"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	user := models.User{}

	c.Bind(&user)

	//database.Instance.Create(&user)

	//useCases.CreateUserService(user)

	repositories.Create(user)

	return c.JSON(http.StatusCreated, user)
}

func GetAll(c echo.Context) error {
	users := []models.User{}
	database.Instance.Find(&users)

	return c.JSON(http.StatusOK, users)
}
