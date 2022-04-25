package main

import (
	"github.com/Manusiabodoh4/Twisocode_Backend/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.CORS())

	routes.NewRoutesAccount(app.Group("/v1/test")).NewCreateRoutes()

	app.Logger.Fatal(app.Start(":4567"))

}
