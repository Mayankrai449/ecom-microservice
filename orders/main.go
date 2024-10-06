package main

import (
	"github.com/Mayankrai449/ecom-microservice/orders/db"
	"github.com/Mayankrai449/ecom-microservice/orders/models"
	"github.com/Mayankrai449/ecom-microservice/orders/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	db.DB_Config()
	db.Migrate(&models.Order{})

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))

}
