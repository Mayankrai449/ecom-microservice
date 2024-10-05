package routes

import (
	"github.com/Mayankrai449/ecom-microservice/users/handlers"
	"github.com/Mayankrai449/ecom-microservice/users/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/health", handlers.CheckHealth)

	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)

	api := e.Group("/api")
	api.Use(middleware.JWTMiddleware)

	api.GET("/user", handlers.GetUserProfile)
	api.PUT("/user", handlers.UpdateUser)
	api.DELETE("/user", handlers.DeleteUser)

}
