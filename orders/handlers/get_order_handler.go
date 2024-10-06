package handlers

import (
	"net/http"

	"github.com/Mayankrai449/ecom-microservice/orders/db"
	"github.com/Mayankrai449/ecom-microservice/orders/models"
	"github.com/Mayankrai449/ecom-microservice/users/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetUserOrders(c echo.Context) error {
	claims := c.Get("user").(*utils.JWTClaim)
	userID := claims.UserID

	var orders []models.Order

	if err := db.GetDB().Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		log.Errorf("Failed to fetch orders for user %d: %v", userID, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch orders")
	}

	log.Infof("Fetched %d orders for user %d", len(orders), userID)
	return c.JSON(http.StatusOK, orders)
}
