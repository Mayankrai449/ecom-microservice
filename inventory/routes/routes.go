package routes

import (
	"github.com/Mayankrai449/ecom-microservice/inventory/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/health", handlers.CheckHealth)

	e.GET("/products", handlers.GetProducts)
	e.GET("/products/:id", handlers.GetProductByID)

	e.POST("/products", handlers.CreateProduct)
	e.PUT("/products/:id", handlers.UpdateProduct)
	e.DELETE("/products/:id", handlers.DeleteProduct)
}
