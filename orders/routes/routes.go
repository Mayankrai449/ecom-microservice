package routes

import (
	"github.com/Mayankrai449/ecom-microservice/orders/handlers"
	pb "github.com/Mayankrai449/ecom-microservice/orders/proto/inventory"
	"github.com/Mayankrai449/ecom-microservice/users/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, inventoryClient pb.InventoryServiceClient) {
	e.GET("/health", handlers.CheckHealth)

	orderHandler := handlers.NewOrderHandler(db, inventoryClient)

	api := e.Group("/api")
	api.Use(middleware.JWTMiddleware)

	api.POST("/orders", orderHandler.CreateOrder)
	api.GET("/orders", handlers.GetUserOrders)
}
