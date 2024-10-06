package routes

import (
	"github.com/Mayankrai449/ecom-microservice/orders/handlers"
	"github.com/Mayankrai449/ecom-microservice/users/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/health", handlers.CheckHealth)

	api := e.Group("/api")
	api.Use(middleware.JWTMiddleware)

	api.POST("/orders", handlers.PostOrder)

}
