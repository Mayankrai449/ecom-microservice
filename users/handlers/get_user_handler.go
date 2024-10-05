package handlers

import (
	"net/http"

	"github.com/Mayankrai449/ecom-microservice/users/db"
	"github.com/Mayankrai449/ecom-microservice/users/models"
	"github.com/Mayankrai449/ecom-microservice/users/utils"
	"github.com/labstack/echo/v4"
)

func GetUserProfile(c echo.Context) error {
	claims := c.Get("user").(*utils.JWTClaim)

	var user models.User
	if err := db.GetDB().First(&user, claims.UserID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User Not Found!")
	}

	user.Password = ""
	return c.JSON(http.StatusOK, user)
}
