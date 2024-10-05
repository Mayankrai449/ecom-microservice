package handlers

import (
	"net/http"

	"github.com/Mayankrai449/ecom-microservice/users/db"
	"github.com/Mayankrai449/ecom-microservice/users/models"
	"github.com/Mayankrai449/ecom-microservice/users/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func UpdateUser(c echo.Context) error {
	claims := c.Get("user").(*utils.JWTClaim)

	req := new(models.UpdateUserRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user models.User
	if err := db.GetDB().First(&user, claims.UserID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User Not Found!")
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to Hash the Password")
		}
		user.Password = string(hashedPassword)
	}

	if err := db.GetDB().Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to Update the User!")
	}

	user.Password = ""
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	claims := c.Get("user").(*utils.JWTClaim)

	var user models.User
	if err := db.GetDB().First(&user, claims.UserID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User Not Found!")
	}

	if err := db.GetDB().Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to Delete the User!")
	}

	return c.JSON(http.StatusAccepted, map[string]string{
		"message": "User Deleted Successfully"})
}
