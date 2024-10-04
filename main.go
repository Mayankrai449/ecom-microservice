package main

import (
	"github.com/Mayankrai449/ecom-microservice/db"
	"github.com/Mayankrai449/ecom-microservice/models"
	"github.com/Mayankrai449/ecom-microservice/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	db.DB_Config()
	db.Migrate(&models.Product{})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))

}
