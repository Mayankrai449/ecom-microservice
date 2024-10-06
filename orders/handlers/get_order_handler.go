package handlers

import (
	"net/http"

	"github.com/Mayankrai449/ecom-microservice/orders/db"
	"github.com/Mayankrai449/ecom-microservice/orders/models"
	"github.com/Mayankrai449/ecom-microservice/users/utils"
	"github.com/labstack/echo/v4"
)

func GetUserOrders(c echo.Context) error {
	claims := c.Get("user").(*utils.JWTClaim)
	userID := claims.UserID

	var orders []models.Order

	if err := db.GetDB().Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch orders")
	}

	if len(orders) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "No orders found for the user")
	}

	return c.JSON(http.StatusOK, orders)
}
