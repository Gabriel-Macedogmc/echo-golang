package main

import (
	"github.com/Gabriel-Macedogmc/echo-golang/configs"
	"github.com/Gabriel-Macedogmc/echo-golang/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, _ := configs.ConnectToDatabase()
	configs.Migrate()

	routes.SetupRoutes(db, e)

	e.Logger.Fatal(e.Start(":3333"))
}
