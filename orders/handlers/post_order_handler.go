package handlers

import (
	"net/http"

	"github.com/Mayankrai449/ecom-microservice/orders/db"
	"github.com/Mayankrai449/ecom-microservice/orders/models"
	"github.com/Mayankrai449/ecom-microservice/users/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func PostOrder(c echo.Context) error {
	claims := c.Get("user").(*utils.JWTClaim)
	userID := claims.UserID

	order := new(models.Order)
	if err := c.Bind(order); err != nil {
		log.Errorf("Failed to bind request: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	order.UserID = userID
	order.OrderStatus = models.StatusPending

	if err := c.Validate(order); err != nil {
		log.Errorf("Validation error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := db.GetDB().Create(order).Error; err != nil {
		log.Errorf("Failed to save order: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create order")
	}

	log.Infof("Order created successfully for user %d", userID)
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Order created successfully",
		"order":   order,
	})
}
