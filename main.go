package main

import (
	"github.com/Mayankrai449/ecom-microservice/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/health", handlers.CheckHealth)

	e.Logger.Fatal(e.Start(":8000"))

}
