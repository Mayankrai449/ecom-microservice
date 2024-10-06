package handlers

import (
	"context"
	"net/http"

	pb "github.com/Mayankrai449/ecom-microservice/orders/proto/inventory"

	"github.com/Mayankrai449/ecom-microservice/orders/models"
	"github.com/Mayankrai449/ecom-microservice/users/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type OrderHandler struct {
	db              *gorm.DB
	inventoryClient pb.InventoryServiceClient
}

func NewOrderHandler(db *gorm.DB, inventoryClient pb.InventoryServiceClient) *OrderHandler {
	return &OrderHandler{
		db:              db,
		inventoryClient: inventoryClient,
	}
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	claims := c.Get("user").(*utils.JWTClaim)
	userID := claims.UserID

	var order models.Order
	if err := c.Bind(&order); err != nil {
		log.Error("Error binding order:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	productIDs := make([]uint32, len(order.Products))
	quantities := make([]uint32, len(order.Products))
	for i, product := range order.Products {
		productIDs[i] = uint32(product.ProductID)
		quantities[i] = uint32(product.Quantity)
	}

	stockResponse, err := h.inventoryClient.CheckStock(context.Background(), &pb.StockRequest{
		ProductIds: productIDs,
		Quantities: quantities,
	})
	if err != nil {
		log.Error("Error checking stock:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to check stock"})
	}

	if !stockResponse.InStock {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": stockResponse.Message})
	}

	order.UserID = userID
	order.OrderStatus = models.StatusPending
	if err := h.db.Create(&order).Error; err != nil {
		log.Error("Error creating order:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create order"})
	}

	return c.JSON(http.StatusCreated, order)
}
