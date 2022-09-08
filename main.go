package main

import (
	"net/http"

	"github.com/Gabriel-Macedogmc/echo-golang/controllers"
	"github.com/Gabriel-Macedogmc/echo-golang/database"
	"github.com/labstack/echo/v4"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	e := echo.New()

	database.ConnectToDatabase()
	database.Migrate()

	e.POST("/users", controllers.Create)

	e.GET("/users/", controllers.GetAll)

	e.Logger.Fatal(e.Start(":3333"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")

	user := User{Id: "2", Name: "Gabriel"}

	if id == user.Id {
		return c.JSON(http.StatusOK, user)
	}
	return nil
}
